# typed: false
# frozen_string_literal: true

module Critic
  module SalesforceFactory
    include StripeForce::Constants

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

    def sf_randomized_name(sf_object_name)
      "REST #{sf_object_name} #{DateTime.now}"
    end

    def create_salesforce_price(sf_product_id: nil, price: nil)
      price ||= 120
      sf_product_id ||= create_salesforce_product

      sf_pricebook_entry_id = sf.create!('PricebookEntry',
        # TODO the default pricebook will need to be dynamic
        "Pricebook2Id" => "01s5e00000BAoBWAA1",

        "Product2Id" => sf_product_id,
        "IsActive" => true,
        "UnitPrice" => price
      )
    end

    def create_salesforce_product(additional_fields: {})
      sf.create!(SF_PRODUCT, {
        "Name" => sf_randomized_name(SF_PRODUCT),
        "Description" => "A great description",

        # this field is an enum/picklist in SF
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => 'Renewable'
      }.merge(additional_fields))
    end

  end
end