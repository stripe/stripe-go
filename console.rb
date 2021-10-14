# frozen_string_literal: true
# typed: strict
require_relative 'config'

include StripeForce::Constants

user = StripeForce::User.new(
  salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
  salesforce_refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
  salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com",

  stripe_account_id: ENV.fetch('STRIPE_ACCOUNT_ID')
)

sf = user.sf_client

Pry.start
