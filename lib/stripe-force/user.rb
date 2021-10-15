# frozen_string_literal: true
# typed: true
module StripeForce
  class User < Sequel::Model
    def sf_client
      client = Restforce.new(
        oauth_token: salesforce_token,
        refresh_token: salesforce_refresh_token,
        instance_url: sf_endpoint,

        client_id: ENV.fetch('SF_CONSUMER_KEY'),
        client_secret: ENV.fetch('SF_CONSUMER_SECRET'),

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
