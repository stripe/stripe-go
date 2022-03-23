# frozen_string_literal: true
# typed: strict

require_relative '../controller_helpers'

module Api
  class Controller < ActionController::API
    include StripeForce::Constants
    include Integrations::ErrorContext
    include Integrations::Log

    include ControllerHelpers
  end
end
