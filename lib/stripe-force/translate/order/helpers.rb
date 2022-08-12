# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module OrderHelpers
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    sig { params(raw_days_until_due: T.any(String, Integer)).returns(Integer) }
    def self.transform_payment_terms_to_days_until_due(raw_days_until_due)
      if raw_days_until_due.is_a?(Integer)
        return raw_days_until_due
      end

      # TODO it is possible for users to customize the options here, we may need to use regex extraction or something at some point
      case raw_days_until_due
      when "Net 15"
        15
      when "Net 30"
        30
      when "Net 45"
        45
      when "Net 60"
        60
      when "Net 90"
        90
      else
        raise StripeForce::Errors::RawUserError.new("unexpected days_until_due option #{raw_days_until_due}")
      end
    end

    sig { params(user: StripeForce::User, original_stripe_price: Stripe::Price).returns(Stripe::Price) }
    def self.duplicate_stripe_price(user, original_stripe_price)
      stripe_price_params = original_stripe_price.to_hash
      stripe_price_params.delete(:id)

      # Stripe::InvalidRequestError: Received unknown parameters: type, object, livemode, created
      stripe_price_params.delete(:type)
      stripe_price_params.delete(:object)
      stripe_price_params.delete(:livemode)
      stripe_price_params.delete(:created)

      # Stripe::InvalidRequestError: You may only specify one of these parameters: unit_amount, unit_amount_decimal.
      stripe_price_params.delete(:unit_amount)

      # same error as above
      # TODO should we conditionally remove the unit amount? Should this be done in the sanitize price params instead?
      stripe_price_params[:tiers]&.each do |tier|
        tier.delete(:unit_amount)

        # the API returns nil for `inf` boundary, but API expects `inf` :(
        if tier.key?(:up_to) && tier[:up_to].nil?
          tier[:up_to] = 'inf'
        end
      end

      stripe_price_params.delete(:active)
      stripe_price = Stripe::Price.construct_from(stripe_price_params)

      if PriceHelpers.tiered_price?(stripe_price)
        stripe_price = PriceHelpers.sanitize_price_tier_params(stripe_price)
      end

      Stripe::Price.create(stripe_price.to_hash, user.stripe_credentials)
    end

    # since the input and output is "fake" stripe subhash, typing here doesn't work
    sig { params(user: StripeForce::User, original_phase_items: T::Array[ContractItemStructure]).returns(T::Array[ContractItemStructure]) }
    def self.ensure_unique_phase_item_prices(user, original_phase_items)
      phase_items = T.cast(original_phase_items.deep_dup, T::Array[ContractItemStructure])

      price_ids = []
      phase_items.each do |phase_item|
        price_id = phase_item.stripe_params[:price]

        if !price_ids.include?(price_id)
          price_ids << price_id
          next
        end

        log.debug "price id is duplicated, creating additional price"
        original_price = Stripe::Price.retrieve({id: price_id, expand: %w{tiers}}, user.stripe_credentials)
        new_price = OrderHelpers.duplicate_stripe_price(user, original_price)
        phase_item.stripe_params[:price] = new_price.id
        log.info "new price created to avoid duplicate price", new_price_id: new_price.id
      end

      phase_items
    end

    sig { params(original_phases: T::Array[Stripe::SubscriptionSchedulePhase]).returns(T::Array[Stripe::SubscriptionSchedulePhase]) }
    def self.sanitize_subscription_schedule_phase_params(original_phases)
      # without deep dupping this will mutate the input
      phases = original_phases.deep_dup

      # TODO report https://jira.corp.stripe.com/browse/PLATINT-1479
      # You can't pass back the phase in it's original format, it must be modified to avoid:
      # 'You passed an empty string for 'phases[0][collection_method]'. We assume empty values are an attempt to unset a parameter; however 'phases[0][collection_method]' cannot be unset. You should remove 'phases[0][collection_method]' from your request or supply a non-empty value.'
      phases.each do |phase|
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

      phases
    end
  end
end
