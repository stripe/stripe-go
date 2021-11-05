# frozen_string_literal: true
# typed: false

# heroku run bundle exec ruby migration.rb
# CREATE DATABASE stripeforce

require 'sequel'

DB = Sequel.connect(ENV.fetch('DATABASE_URL'))

begin
  DB.drop_table :users
  DB.drop_table :poll_timestamps
rescue => e
end

DB.create_table(:users) do
  primary_key :id

  column :field_defaults, :jsonb, default: '{}', null: false
  column :field_mappings, :jsonb, default: '{}', null: false
  column :feature_flags, :jsonb, default: '[]', null: false

  DateTime :created_at, null: false
  DateTime :updated_at, null: false

  String :salesforce_account_id, unique: true, null: false
  String :salesforce_refresh_token, null: false
  String :salesforce_instance_url, null: false
  String :salesforce_token, null: false

  String :stripe_account_id
  String :stripe_refresh_token
  String :stripe_public_token

  TrueClass :enabled, default: true
  TrueClass :livemode, default: false

  String :name
  String :email

  unique [:stripe_account_id, :salesforce_account_id, :livemode]
end

DB.create_table(:poll_timestamps) do
  primary_key :id

  String :stripe_account_id, null: false

  TrueClass :livemode

  # TODO this should be an enum
  # column :integration_record_type, :netsuite_record_types, null: false
  String :integration_record_type
  DateTime :last_polled_at, null: false

  # timestamps
  DateTime :created_at, null: false
  DateTime :updated_at, null: false

  unique [:stripe_account_id, :livemode, :integration_record_type]
end
