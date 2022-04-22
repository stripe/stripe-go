# frozen_string_literal: true
# typed: true
require_relative "boot"

require "action_controller/railtie"
require "action_view/railtie"
require "rails/test_unit/railtie"

# Require the gems listed in Gemfile, including any gems
# you've limited to :test, :development, or :production.
Bundler.require(*Rails.groups)

require_relative '../lib/stripe-force/constants'

module StripeForce
  class Application < Rails::Application
    # Initialize configuration defaults for originally generated Rails version.
    config.load_defaults 6.1

    # Configuration for the application, engines, and railties goes here.
    #
    # These settings can be overridden in specific environments using the files
    # in config/environments, which are processed later.
    #
    # config.time_zone = "Central Time (US & Canada)"
    # config.eager_load_paths << Rails.root.join("extras")

    # Only loads a smaller set of middleware suitable for API only apps.
    # Middleware like session, flash, cookies can be added back manually.
    # Skip views, helpers and assets when generating a new resource.
    config.api_only = true

    # workaround no sessions on API-only applications; we need sessions for oauth
    # https://github.com/omniauth/omniauth#integrating-omniauth-into-your-rails-api
    Rails.application.config.session_store :cookie_store, key: '_stripe-salesforce_session'
    config.middleware.use ActionDispatch::Cookies
    config.middleware.use ActionDispatch::Session::CookieStore, config.session_options

    # https://github.com/stripe/stripe-netsuite/issues/712
    config.force_ssl = true

    config.hosts += ENV.fetch('STRIPEFORCE_HOSTS', '').split(",")

    # TODO I don't believe any emails we send will contain domains
    # config.action_controller.default_url_options = {
    #   host: ENV.fetch('SALESFORCE_APP_DOMAIN'),
    # }

    include StripeForce::Constants

    SALESFORCE_HEADERS = [
      SALESFORCE_ACCOUNT_ID_HEADER,
      SALESFORCE_PACKAGE_NAMESPACE_HEADER,
      SALESFORCE_INSTANCE_TYPE_HEADER,
    ]

    config.lograge.enabled = true
    config.lograge.custom_options = lambda do |event|
      # NOTE lograge doesn't include params by default, we need to manually include them here
      custom_params = {"params" => event.payload[:params].except('controller', 'action')}

      # include headers which identify which account and other details the request is coming from=
      SALESFORCE_HEADERS.each {|h| custom_params[h] = event.payload[:request].headers[h] }

      custom_params.reject {|_, v| v.blank? }
    end
  end
end

ENV['RAILS_CACHE_ID'] = ENV['HEROKU_RELEASE_VERSION']
