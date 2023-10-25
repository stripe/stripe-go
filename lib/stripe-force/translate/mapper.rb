# frozen_string_literal: true
# typed: true

module StripeForce
  class Mapper
    extend T::Sig

    include Integrations::Log
    include Integrations::ErrorContext
    include Integrations::Utilities
    include Integrations::Utilities::StripeUtil

    include StripeForce::Utilities::SalesforceUtil

    attr_reader :user

    sig { params(stripe_record: T.any(Stripe::APIResource, Class), sf_record: T.nilable(Restforce::SObject)).returns(String) }
    def self.mapping_key_for_record(stripe_record, sf_record)
      stripe_record_class = if stripe_record.is_a?(Class) && stripe_record < Stripe::APIResource
        stripe_record
      else
        stripe_record.class
      end

      # Stripe::Charge => charge; InvoiceItem => invoice_item
      # these are the keys used in all of the mapping hashes
      stripe_record_key = stripe_record_class.to_s.demodulize.underscore

      compound_key = stripe_record_class == Stripe::Price && sf_record&.sobject_type == SF_ORDER_ITEM

      if compound_key
        log.info 'using compound key for price order item',
        sf_record = T.must(sf_record)
        stripe_record_key + "_" + Translate::Metadata.sf_object_metadata_name(sf_record)
      else
        stripe_record_key
      end
    end

    sig { params(u: StripeForce::User, cache_service: CacheService).void }
    def initialize(u, cache_service)
      @user = u
      @cache_service = cache_service
    end

    sig { params(record: T.any(Stripe::APIResource, Restforce::SObject), key_path: String).returns(T.nilable(T.any(String, Integer, Float, T::Boolean))) }
    def extract_key_path_for_record(record, key_path)
      if record.is_a?(Stripe::APIResource)
        extract_stripe_resource_field(record, key_path)
      elsif record.is_a?(Restforce::SObject)
        extract_salesforce_object_field(record, key_path)
      end
    end

    def extract_salesforce_object_field(sf_object, key_path)
      components = key_path.split('.').map(&:strip)
      target_object = sf_object

      components.each_with_index do |field_name, i|
        is_last_component = i == components.size - 1

        normalized_field_name = if !is_last_component && !field_name.end_with?('__c') && !field_name.end_with?('Id')
          field_name + "Id"
        else
          field_name
        end

        if target_object.key?(normalized_field_name)
          target_object = target_object[normalized_field_name]
        else
          log.debug 'field does not exist',
            field_component: field_name,
            normalized_field_component: normalized_field_name,
            field_path: key_path,
            target_object: target_object.sobject_type,
            target_object_id: target_object.Id

          target_object = nil
        end

        if target_object.nil?
          log.debug 'field value is nil',
            field_component: field_name,
            field_path: key_path,
            root_object: sf_object.sobject_type,
            root_object_id: sf_object.Id
          break
        end

        # if we aren't at the last component in the key path, sniff for Stripe object references
        # object references are always string, so we can ignore any other object types

        if !is_last_component
          target_class = StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, target_object)

          if target_class
            target_object = @cache_service.silent { @cache_service.get_record_from_cache(target_class, target_object) }
          end
        end
      end

      target_object
    end

    # TODO compound_key is unsused here https://jira.corp.stripe.com/browse/PLATINT-2597
    sig { params(record_to_map: Stripe::APIResource, source_record: T.nilable(Restforce::SObject), compound_key: T.nilable(T::Boolean)).void }
    def apply_mapping(record_to_map, source_record=nil, compound_key: false)
      record_to_map_key = self.class.mapping_key_for_record(record_to_map, source_record)

      # field defaults (static values)
      field_defaults_for_record = @user.field_defaults[record_to_map_key]
      if field_defaults_for_record
        assign_values_from_hash(record_to_map, field_defaults_for_record)
      end

      # field mappings (dynamic values)
      mapping_for_record = @user.field_mappings[record_to_map_key]
      if mapping_for_record && source_record
        mapping_values = build_dynamic_mapping_values(source_record, mapping_for_record)
        assign_values_from_hash(record_to_map, mapping_values)
      end
    end

    sig { params(source_record: Restforce::SObject, mapping_for_record: Hash).returns(Hash) }
    def build_dynamic_mapping_values(source_record, mapping_for_record)
      # build annotations using stripe resource metadata and users table
      # of accepted stripe metadata values and their corresponding NS keys
      mapping_for_record.each_with_object({}) do |(destination_field_name, source_field_path), assignments|
        next if source_field_path.to_s.empty?

        extracted_value = extract_key_path_for_record(source_record, source_field_path)

        log.debug 'extracted value for mapping',
          source_field_path: source_field_path,
          destination_field_name: destination_field_name,
          extracted_value: extracted_value

        next unless extracted_value

        # TODO should we only do this if we run into a specific field suffix?
        #      I feel like this is good default behavior. When would someone NOT want this enabled?

        # convert true/false strings into boolean so the correct type is carried over to the Stripe API
        if extracted_value.to_s.downcase.strip == 'true'
          extracted_value = true
        elsif extracted_value.to_s.downcase.strip == 'false'
          extracted_value = false
        end

        assignments[destination_field_name] = extracted_value
      end
    end

    sig { params(stripe_object: Stripe::StripeObject, sf_record: Restforce::SObject).returns(Stripe::StripeObject) }
    def add_termination_metadata(stripe_object, sf_record)
      sf_record_type = sf_record.sobject_type
      if sf_record_type == SF_ORDER
        key = "subscription_schedule"
      elsif sf_record_type == SF_ORDER_ITEM
        key = "subscription_item"
      else
        raise Integrations::Errors::ImpossibleState.new("Termination metadata is not supported for this sf object type: #{sf_record_type}")
      end

      if self.user.field_defaults[key]
        assign_values_from_hash(stripe_object, self.user.field_defaults[key], only_termination_metadata: true)
      end

      if self.user.field_mappings[key]
        mapping_values = build_dynamic_mapping_values(sf_record, self.user.field_mappings[key])
        assign_values_from_hash(stripe_object, mapping_values, only_termination_metadata: true)
      end

      stripe_object
    end

    sig { params(record: Stripe::StripeObject, field_assignments: Hash, only_termination_metadata: T.nilable(T::Boolean),).void }
    def assign_values_from_hash(record, field_assignments, only_termination_metadata: false)
      field_assignments.each do |raw_field_path, raw_field_value|
        # if only_termination_metadata is true, we are only mapping termination metadata
        if only_termination_metadata
          if raw_field_path.to_s.start_with?(TERMINATION_METADATA_PREFIX)
            set_stripe_resource_field_path(record, raw_field_path.to_s.delete_prefix(TERMINATION_METADATA_PREFIX), raw_field_value)
          end

          next
        end

        # some fields require processing that we do in the object translation directly
        skip_processing = raw_field_path.to_s.start_with?(TERMINATION_METADATA_PREFIX)
        if skip_processing
          next
        end

        set_stripe_resource_field_path(record, raw_field_path, raw_field_value)
      end
    end
  end
end
