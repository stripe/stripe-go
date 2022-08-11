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
