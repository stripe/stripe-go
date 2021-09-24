module StripeForce
  class User < Sequel::Model
    def sf_client
      Restforce.new(
        oauth_token: ENV.fetch('SF_ACCESS_TOKEN'),
        refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
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
    end

    def sf_endpoint
      "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com"
    end
  end
end