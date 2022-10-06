# frozen_string_literal: true
# typed: true

class StripeForce::InitiatePollsJobs
  include Integrations::Log
  extend Integrations::ErrorContext

  @queue = :high

  def self.perform
    set_error_context
    log.info 'queuing poll jobs'

    StripeForce::User.where(enabled: true).extension(:pagination).each_page(50) do |page|
      page.each do |user|
        begin
          queue_polls_for_user(user)
        rescue => e
          Sentry.capture_exception(e)
        end
      end
    end

    set_error_context
    log.info 'poll queuing complete'
  end

  def self.queue_polls_for_user(user)
    set_error_context(user: user)

    # TODO should check if there is a valid NS + Stripe connection

    queue_poll_job_for_user(user, StripeForce::OrderPoller)

    if @user.feature_enabled?(StripeForce::Constants::FeatureFlags::ACCOUNT_POLLING)
      queue_poll_job_for_user(user, StripeForce::AccountPoller)
    end
  end

  def self.queue_poll_job_for_user(user, poller_job)
    log.info 'queuing poll', poll_job: poller_job

    locker = Integrations::Locker.new(user)
    locker.lock_on_user do
      poller_job.perform(user: user, locker: locker)
    end
  end
end
