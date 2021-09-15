require 'rubygems'
require 'bundler'

Bundler.require(:default, :development)
require_relative 'helpers'

Stripe.api_key = ENV['STRIPE_KEY']
Stripe.api_version = '2020-08-27'

"https://#{ENV['SF_INSTANCE']}.salesforce.com/services/oauth2/authorize?response_type=code&client_id=#{ENV['SF_CONSUMER_KEY']}>&redirect_uri=https://login.salesforce.com/"

credential_options = Dir["/Users/mbianco/.sfdx/*.json"]

if credential_options.count > 1
  raise "more than one SF account authenticated"
end

credentials = JSON.parse(File.read(credential_options.first))

client = Restforce.new(
  oauth_token: credentials['accessToken'],
  refresh_token: credentials['refreshToken'],
  instance_url: credentials['instanceUrl'],
  client_id: credentials['clientId'],
  client_secret: '1384510088588713504',
  authentication_callback: Proc.new { |x| Rails.logger.debug x.to_s },

  # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_versions.htm
  # `http https://biancodevorg-dev-ed.my.salesforce.com/services/data`
  api_version: '52.0',

  log_level: :debug
)

Restforce.configure do |config|
  config.log_level = :debug
end

client.create!('Account', Name: 'REST Customer')

binding.pry