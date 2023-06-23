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
          Sentry.capture_exception(e, level: 'error')
        end
      end
    end

    log.info 'poll queuing complete'
  end

  def self.queue_polls_for_user(user)
    set_error_context(user: user)

    unless user.valid_credentials?
      log.info "skipping poll queueing for user with invalid credentials", sf_account_id: user.salesforce_account_id
      return
    end

    unless user.polling_enabled?
      log.info "skipping poll queueing for user with polling disabled", sf_account_id: user.salesforce_account_id
      return
    end

    queue_poll_job_for_user(user, StripeForce::OrderPoller)

    if user.feature_enabled?(StripeForce::Constants::FeatureFlags::ACCOUNT_POLLING)
      queue_poll_job_for_user(user, StripeForce::AccountPoller)
    end
  end

  def self.queue_poll_job_for_user(user, poller_class)
    log.info 'queuing poll', poll_job: poller_class

    StripeForce::SalesforcePollJob.work(user, poller_class)

    Integrations::Metrics::Writer.instance.track_counter('poll.queued', dimensions: {
      livemode: user.livemode,
      stripe_account_id: user.stripe_account_id,
      salesforce_account_id: user.salesforce_account_id,
      poller_job: poller_class.to_s,
    })
  end
end
