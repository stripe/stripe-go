# frozen_string_literal: true
# typed: true
module StripeForce
  class User < Sequel::Model
    plugin :timestamps, update_on_create: true
    plugin :after_initialize
    plugin :defaults_setter

    plugin :serialization, :json, :feature_flags
    plugin :serialization, :json, :field_defaults
    plugin :serialization, :json, :field_mappings

    SF_CONSUMER_KEY = ENV.fetch('SF_CONSUMER_KEY')
    SF_CONSUMER_SECRET = ENV.fetch('SF_CONSUMER_SECRET')

    def after_initialize
      if self.new?
        self.enable_feature(:loud_sandbox_logging)
      end

      self.feature_flags.map!(&:to_sym)
    end

    def sf_client
      client = Restforce.new(
        oauth_token: salesforce_token,
        refresh_token: salesforce_refresh_token,
        instance_url: sf_endpoint,

        client_id: SF_CONSUMER_KEY,
        client_secret: SF_CONSUMER_SECRET,

        # authentication_callback: Proc.new { |x| Rails.logger.debug x.to_s },

        # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_versions.htm
        # `http https://biancodevorg-dev-ed.my.salesforce.com/services/data`
        api_version: '52.0',

        log: true,
        log_level: :debug
      )

      client.authenticate!

      client
    end

    # TODO there's not a practical limit here
    # https://www.infallibletechie.com/2018/12/what-is-concurrent-api-request-limit-in.html
    def concurrent_connections
      10
    end

    def sandbox?
      # [SELECT IsSandbox FROM Organization]
      true
    end

    def in_production?
      livemode && !sandbox?
    end

    def enable_feature(feature, update: false)
      feature = feature.to_sym if !feature.is_a?(Symbol)

      # TODO use Set to avoid duplicates instead?
      if !feature_flags.include?(feature)
        feature_flags << feature

        if update
          self.save(columns: [:feature_flags])
        end
      end
    end

    def disable_feature(feature, update: false)
      feature = feature.to_sym if !feature.is_a?(Symbol)

      if !feature_flags.delete(feature).nil? && update
        save(columns: [:feature_flags])
      end
    end

    def feature_enabled?(feature)
      feature = feature.to_sym if !feature.is_a?(Symbol)

      self.feature_flags.include?(feature)
    end

    def sf_endpoint
      salesforce_instance_url
    end

    def stripe_credentials
      {
        api_key: ENV['STRIPE_CLIENT_SECRET'],
        stripe_account: stripe_account_id,
        stripe_version: '2020-08-27',
      }
    end
  end
end
