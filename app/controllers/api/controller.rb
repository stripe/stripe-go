# frozen_string_literal: true
# typed: strict

module Api
  class Controller < ActionController::API
    include Integrations::ErrorContext
    include Integrations::Log

    SALESFORCE_ACCOUNT_ID_HEADER = 'Salesforce-Account-Id'
  end
end
