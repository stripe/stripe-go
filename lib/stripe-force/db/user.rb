# frozen_string_literal: true
# typed: true
require_relative './user/kms_encryption'

module StripeForce
  class User < Sequel::Model
    extend T::Sig

    include Integrations::Log
    include StripeForce::Constants
    include KMSEncryption

    plugin :timestamps, update_on_create: true
    plugin :after_initialize
    plugin :defaults_setter
    plugin :update_or_create

    plugin :serialization, :json, :feature_flags
    plugin :serialization, :json, :field_defaults
    plugin :serialization, :json, :field_mappings
    plugin :serialization, :json, :connector_settings
    plugin :serialization, :json, :salesforce_object_prefix_mappings

    SF_CONSUMER_KEY = ENV.fetch('SF_CONSUMER_KEY')
    SF_CONSUMER_SECRET = ENV.fetch('SF_CONSUMER_SECRET')

    # platform stripe API tokens
    SF_STRIPE_LIVEMODE_API_KEY = ENV.fetch('STRIPE_API_KEY')
    SF_STRIPE_TESTMODE_API_KEY = ENV.fetch('STRIPE_TEST_API_KEY')

    AWS_KMS_SALESFORCE_CREDENTIALS = ENV.fetch('AWS_KMS_SALESFORCE_CREDENTIALS')

    DEFAULT_CONNECTOR_SETTINGS = {
      api_percentage_limit: 95,
      sync_start_date: nil,
      sync_record_retention: 10_000,
      default_currency: 'USD',
      CONNECTOR_SETTING_POLLING_ENABLED => false,
      CONNECTOR_SETTING_SYNC_START_DATE => nil,
      CONNECTOR_SETTING_CPQ_TERM_UNIT => 'month',
      "filters": {
        SF_ACCOUNT => nil,
        SF_ORDER => "Status = 'Activated'",
        SF_PRODUCT => nil,
      }.stringify_keys.freeze,
    }.stringify_keys.freeze

    kms_encrypted_field :salesforce_token
    kms_encrypted_field :salesforce_refresh_token

    def after_initialize

      if self.new?
        self.enable_feature(FeatureFlags::LOUD_SANDBOX_LOGGING)
        self.enable_feature(FeatureFlags::TEST_CLOCKS)
        self.enable_feature(FeatureFlags::CATCH_ALL_ERRORS)
        self.enable_feature(FeatureFlags::SF_CACHING)
        self.connector_settings = DEFAULT_CONNECTOR_SETTINGS.deep_dup
      end
      self.feature_flags.map!(&:to_sym)
    end

    sig { returns(String) }
    def metadata_prefix
      # only use sandbox prefix when translating to sandbox in livemode
      # https://stripe.slack.com/archives/netsuite/p1446678774000004
      if self.sandbox? && self.livemode
        "salesforce_sandbox_"
      else
        "salesforce_"
      end
    end

    def sf_client
      optional_client_params = {}

      if sandbox?
        optional_client_params[:host] = 'test.salesforce.com'
      end

      @client ||= Restforce.new({
        # this could be expired, if it is the client will automatically refresh it
        oauth_token: self.salesforce_token,
        refresh_token: self.salesforce_refresh_token,
        instance_url: sf_endpoint,

        client_id: SF_CONSUMER_KEY,
        client_secret: SF_CONSUMER_SECRET,

        # authentication_callback: proc {|x| log.info x.to_s },

        # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_versions.htm
        # `http https://biancodevorg-dev-ed.my.salesforce.com/services/data`
        api_version: '58.0',
        log: true,

        log_level: :debug,
      }.merge(optional_client_params))

      # if access token has expired, this will be automatically refreshed by restforce middleware
      # if the token is refreshed, we check for this and persist it at the end of a translation job

      @client
    end

    def persist_refreshed_credentials
      if self.sf_client.options[:oauth_token] != self.salesforce_token
        log.info 'credentials refreshed, updating'
        self.update(salesforce_token: self.sf_client.options[:oauth_token])
      end
    end

    # TODO there's not a practical limit here
    # https://www.infallibletechie.com/2018/12/what-is-concurrent-api-request-limit-in.html
    def concurrent_connections
      10
    end

    sig { params(platform: StripeForce::Constants::Platforms).returns(String) }
    def cached_valid_credential_key(platform)
      "cached-#{platform}-connection-status-#{self.stripe_account_id}-#{self.salesforce_account_id}-#{self.livemode}"
    end

    sig { params(platform: StripeForce::Constants::Platforms, connected: T::Boolean).returns(T::Boolean) }
    def cache_connection_status(platform, connected)
      redis.set(self.cached_valid_credential_key(platform), connected, ex: StripeForce::Constants::CACHED_CREDENTIAL_STATUS_TTL)
      connected
    end

    def get_cached_connection_status(platform)
      cached_value = redis.get(self.cached_valid_credential_key(platform))
      # Redis only stores strings, so we have to do this awkward workaround
      unless cached_value.nil?
        return cached_value == 'true'
      end

      nil
    end

    # TODO: Make this more robust via a cron (caching status, sending emails etc)
    # https://jira.corp.stripe.com/browse/PLATINT-1450
    def valid_credentials?
      valid_credentials_stripe? && valid_credentials_salesforce?
    end

    # Force a check w/o using cached status -- for the connection_statuses endpoint
    def valid_credentials_stripe!
      check_credentials_stripe
    end

    # Check connection status against the cache
    def valid_credentials_stripe?
      # If our cached value has expired, these will return nil and they will be fetched again.
      stripe_credentials_valid = get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE)

      if stripe_credentials_valid.nil?
        check_credentials_stripe
      else
        stripe_credentials_valid
      end
    end

    def check_credentials_stripe
      Stripe::Account.retrieve(
        self.stripe_account_id,
        self.stripe_credentials
      )
      self.cache_connection_status(StripeForce::Constants::Platforms::STRIPE, true)
    rescue Stripe::AuthenticationError, Stripe::PermissionError
      log.info "invalid Stripe credentials"
      self.cache_connection_status(StripeForce::Constants::Platforms::STRIPE, false)
    end

    # Force a check w/o using cached status -- for the connection_statuses endpoint
    def valid_credentials_salesforce!
      check_credentials_salesforce
    end

    def valid_credentials_salesforce?
      # If our cached value has expired, these will return nil and they will be fetched again.
      salesforce_credentials_valid = get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE)

      # Salesforce
      if salesforce_credentials_valid.nil?
        check_credentials_salesforce
      else
        salesforce_credentials_valid
      end
    end

    def check_credentials_salesforce
      # just initializing the client is not enough, a call must be attempted
      self.sf_client.user_info
      self.cache_connection_status(StripeForce::Constants::Platforms::SALESFORCE, true)
    rescue Restforce::UnauthorizedError, Restforce::AuthenticationError
      log.info "invalid Salesforce credentials"
      self.cache_connection_status(StripeForce::Constants::Platforms::SALESFORCE, false)
    end

    def salesforce_namespace
      self.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE]
    end

    def is_multicurrency_org?
      self.connector_settings[CONNECTOR_SETTING_MULTICURRENCY_ENABLED]
    end

    # although you can write the logic to extract this in REST, we need this value to properly
    # establish a connection.
    def salesforce_instance_type
      # if a stripe ID is not set, it's possible that they are early on in the setup process
      if connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE].nil? && self.stripe_account_id
        Integrations::ErrorContext.report_edge_case("instance type should not be empty")
      end

      connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE]
    end

    def last_polled_timestamp_record(sf_record_class: SF_ORDER)
      StripeForce::PollTimestamp.by_user_and_record(self, sf_record_class)
    end

    def last_synced
      poll_timestamp_record = last_polled_timestamp_record
      unless poll_timestamp_record.nil?
        return poll_timestamp_record.last_polled_at.to_i
      end

      nil
    end

    def polling_enabled?
      # For existing users, this setting might not be set.
      if connector_settings[CONNECTOR_SETTING_POLLING_ENABLED].nil?
        # If we have set up their Order poll timestamp, we can safely say this is true otherwise lets say false
        connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = self.last_polled_timestamp_record.nil? ? false : true
      end

      connector_settings[CONNECTOR_SETTING_POLLING_ENABLED]
    end

    # for our purposes a sandbox is anything that isn't a production account
    def sandbox?
      salesforce_instance_type != SFInstanceTypes::PRODUCTION.serialize
    end

    def scratch_org?
      salesforce_instance_type == SFInstanceTypes::SCRATCH_ORG.serialize
    end

    def in_production?
      livemode && !sandbox?
    end

    def required_mappings
      {
        "product" => {},
        "customer" => {},
        "invoice" => {},
        "subscription_schedule" => {
          "start_date" => "#{CPQ_QUOTE}.#{CPQ_QUOTE_SUBSCRIPTION_START_DATE}",
          "iterations" => "#{CPQ_QUOTE}.#{CPQ_QUOTE_SUBSCRIPTION_TERM}",
        },
        "subscription_item" => {
          "quantity" => "Quantity",
        },
        "price" => {
          "unit_amount_decimal" => 'UnitPrice',
        },
        "price_order_item" => {
          "unit_amount_decimal" => 'UnitPrice',
        },
        "coupon" => {},
      }
    end

    # TODO needs to be dynamic based off FF configuration, but for now we don't have any flags :)
    def default_mappings
      {
        "customer" => {
          "name" => "Name",
          "phone" => "Phone",
          "description" => "Description",
          # TODO may need to split BillingStreet into `line2` if multiple lines
          "address.line1" => "BillingStreet",
          "address.city" => "BillingCity",
          "address.state" => "BillingState",
          "address.postal_code" => "BillingPostalCode",
          "address.country" => "BillingCountry",
          "shipping.name" => "Name",
          "shipping.phone" => "Phone",
          "shipping.address.line1" => "ShippingStreet",
          "shipping.address.city" => "ShippingCity",
          "shipping.address.state" => "ShippingState",
          "shipping.address.postal_code" => "ShippingPostalCode",
          "shipping.address.country" => "ShippingCountry",
        },
        "product" => {
          "name" => 'Name',
          "description" => 'Description',
        },

        # price => pricebookentry mapping
        "price" => {
          # default monthly fallback is used if this value is empty
          # note that `Id` is not added to the suffix of the Product2 component, this is added by the mapper

          # We default the interval count to the billing frequency and map it to an integer in transform_salesforce_billing_frequency_to_recurring_interval
          #   this is so we can support quarterly, semiannually etc. billing frequencies. This can be improved:
          #   https://jira.corp.stripe.com/browse/PLATINT-2480
          "recurring.interval_count" => "Product2.#{CPQ_QUOTE_BILLING_FREQUENCY}",
          "recurring.usage_type" => "Product2.#{CPQ_PRODUCT_BILLING_TYPE}",
        },

        "price_order_item" => {
          # these fields have identical names on the product level, which is why we can use the same constants here

          # We default the interval count to the billing frequency and map it to an integer in transform_salesforce_billing_frequency_to_recurring_interval
          #   this is so we can support quarterly, semiannually etc. billing frequencies. This can be improved:
          #   https://jira.corp.stripe.com/browse/PLATINT-2480
          "recurring.interval_count" => CPQ_QUOTE_BILLING_FREQUENCY,
          "recurring.usage_type" => CPQ_PRODUCT_BILLING_TYPE,
        },

        "invoice" => {},

        "coupon" => {
          "name" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Name__c'),
          "amount_off" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Amount_Off__c'),
          "percent_off" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Percent_Off__c'),
          "duration" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Duration__c'),
          "duration_in_months" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Duration_In_Months__c'),
          "max_redemptions" => StripeForce::Utilities::SalesforceUtil.prefixed_stripe_field_(user: self, field_name: 'Max_Redemptions__c'),
        },
      }
    end

    # fields to hide from the Data Mapper UI depending
    # on the feature flags the user has enabled
    def hidden_mapper_fields
      hidden_mapper_fields = []

      if !feature_enabled?(FeatureFlags::INVOICE_RENDERING_TEMPLATE)
        hidden_mapper_fields << "subscription_schedule.default_settings.invoice_settings.rendering.template"
        hidden_mapper_fields << "subscription_schedule.default_settings.invoice_settings.rendering.template_version"
      end

      hidden_mapper_fields
    end

    def hidden_sync_pref_fields
      hidden_sync_pref_fields = []

      if !feature_enabled?(FeatureFlags::NON_ANNIVERSARY_AMENDMENTS)
        hidden_sync_pref_fields << "cpq_prorate_precision"
      end

      hidden_sync_pref_fields
    end

    sig { params(feature: FeatureFlags, update: T::Boolean).void }
    def enable_feature(feature, update: false)
      if !feature_flags.include?(feature.serialize.to_sym)
        feature_flags << feature.serialize.to_sym

        if update
          self.save(columns: [:feature_flags])
        end
      end
    end

    sig { params(feature: FeatureFlags, update: T::Boolean).void }
    def disable_feature(feature, update: false)
      if !feature_flags.delete(feature.serialize.to_sym).nil? && update
        save(columns: [:feature_flags])
      end
    end

    sig { params(feature: FeatureFlags).returns(T::Boolean) }
    def feature_enabled?(feature)
      self.feature_flags.include?(feature.serialize.to_sym) || self.feature_flags.include?(feature.serialize)
    end

    def sf_subdomain
      URI.parse(T.must(salesforce_instance_url)).host&.split('.')&.first
    end

    def sf_endpoint
      salesforce_instance_url
    end

    sig { params(forced_livemode: T.nilable(T::Boolean)).returns(T::Hash[Symbol, T.nilable(String)]) }
    def stripe_credentials(forced_livemode: nil)
      livemode_for_credentials = if forced_livemode.nil?
        self.livemode
      else
        forced_livemode
      end

      {
        api_key: livemode_for_credentials ? SF_STRIPE_LIVEMODE_API_KEY : SF_STRIPE_TESTMODE_API_KEY,
        stripe_account: stripe_account_id,
      }
    end

    # this hash is used to prevent same-time editing via the SF UI
    def configuration_hash
      mappings_as_string = self.field_defaults.to_json + self.field_mappings.to_json

      # take into account all configuration options which are set via the UI
      selective_configuration_json = self.connector_settings.slice(*DEFAULT_CONNECTOR_SETTINGS.keys).to_json

      Digest::SHA1.hexdigest(mappings_as_string + selective_configuration_json)
    end

    def user_specified_where_clause_for_object(poll_type)
      custom_query = self.connector_settings.dig("filters", poll_type)

      if custom_query && custom_query.strip.present?
        log.info 'found custom query for poll type', poll_type: poll_type
        " AND " + custom_query
      else
        ""
      end
    end

    protected def kms_encryption_context(field=nil)
      {
        # NOTE that the account is the SF organization instance ID
        'salesforce_account_id' => self.salesforce_account_id,
      }
    end

    protected def kms_encryption_key(field=nil)
      AWS_KMS_SALESFORCE_CREDENTIALS
    end

    private def redis
      Resque.redis
    end

  end
end
