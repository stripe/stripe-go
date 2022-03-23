# typed: true
# frozen_string_literal: true

module Critic
  module StripeFactory
    extend T::Sig
    extend T::Helpers
    abstract!

    include Kernel
    include StripeForce::Constants

    def create_stripe_id(prefix)
      # NOTE: The number after the underscore has significance for Stripe's internal routing.
      #   While we don't expect these IDs to be used for real API calls, we want to ensure
      #   they don't lead to unexpected behavior if they are.
      random_id = "_1" + SecureRandom.alphanumeric(29)

      if ENV['CIRCLE_NODE_INDEX']
        random_id = "#{random_id}#{ENV['CIRCLE_NODE_INDEX']}"
      end

      prefix.to_s + random_id
    end
  end
end
