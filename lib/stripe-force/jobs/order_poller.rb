# frozen_string_literal: true
# typed: true

class StripeForce::OrderPoller < StripeForce::PollerBase
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
