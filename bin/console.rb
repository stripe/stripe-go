#!/usr/bin/env ruby

# frozen_string_literal: true
# typed: true

require File.expand_path('../config/environment', __dir__)

include StripeForce::Constants

user = StripeForce::User.find(salesforce_account_id: ENV.fetch('SF_INSTANCE_ID'))

# most likely, the local creds will be expired
user ||= StripeForce::User.new(
  salesforce_account_id: ENV.fetch('SF_INSTANCE_UID'),
  salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
  salesforce_refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
  salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE_DOMAIN')}.my.salesforce.com",

  stripe_account_id: ENV.fetch('STRIPE_ACCOUNT_ID')
)

@user = user
@sf = sf = user.sf_client

def example_sf_order
  # sf.find('Order', '8015e000000IJ1rAAG')

  # order with recurring item
  # sf.find('Order', '8015e000000IJDxAAO')

  @sf.find('Order', '8015e000000IJF5AAO')
end

Pry.start
