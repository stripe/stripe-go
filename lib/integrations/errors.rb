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

        # TODO more error information when stripe & salesforce objects are used
        super(message)
      end
    end

    class UnhandledEdgeCase < BaseIntegrationError; end
    class FeatureUsage < BaseIntegrationError; end
    class ImpossibleState < BaseIntegrationError; end
    class UserError < BaseIntegrationError; end
    class MissingRequiredFields < BaseIntegrationError
      attr_reader :missing_salesforce_fields

      def initialize(missing_salesforce_fields:, stripe_resource: nil, salesforce_object: nil, metadata: nil)
        @missing_salesforce_fields = missing_salesforce_fields
        super("missing required fields", stripe_resource: stripe_resource, salesforce_object: salesforce_object, metadata: metadata)
      end
    end
  end
end
