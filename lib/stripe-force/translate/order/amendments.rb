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
      stripe_price = phase_item.price(@user)
      unit_amount_decimal = BigDecimal(stripe_price.unit_amount_decimal)
      prorated_billing_amount = unit_amount_decimal * proration_percentage

      proration_price = OrderHelpers.duplicate_stripe_price(user, stripe_price) do |duplicated_stripe_price|
        duplicated_stripe_price.metadata[StripeForce::Utilities::Metadata.metadata_key(user, MetadataKeys::PRORATION)] = true
        # since we are explicitly doing pricing math here, we should define the max
        duplicated_stripe_price.unit_amount_decimal = prorated_billing_amount.to_s("#{MAX_STRIPE_PRICE_PRECISION}F")

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

      if aggregate_phase_items.all?(&:new_order_line?)
        log.info "all order lines are new, cannot be prorated order"
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
  end
end
