# frozen_string_literal: true
# typed: true

module StripeForce
  class Sanitizer
    extend T::Sig
    include Integrations::Log

    sig { params(user: StripeForce::User).void }
    def initialize(user)
      @user = user
    end

    def sanitize(stripe_record)
      if stripe_record.is_a?(Stripe::Customer)
        sanitize_customer(stripe_record)
      end
    end

    private def sanitize_customer(stripe_customer)
      if stripe_customer[:description]
        stripe_customer[:description] = stripe_customer[:description][0..349]
      end
    end
  end
end
