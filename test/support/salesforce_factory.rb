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
    include Integrations::Log

    TEST_DEFAULT_CONTRACT_TERM = 12
    TEST_DEFAULT_PRICE = 120_00
    TEST_DEFAULT_STANDALONE_PRICE = 100_00

    sig { returns(String) }
    def now_time_formatted_for_salesforce
      format_date_for_salesforce(now_time)
    end

    sig { returns(String) }
    def now_time_in_future_formatted_for_salesforce
      format_date_for_salesforce(now_time_in_future)
    end

    def now_time
      DateTime.now.utc.beginning_of_day
    end

    def now_time_in_future
      DateTime.new(2033, 7, 1)
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
        SF_LAST_MODIFIED_DATE => "2022-09-20T17:53:31.000+0000",
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

    sig { params(additional_fields: Hash).returns(String) }
    def create_salesforce_account(additional_fields: {})
      account_id = sf.create!(SF_ACCOUNT, {
        Name: sf_randomized_name(SF_ACCOUNT),
      }.merge(additional_fields))
    end

    def create_salesforce_contact(contact_email: sf_randomized_id, static_id: true)
      contact_id = sf.create!(SF_CONTACT, {
        LastName: 'Bianco',
        Email: static_id ? create_static_email(email: contact_email) : create_random_email,
      })
    end

    def create_salesforce_opportunity(sf_account_id:, currency_iso_code: nil, additional_fields: {})
      currency_iso_code ||= @user.connector_settings['default_currency']

      if @user.is_multicurrency_org?
        additional_fields = additional_fields.merge({
          StripeForce::Constants::SF_CURRENCY_ISO_CODE => currency_iso_code,
        })
      end

      sf.create!(SF_OPPORTUNITY, {
        Name: sf_randomized_name(SF_OPPORTUNITY),
        "CloseDate": now_time_formatted_for_salesforce,
        StageName: "Closed/Won",
        "AccountId": sf_account_id,
      }.merge(additional_fields))
    end

    def create_salesforce_price(sf_product_id: nil, price: nil, currency_iso_code: nil, additional_fields: {})
      price ||= TEST_DEFAULT_PRICE
      sf_product_id ||= create_salesforce_product
      currency_iso_code ||= @user.connector_settings['default_currency']

      if @user.is_multicurrency_org?
        additional_fields = additional_fields.merge({
          StripeForce::Constants::SF_CURRENCY_ISO_CODE => currency_iso_code,
        })
      end

      sf_pricebook_entry_id = sf.create!(SF_PRICEBOOK_ENTRY, {
        "Pricebook2Id" => default_pricebook_id,
        "Product2Id" => sf_product_id,
        "IsActive" => true,
        "UnitPrice" => price / 100.0,
        'UseStandardPrice' => false,
      }.merge(additional_fields))
    end

    def default_pricebook_id
      # https://help.salesforce.com/s/articleView?id=000326219&type=1
      standard_pricebook = sf.query("Select #{SF_ID} from #{SF_PRICEBOOK} where IsStandard = true").first

      if !standard_pricebook
        raise "could not find standard pricebook"
      end

      standard_pricebook.Id
    end

    def create_salesforce_product(static_id: true, additional_fields: {})
      sf.create!(SF_PRODUCT, {
        "Name" => sf_randomized_name(SF_PRODUCT),
        'IsActive' => true,
        "Description" => "A great description",
        'ProductCode' => static_id ? sf_static_id : sf_randomized_id,
      }.merge(additional_fields))
    end

    def create_salesforce_cpq_segmented_product(additional_product_fields: {}, additional_price_dimension_fields: {}, price: nil)
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price, additional_product_fields: additional_product_fields)

      if additional_product_fields[CPQ_PRODUCT_BILLING_TYPE] == CPQProductBillingTypeOptions::ARREARS.serialize
        sf_product_id, _sf_pricebook_id = salesforce_recurring_metered_produce_with_price(price_in_cents: price)
      end

      # create a price dimension and link the product to it
      sf.create!(CPQ_PRICE_DIMENSION, {
        "Name" => "Yearly Ramp",
        CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize,
        CPQ_QUOTE_LINE_PRODUCT => sf_product_id,
      }.merge(additional_price_dimension_fields))

      sf_product_id
    end

    def create_salesforce_stripe_coupon_quote_association(sf_quote_id:, sf_stripe_coupon_id:)
     sf_stripe_coupon_id ||= create_salesforce_stripe_coupon

     sf.create!(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON_ASSOCIATION), {
       "Quote__c" => sf_quote_id,
       QUOTE_SF_STRIPE_COUPON => sf_stripe_coupon_id,
     }.transform_keys(&method(:prefixed_stripe_field)))
    end

    def create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id:, sf_stripe_coupon_id:)
      sf_stripe_coupon_id ||= create_salesforce_stripe_coupon

      # Our association objects unfortunately have currency fields in multicurrency orgs, so we must pass
      #   the coupon's currency to the association here.
      # This is error checked in the UI with this PR: https://github.com/stripe/stripe-salesforce/pull/1006
      currency_field = {}
      if @user.is_multicurrency_org?
        sf_stripe_coupon = sf_get(sf_stripe_coupon_id)
        currency_field = {
          StripeForce::Constants::SF_CURRENCY_ISO_CODE => sf_stripe_coupon[StripeForce::Constants::SF_CURRENCY_ISO_CODE],
        }
      end

      sf.create!(prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION), {
        "Quote_Line__c" => sf_quote_line_id,
        QUOTE_SF_STRIPE_COUPON => sf_stripe_coupon_id,
      }.merge(currency_field).transform_keys(&method(:prefixed_stripe_field)))
    end

    def create_salesforce_stripe_coupon(additional_fields: {})
      # we want to ensure that either amount_off or percent_off is set and not both
      amount_off = nil
      if additional_fields[SalesforceStripeCouponFields::PERCENT_OFF].nil?
          amount_off = TEST_DEFAULT_PRICE / 2
      end

      sf.create!(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), {
        SalesforceStripeCouponFields::NAME => sf_randomized_name(QUOTE_SF_STRIPE_COUPON),
        SalesforceStripeCouponFields::AMOUNT_OFF => amount_off,
        SalesforceStripeCouponFields::DURATION => 'once',
      }.merge(additional_fields).transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field)))
    end

    def get_sf_order_coupon_from_quote_coupon_id(quote_coupon_id)
      order_coupons = sf.query("SELECT #{SF_ID} FROM #{prefixed_stripe_field(ORDER_SF_STRIPE_COUPON)} WHERE #{prefixed_stripe_field('Quote_Stripe_Coupon_Id__c')} = '#{quote_coupon_id}'")

      # sanity check the order coupon info is the same as the quote coupon
      sf.find(prefixed_stripe_field(ORDER_SF_STRIPE_COUPON), order_coupons.first.Id)
    end

    def get_salesforce_stripe_coupons_associated_to_quote_line(quote_line_id:)
      quote_line_associations = sf.query("Select #{SF_ID} from #{prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION)} where #{prefixed_stripe_field('Quote_Line__c')} = '#{quote_line_id}'")

      if !quote_line_associations
        raise "could not find any stripe coupon quote line associations related to this quote line"
      end

      # there could be multiple coupons associated with a quote line
      coupons = quote_line_associations.map do |quote_line_association|
          association = sf.find(prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION), quote_line_association.Id)
          coupon_id = association[prefixed_stripe_field('Quote_Stripe_Coupon__c')]

          # return the coupon object
          sf.find(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), coupon_id)
      end
      coupons
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
    def salesforce_recurring_product_with_price(price: nil, currency_iso_code: nil, static_id: true, additional_product_fields: {})
      # I don't fully understand how the subscription term on the price iteracts with the billing frequency,
      # but if the term is set to a value which is different than the billing frequency it seems to use the
      # subscription term value. i.e. a yearly billed product
      subscription_term = if additional_product_fields.key?(CPQ_QUOTE_BILLING_FREQUENCY) && additional_product_fields[CPQ_QUOTE_BILLING_FREQUENCY] != CPQBillingFrequencyOptions::MONTHLY.serialize
        # If this field is left blank or not set, Salesforce CPQ typically assumes the subscription term to be a month or treats it as a month-to-month contract
        nil
      else
        1
      end

      # blanking out the subscription type ensures it is a one-time product
      product_id = create_salesforce_product(static_id: static_id, additional_fields: {
        # anything non-nil indicates subscription/recurring pricing
        CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::RENEWABLE,
        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
      }.merge(additional_product_fields))

      pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price, currency_iso_code: currency_iso_code)
      [product_id, pricebook_entry_id]
    end

    def salesforce_standalone_product_with_price(price: TEST_DEFAULT_STANDALONE_PRICE)
      # blanking out the subscription type ensures it is a one-time product
      product_id = create_salesforce_product(additional_fields: {
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::ONE_TIME,
      })

      pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price)

      [product_id, pricebook_entry_id]
    end

    def salesforce_evergreen_product_with_price(price: nil, currency_iso_code: nil, additional_product_fields: {})
      # evergreen products all have subscription term set to 1
      product_id = create_salesforce_product(additional_fields: {
        # anything non-nil indicates subscription/recurring pricing
        CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::EVERGREEN,

        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }.merge(additional_product_fields))

      pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price, currency_iso_code: currency_iso_code)
      [product_id, pricebook_entry_id]
    end

    # returns the full object, not the ID
    def create_subscription_order(sf_product_id: nil, sf_account_id: nil, currency_iso_code: nil, contact_email: sf_randomized_id, additional_fields: {})
      create_salesforce_order(
        sf_product_id: sf_product_id,
        sf_account_id: sf_account_id,
        currency_iso_code: currency_iso_code,
        contact_email: contact_email,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          # one year / 12 months
          CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
        }.merge(additional_fields)
      )
    end

    def add_product_to_cpq_quote(quote_id, sf_product_id:, sf_pricebook_id: nil, product_quantity: nil)
      sf_pricebook_id ||= default_pricebook_id

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

      if product_quantity.present?
        quote_with_product["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = product_quantity
      end

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
    sig { params(sf_quote_id: String, activate_order: T::Boolean).returns(T.untyped) }
    def create_order_from_cpq_quote(sf_quote_id, activate_order: true)
      # it looks like there is additional field validation triggered here when `ordered` is set to true
      # If you recieve a "FIELD_CUSTOM_VALIDATION_EXCEPTION: Required fields are missing: [PricebookEntryId]" here,
      #   this means that SF could not find a pricebook entry for the product you are passing in. This may be because of a
      #   currency mismatch on the opportunity and the pricebook entry you created for the product.
      sf.update!(CPQ_QUOTE, {
        SF_ID => sf_quote_id,
        CPQ_QUOTE_ORDERED => true,
      })

      # for debugging
      # @user.sf_client.find(CPQ_QUOTE, sf_quote_id)

      # TODO this is silly, I should do a simple SOQL query and grab the result
      # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
      # TODO note that looking in the UI is the easiest way to get these magic relational values
      related_orders = sf.get("/services/data/v52.0/sobjects/#{CPQ_QUOTE}/#{sf_quote_id}/SBQQ__Orders__r")
      sf_order = related_orders.body.first

      if activate_order
        sf.update!(SF_ORDER,
          SF_ID => sf_order.Id,
          'Status' => 'Activated'
        )
      end

      sf_order.refresh
      sf_order
    end

    def create_salesforce_quote(sf_pricebook_id: nil, sf_account_id:, currency_iso_code: nil, contact_email: nil, additional_quote_fields: {})
      sf_pricebook_id ||= default_pricebook_id
      opportunity_id = create_salesforce_opportunity(sf_account_id: sf_account_id, currency_iso_code: currency_iso_code)
      contact_id = create_salesforce_contact(contact_email: contact_email)

      # you can create a quote without *any* fields, which seems completely silly
      quote_id = sf.create!(CPQ_QUOTE, {
        CPQ_QUOTE_PRIMARY => true,
        CPQ_QUOTE_OPPORTUNITY => opportunity_id,
        CPQ_QUOTE_PRIMARY_CONTACT => contact_id,
        CPQ_QUOTE_PRICEBOOK => sf_pricebook_id,
      }.merge(additional_quote_fields))
      quote_id
    end

    # https://github.com/sseixas/CPQ-JS
    def create_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, currency_iso_code: nil, contact_email: sf_randomized_id, additional_quote_fields: {})
      if !sf_product_id
        sf_product_id, _ = salesforce_recurring_product_with_price(currency_iso_code: currency_iso_code)
      end

      sf_pricebook_id ||= default_pricebook_id
      sf_account_id ||= create_salesforce_account

      quote_id = create_salesforce_quote(sf_account_id: sf_account_id, currency_iso_code: currency_iso_code, contact_email: contact_email, additional_quote_fields: additional_quote_fields)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

      calculate_and_save_cpq_quote(quote_with_product)

      create_order_from_cpq_quote(quote_id)

      # contract_id = salesforce_client.create!('Contract', accountId: account_id)
      # order_id = salesforce_client.create!('Order', {Status: "Draft", EffectiveDate: "2021-09-21", AccountId: account_id, ContractId: contract_id})
    end

    # for creating evergreen orders specifically. evergreen subscriptions do not have an end date.
    def create_evergreen_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, currency_iso_code: nil, contact_email: sf_randomized_id, additional_quote_fields: {})
      if !sf_product_id
        sf_product_id, _ = salesforce_evergreen_product_with_price(currency_iso_code: currency_iso_code)
      end

      # sf_product_id should point to an evergreen product so will not be nil inside create sf order
      create_salesforce_order(sf_product_id: sf_product_id, sf_account_id: sf_account_id, sf_pricebook_id: sf_pricebook_id, currency_iso_code: currency_iso_code, contact_email: contact_email, additional_quote_fields: additional_quote_fields)
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
        'BillingTerm' => TEST_DEFAULT_CONTRACT_TERM,
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

    def stub_salesforce_query_result(number_of_results: 1)

      # RestForce abstracts away paging, so we can safely mock this as a simple list
      # https://github.com/restforce/restforce/blob/main/lib/restforce/collection.rb#L15
      result_list = []

      number_of_results.times do |_t|
        result_list << Restforce::SObject.new({"Id" => create_salesforce_id})
      end

      Restforce::Data::Client.any_instance.expects(:query).returns(result_list)

      result_list
    end
  end
end
