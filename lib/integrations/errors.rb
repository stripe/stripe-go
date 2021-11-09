# typed: true
# frozen_string_literal: true

module Integrations
  module Errors
    class LockTimeout < StandardError; end
    class DyingWorkerError < StandardError; end

    class BaseIntegrationError < StandardError

    end

    class UnhandledEdgeCase < BaseIntegrationError; end
    class FeatureUsage < BaseIntegrationError; end
    class ImpossibleState < BaseIntegrationError; end
    class MissingRequiredFields < BaseIntegrationError; end

    # https://github.com/getsentry/sentry-ruby/issues/1612
    module ResqueHooks
      def after_perform_send_errors(*_args)
        Sentry.background_worker.drain_and_shutdown
      end

      def on_failure_send_errors(*_args)
        Sentry.background_worker.drain_and_shutdown
      end
    end
  end
end
