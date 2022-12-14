# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module Metadata
    extend T::Sig

    sig { params(sf_object: Restforce::SObject).returns(String) }
    def self.sf_object_metadata_name(sf_object)
      # ex: `SBQQ__OrderItemConsumptionSchedule__c`
      key = sf_object
        .sobject_type
        .underscore
        # remove __c on custom objects to save key size
        .gsub(/__c$/, '')
        # remove sbqq namespace
        .gsub('sbqq__', '')
        # if we are mapping a subfield on order_item, it will be long
        .gsub('order_item_', 'oi_')

      # TODO refactor to be more clear where these strings are coming from
      # namespace_prefixes = [StripeForce::Constants::SalesforceNamespaceOptions::QA, StripeForce::Constants::SalesforceNamespaceOptions::PRODUCTION].map(&:serialize).map(&:underscore)
      if key.include? "qa_stripe_connect"
        key.slice! "qa_stripe_connect__"
      elsif key.include? "stripe_connector"
        key.slice! "stripe_connector__"
      end

      key
    end

    sig { params(user: StripeForce::User, key: T.any(T::Enum, String)).returns(String) }
    def self.metadata_key(user, key)
      if key.is_a?(T::Enum)
        key = key.serialize
      end

      "#{user.metadata_prefix}#{key}"
    end

    sig { params(user: StripeForce::User, sf_object: Restforce::SObject).returns(T::Hash[String, String]) }
    def self.stripe_metadata_for_sf_object(user, sf_object)
      {
        sf_object_metadata_key(user, sf_object) => sf_object.Id,
        sf_object_metadata_link(user, sf_object) => "#{user.sf_endpoint}/#{sf_object.Id}",
      }
    end

    sig { params(user: StripeForce::User, sf_object: Restforce::SObject).returns(String) }
    def self.sf_object_metadata_key(user, sf_object)
      metadata_key(user, sf_object_metadata_name(sf_object)) + '_id'
    end

    sig { params(user: StripeForce::User, sf_object: Restforce::SObject).returns(String) }
    def self.sf_object_metadata_link(user, sf_object)
      metadata_key(user, sf_object_metadata_name(sf_object)) + '_link'
    end
  end
end
