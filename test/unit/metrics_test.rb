# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class MetricsTest < Critic::UnitTest
    before do
      # avoid using `instance` to ensure test mutations below don't effect the state of the global metrics object
      @metrics = StripeForce::Metrics::Writer.send(:new)
    end

    def drain_queue
      @metrics.send(:drain_queue)
    end

    def collate_queue(data)
      @metrics.send(:collate_queue, data)
    end

    it 'properly collates items' do
      5.times { @metrics.track_counter('foo') }
      7.times { @metrics.track_gauge('bar', 1.0) }

      data = drain_queue
      assert_equal(data.length, 12)
      data = collate_queue(data)
      assert_equal(data[:counter].length, 5)
      assert_equal(data[:gauge].length, 7)
    end

    it "handles additional dimensions" do
      7.times {|i| @metrics.track_gauge('bar', 1.0, dimensions: {instigator: "chris", index: i}) }

      data = drain_queue

      assert_equal(7, data.length)
      assert_equal(["chris"] * 7, data.map {|d| d.data[:dimensions][:instigator] })
    end

    it 'sorts items by monotonically increasing timestamp' do
      10.times { @metrics.track_counter('foo') }
      data = drain_queue

      10.times.zip(data) do |n, item|
        item.data[:timestamp] += n
      end

      data.shuffle!

      counter_data = collate_queue(data)[:counter]

      assert(counter_data.first[:timestamp] < counter_data.last[:timestamp])
    end

    it 'drains the queue' do
      assert_equal(drain_queue, [])
      10.times { @metrics.track_counter('foo') }
      assert_equal(drain_queue.length, 10)
      assert_equal(drain_queue, [])
    end

    it 'respects maximum queue length' do
      original_max_queue = @metrics.max_queue_length
      @metrics.max_queue_length = 5
      10.times { @metrics.track_counter('foo') }
      assert_equal(@metrics.queue.length, 5)
      assert_equal(drain_queue.length, 5)
    end

    it 'handles basic functionality' do
      5.times { @metrics.track_counter('foo') }
      7.times { @metrics.track_gauge('bar', 1.0) }

      assert_equal(@metrics.queue.length, 12)
      RestClient::Request.expects(:execute)
      @metrics.send(:process)
      assert_equal(@metrics.queue.length, 0)
    end

    it 'allows metrics to be sent synchronously' do
      5.times { @metrics.track_counter('foo') }
      RestClient::Request.expects(:execute)

      # shutdown thread to ensure metrics can only be send through the primary thread
      @metrics.timer.shutdown

      @metrics.send_metrics_synchronously

      assert_equal(@metrics.queue.length, 0)
    end

    describe '#time' do

      it 'tracks time as a guage' do
        duration = 0.1 # 100 ms
        @metrics.expects(:track_gauge).once.with do |metric, reported_time|
          metric == :hundred_yard_dash &&
            reported_time.between?(
              duration * 1000, # at least as long as the block
              (duration + 0.005) * 1000 # no more than 5 ms more than the block
            )
        end
        @metrics.time(:hundred_yard_dash) { sleep duration }
      end

      it 'tracks time even when there are exceptions' do
        @metrics.expects(:track_gauge).once.with {|metric, _| metric == :rescuers }
        begin
          @metrics.time(:rescuers) { raise 'Joanna!' }
        rescue
          nil
        end
      end

      it 'supports dimensions' do
        @metrics.expects(:track_gauge).once.with do |metric, _, dimensions|
          metric == :do_math && dimensions == {dimensions: {job: 'sales_order_job'}}
        end
        @metrics.time(:do_math, dimensions: {job: 'sales_order_job'}) { 2 + 2 == 4 }
      end
    end
  end
end
