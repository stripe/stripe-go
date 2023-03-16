# typed: true
# frozen_string_literal: true

class StripeForce::BaseJob
  extend T::Helpers
  extend T::Sig

  include Integrations::Log
  extend Integrations::ErrorContext
  extend Integrations::Metrics::ResqueHooks

  extend StripeForce::Utilities::SalesforceUtil

  extend Resque::Plugins::Retry

  @queue = :high

  # the first delay will be 0 seconds, the 2nd will be 60 seconds, etc...
  @backoff_strategy = [0, 60, 300, 600, 1800, 3600, 3600, 3600, 10800, 21600]

  # the delay values will be multiplied by a random float between min and max
  @retry_delay_multiplicand_min = 1.0
  @retry_delay_multiplicand_max = 2.0

  @ignore_exceptions = [Resque::TermException, Integrations::Errors::LockTimeout]
  @retry_exceptions = [Exception, Resque::TermException, Integrations::Errors::LockTimeout]

  give_up_callback do |exception, *args|
    log.error "not retrying job",
      exception: exception,
      job: self.name,
      salesforce_account_id: args[0],
      stripe_user_id: args[1],
      livemode: args[2],
      args: args[3..-1]
  end

  # https://github.com/lantins/resque-retry/blob/dc96d1ea9abbe229cccc810498f05b581d830408/lib/resque/plugins/retry.rb#L80
  # https://github.com/lantins/resque-retry/blob/dc96d1ea9abbe229cccc810498f05b581d830408/test/test_jobs.rb#L323
  def self.inherited(subclass)
    super(subclass)

    %w{@queue @ignore_exceptions @retry_exceptions @backoff_strategy @auto_retry_limit}.each do |variable|
      value = self.instance_variable_get(variable)
      value = value&.dup
      subclass.instance_variable_set(variable, value)
    end
  end

  abstract!

  sig { abstract.params(salesforce_account_id: String, stripe_account_id: String, stripe_livemode: T::Boolean, sf_record_id: String).void }
  def self.perform(salesforce_account_id, stripe_account_id, stripe_livemode, sf_record_id); end

  sig { params(salesforce_account_id: String, stripe_account_id: String, livemode: T::Boolean).returns(StripeForce::User) }
  def self.user_reference(salesforce_account_id, stripe_account_id, livemode)
    user = StripeForce::User.find(
      salesforce_account_id: salesforce_account_id,
      stripe_account_id: stripe_account_id,
      livemode: livemode
    )

    if user.nil?
      raise StripeForce::Errors::UserError.new("#{stripe_account_id} not available for specified mode. livemode=#{livemode}")
    end

    user
  end

  def self.valid_system_credentials!(user)
    unless user.valid_credentials?
      raise Integrations::Errors::InvalidCredentialsError
    end
  end
end
