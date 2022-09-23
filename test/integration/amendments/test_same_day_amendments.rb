# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::SameDayAmendments < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
  end

  def get_proration_invoice_item(stripe_customer_id)
    invoice_items = Stripe::InvoiceItem.list({customer: stripe_customer_id}, @user.stripe_credentials)
    assert_equal(1, invoice_items.count)
    invoice_item = invoice_items.first
    assert_equal("true", invoice_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
    invoice_item
  end

  # https://jira.corp.stripe.com/browse/PLATINT-1808
  it 'supports amending an order today which has already started and is billed yearly' do
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

    # initial order: one product, starting at midnight tomorrow
    # test clock: advanced to noon tomorrow,
    # amendment: add +1 quantity of the same product, starting tomorrow

    # the state we expect:
    #   * Autogenerated (by billing) invoice for the initial subscription
    #   * Invoice items (from us) for the prorated amount
    #       * These amounts should be the full amount of the product since the start date is the same as the original order

    contract_term = TEST_DEFAULT_CONTRACT_TERM
    amendment_term = TEST_DEFAULT_CONTRACT_TERM
    initial_start_date = now_time + 1.day
    amendment_start_date = initial_start_date + (contract_term - amendment_term).month
    amendment_end_date = amendment_start_date + amendment_term.months

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    # translate the initial order THEN advance the clock to the middle of the day tomorrow
    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    # get the test clock from the customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)
    test_clock = advance_test_clock(stripe_customer, (initial_start_date + 12.hours).to_i)

    # now that we've advanced the test clock time, lets create the amendment
    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity by 1
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2

    # midnight of the current day!
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(initial_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    # test_clock = advance_test_clock(stripe_customer, (initial_start_date + 24.hours).to_i)

    # one phase running for less than a day with a single quantity
    # second phase starting midday for a single quantity, billing the prorated amount
    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with a single quantity
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_equal(1, first_phase_item[:quantity])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(12.hours.to_i, first_phase.end_date - initial_start_date.to_i)

    # second phase should start at the end date
    assert_equal(12.hours.to_i, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # second phase should have a second item with a quantity of 1
    # and a prorated amount for the second quantity that was added
    assert_equal(2, second_phase.items.count)
    assert_equal(1, second_phase.add_invoice_items.count)
    first_item_from_initial_phase = T.must(second_phase.items.detect {|i| i[:price] == first_phase_item.price })
    second_item = T.must(second_phase.items.detect {|i| i[:price] != first_phase_item.price })
    assert_equal(1, first_item_from_initial_phase.quantity)
    assert_equal(1, second_item.quantity)

    # ensure the prorated amount is the full billing price for the item
    prorated_item = T.must(second_phase.add_invoice_items.first)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
    assert_equal(TEST_DEFAULT_PRICE, prorated_price.unit_amount_decimal.to_i)
    assert_equal(1, prorated_item.quantity)
  end

  it 'supports amending an order today which has already started and is billed monthly' do
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

    # initial order: one product, billed monthly, starting at midnight tomorrow
    # test clock: advanced to noon tomorrow,
    # amendment: add +1 quantity of the same product, starting tomorrow

    # the state we expect:
    #   * Autogenerated (by billing) invoice for the initial subscription
    #   * Invoice items (from us) for the prorated amount
    #       * These amounts should be the full amount of the product since the start date is the same as the original order

    contract_term = TEST_DEFAULT_CONTRACT_TERM
    amendment_term = TEST_DEFAULT_CONTRACT_TERM
    initial_start_date = now_time + 1.day
    amendment_start_date = initial_start_date + (contract_term - amendment_term).month
    amendment_end_date = amendment_start_date + amendment_term.months

    sf_order = create_salesforce_order(
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    # translate the initial order THEN advance the clock to the middle of the day tomorrow
    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)
    test_clock = advance_test_clock(stripe_customer, (initial_start_date + 12.hours).to_i)

    # now that we've advanced the test clock time, lets create the amendment
    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity by 1
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2

    # midnight of the current day!
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(initial_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    # one phase running for less than a day with a single quantity
    # second phase starting midday for a single quantity, billing the prorated amount
    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with a single quantity
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_equal(1, first_phase_item[:quantity])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(12.hours.to_i, first_phase.end_date - initial_start_date.to_i)

    # second phase should start at the end date
    assert_equal(12.hours.to_i, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # second phase should have a second item with a quantity of 1
    # and a prorated amount for the second quantity that was added
    assert_equal(2, second_phase.items.count)
    assert_equal(1, second_phase.add_invoice_items.count)
    first_item_from_initial_phase = T.must(second_phase.items.detect {|i| i[:price] == first_phase_item.price })
    second_item = T.must(second_phase.items.detect {|i| i[:price] != first_phase_item.price })
    assert_equal(1, first_item_from_initial_phase.quantity)
    assert_equal(1, second_item.quantity)

    # ensure the prorated amount is the full billing price for the item
    prorated_item = T.must(second_phase.add_invoice_items.first)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
    assert_equal(TEST_DEFAULT_PRICE, prorated_price.unit_amount_decimal.to_i)
    assert_equal(1, prorated_item.quantity)

    # advance test clock a couple hours (to move the add_invoice_item to the subscription) and test invoicing
    test_clock = advance_test_clock(stripe_customer, (initial_start_date + 24.hours).to_i)

    invoice_item = get_proration_invoice_item(stripe_customer.id)

    invoice_event = get_invoice_item_event(invoice_item.id)
    prorated_invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event)
    assert_equal(TEST_DEFAULT_PRICE, T.must(prorated_invoice).total)

    # since this is a monthly subscription, advance the clock a month to ensure the next subscription is billed normally
    test_clock = advance_test_clock(stripe_customer, (initial_start_date + 1.month + 1.day).to_i)
    invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials)

    # first monthly invoice, prorated invoice, next month invoice
    assert_equal(3, invoices.count)
    auto_invoice = invoices.data.last
    # last invoice should not be generated by proration logic
    refute(auto_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    # autogen'd invoice should have the default pricing
    assert_equal(TEST_DEFAULT_PRICE, auto_invoice.total)
  end

  # https://jira.corp.stripe.com/browse/PLATINT-1815
  it 'supports amending a order on the start date that has not yet started' do
    # start date is in the future, but amendment date is the same as the initial order start date
  end
end
