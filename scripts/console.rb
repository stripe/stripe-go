#!/usr/bin/env bundle exec ruby
# Description: setup a REPL with a SF user and example records ready to go
# frozen_string_literal: true
# typed: true

require File.expand_path('../config/environment', __dir__)

include StripeForce::Constants

if ARGV[0]
  prefix = "SF_#{ARGV[0].upcase}_"
else
  prefix = "SF_"
end

user = StripeForce::User.find(salesforce_account_id: ENV.fetch(prefix + 'INSTANCE_ID'))

if user
  puts "Using local user reference"
end

# most likely, the local creds will be expired
user ||= StripeForce::User.new(
  salesforce_account_id: ENV.fetch(prefix + 'INSTANCE_ID'),
  salesforce_token: ENV.fetch(prefix + 'ACCESS_TOKEN'),
  salesforce_refresh_token: ENV[prefix + 'REFRESH_TOKEN'],
  salesforce_instance_url: "https://#{ENV.fetch(prefix + 'INSTANCE_DOMAIN')}.my.salesforce.com",

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

def wipe_record_tree(order_id)
  order = sf.find(SF_ORDER, order_id)

  account = sf.find(SF_ACCOUNT, order.AccountId)
end

def delete_sync_records
  result = @sf.query("SELECT Id FROM Sync_Record__c")
  result.each do |sync_record|
    puts "destroying\t#{sync_record.Id}"
    sync_record.destroy
  end
end

# TODO this doesn't seem like it works on some accounts
# https://salesforce.stackexchange.com/questions/186025/how-to-we-get-list-of-installed-packages-and-it-version-number
# sf.query("SELECT Id FROM InstalledSubscriberPackage")

# TODO determine what users have the permission set assigned
# u.sf_client.query("SELECT Id, AssigneeId, Assignee.Name FROM PermissionSetAssignment WHERE PermissionSet.Name = 'Stripe_Connector_Integration_User'")

# dig into field level permissions "Field Permissions"
# u.sf_client.api_get 'sobjects/'
def get_fields_for_object(object_name)
  description = @sf.describe(object_name)
  description['fields'].map(&:name)
end

def get_all(object_name)
  all_fields = get_fields_for_object(object_name).join(',')
  sf.query("SELECT #{all_fields} FROM #{object_name}")
end

require_relative './test/support/salesforce_debugging'
include SalesforceDebugging

user_info = sf.user_info

puts "Salesforce account information:"
puts user_info['username']
puts user_info['urls']['custom_domain']

Pry.start
