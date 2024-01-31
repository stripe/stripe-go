# frozen_string_literal: true
# typed: true

require_relative 'poller_base'

class StripeForce::OrderPoller < StripeForce::PollerBase
  def perform

    execution_time = Time.now.utc

    return if !should_poll?(execution_time, poll_timestamp)

    poll_record = T.must(poll_timestamp)
    end_time = poll_record.last_polled_at

    log.info 'initiating poll',
      from: end_time,
      to: execution_time

    # note that SF fields never contain empty strings, they are reported null instead
    updated_orders = backoff { @user.sf_client.query(generate_soql(end_time, execution_time)) }

    # important to use `size` here and not count because of the paging issue described below
    log.info 'order query complete', size: updated_orders.size

    fail_if_dying_worker!

    # although `updated_orders` looks like a normal array, it's a Restforce::Collection
    # `each` will automatically loop through all pages, but `map` and other array-methods will NOT.
    # make sure `each` is always used here.
    # https://github.com/restforce/restforce/blob/main/lib/restforce/collection.rb#L15
    updated_order_ids = []
    updated_orders.each do |sf_order|
      # refresh lock every 100 records in case we are processing a large batch
      if !locker.nil? && updated_order_ids.length % 100 == 0
        locker.lock_on_poll_job(self.class)
      end

      log.info 'queuing order', sf_order_id: sf_order.Id
      SalesforceTranslateRecordJob.work(@user, sf_order.Id)

      updated_order_ids << sf_order.Id
    end

    # TODO updating the line item does NOT update the order, do we need to worry about this?

    log.info 'poll complete', poll_size: updated_order_ids.count

    poll_record.update(last_polled_at: execution_time)

    updated_order_ids
  end

  private def generate_soql(time_start, time_end)
    # LastModifiedDate has second, but not millisecond, granularity
    <<~EOL
      SELECT Id
      FROM #{poll_type}
      WHERE
      #{prefixed_stripe_field(GENERIC_STRIPE_ID)} = null AND
      LastModifiedDate >= #{time_start.iso8601} AND
      LastModifiedDate < #{time_end.iso8601}
      #{@user.user_specified_where_clause_for_object(poll_type)}
    EOL
  end

  private def poll_type
    SF_ORDER
  end
end
