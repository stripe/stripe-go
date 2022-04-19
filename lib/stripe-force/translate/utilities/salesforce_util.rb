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
      when /^802/
        SF_ORDER_ITEM
      else
        object_prefix = sf_id[0..2]
        object_description = @user.sf_client.api_get('sobjects').body["sobjects"].detect {|o| o["keyPrefix"] == object_prefix }

        # => {"activateable"=>false,
        #   "associateEntityType"=>nil,
        #   "associateParentEntity"=>nil,
        #   "createable"=>true,
        #   "custom"=>true,
        #   "customSetting"=>false,
        #   "deepCloneable"=>false,
        #   "deletable"=>true,
        #   "deprecatedAndHidden"=>false,
        #   "feedEnabled"=>false,
        #   "hasSubtypes"=>false,
        #   "isInterface"=>false,
        #   "isSubtype"=>false,
        #   "keyPrefix"=>"a0z",
        #   "label"=>"Quote",
        #   "labelPlural"=>"Quotes",
        #   "layoutable"=>true,
        #   "mergeable"=>false,
        #   "mruEnabled"=>true,
        #   "name"=>"SBQQ__Quote__c",
        #   "queryable"=>true,
        #   "replicateable"=>true,
        #   "retrieveable"=>true,
        #   "searchable"=>true,
        #   "triggerable"=>true,
        #   "undeletable"=>true,
        #   "updateable"=>true,
        #   "urls"=>
        #    {"compactLayouts"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/describe/compactLayouts",
        #     "rowTemplate"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/{ID}",
        #     "approvalLayouts"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/describe/approvalLayouts",
        #     "describe"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/describe",
        #     "quickActions"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/quickActions",
        #     "layouts"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c/describe/layouts",
        #     "sobject"=>"/services/data/v52.0/sobjects/SBQQ__Quote__c"}}

        # https://help.salesforce.com/s/articleView?id=000325244&type=1
        if !object_description
          raise "unknown object type #{sf_id}"
        end

        # TODO https://jira.corp.stripe.com/browse/PLATINT-1536
        object_description["name"]
      end
    end

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
