# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'integrates a standard subscription order' do
    @user.field_mappings = {
      "subscription_schedule" => {
        "metadata.example_field" => "OrderNumber",
      },
    }
    @user.save

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
    assert_equal(subscription_schedule.metadata['example_field'], sf_order.OrderNumber)

    # subscription phase fields
    subscription = Stripe::Subscription.retrieve(T.cast(subscription_schedule.subscription, String), @user.stripe_credentials)
    assert_match(sf_order.Id, subscription.metadata['salesforce_order_link'])
    assert_equal(subscription.metadata['salesforce_order_id'], sf_order.Id)

    # line-level subscription phase data
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)
    assert_match(sf_order.Id, phase.metadata['salesforce_order_link'])
    assert_equal(phase.metadata['salesforce_order_id'], sf_order.Id)
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

    # if proration is enabled, we would see two lines, but since we are disabling proration by default we'll only see a single line
    assert_equal(1, invoice.lines.count)

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

    # lets ensure a duplicate subscription schedule is not created
    Stripe::SubscriptionSchedule.expects(:create).never

    SalesforceTranslateRecordJob.translate(@user, sf_order)
  end

  it 'does not filter out $0 line items' do
    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price(price: 0)

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    assert_equal(2, phase.items.count)
    assert_equal(0, phase.add_invoice_items.count)

    sf_zero_dollar_price = sf_get(sf_pricebook_id_2)
    stripe_zero_dollar_price_id = sf_zero_dollar_price[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_zero_dollar_price = Stripe::Price.retrieve(stripe_zero_dollar_price_id, @user.stripe_credentials)

    # ensure price was correctly created
    assert_equal(0, stripe_zero_dollar_price.unit_amount)

    phase_item_with_zero_price = T.must(phase.items.detect {|i| i.price == stripe_zero_dollar_price_id })
    assert_equal(sf_pricebook_id_2, stripe_zero_dollar_price.metadata['salesforce_pricebook_entry_id'])

    phase_item_with_normal_price = T.must(phase.items.detect {|i| i.price != stripe_zero_dollar_price_id })
    normal_price = Stripe::Price.retrieve(phase_item_with_normal_price.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_1, normal_price.metadata['salesforce_pricebook_entry_id'])
  end

  it 'skips line items when the skip line item custom field is checked' do
    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price
    sf_product_id_3, sf_pricebook_id_3 = salesforce_recurring_product_with_price

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    # only CPQ fields can be customized within this special quote creation process

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_3)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    sf_line_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(3, sf_line_items.size)
    second_line_item = sf_line_items.detect {|i| i.Product2Id == sf_product_id_2 }

    sf.update!(SF_ORDER_ITEM, {
      SF_ID => second_line_item.Id,
      prefixed_stripe_field(ORDER_LINE_SKIP) => true,
    })

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    # should have three, but one was excluded
    assert_equal(2, phase.items.count)
    assert_equal(0, phase.add_invoice_items.count)

    assert_empty(phase.items.select do |i|
      refute_nil(i.metadata['salesforce_order_item_id'])
      i.metadata['salesforce_order_item_id'] == second_line_item.Id
    end)
  end

  it 'integrates a subscription order with multiple lines' do
    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price
    sf_product_id_3, sf_pricebook_id_3 = salesforce_standalone_product_with_price

    # set the second product up as a metered billing
    sf.update!(SF_PRODUCT, {
      SF_ID => sf_product_id_2,
      CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize,
    })

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    # set first product quantity to 5
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    quote_with_product["lineItems"].first["record"]["SBQQ__Quantity__c"] = 5.0
    calculate_and_save_cpq_quote(quote_with_product)

    # recurring and arrears, should not have a quantity set when passed to stripe
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    # third product should be standalone and not recurring
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

    phase_item_with_five = T.must(phase.items.detect {|i| i.try(:quantity) == 5 })
    price_1 = Stripe::Price.retrieve(phase_item_with_five.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_1, price_1.metadata['salesforce_pricebook_entry_id'])

    phase_item_with_metered = T.must(phase.items.detect {|i| !i.respond_to?(:quantity) })
    price_2 = Stripe::Price.retrieve(phase_item_with_metered.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_2, price_2.metadata['salesforce_pricebook_entry_id'])

    standalone_item = T.must(phase.add_invoice_items.first)
    price_3 = Stripe::Price.retrieve(standalone_item.price, @user.stripe_credentials)
    assert_equal(sf_pricebook_id_3, price_3.metadata['salesforce_pricebook_entry_id'])
  end

  # TODO reuses order line price mapping if the execution halts part way through

  describe 'failures' do
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

      assert_match("The following required fields are missing from this Salesforce record: SBQQ__Quote__c.SBQQ__StartDate__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])
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
      assert_equal(SF_ORDER_ITEM, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

      assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
      assert_match(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::COMPOUND_ID.serialize)])
      assert_match("The following required fields are missing from this Salesforce record: Order.SBQQ__Quote__c.SBQQ__SubscriptionTerm__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])

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
