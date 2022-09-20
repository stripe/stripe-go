# frozen_string_literal: true
# typed: true

module StripeForce::Utilities
  # terrible to suffix this with `Util` but saved us namespace issues
  module StripeUtil
    extend T::Sig
    extend T::Helpers

    include Kernel
    include StripeForce::Constants

    abstract!

    # SF sub terms are generally in month terms
    sig { params(stripe_price: Stripe::Price).returns(Integer) }
    def self.billing_frequency_of_price_in_months(stripe_price)
      # TODO remove assumption that everything is monthly
      if %w{week day}.include?(stripe_price.recurring.interval)
        raise Integrations::Errors::UnhandledEdgeCase.new("unsupported price interval")
      end

      interval_in_months = case stripe_price.recurring.interval
      when 'month'
        1
      when 'year'
        12
      else
        raise Integrations::Errors::UnhandledEdgeCase.new("unexpected stripe pricing interval")
      end

      stripe_price.recurring.interval_count * interval_in_months
    end

    sig { params(user: StripeForce::User, sf_object: Restforce::SObject, action: T.nilable(Symbol)).returns(Hash) }
    def self.generate_idempotency_key_with_credentials(user, sf_object, action=nil)
      mappings_as_string = user.field_mappings.to_json + user.field_mappings.to_json
      mapping_uid = Digest::SHA1.hexdigest(mappings_as_string)

      if user.sandbox?
        # Skip idempotency keys in sandbox. Using an idemptotency key is useful to ensure we do not create duplicate
        # subscriptions but also causes problems when for example a request fails because of a mapping issue
        # (ie days_until_due on invoice settings is not being set) and upon fixing the mapping via the UI and retrying
        # it will fail due to duplicate idempotency key with a different request body.

        # We cannot just add the modified date to the key as that would not have changed in the example above.
        return user.stripe_credentials
      end

      # TODO if the SF object is mutated in a way which changes inputs, we need a new idempotency key, maybe use created date?
      # TODO feature flag to turn this off

      key = "#{mapping_uid}-#{sf_object[SF_ID]}"

      if action
        key = "#{key}-#{action}"
      end

      user.stripe_credentials.merge({idempotency_key: key})
    end

  end
end
