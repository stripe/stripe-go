# frozen_string_literal: true
# typed: false

Rails.application.config.middleware.use OmniAuth::Builder do
  provider :stripe, ENV["STRIPE_CLIENT_ID"], ENV["STRIPE_CLIENT_SECRET"], scope: 'read_write'

  provider :salesforce, ENV['SF_CONSUMER_KEY'], ENV['SF_CONSUMER_SECRET']

  provider OmniAuth::Strategies::SalesforceSandbox,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
  provider OmniAuth::Strategies::SalesforcePreRelease,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
end

# from sinatra
# OmniAuth.config.full_host = ENV.fetch("DOMAIN")
