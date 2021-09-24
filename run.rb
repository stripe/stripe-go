require_relative 'config'

Stripe.api_key = ENV['STRIPE_KEY']
Stripe.api_version = '2020-08-27'

# credentials from the CLI tool is available here: /Users/mbianco/.sfdx/*.json
# puts "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/authorize?response_type=code&client_id=#{ENV['SF_CONSUMER_KEY']}&redirect_uri=https://login.salesforce.com"
# oauth_code = "aPrxD8aRZGb_abrLry4Dl8fC.tyhGkTBCIWTjVKZ59yRhJssNw2ToltBW5zQYzM4JAE3ZeHDdA%3D%3D"
# "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/token?code=#{oauth_code}&grant_type=authorization_code&client_id=#{ENV['SF_CONSUMER_KEY']}&client_secret=#{ENV.fetch('SF_CONSUMER_SECRET')}&redirect_uri=https://login.salesforce.com"
# curl -X POST

Restforce.configure do |config|
  config.log_level = :debug
  # config.log = true
end

# really? Can't set this on an instance or `configure` level?
Restforce.log = ENV.fetch('SALESFORCE_LOG', 'false') == 'true'

user = StripeForce::User.new
StripeForce::OrderPoller.perform(user: user)

# sf.authenticate! will refresh oauth tokens

# oauth_token: ENV.fetch('SF_ACCESS_TOKEN'),
# refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
# instance_url: sf_endpoint,
# "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com"