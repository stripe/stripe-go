# frozen_string_literal: true
# typed: true
class InitiatePollsJobs
  extend SimpleStructuredLogger
  # extend StripeForce::ErrorContext
  # extend StripeForce::ResqueHooks

  @queue = :high

  def self.perform
    # set_error_context

    log.info 'queuing poll jobs'

    StripeForce::User.where(disabled: false).extension(:pagination).each_page(50) do |page|
      page.each {|user| queue_polls_for_user(user) }
    end

    # set_error_context
    log.info 'poll queuing complete'
  end

  def self.queue_polls_for_user(user)
    # set_error_context(user: user)

    StripeForce::OrderPoller.perform(user: user)
  end
end
