# frozen_string_literal: true
# typed: true

module StripeForce::Utilities
  # terrible to suffix this with `Util` but saved us namespace issues
  module StripeUtil
    extend T::Sig
    extend T::Helpers
    include Kernel

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
  end
end
