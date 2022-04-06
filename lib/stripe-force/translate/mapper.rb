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

    sig { params(u: StripeForce::User).void }
    def initialize(u)
      @user = u
    end

    sig { params(record_or_class: T.any(Stripe::APIResource, Restforce::SObject, Class)).returns(String) }
    def mapping_key_for_record(record_or_class)
      if record_or_class.is_a?(Class) && record_or_class < Stripe::APIResource
        record_or_class.to_s.demodulize.underscore
      elsif record_or_class.is_a?(Stripe::APIResource)
        # Stripe::Charge => charge; InvoiceItem => invoice_item
        # these are the keys used in all of the mapping hashes
        record_or_class.class.to_s.demodulize.underscore
      elsif record_or_class.is_a?(Restforce::SObject)
        record_or_class.sobject_type
      end
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

        normalized_field_name = if !is_last_component && !field_name.end_with?('__c')
          field_name + "Id"
        else
          field_name
        end

        if sf_object.key?(normalized_field_name)
          target_object = target_object[normalized_field_name]
        else
          log.info 'field does not exist',
            field_component: field_name,
            field_path: key_path,
            target_object: target_object.class
          target_object = nil
        end

        if target_object.nil?
          break
        end

        # if we aren't at the last component in the key path, sniff for Stripe object references
        # object references are always string, so we can ignore any other object types

        if !is_last_component
          target_class = salesforce_type_from_id(target_object)

          if target_class
            target_object = @user.sf_client.find(target_class, target_object)
          end
        end
      end

      target_object
    end

    sig { params(record_to_map: Stripe::APIResource, source_record: T.nilable(Restforce::SObject)).void }
    def apply_mapping(record_to_map, source_record=nil)
      record_to_map_key = mapping_key_for_record(record_to_map)
      field_defaults_for_record = @user.field_defaults[record_to_map_key]

      # field defaults
      if field_defaults_for_record
        assign_values_from_hash(record_to_map, field_defaults_for_record)
      end

      mapping_for_record = @user.field_mappings[record_to_map_key]

      if mapping_for_record && source_record
        mapping_values = build_dynamic_mapping_values(source_record, mapping_for_record)
        assign_values_from_hash(record_to_map, mapping_values)
      end
    end

    protected def is_record_reference_field?(field_name)
      field_name =~ /_id$/
    end

    protected def is_date_field?(field_name)
      field_name =~ /_date$/
    end

    sig { params(source_record: Restforce::SObject, mapping_for_record: Hash).returns(Hash) }
    def build_dynamic_mapping_values(source_record, mapping_for_record)
      # build annotations using stripe resource metadata and users table
      # of accepted stripe metadata values and their corresponding NS keys
      mapping_for_record.each_with_object({}) do |(destination_field_name, source_field_path), assignments|
        next if source_field_path.to_s.empty?

        extracted_value = extract_key_path_for_record(source_record, source_field_path)

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

    sig { params(record: Stripe::APIResource, field_assignments: Hash).void }
    def assign_values_from_hash(record, field_assignments)
      field_assignments.each do |raw_field_path, v|
        # TODO need to handle nil values
        # TODO think on custom fields with NetSuite suffix `_id`
        # TODO nofify when an existing field is being overwritten

        # TODO unsure if we need to special case this stuff right now
        # if is_record_reference_field?(raw_field_path)
        #   field_path = "#{raw_field_path[0..-4]}="
        # elsif is_date_field?(raw_field_path)
        #   field_path = "#{raw_field_path[0..-6]}="
        # else
        #   field_path = raw_field_path
        # end

        field_path = raw_field_path

        if is_record_reference_field?(field_path) && record.respond_to?(field_path)
          raise 'determine what salesforce record ref transformation requirements are'
          # record.send(field_path, ref_value)
        elsif is_date_field?(field_path) && record.respond_to?(field_path)
          raise 'determine what salesforce date transformation requirements are'
          # record.send(field_path, ref_value)
        else
          set_stripe_resource_field_path(record, field_path, v)
          #   record.send(method, v)
          # else
          #   log.error 'invalid mapping field specified',
          #     mapping_key: k,
          #     mapping_value: v,
          #     record: record
        end
      end
    end

  end
end
