# typed: true
# frozen_string_literal: true

module Integrations::Utilities::Stripe
  extend T::Helpers
  extend T::Sig

  abstract!

  include Integrations::Log
  include Integrations::ErrorContext

  sig { params(stripe_resource: Stripe::APIResource, field_path: String).returns(T.untyped) }
  def extract_stripe_resource_field(stripe_resource, field_path)
    components = field_path.split('.')
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
