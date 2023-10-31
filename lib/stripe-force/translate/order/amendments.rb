# frozen_string_literal: true
# typed: true

# https://docs.google.com/document/d/15d1nO0wkGKGyLS7fRCflI_3LEPQd6yOBhtkDUS6-_hk/edit
class StripeForce::Translate
  # using `Order::Amendment` conflicts will too many other things
  module OrderAmendment
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    def self.inject_backend_proration(original_subscription_phases, backend_proration_phase)
      subscription_phases = Integrations::Utilities::StripeUtil.deep_copy(original_subscription_phases)

      # if there is a phase starting *on* the backend_proration start date, then we can simply add the add_invoice_items to that phase
      exact_phase_match = subscription_phases.detect {|p| p.start_date == backend_proration_phase.start_date }
      if exact_phase_match
        log.info 'exact backend phase match, merging invoice items'
        exact_phase_match.add_invoice_items += backend_proration_phase.add_invoice_items
        return subscription_phases
      end

      raise StripeForce::Errors::RawUserError.new("The backend prorated order amendment case isn't supported.")

      # Is there a phase starting after the backend_proration phase.start_date? If so, we need to add the backend proration phase
      # before that phase and set backend_proration phase.end_date to the start date of the other phase.

      # We know there is at least one amendment, so at this point it must be starting before the start date
      # of the backend proration phase. In this case we need to take the subscription items from that phase
      # and use them instead of the subscription items on the backend proration phase
    end

    sig do
      params(
        user: StripeForce::User,
        original_subscription_phases: T::Array[Stripe::SubscriptionSchedulePhase]
      ).returns([
        # since the objects are *really* not SubscriptionSchedulePhase we can't fool sorbet-runtime
        T.nilable(Stripe::StripeObject),
        T::Array[Stripe::SubscriptionSchedulePhase],
      ])
    end
    def self.extract_backend_proration_phase(user, original_subscription_phases)
      subscription_phases = Integrations::Utilities::StripeUtil.deep_copy(original_subscription_phases)

      # TODO we probably need to adjust the starting and ending dates if we remove the phase

      matching_indexes = T.let([], T::Array[Integer])

      # intentionally mutates the phase objects
      subscription_phases.each_with_index do |phase, index|
        # this isn't a *great* way to detect the backend proration phase since this metadata could be changed by the user
        # a better approach would be to compare the add_invoice_items & subscription_items, but this is surprisingly hard to do
        # because price IDs are not consistent (can't have dup price IDs on a sub), metadata could change, etc. To avoid this
        # complexity initially we simply check for metadata. We know there is exactly one
        if phase[:metadata][StripeForce::Translate::Metadata.metadata_key(user, MetadataKeys::BACKEND_PRORATION)] == "true"
          matching_indexes << index
        end
      end

      if matching_indexes.empty?
        log.warn 'no backend proration phase found'
        return [nil, subscription_phases]
      end

      if matching_indexes.size > 1
        raise Integrations::Errors::TranslatorError.new("more than one backend proration phase found")
      end

      matching_index = matching_indexes.first
      backend_proration = subscription_phases.delete_at(matching_index)

      [backend_proration, subscription_phases]
    end

    # determines if order amendments coterminate with the initial order,
    # this is required by our integration but is allowed by CPQ in some situations
    sig { params(mapper: StripeForce::Mapper, contract_structure: ContractStructure).returns(T::Boolean) }
    def self.contract_co_terminated?(mapper, contract_structure)
      initial_order_end_date = StripeForce::Utilities::SalesforceUtil.calculate_order_end_date(mapper, contract_structure.initial)

      # the end date must be the same for all order amendments
      amendment_orders_end_dates = contract_structure.amendments.map do |sf_order_amendment|
        StripeForce::Utilities::SalesforceUtil.normalize_sf_order_amendment_end_date(mapper: mapper, sf_order_amendment: sf_order_amendment, sf_initial_order: contract_structure.initial)
      end

      is_co_terminated = amendment_orders_end_dates.all? do |sf_amendment_end_date|
        sf_amendment_end_date.to_i == initial_order_end_date.to_i
      end

      if !is_co_terminated
        log.info 'amendment order is not coterminated',
          initial_end_date: initial_order_end_date,
          amendment_end_dates: amendment_orders_end_dates
      end

      is_co_terminated
    end

    sig do
      params(
        stripe_price: Stripe::Price,
        subscription_term: Integer,
        product_subscription_term: Integer,
        billing_frequency: Integer,
        days_prorating: Integer,
      ).returns(BigDecimal)
    end
    def self.calculate_prorated_billing_amount(stripe_price:, subscription_term:, product_subscription_term:, billing_frequency:, days_prorating:)
      # prorated billing amount is months of subscription term that is not included
      prorated_subscription_term = subscription_term % billing_frequency
      proration_percentage = BigDecimal(prorated_subscription_term) / BigDecimal(billing_frequency)

      if days_prorating > 0
        # at this point, we know feature DAY_PRORATIONS is enabled since days_prorating is non-zero
        # and we calculate the partial month value in our prorate multiplier like CPQ-based calculations
        # https://help.salesforce.com/s/articleView?id=sf.cpq_subscriptions_prorate_precision_1.htm&type=5
        proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (subscription_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: product_subscription_term)
      end

      # https://jira.corp.stripe.com/browse/PLATINT-1808
      if days_prorating == 0 && prorated_subscription_term.zero?
        log.warn 'subscription term is equal to billing frequency, amendment is most likely happening on the same day'
        proration_percentage = 1
      end

      unit_amount_decimal = BigDecimal(stripe_price.unit_amount_decimal)
      prorated_billing_amount = unit_amount_decimal * proration_percentage
      prorated_billing_amount
    end

    # NOTE at this point it's assumed that the price is NOT a metered billing item or tiered price
    sig do
      params(
        mapper: StripeForce::Mapper,
        phase_item: ContractItemStructure,
        subscription_term: Integer,
        billing_frequency: Integer,
        days_prorating: Integer,
        backdated_billing_cycles: Integer
      ).returns(Stripe::Price)
    end
    def self.create_prorated_price_from_phase_item(mapper:, phase_item:, subscription_term:, billing_frequency:, days_prorating:, backdated_billing_cycles:)
      # at this point, we know what the billing amount is per billing cycle, since that has already been defined
      # on the price object at the order line. We calculate the percentage of the original line price that should
      # be billed and multiply that against the decimal price on the stripe price.
      # we also have validated that it isn't a tiered or metered billing price
      user = mapper.user
      stripe_price = phase_item.price(user)
      sf_order_item = phase_item.order_line
      sf_order_amendment = user.sf_client.find(SF_ORDER, sf_order_item["OrderId"])
      product_subscription_term = StripeForce::Utilities::SalesforceUtil.determine_quote_line_subscription_term(mapper, sf_order_item, sf_order_amendment)
      prorated_billing_amount = calculate_prorated_billing_amount(stripe_price: stripe_price, subscription_term: subscription_term, product_subscription_term: product_subscription_term, billing_frequency: billing_frequency, days_prorating: days_prorating)

      # if the order is backdated and was synced after a billing cycle, we need to add the amount of the backdated billing cycle
      if user.feature_enabled?(FeatureFlags::BACKDATED_AMENDMENTS) && backdated_billing_cycles > 0
        prorated_billing_amount += stripe_price.unit_amount_decimal.to_d * backdated_billing_cycles
      end

      proration_price = OrderHelpers.duplicate_stripe_price(user, stripe_price) do |duplicated_stripe_price|
        duplicated_stripe_price.metadata[StripeForce::Translate::Metadata.metadata_key(user, MetadataKeys::PRORATION)] = true

        # since we are explicitly doing pricing math here, we should define the max
        duplicated_stripe_price.unit_amount_decimal = prorated_billing_amount.round(MAX_STRIPE_PRICE_PRECISION).to_s("F")

        # this price should be a one-time price, this must be done by removing the recurring field
        Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
          duplicated_stripe_price,
          :recurring
        )

        # TODO we should generate some sort of custom description here, but there's no place to put it
      end

      log.info 'proration price created', stripe_proration_id: proration_price.id

      proration_price
    end

    sig do
      params(
        user: StripeForce::User,
        phase_item: ContractItemStructure,
        subscription_term: Integer,
        billing_frequency: Integer
      ).returns(Hash)
    end
    def self.create_credit_price_data_from_terminated_phase_item(user:, phase_item:, subscription_term:, billing_frequency:)
      # our goal here is to identify the correct amount to credit this user, pass it into a price_data hash to be bubbled to the Subscription Item

      unless phase_item.fully_terminated?
        raise Integrations::Errors::ImpossibleState.new("attempted to create credit price_data for non-terminated line")
      end

      # since we are processing a prorated negative line item, the unused term is the same as the proration period if the line item was positive
      unused_term = subscription_term % billing_frequency
      unused_term_percentage = BigDecimal(unused_term) / BigDecimal(billing_frequency)

      # https://jira.corp.stripe.com/browse/PLATINT-1808
      if unused_term.zero?
        log.warn 'subscription term is equal to billing frequency, amendment is most likely happening on the same day'
        unused_term_percentage = 1
      end

      original_stripe_price = phase_item.price(user)
      unit_amount_decimal = BigDecimal(original_stripe_price.unit_amount_decimal)
      credit_amount = unit_amount_decimal * unused_term_percentage * -1

      # We should eventually use a Price here instead of price data. Details in ticket below
      # https://jira.corp.stripe.com/browse/PLATINT-2090
      price_data = {
        currency: original_stripe_price.currency,
        product: original_stripe_price.product,
        unit_amount_decimal: credit_amount.round(MAX_STRIPE_PRICE_PRECISION).to_s("F"),
        tax_behavior: original_stripe_price.tax_behavior,
        metadata: original_stripe_price.metadata,
      }

      log.info 'parsed credit into price_data', price_data: price_data

      price_data
    end

    sig do
      params(
        user: StripeForce::User,
        mapper: StripeForce::Mapper,
        sf_order_amendment: Restforce::SObject,
        terminated_phase_items: T::Array[ContractItemStructure],
        subscription_term: Integer,
        billing_frequency: Integer,
        is_order_backdated: T::Boolean,
        next_billing_timestamp: T.nilable(Integer),
      ).returns(T::Array[T::Hash[Symbol, T.untyped]])
    end
    # creating one-time invoice items for terminated lines for the unused prorated amount (which has already been billed)
    def self.generate_proration_credits_from_terminated_phase_items(user:, mapper:, sf_order_amendment:, terminated_phase_items:, subscription_term:, billing_frequency:, is_order_backdated:, next_billing_timestamp:)
      negative_invoice_items_for_prorations = []

      terminated_phase_items.each do |phase_item|
        unless phase_item.fully_terminated?
          log.info 'non-terminated phase_item, not creating credit for unused time'
          next
        end

        # metered prices should not create credits since we do not know how much the user is billed in CPQ
        if PriceHelpers.metered_price?(phase_item.price(user))
          log.info 'metered price, not creating credit for unused time',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price.id
          next
        end

        # right now, you can only have tiered pricing specified through consumption schedules, which are only used if the
        # price is metered billed. We may need to support prorated tiered prices in the future, if so this check should be removed
        if PriceHelpers.tiered_price?(phase_item.price)
          log.info 'tiered price, not creating credit for unused time',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        # We do not need to create credits for any items originating in this order amendment, the
        # items creating credits should have been added in a previous order and removed / terminated in this one.
        if phase_item.from_order?(sf_order_amendment)
          log.info 'line item originating from this amendment, not creating credit',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price.id
          next
        end

        log.info 'creating credit for unused time', prorated_order_item_id: phase_item.order_line_id

        # create price data, not price
        price_data = create_credit_price_data_from_terminated_phase_item(
          user: user,
          phase_item: phase_item, # https://jira.corp.stripe.com/browse/PLATINT-2090

          # TODO is there a better way to source these variables? They are required in a couple
          subscription_term: subscription_term,
          billing_frequency: billing_frequency
        )

        credit_stripe_item = Stripe::SubscriptionItem.construct_from({
          metadata: Metadata.stripe_metadata_for_sf_object(user, phase_item.order_line).merge(
            Metadata.metadata_key(user, MetadataKeys::CREDIT) => true,
            Metadata.metadata_key(user, MetadataKeys::PRORATION) => true,
          ),
        })

        mapper.apply_mapping(credit_stripe_item, phase_item.order_line)

        proration_period_start = {type: 'phase_start'}
        proration_period_end = {type: 'subscription_period_end'}
        if is_order_backdated && next_billing_timestamp.present?
          amendment_start_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, sf_order_amendment)
          proration_period_start = {type: 'timestamp', timestamp: amendment_start_date.to_i}
          proration_period_end = {type: 'timestamp', timestamp: next_billing_timestamp}

          # https://admin.corp.stripe.com/gates/billing_subscriptions_open_invoicing_interval
          # https://jira.corp.stripe.com/browse/PLATINT-2450
          if user.feature_enabled?(StripeForce::Constants::FeatureFlags::BILLING_GATE_OPEN_INVOICING_INTERVAL)
            # https://livegrep.corp.stripe.com/view/stripe-internal/pay-server/lib/subscriptions/command/invoicing_period.rb#L26
            proration_period_end[:timestamp] = proration_period_end[:timestamp] - 1 > proration_period_start[:timestamp] ? proration_period_end[:timestamp] - 1 : proration_period_end[:timestamp]
          end
        end

        negative_invoice_items_for_prorations << credit_stripe_item.to_hash.merge({
          quantity: phase_item.reduced_by,
          price_data: price_data,
          period: {
            end: proration_period_end,
            start: proration_period_start,
          },
        })
      end

      negative_invoice_items_for_prorations
    end

    sig do
      params(
        mapper: StripeForce::Mapper,
        sf_order_amendment: Restforce::SObject,
        phase_items: T::Array[ContractItemStructure],
        subscription_term: Integer,
        billing_frequency: Integer,
        backdated_billing_cycles: T.nilable(Integer),
        next_billing_timestamp: T.nilable(Integer),
        days_prorating: Integer,
      ).returns(T::Array[T::Hash[Symbol, T.untyped]])
    end
    def self.generate_proration_items_from_phase_items(mapper:, sf_order_amendment:, phase_items:, subscription_term:, billing_frequency:, backdated_billing_cycles: nil, next_billing_timestamp: nil, days_prorating: 0)
      user = mapper.user
      invoice_items_for_prorations = []

      phase_items.each do |phase_item|
        # metered prices are calculated based on user-provided usage information and therefore cannot be prorated because
        # we do not know the full billing amount ahead of time.
        # TODO I don't like passing the user here, maybe pass in the user with the contract item? Going to see how ugly this feels...
        if PriceHelpers.metered_price?(phase_item.price(user))
          log.info 'metered price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price.id
          next
        end

        # right now, you can only have tiered pricing specified through consumption schedules, which are only used if the
        # price is metered billed. We may need to support prorated tiered prices in the future.
        if PriceHelpers.tiered_price?(phase_item.price)
          log.info 'tiered price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        # if the price is not recurring, there is no need to prorate it since it will be fully billed up front
        if !PriceHelpers.recurring_price?(phase_item.price)
          log.info 'one time price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        # we only want to prorate the items that are unique to this order
        # in other words, if the item needed to be prorated it would have already happened when processing the
        # order that it was originally included on.
        if !phase_item.from_order?(sf_order_amendment)
          log.info 'line item not originated from this amendment, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price.id
          next
        end

        log.info 'prorating order item', prorated_order_item_id: phase_item.order_line_id

        proration_price = OrderAmendment.create_prorated_price_from_phase_item(
          mapper: mapper,
          phase_item: phase_item,
          subscription_term: subscription_term,
          billing_frequency: billing_frequency,
          days_prorating: days_prorating,
          backdated_billing_cycles: backdated_billing_cycles.nil? ? 0 : backdated_billing_cycles,
        )

        proration_stripe_item = Stripe::SubscriptionItem.construct_from({
          metadata: Metadata.stripe_metadata_for_sf_object(user, phase_item.order_line).merge(
            Metadata.metadata_key(user, MetadataKeys::PRORATION) => true
          ),
        })

        mapper.apply_mapping(proration_stripe_item, phase_item.order_line)

        # adjust the proration item dates if this order is backdated
        # since the amendment start date is in the past (different) compared to the current time
        proration_period_start = {type: 'phase_start'}
        proration_period_end = {type: 'subscription_period_end'}
        unless next_billing_timestamp.nil?
          amendment_start_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, sf_order_amendment)
          proration_period_start = {type: 'timestamp', timestamp: amendment_start_date.to_i}
          proration_period_end = {type: 'timestamp', timestamp: next_billing_timestamp}

          # https://admin.corp.stripe.com/gates/billing_subscriptions_open_invoicing_interval
          # https://jira.corp.stripe.com/browse/PLATINT-2450
          if user.feature_enabled?(StripeForce::Constants::FeatureFlags::BILLING_GATE_OPEN_INVOICING_INTERVAL)
            # https://livegrep.corp.stripe.com/view/stripe-internal/pay-server/lib/subscriptions/command/invoicing_period.rb#L26
            proration_period_end[:timestamp] = proration_period_end[:timestamp] - 1 > proration_period_start[:timestamp] ? proration_period_end[:timestamp] - 1 : proration_period_end[:timestamp]
          end
        end

        invoice_items_for_prorations << proration_stripe_item.to_hash.merge({
          quantity: phase_item.quantity,
          price: proration_price.id,
          period: {
            end: proration_period_end,
            start: proration_period_start,
          },
        })
      end

      invoice_items_for_prorations
    end

    sig { params(user: StripeForce::User, aggregate_phase_items: T::Array[ContractItemStructure]).returns(Integer) }
    def self.calculate_billing_frequency_from_phase_items(user, aggregate_phase_items)
      # TODO maybe we add this stuff to the contract stucture? Feels weird to calculate it way down here
      # NOTE right now, all prices on an object must have the exact same billing cycle, this may change in the future
      stripe_price_id = T.must(aggregate_phase_items.first).stripe_params[:price]
      stripe_price = Stripe::Price.retrieve(stripe_price_id, user.stripe_credentials)
      billing_frequency_in_months = StripeForce::Utilities::StripeUtil.billing_frequency_of_price_in_months(stripe_price)
      billing_frequency_in_months
    end

    # used to help determined if a order is prorated after all 'easier' checks have failed to
    # all dates returned are middate utc
    sig { params(user: StripeForce::User, original_subscription_schedule: Stripe::SubscriptionSchedule, billing_frequency: Integer).returns(T::Array[Integer]) }
    def self.calculate_billing_cycle_dates(user, original_subscription_schedule, billing_frequency)
      future_billing_dates = []

      # https://jira.corp.stripe.com/browse/PLATINT-1809
      # the start date of the subscription could be in the future and we may not be able to calculate the upcoming invoice
      if original_subscription_schedule.subscription
        subscription_id = T.cast(original_subscription_schedule.subscription, String)

        begin
          upcoming_invoice = Stripe::Invoice.upcoming({
            subscription: subscription_id,
          }, user.stripe_credentials)

          # all timestamps from stripe are in utc
          next_billing_timestamp = upcoming_invoice.created
          future_billing_dates << next_billing_timestamp
        rescue Stripe::InvalidRequestError => e
          # TODO should report there are two errors for the same thing
          if !%w{invoice_upcoming_none no_upcoming_exception}.include?(e.code)
            raise
          end

          # most likely due to there only being a single invoice in this sub phase which has alread been billed
          log.info 'upcoming invoice api call failed'
        end
      end

      if !next_billing_timestamp
        # if we can't get the next_billing_timestamp from the stripe api, then we'll have to use
        # the start date of the first phase
        next_billing_timestamp = T.must(original_subscription_schedule.phases.first).start_date

        log.info 'upcoming invoice api could not be used, using phase start date for billing cycle', start_date: next_billing_timestamp

        # in this case, we don't know if this timestamp is in the past, let's check before adding it to the list
        customer_id = T.cast(original_subscription_schedule.customer, String)
        if next_billing_timestamp > determine_current_time(user, customer_id)
          future_billing_dates << next_billing_timestamp
        end
      end

      # we don't know the state of the subscription schedule when it is passed in
      # it could be mutated by upstream logic and not yet passed to stripe
      # this is important because we need the end date of the subscription in order to know
      # when to stop calculating the next future billing dates. Since we assume all contracts
      # co-term, we can safely rely on the end date of the last phase being the date we need to calculate to.
      # in order to make sure the sub schedule has the correct ending date, we need to pull it again.
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(original_subscription_schedule.id, user.stripe_credentials)
      last_phase = T.must(subscription_schedule.phases.last)
      final_billing_timestamp = T.must(last_phase.end_date)

      # TODO file a JIRA against the billing team to fix this!
      # NOTE this is going to be a very risky operation, but there's nothing we can do right now
      # add the billing_frequency until we are past the last billing date

      next_billing_date = Time.at(next_billing_timestamp).utc.to_datetime
      is_next_billing_date_eom = StripeForce::Translate::OrderHelpers.is_order_date_eom?(next_billing_date)

      billing_date = next_billing_date
      while billing_date.to_i < final_billing_timestamp
        billing_date += billing_frequency.months

        if is_next_billing_date_eom
          # normalizes the billing date day to the end of month in the case where the final billing timestamp occurs on day-of-month that does not exist
          # in the billing date month
          days_in_month = Date.new(billing_date.year, billing_date.month, -1).day
          normalized_billing_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: billing_date.to_time, anchor_day_of_month: days_in_month)
          future_billing_dates << normalized_billing_date.to_i
        else
          future_billing_dates << billing_date.to_i
        end
      end

      future_billing_dates
    end

    sig do
      params(
        user: StripeForce::User,
        aggregate_phase_items: T::Array[ContractItemStructure],
        subscription_schedule: Stripe::SubscriptionSchedule,
        subscription_term: Integer,
        billing_frequency: Integer,
        # as unix timestamp
        amendment_start_date: Integer
      ).returns(T::Boolean)
    end
    def self.prorated_amendment?(user:, aggregate_phase_items:, subscription_schedule:, subscription_term:, billing_frequency:, amendment_start_date:)
      log.info 'determining if amendment order is prorated'

      if aggregate_phase_items.empty?
        log.info 'no subscription items, cannot be prorated order'
        return false
      end

      # if the subscription term does not match the billing frequency of the stripe item, then there will be some proration
      if (subscription_term % billing_frequency) != 0
        log.info 'billing frequency is not divisible by subscription term, assuming prorated order',
          subscription_term: subscription_term,
          billing_frequency: billing_frequency
        return true
      end

      # is the next billing date aligned with the start of the subscription? if not, there will be prorations
      billing_dates = calculate_billing_cycle_dates(user, subscription_schedule, billing_frequency)

      # we only care about the date, not the time
      # TODO we need to be very careful about date comparison, there's a good chance a nuance here will cause us issues
      if billing_dates.none? {|d| d == amendment_start_date }
        log.info 'start date is not on the next or future billing dates',
          amendment_start_date: amendment_start_date,
          billing_dates: billing_dates
        return true
      end

      false
    end

    sig do
      params(
        user: StripeForce::User,
        stripe_customer_id: String,
        original_phases: T::Array[Stripe::SubscriptionSchedulePhase]
      ).returns(T::Array[Stripe::SubscriptionSchedulePhase])
    end
    def self.delete_past_phases(user, stripe_customer_id, original_phases)
      phases = Integrations::Utilities::StripeUtil.deep_copy(original_phases)
      current_timestamp = determine_current_time(user, stripe_customer_id)

      phases.reject! do |phase|
        if phase.end_date == 'now'
          next false
        end

        in_past = phase.end_date < current_timestamp
        if in_past
          log.info 'removing completed phase', end_date: phase.end_date
        end

        in_past
      end

      phases
    end

    sig { params(user: StripeForce::User, stripe_customer_id: String).returns(Integer) }
    def self.determine_current_time(user, stripe_customer_id)
      stripe_customer = Stripe::Customer.retrieve({id: stripe_customer_id, expand: %w{test_clock}}, user.stripe_credentials)

      if !stripe_customer.test_clock
        return Time.now.utc.to_i
      end

      test_clock = T.cast(stripe_customer.test_clock, Stripe::TestHelpers::TestClock)

      if test_clock.status != "ready"
        raise StripeForce::Errors::RawUserError.new("Test clock still advancing, scheduling a retry.")
      end

      log.info 'Found test clock for customer, using frozen time', stripe_customer: stripe_customer_id, frozen_time: test_clock.frozen_time
      test_clock.frozen_time
    end
  end
end
