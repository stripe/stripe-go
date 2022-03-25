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

    sig { params(field_name: String).returns(String) }
    def prefixed_stripe_field(field_name)
      custom_field_prefix = case (salesforce_namespace = @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE])
      when nil
        report_edge_case("expected namespace to be defined, using fallback")
        ""
      when "c"
        ""
      else
        # `stripeConnector_Stripe_ID__c` is the expected field name
        salesforce_namespace + "__"
      end

      custom_field_prefix + field_name
    end
  end
end
