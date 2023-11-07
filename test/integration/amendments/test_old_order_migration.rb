# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::OldOrderMigration < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
    # note: enabling non_anniversary amendments should not affect anniversary amendments
    # therefore enable for this for the entire test suite
    @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
    @user.enable_feature FeatureFlags::OLD_ORDER_MIGRATIONS, update: true
  end

  it 'migrates an order billed outside of Stripe with annual pricing' do
    # initial order is 1 year with a yearly billed item
    # amendment order 6 months in with three additional items
    yearly_price = 100_00
    contract_term = 12
    amendment_term = 6

    # Initial Order was 5 months ago (offset by a month so the amendment begins in one month)
    initial_start_date = now_time - amendment_term.months + 1.month
    initial_order_end_date = initial_start_date + contract_term.months

    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months

    initial_quantity = 1
    amendment_quantity = initial_quantity + 3

    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                         contact_email: "migrate_out_of_stripe",
                                         additional_fields: {
                                           CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
                                           CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                         })
    sf_contract = create_contract_from_order(sf_order)

    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at + 1.minute)
    end
    # Important that this is set after the order is activated, to simulate the real scenario (even if its only by a few seconds instead of months)
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = Time.now.to_i

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # Increase the quantity to three
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = amendment_quantity
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

    invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription, limit: 100}, @user.stripe_credentials)

    # invoice total should be for only one month
    assert_equal(0, invoices.data.length)

    # get Stripe customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)

    # Fast forward to start of amendment, to see our add_invoice_items get created
    test_clock = advance_test_clock(stripe_customer, (amendment_start_date).to_i)

    # simulate sending the webhook
    events = Stripe::Event.list({
      type: 'invoiceitem.created',
    }, @user.stripe_credentials)

    invoice_events = events.data.select do |event|
      event_object = event.data.object
      event_object.test_clock == test_clock.id && event.request.id.nil?
    end

    assert_equal(1, invoice_events.count)
    invoice_event = invoice_events.first

    invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event))

    # 6 months on 12 month price * 3 new items
    assert_equal(50_00 * 3, invoice.total)
    invoice.pay({paid_out_of_band: true})

    test_clock = advance_test_clock(stripe_customer, (initial_order_end_date + 1.day).to_i)

    invoices = Stripe::Invoice.list({customer: stripe_customer.id, paid: false}, @user.stripe_credentials)

    assert_equal(0, invoices.data.length)
  end

  it 'migrates an order billed outside of Stripe with monthly pricing' do
    # initial order is 1 year with a monthly billed item
    # amendment order 6 months and 15 days in with three additional items
    monthly_price = 100_00
    contract_term = 12
    amendment_term = 6

    # Initial Order was 5 months ago (offset by a month so the amendment begins in one month)
    initial_start_date = now_time - amendment_term.months + 1.month
    initial_order_end_date = initial_start_date + contract_term.months

    amendment_start_date = initial_start_date + (contract_term - amendment_term).months + 15.days
    amendment_end_date = amendment_start_date + amendment_term.months

    initial_quantity = 1
    amendment_quantity = initial_quantity + 3

    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: monthly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
      }
    )

    sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                         contact_email: "migrate_out_of_stripe_monthly",
                                         additional_fields: {
                                           CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
                                           CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                         })
    sf_contract = create_contract_from_order(sf_order)

    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at + 2.minute)
    end
    sleep(5)
    # Important that this is set after the order is activated, to simulate the real scenario (even if its only by a few seconds instead of months)
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = Time.now.to_i

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # Increase the quantity to three
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = amendment_quantity
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term - 1

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

    invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription, limit: 100}, @user.stripe_credentials)

    # invoice total should be for only one month
    assert_equal(0, invoices.data.length)

    # get Stripe customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)

    # Fast forward to start of amendment, to see our add_invoice_items get created
    test_clock = advance_test_clock(stripe_customer, (amendment_start_date).to_i)

    # simulate sending the webhook
    events = Stripe::Event.list({
      type: 'invoiceitem.created',
    }, @user.stripe_credentials)

    invoice_events = events.data.select do |event|
      event_object = event.data.object
      event_object.test_clock == test_clock.id && event.request.id.nil?
    end

    assert_equal(1, invoice_events.count)
    invoice_event = invoice_events.first

    invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event))

    # 15 days on monthly price * 3 new items
    refute_nil(invoice)
    invoice.pay({paid_out_of_band: true})

    # Can only advance a test clock by 2 intervals at a time
    test_clock = advance_test_clock(stripe_customer, test_clock.frozen_time  + 2.months.to_i)
    test_clock = advance_test_clock(stripe_customer, test_clock.frozen_time  + 2.months.to_i)
    test_clock = advance_test_clock(stripe_customer, (initial_order_end_date + 1.day).to_i)

    invoices = Stripe::Invoice.list({customer: stripe_customer.id, paid: false, limit: 100}, @user.stripe_credentials)

    assert_equal(6, invoices.data.length)
  end

  it 'does not create a subscription when the old order does not match the custom soql queries' do
    # initial order is 1 year with a yearly billed item
    # amendment order 6 months in with three additional items
    yearly_price = 100_00
    contract_term = 12
    amendment_term = 6

    # Initial Order was 5 months ago (offset by a month so the amendment begins in one month)
    initial_start_date = now_time - amendment_term.months + 1.month
    initial_order_end_date = initial_start_date + contract_term.months

    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months

    initial_quantity = 1
    amendment_quantity = initial_quantity + 3

    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                         contact_email: "no_create_sub",
                                         additional_fields: {
                                           CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
                                           CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                         })
    sf_contract = create_contract_from_order(sf_order)

    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at + 1.minute)
    end
    # Important that this is set after the order is activated, to simulate the real scenario (even if its only by a few seconds instead of months)
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = Time.now.to_i

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # Increase the quantity to three
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = amendment_quantity
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    @user.connector_settings['filters'][SF_ORDER] = "Id != '#{sf_order.Id}'"

    exception = assert_raises(StripeForce::Errors::UserError) do
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    end

    sf_order.refresh

    # Order should not have translated
    assert_nil(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
  end

  it 'does not translate a subscription when merchant is not gated' do
    # Disable the feature flag
    @user.disable_feature FeatureFlags::OLD_ORDER_MIGRATIONS, update: true

    # initial order is 1 year with a yearly billed item
    # amendment order 6 months in with three additional items
    yearly_price = 100_00
    contract_term = 12
    amendment_term = 6

    # Initial Order was 5 months ago (offset by a month so the amendment begins in one month)
    initial_start_date = now_time - amendment_term.months + 1.month
    initial_order_end_date = initial_start_date + contract_term.months

    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months

    initial_quantity = 1
    amendment_quantity = initial_quantity + 3

    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                         contact_email: "no_translate_sub",
                                         additional_fields: {
                                           CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
                                           CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                         })
    sf_contract = create_contract_from_order(sf_order)

    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at + 1.minute)
    end
    # Important that this is set after the order is activated, to simulate the real scenario (even if its only by a few seconds instead of months)
    @user.connector_settings[CONNECTOR_SETTING_SYNC_START_DATE] = Time.now.to_i

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # Increase the quantity to three
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = amendment_quantity
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh

    # Order should not have translated
    assert_nil(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
  end
end
