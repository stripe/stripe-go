# typed: true
# frozen_string_literal: true

# StripeUtil to avoid namespace madness with Stripe::* objects
module Integrations::Utilities::StripeUtil
  extend T::Helpers
  extend T::Sig

  abstract!

  include Integrations::Log
  include Integrations::ErrorContext

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

  sig { params(stripe_resource: Stripe::APIResource, field_path: String, field_value: T.nilable(T.any(String, Integer, T::Boolean))).void }
  def set_stripe_resource_field_path(stripe_resource, field_path, field_value)
    components = field_path.split('.').map(&:strip)
    target_object = T.let(stripe_resource, Stripe::StripeObject)

    components.each_with_index do |field_name, i|
      last_component = i == components.size - 1

      if !last_component
        if !target_object.respond_to?(field_name)
          target_object[field_name] = {}
        end

        target_object = target_object[field_name]
      else
        target_object[field_name] = field_value
      end
    end
  end

  sig { params(stripe_resource: Stripe::APIResource, field_path: String).returns(T.untyped) }
  def extract_stripe_resource_field(stripe_resource, field_path)
    components = field_path.split('.').map(&:strip)
    target_object = T.let(stripe_resource, T.nilable(Stripe::APIResource))

    components.each_with_index do |field_name, i|
      # NOTE metadata hashes don't need to be accessed using hash notation

      if target_object.respond_to?(field_name.to_sym)
        target_object = target_object.send(field_name.to_sym)
      else
        log.info 'field does not exist',
          field_component: field_name,
          field_path: field_path,
          target_object: target_object.class
        target_object = nil
      end

      if target_object.nil?
        break
      end

      # if we aren't at the last component in the key path, sniff for Stripe object references
      # object references are always string, so we can ignore any other object types

      if i != components.size - 1 && target_object.is_a?(String)
        target_class = case target_object
        when /^sub_/
          ::Stripe::Subscription
        when /^cus_/
          ::Stripe::Customer
        when /^in_/
          ::Stripe::Invoice
        when /^pi_/
          ::Stripe::PaymentIntent
        when /^prod_/
          ::Stripe::Product
        when /^(card_|src_)/
          # UPGRADE_CHECK right now, the full `source` object is included in the charge response, this may not be true in the future
          report_edge_case("card or source referenced in annotator")
        else
          # TODO report on unsupported ID if it looks like a stripe ID
          nil
        end

        if target_class
          target_object = target_class.retrieve(
            target_object,
            @user.stripe_client_credentials
          )
        end
      end
    end

    target_object
  end

end
