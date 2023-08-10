# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::BackdatedOrders < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::BACKDATED_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
  end

  describe 'backdated amendment order' do
    it 'syncs a backdated initial and amendment order billed monthly' do
      # initial order: starts in the past, billed monthly
      # amendment order: starts 1 month later
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time - 2.months - 1.day
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_order_start_date = initial_order_start_date + (contract_term - amendment_term).months
      monthly_price = 24_00

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: monthly_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "syncs_backdated_monthly",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get Stripe customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      # three months of services
      assert_equal(monthly_price * 3, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)
      T.must(first_phase.items.first)

      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # since we use 'now' for the phase date, there will be an offset but ensure it's less than a day
      # assert((first_phase.end_date - now_time.to_i) < SECONDS_IN_DAY)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - first_phase.end_date)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (the original and the addition) and should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # sanity check the generated invoice
      # although we're using the current time to update the subscription schedule
      # the generated invoice should be created for the backdated date

      # advance the test clock
      test_clock = advance_test_clock(stripe_customer, (now_time + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)] == "true"
      end
      assert_equal(1, invoice_events.count)

      proration_invoice_event = invoice_events[0]
      proration_invoice_data = proration_invoice_event.data.object
      assert_equal(1, proration_invoice_event.data.object.quantity)
      assert_equal(monthly_price * 2, proration_invoice_data.amount)
      assert_equal((monthly_price * amendment_term) / 100, sf_order_amendment["TotalAmount"])
      assert_equal("true", proration_invoice_data.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # confirm the proration invoice dates
      assert_equal(amendment_order_start_date.to_i, proration_invoice_data.period.start)
      # the proration invoice end date should be until the next billing cycle date (no missed period)
      assert_equal((now_time - 1.day + 1.month).to_i, proration_invoice_data.period.end)
    end

    it 'syncs a backdated initial and amendment order billed monthly with a half open period' do
      @user.enable_feature FeatureFlags::BILLING_GATE_OPEN_INVOICING_INTERVAL, update: true

      # initial order: starts in the past, billed monthly
      # amendment order: starts 1 month later
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time - 4.months - 1.day
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_order_start_date = initial_order_start_date + (contract_term - amendment_term).months
      monthly_price = 24_00

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: monthly_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "syncs_backdated_monthly_half",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      Timecop.freeze(Time.now + 1.minute)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get Stripe customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      # three months of services
      assert_equal(monthly_price * 5, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)
      T.must(first_phase.items.first)

      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # since we use 'now' for the phase date, there will be an offset but ensure it's less than a day
      assert((first_phase.end_date - now_time.to_i) < SECONDS_IN_DAY)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - first_phase.end_date)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (the original and the addition) and should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # sanity check the generated invoice
      # although we're using the current time to update the subscription schedule
      # the generated invoice should be created for the backdated date

      # advance the test clock
      test_clock = advance_test_clock(stripe_customer, (now_time + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)] == "true"
      end
      assert_equal(1, invoice_events.count)

      proration_invoice_event = invoice_events[0]
      proration_invoice_data = proration_invoice_event.data.object
      assert_equal(1, proration_invoice_event.data.object.quantity)
      assert_equal(monthly_price * 4, proration_invoice_data.amount)
      assert_equal((monthly_price * amendment_term) / 100, sf_order_amendment["TotalAmount"])
      assert_equal("true", proration_invoice_data.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # confirm the proration invoice dates
      assert_equal(amendment_order_start_date.to_i, proration_invoice_data.period.start)
      # the proration invoice end date should be until the next billing cycle date (no missed period)
      assert_equal((now_time + 1.month - 2.days).utc.beginning_of_day.to_i, Time.at(proration_invoice_data.period.end).utc.beginning_of_day.to_i)
    end

    it 'syncs a backdated initial and amendment order billed annually' do
      # initial order: starts in the past, billed annually
      # amendment order: starts 1 month later
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time - 4.months
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_order_start_date = initial_order_start_date + (contract_term - amendment_term).months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "syncs_backdated_annual",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)

      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # since we use 'now' for the phase date, there will be an offset
      # ensure it's less than a day
      assert((first_phase.end_date - now_time.to_i) < SECONDS_IN_DAY)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - first_phase.end_date)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (the original and the addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      assert_equal(1, prorated_item.quantity)
      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      expected_prorated_price = TEST_DEFAULT_PRICE * BigDecimal(amendment_term) / BigDecimal(contract_term)
      assert_equal(TEST_DEFAULT_PRICE * BigDecimal(amendment_term) / BigDecimal(contract_term), prorated_price.unit_amount_decimal.to_i)

      prorated_invoice_amount_in_dollars = BigDecimal(prorated_price.unit_amount_decimal.to_i) / 100
      assert_equal(prorated_invoice_amount_in_dollars, sf_order_amendment["TotalAmount"])

      # get customer
      sf_account = sf_get(sf_order_amendment['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      assert_equal(TEST_DEFAULT_PRICE, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # sanity check the generated invoice
      # although we're using the current time to update the subscription schedule
      # the generated invoice should be created for the backdated date

      # advance the test clock
      test_clock = advance_test_clock(stripe_customer, (now_time + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event_object.unit_amount == expected_prorated_price
      end
      assert_equal(1, invoice_events.count)
      proration_invoice_event = invoice_events[0].data.object
      assert_equal("true", proration_invoice_event.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # verify the dates on the proration invoice dates
      assert_equal(amendment_order_start_date.to_i, proration_invoice_event.period.start)
      assert_equal(initial_order_end_date.to_i, Time.at(proration_invoice_event.period.end).utc.beginning_of_day.to_i)

      assert_equal(expected_prorated_price, proration_invoice_event.unit_amount_decimal.to_i)
      assert_equal("true", proration_invoice_event.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
      assert_equal(expected_prorated_price / 100, sf_order_amendment["TotalAmount"])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_events[0]))
      assert_equal(1, proration_invoice.lines.count)
      assert_equal(expected_prorated_price, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    it 'syncs a backdated initial and amendment order billed annually with a half open phase' do
      @user.enable_feature FeatureFlags::BILLING_GATE_OPEN_INVOICING_INTERVAL, update: true

      # initial order: starts in the past, billed annually
      # amendment order: starts 1 month later
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time - 4.months
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_order_start_date = initial_order_start_date + (contract_term - amendment_term).months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "syncs_backdated_annual_half",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)

      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # since we use 'now' for the phase date, there will be an offset
      # ensure it's less than a day
      # assert((first_phase.end_date - now_time.to_i) < SECONDS_IN_DAY)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - first_phase.end_date)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (the original and the addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      assert_equal(1, prorated_item.quantity)
      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      expected_prorated_price = TEST_DEFAULT_PRICE * BigDecimal(amendment_term) / BigDecimal(contract_term)
      assert_equal(expected_prorated_price, prorated_price.unit_amount_decimal.to_i)

      prorated_invoice_amount_in_dollars = BigDecimal(prorated_price.unit_amount_decimal.to_i) / 100
      assert_equal(prorated_invoice_amount_in_dollars, sf_order_amendment["TotalAmount"])

      # get customer
      sf_account = sf_get(sf_order_amendment['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      assert_equal(TEST_DEFAULT_PRICE, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # sanity check the generated invoice
      # although we're using the current time to update the subscription schedule
      # the generated invoice should be created for the backdated date

      # advance the test clock
      test_clock = advance_test_clock(stripe_customer, (now_time + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event_object.unit_amount == expected_prorated_price
      end
      assert_equal(1, invoice_events.count)
      proration_invoice_event = invoice_events[0].data.object
      assert_equal("true", proration_invoice_event.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # verify the dates on the proration invoice dates
      assert_equal(amendment_order_start_date.to_i, proration_invoice_event.period.start)
      assert_equal((initial_order_end_date - 1.day).to_i, Time.at(proration_invoice_event.period.end).utc.beginning_of_day.to_i)

      assert_equal(expected_prorated_price, proration_invoice_event.unit_amount_decimal.to_i)
      assert_equal("true", proration_invoice_event.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
      assert_equal(expected_prorated_price / 100, sf_order_amendment["TotalAmount"])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_events[0]))
      assert_equal(1, proration_invoice.lines.count)
      assert_equal(expected_prorated_price, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    it 'syncs a backdated initial and amendment order billed monthly with a billing cycle passed' do
      # initial order: starts almost three months in the past, billed monthly
      # amendment order: starts one month after, in the past
      contract_term = 12
      amendment_term = 11
      backdated_months = 3
      initial_order_start_date = now_time - backdated_months.months + 1.days
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_order_start_date = initial_order_start_date + (contract_term - amendment_term).months
      monthly_price = 111_00

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: monthly_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "syncs_backdated_month_bill_pass",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get Stripe customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      # assert_equal(monthly_price * backdated_months, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create the sf amendment order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the sf amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # get the generated subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # sanity check subscription schedule first phase
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # since we use 'now' for the phase date, there will be an offset but ensure it's less than a day
      assert((first_phase.end_date - now_time.to_i) < SECONDS_IN_DAY)
      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      T.must(first_phase.items.first)
      assert_empty(first_phase.add_invoice_items)

      # sanity check subscription schedule second phase
      second_phase = T.must(stripe_subscription_schedule.phases.last)
      assert_equal(0, second_phase.start_date - first_phase.end_date)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)
      # second phase should have two items (the original and the addition) and should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # advance the test clock
      test_clock = advance_test_clock(stripe_customer, (now_time + 10.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)] == "true"
      end
      assert_equal(1, invoice_events.count)

      proration_invoice_event = invoice_events[0].data.object
      assert_equal(monthly_price * 2, proration_invoice_event.unit_amount_decimal.to_i)
      assert_equal("true", proration_invoice_event.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
      assert_equal((monthly_price * amendment_term) / 100, sf_order_amendment["TotalAmount"])

      # verify the dates on the proration invoice dates
      assert_equal(amendment_order_start_date.to_i, proration_invoice_event.period.start)
      assert_equal((now_time + 1.days).to_i, Time.at(proration_invoice_event.period.end).utc.beginning_of_day.to_i)
    end
  end

  describe 'skip initial invoice of backdated initial order' do
    it 'sync a backdated initial order billing monthly skipping the initial invoice' do
      @user.field_defaults = {
        "customer" => {
          "email" => "test@test.com",
        },
        "subscription_schedule" => {
          "default_settings.invoice_settings.days_until_due" => "Net-15",
          "default_settings.collection_method" => "send_invoice", # note: you can only specify 'days_until_due' if invoice collection method is 'send_invoice'
        },
      }
      @user.save

      # initial order is one year contract with a monthly product
      # amendment order in three months adding an additional item
      contract_term = 12
      amendment_term = 8

      # initial order starts three months + two days ago
      initial_order_start_date = (now_time - 2.days) - 3.months
      initial_order_end_date = initial_order_start_date + contract_term.months

      # amendment order starts 4 months into contract
      amendment_start_date = initial_order_start_date + 4.month

      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      # create and sync a backdated initial order
      sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                           contact_email: "skip_initial_invoice",
                                           additional_fields: {
                                             CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
                                             CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                           })

      # set SKIP_PAST_INITIAL_INVOICES to true
      sf.update!(SF_ORDER,
        SF_ID => sf_order.Id,
        prefixed_stripe_field(SKIP_PAST_INITIAL_INVOICES) => true
      )

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

      # confirm no intial invoice was created for the backdated period
      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)
      assert_equal(0, invoices.data.length)

      # get Stripe customer and advance the test clock
      stripe_customer = stripe_get(subscription_schedule.customer)
      refute_nil(stripe_customer.test_clock)
      advance_test_clock(stripe_customer, (now_time + 1.months).to_i)

      # confirm the due date of the next invoice is 15 days from when the subscription started (not synced)
      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription, limit: 100}, @user.stripe_credentials)
      assert_equal(1, invoices.data.length)

      # TODO @nada, need to confirm if this is the correct behavior
      # currently we skip the current invoice if synced on anniversary day
      next_invoice = invoices.data.first
      days_until_due = 15
      assert_equal(((now_time - 2.days) + 1.month + days_until_due.days).utc.beginning_of_day.to_i, Time.at(next_invoice.due_date).utc.beginning_of_day.to_i)
      assert_equal(TEST_DEFAULT_PRICE, next_invoice.amount_due)
      assert_equal(1, next_invoice.lines.data.length)
      next_invoice.pay({'paid_out_of_band': true})

      # create and sync an amendment order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      # Increase the quantity to three
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # confirm the amendment order was processed correctly
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
      assert_equal(2, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start at the amendment start date
      # and have two products with total quantity of 2
      assert(second_phase.start_date.to_i, amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)
      # second phase should have two items with a total quantity of 2
      assert_equal(2, second_phase.items.count)
      second_phase_item_1 = T.must(second_phase.items.first)
      second_phase_item_2 = T.must(second_phase.items.second)
      assert_equal(1, second_phase_item_1.quantity)
      assert_equal(1, second_phase_item_2.quantity)

      # fast forward to start of amendment, to see our add_invoice_items get created for the additional item
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.month).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id
      end

      assert_equal(1, invoice_events.count)
      invoice_event = invoice_events.first
      assert_equal(TEST_DEFAULT_PRICE, invoice_event.data.object.amount)
    end

    it 'sync a backdated initial order billing annually skipping the initial invoice' do
      @user.field_defaults = {
        "customer" => {
          "email" => "test@test.com",
        },
        "subscription_schedule" => {
          "default_settings.invoice_settings.days_until_due" => "Net-45",
          "default_settings.collection_method" => "send_invoice", # note: you can only specify 'days_until_due' if invoice collection method is 'send_invoice'
        },
      }
      @user.save

      # initial order is one year contract with a monthly product
      # amendment order in three months adding an additional item
      contract_term = 12
      amendment_term = 8

      # initial order starts three months + 10 days ago
      initial_order_start_date = (now_time - 10.days) - 3.months
      initial_order_end_date = initial_order_start_date + contract_term.months

      # amendment order starts 4 months into contract
      amendment_start_date = initial_order_start_date + 4.month

      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create and sync a backdated initial order
      sf_order = create_subscription_order(sf_product_id: sf_product_id,
                                           contact_email: "syncs_backdated_initial_skip",
                                           additional_fields: {
                                             CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
                                             CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                           })

      # set SKIP_PAST_INITIAL_INVOICES to true
      sf.update!(SF_ORDER,
        SF_ID => sf_order.Id,
        prefixed_stripe_field(SKIP_PAST_INITIAL_INVOICES) => true
      )

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

      # get Stripe customer and advance the test clock
      stripe_customer = stripe_get(subscription_schedule.customer)
      refute_nil(stripe_customer.test_clock)
      advance_test_clock(stripe_customer, (now_time + 1.months).to_i)

      # confirm no intial invoice was created for the backdated period
      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)
      assert_equal(0, invoices.data.length)

      # create and sync an amendment order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # confirm the amendment order was processed correctly
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
      assert_equal(2, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start at the amendment start date
      # and have two products with total quantity of 2
      assert(second_phase.start_date.to_i, amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)
      # second phase should have two items with a total quantity of 2
      assert_equal(2, second_phase.items.count)
      second_phase_item_1 = T.must(second_phase.items.first)
      second_phase_item_2 = T.must(second_phase.items.second)
      assert_equal(1, second_phase_item_1.quantity)
      assert_equal(2, second_phase_item_2.quantity)

      # fast forward to start of amendment, to see our add_invoice_items get created for the additional item
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.month).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id
      end

      assert_equal(1, invoice_events.count)
      invoice_event = invoice_events.first
      StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)
      assert_equal(1, invoices.data.length)
      next_invoice = invoices.data.first
      assert_equal(2 * TEST_DEFAULT_PRICE * BigDecimal(amendment_term) / BigDecimal(contract_term), next_invoice.amount_due)
      assert_equal("true", next_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
      assert_equal(1, next_invoice.lines.count)
      invoice_line = next_invoice.lines.first
      assert_equal(2, invoice_line.quantity)
    end
  end
end
