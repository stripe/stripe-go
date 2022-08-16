# frozen_string_literal: true
# typed: true

module StripeForce::Utilities
  module Metadata
    extend T::Sig
    include StripeForce::Constants

    def stripe_metadata_for_sf_object(sf_object)
      {
        sf_object_metadata_key(sf_object) => sf_object.Id,
        sf_object_metadata_link(sf_object) => "#{@user.sf_endpoint}/#{sf_object.Id}",
      }
    end

    def sf_object_metadata_key(sf_object)
      metadata_key(sf_object_metadata_name(sf_object)) + '_id'
    end

    def sf_object_metadata_link(sf_object)
      metadata_key(sf_object_metadata_name(sf_object)) + '_link'
    end

    def sf_object_metadata_name(sf_object)
      # ex: `SBQQ__OrderItemConsumptionSchedule__c`
      sf_object
        .sobject_type
        .underscore
        # remove __c on custom objects to save key size
        .gsub(/__c$/, '')
        # remove sbqq namespace
        .gsub('sbqq__', '')
        # if we are mapping a subfield on order_item, it will be long
        .gsub('order_item_', 'oi_')
    end

    # TODO we should use the pattern more, so we can reference the pure-functional methods directly
    module_function :sf_object_metadata_name

    def metadata_key(key)
      "#{@user.metadata_prefix}#{key}"
    end

    sig { params(user: StripeForce::User, key: String).returns(String) }
    def self.metadata_key(user, key)
      "#{user.metadata_prefix}#{key}"
    end
  end
end
