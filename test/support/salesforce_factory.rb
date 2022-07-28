# typed: true
# frozen_string_literal: true

module Critic
  module SalesforceFactory
    extend T::Sig
    extend T::Helpers
    abstract!

    include Kernel
    include StripeForce::Constants
    include Minitest::Assertions

    sig { abstract.returns(T.untyped) }
    def sf; end

    def now_time_formatted_for_salesforce
      format_date_for_salesforce(now_time)
    end

    def now_time
      DateTime.now.utc.beginning_of_day
    end

    def format_date_for_salesforce(date)
      date.strftime("%Y-%m-%d")
    end

    def create_salesforce_id
      SecureRandom.alphanumeric(18)
    end

    def create_mock_salesforce_order
      id = create_salesforce_id
      Restforce::SObject.new({"attributes" => {"type" => SF_ORDER, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER}/#{id}", "Id" => id}})
    end

    def create_mock_salesforce_order_item
      id = create_salesforce_id
      Restforce::SObject.new({"attributes" => {"type" => SF_ORDER_ITEM, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER_ITEM}/#{id}", "Id" => id}})
    end

    def create_mock_salesforce_customer
      id = create_salesforce_id
      Restforce::SObject.new({"attributes" => {"type" => "Account", "url" => "/services/data/v52.0/sobjects/Account/#{id}", "Id" => id}})
    end

    def sf_randomized_name(sf_object_name)
      "REST #{sf_object_name} #{DateTime.now}"
    end

    def create_salesforce_account(additional_fields: {})
      account_id = sf.create!(SF_ACCOUNT, {
        Name: sf_randomized_name(SF_ACCOUNT),
      }.merge(additional_fields))
    end

    def create_salesforce_contact
      contact_id = sf.create!(SF_CONTACT, {
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
      }.merge(additional_fields))
    end

    def salesforce_recurring_metered_produce_with_price
      salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize,
        }
      )
    end

    def salesforce_recurring_product_with_price(price: nil, additional_product_fields: {})
      # blanking out the subscription type ensures it is a one-time product
      product_id = create_salesforce_product(additional_fields: {
        # anything non-nil indicates subscription/recurring pricing
        CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::RENEWABLE,

        # default term of one month
        CPQ_QUOTE_SUBSCRIPTION_TERM => 1,

        # one month
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
      }.merge(additional_product_fields))

      pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price)
      [product_id, pricebook_entry_id]
    end

    def salesforce_standalone_product_with_price(price: 100_00)
      # blanking out the subscription type ensures it is a one-time product
      product_id = create_salesforce_product(additional_fields: {
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::ONE_TIME,
      })

      pricebook_entry_id = create_salesforce_price(sf_product_id: product_id)

      [product_id, pricebook_entry_id]
    end

    def create_subscription_order(sf_product_id: nil, sf_account_id: nil)
      create_salesforce_order(
        sf_product_id: sf_product_id,
        sf_account_id: sf_account_id,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          # one year / 12 months
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        }
      )
    end

    def add_product_to_cpq_quote(quote_id, sf_product_id:, sf_pricebook_id: nil)
      sf_pricebook_id ||= default_pricebook_id

      # this error indicates that the SF account disconnected from the steelbrick REST service
      # reauth via: https://brick-rest.steelbrick.com/oauth/auth
      # Restforce::ErrorCode::ApexError: APEX_ERROR: SBQQ.RestClient.RefreshTokenNilException: Invalid nil argument: OAuth Refresh Token

      # https://developer.salesforce.com/docs/atlas.en-us.cpq_api_dev.meta/cpq_api_dev/cpq_api_read_quote.htm
      # get CPQ version of the quote
      quote_data = JSON.parse(sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body)

      cpq_product_representation = JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ProductAPI.ProductLoader&uid=#{sf_product_id}", {
        context: {
          # currencyCode:
          pricebookId: sf_pricebook_id,
        }.to_json,
      }).body)

      # TODO how can I specify quantity?
      # https://gist.github.com/paustint/bd18bd281134a180e014829b49ed043a
      quote_with_product = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteProductAdder', {
        context: {
          "quote": quote_data,
          "products": [
            cpq_product_representation,
          ],
          "groupKey": 0,
          "ignoreCalculate": false,
        }.to_json,
      }).body)

      quote_with_product
    end

    def calculate_and_save_cpq_quote(quote_data)
      # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_calculate_final.htm
      calculated_quote = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteCalculator', {
        # "context": quote_with_product.to_json
        # "context": saved_quote.to_json
        # "context": sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body
        "context": {"quote" => quote_data}.to_json,
      }).body)

      # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_save_final.htm
      saved_quote = JSON.parse(sf.post('/services/apexrest/SBQQ/ServiceRouter', {
        "saver": "SBQQ.QuoteAPI.QuoteSaver",
        "model": calculated_quote.to_json,
      }).body)

      refute(saved_quote["calculationRequired"])

      saved_quote.dig("record", "Id")
    end

    # returns full SF order object
    sig { params(sf_quote_id: String).returns(T.untyped) }
    def create_order_from_cpq_quote(sf_quote_id)
      # it looks like there is additional field validation triggered here when `ordered` is set to true
      sf.update!(CPQ_QUOTE, {
        SF_ID => sf_quote_id,
        CPQ_QUOTE_ORDERED => true,
      })

      sf_quote = sf.find(CPQ_QUOTE, sf_quote_id)

      # TODO this is silly, I should do a simple SOQL query and grab the result
      # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
      # TODO note that looking in the UI is the easiest way to get these magic relational values
      related_orders = sf.get("/services/data/v52.0/sobjects/#{CPQ_QUOTE}/#{sf_quote_id}/SBQQ__Orders__r")
      sf_order = related_orders.body.first

      sf.update!(SF_ORDER,
        SF_ID => sf_order.Id,
        'Status' => 'Activated'
      )

      sf_order.refresh
      sf_order
    end

    def create_salesforce_quote(sf_pricebook_id: nil, sf_account_id:, additional_quote_fields: {})
      sf_pricebook_id ||= default_pricebook_id
      opportunity_id = create_salesforce_opportunity(sf_account_id: sf_account_id)
      contact_id = create_salesforce_contact

      # you can create a quote without *any* fields, which seems completely silly
      quote_id = sf.create!(CPQ_QUOTE, {
        CPQ_QUOTE_PRIMARY => true,
        CPQ_QUOTE_OPPORTUNITY => opportunity_id,
        CPQ_QUOTE_PRIMARY_CONTACT => contact_id,
        CPQ_QUOTE_PRICEBOOK => sf_pricebook_id,
      }.merge(additional_quote_fields))
    end

    # https://github.com/sseixas/CPQ-JS
    def create_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, additional_quote_fields: {})
      if !sf_product_id
        sf_product_id, _ = salesforce_recurring_product_with_price
      end

      sf_pricebook_id ||= default_pricebook_id
      sf_account_id ||= create_salesforce_account

      quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: additional_quote_fields)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

      calculate_and_save_cpq_quote(quote_with_product)

      create_order_from_cpq_quote(quote_id)

      # contract_id = salesforce_client.create!('Contract', accountId: account_id)
      # order_id = salesforce_client.create!('Order', {Status: "Draft", EffectiveDate: "2021-09-21", AccountId: account_id, ContractId: contract_id})
    end
  end
end
