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

    log.info 'queuing poll', poll_job: StripeForce::OrderPoller

    # TODO spit out to a separate job
    locker = Integrations::Locker.new(user)
    locker.lock_on_user do
      StripeForce::OrderPoller.perform(
        user: user,
        locker: locker
      )
    end
  end
end
