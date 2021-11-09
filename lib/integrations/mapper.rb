# frozen_string_literal: true
# typed: true

require_relative './utilities'

module Integrations
  class Mapper
    extend T::Sig

    include Integrations::Log
    include Integrations::ErrorContext
    include Integrations::Utilities
    include Integrations::Utilities::Stripe

    attr_reader :user

    sig { params(u: StripeForce::User).void }
    def initialize(u)
      @user = u
    end

    sig { params(record: T.any(Stripe::APIResource, Restforce::SObject)).returns(String) }
    def mapping_key_for_record(record)
      if record.is_a?(Stripe::APIResource)
        # Stripe::Charge => charge; InvoiceItem => invoice_item
        record.class.to_s.demodulize.underscore
      elsif record.is_a?(Restforce::SObject)
        record.sobject_type
      end
    end

    sig { params(record: T.any(Stripe::APIResource, Restforce::SObject), key_path: String).returns(T.nilable(String)) }
    def extract_key_path_for_record(record, key_path)
      if record.is_a?(Stripe::APIResource)
        extract_stripe_resource_field(record, key_path)
      elsif record.is_a?(Restforce::SObject)
        raise 'dot path not supported! Implement!' if key_path.include?('.')
        record[key_path]
      end
    end

    sig { params(record_to_map: Stripe::APIResource, source_record: T.nilable(Restforce::SObject)).void }
    def map(record_to_map, source_record=nil)
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

    protected def build_dynamic_mapping_values(source_record, mapping_for_record)
      # build annotations using stripe resource metadata and users table
      # of accepted stripe metadata values and their corresponding NS keys
      mapping_for_record.each_with_object({}) do |(key_path, destination_field_name), assignments|
        next if destination_field_name.to_s.empty?

        extracted_value = extract_key_path_for_record(source_record, key_path)

        next unless extracted_value

        # TODO should we only do this if we run into a specific field suffix?
        #      I feel like this is good default behavior. When would someone NOT want this enabled?

        # convert true/false strings into boolean so the correct
        if extracted_value.to_s.downcase.strip == 'true'
          extracted_value = true
        elsif extracted_value.to_s.downcase.strip == 'false'
          extracted_value = false
        end

        assignments[destination_field_name] = extracted_value
      end
    end

    protected def assign_values_from_hash(record, field_assignments)
      field_assignments.each do |k, v|
        # TODO need to handle nil values
        # TODO think on custom fields with NetSuite suffix `_id`
        # TODO nofify when an existing field is being overwritten

        method = "#{k}=".to_sym

        if is_record_reference_field?(k)
          ref_method = "#{k[0..-4]}=".to_sym
        elsif is_date_field?(k)
          ref_method = "#{k[0..-6]}=".to_sym
        end

        if is_record_reference_field?(k) && record.respond_to?(ref_method)
          raise 'determine what salesforce record ref transformation requirements are'
          # record.send(ref_method, ref_value)
        elsif is_date_field?(k) && record.respond_to?(ref_method)
          raise 'determine what salesforce date transformation requirements are'
          # record.send(ref_method, ref_value)
        elsif record.respond_to?(method)
          record.send(method, v)
        else
          log.error 'invalid mapping field specified',
            mapping_key: k,
            mapping_value: v,
            record: record
        end
      end
    end

  end
end
