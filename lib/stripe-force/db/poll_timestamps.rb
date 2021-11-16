# frozen_string_literal: true
# typed: true

module StripeForce
  class PollTimestamp < Sequel::Model
    plugin :timestamps, update_on_create: true

    def self.by_user_and_record(user, ns_record_class)
      self.find(build_timestamp_search(user, ns_record_class))
    end

    def self.build_with_user_and_record(user, ns_record_class)
      self.new(build_timestamp_search(user, ns_record_class))
    end

    def self.build_timestamp_search(user, ns_record_class)
      {
        salesforce_account_id: user.salesforce_account_id,
        livemode: user.livemode,
        # sandbox: user.sandbox?,
        integration_record_type: ns_record_class.to_s,
      }
    end

  end
end
