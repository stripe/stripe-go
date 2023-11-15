# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  include StripeForce::Constants

  def validate_evergreen_order_items(sf_order_items)
    sf_order_items.each do |sf_order_item|
      # subscription term field can only be 1
      if sf_order_item[CPQ_QUOTE_SUBSCRIPTION_TERM] != 1 || sf_order_item[CPQ_DEFAULT_SUBSCRIPTION_TERM] != 1
        throw_user_failure!(
          salesforce_object: sf_order_item,
          message: "Evergreen Salesforce orders should have default subscription term and quote subscription term equal to 1."
        )
      end
    end
  end

  def sf_order_contains_evergreen_order_items(sf_order_items, sf_order)
    evergreen_items = sf_order_items.select {|item| item[CPQ_PRODUCT_SUBSCRIPTION_TYPE] == CPQProductSubscriptionTypeOptions::EVERGREEN.serialize }
    renewable_items = sf_order_items.select {|item| item[CPQ_PRODUCT_SUBSCRIPTION_TYPE] == CPQProductSubscriptionTypeOptions::RENEWABLE.serialize }

    return false if evergreen_items.empty?

    unless renewable_items.empty?
      throw_user_failure!(
        salesforce_object: sf_order,
        message: "Salesforce orders with both Evergreen and Renewable items are not yet supported."
      )
    end

    true
  end

  def create_stripe_subscription_from_initial_evergreen_sf_order(sf_order, subscription_items, stripe_customer)
    subscription = Stripe::Subscription.construct_from({
      cancel_at_period_end: false,
      customer: stripe_customer.id,
      metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order),
      items: subscription_items.map(&:stripe_params),
    })

    subscription_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order, Stripe::SubscriptionSchedule)

    mapper.assign_values_from_hash(subscription, subscription_params)
    apply_mapping(subscription, sf_order)

    # Subscriptions can only start now or in future => https://jira.corp.stripe.com/browse/PLATINT-2862
    Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
      subscription,
      [:end_behavior, :iterations]
    )

    stripe_customer_id = T.cast(subscription.customer, String)
    subscription_start_time = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, T.must(sf_order)).utc.beginning_of_day.to_i
    current_time = Time.at(OrderAmendment.determine_current_time(@user, stripe_customer_id)).utc.beginning_of_day.to_i

    if subscription_start_time < current_time
      throw_user_failure!(
        salesforce_object: sf_order,
        message: "Backdated evergreen Salesforce orders are not yet supported."
      )
    end

    if subscription_start_time == current_time
      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
        subscription,
        [:start_date]
      )

      subscription = Stripe::Subscription.create(
        subscription.to_hash,
        @user.stripe_credentials
      )

      update_sf_stripe_id(sf_order, subscription)
      log.info 'created stripe subscription for evergreen order', stripe_subscription_id: subscription.id
    else
      sf_order_items = order_lines_from_order(sf_order)
      invoice_items, subscription_items = phase_items_from_order_lines(sf_order_items)

      subscription_schedule = Stripe::SubscriptionSchedule.construct_from({
        end_behavior: StripeEndBehavior::RELEASE.serialize,
        metadata: Metadata.stripe_metadata_for_sf_object(@user, sf_order),
      })

      subscription_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_order, Stripe::SubscriptionSchedule)

      mapper.assign_values_from_hash(subscription_schedule, subscription_params)
      apply_mapping(subscription_schedule, sf_order)

      subscription_schedule_phases = generate_phases_for_initial_order(
        sf_order: sf_order,
        invoice_items: invoice_items,
        subscription_items: subscription_items,
        subscription_schedule: subscription_schedule,
        stripe_customer: stripe_customer
      )

      if subscription_schedule_phases.is_a?(Stripe::Invoice)
        return subscription_schedule_phases
      end

      subscription_start_date_as_timestamp = StripeForce::Utilities::SalesforceUtil.salesforce_date_to_unix_timestamp(subscription_schedule.start_date)
      subscription_schedule.start_date = subscription_start_date_as_timestamp

      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
        subscription_schedule,
        :iterations
      )

      subscription_schedule.customer = stripe_customer.id
      subscription_schedule.phases = subscription_schedule_phases.compact

      # there should only be one subscription schedule phase
      if subscription_schedule.phases.count != 1
        raise Integrations::Errors::ImpossibleState.new("Evergreen order generated subscription schedule with more than one phase.")
      end

      # set cancelation date to be immediately after phases get created
      # so subscription schedule gets released asap
      subscription_schedule.phases.first[:end_date] = (subscription_start_time + 1).to_i

      sanitize(subscription_schedule)

      subscription_schedule = Stripe::SubscriptionSchedule.create(
        subscription_schedule.to_hash,
        @user.stripe_credentials
      )

      update_sf_stripe_id(sf_order, subscription_schedule)
      log.info 'created stripe subscription schedule for evergreen order', stripe_subscription_id: subscription_schedule.id
    end

    subscription
  end

  sig { params(contract_structure: ContractStructure).returns(T.nilable(Stripe::Subscription)) }
  def update_subscription_from_order_amendments(contract_structure)
    log.info "processing evergreen Salesforce order amendments"

    sf_order = contract_structure.initial
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    if stripe_id.nil?
      raise Integrations::Errors::UserError.new("Could not find corresponding Stripe subscription for initial evergreen Salesforce order.")
    end

    # if amendment is for an evergreen order whose stripe id is a subscription schedule
    # need to check if subscription has been released and change the stripe id to the subscription
    if stripe_id.start_with?('sub_sched')
      subscription_schedule = retrieve_from_stripe(
        Stripe::SubscriptionSchedule,
        sf_order
      )

      if !subscription_schedule
        raise Integrations::Errors::ImpossibleState.new("Could not find corresponding Stripe subscription subscription for a scheduled evergreen Salesforce order")
      end

      subscription_schedule = T.cast(subscription_schedule, Stripe::SubscriptionSchedule)

      if subscription_schedule.released_subscription.nil?
        throw_user_failure!(
          salesforce_object: sf_order,
          message: 'Amending a Salesforce evergreen order that has not started is not yet supported.'
        )
      end

      subscription_id = T.must(subscription_schedule.released_subscription)
      subscription = Stripe::Subscription.retrieve(subscription_id, @user.stripe_credentials)
      update_sf_stripe_id(sf_order, subscription)

      # refresh to include updated subscription id instead of subscription schedule id
      sf_order.refresh
    end

    subscription = retrieve_from_stripe(Stripe::Subscription, sf_order)
    if !subscription
      raise StripeForce::Errors::RawUserError.new("Could not find corresponding Stripe subscription for initial evergreen Salesforce order.")
    end

    subscription = T.cast(subscription, Stripe::Subscription)
    stripe_customer_id = T.cast(subscription.customer, String)
    subscription_id = subscription.id

    sf_initial_order_items = order_lines_from_order(contract_structure.initial)
    is_initial_order_renewal = StripeForce::Utilities::SalesforceUtil.is_renewal_order(@cache_service, contract_structure.initial)
    _, initial_order_phase_items = phase_items_from_order_lines(sf_initial_order_items, is_initial_order_renewal)

    previous_phase_items = initial_order_phase_items
    amendments = contract_structure.amendments
    amendments.each do |sf_order_amendment|
      # do not support adding renewable products to evergreen order
      sf_order_items = order_lines_from_order(sf_order_amendment)
      if !sf_order_contains_evergreen_order_items(sf_order_items, sf_order_amendment)
        throw_user_failure!(
          salesforce_object: sf_order_amendment,
          message: "Adding Salesforce products with a 'Renewable' type to evergreen orders isn't supported."
        )
      end

      # ignore amendment if already synced before (aka has stripe id)
      order_amendment_subscription_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      if order_amendment_subscription_id.present?
        log.info "order amendment already translated, skipping", sf_order_amendment_id: sf_order_amendment.Id

        _, subscription_items = build_phase_items_from_order_amendment(previous_phase_items, sf_order_amendment)
        previous_phase_items = subscription_items
        next
      end

      # cannot make new amendment to canceled subscription
      if subscription.status == 'canceled'
        # if attempting to make a new amendment to a canceled subscription schedule, raise error
        throw_user_failure!(
          salesforce_object: sf_order_amendment,
          message: "The Stripe subscription for the evergreen order has already been cancelled and can't be modified."
        )
      end

      # logic to process new amendment to evergreen order
      _, subscription_items = build_phase_items_from_order_amendment(
        previous_phase_items,
        sf_order_amendment
      )

      sf_order_items = order_lines_from_order(sf_order_amendment)
      _, amendment_subscription_items = phase_items_from_order_lines(sf_order_items)

      is_order_terminated = subscription_items.all?(&:fully_terminated?)
      if !is_order_terminated
        # amendment is not a full termination
        amendment_start_time = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, T.must(sf_order_amendment)).utc.beginning_of_day.to_i
        current_time = Time.at(OrderAmendment.determine_current_time(@user, stripe_customer_id)).utc.beginning_of_day.to_i

        if amendment_start_time == current_time
          Stripe::Subscription.update(
            subscription_id,
            {items: Integrations::Utilities::StripeUtil.deep_copy(amendment_subscription_items).map(&:stripe_params)},
            @user.stripe_credentials
          )

          log.info 'added subscription items to stripe subscription for evergreen order', stripe_subscription_id: subscription.id
        else
          throw_user_failure!(
            salesforce_object: sf_order,
            message: "Amendment(s) with start dates in the future was not processed by the connector."
          )
        end

        previous_phase_items = subscription_items
      else
        # amendment is for cancelation
        log.info 'amendment is for termination', amendment_id: sf_order_amendment.Id

        # the first amendment must be for cancelation if code gets to this point
        # check termination date using start date of amendment
        amendment_start_time = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, T.must(amendments.first)).beginning_of_day.to_i
        current_time = Time.at(OrderAmendment.determine_current_time(@user, stripe_customer_id)).utc.beginning_of_day.to_i

        if amendment_start_time == current_time
          # order is to be canceled right now
          log.info "Canceling Stripe subscription", stripe_object: subscription.id, salesforce_object: sf_order_amendment

          subscription.cancel(
            {
              invoice_now: false,
              prorate: false,
            },
            @user.stripe_credentials
          )
        else
          # order is to be canceled in the future
          log.info "Updating Stripe subscription", stripe_object: subscription.id, salesforce_object: sf_order_amendment

          Stripe::Subscription.update(
            subscription.id,
            {
              cancel_at: amendment_start_time,
              proration_behavior: StripeProrationBehavior::NONE.serialize,
            },
            @user.stripe_credentials,
          )
        end

        # serves as a marker that the amendment has been synced
        update_sf_stripe_id(sf_order_amendment, subscription)
      end
    end

    subscription
  end
end
