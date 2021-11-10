# frozen_string_literal: true
# typed: true
module Integrations
  module Metrics
    include Log

    def metrics
      Integrations::Metrics::Writer.instance
    end

    # adds class method when the method is included for convenience
    def self.included(klass)
      klass.class_eval do
        def self.log
          Integrations::Metrics::Writer.instance
        end
      end
    end

    QueuedMessage = Struct.new(:category, :data)

    class Writer
      include Singleton
      include Integrations::Log

      def initialize
        @queue = Queue.new

        @token = ENV.fetch('SIGNALFX_TOKEN')
        @url = "#{ENV.fetch('SIGNALFX_URL_BASE')}/datapoint"

        @max_queue_length = 500

        # TODO timeout_interval doesn't have any effect https://jira.corp.stripe.com/browse/DATAIO-213
        @timer = Concurrent::TimerTask.new(execution_interval: 1, timeout_interval: 10) do
          process
        end

        @timer.execute
      end

      def track_counter(metric, value: 1, dimensions: {})
        send_async(:counter, build_args(metric, value, dimensions))
      end

      def track_gauge(metric, value, dimensions: {})
        send_async(:gauge, build_args(metric, value, dimensions))
      end

      def send_metrics_synchronously
        process
      end

      def time(name, dimensions: {})
        start = Time.now
        yield
      ensure
        start = T.must(start)
        duration_in_seconds = Time.now - start

        # sfx assumes milliseconds, so we need to convert before sending
        track_gauge(name, (duration_in_seconds * 1_000), dimensions: dimensions)

        # return is required inside of `ensure` to return a value
        return duration_in_seconds
      end

      # For testing, poking with REPL
      attr_reader :timer, :queue
      attr_accessor :max_queue_length

      private def send_async(category, data)
        if @queue.length >= @max_queue_length
          log.warn 'dropping metric due to queue size'
          return
        end

        @queue << QueuedMessage.new(category, data)
      end

      private def build_args(metric, value, dimensions)
        {
          metric: "suitesync.#{metric}",
          dimensions: build_dimensions.merge(dimensions).transform_values(&:to_s),
          value: value,
          timestamp: (Time.now.utc.to_f * 1000).floor,
        }
      end

      private def build_dimensions
        {
          heroku_app: ENV.fetch('HEROKU_APP_NAME', ''),
        }
      end

      private def drain_queue
        rv = []
        begin
          until @queue.empty?
            rv << @queue.pop(true)
          end
        rescue ThreadError
          # - Somebody else is draining the queue as well
        end
        rv
      end

      private def collate_queue(items)
        rv = Hash.new {|h, k| h[k] = [] }

        items.each do |item|
          rv[item.category] << item.data
        end

        rv.each_value do |v|
          v.sort_by! {|x| x[:timestamp] }
        end

        rv
      end

      private def process
        data = collate_queue(drain_queue)
        return if data.empty?

        send_signal(data)
      end

      private def send_signal(payload)
        return if @token.nil? || @token.empty?

        headers = {
          "X-SF-TOKEN" => @token,
          content_type: :json,
        }

        begin
          RestClient::Request.execute(
            method: :post,
            url: @url,
            payload: payload.to_json,
            headers: headers,

            # surprisingly, the default open timeout on ruby is 60s
            # given that this is run at the end of a job and will block the next job
            # from being picked up, we should *not* block in the tail case where sfx is not responding
            # https://www.exceptionalcreatures.com/bestiary/Net/OpenTimeout.html
            open_timeout: 5,
            timeout: 15
          )
        rescue RestClient::BadRequest,
               RestClient::BadGateway,
               RestClient::ServiceUnavailable,
               RestClient::Exceptions::ReadTimeout,
               RestClient::Exceptions::OpenTimeout,
               RestClient::GatewayTimeout,
               OpenSSL::SSL::SSLError => e
          Sentry.capture_exception(e, level: 'warning', extra: {signalfx_error: true, payload: payload})
          log.error 'error sending metrics to signalfx', exception_message: e.message
        end
      end

    end

    # Resque calls these functions automatically based on the method name. They must be `extended` so they are defined as class methods
    # Removing the `after_perform` or `on_failure` prefix will stop automatically firing these functions
    # https://github.com/resque/resque/blob/master/docs/HOOKS.md
    module ResqueHooks
      def after_perform_send_metrics(*_args)
        Integrations::Metrics::Writer.instance.send_metrics_synchronously
      end

      def on_failure_send_metrics(*_args)
        Integrations::Metrics::Writer.instance.send_metrics_synchronously
      end
    end
  end
end
