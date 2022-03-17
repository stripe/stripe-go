# frozen_string_literal: true
# typed: true

module StripeForce::Utilities
  module Metadata
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
      sf_object.sobject_type.underscore
    end

    def metadata_key(key)
      "#{@user.metadata_prefix}#{key}"
    end
  end
end
