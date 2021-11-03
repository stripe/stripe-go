# typed: strict
module Integrations
  module Errors
    class BaseIntegrationError < StandardError

    end

    class UnhandledEdgeCase < BaseIntegrationError; end
    class FeatureUsage < BaseIntegrationError; end
    class ImpossibleState < BaseIntegrationError; end
  end
end