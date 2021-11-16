# frozen_string_literal: true
# typed: true

require_relative 'poller_base'

class StripeForce::OrderPoller < StripeForce::PollerBase
  def perform
    # TODO lock on job and user in a separate job

    execution_time = Time.now.utc
    poll_record = poll_timestamp

    return if !should_poll?(execution_time, poll_record)

    poll_record = T.must(poll_record)

    log.info 'initiating poll', from: poll_record.last_polled_at, to: execution_time

    # TODO filter by activated orders without a Stripe ID
    updated_orders = sf.get_updated(SF_ORDER, poll_record.last_polled_at, execution_time)

    # TODO anything else but "ids" in the hash? Should verify this
    updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

    fail_if_dying_worker!

    # TODO should refresh the poll lock
    # TODO updating the line item does NOT update the order

    updated_orders.each do |sf_order_id|
      log.info 'queuing order', sf_order_id: sf_order_id
      SalesforceTranslateRecordJob.work(@user, SF_ORDER, sf_order_id)
    end

    log.info 'poll complete'

    poll_record.update(last_polled_at: execution_time)

    updated_orders
  end

  private def poll_type
    SF_ORDER
  end
end
