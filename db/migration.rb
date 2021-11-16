# frozen_string_literal: true
# typed: false

# heroku run bundle exec ruby migration.rb
# CREATE DATABASE stripeforce; CREATE DATABASE test_stripeforce

require 'sequel'

is_test_environment = ENV['RAILS_ENV'] && ENV['RAILS_ENV'] == 'test'

url = if is_test_environment
  ENV.fetch('TEST_DATABASE_URL')
else
  ENV.fetch('DATABASE_URL')
end

DB = Sequel.connect(url)

begin
  DB.drop_table :users
  DB.drop_table :poll_timestamps
rescue => e
  puts "error dropping tables"
end

DB.create_table(:users) do
  primary_key :id

  column :field_defaults, :jsonb, default: '{}', null: false
  column :field_mappings, :jsonb, default: '{}', null: false
  column :feature_flags, :jsonb, default: '[]', null: false

  DateTime :created_at, null: false
  DateTime :updated_at, null: false

  String :salesforce_account_id, unique: true, null: false
  String :salesforce_refresh_token
  String :salesforce_instance_url
  String :salesforce_token
  String :salesforce_organization_key, unique: true

  String :stripe_account_id
  String :stripe_refresh_token
  String :stripe_public_token

  TrueClass :enabled, default: true, null: false
  TrueClass :livemode, default: false, null: false

  String :name
  String :email

  unique [:stripe_account_id, :salesforce_account_id, :livemode]
end

DB.create_table(:poll_timestamps) do
  primary_key :id

  String :salesforce_account_id, null: false
  TrueClass :livemode, null: false

  # TODO this should be an enum
  # column :integration_record_type, :netsuite_record_types, null: false
  String :integration_record_type, null: false
  DateTime :last_polled_at, null: false

  # timestamps
  DateTime :created_at, null: false
  DateTime :updated_at, null: false

  unique [:salesforce_account_id, :livemode, :integration_record_type]
end
