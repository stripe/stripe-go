# frozen_string_literal: true
# typed: true
class StripeForce::InitiatePollsJobs
  extend SimpleStructuredLogger
  # extend StripeForce::ErrorContext
  # extend StripeForce::ResqueHooks

  @queue = :high

  def self.perform
    # set_error_context

    log.info 'queuing poll jobs'

    # TODO need to add disabled flag
    # StripeForce::User.where(disabled: false).extension(:pagination).each_page(50) do |page|
    #   page.each {|user| queue_polls_for_user(user) }
    # end
    StripeForce::User.all.each do |user|
      queue_polls_for_user(user)
    end

    # set_error_context
    log.info 'poll queuing complete'
  end

  def self.queue_polls_for_user(user)
    # set_error_context(user: user)

    # TODO spit out to a separate job
    StripeForce::OrderPoller.perform(user: user)
  end
end
