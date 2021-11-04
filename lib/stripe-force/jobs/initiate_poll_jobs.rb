# frozen_string_literal: true
# typed: true
class StripeForce::InitiatePollsJobs
  include Integrations::Log
  extend Integrations::ErrorContext
  extend Integrations::Metrics::ResqueHooks

  @queue = :high

  def self.perform
    set_error_context

    log.info 'queuing poll jobs'

    StripeForce::User.where(enabled: true).select_all.extension(:pagination).each_page(50) do |page|
      page.each {|user| queue_polls_for_user(user) }
    end

    set_error_context
    log.info 'poll queuing complete'
  end

  def self.queue_polls_for_user(user)
    set_error_context(user: user)

    # TODO spit out to a separate job
    StripeForce::OrderPoller.perform(user: user)
  end
end
