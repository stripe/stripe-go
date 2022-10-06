# frozen_string_literal: true
# typed: true

require_relative 'poller_base'

class StripeForce::AccountPoller < StripeForce::PollerBase
  def perform
    locker.lock_on_poll_job(self.class)

    current_execution_time = Time.now.utc
    poll_record = poll_timestamp

    return if !should_poll?(current_execution_time, poll_record)

    poll_record = T.must(poll_record)
    last_polled_at = poll_record.last_polled_at

    log.info 'initiating poll',
      from: last_polled_at,
      to: current_execution_time

    updated_sf_accounts = backoff { @user.sf_client.query(generate_soql(last_polled_at, current_execution_time)) }

    sf_account_ids_to_sync = []
    updated_sf_accounts.each do |sf_account|
      sf_account_ids_to_sync << sf_account.Id
    end

    log.info 'account query complete', size: sf_account_ids_to_sync.size

    fail_if_dying_worker!

    sf_account_ids_to_sync.each do |sf_account_id|
      log.info 'queuing account', sf_account_id: sf_account_id
      SalesforceTranslateRecordJob.work(@user, sf_account_id)
    end

    log.info 'poll complete', poll_size: sf_account_ids_to_sync.count

    poll_record.update(last_polled_at: current_execution_time)

    sf_account_ids_to_sync
  end

  # soql query to fetch all the sf account ids that:
  #  have our custom stripe field, indicating the sf account has been synced to stripe
  #  have a last modified time between the start and end time
  private def generate_soql(time_start, time_end)
    <<~EOL
      SELECT Id
      FROM #{poll_type}
      WHERE
      #{prefixed_stripe_field(GENERIC_STRIPE_ID)} != null AND
      LastModifiedDate >= #{time_start.iso8601} AND
      LastModifiedDate < #{time_end.iso8601}
    EOL
  end

  private def poll_type
    SF_ACCOUNT
  end
end
