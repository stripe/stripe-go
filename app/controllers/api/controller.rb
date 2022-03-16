# frozen_string_literal: true
# typed: strict

module Api
  class Controller < ActionController::API
    include StripeForce::Constants
    include Integrations::ErrorContext
    include Integrations::Log
  end
end
