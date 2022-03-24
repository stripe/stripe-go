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

    def stripe_class_from_id(stripe_object_id, raise_on_missing: true)
      # Setting raise_on_missing to false should only be used on internal
      # support tooling.
      case stripe_object_id
      when /^re_/, /^pyr_/
        Stripe::Refund
      when /^tr_/
        Stripe::Transfer
      when /^po_/
        Stripe::Payout
      when /^ch_/, /^py_/
        Stripe::Charge
      when /^dp_/, /^pdp_/, /^du_/
        Stripe::Dispute
      when /^or_/
        Stripe::Order
      when /^in_/
        Stripe::Invoice
      # customer IDs can be specified by the user, but rarely are
      when /^cus_/
        Stripe::Customer
      when /^cbtxn_/
        Stripe::CustomerBalanceTransaction
      when /^cn_/
        Stripe::CreditNote
      when /^txn_/
        Stripe::BalanceTransaction
      when /^ii_/
        Stripe::InvoiceItem
      when /^si_/
        Stripe::SubscriptionItem
      # plan IDs are often specified by the user
      when /^plan_/
        Stripe::Plan
      when /^prod_/
        Stripe::Product
      when /^sku_/
        Stripe::SKU
      when /^seti_/
        Stripe::SetupIntent
      when /^pi_/
        Stripe::PaymentIntent
      when /^pm_/
        Stripe::PaymentMethod
      # coupons do not have a prefix since the ID is often exposed to the user
      # https://github.com/stripe/stripe-netsuite/issues/1658
      else
        raise "unknown stripe id: #{stripe_object_id}" if raise_on_missing
        nil
      end
    end

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
