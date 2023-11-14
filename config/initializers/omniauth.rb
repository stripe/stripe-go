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
class OmniAuth::Strategies::StripeAbstraction < OmniAuth::Strategies::Stripe
  # This is so we can feed our state value through the Stripe auth protocol.
  def authorize_params
    params = super

    params[:state] = session[:state]
    session["omniauth.state"] = session[:state]

    params
  end
end
class OmniAuth::Strategies::StripeTestMode < OmniAuth::Strategies::StripeAbstraction

end
class OmniAuth::Strategies::StripeLiveMode < OmniAuth::Strategies::StripeAbstraction
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

  # same oauth consumer keys are used for sandbox & prod, but different destination URLs
  provider :salesforce, ENV.fetch('SF_CONSUMER_KEY'), ENV.fetch('SF_CONSUMER_SECRET')

  provider OmniAuth::Strategies::SalesforceSandbox,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']

  provider OmniAuth::Strategies::SalesforcePreRelease,
    ENV['SF_CONSUMER_KEY'],
    ENV['SF_CONSUMER_SECRET']
end
