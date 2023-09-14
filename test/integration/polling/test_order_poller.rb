# frozen_string_literal: true
# typed: ignore

require_relative '../../test_helper'

class Critic::OrderPollerTest < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || DateTime.now.utc)
    @one_min_in_future = VCR.current_cassette.originally_recorded_at ? VCR.current_cassette.originally_recorded_at + 60.seconds : DateTime.now.utc + 60.seconds

    @user = make_user
    # Default to enable polling and individual tests can disable if needed
    @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = true
    @user.save
    inline_job_processing!
  end

  it 'polls orders and does not allow two polls to run at once' do
    initial_poll = set_initial_poll_timestamp(SF_ORDER)
    initial_poll.update(last_polled_at: now_time - POLL_FREQUENCY - 1.minutes)

    # although it's slow, we want to actually hit the live salesforce API
    # to test our query generation logic

    sf_order = create_salesforce_order(contact_email: "one_poll_only")
    contains_order = false

    StripeForce::Translate.expects(:perform).at_least_once.with do |kwargs|
      # return true to skip actually checking this case
      if kwargs[:sf_object].Id != sf_order.Id
        next true
      end

      # make sure the poll actually includes the order we expect
      contains_order = true

      # lock error should be raised if we try lock on another instance of the job
      exception = assert_raises(Integrations::Errors::LockTimeout) do
        initial_poll.update(last_polled_at: now_time - POLL_FREQUENCY - 1.minute)
        StripeForce::SalesforcePollJob.perform(@user.salesforce_account_id, @user.stripe_account_id, @user.livemode, "StripeForce::OrderPoller")
      end

      # key is included in exception, and should include the job
      # we were trying to lock, not the user key
      assert_match("StripeForce::OrderPoller", exception.message)

      true
    end

    # ensures there are no weird propogation issues in SOQL, we've seen some strange timeouts here
    sleep(5)

    # There is some latency to create the order in salesforce, so we need to forward time a few seconds to ensure this order is caught in our poll during recording
    Timecop.freeze(@one_min_in_future)

    StripeForce::SalesforcePollJob.work(@user, StripeForce::OrderPoller)

    assert(contains_order)
  end

  it 'does not poll for a user with polling disabled' do
    @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = false
    @user.save

    StripeForce::SalesforcePollJob.expects(:work).never

    StripeForce::InitiatePollsJobs.perform
  end

  it 'polls for a user with polling enabled' do
    # polling_enabled = true in before step

    StripeForce::SalesforcePollJob.expects(:work).once

    StripeForce::InitiatePollsJobs.perform
  end

  it 'creates a poll timestamp for a user with polling enabled' do
    # polling_enabled = true in before step

    assert_nil(@user.last_polled_timestamp_record)
    assert_nil(@user.last_synced)

    sf_order = create_salesforce_order(contact_email: "creates_a_poll_timestamp")

    SalesforceTranslateRecordJob.expects(:work).at_least_once
    StripeForce::SalesforcePollJob.work(@user, StripeForce::OrderPoller)

    refute_nil(@user.last_polled_timestamp_record)
  end


  it 'creates a poll timestamp for a user with polling enabled and a sync start date defined' do
    # polling_enabled = true in before step
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = "1662681600"
    @user.save

    assert_nil(@user.last_polled_timestamp_record)
    assert_nil(@user.last_synced)

    sf_order = create_salesforce_order(contact_email: "creates_a_poll_timestamp_sync")

    SalesforceTranslateRecordJob.expects(:work).at_least_once
    StripeForce::SalesforcePollJob.work(@user, StripeForce::OrderPoller)

    timestamp_record = @user.last_polled_timestamp_record
    refute_nil(timestamp_record)
    assert_equal(@user.last_synced, timestamp_record.last_polled_at.to_i)
  end

  it 'uses the sync start date' do
    # polling_enabled = true in before step
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = "1662681600"
    @user.save

    assert_nil(@user.last_polled_timestamp_record)
    assert_nil(@user.last_synced)

    op = StripeForce::OrderPoller.new(@user)
    should_poll = op.should_poll?(Time.now.utc, @user.last_polled_timestamp_record)

    assert(should_poll)

    timestamp_record = @user.last_polled_timestamp_record
    refute_nil(timestamp_record)
    assert_equal(@user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE].to_i, timestamp_record.last_polled_at.to_i)
  end

  it 'allows for query customizations' do
    random_address = VCR.current_cassette.originally_recorded_at ? VCR.current_cassette.originally_recorded_at.to_s : DateTime.now.utc.to_s

    @user.connector_settings['filters'][SF_ORDER] = "Account.BillingStreet = '#{random_address}'"

    # must persist user record in order for initial poll job to pick it up
    @user.save

    initial_poll = set_initial_poll_timestamp(SF_ORDER)
    initial_poll.update(last_polled_at: now_time - 5.minutes)

    sf_account_id = create_salesforce_account(additional_fields: {
      "BillingStreet" => random_address,
    })

    sf_account_bad_id = create_salesforce_account(additional_fields: {
      "BillingStreet" => "bad",
    })

    sf_order_1 = create_salesforce_order(sf_account_id: sf_account_id, contact_email: "query_customizations")
    sf_order_2 = create_salesforce_order(sf_account_id: sf_account_bad_id, contact_email: "query_customizations_2")

    contains_order = false
    excludes_condition = true

    # because of the randomized query condition, only a single record should be returned
    # this implies that the dynamic condition specified at the top of this test is applied
    StripeForce::Translate.expects(:perform).once.with do |kwargs|
      assert_equal(kwargs[:sf_object].Id, sf_order_1.Id)
    end

    # ensure the ending time is not in the same second
    sleep(1)
    one_min_in_future = VCR.current_cassette.originally_recorded_at ? VCR.current_cassette.originally_recorded_at + 60.seconds : DateTime.now.utc + 60.seconds
    Timecop.freeze(one_min_in_future)

    StripeForce::SalesforcePollJob.perform(@user.salesforce_account_id, @user.stripe_account_id, @user.livemode, "StripeForce::OrderPoller")
  end

  it 'refreshes the poll lock if there are locks of records to process' do
    Timecop.return

    inline_job_processing!

    initial_poll = set_initial_poll_timestamp(SF_ORDER)

    # lock is refreshed every 100 records
    record_total = 105
    stub_salesforce_query_result(number_of_results: record_total)

    lock_key = nil
    initial_lock_value = nil
    final_lock_value = nil
    queued_count = 0

    SalesforceTranslateRecordJob.expects(:work).at_least_once.with do |*args|
      # sleep for ~second to ensure the expiration time will change when refreshed
      sleep 0.01

      queued_count += 1

      if queued_count == record_total
        final_lock_value = redis.get(lock_key)
      end

      # the first time `with` is run the poll job lock value would have already been set
      # the key of the poll job lock key contains the poll job class, let's search redis keys
      # for a poll job lock key and retrieve the expiration value to ensure it's refreshed
      # later in the processing loop
      if initial_lock_value.nil?
        lock_key = redis.keys.detect {|k| k.include?('StripeForce::OrderPoller') }
        initial_lock_value = redis.get(lock_key)
      end

      args[0].id == @user.id
    end

    StripeForce::SalesforcePollJob.work(@user, StripeForce::OrderPoller)

    refute_nil(initial_lock_value)
    refute_nil(final_lock_value)
    assert(final_lock_value > initial_lock_value)
  end

  it 'does not run polls if salesforce credentials are invalid' do
    StripeForce::BaseJob.stubs(:valid_system_credentials!).raises(Integrations::Errors::InvalidCredentialsError)

    StripeForce::OrderPoller.expects(:perform).never

    assert_raises(Integrations::Errors::InvalidCredentialsError) { StripeForce::SalesforcePollJob.work(@user, StripeForce::OrderPoller) }
  end

  # order poller is used for the remainder of the test suite since it is most commonly used
  describe 'basic test for all pollers' do

    # TODO when we add support for additional poll types test them here
    it 'StripeForce::OrderPoller' do

      initial_poll = set_initial_poll_timestamp(SF_ORDER)

      stubbed_ids = stub_salesforce_query_result

      SalesforceTranslateRecordJob.expects(:work).once.with do |*args|
        args[0] == @user && args[1] == stubbed_ids.first.Id
      end

      order_poller = StripeForce::OrderPoller.new(@user)
      order_poller.perform
    end

    it 'StripeForce::AccountPoller' do

      initial_poll = set_initial_poll_timestamp(SF_ACCOUNT)

      stubbed_ids = stub_salesforce_query_result

      SalesforceTranslateRecordJob.expects(:work).once.with do |*args|
        args[0] == @user && args[1] == stubbed_ids.first.Id
      end

      account_poller = StripeForce::AccountPoller.new(@user)
      account_poller.perform
    end

  end


  describe 'general tests using order poller' do
    it 'does not poll if polling is disabled' do
      @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = false
      @user.save

      Restforce::Data::Client.any_instance.expects(:get_updated).never

      locker = Integrations::Locker.new(@user)
      invoice_poller = StripeForce::OrderPoller.perform(user: @user, locker: locker)

      assert_equal(0, StripeForce::PollTimestamp.count)
    end

    it 'polls if polling is enabled' do
      initial_poll = set_initial_poll_timestamp(SF_ORDER).last_polled_at

      SalesforceTranslateRecordJob.expects(:perform)

      stub_salesforce_query_result

      locker = Integrations::Locker.new(@user)
      Timecop.freeze(@one_min_in_future)
      StripeForce::OrderPoller.perform(user: @user, locker: locker)

      assert_equal(1, StripeForce::PollTimestamp.count)

      poll_timestamp = StripeForce::PollTimestamp.by_user_and_record(
        @user,
        SF_ORDER
      )

      assert(poll_timestamp.last_polled_at - initial_poll > initial_poll_delta)
      assert_equal(SF_ORDER, poll_timestamp.integration_record_type)
    end
  end
end