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
    include Critic::CommonHelpers

    sig { returns(String) }
    def now_time_formatted_for_salesforce
      format_date_for_salesforce(now_time)
    end

    def now_time
      DateTime.now.utc.beginning_of_day
    end

    def format_date_for_salesforce(date)
      date.strftime("%Y-%m-%d")
    end

    sig { params(prefix: T.nilable(String)).returns(String) }
    def create_salesforce_id(prefix: nil)
      (prefix || "") + SecureRandom.alphanumeric(prefix ? 15 : 18)
    end

    def create_mock_salesforce_order
      id = create_salesforce_id(prefix: "802")
      Restforce::SObject.new({
        "attributes" => {"type" => SF_ORDER, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER}/#{id}"},
        "Id" => id,
        "IsDeleted" => false,
      })
    end

    def create_mock_salesforce_order_item(additional_attributes: {})
      id = create_salesforce_id(prefix: "0Mh")
      Restforce::SObject.new({
        "Id" => id,
        "IsDeleted" => false,
        "attributes" => {"type" => SF_ORDER_ITEM, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER_ITEM}/#{id}"},
      }.merge(additional_attributes))
    end

    def create_mock_salesforce_customer
      # TODO should move all prefixes to an enum
      id = create_salesforce_id(prefix: "003")
      Restforce::SObject.new({
        "attributes" => {"type" => "Account", "url" => "/services/data/v52.0/sobjects/Account/#{id}"},
        "IsDeleted" => false,
        "Id" => id,
      })
    end

    def create_salesforce_account(additional_fields: {})
      account_id = sf.create!(SF_ACCOUNT, {
        Name: sf_randomized_name(SF_ACCOUNT),
      }.merge(additional_fields))
    end

    def create_salesforce_contact
      contact_id = sf.create!(SF_CONTACT, {
        LastName: 'Bianco',
        Email: create_random_email,
      })
    end

    def create_salesforce_opportunity(sf_account_id:, additional_fields: {})
      sf.create!(SF_OPPORTUNITY, {
        Name: sf_randomized_name(SF_OPPORTUNITY),
        "CloseDate": now_time_formatted_for_salesforce,
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
        'ProductCode' => sf_randomized_id,
      }.merge(additional_fields))
    end

    def salesforce_recurring_metered_produce_with_price(price_in_cents: nil)
      salesforce_recurring_product_with_price(
        price: price_in_cents,
        additional_product_fields: {
          CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize,
        }
      )
    end

    # TODO `price` should be `price_in_cents`
    def salesforce_recurring_product_with_price(price: nil, additional_product_fields: {})
      # I don't fully understand how the subscription term on the price iteracts with the billing frequency,
      # but if the term is set to a value which is different than the billing frequency it seems to use the
      # subscription term value. i.e. a yearly billed product
      subscription_term = if additional_product_fields.key?(CPQ_QUOTE_BILLING_FREQUENCY) && additional_product_fields[CPQ_QUOTE_BILLING_FREQUENCY] != CPQBillingFrequencyOptions::MONTHLY.serialize
        nil
      else
        1
      end

      # blanking out the subscription type ensures it is a one-time product
      product_id = create_salesforce_product(additional_fields: {
        # anything non-nil indicates subscription/recurring pricing
        CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::RENEWABLE,

        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,

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
      log.info 'calculate and save quote'

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

    def create_recurring_per_unit_tiered_price
      # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
      product_id, pricebook_entry_id = salesforce_recurring_product_with_price

      consumption_schedule_id = create_salesforce_consumption_schedule

      # first tier is priced at per-seat pricing, the rest are not
      # this test also tested >2 tiered pricing
      consumption_rate_id_1 = create_salesforce_consumption_rate(consumption_schedule_id)
      consumption_rate_id_2 = create_salesforce_consumption_rate(consumption_schedule_id, {
        "LowerBound" => 2,
        "UpperBound" => 10,
        "Price" => 20,
      })
      consumption_rate_id_3 = create_salesforce_consumption_rate(consumption_schedule_id, {
        "LowerBound" => 10,
        "UpperBound" => nil,
        "Price" => 30,
      })

      activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

      [product_id, pricebook_entry_id]
    end

    def create_salesforce_consumption_schedule(additional_fields={})
      sf.create!(SF_CONSUMPTION_SCHEDULE, {
        'Name' => sf_randomized_name(SF_CONSUMPTION_SCHEDULE),
        'Type' => 'Range',
        'RatingMethod' => 'Tier',
        'SBQQ__Category__c' => 'Rates',
        'BillingTerm' => 12,
        'BillingTermUnit' => "Month",
        # cannot create the consumption schedule as activated
        # 'IsActive' => true,
      }.merge(additional_fields))
    end

    def create_salesforce_consumption_rate(consumption_schedule_id, additional_fields={})
      @consumption_schedule_processing_order ||= Hash.new(0)
      @consumption_schedule_processing_order[consumption_schedule_id] += 1

      sf.create!(SF_CONSUMPTION_RATE, {
        "ConsumptionScheduleId" => consumption_schedule_id,
        "ProcessingOrder" => @consumption_schedule_processing_order[consumption_schedule_id],
        "PricingMethod" => "PerUnit",
        "LowerBound" => 1,
        "UpperBound" => 2,
        "Price" => 10,
      }.merge(additional_fields))
    end

    def activate_and_link_consumption_schedule(consumption_schedule_id, product_id)
      sf.update!(SF_CONSUMPTION_SCHEDULE, {
        SF_ID => consumption_schedule_id,
        'IsActive' => true,
      })

      # joining record, must be an activated consumption schedule
      sf.create!('ProductConsumptionSchedule', {
        'ConsumptionScheduleId' => consumption_schedule_id,
        'ProductId' => product_id,
      })
    end
  end
end
