require 'dotenv/load'

require 'rubygems'
require 'bundler'

$stdout.sync = true

module StripeForce

end

Bundler.require(:default, :development)

Sentry.init do |config|
  config.dsn = ENV.fetch('SENTRY_DSN')
  config.traces_sample_rate = 0.5
end

# CREATE DATABASE stripeforce
DB = Sequel.connect(ENV.fetch('DATABASE_URL'))

Restforce.configure do |config|
  config.log_level = :debug
  # config.log = true
end

# really? Can't set this on an instance or `configure` level?
Restforce.log = ENV.fetch('SALESFORCE_LOG', 'false') == 'true'

require_relative 'constants'
require_relative 'user'
require_relative 'translate'
require_relative 'polling'
