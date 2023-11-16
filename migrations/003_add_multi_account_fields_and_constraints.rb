# Usage: bundle exec sequel -E -m migrations $DATABASE_URL

# frozen_string_literal: true
# typed: false

require 'sequel'

Sequel.migration do
  up do
    alter_table(:users) do
      drop_constraint('users_salesforce_account_id_key')
      add_column :is_default_account_config, :boolean, default: true, null: false
    end
  end

  down do
    alter_table(:users) do
      add_unique_constraint(:salesforce_account_id, name: 'users_salesforce_account_id_key')
      drop_column :is_default_account_config
    end
  end
end
