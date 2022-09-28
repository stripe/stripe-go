# Usage: bundle exec sequel -E -m migrations $DATABASE_URL

# frozen_string_literal: true
# typed: false

require 'sequel'

Sequel.migration do
  up do
    add_column :users, :salesforce_object_prefix_mappings, :jsonb, default: '{}', null: false
  end

  down do
    drop_column :users, :salesforce_object_prefix_mappings
  end
end
