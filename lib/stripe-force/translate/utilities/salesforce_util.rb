# typed: true
# frozen_string_literal: true

# SalesforceUtil to avoid namespace madness with Stripe::* objects
module StripeForce::Utilities
  module SalesforceUtil
    extend T::Sig
    include Kernel

    include Integrations::Log
    include StripeForce::Constants

    # SF dates have no TZ data and come in as a simple 'YYYY-MM-DD'
    # Stripe APIs speak UTC, so we convert to UTC + unix timestamp
    sig { params(date_string: String).returns(Integer) }
    def self.salesforce_date_to_unix_timestamp(date_string)
      salesforce_date_to_beginning_of_day(date_string).to_i
    end

    sig { params(datetime: T.any(Time, DateTime)).returns(Integer) }
    def self.datetime_to_unix_timestamp(datetime)
      datetime.utc.beginning_of_day.to_i
    end

    sig { params(date_string: String).returns(Time) }
    def self.salesforce_date_to_beginning_of_day(date_string)
      DateTime.parse(date_string).utc.beginning_of_day
    end


    sig { params(user: StripeForce::User, sf_id: String).returns(String) }
    def self.salesforce_type_from_id(user, sf_id)
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
      when /^800/
        SF_CONTRACT
      when /^801/
        SF_ORDER
      when /^802/
        SF_ORDER_ITEM
      when /^0Mh/
        SF_CONSUMPTION_SCHEDULE
      when /^0Mo/
        SF_CONSUMPTION_RATE
      else
        object_prefix = T.must(sf_id[0..2])

        # check if we have cached this object's prefix mapping before
        user_prefix_mappings = user.salesforce_object_prefix_mappings

        if user_prefix_mappings[object_prefix].present?
          user_prefix_mappings[object_prefix]
        else
          # we have not seen this object prefix before, fetch the object info
          object_description = user
            .sf_client
            .api_get('sobjects')
            .body["sobjects"]
            .detect {|o| o["keyPrefix"] == object_prefix }

          # https://help.salesforce.com/s/articleView?id=000325244&type=1
          if !object_description
            raise ArgumentError.new("unknown object type #{sf_id}")
          end

          # cache the object's prefix to name value
          object_type = object_description["name"]
          user.salesforce_object_prefix_mappings[object_prefix] = object_type
          user.save

          object_type
        end
      end
    end

    sig { params(origin_salesforce_object: Restforce::SObject, salesforce_object: Restforce::SObject).returns(String) }
    def generate_compound_external_id(origin_salesforce_object, salesforce_object)
      "#{@origin_salesforce_object.Id}-#{salesforce_object.Id}"
    end

    sig { params(field_name: String).returns(String) }
    def prefixed_stripe_field(field_name)
      custom_field_prefix = case (salesforce_namespace = @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE])
      when nil
        Integrations::ErrorContext.report_edge_case("expected namespace to be defined, using fallback")
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

    sig { params(options: {}).returns(T.untyped) }
    def backoff(options={})
      count = 0

      options[:attempts] ||= if ENV['SALESFORCE_BACKOFF_ATTEMPTS'].nil?
        MAX_SF_RETRY_ATTEMPTS
      else
        ENV['SALESFORCE_BACKOFF_ATTEMPTS'].to_i
      end

      begin
        count += 1

        # runs the block (if given) and returns if no errors raised
        yield if block_given?
      rescue Restforce::ErrorCode::UnableToLockRow,
             Restforce::ServerError,
             Restforce::NotFoundError,
             Faraday::ConnectionFailed,
             Faraday::TimeoutError,
             Restforce::ResponseError,
             Restforce::UnauthorizedError => e

        # log & raise error if all retries fail
        if count >= options[:attempts]
          log.warn 'finished retrying SF operation, raising error',
            attempt: count,
            error_class: e.class.to_s,
            error_message: e.message
          raise e
        end

        sleep(count * count)
        retry
      end
    end

    # param_mapping: { stripe_key_name => salesforce_field_name }
    sig { params(mapper: StripeForce::Mapper, sf_record: Restforce::SObject, stripe_record_or_class: T.any(Class, Stripe::APIResource)).returns(Hash) }
    def self.extract_salesforce_params!(mapper, sf_record, stripe_record_or_class)
      stripe_mapping_key = StripeForce::Mapper.mapping_key_for_record(stripe_record_or_class, sf_record)

      user = mapper.user
      required_mappings = user.required_mappings[stripe_mapping_key]

      if required_mappings.nil?
        raise "expected mappings for #{stripe_mapping_key} but they were nil"
      end

      # first, let's pull required mappings and check if there's anything missing
      required_data = mapper.build_dynamic_mapping_values(sf_record, required_mappings)

      missing_stripe_fields = required_mappings.select {|k, _v| required_data[k].nil? }

      if missing_stripe_fields.present?
        missing_salesforce_fields = missing_stripe_fields.keys.map {|k| required_mappings[k] }

        raise Integrations::Errors::MissingRequiredFields.new(
          salesforce_object: sf_record,
          missing_salesforce_fields: missing_salesforce_fields
        )
      end

      # then, let's extract optional fields and then merge them in
      default_mappings = user.default_mappings[stripe_mapping_key]
      return required_data if default_mappings.blank?

      optional_data = mapper.build_dynamic_mapping_values(sf_record, default_mappings)

      required_data.merge(optional_data)
    end

    sig { params(mapper: StripeForce::Mapper, sf_order: Restforce::SObject).returns(Integer) }
    def self.extract_subscription_term_from_order!(mapper, sf_order)
      user = mapper.user

      subscription_term_stripe_path = ['subscription_schedule', 'iterations']
      subscription_term_order_path = user.field_mappings.dig(*subscription_term_stripe_path) ||
        user.required_mappings.dig(*subscription_term_stripe_path)

      quote_subscription_term = T.cast(mapper.extract_key_path_for_record(sf_order, subscription_term_order_path), T.nilable(T.any(String, Float)))

      if quote_subscription_term.nil?
        raise Integrations::Errors::MissingRequiredFields.new(
          salesforce_object: sf_order,
          missing_salesforce_fields: [subscription_term_order_path]
        )
      end

      # it's looking like these values are never really aligned and we should ignore the line item
      if sf_order[CPQ_QUOTE_SUBSCRIPTION_TERM] == quote_subscription_term
        Integrations::ErrorContext.report_edge_case("subscription term on quote matches line item")
      end

      if !Integrations::Utilities::StripeUtil.is_integer_value?(quote_subscription_term)
        raise StripeForce::Errors::RawUserError.new("Subscription term is specified as a decimal value")
      end

      quote_subscription_term.to_i
    end
  end
end
