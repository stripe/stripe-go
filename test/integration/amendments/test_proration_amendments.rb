# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::ProratedAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
  end

  it 'creates a new phase with a duration longer than the billing frequency, but not divisible' do
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

    # initial order: 2yr contract, one annual item
    # second order: +2 quantity, revising the existing item, 6-24mo

    yearly_price = 120_00
    contract_term = 24
    initial_order_start_date = now_time
    initial_order_end_date = initial_order_start_date + contract_term.months

    amendment_term = 18
    amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months

    # normalize the amendment_end_date so test doesn't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => nil,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      contact_email: "new_phase_duration_longer_freq",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
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
    assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
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
    assert_equal(yearly_price.to_s, first_phase_item_price.unit_amount_decimal)
    assert_equal(second_phase_item_additive_price.unit_amount_decimal, first_phase_item_price.unit_amount_decimal)
    assert_equal(second_phase_item_additive_price.type, first_phase_item_price.type)
    assert_equal(second_phase_item_additive_price.product, first_phase_item_price.product)

    # now, let's take a look at the prorated items!
    assert_equal(1, second_phase.add_invoice_items.count)
    prorated_item = T.unsafe(second_phase.add_invoice_items.first)
    assert_equal(2, prorated_item.quantity)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)

    # check additional fields added to the proration invoice item
    assert_equal("subscription_period_end", prorated_item.period.end.type)
    assert_equal("phase_start", prorated_item.period.start.type)
    assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
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

    test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

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

    invoice_period_start = invoice_event.data.object.period.start
    invoice_period_end = invoice_event.data.object.period.end

    # sanity check the usage period of this proration invoice
    # should be from the amendment start to the next billing cycle date
    assert_equal(amendment_start_date.to_i, invoice_period_start)
    # the proration period end should align with the next billing cycle
    # since this is billed 'Annually', add a year to the initial order start date to get the start of the next billing cycle
    assert_equal((initial_order_start_date + 1.year).to_i, invoice_period_end)

    invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event))

    assert_equal(1, invoice.lines.count)
    invoice_line = invoice.lines.first
    assert_equal(2, invoice_line.quantity)
    assert_equal(60_00 * 2, invoice.total)
    assert_equal("true", invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
  end

  # NOTE this was the first test written and has more extensive edge cases than other amendment tests
  it 'creates a new phase with a duration shorter than the billing frequency' do
    # initial order: 2yr contract (two billing cycles), one annual item
    # second order: +2 quantity, revising the existing item, 18-24mo

    yearly_price = 120_00
    contract_term = 24

    initial_order_start_date = now_time
    initial_order_end_date = initial_order_start_date + contract_term
    amendment_start_date = now_time + 18.months
    amendment_term = 6
    amendment_end_date = amendment_start_date + amendment_term.months

    # normalize the amendment_end_date so test doesn't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      contact_email: "new_phase_duration_shorter_freq",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3

    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)
    sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

    # api preconditions: no end date calculated on orders, end date IS calculated on the contract
    assert_nil(sf_order_amendment["EndDate"])
    # NOTE is seems like the start and end date of the contract are NOT tied to the start date of the
    assert_equal(format_date_for_salesforce(initial_order_start_date), sf_order_amendment_contract['StartDate'])
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
    assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
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
    assert_equal("subscription_period_end", prorated_item.period.end.type)
    assert_equal("phase_start", prorated_item.period.start.type)
    assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
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

  it 'excludes tiered pricing, and one-time invoice items' do

  end

  it 'excludes metered billing from proration calculation' do
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true

      # initial order: 2yr contract (two billing cycles), silver metered item
      # amendment order: two months into contract, upgrade to gold metered item
      silver_tier_price = 1200_00
      gold_tier_price = 2400_00

      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months
      # normalize the amendment_end_date so test doesn't fail at EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_metered_silver_product_id, _sf_metered_silver_pricebook_id = salesforce_recurring_metered_produce_with_price(price_in_cents: silver_tier_price)
      sf_order = create_subscription_order(sf_product_id: sf_metered_silver_product_id, contact_email: "exclude_metered_billing")

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # create amendment to remove silver metered and upsell to gold metered product
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # remove silver item
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      # # add gold metered product
      sf_metered_gold_product_id, _sf_metered_gold_pricebook_id = salesforce_recurring_metered_produce_with_price(price_in_cents: gold_tier_price)
      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_metered_gold_product_id)
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(2, subscription_schedule.phases.count)

      # validate the first phase
      first_phase = T.must(subscription_schedule.phases.first)
      # first phase should have one item (silver metered item)
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      first_phase_item_price = Stripe::Price.retrieve(T.cast(first_phase_item.price, String), @user.stripe_credentials)
      assert_equal(silver_tier_price.to_s, first_phase_item_price.unit_amount_decimal)

      # validate the second phase
      second_phase = T.must(subscription_schedule.phases[1])
      # second phase should have two items (silver and gold metered items)
      # and no proration items
      assert_equal(1, second_phase.items.count)
      second_phase_item = T.must(second_phase.items.first)
      second_phase_item_price = Stripe::Price.retrieve(T.cast(second_phase_item.price, String), @user.stripe_credentials)
      assert_equal(gold_tier_price.to_s, second_phase_item_price.unit_amount_decimal)
      assert_equal('metered', second_phase_item_price.recurring.usage_type)

      assert_equal(0, second_phase.add_invoice_items.count)
  end

  # also known as a 'cross sell'
  # https://jira.corp.stripe.com/browse/PLATINT-1831
  it 'creates a new phase with prorations when all items are new' do
    # initial order: 2yr contract, one annual item
    # second order: additional product, 5-24mo

    yearly_price = 120_00
    yearly_price_2 = 150_00
    contract_term = TEST_DEFAULT_CONTRACT_TERM * 2
    amendment_term = 19
    amendment_start_date = now_time + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months
    initial_start_date = now_time

    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
      price: yearly_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => nil,
      }
    )

    sf_product_id_2, _sf_pricebook_id_2 = salesforce_recurring_product_with_price(
      price: yearly_price_2,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => nil,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      contact_email: "all_items_new_prorations",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
    sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

    amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id_2)

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

    # now, let's take a look at the prorated items!
    assert_equal(1, second_phase.add_invoice_items.count)
    prorated_item = T.unsafe(second_phase.add_invoice_items.first)
    assert_equal(1, prorated_item.quantity)
    prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)

    first_phase_item = T.must(first_phase.items.first)
    first_phase_item_price = Stripe::Price.retrieve(T.cast(first_phase_item.price, String), @user.stripe_credentials)
    second_phase_item = T.must(second_phase.items.detect {|i| i.price == first_phase_item.price })
    second_phase_item_additive = T.must(second_phase.items.detect {|i| i.price != first_phase_item.price })
    second_phase_item_additive_price = Stripe::Price.retrieve(T.cast(second_phase_item_additive.price, String), @user.stripe_credentials)

    # check additional fields added to the proration invoice item
    assert_equal("subscription_period_end", prorated_item.period.end.type)
    assert_equal("phase_start", prorated_item.period.start.type)
    assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
    assert_equal(second_phase_item_additive.metadata['salesforce_order_item_id'], prorated_item.metadata['salesforce_order_item_id'])

    assert_equal('one_time', prorated_price.type)
    assert_equal(prorated_price.product, second_phase_item_additive_price.product)
    assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
    assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
    assert_equal("true", prorated_price.metadata['salesforce_proration'])
    assert_equal(second_phase_item_additive_price.id, prorated_price.metadata['salesforce_original_stripe_price_id'])

    # since this is an 19mo prorated item we should only bill for 7mo since the rest will be billed by stripe
    assert_equal((yearly_price_2 / 12 * 7).to_s, prorated_price.unit_amount_decimal)
  end

  describe 'supports negative order items' do
    # When we ecounter terminated items, we will create a negative line item in Stripe for their associated credit.
    it 'supports an upsell, with credit for unused time' do
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 2yr contract (two billing cycles), silver tier (one item)
      # second order: upgrade to platinum tier (one item) two months into contract
      #                 - cancels remaining silver tier (issuing credit via negative line item for $1000)
      #                 - creates an invoice item for gold tier for remaining time ($2000)
      #                 - resulting invoice will have a total of $1000, ie $2000 prorated amount - $1000 credit

      yearly_silver_tier_price = 1200_00
      yearly_gold_tier_price = 2400_00
      contract_term = 24
      amendment_term = 22

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, _sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_gold_tier_product_id, _sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_gold_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "supports_upsell",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_gold_tier_product_id)
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      create_contract_from_order(sf_order_amendment)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(1200_00, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling

      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(2, invoice_events.count)
      proration_invoice_event = invoice_events[1]

      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))

      assert_equal(2, proration_invoice.lines.count)
      assert_equal(1000_00, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    it 'supports an upsell for a single year contract' do
      # Same as above test, but a single year contract (ie one billing cycle)
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 1yr contract (one billing cycle), silver tier (one item)
      # second order: upgrade to gold tier (one item) two months into contract
      #                 - cancels remaining silver tier (issuing credit via negative line item for $1000)
      #                 - creates an invoice item for gold tier for remaining time ($2000)
      #                 - resulting invoice will have a total of $1000, ie $2000 prorated amount - $1000 credit

      yearly_silver_tier_price = 1200_00
      yearly_gold_tier_price = 2400_00
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_gold_tier_product_id, sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_gold_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "supports_upsell_year",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_gold_tier_product_id)
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # Get SubscriptionSchedule Id
      sf_order.refresh
      stripe_subscription_schedule_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data

      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(1200_00, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling

      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(2, invoice_events.count)
      proration_invoice_event = invoice_events[1]
      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))

      assert_equal(2, proration_invoice.lines.count)
      assert_equal(1000_00, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    it 'supports an upsell for a two year contract, in the second year' do
      # Same as first negative invoice item test, but the amendment is in the second year
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 2yr contract (two billing cycles), silver tier (one item)
      # second order: upgrade to platinum tier (one item) 18 months into contract (six months into second billing cycle)
      #                 - cancels remaining silver tier (issuing credit via negative line item for $600)
      #                 - creates an invoice item for platinum tier for remaining time ($2500)
      #                 - resulting invoice will have a total of $1900, ie $2500 prorated amount - $600 credit

      yearly_silver_tier_price = 1200_00
      yearly_platinum_tier_price = 5000_00
      contract_term = 24
      amendment_term = 6

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_platinum_tier_product_id, sf_platinum_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_platinum_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "supports_upsell_two_year",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0

      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_platinum_tier_product_id)
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      Timecop.freeze(Time.now + 1.minute)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data

      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(1200_00, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling

      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(2, invoice_events.count)
      proration_invoice_event = invoice_events[1]

      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))

      assert_equal(2, proration_invoice.lines.count)
      assert_equal(1900_00, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    # https://jira.corp.stripe.com/browse/PLATINT-2091
    it 'fully terminates an order, without issuing credit' do
      # Does not create a credit for a fully terminated order
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 1yr contract (one billing cycle), silver tier (one item)
      # second order: cancel silver tier (issuing credit via negative line item for $1000)

      yearly_silver_tier_price = 1200_00
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, _sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "full_terminate",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0

      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      Timecop.freeze(Time.now + 1.minute)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data

      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(1200_00, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(0, invoice_events.count)
    end

    it 'supports an upsell for a single year contract with day prorations enabled' do
      # Single year contract (ie one billing cycle)
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 1yr contract (one billing cycle), silver tier (one item)
      # second order: upgrade to gold tier (one item) two months into contract
      #                 - cancels remaining silver tier (issuing credit via negative line item)
      #                 - creates an invoice item for gold tier for remaining time

      yearly_silver_tier_price = 398_45
      contract_term = 12
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      yearly_gold_tier_price = 9961_19
      amendment_term = 9
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term - 1).months + 10.days
      amendment_end_date = initial_order_end_date

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      # create the two products
      billing_frequency_serialized = CPQBillingFrequencyOptions::ANNUAL.serialize
      sf_silver_tier_product_id, _sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_gold_tier_product_id, _sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_gold_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "supports_upsell_year_day_prorations_18",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)

      # cancel the silver tier product
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      # add the gold tier product
      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_gold_tier_product_id)
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      create_contract_from_order(sf_order_amendment)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # Get Subscription Schedule Id
      sf_order.refresh
      stripe_subscription_schedule_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)

      assert_equal(2, stripe_subscription_schedule.phases.count)

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)
      initial_invoice = invoices.first
      assert_equal(yearly_silver_tier_price, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(2, invoice_events.count)
      proration_invoice_event = invoice_events[1]
      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))
      assert_equal(2, proration_invoice.lines.count)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])

      # first calculate the prorated price of the gold product
      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)

      # then calculate the prorated terminated order item product
      prorated_credit_silver = yearly_silver_tier_price * proration_percentage * -1
      assert_equal(prorated_credit_silver.round, proration_invoice.lines.data.first.amount)
      assert_equal("true", proration_invoice.lines.data.first.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::CREDIT)])
      assert_equal("true", proration_invoice.lines.data.first.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      prorated_amount_gold = yearly_gold_tier_price * proration_percentage
      assert_equal(prorated_amount_gold.round, proration_invoice.lines.data.second.amount)
      assert_equal("true", proration_invoice.lines.data.second.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # verify proration invoice total
      assert_equal(prorated_amount_gold.round + prorated_credit_silver.round, proration_invoice.total)

      # sanity check this proration calc aligns with what is in Salesforce CPQ
      # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      prorated_invoice_amount_in_dollars = BigDecimal(proration_invoice.total) / 100
      assert_equal(sf_order_amendment["TotalAmount"], prorated_invoice_amount_in_dollars)
    end

    it 'terminates one of an orders items without adding a new one' do
      @user.field_defaults['price'] = {
        'metadata.field_custom' => 'value_custom',
      }
      @user.save

      # ie, issues a prorated credit without a prorated debit
      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: 1yr contract (one billing cycle), silver tier (one item), gold tier (one item)
      # second order: remove silver tier (one item) two months into contract
      #                 - cancels remaining silver tier (issuing credit via negative line item for $1000)
      #                 - resulting invoice will have a total of -$1000, ie $1000 prorated credit

      yearly_silver_tier_price = 1200_00
      yearly_gold_tier_price = 2400_00
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_gold_tier_product_id, sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_gold_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_account_id = create_salesforce_account

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "terminate_item",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # add our two products
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_silver_tier_product_id)
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_gold_tier_product_id)
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0

      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      assert_equal(3600_00, initial_invoice.amount_due)
      assert_equal(2, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling

      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(1, invoice_events.count)
      proration_credit_invoice_event = invoice_events.first

      assert_equal("true", proration_credit_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::CREDIT)])
      assert_equal("true", proration_credit_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_credit_invoice_event))
      assert_equal(1, proration_invoice.lines.count)
      assert_equal(-1000_00, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
      assert_equal("value_custom", proration_invoice.lines.first.price.metadata["field_custom"])
    end

    it 'supports credits on monthly billing cycles' do
      # TODO
      pass
    end

    it 'does not issue credit for merchant not gated into feature' do
      # initial order: 2yr contract (two billing cycles), silver tier (one item)
      # second order: upgrade to platinum tier (one item) two months into contract
      #                 - cancels remaining silver tier
      #                 - creates an invoice item for gold tier for remaining time ($2000)
      #                 - resulting invoice will have a total of $2000, ie $2000 prorated amount with no credit

      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      yearly_silver_tier_price = 1200_00
      yearly_gold_tier_price = 2400_00
      contract_term = 24
      amendment_term = 22

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_silver_tier_product_id, sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_silver_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_gold_tier_product_id, sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: yearly_gold_tier_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_salesforce_order(
        sf_product_id: sf_silver_tier_product_id,
        contact_email: "no_issue_credit",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # cancel silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0

      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_gold_tier_product_id)

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # Get SubscriptionSchedule Id
      sf_order.refresh
      stripe_subscription_schedule_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data

      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(1200_00, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling

      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(1, invoice_events.count)
      proration_invoice_event = invoice_events.first

      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))

      assert_equal(1, proration_invoice.lines.count)
      assert_equal(2000_00, proration_invoice.total)
      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end

    it 'terminate an item that had quantity > 1' do
      # initial order: 24 quantity gold tier ($24,000 / year), billed anually on 1 year contract
      # ammendment order: remove all 24 of gold tier and add a single silver from month 2 onwards
      #                   - Credit for *$1000 * 24) * (11 / 12) = $22,000
      #                   - Debit for $120 * (11 / 12) = $110
      #                   - Invoice total of -$21,890

      @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      contract_term = 12
      amendment_term = 11

      gold_quantity = 24
      gold_unit_price = 1000_00
      gold_total = gold_unit_price * gold_quantity

      silver_unit_price = 120_00

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term
      amendment_start_date = now_time + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      sf_account_id = create_salesforce_account

      # normalize the amendment_end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_gold_tier_product_id, sf_gold_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: gold_unit_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_silver_tier_product_id, sf_silver_tier_pricebook_id = salesforce_recurring_product_with_price(
        price: silver_unit_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "term_item",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # set first product quantity
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_gold_tier_product_id)
      quote_with_product["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = gold_quantity
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # remove gold and add silver tier
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0

      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_silver_tier_product_id)

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

      # Translate
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # Get SubscriptionSchedule Id
      sf_order.refresh
      stripe_subscription_schedule_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)

      # Get Customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # Pay off initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data

      assert_equal(1, invoices.length)

      initial_invoice = invoices.first

      assert_equal(gold_total, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)

      initial_invoice.pay({'paid_out_of_band': true})

      # now, let's advance the clock and pretent like we are in the future to fully test autobilling
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end

      assert_equal(2, invoice_events.count)
      proration_invoice_event = invoice_events[1]

      assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      proration_invoice = T.must(StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, proration_invoice_event))

      assert_equal(2, proration_invoice.lines.count)

      proration_multiplier = 11.to_f / 12
      gold_prorated_credit = (gold_unit_price * gold_quantity) * proration_multiplier * -1
      silver_prorated_debit = silver_unit_price * proration_multiplier

      assert_equal(gold_prorated_credit, proration_invoice.lines.first.amount)
      assert_equal(silver_prorated_debit + gold_prorated_credit, proration_invoice.total)

      assert_equal("true", proration_invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
    end
  end

  describe 'non_anniversary_amendments' do
    it 'translates non-anniversary amendment order billed monthly' do
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed monthly
      # amendment order: starts 3 days later (non-anniversary day)
      contract_term = 12
      amendment_term = 11

      initial_order_start_date = now_time
      _initial_order_end_date = initial_order_start_date + contract_term.months

      # amendment order starts on non-anniversary day
      amendment_start_date = initial_order_start_date + 3.days

      # create the initial sf order
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "non_anniversary_monthly",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)
      initial_invoice = invoices.first
      assert_equal(TEST_DEFAULT_PRICE, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create and sync the sf order amendment to increase +1
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate amendment order and confirm no errors raised
      exception = assert_raises(Integrations::Errors::TranslatorError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("Calculated price multiplier differs from cpq price multiplier.", exception.message)

      # Note: the below is commented out since this will fail CI since the below is for when CPQ Subscription Prorate equals 'Month'
      # sf_order_amendment.refresh

      # fetch the subscription schedule
      # stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      # stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      # assert_equal(2, stripe_subscription_schedule.phases.count)

      # # verify phase dates
      # first_phase = T.must(stripe_subscription_schedule.phases.first)
      # assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      # assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # # first phase should have one sub item and no proration items
      # assert_equal(1, first_phase.items.count)
      # assert_empty(first_phase.add_invoice_items)

      # second_phase = T.must(stripe_subscription_schedule.phases[1])
      # assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      # assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # # second phase should have two items (one original, one addition)
      # # it should have a single proration item
      # assert_equal(2, second_phase.items.count)
      # assert_equal(1, second_phase.add_invoice_items.count)

      # prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      # assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
      # assert_equal(1, prorated_item.quantity)

      # # check additional fields added to the proration invoice item
      # assert_equal("phase_start", prorated_item.period.start.type)
      # assert_equal("subscription_period_end", prorated_item.period.end.type)
      # assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      # assert_equal('one_time', prorated_price.type)
      # assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      # assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      # assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # now let's advance the test clock
      # test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # # simulate sending the webhook
      # events = Stripe::Event.list({
      #   type: 'invoiceitem.created',
      # }, @user.stripe_credentials)

      # invoice_events = events.data.select do |event|
      #   event_object = event.data.object
      #   event_object.test_clock == test_clock.id && event.request.id.nil?
      # end
      # assert_equal(1, invoice_events.count)

      # # verify proration invoice total
      # # partial month equal a whole month for proration calculation
      # proration_invoice_event = invoice_events[0]
      # prorated_invoice_amount = proration_invoice_event.data.object.amount
      # assert_equal(TEST_DEFAULT_PRICE, prorated_invoice_amount)

      # # sanity check this proration calc aligns with what is in Salesforce CPQ
      # # the Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Prorated Invoice Amount
      # price_in_dollars = TEST_DEFAULT_PRICE / 100
      # total_list_price = price_in_dollars * contract_term
      # previously_billed_amount = price_in_dollars * (contract_term - amendment_term)
      # prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      # cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      # assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end

    it 'translates non-anniversary amendment order billed annually' do
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed annually
      # amendment order: starts 1 months + 3 days later on a non-anniversary day
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months - (1.month - 3.days)

      # create the initial sf order
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "non_anniversary_annually",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)
      initial_invoice = invoices.first
      assert_equal(TEST_DEFAULT_PRICE, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create and sync the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate order
      exception = assert_raises(Integrations::Errors::TranslatorError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("Calculated price multiplier differs from cpq price multiplier.", exception.message)

      # Note: the below is commented out since this will fail CI since the below is for when CPQ Subscription Prorate equals 'Month'

      # sf_order_amendment.refresh

      # # get customer
      # sf_account = sf_get(sf_order['AccountId'])
      # stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      # stripe_customer = stripe_get(stripe_customer_id)
      # refute_nil(stripe_customer.test_clock)

      # invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      # assert_equal(1, invoices.length)

      # now let's advance the clock
      # test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.day).to_i)

      # # simulate sending the webhook
      # events = Stripe::Event.list({
      #   type: 'invoiceitem.created',
      # }, @user.stripe_credentials)

      # invoice_events = events.data.select do |event|
      #   event_object = event.data.object
      #   event_object.test_clock == test_clock.id && event.request.id.nil?
      # end
      # assert_equal(1, invoice_events.count)

      # # verify this is a proration invoice
      # proration_invoice_event = invoice_events[0]
      # assert_equal("true", proration_invoice_event.data.object.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      # # verify proration invoice total
      # prorated_invoice_amount = proration_invoice_event.data.object.amount
      # assert_equal(TEST_DEFAULT_PRICE * BigDecimal(amendment_term + 1) / contract_term, prorated_invoice_amount)

      # # sanity check this proration calc aligns with what is in Salesforce CPQ
      # # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      # price_in_dollars = TEST_DEFAULT_PRICE / 100
      # total_list_price = price_in_dollars
      # previously_billed_amount = total_list_price
      # prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      # cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      # assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end
  end

  describe 'non-anniversary amendments with day porations' do
    it 'translates non-anniversary amendment order billed monthly with day proration enabled' do
      # Note either feature DAY_PRORATIONS should only be enabled or connector settings needs to be set to 'month_day"
      # for connector to use PQ Subscription Prorate Precision is set to 'Month + Day'
      @user.disable_feature FeatureFlags::DAY_PRORATIONS, update: true
      @user.connector_settings[CONNECTOR_SETTING_CPQ_PRORATE_PRECISION] = 'month+day'
      @user.save

      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed monthly
      # amendment order: starts 1 month and 4 days later on a non-anniversary day
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time # march 22
      initial_order_end_date = initial_order_start_date + contract_term.months

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + 1.months + 22.days

      # create the initial sf order
      price = 377.63
      billing_frequency_serialized = CPQBillingFrequencyOptions::MONTHLY.serialize
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "non_anniversary_day_proration_6",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)
      initial_invoice = invoices.first
      assert_equal(price, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create and sync the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # verify phase dates
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (one original, one addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      # now let's advance the test clock
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end
      assert_equal(1, invoice_events.count)

      # verify proration invoice total
      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: 1)
      prorated_amount = price * proration_percentage
      prorated_invoice_amount = invoice_events[0].data.object.amount
      assert_equal(prorated_amount.round, prorated_invoice_amount)

      # sanity check this proration calc aligns with what is in Salesforce CPQ
      # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      price_in_dollars = price / 100
      total_list_price = price_in_dollars * contract_term
      previously_billed_amount = price_in_dollars * (contract_term - amendment_term)
      prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end

    it 'translates non-anniversary amendment order billed monthly with day proration enabled and salesforce precision' do
      @user.connector_settings[CONNECTOR_SETTING_CPQ_PRORATE_PRECISION] = 'month+day'
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
      @user.enable_feature FeatureFlags::SALESFORCE_PRECISION, update: true
      @user.save

      # initial order: starts now, billed annually
      # amendment order: starts 1 month and 4 days later on a non-anniversary day
      contract_term = 12
      amendment_term = 10

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + 1.months + 12.days

      # create the initial sf order
      salesforce_price = 377_63
      billing_frequency_serialized = CPQBillingFrequencyOptions::ANNUAL.serialize
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: salesforce_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "non_anniversary_day_proration_salesforce_precision_18",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)
      initial_invoice = invoices.first
      assert_equal(salesforce_price, initial_invoice.amount_due)
      assert_equal(1, initial_invoice.lines.data.length)
      initial_invoice.pay({'paid_out_of_band': true})

      # create and sync the sf order amendment to increase quantity by 49 (total 50)
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 50
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # verify phase dates
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (one original, one addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      # verify the unit amount on the phase has salesforce precision
      sf_order_items = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      sf_unit_price = sf_order_items.first["ListPrice"]
      phase_item_price = Stripe::Price.retrieve(T.must(second_phase.items.second)["price"], @user.stripe_credentials)
      assert_equal(sf_unit_price, phase_item_price["unit_amount_decimal"].to_d / 100)

      # now let's advance the test clock
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end
      assert_equal(1, invoice_events.count)

      # verify proration invoice total
      prorated_invoice_amount = invoice_events[0].data.object.amount
      assert_equal(sf_order_amendment["TotalAmount"], BigDecimal(prorated_invoice_amount) / 100)

      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: TEST_DEFAULT_CONTRACT_TERM)
      calculated_prorated_amount = (BigDecimal(salesforce_price) * proration_percentage / 100).round(2) * 100 * 49
      assert_equal(calculated_prorated_amount.round, prorated_invoice_amount)
    end

    it 'translates non-anniversary amendment order billed annually with day prorations enabled' do
      # Note feature DAY_PRORATIONS should only be enabled if CPQ Subscription Prorate Precision is set to 'Month + Day'
      @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      cache_service = CacheService.new(@user)
      mapper = StripeForce::Mapper.new(@user, cache_service)

      # initial order: starts now, billed annually
      # amendment order: starts 4 days later (on a non-anniversary day)
      contract_term = 12

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + 4.days

      # create the initial sf order
      billing_frequency_serialized = CPQBillingFrequencyOptions::ANNUAL.serialize
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "non_anniversary_annually_day_proration",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # create and sync the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      # note this is on purpose to test that non anniversary amendment work without a subscription term
      assert_nil(amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM])

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # verify phase dates
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (one original, one addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      # now let's advance the test clock
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end
      assert_equal(1, invoice_events.count)

      amendment_term = StripeForce::Utilities::SalesforceUtil.calculate_subscription_term(mapper, sf_order_amendment)
      assert_equal(11, amendment_term)
      # verify proration invoice total
      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      prorated_amount = TEST_DEFAULT_PRICE * proration_percentage
      prorated_invoice_amount = invoice_events[0].data.object.amount
      assert_equal(prorated_amount.round, prorated_invoice_amount)

      # sanity check this proration calc aligns with what is in Salesforce CPQ
      # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      price_in_dollars = TEST_DEFAULT_PRICE / 100
      total_list_price = price_in_dollars
      previously_billed_amount = price_in_dollars
      prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end

    it 'translates non-anniversary amendment order billed semi-annually with day prorations enabled' do
      # Note feature DAY_PRORATIONS should only be enabled if CPQ Subscription Prorate Precision is set to 'Month + Day'
      @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed semi-annually
      # amendment order: starts 20 days later on a non-anniversary day
      contract_term = 24
      amendment_term = 23

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + 20.days

      # create the initial sf order
      annual_price = 100_00
      billing_frequency_serialized = CPQBillingFrequencyOptions::SEMIANNUAL.serialize
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: annual_price,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "semi_annual_day_proration_2",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,

        }
      )

      # create and sync the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # now let's advance the test clock
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end
      assert_equal(1, invoice_events.count)

      # verify proration invoice total
      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      prorated_amount = annual_price * proration_percentage
      prorated_invoice_amount = invoice_events[0].data.object.amount
      assert_equal(prorated_amount.round, prorated_invoice_amount)

      # sanity check this proration calc aligns with what is in Salesforce CPQ
      # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      price_in_dollars = annual_price / 100
      total_list_price = price_in_dollars * (contract_term / CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      previously_billed_amount = price_in_dollars * BigDecimal(billing_frequency) / BigDecimal(CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end

    it 'translates non-anniversary amendment order billed quarterly with day prorations enabled' do
      # Note feature DAY_PRORATIONS should only be enabled if CPQ Subscription Prorate Precision is set to 'Month + Day'
      @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
      @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed quarterly
      # amendment order: starts 1 month + 3 days later on a non-anniversary day
      contract_term = 24
      amendment_term = 22

      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      # intentially start amendment order on non-anniversary day
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term - 1).months + 3.day

      # create the initial sf order
      billing_frequency_serialized = CPQBillingFrequencyOptions::QUARTERLY.serialize
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        price: TEST_DEFAULT_PRICE,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => billing_frequency_serialized,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        contact_email: "quarterly_day_proration",
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # create and sync the sf order amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # fetch the subscription schedule
      stripe_subscription_schedule_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(2, stripe_subscription_schedule.phases.count)

      # verify phase dates
      first_phase = T.must(stripe_subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # first phase should have one sub item and no proration items
      assert_equal(1, first_phase.items.count)
      assert_empty(first_phase.add_invoice_items)

      second_phase = T.must(stripe_subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - initial_order_end_date.to_i)

      # second phase should have two items (one original, one addition)
      # it should have a single proration item
      assert_equal(2, second_phase.items.count)
      assert_equal(1, second_phase.add_invoice_items.count)

      # sanity check regular line item has quarterly interval
      recurring_price = Stripe::Price.retrieve(T.cast(T.must(second_phase.items.first).price, String), @user.stripe_credentials)
      assert_equal('recurring', recurring_price.type)
      assert_equal(3, recurring_price['recurring']['interval_count'])

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
      assert_equal(1, prorated_item.quantity)

      # check additional fields added to the proration invoice item
      assert_equal("phase_start", prorated_item.period.start.type)
      assert_equal("subscription_period_end", prorated_item.period.end.type)
      assert_equal("true", prorated_item.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # get customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # now let's advance the test clock
      test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 2.day).to_i)

      # simulate sending the webhook
      events = Stripe::Event.list({
        type: 'invoiceitem.created',
      }, @user.stripe_credentials)

      invoice_events = events.data.select do |event|
        event_object = event.data.object
        event_object.test_clock == test_clock.id && event.request.id.nil?
      end
      assert_equal(1, invoice_events.count)

      # verify proration invoice total
      days_prorating = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: amendment_start_date,
        sf_order_end_date: initial_order_end_date,
        sf_order_subscription_term: amendment_term)
      billing_frequency = StripeForce::Translate::PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(billing_frequency_serialized)
      proration_percentage = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: (amendment_term % billing_frequency), partial_month_days: days_prorating, product_subscription_term: CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      prorated_amount = TEST_DEFAULT_PRICE * proration_percentage
      prorated_invoice_amount = invoice_events[0].data.object.amount
      assert_equal(prorated_amount.round, prorated_invoice_amount)

      # sanity check this proration calc aligns with what is in Salesforce CPQ
      # Salesforce Prorated List Price = List Price - Already Billed Amount + CPQ Proration Amount
      price_in_dollars = TEST_DEFAULT_PRICE / 100
      total_list_price = price_in_dollars * (contract_term / CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM)
      previously_billed_amount = price_in_dollars * (BigDecimal(billing_frequency) / BigDecimal(CPQ_QUOTE_LINE_DEFAULT_SUBSCRIPTION_TERM))
      prorated_invoice_amount_in_dollars = BigDecimal(prorated_invoice_amount) / 100

      cpq_prorated_list_price = total_list_price - previously_billed_amount + prorated_invoice_amount_in_dollars
      assert_equal(sf_order_amendment["TotalAmount"], cpq_prorated_list_price)
    end
  end
end
