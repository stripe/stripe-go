# frozen_string_literal: true
# typed: true

module StripeForce
  class User < Sequel::Model
    extend T::Sig

    include StripeForce::Constants
    include Integrations::ErrorContext

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

    def after_initialize
      if self.new?
        self.enable_feature(:loud_sandbox_logging)
        self.connector_settings = {
          api_percentage_limit: 95,
          sync_start_date: Time.now.to_i,
          sync_record_retention: 10_000,
          default_currency: 'USD',
          CONNECTOR_SETTING_CPQ_TERM_UNIT => 'month',
        }
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
        oauth_token: salesforce_token,
        refresh_token: salesforce_refresh_token,
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

      # TODO should we conditionally do this?
      # @client.authenticate!

      @client
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

    # although you can write the logic to extract this in REST, we need this value to properly
    # establish a connection.
    def salesforce_instance_type
      if connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE].nil?
        report_edge_case("instance type should not be empty")
      end

      connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE]
    end

    # for our purposes a sandbox is anything that isn't a production account
    def sandbox?
      salesforce_instance_type != SFInstanceTypes::PRODUCTION.serialize
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
        "price" => {
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
          # TODO setting custom Ids may not be the best idea here
          # "id" => "Id",
          "name" => 'Name',
          "description" => 'Description',
        },
        "price" => {
          # default monthly fallback is used if this value is empty
          "recurring.interval_count" => CPQ_QUOTE_BILLING_FREQUENCY,
          "recurring.usage_type" => CPQ_PRODUCT_BILLING_TYPE,
        },
        "invoice" => {},
      }
    end

    sig { params(feature: Symbol, update: T::Boolean).void }
    def enable_feature(feature, update: false)
      if !feature_flags.include?(feature)
        feature_flags << feature

        if update
          self.save(columns: [:feature_flags])
        end
      end
    end

    sig { params(feature: Symbol, update: T::Boolean).void }
    def disable_feature(feature, update: false)
      if !feature_flags.delete(feature).nil? && update
        save(columns: [:feature_flags])
      end
    end

    sig { params(feature: Symbol).returns(T::Boolean) }
    def feature_enabled?(feature)
      self.feature_flags.include?(feature)
    end

    def sf_subdomain
      URI.parse(T.must(salesforce_instance_url)).host&.split('.')&.first
    end

    def sf_endpoint
      salesforce_instance_url
    end

    def stripe_credentials
      {
        api_key: ENV.fetch('STRIPE_CLIENT_SECRET'),
        stripe_account: stripe_account_id,
        stripe_version: '2020-08-27',
      }
    end
  end
end
