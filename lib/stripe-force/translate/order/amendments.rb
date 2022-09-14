# frozen_string_literal: true
# typed: true

# using `Order::Amendment` conflicts will too many other things
class StripeForce::Translate
  module OrderAmendment
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    # for now, let's not support non-co-terminated contracts
    sig { params(contract_structure: ContractStructure).returns(T::Boolean) }
    def self.contract_co_terminated?(contract_structure)
      target_end_date = calculate_order_end_date(contract_structure.initial)

      contract_structure.amendments.all? do |sf_order|
        calculate_order_end_date(sf_order) == target_end_date
      end
    end

    def self.calculate_order_end_date(sf_order)
      # get start date
      # add subscription term
    end

    # NOTE at this point it's assumed that the price is NOT a metered billing item or tiered price
    def self.create_prorated_price_from_phase_item(user:, phase_item:, subscription_term:, billing_frequency:)
      # at this point, we know what the billing amount is per billing cycle, since that has alread been defined
      # on the price object at the order line. We calculate the percentage of the original line price that should
      # be billed and multiply that against the decimal price on the stripe price.

      # TODO validate that it isn't a tiered or metered billing price

      # TODO we'll need to have some sort of logic for backend prorations to calculate the amount before
      #      a billing cycle that needs to be billed for, but let's deal with that later...
      prorated_subscription_term = subscription_term % billing_frequency
      proration_percentage = BigDecimal(prorated_subscription_term) / BigDecimal(billing_frequency)
      stripe_price = phase_item.price(user)
      unit_amount_decimal = BigDecimal(stripe_price.unit_amount_decimal)
      prorated_billing_amount = unit_amount_decimal * proration_percentage

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
        sf_order_amendment: Restforce::SObject,
        phase_items: T::Array[ContractItemStructure],
        subscription_term: Integer,
        billing_frequency: Integer
      ).returns(T::Array[T::Hash[Symbol, T.untyped]])
    end
    def self.generate_proration_items_from_phase_items(user:, sf_order_amendment:, phase_items:, subscription_term:, billing_frequency:)
      invoice_items_for_prorations = []

      phase_items.each do |phase_item|
        # TODO I don't like passing the user here, maybe pass in the user with the contract item? Going to see how ugly this feels...
        if PriceHelpers.metered_price?(phase_item.price(user))
          log.info 'metered price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        if PriceHelpers.tiered_price?(phase_item.price)
          log.info 'tiered price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        if !PriceHelpers.recurring_price?(phase_item.price)
          log.info 'one time price, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price
          next
        end

        # we only want to prorate the items that are unique to this order
        if !phase_item.from_order?(sf_order_amendment)
          log.info 'line item not originated from this amendment, not prorating',
            prorated_order_item_id: phase_item.order_line_id,
            price_id: phase_item.price.id
          next
        end

        log.info 'prorating order item', prorated_order_item_id: phase_item.order_line_id

        proration_price = OrderAmendment.create_prorated_price_from_phase_item(
          user: user,
          phase_item: phase_item,

          # TODO is there a better way to source these variables? They are required in a couple
          subscription_term: subscription_term,
          billing_frequency: billing_frequency
        )

        proration_stripe_item = Stripe::SubscriptionItem.construct_from({
          metadata: Metadata.stripe_metadata_for_sf_object(user, phase_item.order_line).merge(
            Metadata.metadata_key(user, MetadataKeys::PRORATION) => true
          ),
        })

        StripeForce::Mapper.apply_mapping(user, proration_stripe_item, phase_item.order_line)

        invoice_items_for_prorations << proration_stripe_item.to_hash.merge({
          quantity: phase_item.quantity,
          price: proration_price.id,
          period: {
            end: {
              type: 'phase_end',
            },
            start: {
              type: 'phase_start',
            },
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

    # all dates returned are middate utc
    sig { params(user: StripeForce::User, original_subscription_schedule: Stripe::SubscriptionSchedule, billing_frequency: Integer).returns(T::Array[DateTime]) }
    def self.calculate_billing_cycle_dates(user, original_subscription_schedule, billing_frequency)
      subscription_id = T.cast(original_subscription_schedule.subscription, String)

      # TODO hide these API errors
      upcoming_invoice = Stripe::Invoice.upcoming({
        subscription: subscription_id,
      }, user.stripe_credentials)

      # all timestamps are in utc
      next_billing_timestamp = upcoming_invoice.created

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
      future_billing_dates = [next_billing_timestamp]

      billing_date = next_billing_date
      while billing_date.to_i <= final_billing_timestamp
        billing_date += billing_frequency.months
        future_billing_dates << billing_date.to_i
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
      if aggregate_phase_items.empty?
        log.info 'no subscription items, cannot be a proration'
        return false
      end

      # if the subscription term does not match the billing frequency of the stripe item, then there will be some proration
      if (subscription_term % billing_frequency) != 0
        log.info 'billing frequency is not divisible by subscription term, assuming proration',
          subscription_term: subscription_term,
          billing_frequency: billing_frequency
        return true
      end

      # is the next billing date aligned with the start of the subscription? if not, there will be prorations
      billing_dates = calculate_billing_cycle_dates(user, subscription_schedule, billing_frequency)

      # we only care about the date, not the time
      # TODO we need to be very careful about date comparison, there's a good chance a nuance here will cause us issues
      if billing_dates.none? {|d| d == amendment_start_date }
        # TODO need to do more thinking here, but I think this specific case may be impossible since we are enforcing coterm
        #      let's track this and then possibly remove this codepath in the future
        Integrations::ErrorContext.report_edge_case("start date is not on the next or future billing dates")

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
        raise StripeForce::Errors::RawUserError.new("Test clock still advancing, scheduling a retry")
      end

      test_clock.frozen_time
    end
  end
end
