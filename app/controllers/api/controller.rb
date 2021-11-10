# frozen_string_literal: true
# typed: strict

module Api
  class Controller < ActionController::API
    include Integrations::ErrorContext
    include Integrations::Log
  end
end
