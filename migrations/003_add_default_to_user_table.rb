# Usage: bundle exec sequel -E -m migrations $DATABASE_URL

# frozen_string_literal: true
# typed: false

require 'sequel'

Sequel.migration do
  up do
    add_column :users, :is_default_account_config, :boolean, default: true, null: false
  end

  down do
    drop_column :users, :is_default_account_config
  end
end
