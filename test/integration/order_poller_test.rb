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

    initial_poll
  end

  # it 'polls invoices and does not allow two polls to run at once' do
  #   initial_poll = set_initial_poll_timestamp(SF_ORDER)

  #   # stub_netsuite_batch_search_result(NetSuite::Records::Invoice)

  #   StripeForce::Translate.expects(:perform).once.with do |*args|
  #     # lock error should be raised if we try lock on another instance of the job
  #     exception = assert_raises(StripeSuite::Errors::LockTimeout) do
  #       NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)
  #     end

  #     # key is included in exception, and should include the job
  #     # we were trying to lock, not the user key
  #     assert_match("SuiteSync::InvoicePoller", exception.message)

  #     args[0].id == @user.id &&
  #       args[1] == SuiteSync::InvoiceProcessor &&
  #       args[2] == ns_record_id
  #   end

  #   NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)
  # end

  # it 'refreshes the poll lock if there are locks of records to process' do
  #   inline_job_processing!

  #   initial_poll = set_initial_poll_timestamp(NetSuite::Records::Invoice)

  #   # lock is refreshed every 100 records
  #   record_total = 105
  #   stub_netsuite_batch_search_result(NetSuite::Records::Invoice, number_of_results: record_total)

  #   lock_key = nil
  #   initial_lock_value = nil
  #   final_lock_value = nil
  #   queued_count = 0

  #   NetsuiteProcessRecordJob.expects(:work).at_least_once.with do |*args|
  #     # sleep for ~second to ensure the expiration time will change when refreshed
  #     sleep 0.01

  #     queued_count += 1

  #     if queued_count == record_total
  #       final_lock_value = redis.get(lock_key)
  #     end

  #     # the first time `with` is run the poll job lock value would have already been set
  #     # the key of the poll job lock key contains the poll job class, let's search redis keys
  #     # for a poll job lock key and retrieve the expiration value to ensure it's refreshed
  #     # later in the processing loop
  #     if initial_lock_value.nil?
  #       lock_key = redis.keys.detect {|k| k.include?('SuiteSync::InvoicePoller') }
  #       initial_lock_value = redis.get(lock_key)
  #     end

  #     args[0].id == @user.id &&
  #       args[1] == SuiteSync::InvoiceProcessor
  #   end

  #   NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)

  #   refute_nil(initial_lock_value)
  #   refute_nil(final_lock_value)
  #   assert(final_lock_value > initial_lock_value)
  # end

  # it 'does not run polls if netsuite credentials are invalid' do
  #   @user.netsuite_email = 'bad-email@example.com'
  #   @user.netsuite_password = 'bad-password'

  #   clear_token_based_auth(@user)

  #   assert_equal(false, @user.status.valid_netsuite_credentials?(refresh: true))

  #   SuiteSync::InvoicePoller.expects(:perform).never

  #   NetsuitePollJob.work(@user, SuiteSync::InvoicePoller)
  # end

  # order poller is used for the remainder of the test suite since it is most commonly used
  describe 'basic test for all pollers' do
    def expect_tracked_search_results(records:, poller:)
      StripeSuite::Metrics::Writer.instance.expects(:track_counter).once.with do |*args|
        args.first == StripeSuite::Metrics::NETSUITE_POLL_RECORDS_COUNTER &&
          args.last == {value: records, dimensions: {poller: poller}}
      end
    end

    def expect_timed_search(poller:)
      StripeSuite::Metrics::Writer.instance.expects(:track_gauge).once.with do |metric, _, dimensions|
        metric == StripeSuite::Metrics::NETSUITE_SEARCH_TIME &&
          dimensions == {dimensions: {poller: poller}}
      end
    end

    it 'SuiteSync::CashSalePoller' do
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

    it 'SuiteSync::CashRefundPoller' do
      initial_poll = set_initial_poll_timestamp(NetSuite::Records::CashRefund)

      stub_netsuite_batch_search_result(NetSuite::Records::CashRefund)

      NetsuiteProcessRecordJob.expects(:work).once.with do |*args|
        args[0].id == @user.id &&
          args[1] == SuiteSync::CashRefundProcessor &&
          args[2] == ns_record_id
      end

      expect_tracked_search_results(records: 1, poller: "SuiteSync::CashRefundPoller")
      expect_timed_search(poller: "SuiteSync::CashRefundPoller")

      cash_refund_poller = SuiteSync::CashRefundPoller.new(user: @user)
      cash_refund_poller.perform
    end

    it 'SuiteSync::CreditMemoPoller' do
      initial_poll = set_initial_poll_timestamp(NetSuite::Records::CreditMemo)

      stub_netsuite_batch_search_result(NetSuite::Records::CreditMemo)

      NetsuiteProcessRecordJob.expects(:work).once.with do |*args|
        args[0].id == @user.id &&
          args[1] == SuiteSync::CreditMemoProcessor &&
          args[2] == ns_record_id
      end

      expect_tracked_search_results(records: 1, poller: "SuiteSync::CreditMemoPoller")
      expect_timed_search(poller: "SuiteSync::CreditMemoPoller")

      credit_memo_poller = SuiteSync::CreditMemoPoller.new(user: @user)
      credit_memo_poller.perform
    end

    it 'SuiteSync::SalesOrderPoller' do
      initial_poll = set_initial_poll_timestamp(NetSuite::Records::SalesOrder)

      stub_netsuite_batch_search_result(NetSuite::Records::SalesOrder)

      NetsuiteProcessRecordJob.expects(:work).once.with do |*args|
        args[0].id == @user.id &&
          args[1] == SuiteSync::SalesOrderProcessor &&
          args[2] == ns_record_id
      end

      expect_tracked_search_results(records: 1, poller: "SuiteSync::SalesOrderPoller")
      expect_timed_search(poller: "SuiteSync::SalesOrderPoller")

      sales_order_poller = SuiteSync::SalesOrderPoller.new(user: @user)
      sales_order_poller.perform
    end

    it 'SuiteSync::CustomerRefundPoller' do
      initial_poll = set_initial_poll_timestamp(NetSuite::Records::CustomerRefund)

      stub_netsuite_batch_search_result(NetSuite::Records::CustomerRefund)

      NetsuiteProcessRecordJob.expects(:work).once.with do |*args|
        args[0].id == @user.id &&
          args[1] == SuiteSync::CustomerRefundProcessor &&
          args[2] == ns_record_id
      end

      expect_tracked_search_results(records: 1, poller: "SuiteSync::CustomerRefundPoller")
      expect_timed_search(poller: "SuiteSync::CustomerRefundPoller")

      customer_refund_poller = SuiteSync::CustomerRefundPoller.new(user: @user)
      customer_refund_poller.perform
    end
  end

  describe 'general tests using invoice poller' do
    it 'does not poll if no initial timestamp is set' do
      Restforce::Data::Client.any_instance.expects(:get_updated).never

      invoice_poller = StripeForce::OrderPoller.perform(user: @user)

      assert_equal(0, StripeForce::PollTimestamp.count)
    end

    it 'polls if a timestamp is set' do
      initial_poll = set_initial_poll_timestamp(SF_ORDER)

      stub_netsuite_batch_search_result(NetSuite::Records::Invoice)

      NetsuiteProcessRecordJob.expects(:work).once

      SuiteSync::InvoicePoller.perform(user: @user)

      assert_equal(1, SuiteSync::NetsuitePollTimestamp.count)

      poll_timestamp = SuiteSync::NetsuitePollTimestamp.by_user_and_record(
        @user,
        NetSuite::Records::Invoice
      )

      assert(poll_timestamp.last_polled_at - initial_poll > initial_poll_delta)
      assert_equal("NetSuite::Records::Invoice", poll_timestamp.netsuite_record_type)
    end

    it 'polls unpaid invoices if a timestamp is set' do
      @user.enable_feature(:poll_paid_invoices)
      set_initial_poll_timestamp(NetSuite::Records::Invoice)

      # TODO why aren't we using the convenience method for stubbing here?

      search_results = OpenStruct.new(
        total_records: 1,
        total_pages: 1,
        page_index: 1
      )

      search_results.expects(:results_in_batches).twice.yields([NetSuite::Records::Invoice.new(internal_id: ns_record_id)])
      NetSuite::Records::Invoice.expects(:search).twice.returns(search_results)

      NetsuiteProcessRecordJob.expects(:work).twice

      SuiteSync::InvoicePoller.perform(user: @user)

      assert_equal(1, SuiteSync::NetsuitePollTimestamp.count)
    end

    # it 'polls credit memos if a timestamp is set'

    it 'properly calculates the netsuite time drift' do
      invoice_poller = SuiteSync::InvoicePoller.new(user: @user)
      drift = invoice_poller.send(:netsuite_drift)

      refute_nil(drift)
    end
  end

  describe 'credit memo poller' do
    it 'adjusts poll window by a delay amount' do
      @user.netsuite_refund_delay = 60
      @user.enable_feature(:poll_credit_memos)

      initial_poll = set_initial_poll_timestamp(NetSuite::Records::CreditMemo)
      current_time = DateTime.now

      search_results = OpenStruct.new(
        total_records: 1,
        total_pages: 1,
        page_index: 1
      )

      NetsuiteProcessRecordJob.expects(:work).once

      # not using convience mocking method since we are mocking out the search

      search_results.expects(:results_in_batches).once.yields([NetSuite::Records::CreditMemo.new(internal_id: ns_record_id)])
      NetSuite::Records::CreditMemo.expects(:search).with do |criteria|
        basic_criteria = criteria[:criteria][:basic]

        within_criteria = basic_criteria.detect {|c| c[:field] == 'lastModifiedDate' && c[:operator] == 'within' }

        refute_nil(within_criteria)

        after_time = DateTime.parse(within_criteria[:value].first)
        before_time = DateTime.parse(within_criteria[:value].last)

        # 30s for drift & 15 second adjustment
        poll_adjustment = initial_poll.to_i - after_time.to_i - @user.netsuite_refund_delay
        assert(poll_adjustment < 30,
          "poll adjustment (#{poll_adjustment}) is not within a reasonsable window. Most likely this is an intermittent NetSuite failure. " \
          "but if it continues to occur it could indicate a bug (or that NetSuite is taking forever to accept a open TCP connection)."
        )
        assert((DateTime.now.to_i - before_time.to_i - @user.netsuite_refund_delay).abs < 1, "before window beyond reasonsable window")
      end.returns(search_results)

      SuiteSync::CreditMemoPoller.perform(user: @user)
    end
  end

end
