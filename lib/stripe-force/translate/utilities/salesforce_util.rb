# typed: true
# frozen_string_literal: true

# StripeUtil to avoid namespace madness with Stripe::* objects
module StripeForce::Utilities
  module SalesforceUtil
    extend T::Helpers
    extend T::Sig
    include Kernel

    abstract!

    include Integrations::ErrorContext
    include StripeForce::Constants

    module ClassMethods
      extend T::Sig
      extend T::Helpers
      abstract!

      include Kernel
      include StripeForce::Constants

      sig { params(sf_id: String).returns(String) }
      def salesforce_type_from_id(sf_id)
        case sf_id
        when /^01s/
          SF_PRICEBOOK
        when /^01t/
          SF_PRODUCT
        when /^01u/
          SF_PRICEBOOK_ENTRY
        when /^001/
          SF_ACCOUNT
        when /^003/
          SF_CONTACT
        when /^801/
          SF_ORDER
        when /^a0z/
          CPQ_QUOTE
        when /^802/
          SF_ORDER_ITEM
        else
          # https://help.salesforce.com/s/articleView?id=000325244&type=1
          raise "unknown object type #{sf_id}"
        end
      end
    end

    mixes_in_class_methods(ClassMethods)

    # bad pattern, but we need the purely-functional class methods copied to the instance
    include ClassMethods

    sig { params(field_name: String).returns(String) }
    def prefixed_stripe_field(field_name)
      custom_field_prefix = case (salesforce_namespace = @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE])
      when nil
        report_edge_case("expected namespace to be defined, using fallback")
        ""
      when "c"
        ""
      when *SalesforceNamespaceOptions.values.map(&:serialize)
        # `stripeConnector_Stripe_ID__c` is the expected field name
        salesforce_namespace + "__"
      else
        raise "invalid namespace provided #{custom_field_prefix}"
      end

      custom_field_prefix + field_name
    end
  end
end
