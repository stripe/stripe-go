# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::ProratedAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'creates a new phase with a duration longer than the billing frequency, but not divisible' do
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

    # initial order: 2yr contract, one annual item
    # second order: +2 quantity, revising the existing item, 6-24mo

    yearly_price = 120_00
    contract_term = 24
    amendment_term = 18
    amendment_start_date = now_time + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months
    initial_start_date = now_time

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => nil,
        # CPQ_QUOTE_SUBSCRIPTION_TERM => 12,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity
    amendment_data["lineItems"].first["record"]["SBQQ__Quantity__c"] = 3

    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)
    sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # first phase should have one sub item and no proration items
    assert_equal(1, first_phase.items.count)
    assert_empty(first_phase.add_invoice_items)

    # second phase should have two items (one original, one addition)
    # it should have a single proration item
    assert_equal(2, second_phase.items.count)
    assert_equal(1, second_phase.add_invoice_items.count)

    first_phase_item = T.must(first_phase.items.first)
    second_phase_item = T.must(second_phase.items.detect {|i| i.quantity == 1 })
    second_phase_item_additive = T.must(second_phase.items.detect {|i| i.quantity == 2 })

    assert_equal(1, first_phase_item.quantity)
    assert_equal(1, second_phase_item.quantity)
    assert_equal(2, second_phase_item_additive.quantity)

    # metadata on the phase item representing the same order line should be identical
    # additionally, the price IDs should be the same since the pricing information did not change
    assert_equal(first_phase_item.metadata, second_phase_item.metadata)
    assert_equal(first_phase_item.price, second_phase_item.price)

    # the additive price order line should have different metadata, but the exact same pricing information
    refute_equal(second_phase_item_additive.metadata, first_phase_item.metadata)
    second_phase_item_additive_price = Stripe::Price.retrieve(T.cast(second_phase_item_additive.price, String), @user.stripe_credentials)
    first_phase_item_price = Stripe::Price.retrieve(T.cast(first_phase_item.price, String), @user.stripe_credentials)
    assert_equal(second_phase_item_additive_price.unit_amount_decimal, first_phase_item_price.unit_amount_decimal)
    assert_equal(second_phase_item_additive_price.type, first_phase_item_price.type)
    assert_equal(second_phase_item_additive_price.product, first_phase_item_price.product)

    # now, let's take a look at the prorated items!
    assert_equal(1, second_phase.add_invoice_items.count)
    prorated_item = T.unsafe(second_phase.add_invoice_items.first)
    assert_equal(2, prorated_item.quantity)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)

    # check additional fields added to the proration invoice item
    assert_equal("phase_end", prorated_item.period.end.type)
    assert_equal("phase_start", prorated_item.period.start.type)
    assert_equal("true", prorated_item.metadata[StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
    assert_equal(second_phase_item_additive.metadata['salesforce_order_item_id'], prorated_item.metadata['salesforce_order_item_id'])

    assert_equal('one_time', prorated_price.type)
    assert_equal(prorated_price.product, first_phase_item_price.product)
    assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
    assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
    assert_equal("true", prorated_price.metadata['salesforce_proration'])
    assert_equal(second_phase_item_additive_price.id, prorated_price.metadata['salesforce_original_stripe_price_id'])

    # since this is an 18mo prorated item we should only bill for 6mo since the rest will be billed by stripe
    assert_equal((yearly_price / (12 / 6)).to_s, prorated_price.unit_amount_decimal)

    # now, let's advance the clock and pretent like we are in the future to fully test autobilling
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)

    test_clock = Stripe::TestHelpers::TestClock.retrieve(stripe_customer.test_clock, @user.stripe_credentials)
    test_clock.advance(frozen_time: (amendment_start_date + 1.day).to_i)

    # test clocks can take some time...
    wait_until(timeout: 5.minutes) do
      test_clock.refresh
      test_clock.status == 'ready'
    end

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

    assert_equal(1, invoice.lines.count)
    invoice_line = invoice.lines.first
    assert_equal(2, invoice_line.quantity)
    assert_equal(60_00 * 2, invoice.total)
    assert_equal("true", invoice.metadata[StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
  end

  # NOTE this was the first test written and has more extensive edge cases than other amendment tests
  it 'creates a new phase with a duration shorter than the billing frequency' do
    # initial order: 2yr contract (two billing cycles), one annual item
    # second order: +2 quantity, revising the existing item, 18-24mo

    yearly_price = 120_00
    contract_term = 24
    amendment_start_date = now_time + 18.months
    amendment_term = 6
    amendment_end_date = (amendment_start_date + amendment_term.months)
    initial_start_date = now_time

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity
    amendment_data["lineItems"].first["record"]["SBQQ__Quantity__c"] = 3

    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)
    sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

    # api preconditions: no end date calculated on orders, end date IS calculated on the contract
    assert_nil(sf_order_amendment["EndDate"])
    # TODO understand why the contract end date is one day before Stripe, pretty certain this is due to differences in what "end" really is
    assert_equal(format_date_for_salesforce(amendment_end_date - 1.day), sf_order_amendment_contract['EndDate'])

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # first phase should have one sub item and no proration items
    assert_equal(1, first_phase.items.count)
    assert_empty(first_phase.add_invoice_items)

    # second phase should have two items (one original, one addition)
    # it should have a single proration item
    assert_equal(2, second_phase.items.count)
    assert_equal(1, second_phase.add_invoice_items.count)

    first_phase_item = T.must(first_phase.items.first)
    second_phase_item = T.must(second_phase.items.detect {|i| i.quantity == 1 })
    second_phase_item_additive = T.must(second_phase.items.detect {|i| i.quantity == 2 })

    assert_equal(1, first_phase_item.quantity)
    assert_equal(1, second_phase_item.quantity)
    assert_equal(2, second_phase_item_additive.quantity)

    # metadata on the phase item representing the same order line should be identical
    # additionally, the price IDs should be the same since the pricing information did not change
    assert_equal(first_phase_item.metadata, second_phase_item.metadata)
    assert_equal(first_phase_item.price, second_phase_item.price)

    # the additive price order line should have different metadata, but the exact same pricing information
    refute_equal(second_phase_item_additive.metadata, first_phase_item.metadata)
    second_phase_item_additive_price = Stripe::Price.retrieve(T.cast(second_phase_item_additive.price, String), @user.stripe_credentials)
    first_phase_item_price = Stripe::Price.retrieve(T.cast(first_phase_item.price, String), @user.stripe_credentials)
    assert_equal(second_phase_item_additive_price.unit_amount_decimal, first_phase_item_price.unit_amount_decimal)
    assert_equal(second_phase_item_additive_price.type, first_phase_item_price.type)
    assert_equal(second_phase_item_additive_price.product, first_phase_item_price.product)

    # let's make sure the pricing information is actually correct
    assert_equal(yearly_price.to_s, first_phase_item_price.unit_amount_decimal)
    assert_equal('recurring', first_phase_item_price.type)
    assert_equal('month', first_phase_item_price.recurring.interval)
    assert_equal(12, first_phase_item_price.recurring.interval_count)

    # the additive line price should be auto-archived, but not the first
    assert_nil(first_phase_item_price.metadata['salesforce_auto_archive'])
    assert(first_phase_item_price.active)
    assert_equal("true", second_phase_item_additive_price.metadata['salesforce_auto_archive'])
    refute(second_phase_item_additive_price.active)

    # now, let's take a look at the prorated items!
    assert_equal(1, second_phase.add_invoice_items.count)
    prorated_item = T.unsafe(second_phase.add_invoice_items.first)
    assert_equal(2, prorated_item.quantity)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)

    # check additional fields added to the proration invoice item
    assert_equal("phase_end", prorated_item.period.end.type)
    assert_equal("phase_start", prorated_item.period.start.type)
    assert_equal("true", prorated_item.metadata[StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
    assert_equal(second_phase_item_additive.metadata['salesforce_order_item_id'], prorated_item.metadata['salesforce_order_item_id'])

    assert_equal('one_time', prorated_price.type)
    assert_equal((yearly_price / (12 / amendment_term)).to_s, prorated_price.unit_amount_decimal)
    assert_equal(prorated_price.product, first_phase_item_price.product)
    assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
    assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
    assert_equal("true", prorated_price.metadata['salesforce_proration'])
    assert_equal(second_phase_item_additive_price.id, prorated_price.metadata['salesforce_original_stripe_price_id'])
  end

  # scenario where the start date is not on a billing cycle boundary but the subscription_term % billing_frequency == 0
  # TODO I am struggling to think of a situation where this could be true, might be impossible with the coterm requirement
  # initial order: 1yr contract, one quarterly item
  # second order: +1 quantity, revising the existing item, 6-24mo

  it 'maps metadata from the order line to the proration line item' do

  end

  it 'bills the incremental quantity at a new price' do

  end

  it 'errors on a non-cotermed invoice' do

  end

  it 'excludes metered billing, tiered pricing, and one-time invoice items from proration calculation' do

  end

  it 'creates a new phase without prorations when all items are new' do
    # if all line items are new
  end
end
