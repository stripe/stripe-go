# typed: true
# frozen_string_literal: true

module Critic
  module SalesforceFactory
    def create_salesforce_id
      SecureRandom.alphanumeric(18)
    end

    def create_mock_salesforce_order
      id = create_salesforce_id
      Restforce::SObject.new({"attributes" => {"type" => "Order", "url" => "/services/data/v52.0/sobjects/Order/#{id}", "Id" => id}})
    end


    def create_mock_salesforce_customer
      id = create_salesforce_id
      Restforce::SObject.new({"attributes" => {"type" => "Account", "url" => "/services/data/v52.0/sobjects/Account/#{id}", "Id" => id}})
    end

  end
end