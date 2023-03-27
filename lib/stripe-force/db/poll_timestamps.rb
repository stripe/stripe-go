# frozen_string_literal: true
# typed: true

module StripeForce
  class PollTimestamp < Sequel::Model
    include StripeForce::Constants
    plugin :timestamps, update_on_create: true

    def self.by_user_and_record(user, sf_record_class)
      self.find(build_timestamp_search(user, sf_record_class))
    end

    def self.build_with_user_and_record(user, sf_record_class)

      timestamp = nil
      unless user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE].nil?
        timestamp = user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE].to_i
      end

      timestamp ||= (Time.now - 5.minutes).to_i

      self.new(build_timestamp_search(user, sf_record_class).merge(last_polled_at: Time.at(timestamp)))
    end

    def self.build_timestamp_search(user, sf_record_class)
      {
        salesforce_account_id: user.salesforce_account_id,
        livemode: user.livemode,
        integration_record_type: sf_record_class.to_s,
      }
    end

  end
end
