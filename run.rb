require 'dotenv/load'

require 'rubygems'
require 'bundler'

module StripeForce

end


Bundler.require(:default, :development)

require_relative 'user'
require_relative 'translate'

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

def poll_orders
  user = StripeForce::User.new
  sf = user.sf_client

  updated_orders = sf.get_updated('Order', DateTime.now - 1, DateTime.now)
  # anything else but "ids" in the hash?
  updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

  updated_orders.each do |sf_order_id|
    sf_order = sf.find('Order', sf_order_id)
    StripeForce::Translate.perform(user: user, sf_object: sf_order)
  end
end

poll_orders

# sf.authenticate! will refresh oauth tokens
