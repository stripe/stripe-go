# typed: true
# frozen_string_literal: true

class StripeForce::BaseJob
  extend T::Helpers
  extend T::Sig

  include Integrations::Log
  extend Integrations::ErrorContext
  extend Integrations::Metrics::ResqueHooks
  extend Integrations::Errors::ResqueHooks

  @queue = :high
  @retry_limit = 71

  @ignore_exceptions = [Resque::TermException, Integrations::Errors::LockTimeout]
  @retry_exceptions = [Exception, Resque::TermException, Integrations::Errors::LockTimeout]

  # https://github.com/lantins/resque-retry/blob/dc96d1ea9abbe229cccc810498f05b581d830408/lib/resque/plugins/retry.rb#L80
  # https://github.com/lantins/resque-retry/blob/dc96d1ea9abbe229cccc810498f05b581d830408/test/test_jobs.rb#L323
  def self.inherited(subclass)
    super(subclass)

    %w{@queue @ignore_exceptions @retry_exceptions @retry_delay @retry_limit @auto_retry_limit}.each do |variable|
      value = self.instance_variable_get(variable)
      value = value&.dup
      subclass.instance_variable_set(variable, value)
    end
  end

  abstract!

  sig { abstract.params(stripe_account_id: String, stripe_livemode: T::Boolean, sf_record_type: String, sf_record_id: String).void }
  def self.perform(stripe_account_id, stripe_livemode, sf_record_type, sf_record_id); end

  sig { params(stripe_account_id: String, livemode: T::Boolean).returns(StripeForce::User) }
  def self.user_reference(stripe_account_id, livemode)
    user = StripeForce::User.find(stripe_account_id: stripe_account_id, livemode: livemode)

    if user.nil?
      raise StripeForce::Errors::UserError.new("#{stripe_account_id} not available for specified mode. livemode=#{livemode}")
    end

    user
  end

  def self.valid_system_credentials!(user)
    true
  end
end
