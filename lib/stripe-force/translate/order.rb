# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_order(sf_object)
    contract_structure = extract_contract_from_order(sf_object)

    create_stripe_transaction_from_sf_order(contract_structure.initial)

    update_subscription_phases_from_order_amendments(contract_structure)
  end

  # give an order (amendment or initial order) determine the initial order and all amendments
  sig { params(sf_order: T.untyped).returns(ContractStructure) }
  def extract_contract_from_order(sf_order)
    # in a fresh CPQ account there is only NEW and AMENDMENT types, but users can and do customize these types
    # for our purposes, we only care about new and not-new (amendment types) so we don't need to filter these aggressively
    if ![OrderTypeOptions::AMENDMENT, OrderTypeOptions::NEW].map(&:serialize).include?(sf_order.Type)
      report_edge_case('unexpected order type', metadata: {type: sf_order.Type})
    end

    # if the original order, then it will have been contracted if it has additional phases/amendments
    # if it hasn't been contracted, then we know there's no amendments to look up
    if sf_order.Type == OrderTypeOptions::NEW.serialize && !sf_order[SF_ORDER_CONTRACTED]
      log.info 'order is not contracted, assuming only initial order'
      return ContractStructure.new(initial: sf_order)
    end

    # if not new, then it's an amendment. Remember each account can have different Types, which is why we don't do an exaustive type cehck here
    if sf_order.Type != OrderTypeOptions::NEW.serialize
      initial_order = extract_initial_order_from_amendment(sf_order)
    else
      log.info 'initial order provided'
      initial_order = sf_order
    end

    # we ignore `sf_order` at this point onward: we have the initial order and extract all amendments

    log.info 'initial order, finding amendments', initial_order_id: initial_order.Id

    # at this point, we have a reference to the initial order and know it has been contracted

    # if the initial order has been contracted then we need to find all amended orders.
    # All order amendments contain a reference to the contract in the opportunity associated
    # the new order. This seems to be the only common reference between orders.

    # first, let's find the contract related to this order
    # the order field on the contract *could* contain a reference, but is overwritten if an
    # order amendment is contracted. However, the quote reference on the contract is NOT overwritten.
    # It will always be the first quote that generated the contract.

    sf_contract_id = extract_contract_id_from_initial_order(initial_order)

    if !sf_contract_id
      return ContractStructure.new(initial: initial_order)
    end

    log.info 'contract for order amendment found', contract_id: sf_contract_id

    # TODO we are hardcoding the subscription start date here as the ordering key, we should pull this dynamically from the mapper
    # important for results to be ordered with subscription date ascending
    order_amendment_query = sf.query(
      <<~EOL
        SELECT Id FROM #{SF_ORDER}
        WHERE Opportunity.SBQQ__AmendedContract__c = '#{sf_contract_id}'
        ORDER BY SBQQ__Quote__r.#{CPQ_QUOTE_SUBSCRIPTION_START_DATE} ASC
      EOL
    )

    log.info 'order amendments found',
      count: order_amendment_query.count,
      amendmend_ids: order_amendment_query.map(&:Id)

    # TODO use cache_service
    order_amendments = order_amendment_query.map {|i| sf.find(SF_ORDER, i[SF_ID]) }

    ContractStructure.new(
      initial: initial_order,
      amendments: order_amendments
    )
  end

  # TODO How can we organize code to support CPQ & non-CPQ use-cases? how can this be abstracted away from the order?
  def create_stripe_transaction_from_sf_order(sf_order)
    log.info 'translating order', salesforce_object: sf_order

    if sf_order.Type != OrderTypeOptions::NEW.serialize
      raise Integrations::Errors::ImpossibleState.new("only new orders should be passed for transaction generation #{sf_order.Id}")
    end

    stripe_transaction = retrieve_from_stripe(Stripe::SubscriptionSchedule, sf_order)
    stripe_transaction ||= retrieve_from_stripe(Stripe::Invoice, sf_order)

    if !stripe_transaction.nil?
      log.info 'order already translated',
        stripe_transaction_id: stripe_transaction.id,
        salesforce_order_id: sf_order.Id
      return
    end

    # TODO use cache_service
    sf_account = sf.find(SF_ACCOUNT, sf_order[SF_ORDER_ACCOUNT])

    stripe_customer = create_customer_from_sf_account(sf_account)

    sf_order_items = order_lines_from_order(sf_order)
    invoice_items, subscription_items = phase_items_from_order_lines(sf_order_items)
    is_recurring_order = !subscription_items.empty?

    order_update_params = {}

    if is_recurring_order
      log.info 'recurring items found, creating subscription schedule'

      subscription_params = extract_salesforce_params!(sf_order, Stripe::SubscriptionSchedule)

      # TODO should file an API papercut for this
      # when creating the subscription schedule the start_date must be specified on the heaer
      # when updating it, it is specified on the individual phase object
      subscription_start_date = subscription_params['start_date']
      subscription_params['start_date'] = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(subscription_start_date)

      # TODO this should really be done *before* generating the line items and therefore creating prices
      phase_iterations = transform_iterations_by_billing_frequency(
        # TODO is the restforce gem somehow formatting everything as a float? Or is this is the real value returned from SF?
        subscription_params.delete('iterations').to_i,
        T.must(subscription_items.first).stripe_params[:price]
      )

      # TODO we should remove all zero-quantity items
      # TODO we should have a check to ensure all quantities are positive
      # TODO we should check if subscription is backddated, if it is, then we should only proceed if the user has a specific flag enabled

      # initial order, so there is only a single phase
      initial_phase = {
        add_invoice_items: invoice_items.map(&:stripe_params),
        items: subscription_items.map(&:stripe_params),
        iterations: phase_iterations,

        metadata: stripe_metadata_for_sf_object(sf_order),
      }

      # TODO this needs to be gated and synced with the specific flag that CF is using
      if subscription_start_date >= DateTime.now
        # when `sub_sched_backdating_anchors_on_backdate` is enabled prorating the initial phase
        # does NOT actually prorate it, but instead bills the full amount of the subscription item
        # and locks the billing anchor to the start date of the subscription. Without this flag the
        # billing achor is disconnected from the start date of the invoice
        initial_phase[:proration_behavior] = 'none'
      end

      # TODO add mapping support against the subscription schedule phase

      # TODO subs in SF must always have an end date
      stripe_transaction = Stripe::SubscriptionSchedule.construct_from({
        customer: stripe_customer.id,

        # TODO this should be specified in the defaults hash... we should create a defaults hash https://jira.corp.stripe.com/browse/PLATINT-1501
        end_behavior: 'cancel',

        # initial order will only ever contain a single phase
        phases: [initial_phase],

        metadata: stripe_metadata_for_sf_object(sf_order),
      })

      mapper.assign_values_from_hash(stripe_transaction, subscription_params)
      apply_mapping(stripe_transaction, sf_order)

      # TODO the idempotency key here is not perfect, need to refactor and use a job UID or something

      catch_stripe_api_errors(sf_order) do
        stripe_transaction = Stripe::SubscriptionSchedule.create(
          stripe_transaction.to_hash,
          @user.stripe_credentials.merge(idempotency_key: sf_order[SF_ID])
        )
      end
    else
      log.info 'no recurring items found, creating a one-time invoice'

      # TODO there has got to be a way to include the lines on the invoice item create call
      invoice_items.each do |invoice_item_params|
        # TODO we should wrap these in `catch_stripe_api_errors`
        # TODO idempotency keys https://jira.corp.stripe.com/browse/PLATINT-1474
        Stripe::InvoiceItem.create(
          {customer: stripe_customer}.merge(invoice_item_params.stripe_params),
          @user.stripe_credentials
        )
      end

      stripe_transaction = create_stripe_object(Stripe::Invoice, sf_order, additional_stripe_params: {
        customer: stripe_customer.id,
      })

      stripe_transaction.finalize_invoice

      order_update_params[prefixed_stripe_field(ORDER_INVOICE_PAYMENT_LINK)] = stripe_transaction.hosted_invoice_url
    end

    log.info 'stripe subscription or invoice created', stripe_resource_id: stripe_transaction.id

    update_sf_stripe_id(sf_order, stripe_transaction, additional_salesforce_updates: order_update_params)

    stripe_transaction
  end

  sig { params(contract_structure: ContractStructure).void }
  def update_subscription_phases_from_order_amendments(contract_structure)
    return if contract_structure.amendments.empty?

    # refresh to include subscription reference on the Stripe ID field in SF if the order was just translated
    contract_structure.initial.refresh

    # TODO should pull from a sync record in the future
    subscription_schedule = retrieve_from_stripe(
      Stripe::SubscriptionSchedule,
      contract_structure.initial
    )

    if !subscription_schedule
      raise Integrations::Errors::ImpossibleState.new("initial order should always be present")
    end

    subscription_schedule = T.cast(subscription_schedule, Stripe::SubscriptionSchedule)

    # TODO we should short-circuit our logic here sooner, but for now let's just track this in the logs sooner
    if subscription_schedule.status == "canceled"
      report_edge_case('subscription is cancelled, it cannot be modified')
    end

    subscription_schedule = T.cast(subscription_schedule, Stripe::SubscriptionSchedule)

    # at this point, the initial order would have already been translated
    # and a corresponding subscription schedule created.

    # Order amendments contain a negative item if they are adjusting a previous line item.
    # In order to determine what the line items should be for a given phase we need to aggregate
    # all previous line items together (i.e. the lines contained in the initial order and all order amendments).

    # How we will know if the line items are the same? Price references won't work.
    # The price of a line item can change across amendments. There's a special field
    # on an amendmended line item that we can leverage: `SBQQ__RevisedOrderProduct__c`

    last_phase = T.must(subscription_schedule.phases.last)

    # use the current state of the subscription schedule as the starting point for all aggregate line items
    # TODO there is some risk here that if we don't do the calculation properly right away
    # TODO should we conver the `items` stripe objects into hashes? Probably more consistent this way.

    # NOTE `to_hash` is a special method on the stripe object which handles nested objects and is NOT the same as `to_h`
    aggregate_phase_items = last_phase.items.map(&:to_hash).map {|h| ContractItemStructure.new_from_created_phase_item(h) }
    subscription_phases = subscription_schedule.phases

    is_subscription_schedule_cancelled = T.let(false, T::Boolean)

    # SF does not enforce mutation restrictions. It's possible to go in and modify anything you want in Salesforce
    # for this reason we should NOT mutate the existing phases of a subscription. This could result in updating quantity, metadata, etc
    # that was updated in Salesforce changing phase data that the user does not expect to be changed.

    contract_structure.amendments.each_with_index do |sf_order_amendment, index|
      sf_order_items = order_lines_from_order(sf_order_amendment)
      invoice_items_for_order_item, subscription_items_for_order_item = phase_items_from_order_lines(sf_order_items)

      aggregate_phase_items = merge_subscription_line_items(
        aggregate_phase_items,
        subscription_items_for_order_item
      )

      is_recurring_order = !aggregate_phase_items.empty?

      if !is_recurring_order
        raise Integrations::Errors::UnhandledEdgeCase.new("order amendments representing all one-time invoices")
      end

      # TODO this probably needs to be tweaked after item merging is working
      is_order_terminated = aggregate_phase_items.all?(&:is_terminated?)

      if is_order_terminated && !invoice_items_for_order_item.empty?
        raise Integrations::Errors::UnhandledEdgeCase.new("one-time invoice items but terminated order")
      end

      if is_order_terminated && contract_structure.amendments.size - 1 != index
        raise Integrations::Errors::UnhandledEdgeCase.new("order terminated, but there's more amendmends")
      end

      # if it's not terminated, it could be partially terminated
      if !is_order_terminated
        aggregate_phase_items.reject! do |phase_item|
          if phase_item.is_terminated?
            log.info 'line iterm terminated', terminated_order_item_id: phase_item.order_line&.Id
            true
          else
            false
          end
        end
      end

      # TODO should probably use a completely different key/mapping for the phase items
      phase_params = extract_salesforce_params!(sf_order_amendment, Stripe::SubscriptionSchedule)

      phase_params['start_date'] = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(phase_params['start_date'])
      phase_params['iterations'] = transform_iterations_by_billing_frequency(
        phase_params['iterations'].to_i,
        T.must(aggregate_phase_items.first).stripe_params[:price]
      )

      # this loop excludes the initial phase, which is why we are subtracting by 1
      if subscription_phases.count - 1 > index
        log.info 'phase already exists, checking for diff and skipping'
        # TODO check for diff and log
        next
      end

      new_phase = Stripe::StripeObject.construct_from({
        add_invoice_items: invoice_items_for_order_item.map(&:stripe_params),

        # this is important, otherwise multiple phase changes in a single job run will use the same aggregate phase items
        items: aggregate_phase_items.deep_dup.map(&:stripe_params),

        # TODO should be moved to global defaults
        proration_behavior: 'none',

        metadata: stripe_metadata_for_sf_object(sf_order_amendment),
      }.merge(phase_params))

      previous_phase = T.must(subscription_phases[index])

      # TODO dynamic metadata on the phase?

      # TODO this is very naive... something better here?
      # align date boundaries of the schedules
      previous_phase.end_date = new_phase.start_date

      # TODO I wonder if we can do something smarter here: if the invoice has not been paid/billed, do XYZ?
      # this is a special case: subscription is cancelled on the same day, the intention here is to not bill the user at all
      is_subscription_schedule_cancelled = is_order_terminated &&
        previous_phase.start_date == previous_phase.end_date &&
        contract_structure.amendments.count == 1

      # if the order is terminated, updating the last phase end date and NOT adding another phase is all that needs to be done
      if !is_order_terminated
        subscription_phases << new_phase
      end

      # TODO report https://jira.corp.stripe.com/browse/PLATINT-1479
      # You can't pass back the phase in it's original format, it must be modified to avoid:
      # 'You passed an empty string for 'phases[0][collection_method]'. We assume empty values are an attempt to unset a parameter; however 'phases[0][collection_method]' cannot be unset. You should remove 'phases[0][collection_method]' from your request or supply a non-empty value.'
      subscription_phases.each do |phase|
        phase
          .keys
          # all fields that are nil from the API should be removed before sending to the API
          .select {|field_sym| phase.send(field_sym).nil? }
          .each do |field_sym|
            Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
              phase,
              field_sym
            )
          end
      end

      # NOTE intentional decision here NOT to update any other subscription fields
      if is_subscription_schedule_cancelled
        # TODO should we add additional metadata here?
        log.info 'cancelling subscription immediately'

        catch_stripe_api_errors(sf_order_amendment) do
          subscription_schedule.cancel(
            invoice_now: false,
            prorate: false
          )
        end
      else
        log.info 'adding phase', sf_order_amendment_id: sf_order_amendment.Id

        # TODO wrap in error context
        subscription_schedule.proration_behavior = 'none'
        subscription_schedule.phases = subscription_phases

        catch_stripe_api_errors(sf_order_amendment) do
          subscription_schedule.save
        end
      end

      update_sf_stripe_id(sf_order_amendment, subscription_schedule)
    end
  end

  # TODO will most likely throw this out, but we'll need to model something like this for terminations
  sig do
    params(
      aggregate_phase_items: T::Array[StripeForce::Translate::ContractItemStructure],
      new_phase_items: T::Array[StripeForce::Translate::ContractItemStructure]
    ).returns(T::Array[StripeForce::Translate::ContractItemStructure])
  end
  def merge_subscription_line_items(aggregate_phase_items, new_phase_items)
    # avoid mutating the input value
    aggregate_phase_items = aggregate_phase_items.dup

    # now we have the stripe representation of all of the line items on the order
    new_phase_items.each do |new_subscription_item|
      # subscription_item entries are still in our internal phase item structure at this point
      # they get converted into StripeObjects downstream
      existing_phase_item = if new_subscription_item.original_order_line_id.present?
        aggregate_phase_items.detect do |possible_existing_item|
          # TODO I don't like relying on the metadata here; maybe we could regenerate the line items for the first order and use that representation instead?
          #      or maybe in the future we could use an internal sync record to pull references? Either way, this should change.
          possible_existing_item.stripe_params[:metadata][:salesforce_order_item_id] == new_subscription_item.original_order_line_id
        end
      else
        log.info 'line is not revising a previous line item'
        nil
      end

      if existing_phase_item.nil? && new_subscription_item.original_order_line_id
        throw_user_failure!(
          salesforce_object: T.must(new_subscription_item.order_line),
          message: "Any order items, revising order items in a previous order, must not be skipped in the previous order." \
                   " Order line with ID '#{new_subscription_item.original_order_line_id}' could not be found."
        )
      end

      if !existing_phase_item.nil?
        new_subscription_item.append_previous_phase_item(existing_phase_item)

        # delete the old item from the phase item list, the new `subscription_item` will supercede it
        deleted_count = aggregate_phase_items.delete_if {|i| i.object_id == existing_phase_item.object_id }.count

        log.info 'removed old phase items', count: deleted_count
      end

      aggregate_phase_items << new_subscription_item
    end

    aggregate_phase_items
  end

  # retrieves all line items, filters out skipped order lines
  def order_lines_from_order(sf_order)
    # TODO should move to cache service
    # TODO could include the order line skip field and skip lines before pulling them
    sf_order_items = sf.query(
      <<~EOL
        SELECT Id
        FROM #{SF_ORDER_ITEM}
        WHERE OrderID = '#{sf_order.Id}'
      EOL
    ).map(&:Id).map do |order_line_id|
      sf.find(SF_ORDER_ITEM, order_line_id)
    end

    sf_order_items.select do |sf_order_item|
      # never expect this to occur
      if sf_order_item.IsDeleted || !sf_order_item.SBQQ__Activated__c
        report_edge_case("order line is deleted or not activated")
      end

      should_keep = sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)].nil? ||
        !sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)]

      if !should_keep
        log.info 'order line marked as skipped'
      end

      should_keep
    end
  end

  sig do
    params(sf_order_lines: T::Array[Restforce::SObject])
    .returns([T::Array[ContractItemStructure], T::Array[ContractItemStructure]])
  end
  def phase_items_from_order_lines(sf_order_lines)
    invoice_items = []
    subscription_items = []

    sf_order_lines.map do |sf_order_item|
      price = create_price_for_order_item(sf_order_item)

      # could occur if a coupon is required for a negative amount, although this should probably be built into the `price` method instead
      next if price.nil?

      # a phase item is NOT a subscription item, but they are close, and right now the data mapper & revelant keys uses these across the codebase
      # I wonder if they will eventually converge in the public Stripe API over time?
      phase_item = Stripe::SubscriptionItem.construct_from({
        price: price.id,
        metadata: stripe_metadata_for_sf_object(sf_order_item),
      })

      phase_item_params = extract_salesforce_params!(sf_order_item, Stripe::SubscriptionItem)
      mapper.assign_values_from_hash(phase_item, phase_item_params)
      apply_mapping(phase_item, sf_order_item)

      # TODO add test case for this
      unless Integrations::Utilities::StripeUtil.is_integer_value?(phase_item.quantity)
        throw_user_failure!(
          salesforce_object: sf_order_item,
          message: "Quantity specified as a decimal value. Only integers are supported."
        )
      end

      # we know the quantity is not a float, so we can force it to an integer
      phase_item.quantity = phase_item.quantity.to_i

      phase_item_struct = ContractItemStructure.new(
        # TODO maybe consider including a boolean for this instead?
        # is_recurring: recurring_item?(sf_order_item),

        stripe_params: phase_item.to_hash,
        order_line: sf_order_item,

        original_order_line_id: sf_order_item[SF_ORDER_ITEM_REVISED_ORDER_PRODUCT],
        quantity: phase_item[:quantity]
      )

      # quantity cannot be specified if usage type is metered
      if PriceHelpers.metered_price?(price)
        # allowing > 1 quantities may cause issues with terminations and generally indicates a data issue on the customers end
        if phase_item_struct.stripe_params[:quantity] > 1
          throw_user_failure!(
            salesforce_object: sf_order_item,
            message: "Order lines with a price configured for metered billing cannot have a quantity greater than 1."
          )
        end

        # setting quantity to null generates an empty string when `to_hash` is called
        # which will throw an API error when this is passed to Stripe, which is we need this hack
        phase_item_struct.stripe_params.delete(:quantity)
      end

      # TODO this helper uses a hardcoded field to determine if the line is recurring or not; it should be dynamic instead
      if recurring_item?(sf_order_item)
        subscription_items << phase_item_struct
      else
        # TODO metadata is currently not supported here https://jira.corp.stripe.com/browse/PLATINT-1609
        phase_item_struct.stripe_params.delete(:metadata)

        invoice_items << phase_item_struct
      end
    end

    [invoice_items, subscription_items]
  end

  sig { params(sf_initial_order: Restforce::SObject).returns(T.nilable(String)) }
  def extract_contract_id_from_initial_order(sf_initial_order)
    # if this occurs, the user's CPQ is not configured properly/as we assume
    if sf_initial_order[SF_ORDER_QUOTE].blank?
      raise Integrations::Errors::ImpossibleState.new("no quote associated with orde1r")
    end

    contract_query = sf.query(
      <<~EOL
        SELECT #{SF_ID}
        FROM #{SF_CONTRACT}
        WHERE #{SF_CONTRACT_QUOTE_ID} = '#{sf_initial_order[SF_ORDER_QUOTE]}'
      EOL
    )

    if contract_query.size > 1
      raise Integrations::Errors::ImpossibleState.new("more than one contract associated with order")
    end

    # this can occur if contracts are processed async
    if contract_query.count.zero?
      log.info 'order is contracted, but no contract is associated'
      return nil
    end

    contract_query.first[SF_ID]
  end

  def extract_initial_order_from_amendment(sf_order)
    # in the case of an amendment, the associated opportunity contains a reference to the contract
    # which contains a reference to the original quote, which references the orginal order (initial non-amendment order)
    initial_quote_query = sf.query(
      <<~EOL
        SELECT Opportunity.SBQQ__AmendedContract__r.#{SF_CONTRACT_QUOTE_ID}
        FROM #{SF_ORDER}
        WHERE Id = '#{sf_order.Id}'
      EOL
    )

    if initial_quote_query.count.zero?
      raise Integrations::Errors::ImpossibleState.new("order amendments should always be associated with the initial quote")
    end

    # TODO this should never happen and should be removed
    if initial_quote_query.count > 1
      raise Integrations::Errors::ImpossibleState.new("exact ID match yields two records")
    end

    # the contract tied to the amended order has the ID of the quote of the original order
    # let's get that ID, then we can pull the order tied to that original quote
    sf_original_quote_id = initial_quote_query.first.dig(
      SF_OPPORTUNITY,
      # TODO should pull into a constant
      "SBQQ__AmendedContract__r",
      SF_CONTRACT_QUOTE_ID
    )

    log.info 'initial quote found', quote_id: sf_original_quote_id

    initial_order_query = sf.query(
      <<~EOL
        SELECT Id
        FROM #{SF_ORDER}
        WHERE SBQQ__Quote__c = '#{sf_original_quote_id}'
      EOL
    )

    if initial_order_query.count.zero?
      raise Integrations::Errors::ImpossibleState.new("initial order should be associated with an initial quote")
    end

    # TODO this should never happen and should be removed
    if initial_order_query.count > 1
      raise Integrations::Errors::ImpossibleState.new("exact ID match yields two records")
    end

    initial_order_id = initial_order_query.first[SF_ID]

    log.info 'found initial order', initial_order_id: initial_order_id

    sf.find(SF_ORDER, initial_order_id)
  end
end
