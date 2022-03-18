# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  def salesforce_recurring_product_with_price
    external_id = "recurring-product"

    # TODO we should add a field to track the external ID
    # sf.find(SF_PRODUCT, external_id, 'THE_FIELD')

    # blanking out the subscription type ensures it is a one-time product
    product_id = create_salesforce_product
    pricebook_entry = create_salesforce_price(sf_product_id: product_id)
    product_id
  end

  def salesforce_standalone_product_with_price_id(price: 100_00)
    external_id = "standalone-product-id-#{price}"

    # TODO we should add a field to track the external ID
    # product_id = sf.find(SF_PRODUCT, external_id, 'THE_FIELD__C')
    # return product_id if product_id

    # blanking out the subscription type ensures it is a one-time product
    product_id = create_salesforce_product(additional_fields: {
      CPQ_PRODUCT_SUBSCRIPTION_TYPE => "",
    })

    pricebook_entry = create_salesforce_price(sf_product_id: product_id)

    product_id
  end

  # https://github.com/sseixas/CPQ-JS
  def create_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, additional_order_fields: {})
    sf_pricebook_id ||= default_pricebook_id
    sf_account_id ||= create_salesforce_account

    opportunity_id = create_salesforce_opportunity(sf_account_id: sf_account_id)
    contact_id = create_salesforce_contact

    # you can create a quote without *any* fields, which seems completely silly
    quote_id = sf.create!(CPQ_QUOTE, {
      CPQ_QUOTE_PRIMARY => true,

      "SBQQ__Opportunity2__c": opportunity_id,
      'SBQQ__PrimaryContact__c' => contact_id,
      'SBQQ__PricebookId__c' => sf_pricebook_id,
    }.merge(additional_order_fields))

    # this error indicates that the SF account disconnected from the steelbrick REST service
    # reauth via: https://brick-rest.steelbrick.com/oauth/auth
    # Restforce::ErrorCode::ApexError: APEX_ERROR: SBQQ.RestClient.RefreshTokenNilException: Invalid nil argument: OAuth Refresh Token

    # https://developer.salesforce.com/docs/atlas.en-us.cpq_api_dev.meta/cpq_api_dev/cpq_api_read_quote.htm
    # get CPQ version of the quote
    cpq_quote_representation = JSON.parse(sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body)

    cpq_product_representation = JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ProductAPI.ProductLoader&uid=#{sf_product_id}", {
      context: {
        # productId: product_id,
        pricebookId: sf_pricebook_id,
        # currencyCode:
      }.to_json,
    }).body)

    # TODO how can I specify quantity?
    # https://gist.github.com/paustint/bd18bd281134a180e014829b49ed043a
    quote_with_product = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteProductAdder', {
      context: {
        "quote": cpq_quote_representation,
        "products": [
          cpq_product_representation,
        ],
        "groupKey": 0,
        "ignoreCalculate": true,
        # quote: {
        #   record: {
        #     Id: quote_id,
        #     attributes: {
        #       type: CPQ_QUOTE,
        #     }
        #   }
        # },
        # products: [
        #   {
        #     record: {
        #       Id: product_id
        #     }
        #   }
        # ],
        # ignoreCalculate: true
      }.to_json,
    }).body)

    # calculated_quote = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteCalculator', {
    #   "context": { "quote": updated_quote["record"] }.to_json
    # }).body)


    # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_calculate_final.htm
    calculated_quote = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteCalculator', {
      # "context": quote_with_product.to_json
      # "context": saved_quote.to_json
      # "context": sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body
      "context": {"quote" => quote_with_product}.to_json,
    }).body)

    # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_save_final.htm
    saved_quote = JSON.parse(sf.post('/services/apexrest/SBQQ/ServiceRouter', {
      "saver": "SBQQ.QuoteAPI.QuoteSaver",
      "model": calculated_quote.to_json,
    }).body)

    # sf.create!(CPQ_QUOTE_LINE, {
    #   CPQ_QUOTE => quote_id,
    #   CPQ_QUOTE_LINE_PRODUCT => product_id,
    #   CPQ_QUOTE_LINE_PRICEBOOK_ENTRY => pricebook_entry_id
    # })

    # give CPQ some time to calculate...
    # sleep(5)

    # it looks like there is additional field validation triggered here when `ordered` is set to true
    sf.update!(CPQ_QUOTE, 'Id' => quote_id, CPQ_QUOTE_ORDERED => true)

    sf_quote = sf.find(CPQ_QUOTE, quote_id)

    # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
    # TODO note that looking in the UI is the easiest way to get these magic relational values
    related_orders = sf.get("/services/data/v52.0/sobjects/#{CPQ_QUOTE}/#{quote_id}/SBQQ__Orders__r")
    sf_order = related_orders.body.first

    sf.update!('Order', 'Id' => sf_order.Id, 'Status' => 'Activated')

    # TODO need refresh here
    sf.find('Order', sf_order.Id)

    # contract_id = salesforce_client.create!('Contract', accountId: account_id)
    # order_id = salesforce_client.create!('Order', {Status: "Draft", EffectiveDate: "2021-09-21", AccountId: account_id, ContractId: contract_id})
  end

  it 'integrates a subscription order' do
    price = 240_00

    sf_product_id = create_salesforce_product
    sf_pricebook_entry_id = create_salesforce_price(sf_product_id: sf_product_id, price: price)

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_order_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => "2021-10-01",
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    # TODO add `refresh` to salesforce library
    sf_order = sf.find(SF_ORDER, sf_order.Id)
    sf_product = sf.find(SF_PRODUCT, sf_product_id)
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[GENERIC_STRIPE_ID]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

    # basic customer creation check
    # without a 1:1 contact relationship, the email is nil
    refute_empty(customer.name)
    assert_nil(customer.email)

    assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
    assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)
    # TODO customer address assertions once mapping is complete

    # top level subscription fields
    assert_match(sf_order.Id, subscription_schedule.metadata['salesforce_order_link'])
    assert_equal(subscription_schedule.metadata['salesforce_order_id'], sf_order.Id)

    # line-level subscription phase data
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    assert_equal(1, phase.items.count)
    phase_item = T.must(phase.items.first)
    assert_equal(sf_pricebook_entry[GENERIC_STRIPE_ID], phase_item.plan)
    assert_equal(sf_pricebook_entry[GENERIC_STRIPE_ID], phase_item.price)
    assert_equal(1, phase_item.quantity)

    invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

    assert_equal(1, invoices.count)
    invoice = invoices.first

    # the stripe invoice is not 1:1 linked with any SF record
    assert_empty(invoice.metadata.to_h)

    # should be two lines since we are backdating the invoice
    assert_equal(2, invoice.lines.count)

    # TODO test first line as well

    line = invoice.lines.data.last
    assert_equal(1, line.quantity)
    assert_equal(price, line.amount)

    # right now, price translation is tied to both a pricebook and order line in SF
    # test the price translation logic here right now instead of in a separate price test
    stripe_price = line.price
    assert_equal(price, stripe_price.unit_amount)
    assert_equal("recurring", stripe_price.type)
    assert_equal("month", stripe_price.recurring.interval)
    assert_equal(1, stripe_price.recurring.interval_count)
    assert_equal("licensed", stripe_price.recurring.usage_type)

    assert_match(@user.salesforce_instance_url, stripe_price.metadata['salesforce_pricebook_entry_link'])
    assert_match(sf_pricebook_entry_id, stripe_price.metadata['salesforce_pricebook_entry_link'])
    assert_equal(stripe_price.metadata['salesforce_pricebook_entry_id'], sf_pricebook_entry_id)
  end

  describe 'failures' do
    it 'gracefully fails when the subscription term and start date does not exist' do
      sf_product_id = salesforce_recurring_product_with_price
      sf_order = create_salesforce_order(sf_product_id: sf_product_id)

      assert_raises(Integrations::Errors::MissingRequiredFields) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      # TODO assert on specific fields which are missing
      # TODO assert on custom object in SF
    end
  end

  # TODO multiple quantity
  # TODO quantity as float
  # TODO start date in the future
  # TODO missing fields / failure
  # TODO subscription term specified

  it 'integrates a invoice order' do
    @user.update(field_mappings: {
      customer: {
        # if accounts are mapped to customer, there is no default email field
        "Description": "email",
      },
    })

    sf_account_id = create_salesforce_account(additional_fields: {
      "Description" => "#{Time.now.to_i}@example.com",
    })

    sf_order = create_salesforce_order(sf_account_id: sf_account_id, sf_product_id: salesforce_standalone_product_with_price_id)

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order = sf.find(SF_ORDER, sf_order.Id)

    stripe_invoice_id = sf_order[GENERIC_STRIPE_ID]
    invoice = Stripe::Invoice.retrieve(stripe_invoice_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(invoice.customer, @user.stripe_credentials)

    refute_empty(customer.email)

    assert_equal(1, invoice.lines.count)

    line = invoice.lines.first
    assert_equal("one_time", line.price.type)

    puts sf_order.Id
  end
end
