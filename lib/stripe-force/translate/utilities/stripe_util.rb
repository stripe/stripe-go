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
      last_modified_as_timestamp = DateTime.parse(sf_object[SF_LAST_MODIFIED_DATE]).to_i
      key = "#{user.mapping_hash}-#{sf_object[SF_ID]}-#{last_modified_as_timestamp}"

      if action
        key = "#{key}-#{action}"
      end

      user.stripe_credentials.merge({idempotency_key: key})
    end

  end
end
