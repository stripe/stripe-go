# frozen_string_literal: true
# typed: false

# defines these POST routes for us:
#   - /auth/salesforcesandbox
#   - /auth/salesforce
#   - /auth/stripe
#
# callbacks are explicitly defined in routes.rb. The GET versions of these
# routes are defined in routes.rb. This is to work out around a oauth security issue.

Rails.application.config.middleware.use OmniAuth::Builder do
  provider :stripe, ENV["STRIPE_CLIENT_ID"], ENV["STRIPE_CLIENT_SECRET"], scope: 'read_write'

  provider :salesforce, ENV.fetch('SF_CONSUMER_KEY'), ENV.fetch('SF_CONSUMER_SECRET')

  # same oauth consumer keys are used for sandbox & prod, but different destination URLs
  provider OmniAuth::Strategies::SalesforceSandbox,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
  provider OmniAuth::Strategies::SalesforcePreRelease,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
end
