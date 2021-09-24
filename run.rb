require_relative 'config'

# credentials from the CLI tool is available here: /Users/mbianco/.sfdx/*.json
# puts "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/authorize?response_type=code&client_id=#{ENV['SF_CONSUMER_KEY']}&redirect_uri=https://login.salesforce.com"
# oauth_code = "aPrxD8aRZGb_abrLry4Dl8fC.tyhGkTBCIWTjVKZ59yRhJssNw2ToltBW5zQYzM4JAE3ZeHDdA%3D%3D"
# "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/token?code=#{oauth_code}&grant_type=authorization_code&client_id=#{ENV['SF_CONSUMER_KEY']}&client_secret=#{ENV.fetch('SF_CONSUMER_SECRET')}&redirect_uri=https://login.salesforce.com"
# curl -X POST


# sf.authenticate! will refresh oauth tokens

user = StripeForce::User.new(
  salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
  salesforce_refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
  salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com",
)

sf = user.sf_client

# sf_object = user.sf_client.find("Order", "")
# StripeForce::Translate.perform(user: user, sf_object: sf_object)

# StripeForce::OrderPoller.perform(user: user)

translator = StripeForce::Translate.new(user)
translator.create_customer_from_sf_order(user.sf_client.find("Order", "8015e000000IJDxAAO"))
binding.pry