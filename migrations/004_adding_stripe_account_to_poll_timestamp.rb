# Usage: bundle exec sequel -E -m migrations $DATABASE_URL


# frozen_string_literal: true
# typed: false

# Note: the alters are separated due to race conditions. Dropping and adding constraints occurred simultaneously without separation

Sequel.migration do
    up do
      alter_table(:poll_timestamps) do
        add_column :stripe_account_id, String
      end

      alter_table(:poll_timestamps) do
        drop_constraint('poll_timestamps_salesforce_account_id_livemode_integration__key')
      end

      alter_table(:poll_timestamps) do
        add_unique_constraint([:salesforce_account_id, :livemode, :stripe_account_id, :integration_record_type], name: 'salesforce_livemode_stripe_integration_key')
      end
    end

    down do
      alter_table(:poll_timestamps) do
        drop_constraint('salesforce_livemode_stripe_integration_key')
      end

      alter_table(:poll_timestamps) do
        add_unique_constraint([:salesforce_account_id, :livemode, :integration_record_type], name: 'poll_timestamps_salesforce_account_id_livemode_integration__key')
      end

      alter_table(:poll_timestamps) do
        drop_column :stripe_account_id
      end
    end
end
