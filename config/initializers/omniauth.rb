# frozen_string_literal: true
# typed: false
Rails.application.config.middleware.use OmniAuth::Builder do
  provider :salesforce, ENV['SF_CONSUMER_KEY'], ENV['SF_CONSUMER_SECRET']
  provider :stripe, ENV["STRIPE_CLIENT_ID"], ENV["STRIPE_CLIENT_SECRET"], scope: 'read_write'
end

# from sinatra
# OmniAuth.config.full_host = ENV.fetch("DOMAIN")
