# frozen_string_literal: true
# typed: true

module StripeForce
  module Sanitizer
    extend T::Sig
    include Integrations::Log
    include StripeForce::Constants

    def self.sanitize(stripe_record)
      if stripe_record.is_a?(Stripe::Customer)
        sanitize_customer(stripe_record)
      end

      if stripe_record.is_a?(Stripe::Price)
        sanitize_price(stripe_record)
      end

      # Users may send iterations over as a decimal, we will round down.
      # https://jira.corp.stripe.com/browse/PLATINT-2050
      if stripe_record.is_a?(Stripe::SubscriptionSchedule)
        sanitize_subscription_schedule(stripe_record)
      end

      if stripe_record.is_a?(Stripe::Subscription)
        # TODO: do something maybe? this is the case of old pre-integration orders
        # also the case of evergreen orders
        # https://jira.corp.stripe.com/browse/PLATINT-2538
      end
    end

    private_class_method def self.sanitize_subscription_schedule(stripe_subscription_schedule)
      if stripe_subscription_schedule[:prebilling] && stripe_subscription_schedule[:prebilling][:iterations]
        stripe_subscription_schedule[:prebilling][:iterations] = stripe_subscription_schedule[:prebilling][:iterations].to_i

        # Passing prebilling.iterations = '0' results in a Stripe::InvalidRequestError since the value must be greater than or equal to 1,
        # but we want to enable users to map to a value of '0' (instead of nil) to represent no prebilling
        if stripe_subscription_schedule[:prebilling][:iterations] == 0
          stripe_subscription_schedule[:prebilling][:iterations] = nil
        end
      end
    end

    private_class_method def self.sanitize_price(stripe_price)
      stripe_price_unit_amount_decimal = stripe_price[:unit_amount_decimal]
      if stripe_price_unit_amount_decimal && !Integrations::Utilities::StripeUtil.is_integer_value?(stripe_price_unit_amount_decimal)
        # if a string, then we want to convert to a precision float for careful comparison
        if stripe_price_unit_amount_decimal.is_a?(String)
          stripe_price_unit_amount_decimal = BigDecimal(stripe_price_unit_amount_decimal)
        end

        # Stripe only supports 12 digits
        stripe_price[:unit_amount_decimal] = stripe_price_unit_amount_decimal.round(MAX_STRIPE_PRICE_PRECISION)
      end
    end

    private_class_method def self.sanitize_customer(stripe_customer)
      if stripe_customer[:description]
        stripe_customer[:description] = stripe_customer[:description][0..349]
      end

      if stripe_customer[:shipping] && stripe_customer[:shipping][:phone]
        stripe_customer[:shipping][:phone] = stripe_customer[:shipping][:phone][0..19]
      end

      # passing a partial shipping hash will trigger an error, remove the shipping hash entirely if it's only partial
      if !stripe_customer.shipping.respond_to?(:address) || stripe_customer.shipping.address.to_h.empty?
        log.info 'no address on shipping hash, removing'
        stripe_customer.shipping = {}
      end

      # currently, the connector only supports one custom field but we should remove this once we support multiple
      if stripe_customer[:invoice_settings] && stripe_customer[:invoice_settings][:custom_fields]
        stripe_customer[:invoice_settings][:custom_fields] = [stripe_customer[:invoice_settings][:custom_fields]]
      end
    end
  end
end
