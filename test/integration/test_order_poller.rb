# frozen_string_literal: true
# typed: ignore

require_relative '../test_helper'

class Critic::OrderPollerTest < Critic::UnitTest
  before do
    @user = make_user
    inline_job_processing!
  end

  def initial_poll_delta; 60 * 60 * 24 end

  def set_initial_poll_timestamp(sf_class)
    initial_poll = DateTime.now - initial_poll_delta

    poll_timestamp = StripeForce::PollTimestamp.build_with_user_and_record(
      @user,
      sf_class
    )

    poll_timestamp.last_polled_at = initial_poll
    poll_timestamp.save

    poll_timestamp
  end

  it 'polls orders and does not allow two polls to run at once' do
    # must persist user record in order for initial poll job to pick it up
    @user.save

    initial_poll = set_initial_poll_timestamp(SF_ORDER)
    initial_poll.update(last_polled_at: DateTime.now - 5.minutes)

    # although it's slow, we want to actually hit the live salesforce API
    # to test our query generation logic

    sf_order = create_salesforce_order
    contains_order = false

    StripeForce::Translate.expects(:perform).at_least_once.with do |kwargs|
      # return true to skip actually checking this case
      if kwargs[:sf_object].Id != sf_order.Id
        next true
      end

      # lock error should be raised if we try lock on another instance of the job
      exception = assert_raises(Integrations::Errors::LockTimeout) do
        initial_poll.update(last_polled_at: DateTime.now - 5.minutes)
        StripeForce::InitiatePollsJobs.perform
      end

      # key is included in exception, and should include the job
      # we were trying to lock, not the user key
      assert_match("StripeForce::OrderPoller", exception.message)

      contains_order = true
      true
    end

    sleep(1)

    # TODO right now, this works since we only have a single order poll, this will need to be specific to orders in the future
    StripeForce::InitiatePollsJobs.perform

    assert(contains_order)
  end

  it 'allows for query customizations' do
    random_address = SecureRandom.alphanumeric(16)

    @user.connector_settings['filters'][SF_ORDER] = "Account.BillingStreet = '#{random_address}'"

    # must persist user record in order for initial poll job to pick it up
    @user.save

    initial_poll = set_initial_poll_timestamp(SF_ORDER)
    initial_poll.update(last_polled_at: DateTime.now - 5.minutes)

    sf_account_id = create_salesforce_account(additional_fields: {
      "BillingStreet" => random_address,
    })

    sf_account_bad_id = create_salesforce_account(additional_fields: {
      "BillingStreet" => "bad",
    })

    sf_order_1 = create_salesforce_order(sf_account_id: sf_account_id)
    sf_order_2 = create_salesforce_order(sf_account_id: sf_account_bad_id)

    contains_order = false
    excludes_condition = true

    # because of the randomized query condition, only a single record should be returned
    # this implies that the dynamic condition specified at the top of this test is applied
    StripeForce::Translate.expects(:perform).once.with do |kwargs|
      assert_equal(kwargs[:sf_object].Id, sf_order_1.Id)
    end

    # ensure the ending time is not in the same second
    sleep(1)

    # TODO right now, this works since we only have a single order poll, this will need to be specific to orders in the future
    StripeForce::InitiatePollsJobs.perform
  end

  it 'refreshes the poll lock if there are locks of records to process' do
    skip("to implement after we split out the poll job into a separate resque job")

    inline_job_processing!

    initial_poll = set_initial_poll_timestamp(NetSuite::Records::Invoice)

    # lock is refreshed every 100 records
    record_total = 105
    stub_netsuite_batch_search_result(NetSuite::Records::Invoice, number_of_results: record_total)

    lock_key = nil
    initial_lock_value = nil
    final_lock_value = nil
    queued_count = 0

    NetsuiteProcessRecordJob.expects(:work).at_least_once.with do |*args|
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
        lock_key = redis.keys.detect {|k| k.include?('SuiteSync::InvoicePoller') }
        initial_lock_value = redis.get(lock_key)
      end

      args[0].id == @user.id &&
        args[1] == SuiteSync::InvoiceProcessor
    end

    NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)

    refute_nil(initial_lock_value)
    refute_nil(final_lock_value)
    assert(final_lock_value > initial_lock_value)
  end

  it 'does not run polls if salesforce credentials are invalid' do
    skip("implement after polling is in a separate job")

    refute(@user.status.valid_netsuite_credentials?(refresh: true))

    SuiteSync::InvoicePoller.expects(:perform).never

    NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)
  end

  # order poller is used for the remainder of the test suite since it is most commonly used
  describe 'basic test for all pollers' do
    def expect_tracked_search_results(records:, poller:)
      StripeSuite::Metrics::Writer.instance.expects(:track_counter).once.with do |*args|
        args.first == StripeSuite::Metrics::NETSUITE_POLL_RECORDS_COUNTER &&
          args.last == {value: records, dimensions: {poller: poller}}
      end
    end

    def expect_timed_search(poller:)
      Integrations::Metrics::Writer.instance.expects(:track_gauge).once.with do |metric, _, dimensions|
        metric == Integrations::Metrics::NETSUITE_SEARCH_TIME &&
          dimensions == {dimensions: {poller: poller}}
      end
    end

    # TODO when we add support for additional poll types test them here
    it 'SuiteSync::CashSalePoller' do
      skip("stop referencing netsuite")

      initial_poll = set_initial_poll_timestamp(NetSuite::Records::CashSale)

      stub_netsuite_batch_search_result(NetSuite::Records::CashSale)

      NetsuiteProcessRecordJob.expects(:work).once.with do |*args|
        args[0].id == @user.id &&
          args[1] == SuiteSync::CashSaleProcessor &&
          args[2] == ns_record_id
      end

      expect_tracked_search_results(records: 1, poller: "SuiteSync::CashSalePoller")
      expect_timed_search(poller: "SuiteSync::CashSalePoller")

      cash_sale_poller = SuiteSync::CashSalePoller.new(user: @user)
      cash_sale_poller.perform
    end

  end

  describe 'general tests using order poller' do
    it 'does not poll if no initial timestamp is set' do
      Restforce::Data::Client.any_instance.expects(:get_updated).never

      locker = Integrations::Locker.new(@user)
      invoice_poller = StripeForce::OrderPoller.perform(user: @user, locker: locker)

      assert_equal(0, StripeForce::PollTimestamp.count)
    end

    it 'polls if a timestamp is set' do
      initial_poll = set_initial_poll_timestamp(SF_ORDER).last_polled_at

      SalesforceTranslateRecordJob.expects(:perform)

      Restforce::Data::Client.any_instance.expects(:query).returns([
        Restforce::SObject.new({"Id" => create_salesforce_id}),
      ])

      locker = Integrations::Locker.new(@user)
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
