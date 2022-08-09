# Usage: bundle exec sequel -E -m db/ $DATABASE_URL

# frozen_string_literal: true
# typed: false

require 'sequel'

Sequel.migration do
  up do
    create_table(:users) do
      primary_key :id

      column :field_defaults, :jsonb, default: '{}', null: false
      column :field_mappings, :jsonb, default: '{}', null: false
      column :feature_flags, :jsonb, default: '[]', null: false
      column :connector_settings, :jsonb, default: '{}', null: false

      DateTime :created_at, null: false
      DateTime :updated_at, null: false

      String :salesforce_account_id, unique: true, null: false
      String :salesforce_instance_url
      String :salesforce_organization_key, unique: true

      column :encrypted_salesforce_token, :text
      column :encrypted_salesforce_refresh_token, :text

      String :stripe_account_id
      String :stripe_public_token

      TrueClass :enabled, default: true, null: false
      TrueClass :livemode, default: false, null: false

      String :name
      String :email

      unique [:stripe_account_id, :salesforce_account_id, :livemode]
    end

    create_table(:poll_timestamps) do
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
  end

  down do
    drop_table :users
    drop_table :poll_timestamps
  end
end
