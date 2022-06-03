# frozen_string_literal: true
# typed: true

require_relative 'poller_base'

class StripeForce::OrderPoller < StripeForce::PollerBase
  def perform
    # TODO lock on job in a separate wrapper job
    locker.lock_on_poll_job(self.class)

    execution_time = Time.now.utc
    poll_record = poll_timestamp

    return if !should_poll?(execution_time, poll_record)

    poll_record = T.must(poll_record)

    log.info 'initiating poll', from: poll_record.last_polled_at, to: execution_time

    # note that SF fields never contain empty strings, they are reported null instead
    updated_orders = sf.query(
      <<~EOL
        SELECT Id
        FROM #{SF_ORDER}
        WHERE
        #{prefixed_stripe_field(GENERIC_STRIPE_ID)} = null AND
        LastModifiedDate >= #{poll_record.last_polled_at.iso8601} AND
        LastModifiedDate < #{execution_time.iso8601}
        #{user_specified_where_clause_for_object}
      EOL
    )

    # important to use `size` here and not count because of the paging issue described below
    log.info 'order query complete', size: updated_orders.size

    # although `updated_orders` looks like a normal array, it's a Restforce::Collection
    # `each` will automatically loop through all pages, but `map` and other array-methods will NOT.
    # make sure `each` is always used here.
    # https://github.com/restforce/restforce/blob/685900a6c3d24f860f18b38448eda7b39187b2fe/lib/restforce/collection.rb#L15

    updated_order_ids = []
    updated_orders.each do |sf_order|
      updated_order_ids << sf_order.Id
    end

    fail_if_dying_worker!

    # TODO should refresh the poll lock
    # TODO updating the line item does NOT update the order, do we need to worry about this?

    updated_order_ids.each do |sf_order_id|
      log.info 'queuing order', sf_order_id: sf_order_id
      SalesforceTranslateRecordJob.work(@user, SF_ORDER, sf_order_id)
    end

    log.info 'poll complete', poll_size: updated_order_ids.count

    poll_record.update(last_polled_at: execution_time)

    updated_order_ids
  end

  private def user_specified_where_clause_for_object
    custom_query = @user.connector_settings.dig("filters", poll_type)

    if custom_query && custom_query.strip.present?
      log.info 'adding custom query to poll', poll_type: poll_type
      " AND " + custom_query
    else
      ""
    end
  end

  private def poll_type
    SF_ORDER
  end
end
