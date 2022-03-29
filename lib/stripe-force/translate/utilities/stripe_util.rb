# frozen_string_literal: true
# typed: true

module StripeForce::Utilities
  # terrible to suffix this with `Util` but saved us namespace issues
  module StripeUtil
    extend T::Sig
    extend T::Helpers
    include Kernel

    abstract!

    include Integrations::Log
    include Integrations::Utilities::StripeUtil

    # assume incoming ID could be anything! It's essentially free-form text from the user
    def stripe_object_from_id(stripe_object_id)
      stripe_transaction_class = stripe_class_from_id(stripe_object_id, raise_on_missing: false)
      stripe_retrieve_params = {id: stripe_object_id}

      if stripe_transaction_class.nil?
        log.error 'no valid stripe object detected', invalid_value: stripe_object_id
        return
      end

      unless [Stripe::SubscriptionSchedule, Stripe::Invoice].include?(stripe_transaction_class)
        log.error 'invalid stripe class provided', invalid_class: stripe_transaction_class
        return
      end

      begin
        stripe_transaction = stripe_transaction_class.retrieve(stripe_retrieve_params, @user.stripe_credentials)
      rescue Stripe::InvalidRequestError => e
        normalized_error_message = e.message&.downcase
        if ['no such payment_intent', 'no such charge'].any? {|m| normalized_error_message&.include?(m) }
          log.error 'invalid charge or payment_intent', stripe_object_id: stripe_object_id
          return
        end

        raise(e)
      end

      stripe_transaction
    end
  end
end
