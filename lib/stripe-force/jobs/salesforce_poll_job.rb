# frozen_string_literal: true
# typed: true

class StripeForce::SalesforcePollJob < StripeForce::BaseJob
    extend T::Sig
    extend T::Helpers

    include Integrations::Log
    include Integrations::ErrorContext

    # We shouldn't schedule retries of the poller, a new instance will be re-queued within 3 minutes
    @retry_limit = 0

    def self.work(user, poll_class)
      Resque.enqueue(
        self,
        user.salesforce_account_id,
        user.stripe_account_id,
        user.livemode,
        poll_class.to_s
      )
    end

    def self.perform(salesforce_account_id, stripe_account_id, livemode, poll_class_name)

      user = user_reference(salesforce_account_id, stripe_account_id, livemode)

      set_error_context(user: user, poll_job: poll_class_name)

      poll_class = allowlist_job_constantize(poll_class_name)

      valid_system_credentials!(user)

      locker = Integrations::Locker.new(user)
      locker.lock_on_user do
        locker.lock_on_poll_job(poll_class)
        poll_class.perform(user: user, locker: locker)
      end

      log.info 'poll complete'

      Integrations::Metrics::Writer.instance.track_counter('poll.processed', dimensions: {
        livemode: user.livemode,
        stripe_account_id: user.stripe_account_id,
        salesforce_account_id: user.salesforce_account_id,
        poller_job: poll_class_name,
      })
    end

    def self.allowlist_job_constantize(name)
      case name
      when 'StripeForce::OrderPoller'
        StripeForce::OrderPoller
      when "StripeForce::AccountPoller"
        StripeForce::AccountPoller
      else
        raise Integrations::Errors::ImpossibleState.new("unexpected salesforce poll class #{name}")
      end
    end
end
