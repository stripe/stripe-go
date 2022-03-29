# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  def salesforce_recurring_product_with_price
    # blanking out the subscription type ensures it is a one-time product
    product_id = create_salesforce_product(additional_fields: {
      # anything non-nil indicates subscription/recurring pricing
      CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

      CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::RENEWABLE,

      # default term of one year
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12,

      # one month
      CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
    })

    pricebook_entry_id = create_salesforce_price(sf_product_id: product_id)
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

  def create_subscription_order
    create_salesforce_order(additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now.strftime("%Y-%m-%d"),
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })
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
  end

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

    sf.update!('Order', 'Id' => sf_order.Id, 'Status' => 'Activated')

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

  it 'integrates a standard subscription order' do
    price = 120_00
    terms = 12

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => "2021-10-01",
        CPQ_QUOTE_SUBSCRIPTION_TERM => terms,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    # TODO add `refresh` to salesforce library
    sf_order = sf.find(SF_ORDER, sf_order.Id)
    sf_product = sf.find(SF_PRODUCT, sf_product_id)
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
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
    # NOTE iterations does not exist on the phase! https://jira.corp.stripe.com/browse/PLATINT-1479
    # TODO I have no idea why the math requires rounding here. This doesn't make any sense. https://jira.corp.stripe.com/browse/PLATINT-1480
    phase_iterations = ((phase.end_date - phase.start_date) / 1.month.to_f).round
    assert_equal(terms, phase_iterations)

    assert_equal(1, phase.items.count)
    phase_item = T.must(phase.items.first)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.plan)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.price)
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

  it 'integrates a subscription order with multiple lines' do
    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price
    sf_product_id_3, sf_pricebook_id_3 = salesforce_standalone_product_with_price

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    # set first product quantity to 5
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    quote_with_product["lineItems"].first["record"]["SBQQ__Quantity__c"] = 5.0
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_3)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    assert_equal(2, phase.items.count)
    assert_equal(1, phase.add_invoice_items.count)

    phase_item_with_five = T.must(phase.items.detect {|i| i.quantity == 5 })
    price_1 = Stripe::Price.retrieve(phase_item_with_five.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_1, price_1.metadata['salesforce_pricebook_entry_id'])

    phase_item_with_one = T.must(phase.items.detect {|i| i.quantity == 1 })
    price_2 = Stripe::Price.retrieve(phase_item_with_one.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_2, price_2.metadata['salesforce_pricebook_entry_id'])

    standalone_item = T.must(phase.add_invoice_items.first)
    price_3 = Stripe::Price.retrieve(standalone_item.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_3, price_3.metadata['salesforce_pricebook_entry_id'])
  end

  describe 'integration overrides' do
    it 'integrates a subscription order when an invalid ID is entered in the stripe ID field' do
      sf_order = create_subscription_order

      sf.update!(sf_order.sobject_type, {
        SF_ID => sf_order.Id,
        # insert a random ID that does not represent a stripe object
        prefixed_stripe_field(GENERIC_STRIPE_ID) => SecureRandom.alphanumeric(16),
      })

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh

      assert_match(/^sub_sched_/, sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
      schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
    end

    it 'ignores a valid stripe ID of the wrong object type' do
      stripe_customer = Stripe::Customer.create({}, @user.stripe_credentials)

      sf_order = create_salesforce_order(additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now.strftime("%Y-%m-%d"),
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      })

      sf.update!(sf_order.sobject_type, {
        SF_ID => sf_order.Id,
        # insert a ID that represents a valid stripe object of the wrong type
        prefixed_stripe_field(GENERIC_STRIPE_ID) => stripe_customer.id,
      })

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh

      assert_match(/^sub_sched_/, sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
      schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
    end

    it 'accepts a valid stripe ID and does not create a additional object in Stripe'
  end

  describe 'failures' do
    def get_sync_record_by_primary_id(primary_id)
      sync_record_results = sf.query("SELECT Id FROM #{prefixed_stripe_field(SYNC_RECORD)} WHERE #{prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)} = '#{primary_id}'")
      assert_equal(1, sync_record_results.count)

      sync_record_id = sync_record_results.first.Id
      sync_record = sf.find(prefixed_stripe_field(SYNC_RECORD), sync_record_id)
    end

    it 'creates a sync failure when the start date does not exist' do
      sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price
      sf_order = create_salesforce_order(
        sf_product_id: sf_product_id,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
          # intentional ommission of the start date
        }
      )

      exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      sync_record = get_sync_record_by_primary_id(sf_order.Id)

      assert_match("The following required fields are missing from this Salesforce record: SBQQ__StartDate__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])
    end

    # it looks like there are field validations in place to protect against the term being nil'd out on the order line
    it 'creates a user error when the subscription term does not exist on the order line level' do
      sf_order = create_salesforce_order(additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now.strftime("%Y-%m-%d"),
        # omit subscription term
      })

      exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      sync_record = get_sync_record_by_primary_id(sf_order.Id)

      assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(CPQ_QUOTE, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

      assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
      assert_match(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::COMPOUND_ID.serialize)])
      assert_match("The following required fields are missing from this Salesforce record: SBQQ__SubscriptionTerm__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])

      assert_match(@user.salesforce_instance_url, sync_record[prefixed_stripe_field('Primary_Record__c')])
      assert_match(@user.salesforce_instance_url, sync_record[prefixed_stripe_field('Secondary_Record__c')])
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
        "email": "Description",
      },
    })

    sf_account_id = create_salesforce_account(additional_fields: {
      # an email is required for creating an invoice without a payment method
      "Description" => "#{Time.now.to_i}@example.com",
    })

    sf_product_id, _ = salesforce_standalone_product_with_price

    sf_order = create_salesforce_order(
      sf_account_id: sf_account_id,
      sf_product_id: sf_product_id
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_invoice_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    invoice = Stripe::Invoice.retrieve(stripe_invoice_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(invoice.customer, @user.stripe_credentials)

    refute_empty(customer.email)

    assert_equal(1, invoice.lines.count)

    line = invoice.lines.first
    assert_equal("one_time", line.price.type)
  end
end
