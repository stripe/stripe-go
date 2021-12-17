#!/usr/bin/env bundle exec ruby
# Description: setup a REPL with a SF user and example records ready to go
# frozen_string_literal: true
# typed: true

require File.expand_path('../config/environment', __dir__)

include StripeForce::Constants

user = StripeForce::User.find(salesforce_account_id: ENV.fetch('SF_INSTANCE_ID'))

# most likely, the local creds will be expired
user ||= StripeForce::User.new(
  salesforce_account_id: ENV.fetch('SF_INSTANCE_ID'),
  salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
  salesforce_refresh_token: ENV['SF_REFRESH_TOKEN'],
  salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE_DOMAIN')}.my.salesforce.com",

  stripe_account_id: ENV.fetch('STRIPE_ACCOUNT_ID')
)

if user.salesforce_instance_url.blank?
  puts "ERROR: invalid instance URL, local user may be corrupted"
end

@user = user
@sf = sf = user.sf_client

def translate_order(order_id)
  locker = Integrations::Locker.new(@user)
  sf_order = @sf.find(SF_ORDER, order_id)
  StripeForce::Translate.perform(user: @user, sf_object: sf_order, locker: locker)
end

def loud_sf_logging
  ENV['SALESFORCE_LOG'] = 'true'
  @user.refresh
  @sf = sf = @user.sf_client
end

def example_sf_order
  # sf.find('Order', '8015e000000IJ1rAAG')

  # order with recurring item
  # sf.find('Order', '8015e000000IJDxAAO')

  @sf.find('Order', '8015e000000IJF5AAO')
  # @sf.find('Order', '8015e000000IIpgAAG')
end

Pry.start
