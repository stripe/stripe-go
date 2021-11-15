# frozen_string_literal: true
# typed: strict

module Api
  class Controller < ActionController::API
    include Integrations::ErrorContext
    include Integrations::Log

    SALESFORCE_ACCOUNT_ID_HEADER = 'Salesforce-Account-Id'
    SALESFORCE_KEY_HEADER = 'Salesforce-Key'
  end
end
