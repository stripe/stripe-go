# frozen_string_literal: true
# typed: true

class StripeForce::PollerBase
  extend T::Sig
  extend T::Helpers

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities

  include StripeForce::Constants
  include StripeForce::Utilities::SalesforceUtil

  abstract!

  attr_accessor :locker

  def self.perform(user:, locker:)
    interactor = self.new(user)
    interactor.locker = locker
    interactor.perform
  end

  def initialize(user)
    @user = user
    set_error_context(user: user)
  end

  sig { params(execution_time: Time, poll_record: T.nilable(StripeForce::PollTimestamp)).returns(T::Boolean) }
  def should_poll?(execution_time, poll_record)
    fail_if_dying_worker!

    if poll_record.nil?
      log.warn 'no initial poll timestamp defined, not performing'
      return false
    end

    should_poll = execution_time.to_i - poll_record.last_polled_at.to_i > POLL_FREQUENCY

    if !should_poll
      log.info 'skipping poll'
    end

    should_poll
  end

  sig { returns(T.nilable(StripeForce::PollTimestamp)) }
  def poll_timestamp
    StripeForce::PollTimestamp.by_user_and_record(
      @user,
      poll_type
    )
  end

  def sf
    @user.sf_client
  end

  sig { abstract.returns(String) }
  def poll_type; end

  sig { abstract.returns(Array) }
  def perform; end
end
