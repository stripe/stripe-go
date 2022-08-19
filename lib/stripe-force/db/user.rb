# frozen_string_literal: true
# typed: true
require_relative './user/kms_encryption'

module StripeForce
  class User < Sequel::Model
    extend T::Sig

    include StripeForce::Constants
    include Integrations::ErrorContext
    include KMSEncryption

    plugin :timestamps, update_on_create: true
    plugin :after_initialize
    plugin :defaults_setter
    plugin :update_or_create

    plugin :serialization, :json, :feature_flags
    plugin :serialization, :json, :field_defaults
    plugin :serialization, :json, :field_mappings
    plugin :serialization, :json, :connector_settings

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
        api_version: '52.0',

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

    # TODO should move to a status service
    def has_cpq_installed?
      sf_client.query("SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE 'SBQQ%'").count >= 1
    end

    def salesforce_namespace
      self.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE]
    end

    # although you can write the logic to extract this in REST, we need this value to properly
    # establish a connection.
    def salesforce_instance_type
      # if a stripe ID is not set, it's possible that they are early on in the setup process
      if connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE].nil? && self.stripe_account_id
        report_edge_case("instance type should not be empty")
      end

      connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE]
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
          "recurring.interval_count" => "Product2.#{CPQ_QUOTE_BILLING_FREQUENCY}",
          "recurring.usage_type" => "Product2.#{CPQ_PRODUCT_BILLING_TYPE}",
        },

        "price_order_item" => {
          # these fields have identical names on the product level, which is why we can use the same constants here
          "recurring.interval_count" => CPQ_QUOTE_BILLING_FREQUENCY,
          "recurring.usage_type" => CPQ_PRODUCT_BILLING_TYPE,
        },

        "invoice" => {},
      }
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
      self.feature_flags.include?(feature.serialize.to_sym)
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

    protected def kms_encryption_context(field=nil)
      {
        'salesforce_account_id' => self.salesforce_account_id,
      }
    end

    protected def kms_encryption_key(field=nil)
      AWS_KMS_SALESFORCE_CREDENTIALS
    end

  end
end
