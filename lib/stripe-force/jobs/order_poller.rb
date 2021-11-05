# frozen_string_literal: true
# typed: true

module Integrations
  module Utilities
    include Kernel

    def fail_if_dying_worker!
      raise Integrations::Errors::DyingWorkerError if Resque.heroku_will_terminate?
    end

    def feature?(flag)
      @user.feature_enabled?(flag)
    end
  end
end

class StripeForce::PollBase
  extend T::Sig
  extend T::Helpers

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities

  include StripeForce::Constants

  abstract!

  attr_accessor :locker

  def self.perform(user:, locker: nil)
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

    should_poll = execution_time.to_i - poll_record.last_polled_at.to_i > poll_frequency

    if !should_poll
      log.debug 'skipping poll'
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

  def poll_frequency
    60 * 3
  end

  def sf
    @user.sf_client
  end

  sig { abstract.returns(String) }
  def poll_type; end

  sig { abstract.returns(Array) }
  def perform; end
end

class StripeForce::OrderPoller < StripeForce::PollBase
  def perform
    # TODO lock on job and user in a separate job

    execution_time = Time.now.utc
    poll_record = poll_timestamp

    return if !should_poll?(execution_time, poll_record)

    poll_record = T.must(poll_record)

    log.info 'initiating poll', from: poll_record.last_polled_at, to: execution_time

    updated_orders = sf.get_updated(SF_ORDER, poll_record.last_polled_at, execution_time)
    # TODO anything else but "ids" in the hash? Should verify this
    updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

    fail_if_dying_worker!

    # TODO updating the line item does NOT update the order

    updated_orders.each do |sf_order_id|
      log.info 'translating order', sf_order_id: sf_order_id

      # TODO should split out into a separate job
      sf_order = sf.find(SF_ORDER, sf_order_id)
      StripeForce::Translate.perform(user: @user, sf_object: sf_order)
    end

    log.info 'poll complete'

    poll_record.update(last_polled_at: execution_time)

    updated_orders
  end

  private def poll_type
    SF_ORDER
  end
end
