# frozen_string_literal: true
# typed: false

# defines these POST routes for us:
#   - /auth/salesforcesandbox
#   - /auth/salesforce
#   - /auth/stripelivemode
#   - /auth/stripetestmode
#
# callbacks are explicitly defined in routes.rb. The GET versions of these
# routes are defined in routes.rb. This is to work out around a oauth security issue.


# We need to fork the Stripe strategy into two identical ones so we can use different keys
class OmniAuth::Strategies::StripeTestMode < OmniAuth::Strategies::Stripe
  extend OmniAuth::Strategy::ClassMethods
end
class OmniAuth::Strategies::StripeLiveMode < OmniAuth::Strategies::Stripe
  extend OmniAuth::Strategy::ClassMethods
end

class OmniAuth::Strategies::StripeTestModeV2 < OmniAuth::Strategies::Stripe
  extend OmniAuth::Strategy::ClassMethods

  def request_path
    '/auth/v2/stripetestmode'
  end

  def callback_path
    '/auth/v2/stripetestmode/callback'
  end

end

class OmniAuth::Strategies::StripeLiveModeV2 < OmniAuth::Strategies::Stripe
  extend OmniAuth::Strategy::ClassMethods

  def request_path
    '/auth/v2/stripelivemode'
  end

  def callback_path
    '/auth/v2/stripelivemode/callback'
  end

end

class OmniAuth::Strategies::SalesforceSandboxV2 < OmniAuth::Strategies::SalesforceSandbox
  extend OmniAuth::Strategy::ClassMethods

  def request_path
    '/auth/v2/salesforcesandbox'
  end

  def callback_path
    '/auth/v2/salesforcesandbox/callback'
  end

end

class OmniAuth::Strategies::SalesforceV2 < OmniAuth::Strategies::Salesforce
  extend OmniAuth::Strategy::ClassMethods

  def request_path
    '/auth/v2/salesforce'
  end

  def callback_path
    '/auth/v2/salesforce/callback'
  end

end

Rails.application.config.middleware.use OmniAuth::Builder do
  # the key type (livemode, testmode) must match the `STRIPE_CLIENT_ID` environment in Stripe
  provider OmniAuth::Strategies::StripeLiveMode,
    ENV.fetch("STRIPE_CLIENT_ID"),
    ENV.fetch("STRIPE_API_KEY"),
    scope: 'read_write'

  provider OmniAuth::Strategies::StripeTestMode,
    ENV.fetch("STRIPE_TEST_CLIENT_ID"),
    ENV.fetch("STRIPE_TEST_API_KEY"),
    scope: 'read_write'

  provider OmniAuth::Strategies::StripeLiveModeV2,
    ENV.fetch("STRIPE_CLIENT_ID"),
    ENV.fetch("STRIPE_API_KEY"),
    scope: 'read_write',
    path_prefix: '/auth/v2/stripelivemode'

  provider OmniAuth::Strategies::StripeTestModeV2,
    ENV.fetch("STRIPE_TEST_CLIENT_ID"),
    ENV.fetch("STRIPE_TEST_API_KEY"),
    scope: 'read_write',
    path_prefix: '/auth/v2/testmode'

  provider :salesforce, ENV.fetch('SF_CONSUMER_KEY'), ENV.fetch('SF_CONSUMER_SECRET')
  provider OmniAuth::Strategies::SalesforceV2,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']

  # same oauth consumer keys are used for sandbox & prod, but different destination URLs
  provider OmniAuth::Strategies::SalesforceSandbox,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
  provider OmniAuth::Strategies::SalesforceSandboxV2,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']

  provider OmniAuth::Strategies::SalesforcePreRelease,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
end
