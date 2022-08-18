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
    is_order_amendment = is_order_amendment?(sf_order)

    # if the original order, then it will have been contracted if it has additional phases/amendments
    # if it hasn't been contracted, then we know there's no amendments to look up
    if !is_order_amendment && !sf_order[SF_ORDER_CONTRACTED]
      log.info 'order is not contracted, assuming only initial order'
      return ContractStructure.new(initial: sf_order)
    end

    # if not new, then it's an amendment. Remember each account can have different Types, which is why we don't do an exaustive type cehck here
    if is_order_amendment
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
    order_amendments = order_amendments_for_contract(sf_contract_id)

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

    stripe_customer = translate_account(sf_account)

    sf_order_items = order_lines_from_order(sf_order)
    invoice_items, subscription_items = phase_items_from_order_lines(sf_order_items)
    is_recurring_order = !subscription_items.empty?

    if !is_recurring_order
      return create_stripe_invoice_from_order(stripe_customer, invoice_items, sf_order)
    end

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

    # this should never happen, but we are still learning about CPQ
    # if metered billing, quantity is not set, so we set to 1
    if OrderHelpers.extract_all_items_from_subscription_schedule(stripe_transaction).map {|l| l[:quantity] || 1 }.any?(&:zero?)
      report_edge_case("quantity is zero on initial subscription schedule")
    end

    # https://jira.corp.stripe.com/browse/PLATINT-1731
    days_until_due = stripe_transaction.[](:default_settings)&.[](:invoice_settings)&.[](:days_until_due)
    if days_until_due
      stripe_transaction.default_settings.invoice_settings.days_until_due = OrderHelpers.transform_payment_terms_to_days_until_due(days_until_due)
    end

    # TODO the idempotency key here is not perfect, need to refactor and use a job UID or something
    stripe_transaction = Stripe::SubscriptionSchedule.create(
      stripe_transaction.to_hash,
      @user.stripe_credentials.merge(idempotency_key: sf_order[SF_ID])
    )

    log.info 'stripe subscription or invoice created', stripe_resource_id: stripe_transaction.id

    update_sf_stripe_id(sf_order, stripe_transaction)

    if stripe_transaction.is_a?(Stripe::SubscriptionSchedule)
      PriceHelpers.auto_archive_prices_on_subscription_schedule(@user, stripe_transaction)
    end

    stripe_transaction
  end

  sig do
    params(
      previous_phase_items: T::Array[ContractItemStructure],
      order_amendment: Restforce::SObject
    ).returns([T::Array[ContractItemStructure], T::Array[ContractItemStructure]])
  end
  def build_phase_items_from_order_amendment(previous_phase_items, order_amendment)
    sf_order_items = order_lines_from_order(order_amendment)
    invoice_items_in_order, subscription_items_in_order = phase_items_from_order_lines(sf_order_items)

    aggregate_phase_items = merge_subscription_line_items(
      previous_phase_items,
      subscription_items_in_order
    )

    is_recurring_order = !aggregate_phase_items.empty?

    if !is_recurring_order
      raise Integrations::Errors::UnhandledEdgeCase.new("order amendments representing all one-time invoices")
    end

    # TODO this should be moved to a helper
    is_order_terminated = aggregate_phase_items.all?(&:fully_terminated?)

    if is_order_terminated && !invoice_items_in_order.empty?
      raise Integrations::Errors::UnhandledEdgeCase.new("one-time invoice items but terminated order")
    end

    [invoice_items_in_order, aggregate_phase_items]
  end

  sig { params(contract_structure: ContractStructure).void }
  def update_subscription_phases_from_order_amendments(contract_structure)
    return if contract_structure.amendments.empty?

    # refresh to include subscription reference on the Stripe ID field in SF if the order was just translated
    contract_structure.initial.refresh

    # TODO use application sync record
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

    # at this point, the initial order would have already been translated
    # and a corresponding subscription schedule created.

    subscription_schedule = T.cast(subscription_schedule, Stripe::SubscriptionSchedule)

    # Order amendments contain a negative item if they are adjusting a previous line item.
    # If they are adjusting a previous line item

    # How we will know if the line items are the same? Price references won't work.
    # The price of a line item can change across amendments. There's a special field
    # on an amendmended line item that we can leverage: `SBQQ__RevisedOrderProduct__c`.
    # However, this field always references the *first* order line, not the last revised
    # order line.

    # do NOT use the current state of the subscription schedule at all. This could cause some issues:
    # if the old SF data was mutated in some way (even metadata!) that data will be used, which could
    # cause weird unintended side effects. However, the state of the Stripe side would be too hard to infer.
    # the user could mutate the phase data out-of-band, so we need to (a) limit the phase we are editing
    # (b) recreate each phase data from the raw order line data.

    sf_initial_order_items = order_lines_from_order(contract_structure.initial)
    _, aggregate_phase_items = phase_items_from_order_lines(sf_initial_order_items)

    subscription_phases = subscription_schedule.phases

    # SF does not enforce mutation restrictions. It's possible to go in and modify anything you want in Salesforce
    # for this reason we should NOT mutate the existing phases of a subscription. This could result in updating quantity, metadata, etc
    # that was updated in Salesforce changing phase data that the user does not expect to be changed.

    # TODO should break this out into a separate method and wrap in `catch_errors_with_salesforce_context`
    contract_structure.amendments.each_with_index do |sf_order_amendment, index|
      log.info 'processing amendment', salesforce_object: sf_order_amendment, index: index

      invoice_items_in_order, aggregate_phase_items = build_phase_items_from_order_amendment(aggregate_phase_items, sf_order_amendment)
      is_order_terminated = aggregate_phase_items.all?(&:fully_terminated?)

      if is_order_terminated && contract_structure.amendments.size - 1 != index
        raise Integrations::Errors::UnhandledEdgeCase.new("order terminated, but there's more amendments")
      end

      # this loop excludes the initial phase, which is why we are subtracting by 1
      if subscription_phases.count - 1 > index
        log.info 'phase already exists, checking for diff and skipping'
        # TODO check for diff and log
        next
      end

      # TODO price ID dup check on the invoice items
      aggregate_phase_items = OrderHelpers.ensure_unique_phase_item_prices(@user, aggregate_phase_items)

      # TODO should probably use a completely different key/mapping for the phase items
      phase_params = extract_salesforce_params!(sf_order_amendment, Stripe::SubscriptionSchedule)

      phase_params['start_date'] = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(phase_params['start_date'])

      # TODO check for float value
      subscription_term_from_sales_force = phase_params['iterations'].to_i



      # if the order is terminated this is not used
      phase_params['iterations'] = transform_iterations_by_billing_frequency(
        subscription_term_from_sales_force,
        T.must(aggregate_phase_items.first).stripe_params[:price]
      )

      aggregate_phase_items = OrderHelpers.remove_terminated_lines(aggregate_phase_items)

      new_phase = Stripe::StripeObject.construct_from({
        add_invoice_items: invoice_items_in_order.map(&:stripe_params),

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

      subscription_phases = OrderHelpers.sanitize_subscription_schedule_phase_params(subscription_phases)

      # NOTE intentional decision here NOT to update any other subscription fields
      catch_errors_with_salesforce_context(secondary: sf_order_amendment) do
        if is_subscription_schedule_cancelled
          # TODO should we add additional metadata here?
          log.info 'cancelling subscription immediately'

          subscription_schedule.cancel(
            invoice_now: false,
            prorate: false
          )
        else
          log.info 'adding phase', sf_order_amendment_id: sf_order_amendment.Id

          # TODO wrap in error context
          subscription_schedule.proration_behavior = 'none'
          subscription_schedule.phases = subscription_phases
          subscription_schedule.save
        end
      end

      update_sf_stripe_id(sf_order_amendment, subscription_schedule)
    end

    PriceHelpers.auto_archive_prices_on_subscription_schedule(@user, subscription_schedule)
  end

  sig do
    params(
      original_aggregate_phase_items: T::Array[StripeForce::Translate::ContractItemStructure],
      new_phase_items: T::Array[StripeForce::Translate::ContractItemStructure]
    ).returns(T::Array[StripeForce::Translate::ContractItemStructure])
  end
  def merge_subscription_line_items(original_aggregate_phase_items, new_phase_items)
    # avoid mutating the input value
    aggregate_phase_items = original_aggregate_phase_items.dup

    termination_lines, additive_lines = new_phase_items.partition(&:termination?)

    additive_lines.each do |new_subscription_item|
      log.info 'adding new line item'
      aggregate_phase_items << new_subscription_item
    end

    aggregate_phase_items = terminate_subscription_line_items(aggregate_phase_items, termination_lines)
    aggregate_phase_items
  end

  sig do
    params(
      original_aggregate_phase_items: T::Array[StripeForce::Translate::ContractItemStructure],
      termination_lines: T::Array[StripeForce::Translate::ContractItemStructure]
    ).returns(T::Array[StripeForce::Translate::ContractItemStructure])
  end
  def terminate_subscription_line_items(original_aggregate_phase_items, termination_lines)
    aggregate_phase_items = original_aggregate_phase_items.dup
    revision_map = T.let({}, T::Hash[String, T::Array[ContractItemStructure]])

    aggregate_phase_items.select(&:new_order_line?).each do |origin_order_line|
      if revision_map[origin_order_line.order_line_id].present?
        raise Integrations::Errors::ImpossibleState.new("should never be more than a single revised order ID match")
      end

      revision_map[origin_order_line.order_line_id] ||= [origin_order_line]
    end

    # discover one-to-many mapping between line items
    aggregate_phase_items.
      # only include revisions
      reject(&:new_order_line?)
      .each do |revision_order_line|
        origin_order_line_id = T.must(revision_order_line.revised_order_line_id)
        revised_item_list = T.must(revision_map[origin_order_line_id])

        if revised_item_list.map(&:order_line_id).include?(revision_order_line.order_line_id)
          raise ArgumentError.new("an order line ID should not be listed twice")
        else
          revised_item_list << revision_order_line
        end
      end

    log.debug "order amendment revision map", revision_map: revision_map

    # now let's terminate the related line items
    termination_lines.each do |termination_line|
      origin_order_line_id = T.must(termination_line.revised_order_line_id)
      fifo_order_line_stack = revision_map[origin_order_line_id]

      if !fifo_order_line_stack || fifo_order_line_stack.empty?
        throw_user_failure!(
          salesforce_object: T.must(termination_line.order_line),
          message: "Any order items, revising order items in a previous order, must not be skipped in the previous order." \
                   " Order line with ID '#{termination_line.revised_order_line_id}' could not be found."
        )
      end

      if termination_line.quantity > 0
        raise "quantity on termination lines should never be positive"
      end

      termination_line.quantity.abs.times do
        fifo_remaining_line = fifo_order_line_stack
          .reject(&:fully_terminated?)
          .reverse
          .first

        if fifo_remaining_line
          log.info 'reducing quantity on line', reducing_line: fifo_remaining_line.order_line_id
          fifo_remaining_line.reduce_quantity
        else
          # this should never happen
          raise StripeForce::Errors::RawUserError.new("termination quantity is greater than the aggregate quantity for the line item")
        end
      end
    end

    aggregate_phase_items
  end

  # retrieves all line items, filters out skipped order lines
  sig { params(sf_order: Restforce::SObject).returns(Array) }
  def order_lines_from_order(sf_order)
    # we could add our custom fields to the SOQL query, but then we wouldn't get the logging below
    # TODO use cache_service
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

      phase_item_struct = ContractItemStructure.from_order_line_and_params(
        sf_order_item,
        phase_item.to_hash,
      )

      # quantity cannot be specified if usage type is metered
      if PriceHelpers.metered_price?(price)
        # allowing > 1 quantities may cause issues with terminations and generally indicates a data issue on the customers end
        if phase_item_struct.quantity > 1
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
      if PriceHelpers.recurring_price?(price)
        subscription_items << phase_item_struct
      else
        # TODO this sanitization should be moved somewhere else
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

  # TODO move to order amendment helpers
  # if an order does not have a 'AmendedContract' relationship than it is a initial order
  sig { params(sf_order: Restforce::SObject).returns(T::Boolean) }
  def is_order_amendment?(sf_order)
    order_with_amended_contract_query = sf.query(
      # include `Type`, `OpportunityId` for debugging purposes
      <<~EOL
        SELECT Type, OpportunityId,
               Opportunity.SBQQ__AmendedContract__c
        FROM #{SF_ORDER}
        WHERE Id = '#{sf_order.Id}'
      EOL
    )

    if order_with_amended_contract_query.size.zero?
      raise Integrations::Errors::ImpossibleInternalError.new("query should never return an empty result")
    end

    if order_with_amended_contract_query.size > 1
      raise Integrations::Errors::ImpossibleInternalError.new("query should only return a single result")
    end

    order_with_amended_contract = order_with_amended_contract_query.first

    amended_contract_id = order_with_amended_contract.dig(SF_OPPORTUNITY, "SBQQ__AmendedContract__c")
    is_order_amendment = amended_contract_id.present?

    if !OrderTypeOptions.values.map(&:serialize).include?(order_with_amended_contract.Type)
      log.warn 'order type is not standard', order_type: order_with_amended_contract.Type
    end

    if is_order_amendment && order_with_amended_contract.Type == OrderTypeOptions::NEW.serialize
      report_edge_case("order is determined to be an amendment, but type is new")
    end

    is_order_amendment
  end

  # if we are given an order amendment, determine the original/initial/non-amendment order
  # there is no direct connection between the amendment and the initial order, so we must run
  # these dispirate queries to create a connection.
  def extract_initial_order_from_amendment(sf_order_amendment)
    # in the case of an amendment, the associated opportunity contains a reference to the contract
    # which contains a reference to the original quote, which references the orginal order (initial non-amendment order)
    initial_quote_query = sf.query(
      # include `OpportunityId`, `SBQQ__AmendedContract__c` for debugging purposes
      <<~EOL
        SELECT OpportunityId, Opportunity.SBQQ__AmendedContract__c,
               Opportunity.SBQQ__AmendedContract__r.#{SF_CONTRACT_QUOTE_ID}
        FROM #{SF_ORDER}
        WHERE Id = '#{sf_order_amendment.Id}'
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

    if sf_original_quote_id.blank?
      log.warn "related records while retrieve missing quote",
        opportunity_id: initial_quote_query.first.dig("OpportunityId"),
        amended_contract: initial_quote_query.first.dig(SF_OPPORTUNITY, "SBQQ__AmendedContract__c")

      raise StripeForce::Errors::RawUserError.new("Could not find initial quote associated with order amendment")
    end

    log.info 'quote tied to initial order found', quote_id: sf_original_quote_id

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

  # TODO this should be a function in a helper module
  sig { params(sf_contract_id: String).returns(Array) }
  def order_amendments_for_contract(sf_contract_id)
    # TODO we are hardcoding the subscription start date here as the ordering key, we should pull this dynamically from the mapper
    # important for results to be ordered with subscription date ascending'
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
    order_amendment_query.map {|i| sf.find(SF_ORDER, i[SF_ID]) }
  end
end
