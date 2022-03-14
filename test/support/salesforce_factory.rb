# typed: true
# frozen_string_literal: true

module Critic
  module SalesforceFactory
    extend T::Sig
    extend T::Helpers
    abstract!

    include Kernel
    include StripeForce::Constants

    sig { abstract.returns(T.untyped) }
    def sf; end

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

    def create_salesforce_account
      account_id = sf.create!(SF_ACCOUNT, Name: sf_randomized_name(SF_ACCOUNT))
    end

    def create_salesforce_contact
      sf.create!('Contact', {
        LastName: 'Bianco',
        Email: "#{DateTime.now.to_i}@example.com",
      })
    end

    def create_salesforce_opportunity(sf_account_id:, additional_fields: {})
      sf.create!(SF_OPPORTUNITY, {
        Name: sf_randomized_name(SF_OPPORTUNITY),
        "CloseDate": DateTime.now.iso8601,
        StageName: "Closed/Won",
        "AccountId": sf_account_id,
      }.merge(additional_fields))
    end

    def create_salesforce_price(sf_product_id: nil, price: nil)
      price ||= 120_00
      sf_product_id ||= create_salesforce_product

      sf_pricebook_entry_id = sf.create!(SF_PRICEBOOK_ENTRY,
        "Pricebook2Id" => default_pricebook_id,
        "Product2Id" => sf_product_id,

        "IsActive" => true,
        "UnitPrice" => price / 100.0,
        'UseStandardPrice' => false,
      )
    end

    def default_pricebook_id
      # https://help.salesforce.com/s/articleView?id=000326219&type=1
      standard_pricebook = sf.query("Select Id from #{SF_PRICEBOOK} where IsStandard = true").first

      if !standard_pricebook
        raise "could not find standard pricebook"
      end

      standard_pricebook.Id
    end

    def create_salesforce_product(additional_fields: {})
      sf.create!(SF_PRODUCT, {
        "Name" => sf_randomized_name(SF_PRODUCT),
        'IsActive' => true,
        "Description" => "A great description",
        'ProductCode' => "Prod#{Time.now.to_i}",

        # this field is an enum/picklist in SF
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQ_PRODUCT_SUBSCRIPTION_TYPE_RENEWABLE,
      }.merge(additional_fields))
    end

  end
end
