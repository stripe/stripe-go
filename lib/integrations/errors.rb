# typed: true
# frozen_string_literal: true

module Integrations
  module Errors
    class LockTimeout < StandardError; end
    class DyingWorkerError < StandardError; end

    class BaseIntegrationError < StandardError
      attr_reader :stripe_resource
      attr_reader :salesforce_object
      attr_reader :metadata

      def initialize(message, stripe_resource: nil, salesforce_object: nil, metadata: nil)
        @stripe_resource = stripe_resource
        @salesforce_object = salesforce_object
        @metadata = metadata

        sf_message = if salesforce_object && salesforce_object.class == Class
          "salesforce_object_type=#{salesforce_object}"
        elsif salesforce_object
          "salesforce_object_type=#{salesforce_object.sobject_type} salesforce_object_id=#{salesforce_object.Id} "
        else
          ""
        end

        stripe_message = if stripe_resource
          "stripe_resource_id=#{stripe_resource.id}"
        else
          ""
        end

        super([message, stripe_message, sf_message].join(' ').strip)
      end

      # this is a magic method which is pulled into the event sent to sentry
      # https://github.com/getsentry/sentry-ruby/blob/0f66522fec93bd03d5f3a3c9e9c0a1d759013b64/sentry-ruby/lib/sentry/error_event.rb#L32
      def sentry_context
        @metadata || {}
      end
    end

    class UnhandledEdgeCase < BaseIntegrationError; end
    class FeatureUsage < BaseIntegrationError; end
    class ImpossibleState < BaseIntegrationError; end
    class UserError < BaseIntegrationError; end
    class TranslatorError < BaseIntegrationError; end
    class MissingRequiredFields < BaseIntegrationError
      attr_reader :missing_salesforce_fields

      def initialize(missing_salesforce_fields:, stripe_resource: nil, salesforce_object: nil, metadata: {})
        @missing_salesforce_fields = missing_salesforce_fields

        # include in metadata in order to propogate to sentry
        metadata = metadata.merge(missing_salesforce_fields: missing_salesforce_fields)

        # TODO should remove missing fields from prod to avoid tons of different-looking errors
        super("missing required fields #{missing_salesforce_fields.inspect}", stripe_resource: stripe_resource, salesforce_object: salesforce_object, metadata: metadata)
      end
    end
  end
end
