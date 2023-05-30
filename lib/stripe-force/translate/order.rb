# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  include StripeForce::Constants

  def translate_order(sf_object)
    locker.lock_salesforce_record(sf_object)

    # get initial order and all amendments
    contract_structure = extract_contract_from_order(sf_object)

    # in case the amendment order was synced but the initial order was not ready to be synced
    if !orders_respect_custom_filters([contract_structure.initial])
      log.info 'initial order does not respect users custom order filters. not syncing: ', order_ids: contract_structure.initial
      raise StripeForce::Errors::RawUserError.new("Attempted to sync amendment order when initial order was skipped because it didn't match custom sync filters")
    end

    # process the initial order
    create_stripe_transaction_from_sf_order(contract_structure.initial)

    # if there are amendment orders for this initial order, ensure the order respects
    # the user's custom order filters before syncing
    if !orders_respect_custom_filters(contract_structure.amendments)
      log.info 'not all amendment orders respect users custom order filters. not syncing orders: ', order_ids: contract_structure.amendments
      return
    end

    # after that is created, then process all amendments
    update_subscription_phases_from_order_amendments(contract_structure)
  end

  # we need to ensure that both the initial order and the amendment order respect the user's custom order filters
  # regardless of which one is picked up by the connector
  sig { params(sf_orders: T::Array[Restforce::SObject]).returns(T::Boolean) }
  def orders_respect_custom_filters(sf_orders)
    sf_orders.each do |sf_order|
      sf_order_id = sf_order.Id
      user_specified_order_filters = @user.user_specified_where_clause_for_object(SF_ORDER)
      results = backoff { @user.sf_client.query("SELECT Id FROM #{SF_ORDER} WHERE Id = '#{sf_order_id}' #{user_specified_order_filters}") }

      # if one of the order amendments does not meet the custom filters, do not sync over all of them
      if results.empty?
        return false
      end
    end
    true
  end

  # give an order (amendment or initial order) determine the initial order and all amendments
  sig { params(sf_order: T.untyped).returns(ContractStructure) }
  def extract_contract_from_order(sf_order)
    is_order_amendment = StripeForce::Translate::OrderHelpers.is_order_amendment?(@user, sf_order)

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

  def create_stripe_transaction_from_sf_order(sf_order)
    log.info 'translating order', salesforce_object: sf_order

    # this check is rigorously enforced upstream
    if sf_order.Type != OrderTypeOptions::NEW.serialize
      log.warn "order is not new, but is treated as such, type field could be customized"
    end

    stripe_transaction = retrieve_from_stripe(Stripe::SubscriptionSchedule, sf_order)
    stripe_transaction ||= retrieve_from_stripe(Stripe::Invoice, sf_order)

    if !stripe_transaction.nil?
      log.info 'initial order is already translated',
        stripe_transaction_id: stripe_transaction.id,
        salesforce_order_id: sf_order.Id
      return
    end

    sf_account = cache_service.get_record_from_cache(SF_ACCOUNT, sf_order[SF_ORDER_ACCOUNT])
    stripe_customer = translate_account(sf_account)

    sf_order_items = order_lines_from_order(sf_order)
    invoice_items, subscription_items = phase_items_from_order_lines(sf_order_items)

    is_recurring_order = !subscription_items.empty?
    if !is_recurring_order
      return create_stripe_invoice_from_order(stripe_customer, invoice_items, sf_order)
    end

    log.info 'recurring items found, creating subscription schedule'

    subscription_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order, Stripe::SubscriptionSchedule)

    subscription_schedule = Stripe::SubscriptionSchedule.construct_from({
      # TODO this should be specified in the defaults hash... we should create a defaults hash https://jira.corp.stripe.com/browse/PLATINT-1501
      end_behavior: 'cancel',
      metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order),
    })

    mapper.assign_values_from_hash(subscription_schedule, subscription_params)
    apply_mapping(subscription_schedule, sf_order)

    subscription_schedule_phases = generate_phases_for_initial_order(
      sf_order: sf_order,
      invoice_items: invoice_items,
      subscription_items: subscription_items,
      subscription_schedule: subscription_schedule,
      stripe_customer: stripe_customer
    )

    # TODO refactor once caching is stable, more notes in the `generate_phases_for_initial_rder`
    if subscription_schedule_phases.is_a?(Stripe::Invoice)
      return subscription_schedule_phases
    end

    # TODO should file a Stripe API papercut for this, weird to have the start date on the header but the end date on the phase
    # when creating the subscription schedule the start_date must be specified on the header
    # when updating it, it is specified on the individual phase object

    subscription_start_date_as_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(subscription_schedule.start_date)
    subscription_schedule.start_date = subscription_start_date_as_timestamp

    Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
      subscription_schedule,
      :iterations
    )

    # TODO add mapping support against the subscription schedule phase

    subscription_schedule.customer = stripe_customer.id
    subscription_schedule.phases = subscription_schedule_phases.compact

    # this should never happen, but we are still learning about CPQ
    # if metered billing, quantity is not set, so we set to 1
    if OrderHelpers.extract_all_items_from_subscription_schedule(subscription_schedule).map {|l| l[:quantity] || 1 }.any?(&:zero?)
      Integrations::ErrorContext.report_edge_case("quantity is zero on initial subscription schedule")
    end

    # https://jira.corp.stripe.com/browse/PLATINT-1731
    days_until_due = subscription_schedule.[](:default_settings)&.[](:invoice_settings)&.[](:days_until_due)
    if days_until_due
      subscription_schedule.default_settings.invoice_settings.days_until_due = OrderHelpers.transform_payment_terms_to_days_until_due(days_until_due)
    end

    sanitize(subscription_schedule)

    subscription_schedule = Stripe::SubscriptionSchedule.create(
      subscription_schedule.to_hash,
      @user.stripe_credentials
    )

    log.info 'stripe subscription schedule created', stripe_subscription_schedule_id: subscription_schedule.id

    update_sf_stripe_id(sf_order, subscription_schedule)

    PriceHelpers.auto_archive_prices_on_subscription_schedule(@user, subscription_schedule)

    if @user.feature_enabled?(FeatureFlags::STRIPE_REVENUE_CONTRACT)
      create_revenue_contract_from_sub_schedule(subscription_schedule, sf_order, invoice_items, subscription_items)
    end

    stripe_transaction
  end

  # it is important that the subscription schedule is passed in before the start_date is transformed
  sig do
    params(
      sf_order: T.untyped,
      invoice_items: T::Array[ContractItemStructure],
      subscription_items: T::Array[ContractItemStructure],
      subscription_schedule: Stripe::SubscriptionSchedule,
      stripe_customer: Stripe::Customer
    ).returns(T.any(Stripe::Invoice, [Hash, T.nilable(Hash)]))
  end
  def generate_phases_for_initial_order(sf_order:, invoice_items:, subscription_items:, subscription_schedule:, stripe_customer:)
    string_start_date_from_salesforce = subscription_schedule['start_date']
    # TODO should avoid using blind parse and instead enforce a strict SF datetime format
    start_date_from_salesforce = DateTime.parse(string_start_date_from_salesforce)
    subscription_start_date_as_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(string_start_date_from_salesforce)

    subscription_term_from_salesforce = subscription_schedule['iterations'].to_i

    # originally `iterations` was used, but this fails when subscription term is less than a single billing cycle
    initial_phase_end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
      start_date_from_salesforce + subscription_term_from_salesforce.months
    )

    billing_frequency = OrderAmendment.calculate_billing_frequency_from_phase_items(@user, subscription_items)

    is_initial_order_backend_prorated, _is_initial_order_frontend_prorated = OrderHelpers.prorated_initial_order?(
      phase_items: subscription_items,
      subscription_term: subscription_term_from_salesforce,
      billing_frequency: billing_frequency
    )

    # TODO we should have a check to ensure all quantities are positive
    # TODO we should check if subscription is backddated, if it is, then we should only proceed if the user has a specific flag enabled

    # initial order, so there is only a single phase
    initial_phase = {
      add_invoice_items: invoice_items.map(&:stripe_params),
      items: subscription_items.map(&:stripe_params),
      # TODO so annoying we cannot put the start_date here https://jira.corp.stripe.com/browse/PLATINT-1479
      end_date: initial_phase_end_date,
      metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order),
    }

    if @user.feature_enabled?(FeatureFlags::COUPONS)
      initial_phase["discounts"] = stripe_discounts_for_sf_object(sf_object: sf_order)
    end

    prorated_phase = nil

    # TODO this needs to be gated and synced with the specific flag that CF is using
    if subscription_start_date_as_timestamp >= OrderAmendment.determine_current_time(@user, stripe_customer.id)
      # when `sub_sched_backdating_anchors_on_backdate` is enabled prorating the initial phase
      # does NOT actually prorate it, but instead bills the full amount of the subscription item
      # and locks the billing anchor to the start date of the subscription. Without this flag the
      # billing achor is disconnected from the start date of the invoice
      initial_phase[:proration_behavior] = 'none'
    end

    if is_initial_order_backend_prorated
      # create a new phase to store these items, this will be custom for the initial order
      # pull this '2nd phase creation' out into a separate method so we can use it on the amendment side of things for amendments

      # TODO next, we need to support prorated order amendments
      # make sure the initial state on the order amendments includes the second phase
      # we may need to do some weird phase merging because the initial order phase could be after the order amendment phase

      invoice_items_for_prorations = OrderAmendment.generate_proration_items_from_phase_items(
        mapper: mapper,
        sf_order_amendment: sf_order,
        phase_items: subscription_items,
        subscription_term: subscription_term_from_salesforce,
        billing_frequency: billing_frequency,
      )

      # add special metadata to line items as well
      invoice_items_for_prorations.each do |invoice_item|
        invoice_item[:metadata][StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::BACKEND_PRORATION)] = true
        invoice_item[:period][:end][:type] = 'phase_end'
      end

      backend_proration_term = subscription_term_from_salesforce % billing_frequency
      backend_proration_start_date = start_date_from_salesforce + (subscription_term_from_salesforce - backend_proration_term).months
      backend_proration_start_date_timestamp = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(backend_proration_start_date)

      prorated_phase = {
        # TODO https://jira.corp.stripe.com/browse/PLATINT-1501
        # we do not want to bill the user for a entire billing cycle, so we turn off prorations
        proration_behavior: 'none',

        # in the prorated phase, the item list will be exactly the same as the initial phase
        items: subscription_items.map(&:stripe_params),

        # on initial creation, you cannot specify a start_date!
        end_date: initial_phase_end_date,
        add_invoice_items: invoice_items_for_prorations,
        metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order),
      }

      prorated_phase[:metadata][StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::BACKEND_PRORATION)] = true

      initial_phase['end_date'] = backend_proration_start_date_timestamp
    end

    [initial_phase, prorated_phase]
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

  sig { params(contract_structure: ContractStructure).returns(T.nilable(Stripe::SubscriptionSchedule)) }
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

    if subscription_schedule.status == "canceled"
      Integrations::ErrorContext.report_edge_case('subscription is cancelled, it cannot be modified')
    end

    # at this point, the initial order would have already been translated
    # and a corresponding subscription schedule created

    # verify that all the amendment orders co-terminate with the initial order
    if !OrderAmendment.contract_co_terminated?(mapper, contract_structure)
      raise StripeForce::Errors::RawUserError.new("Order amendments must coterminate with the initial order")
    end

    # Order amendments contain a negative item if they are adjusting a previous line item.

    # How we will know if the line items are the same? Price references won't work.
    # The price of a line item can change across amendments. There's a special field
    # on an amendmended line item that we can leverage: `SBQQ__RevisedOrderProduct__c`.
    # However, this field always references the *first* order line, not the last revised
    # order line.

    # Do NOT use the current state of the subscription schedule at all. This could cause some issues:
    # if the old SF data was mutated in some way (even metadata!) that data will be used, which could
    # cause weird unintended side effects. However, the state of the Stripe side would be too hard to infer.
    # the user could mutate the phase data out-of-band, so we need to (a) limit the phase we are editing
    # (b) recreate each phase data from the raw order line data.

    sf_initial_order_items = order_lines_from_order(contract_structure.initial)
    _, initial_order_phase_items = phase_items_from_order_lines(sf_initial_order_items)

    # another complexity here is 'backend' prorated order amendments (i.e. 18mo contract, billed yearly)
    # in this case, the initial order creates two phases. Effectively, there is an amendment without an order.
    # However, this 'amendment' does *not* contain the latest subscription items. The latest amendment always contains
    # the latest aggregated subscription items. However, we can't lose the add_invoice_items contained in the 2nd phase
    # of a backend order amendment, since that is what will cause the non-subscription (prorated) component of the contract
    # to be billed for.

    is_initial_order_backend_prorated, _is_initial_order_frontend_prorated = OrderHelpers.prorated_initial_order?(
      phase_items: initial_order_phase_items,
      subscription_term: StripeForce::Utilities::SalesforceUtil.extract_subscription_term_from_order!(@mapper, contract_structure.initial),
      billing_frequency: OrderAmendment.calculate_billing_frequency_from_phase_items(@user, initial_order_phase_items)
    )

    # TODO we generally make the assumption that the connector is the only system modifying phases, we should make
    #      this assumption explicit and error when we notice this is not the case
    subscription_phases = subscription_schedule.phases

    if is_initial_order_backend_prorated
      backend_proration, subscription_phases = OrderAmendment.extract_backend_proration_phase(@user, subscription_phases)
    end

    # SF does not enforce mutation restrictions. It's possible to go in and modify anything you want in Salesforce
    # for this reason we should NOT mutate the existing phases of a subscription. This could result in updating quantity, metadata, etc
    # that was updated in Salesforce changing phase data that the user does not expect to be changed.

    # TODO should break this out into a separate method and wrap in `catch_errors_with_salesforce_context`

    # there is more than one amendment in this run
    if contract_structure.amendments.count > 1
      log.info 'processing stacked amendments'
    end

    previous_phase_items = initial_order_phase_items
    contract_structure.amendments.each_with_index do |sf_order_amendment, index|
      locker.lock_salesforce_record(sf_order_amendment)

      log.info 'processing amendment', sf_order_amendment_id: sf_order_amendment.Id, index: index
      invoice_items_in_order, aggregate_phase_items = build_phase_items_from_order_amendment(
        previous_phase_items,
        sf_order_amendment
      )

      # TODO replace with local sync record call in the future
      order_amendment_subscription_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      if order_amendment_subscription_id.present?
        log.info "order amendment already translated, skipping", sf_order_amendment_id: sf_order_amendment.Id, index: index

        # it's important we update the previous phase items since the latest amendment depends on previous amendments
        aggregate_phase_items, _terminated_phase_items = OrderHelpers.remove_terminated_lines(aggregate_phase_items)
        previous_phase_items = aggregate_phase_items
        next
      end


      is_order_terminated = aggregate_phase_items.all?(&:fully_terminated?)
      if is_order_terminated
        log.info 'amendment is termination order', sf_order_amendment_id: sf_order_amendment.Id

        if contract_structure.amendments.count - 1 != index
          raise StripeForce::Errors::RawUserError.new("Processing a termination order, but there's more amendments.")
        end
      end

      # this loop excludes the initial phase, which is why we are subtracting by 1
      # note this can happen with backend initial order prorations or if the user has amendment the contract previously
      if subscription_phases.count - 1 > index
        log.info 'number of subscription phase count is greater than this amendment index', subscription_phase_count: subscription_phases.count, amendment_index: index
      end

      # TODO price ID dup check on the invoice items
      # make sure to run this *after* checking if the amendment is already processed, otherwise
      # we'll create price IDs which will never be archived since they won't be used in an active phase
      aggregate_phase_items = OrderHelpers.ensure_unique_phase_item_prices(@user, aggregate_phase_items)

      # TODO should probably use a completely different key/mapping for the phase items
      phase_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order_amendment, Stripe::SubscriptionSchedule)

      string_start_date_from_salesforce = phase_params['start_date']
      sf_order_amendment_start_date_as_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(string_start_date_from_salesforce)
      phase_params['start_date'] = sf_order_amendment_start_date_as_timestamp

      # TODO should probably move iteration extraction to another helper
      subscription_term_from_salesforce = phase_params.delete('iterations').to_i

      # originally `iterations` was used, but this fails when subscription term is less than a single billing cycle

      # we need to normalize the amendment order end date to take into account special cases
      # that can occur when then initial order starts on a day of the month that doesn't exist in the amendment month
      sf_order_amendment_end_date = StripeForce::Utilities::SalesforceUtil.normalize_sf_order_amendment_end_date(mapper: mapper, sf_order_amendment: sf_order_amendment, sf_initial_order: contract_structure.initial)
      # at this point, we have verified the order amendment end date is equal to the initial order
      phase_params['end_date'] = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(sf_order_amendment_end_date)

      # TODO should we validate the end date vs the subscription schedule?

      billing_frequency = OrderAmendment.calculate_billing_frequency_from_phase_items(@user, aggregate_phase_items)

      aggregate_phase_items, terminated_phase_items = OrderHelpers.remove_terminated_lines(aggregate_phase_items)

      # determine if this is a backdated order since this has implications on how we prorate
      stripe_customer_id = T.cast(subscription_schedule.customer, String)
      current_time = OrderAmendment.determine_current_time(@user, stripe_customer_id)
      normalized_current_time = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(Time.at(current_time))
      is_order_backdated = sf_order_amendment_start_date_as_timestamp < normalized_current_time && @user.feature_enabled?(StripeForce::Constants::FeatureFlags::BACKDATED_AMENDMENTS)

      backdated_billing_cycles = nil
      next_billing_timestamp = nil
      if is_order_backdated
        log.info 'processing a backdated amendment order'
        backdated_billing_cycles = 0
        subscription_schedule_start = T.must(subscription_schedule.phases.first).start_date
        next_billing_timestamp = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(Time.at(subscription_schedule_start))

        # determine if a billing cycle has passed between the amendment start date and current time
        while next_billing_timestamp <= normalized_current_time
          if next_billing_timestamp > sf_order_amendment_start_date_as_timestamp
            backdated_billing_cycles += 1
          end
          next_billing_timestamp = (Time.at(next_billing_timestamp).utc.beginning_of_day + billing_frequency.months).to_i
        end

        # If a subscription is backdated by 3 months and began on January 30th, our next_billing_timestamp will be March 27th instead of March 30th.
        # ie January 30th, + 1 billing cycle (1 month) will result in Febuary 27th, the second iteration will result in March 27th (instead of March 30th).
        next_billing_datetime = Time.at(next_billing_timestamp).utc
        sf_order_amendment_start_date_datetime = Time.at(sf_order_amendment_start_date_as_timestamp).utc

        if next_billing_datetime.day != sf_order_amendment_start_date_datetime.day
          next_billing_timestamp = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: next_billing_datetime, anchor_day_of_month: sf_order_amendment_start_date_datetime.day).to_i
        end
      end

      negative_invoice_items = []
      if @user.feature_enabled?(FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT)
        # https://jira.corp.stripe.com/browse/PLATINT-2092
        # For now, we will only issue credits for terminated lines (to support the most requested feature of license expansion / upgrades)
        #     in the future, we could augment the below method to take in all aggregate_phase items and account for partial line termination credits.
        #     i.e. when a customer is moving from 50 to 30 quantity of a subscription item  half way through a billing cycle. The merchant may want to issue credit
        #     for the 20 units removed. However, for now we would only issue credit if they terminate the line by removing all 50 units.
        negative_invoice_items = OrderAmendment.generate_proration_credits_from_terminated_phase_items(
          user: @user,
          mapper: mapper,
          sf_order_amendment: sf_order_amendment,
          terminated_phase_items: terminated_phase_items,
          subscription_term: subscription_term_from_salesforce,
          billing_frequency: billing_frequency,
          is_order_backdated: is_order_backdated,
          next_billing_timestamp: next_billing_timestamp,
        )
      end

      # TODO validate that all prices have the same recurrence? Stripe does this downstream,
      #      but at this point we assume that this check has already done, so it may make sense
      #      to do this check more explicitly.

      # at this point, we have the finalized list of non-prorated order lines
      # this means all price data has been mapped and converted into Stripe line items
      # and we can calculate the finalized billing cycle of the order amendment

      if !is_order_terminated
        # if the amendment is prorated, then all line items will have prorated component
        is_prorated = OrderAmendment.prorated_amendment?(
          user: @user,
          aggregate_phase_items: aggregate_phase_items,
          subscription_schedule: subscription_schedule,

          # these params in an ideal world should be pulled from the `subscription_schedule`, but
          # they are tricky to extract without additional API calls
          subscription_term: subscription_term_from_salesforce,
          billing_frequency: billing_frequency,
          amendment_start_date: sf_order_amendment_start_date_as_timestamp,
        )
      end

      invoice_items_for_prorations = []
      if !is_order_terminated && is_prorated
        log.info 'amendment order is prorated', sf_order_amendment_id: sf_order_amendment.Id, index: index

        # the subscription term represents the number of whole months
        # the days prorating represents the partial month (or days) remaining
        subscription_term = subscription_term_from_salesforce
        days_prorating = 0

        if @user.feature_enabled?(FeatureFlags::NON_ANNIVERSARY_AMENDMENTS)
          days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
            sf_order_start_date: StripeForce::Utilities::SalesforceUtil.salesforce_date_to_beginning_of_day(string_start_date_from_salesforce),
            sf_order_end_date: sf_order_amendment_end_date,
            sf_order_subscription_term: subscription_term)

          # CPQ Proration Calculations
          # https://help.salesforce.com/s/articleView?id=sf.cpq_subscriptions_prorate_precision_1.htm&type=5
          # in CPQ, the proration multiple is the number of whole months plus a decimal for any partial month at the end of the term
          # which is represented by the subscription term and the days below
          # depending on the CPQ setting <=> feature flag enabled, the calculation will differ
          if days > 0
            # if feature DAY_PRORATIONS is enabled, set the number of days to prorate
            # else, a partial month equals a whole month so add one to the subscription term
            if @user.feature_enabled?(FeatureFlags::DAY_PRORATIONS)
              log.info 'prorating line items by days', days_prorating: days
              days_prorating = days
            else
              # the subscription term represents the number of whole months
              # plus a decimal for any partial month at the end of the term if your term contains a partial month (days_prorating > 0)
              # so we round the number of months to the nearest whole number (if there are days) by adding one since the subscription_term
              log.info 'prorating line items by months but accounting for partial month', days_prorating: days
              subscription_term += 1
            end
          end
        end

        invoice_items_for_prorations = OrderAmendment.generate_proration_items_from_phase_items(
          mapper: mapper,
          sf_order_amendment: sf_order_amendment,
          phase_items: aggregate_phase_items,
          subscription_term: subscription_term,
          billing_frequency: billing_frequency,
          days_prorating: days_prorating,
          backdated_billing_cycles: backdated_billing_cycles,
          next_billing_timestamp: next_billing_timestamp,
        )
      end

      # for debugging
      aggregate_phase_items.each do |phase_item|
        log.info 'including item on order', order_line_id: phase_item.order_line_id
      end

      new_phase = Stripe::StripeObject.construct_from({
        add_invoice_items: invoice_items_in_order.map(&:stripe_params) + invoice_items_for_prorations + negative_invoice_items,

        # `deep_copy` is important, otherwise multiple phase changes in a
        # single job run will use the same aggregate phase items
        items: Integrations::Utilities::StripeUtil.deep_copy(aggregate_phase_items).map(&:stripe_params),

        # TODO move to global defaults https://jira.corp.stripe.com/browse/PLATINT-1501
        proration_behavior: 'none',

        metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order_amendment),
      }.merge(phase_params))

      if @user.feature_enabled?(FeatureFlags::COUPONS)
        new_phase["discounts"] = stripe_discounts_for_sf_object(sf_object: sf_order_amendment)
      end

      # if the time ranges are identical, then the previous phase should be removed
      # the previous phases subscription items should be overwritten by the latest phase calculation
      # but any one-off items would be lost without "merging" these items
      previous_phase = T.must(subscription_phases.last)

      # it's important to check this before setting that start date to 'now' below
      is_identical_to_previous_phase_time_range = previous_phase.start_date == new_phase.start_date &&
        previous_phase.end_date == new_phase.end_date

      # if the current day is the same day as the start day, then use 'now'
      is_same_day = normalized_current_time == new_phase.start_date
      if is_same_day
        log.info 'amendment starts on the current day, using now'
        new_phase.start_date = 'now'
      elsif @user.feature_enabled?(FeatureFlags::BACKDATED_AMENDMENTS) && is_order_backdated
        # if this is a backdated amendment, then use the current time to update the subscription schedule
        log.info 'backdated amendment, using now'
        new_phase.start_date = 'now'
      end

      should_merge_phases = is_identical_to_previous_phase_time_range && !is_same_day && !is_order_backdated
      if should_merge_phases && !previous_phase.add_invoice_items.empty?
        log.info 'previous phase is identical, merging invoice items'
        new_phase.add_invoice_items += previous_phase.add_invoice_items
      end

      # align date boundaries of the schedules
      previous_phase.end_date = new_phase.start_date

      # TODO I wonder if we can do something smarter here: if the invoice has not been paid/billed, do XYZ?
      # this is a special case: subscription is cancelled on the same day, the intention here is to not bill the user at all
      is_subscription_schedule_cancelled = is_order_terminated && (is_same_day || is_order_backdated)

      # if the order is terminated, updating the last phase end date and NOT adding another phase is all that needs to be done
      if !is_order_terminated
        # !is_same_day is important since these items may have already been invoiced
        if should_merge_phases
          # https://jira.corp.stripe.com/browse/PLATINT-1815
          log.info 'previous phase identical, removing previous phase'
          subscription_phases.delete_at(subscription_phases.count - 1)
        end

        # TODO: https://jira.corp.stripe.com/browse/PLATINT-2091
        # We are only adding a new phase if an order is not fully terminated, therefore our negative invoice items are never
        # created because the new phase does not get created. If merchants wish to issue credits for fully cancelled orders we should
        # think about conditionally making a call to create our negative_invoice_items here.
        subscription_phases << new_phase
      end

      subscription_phases = OrderHelpers.sanitize_subscription_schedule_phase_params(subscription_phases)

      # https://jira.corp.stripe.com/browse/PLATINT-1832
      subscription_phases = OrderAmendment.delete_past_phases(@user, stripe_customer_id, subscription_phases)

      # TODO we do not currently map to the subscription schedule (again) when there is an amendment order
      # we should consider remapping the subscription schedule when there is an amendment order but for now we will map specific fields
      subscription_schedule = apply_amendment_order_mappings(mapper, subscription_schedule, sf_order_amendment)

      # NOTE intentional decision here NOT to update any other subscription fields
      catch_errors_with_salesforce_context(secondary: sf_order_amendment) do
        if is_subscription_schedule_cancelled
          log.info 'cancelling subscription immediately', sf_order_amendment_id: sf_order_amendment

          # NOTE the intention here is to void/reverse out the entire contract, this is the closest API call we have
          subscription_schedule.cancel(
            {
              invoice_now: false,
              prorate: false,
            },
            @user.stripe_credentials
          )
        else
          log.info 'adding phase',
            sf_order_amendment_id: sf_order_amendment.Id,
            start_date: new_phase.start_date,
            end_date: new_phase.end_date

          # we do NOT want the next amendment loop to use the version of subscription phases with the backend proration in place
          final_subscription_phases = if is_initial_order_backend_prorated
            OrderAmendment.inject_backend_proration(subscription_phases, backend_proration)
          else
            subscription_phases
          end

          # `none` is really important to ensure that the user is not billed by stripe for any prorated amounts
          subscription_schedule.proration_behavior = 'none'
          subscription_schedule.phases = final_subscription_phases

          # note: to support stacked amendments, we want to update the local sub_schedule and sub_phases
          # because  Stripe converts 'now' to a timestamp
          # and we want to use that timestamp when there is a stacked amendment
          subscription_schedule = T.cast(subscription_schedule.save({}, @user.stripe_credentials), Stripe::SubscriptionSchedule)
        end
      end

      update_sf_stripe_id(sf_order_amendment, subscription_schedule)
      subscription_phases = T.cast(subscription_schedule.phases, T::Array[Stripe::SubscriptionSchedulePhase])
      previous_phase_items = aggregate_phase_items
    end

    PriceHelpers.auto_archive_prices_on_subscription_schedule(@user, subscription_schedule)

    subscription_schedule
  end

  sig do
    params(mapper: StripeForce::Mapper, subscription_schedule: Stripe::SubscriptionSchedule, sf_order_amendment: Restforce::SObject).returns(Stripe::SubscriptionSchedule)
  end
  def apply_amendment_order_mappings(mapper, subscription_schedule, sf_order_amendment)
    # a different invoice_rendering_template could be used between the initial and amendment order therefore extract the mapped values again
    if @user.feature_enabled?(StripeForce::Constants::FeatureFlags::INVOICE_RENDERING_TEMPLATE)
      invoice_rendering_template = StripeForce::Utilities::SalesforceUtil.extract_optional_fields_from_order(mapper, sf_order_amendment, ['subscription_schedule', 'default_settings.invoice_settings.rendering.template'])

      if !invoice_rendering_template.nil?
        invoice_rendering_template_version = StripeForce::Utilities::SalesforceUtil.extract_optional_fields_from_order(mapper, sf_order_amendment, ['subscription_schedule', 'default_settings.invoice_settings.rendering.template_version'])
        subscription_schedule[:default_settings][:invoice_settings][:rendering] = {"template": invoice_rendering_template, "template_version": invoice_rendering_template_version}
      end
    end

    # update days_until_due on sub schedule
    days_until_due = StripeForce::Utilities::SalesforceUtil.extract_optional_fields_from_order(mapper, sf_order_amendment, ['subscription_schedule', 'default_settings.invoice_settings.days_until_due'])
    if days_until_due.present?
      days_until_due = T.cast(days_until_due, T.any(String, Integer, Float))
      subscription_schedule[:default_settings][:invoice_settings][:days_until_due] = OrderHelpers.transform_payment_terms_to_days_until_due(days_until_due)
    end

    subscription_schedule
  end

  sig do
    params(
      original_aggregate_phase_items: T::Array[StripeForce::Translate::ContractItemStructure],
      new_phase_items: T::Array[StripeForce::Translate::ContractItemStructure]
    ).returns(T::Array[StripeForce::Translate::ContractItemStructure])
  end
  def merge_subscription_line_items(original_aggregate_phase_items, new_phase_items)
    # no side effects, please!
    aggregate_phase_items = original_aggregate_phase_items.dup

    termination_lines, additive_lines = new_phase_items.partition(&:termination?)

    additive_lines.each do |new_subscription_item|
      log.info 'adding new line item', order_line_id: new_subscription_item.order_line_id
      aggregate_phase_items << new_subscription_item
    end

    # NOTE this terminates the lines, but does NOT remove them
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

    # line items that are "new" (i.e. not revising anything) are "origin" lines which future
    # revisions should be mapped to
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

    log.info "order amendment revision map",
      revision_map: revision_map.transform_values {|ci| ci.map(&:order_line_id) }

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

      log.info 'reducing quantity on line', reducing_line: termination_line.order_line_id, quantity: termination_line.quantity
      termination_line.quantity.abs.times do
        fifo_remaining_line = fifo_order_line_stack
          .reject(&:fully_terminated?)
          .reverse
          .first

        if fifo_remaining_line
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
    sf_order_items = cache_service.get_related_records_from_cache(
      sf_order[SF_ID],
      :OrderId,
      SF_ORDER_ITEM,
    )

    sf_order_items.select do |sf_order_item|
      # never expect this to occur
      if sf_order_item.IsDeleted || !sf_order_item.SBQQ__Activated__c
        Integrations::ErrorContext.report_edge_case("order line is deleted or not activated")
      end

      should_keep = sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)].nil? ||
        !sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)]

      if !should_keep
        log.info 'order line marked as skipped', sf_order_line_id: sf_order_item.Id
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
      # this is a critical step: this performs the complicated logic of comparing pricebook & order line prices
      # and creating a new price if needed
      price = catch_errors_with_salesforce_context(secondary: sf_order_item) do
        create_price_for_order_item(sf_order_item)
      end

      # could occur if a coupon is required for a negative amount, although this should probably be built into the `price` method instead
      next if price.nil?

      # a phase item is NOT a subscription item, but they are close, and right now the data mapper & revelant keys uses these across the codebase
      # I wonder if they will eventually converge in the public Stripe API over time?
      phase_item = Stripe::SubscriptionItem.construct_from({
        price: price.id,
        metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order_item),
      })

      if @user.feature_enabled?(FeatureFlags::COUPONS)
        phase_item.discounts = stripe_discounts_for_sf_object(sf_object: sf_order_item)
      end

      phase_item_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order_item, Stripe::SubscriptionItem)
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

      # quantity cannot be specified if usage type is metered, we have to delete the quantity param before sending to Stripe
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
        invoice_items << phase_item_struct
      end
    end

    [invoice_items, subscription_items]
  end

  sig { params(sf_initial_order: Restforce::SObject).returns(T.nilable(String)) }
  def extract_contract_id_from_initial_order(sf_initial_order)
    # if this occurs, the user's CPQ is not configured properly/as we assume
    if sf_initial_order[SF_ORDER_QUOTE].blank?
      raise StripeForce::Errors::RawUserError.new("No quote associated with an order. Orders pushed to Stripe must have a related CPQ Quote.")
    end

    contract = cache_service.get_related_record_from_cache(
      sf_initial_order[SF_ORDER_QUOTE],
      SF_CONTRACT_QUOTE_ID.to_sym,
      SF_CONTRACT
    )

    # this can occur if contracts are processed async
    if contract.nil?
      log.info 'order is contracted, but no contract is associated'
      return nil
    end

    contract[SF_ID]
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

    cache_service.get_record_from_cache(SF_ORDER, initial_order_id)
  end

  # TODO this should be a function in a helper module
  sig { params(sf_contract_id: String).returns(Array) }
  def order_amendments_for_contract(sf_contract_id)

    # TODO: use this and add a sort func
    # order_amendments = cache_service.search_cache_via_nested_field(
    #   SF_ORDER,
    #   [SF_OPPORTUNITY, CPQ_AMENDED_CONTRACT],
    #   sf_contract_id
    # )

    order_amendment_query = sf.query(
      <<~EOL
        SELECT Id FROM #{SF_ORDER}
        WHERE Opportunity.SBQQ__AmendedContract__c = '#{sf_contract_id}'
        ORDER BY SBQQ__Quote__r.SBQQ__StartDate__c, LastModifiedDate  ASC
      EOL
    )

    log.info 'order amendments found',
      count: order_amendment_query.count,
      amendmend_ids: order_amendment_query.map(&:Id)

    order_amendment_query.map {|i| cache_service.get_record_from_cache(SF_ORDER, i[SF_ID]) }
  end
end
