# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::OrderAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
    # note: enabling non_anniversary amendments should not affect anniversary amendments
    # therefore enable for this for the entire test suite
    @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
    @user.enable_feature FeatureFlags::BACKDATED_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
  end

  it 'creates a new phase from an order amendment with monthly billed products' do
    # initial order: 1yr contract, monthly billed
    # amendment: starts month 9, lasts 3 months, adds quantity 2
    monthly_price = 10_00
    contract_term = TEST_DEFAULT_CONTRACT_TERM
    initial_start_date = now_time
    initial_order_end_date = initial_start_date + contract_term

    amendment_term = 3
    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months
    # normalize the amendment_end_date so test doesn't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(price: monthly_price)
    sf_order = create_subscription_order(sf_product_id: sf_product_id, additional_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
    })
    sf_contract = create_contract_from_order(sf_order)

    # although the associated order is contracted, this does not
    # contract the associated opportunity. A user can opt to contract
    # through the opportunity, but we require that they contract the order
    sf_opportunity = sf_get(sf_contract.SBQQ__Opportunity__c)
    refute(sf_opportunity['SBQQ__Contracted__c'])

    # api precondition: initial orders have a nil contract ID
    sf_order.refresh
    assert_nil(sf_order.ContractId)

    # the contract should reference the initial order that was created
    assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

    # increase quantity by 2
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
    sf_order_amendment = create_order_from_quote_data(amendment_data)

    # api preconditions: the contract ID on the amendment is STILL empty right after the contract is created
    assert_nil(sf_order_amendment.ContractId)
    refute(sf_order_amendment[SF_ORDER_CONTRACTED])

    # https://help.salesforce.com/s/articleView?id=cpq_amend_contracts.htm&type=5&language=en_US
    # "The amendment opportunity updates its Amended Contract and Primary Quote fields with links to the contract you amended and your original quote."
    sf_amendment_opportunity = sf_get(sf_order_amendment.OpportunityId)
    assert_equal(sf_amendment_opportunity.SBQQ__AmendedContract__c, sf_contract.Id)

    # confirmed with users that they contract order amendments
    sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)
    # contracting the order amendment links it to the same contract object
    assert_equal(sf_order_amendment_contract.Id, sf_contract.Id)
    # the quote reference seems to stay the same
    assert_equal(sf_order_amendment_contract[SF_CONTRACT_QUOTE_ID], sf_contract[SF_CONTRACT_QUOTE_ID])
    # however, the order reference on the contract is updated as well!
    refute_equal(sf_order_amendment_contract[SF_CONTRACT_ORDER_ID], sf_contract[SF_CONTRACT_ORDER_ID])

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with a quantity of one
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_equal(1, first_phase_item.quantity)

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # second phase should have a second item with a quantity of 3
    # the order line in SF will have 2 (added 2 to have a net quantity of 3)
    assert_equal(2, second_phase.items.count)
    assert_equal(0, second_phase.add_invoice_items.count)
    second_phase_item_1 = T.must(second_phase.items.first)
    second_phase_item_2 = T.must(second_phase.items[1])
    assert_equal(1, second_phase_item_1.quantity)
    assert_equal(2, second_phase_item_2.quantity)

    # prices should be the same, but the order line reference is different
    # NOTE if proration is set to day, the pricing will be customized and this check will fail
    assert_equal(first_phase_item.price, second_phase_item_1.price)
    refute_equal(first_phase_item.price, second_phase_item_2.price)
    assert_equal(first_phase_item.metadata['salesforce_order_item_id'], second_phase_item_1.metadata['salesforce_order_item_id'])
    refute_equal(first_phase_item.metadata['salesforce_order_item_id'], second_phase_item_2.metadata['salesforce_order_item_id'])

    # let's be paranoid here and make sure the price is exactly what we expect
    phase_price = Stripe::Price.retrieve(T.cast(first_phase_item.price, String), @user.stripe_credentials)
    assert_equal(monthly_price.to_s, phase_price.unit_amount_decimal)
    assert_equal('recurring', phase_price.type)
    assert(phase_price.active)

    # enforces none of the duplicate/auto-archive/etc metadata is in place
    phase_price_metadata = phase_price.metadata.to_hash
    phase_price_metadata.delete(:salesforce_pricebook_entry_id)
    phase_price_metadata.delete(:salesforce_pricebook_entry_link)
    assert_empty(phase_price_metadata)
  end

  # usage products do NOT have a quantity in Stripe, which introduces additional complexity
  it 'creates a new phase from an order amendment removing a metered billing product' do
    # initial order: one metered
    # amendment: remove metered item, add non-metered

    amendment_term = 6
    initial_order_start_date = now_time
    initial_order_end_date = initial_order_start_date + TEST_DEFAULT_CONTRACT_TERM
    amendment_start_date = initial_order_start_date + 6.months
    amendment_end_date = amendment_start_date + amendment_term.months
    # normalize the amendment_end_date so tests don't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_metered_product_id, _sf_metered_pricebook_id = salesforce_recurring_metered_produce_with_price
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

    sf_order = create_subscription_order(sf_product_id: sf_metered_product_id)
    sf_contract = create_contract_from_order(sf_order)

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # remove metered billing item completely
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
    sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

    # add non-metered product
    amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with no quantity, since it is a metered product
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_nil(first_phase_item[:quantity])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
    assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # second phase should have a second item with a quantity of 1
    assert_equal(1, second_phase.items.count)
    assert_equal(0, second_phase.add_invoice_items.count)
    subscription_item = T.must(second_phase.items.detect {|i| !i[:quantity].nil? })
    assert_equal(1, subscription_item.quantity)
  end

  it 'does not treat a order amendment as a new order if it has not been contracted' do
    # if an amendment has not been contracted we should throw an error
  end

  it 'creates a new phase from an order amendment adding a non-metered product to a metered product' do
    # initial order: one metered
    # amendment: keep metered item, add non-metered

    contract_term = TEST_DEFAULT_CONTRACT_TERM
    amendment_term = 6
    initial_start_date = now_time
    initial_order_end_date = initial_start_date + contract_term
    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months

    # normalize the amendment_end_date so test doesn't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_metered_product_id, _sf_metered_pricebook_id = salesforce_recurring_metered_produce_with_price
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

    sf_order = create_subscription_order(sf_product_id: sf_metered_product_id)
    sf_contract = create_contract_from_order(sf_order)

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
    sf_quote_id = calculate_and_save_cpq_quote(amendment_data)

    # add non-metered product
    amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with no quantity, since it is a metered product
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_nil(first_phase_item[:quantity])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
    assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

    # second phase should have a second item with a quantity of 2
    assert_equal(2, second_phase.items.count)
    assert_equal(0, second_phase.add_invoice_items.count)
    subscription_item = T.must(second_phase.items.detect {|i| !i[:quantity].nil? })
    metered_item = T.must(second_phase.items.detect {|i| i[:quantity].nil? })
    assert_equal(1, subscription_item.quantity)
  end

  it 'creates a new phase with a duration equal to billing frequency' do
    start_date = now_time + 12.months
    amendment_term = 12
    end_date = start_date + amendment_term.months
    initial_start_date = now_time

    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        # 2yr term, two billing cycles
        CPQ_QUOTE_SUBSCRIPTION_TERM => 24.0,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
    sf_order_amendment = create_order_from_quote_data(amendment_data)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(2, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should start now and end in 9mo
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - start_date.to_i)

    # second phase should start at the end date
    assert_equal(0, second_phase.start_date - start_date.to_i)
    assert_equal(0, second_phase.end_date - end_date.to_i)

    assert_equal(1, first_phase.items.count)
    assert_equal(2, second_phase.items.count)

    first_phase_item = T.must(first_phase.items.first)
    second_phase_item = T.must(second_phase.items.detect {|i| i.quantity == 2 })
    second_phase_item_single_quantity = T.must(second_phase.items.detect {|i| i.quantity == 1 })

    assert_equal(1, first_phase_item.quantity)
    assert_equal(1, second_phase_item_single_quantity.quantity)
    assert_equal(2, second_phase_item.quantity)

    refute_equal(first_phase_item.price, second_phase_item.price)

    # price should be identical for the initial product
    assert_equal(first_phase_item.price, second_phase_item_single_quantity.price)
  end

  # this can occur if contracts are created async and the apex jobs are backed up
  # in this case, we'll just process the initial order we have without pulling the contract
  it 'does not fail if a contract does not yet exist for an order' do
    sf_order = create_subscription_order
    sf_contract = create_contract_from_order(sf_order)

    sf.destroy(SF_CONTRACT, sf_contract.Id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)
  end

  it 'supports adding one-time line items on a order amendment'
  it 'supports adding multiple one-time items of the pricebook id to an order amendment'

  # https://jira.corp.stripe.com/browse/PLATINT-1809
  describe 'subscription without a billing cycle' do
    it 'supports amending an order when the start date is in the future' do
      # initial order: starts in 1 month, standard product
      # amendment: starts in 9 months, qty +1

      contract_term = TEST_DEFAULT_CONTRACT_TERM
      amendment_term = 3
      initial_order_start_date = now_time + 1.month
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
        }
      )

      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # we will skip the upcoming invoice API call when this value is nil
      assert_nil(subscription_schedule.subscription)

      # unit test the upcoming calculation
      billing_cycle_dates = StripeForce::Translate::OrderAmendment.calculate_billing_cycle_dates(@user, subscription_schedule, 12)
      # they should all be in the future
      assert(billing_cycle_dates.all? {|ts| ts >= now_time.to_i })
      # first cycle should be midnight of the start date of the subscription
      assert_equal(initial_order_start_date.to_i, billing_cycle_dates.first)
      # there should be two billing dates: one right now, and the other at the end of the cycle
      # we may be need to change this: the second date isn't actually billed, but it is a billing cycle boundary
      assert_equal(2, billing_cycle_dates.count)
      # billing dates should be a year apart
      assert_includes([365, 366], (T.must(billing_cycle_dates[1]) - T.must(billing_cycle_dates[0])) / (60 * 60 * 24))

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # increase quantity by 1
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      subscription_schedule.refresh

      assert_equal(2, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases[1])

      # first phase should start now and end in 9mo
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # second phase should start at the end date
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

      assert_equal(1, first_phase.items.count)
      assert_equal(2, second_phase.items.count)

      first_phase_item = T.must(first_phase.items.first)
      second_phase_item = T.must(second_phase.items.detect {|i| i.price != first_phase_item.price })

      # out of an abundance of caution, ensure the correct prorated amount is calculated
      refute_empty(second_phase.add_invoice_items)
      assert_equal(1, second_phase.add_invoice_items.count)
      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)
      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)

      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      assert_equal("true", prorated_price.metadata['salesforce_proration'])
      assert_equal(second_phase_item.price, prorated_price.metadata['salesforce_original_stripe_price_id'])
    end

    it 'supports amending when there is a single billing cycle which has already passed' do
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts on Sept 30, billed yearly
      # amendment: starts in 5 months, Feb 28
      # test clock: advanced 1 month in the future so invoice already has passed

      contract_term = TEST_DEFAULT_CONTRACT_TERM
      amendment_term = 6
      # 1 year in the future, so our test clock advancement does not result in the past
      initial_order_start_date = DateTime.new(Time.now.year + 1, 10, 30).utc.beginning_of_day
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months
      amendment_end_date = initial_order_start_date + TEST_DEFAULT_CONTRACT_TERM.months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order = create_subscription_order(sf_product_id: sf_product_id, additional_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
      })

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # advance the clock past the first billing cycle
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)
      advance_test_clock(stripe_customer, (initial_order_start_date + 1.month).to_i)

      # fetch the corresponding subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # we expect the invoice api to fail
      # Stripe::InvalidRequestError: No upcoming invoices for customer: cus_MQpWiftgQua7UT
      no_upcoming_exception = assert_raises(Stripe::InvalidRequestError) do
        Stripe::Invoice.upcoming({
          subscription: subscription_schedule.subscription,
        }, @user.stripe_credentials)
      end
      assert_match("No upcoming invoices", no_upcoming_exception.message)

      # unit test the upcoming calculation
      billing_cycle_dates = StripeForce::Translate::OrderAmendment.calculate_billing_cycle_dates(@user, subscription_schedule, 12)

      # they should all be in the future
      assert(billing_cycle_dates.all? {|ts| ts >= now_time.to_i })
      assert(initial_order_start_date < Time.now.to_i)
      # first cycle should be a year apart from the billing cycle that has already passed (account for leap years)
      assert_includes([365, 366], (T.must(billing_cycle_dates.first) - initial_order_start_date.to_i) / (60 * 60 * 24))
      # there should be a single billing dates: one in the future, since the past one has already passed due to the test clock
      # we may be need to change this: this isn't actually billed, but it is a billing cycle boundary
      assert_equal(1, billing_cycle_dates.count)

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # increase quantity by 2
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      subscription_schedule.refresh
      assert_equal(2, subscription_schedule.phases.count)

      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases[1])

      # first phase should start now and end in 9mo
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # second phase should start at the end date
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

      assert_equal(1, first_phase.items.count)
      assert_equal(2, second_phase.items.count)

      first_phase_item = T.must(first_phase.items.first)
      second_phase_item = T.must(second_phase.items.detect {|i| i.price != first_phase_item.price })

      # out of an abundance of caution, ensure the correct prorated amount is calculated
      refute_empty(second_phase.add_invoice_items)
      assert_equal(1, second_phase.add_invoice_items.count)

      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      assert_equal("true", prorated_price.metadata['salesforce_proration'])
      assert_equal(second_phase_item.price, prorated_price.metadata['salesforce_original_stripe_price_id'])
    end
  end

  describe '#contract_co_terminated' do
    it 'order amendment does not coterminating with initial order' do
      # initial order: starts now, billed yearly
      # amendment: starts in 6 months, ends 7 months later (so one month after initial order)

      # setup
      initial_order_start_date = now_time
      # amendment term is intentionally longer than the original end date
      amendment_term = 7
      amendment_start_date = initial_order_start_date + 6.months

      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      sf_order.refresh

      # the contract should reference the initial order that was created
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # quote is generated by CPQ API, so set these fields manually
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate order
      contracts_not_coterminating_error = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("order amendments must coterminate with the initial order", contracts_not_coterminating_error.message.downcase)
    end

    it 'amendment order coterminates with initial order even though initial order day of month does not exist in amendment month' do
      # initial order: starts Sept 29, billed yearly
      # amendment: starts 5 months later, on Feb 28, and ends 7 months later (should co-terminate with initial order)

      # Pick an initial order date that will result in the amendment start date landing on the 29th of February in a non-leap year (ie the 29th doesn't exist)
      # If this test fails, it's most likely because we have passed initial_order_start_date + 5.months as initial_order_start_date is hard coded.
      # To fix, just bump initial_order_start_date a year forward so long as that year + 1 is not a leap year.
      initial_order_start_date = DateTime.new(2026, 9, 30).utc.beginning_of_day
      amendment_start_date = initial_order_start_date + 5.months
      amendment_term = 7

      # create the initial sf order
      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
        }
      )
      sf_contract = create_contract_from_order(sf_order)
      sf_order.refresh

      # the contract should reference the initial order that was created
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # quote is generated by CPQ API, so set these fields manually
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # create the sf order amendment
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate order and confirm no error raised
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    end

    it 'amendment order does not coterminate with initial order but both end in same month' do
      # initial order: starts Sept 27, billed yearly => ends Sept 27, 2023
      # amendment: starts 5 months + 1 day later, on Feb 28 => ends Sept 28, 2023

      # specifically pick an initial order day of month that does not exist in the amendment month
      initial_order_start_date = DateTime.new(2022, 9, 27).utc.beginning_of_day
      amendment_start_date = initial_order_start_date + 5.months + 1.day
      amendment_term = 7

      # create the initial sf order
      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
        }
      )
      sf_contract = create_contract_from_order(sf_order)
      sf_order.refresh

      # the contract should reference the initial order that was created
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # quote is generated by CPQ API, so set these fields manually
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # create the sf order amendment
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate order
      contracts_not_coterminating_error = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("amendment order subscription term does not equal number of whole months between start and end date", contracts_not_coterminating_error.message.downcase)
    end

    it 'non anniversary amendment order coterminates with initial order' do
      # initial order: starts Feb 15, year contracts, billed semi-annually
      initial_order_contract_term = TEST_DEFAULT_CONTRACT_TERM
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + initial_order_contract_term.months

      # first amendment order: starts 6 days later, and should co-terminate with initial order
      # even though the amendment starts on a different day of month than the initial order
      amendment_start_date_1 = initial_order_start_date + 6.days
      amendment_term_1 = 11

      # second amendment order: starts 2 months + 3 days later, and should co-terminate with initial order
      # even though the amendment starts on a different day of month than the initial order
      amendment_start_date_2 = initial_order_start_date + 2.months + 3.days
      amendment_term_2 = 9

      # create the initial sf order
      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::SEMIANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => initial_order_contract_term,
        }
      )

      # the contract should reference the initial order that was created
      sf_contract = create_contract_from_order(sf_order)
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # create and translate the first sf order amendment
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date_1)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term_1
      sf_order_amendment_1 = create_order_from_quote_data(amendment_data)

      # create and translate the second sf amendment order
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date_2)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term_2
      sf_order_amendment_2 = create_order_from_quote_data(amendment_data)

      # translate the order and confirm no error raised
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      sf_order_amendment_1.refresh
      sf_order_amendment_2.refresh

      # confirm subscription schedule phase dates are what we expect
      assert_equal(sf_order_amendment_1[prefixed_stripe_field(GENERIC_STRIPE_ID)], sf_order_amendment_2[prefixed_stripe_field(GENERIC_STRIPE_ID)])
      stripe_id = sf_order_amendment_1[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(3, subscription_schedule.phases.count)

      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date_1.to_i)

      second_phase = T.must(subscription_schedule.phases[1])
      assert_equal(0, second_phase.start_date - amendment_start_date_1.to_i)
      assert_equal(0, second_phase.end_date - amendment_start_date_2.to_i)

      third_phase = T.must(subscription_schedule.phases[2])
      assert_equal(0, third_phase.start_date - amendment_start_date_2.to_i)
      assert_equal(0, third_phase.end_date - initial_order_end_date.to_i)
    end

    it 'non anniversary amendment order missing CPQ_QUOTE_SUBSCRIPTION_END_DATE does not coterminate with initial order' do
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed semi-annually
      # amendment order: starts 10 days later and does not coterminate with initial order
      contract_term = 12
      initial_order_start_date = now_time
      amendment_order_start_date = initial_order_start_date + 15.days
      amendment_term = 11

      # create the initial sf order
      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::SEMIANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # the contract should reference the initial order that was created
      sf_contract = create_contract_from_order(sf_order)
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # create the sf order amendment
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_END_DATE] = nil
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate order and confirm error raised
      contracts_not_coterminating_error = assert_raises(Integrations::Errors::MissingRequiredFields) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("missing required fields", contracts_not_coterminating_error.message.downcase)
    end

    it 'non anniversary amendment order with invalid subscription term does not coterminate with initial order' do
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # initial order: starts now, billed semi-annually
      # amendment order: starts 10 days later and does not coterminate with initial order
      contract_term = 24
      initial_order_start_date = now_time
      amendment_order_start_date = initial_order_start_date + 2.months + 1.days

      # intentially set this subscription term to be wrong
      # the subscription term should be the number of whole months remaining in the amendment term
      long_invalid_amendment_term = 22
      short_invalid_amendment_term = 20

      # create the initial sf order
      sf_order = create_subscription_order(
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::SEMIANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # the contract should reference the initial order that was created
      sf_contract = create_contract_from_order(sf_order)
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      # create the sf order amendment
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = long_invalid_amendment_term
      sf_order_amendment_1 = create_order_from_quote_data(amendment_data)

      # create another amendment term and ensure it still errors
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = short_invalid_amendment_term
      sf_order_amendment_2 = create_order_from_quote_data(amendment_data)

      # translate order and confirm error raised
      contracts_not_coterminating_error = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order.Id)
      end
      assert_match('amendment order subscription term does not equal number of whole months between start and end date', contracts_not_coterminating_error.message.downcase)
    end
  end

  describe 'coupons' do
    it 'phase coupons are not copied over to new phase on amendment orders' do
      @user.enable_feature(FeatureFlags::COUPONS)

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      contract_term = TEST_DEFAULT_CONTRACT_TERM
      amendment_term = 6
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).month
      amendment_end_date = amendment_start_date + amendment_term.months

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # create a quote with a product
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
      sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # retrieve the quote line
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      assert_equal(1, quote_lines.size)
      quote_line_id = quote_lines.first.Id

      # create two coupons and attach one to the quote and the other to the quote line
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '25% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 25,
      })
      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$10 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 10,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # create and translate the initial sf order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # now create and translate an amendment
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # increase quantity of the line item by 1
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # translate the sf amendment
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # fetch the corresponding stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(2, subscription_schedule.phases.count)

      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases[1])

      # first phase should have one discount
      assert_equal(1, first_phase.discounts.count)

      # sanity check it's the coupon we attached
      original_phase_coupon_id = T.must(first_phase.discounts.first).coupon
      stripe_coupon = Stripe::Coupon.retrieve(original_phase_coupon_id, @user.stripe_credentials)
      sf_amount_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_amount_off_coupon_id)
      assert_equal(sf_amount_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])

      # first phase should have one item with a single quantity
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item[:quantity])

      # first phase item should have one coupon
      assert_equal(1, first_phase_item.discounts.count)

      # sanity check the phase item coupon data
      original_phase_item_coupon_id = T.must(first_phase_item.discounts.first).coupon
      stripe_coupon = Stripe::Coupon.retrieve(original_phase_item_coupon_id, @user.stripe_credentials)
      assert_equal(25, stripe_coupon.percent_off)
      assert_equal("once", stripe_coupon.duration)
      sf_percent_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_percent_off_coupon_id)
      assert_equal(sf_percent_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])

      # second phase should start at the end date
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

      # second phase should have two items
      second_phase_items = second_phase.items
      assert_equal(2, second_phase_items.count)

      # one of the second phase items should have a coupon
      phase_items_with_discount = second_phase_items.reject! {|d| d.discounts.empty? }
      assert_equal(1, T.must(phase_items_with_discount).count)

      # sanity check the second phase item coupon
      phase_item_discount = T.must(T.must(phase_items_with_discount).first).discounts
      stripe_coupon = Stripe::Coupon.retrieve(T.must(phase_item_discount.first).coupon, @user.stripe_credentials)
      assert_equal(25, stripe_coupon.percent_off)
      assert_equal("once", stripe_coupon.duration)
      assert_equal(sf_percent_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])
    end

    it 'supports adding new phase and phase item coupons to an sf amendment order' do
      @user.enable_feature(FeatureFlags::COUPONS)

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      contract_term = TEST_DEFAULT_CONTRACT_TERM
      amendment_term = 6
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).month
      amendment_end_date = amendment_start_date + amendment_term.months

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # create a quote with a product
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
      sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # create and translate the initial sf order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # now create an amendment order
      sf_contract = create_contract_from_order(sf_order)

      # increase quantity by 2
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # retrieve the amendment quote and quote line Id
      sf_amendment_quote_id = calculate_and_save_cpq_quote(amendment_data)
      quote_lines = sf_get_related(sf_amendment_quote_id, CPQ_QUOTE_LINE)
      assert_equal(1, quote_lines.size)
      sf_amendment_quote_line_id = quote_lines.first.Id

      # create two coupons and attach one to the quote and the other to quote line
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '26.7% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 26.7,
      })
      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$10 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 10,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: sf_amendment_quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_amendment_quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # create and translate amendment order
      sf_order_amendment = create_order_from_quote_data(amendment_data)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # fetch the corresponding stripe subscription schedule
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(2, subscription_schedule.phases.count)

      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases[1])

      # first phase should have one item and no coupons
      first_phase_items = first_phase.items
      assert_equal(1, first_phase_items.count)

      # first phase should have no discounts
      assert_empty(first_phase.discounts)
      # first phase item should no discounts
      assert_empty(T.must(first_phase_items.first).discounts)

      # second phase should start at the initial order end date
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

      # second phase should have one coupon
      assert_equal(1, second_phase.discounts.count)

      # second phase item should have one item with a single quantity
      assert_equal(2, second_phase.items.count)
      assert_equal(0, second_phase.add_invoice_items.count)

      # sanity check the phase coupon info
      amendment_phase_coupon_id = T.must(second_phase.discounts.first).coupon
      stripe_coupon = Stripe::Coupon.retrieve(amendment_phase_coupon_id, @user.stripe_credentials)
      assert_equal(10 * 100, stripe_coupon.amount_off)
      sf_amount_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_amount_off_coupon_id)
      assert_equal(sf_amount_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])

      # one of the two phase items should have a coupon
      items_with_coupons = second_phase.items.select {|item| item.discounts.count > 0 }
      assert_equal(1, items_with_coupons.count)

      # sanity check the phase item coupons
      amendment_phase_item_coupon_id = T.must(T.must(items_with_coupons.first).discounts.first).coupon
      stripe_coupon = Stripe::Coupon.retrieve(amendment_phase_item_coupon_id, @user.stripe_credentials)
      assert_in_delta(26.7, stripe_coupon.percent_off)
      assert_equal("once", stripe_coupon.duration)
      sf_percent_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_percent_off_coupon_id)
      assert_equal(sf_percent_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])
    end
  end

  describe 'invoice rendering template' do
    it 'updates the invoice rendering template for an amendment order' do
      @user.enable_feature(StripeForce::Constants::FeatureFlags::INVOICE_RENDERING_TEMPLATE)
      # setup
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months
      invoice_rendering_template_id = create_invoice_rendering_template

      @user.field_defaults = {
        "subscription_schedule" => {
          "default_settings.invoice_settings.rendering.template" => invoice_rendering_template_id,
        },
      }
      @user.save

      # create and translate a Salesforce initial order
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price
      sf_account_id = create_salesforce_account
      sf_order = create_salesforce_order(
        sf_product_id: sf_product_id,
        sf_account_id: sf_account_id,

        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
        }
      )

      SalesforceTranslateRecordJob.translate(@user, sf_order)
      sf_order.refresh

      # fetch the generated subscription and verify the invoice rendering template is set
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(invoice_rendering_template_id, subscription_schedule.default_settings.invoice_settings.rendering.template)

      # update the template
      @user.field_defaults = {
        "subscription_schedule" => {
          "default_settings.invoice_settings.rendering.template" => 'invalid_id',
        },
      }
      @user.save

      # create an amendment order and translate
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the sf amendment
      template_does_not_exist_error = assert_raises(Stripe::InvalidRequestError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match('no such invoice rendering template', template_does_not_exist_error.message.downcase)
    end
  end

  describe 'custom order filters' do
    it 'does not translate an amendment order if it does not meet user\'s custom filters' do
      # setup
      contract_term = 12
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date

      # create ('Activated') initial order
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price
      sf_order = create_subscription_order(sf_product_id: sf_product_id)
      sf_contract = create_contract_from_order(sf_order)

      # quote is generated by CPQ API, so set these fields manually
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = contract_term

      # create ('Draft') amendment order
      sf_quote_id = calculate_and_save_cpq_quote(amendment_quote)
      sf_amendment_order = create_draft_order_from_quote(sf_quote_id)

      # translate the amendment order
      StripeForce::Translate.perform_inline(@user, sf_amendment_order.Id)
      sf_amendment_order.refresh

      # we expect that the initial order was translated but not
      # the amendment order since it doesn't meet the user's custom order filters
      assert_nil(sf_amendment_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])

      sf_order.refresh
      stripe_subscription_schedule_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      # confirm only the initial order was translated and not the amendment order
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_subscription_schedule_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)
    end

    it 'does not translate either initial or amendment order if the initial order does not respect user\'s custom filters' do
      # set the user's custom filters such that the initial order will not respect them
      @user.connector_settings["filters"] = {SF_ORDER => "Status = 'NotRealStatus'"}
      @user.save

      # setup
      contract_term = 12
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date

      # create a quote with a product
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price
      sf_order = create_subscription_order(sf_product_id: sf_product_id)
      sf_contract = create_contract_from_order(sf_order)

      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = contract_term
      sf_amendment_order = create_order_from_quote_data(amendment_quote)

      # translate the amendment order
      exception = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_amendment_order.Id)
      end
      assert_match("attempted to sync amendment order when initial order was skipped because it didn't match custom sync filters", exception.message.downcase)
    end
  end

  describe 'update subscription schedule fields on amendments' do
    it 'update days_until_due on subscription schedule' do
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
      @user.field_defaults = {
        "customer" => {
          "email" => "test@test.com",
        },
        "subscription_schedule" => {
          "default_settings.invoice_settings.days_until_due" => "Net-30",
          "default_settings.collection_method" => "send_invoice", # note: you can only specify 'days_until_due' if invoice collection method is 'send_invoice'
        },
      }
      @user.save

      # setup
      contract_term = 12
      amendment_term = 11
      initial_order_start_date = now_time
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
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # translate initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      # get Stripe customer
      sf_account = sf_get(sf_order['AccountId'])
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      stripe_customer = stripe_get(stripe_customer_id)
      refute_nil(stripe_customer.test_clock)

      # check and pay off the initial invoice
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(1, invoices.length)

      initial_invoice = invoices.first
      assert_equal((now_time + 30.days).to_i, Time.at(initial_invoice.due_date).utc.beginning_of_day.to_i)
      initial_invoice.pay({'paid_out_of_band': true})

      # create the sf amendment order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_order_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # update days_until_due
      @user.field_defaults = {
        "subscription_schedule" => {
          "default_settings.invoice_settings.days_until_due" => "Net-15",
        },
      }
      @user.save

      # translate the sf amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order_amendment.refresh

      # advance the test clock
      advance_test_clock(stripe_customer, (amendment_order_start_date + 10.day).to_i)

      # confirm the new invoice has the update days_until_due
      invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
      assert_equal(2, invoices.length)

      proration_invoice = invoices.first {|invoice| invoice.status != "paid" }
      assert_equal((amendment_order_start_date + 15.days).to_i, Time.at(proration_invoice.due_date).utc.beginning_of_day.to_i)
    end
  end

  describe 'stacked amendments' do
    it 'syncs stacked amendments - one backdated and the other starts in the future' do
      # initial order: 1yr contract, billed annually
      # amendment 1: started 3 months ago
      # amendment 2: starts 1 month in the future
      contract_term = TEST_DEFAULT_CONTRACT_TERM
      initial_order_start_date = now_time - 4.months
      initial_order_end_date = initial_order_start_date + contract_term.months

      amendment_1_term = 9
      amendment_1_start_date = initial_order_start_date + (contract_term - amendment_1_term).months
      amendment_1_end_date = amendment_1_start_date + amendment_1_term.months
      # normalize the end_date so test doesn't fail EOM
      amendment_1_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_1_end_date, anchor_day_of_month: initial_order_end_date.day)

      amendment_2_term = 7
      amendment_2_start_date = initial_order_start_date + (contract_term - amendment_2_term).months
      # normalize the end_date so test doesn't fail EOM
      amendment_2_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_1_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the first amendment to increase quantity (+3)
      sf_contract_1 = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_1)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 4
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_1_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_1_term
      sf_order_amendment_1 = create_order_from_quote_data(amendment_quote)

      # create the second amendment to decrease quantity (-3) and add a standalone product
      sf_contract_2 = create_contract_from_order(sf_order_amendment_1)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_2)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 1
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_2_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_2_term
      sf_quote_id = calculate_and_save_cpq_quote(amendment_quote)

      # add standalone product
      sf_standalone_product_id, _sf_pricebook_id = salesforce_standalone_product_with_price
      amendment_data = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_standalone_product_id)
      _sf_order_amendment_2 = create_order_from_quote_data(amendment_data)

      # translate the all the orders (initial order and two amendments)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      # fetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(3, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)
      third_phase = T.must(subscription_schedule.phases.last)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start 'now' (since it was a backdated amendment)
      # and have two products with total quantity of 4
      assert_equal(0, second_phase.start_date.to_i - first_phase.end_date)
      assert_equal(0, second_phase.end_date - third_phase.start_date.to_i)
      # second phase should have a second item with a quantity of 3

      # TODO @nada why this is isn't one item with quantity 3
      assert_equal(2, second_phase.items.count)
      second_phase_item_1 = T.must(second_phase.items.first)
      second_phase_item_2 = T.must(second_phase.items.second)
      assert_equal(1, second_phase_item_1.quantity)
      assert_equal(3, second_phase_item_2.quantity)

      # prorate the added items added since the amendment was backdated and missed a billing cycle
      assert_equal(1, second_phase.add_invoice_items.count)
      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(3, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_1_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # third phase should start 1 month from now
      assert_equal(0, third_phase.start_date.to_i - second_phase.end_date.to_i)
      assert_equal(0, third_phase.end_date - initial_order_end_date.to_i)

      assert_equal(1, third_phase.items.count)
      T.must(third_phase.items.detect {|i| i[:quantity] == 1 })

      # there should be two invoice items - one for the standalone price and credit item for decrease quantity
      assert_equal(2, third_phase.add_invoice_items.count)
      standalone_item = T.must(third_phase.add_invoice_items.detect {|i| i[:quantity] == 1 })
      standalone_stripe_price = Stripe::Price.retrieve(T.cast(standalone_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', standalone_stripe_price.type)
      assert_equal(TEST_DEFAULT_STANDALONE_PRICE.round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(standalone_stripe_price.unit_amount_decimal))

      credit_item = T.must(third_phase.add_invoice_items.detect {|i| i[:quantity] == 3 })
      credit_stripe_price = Stripe::Price.retrieve(T.cast(credit_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', credit_stripe_price.type)
      assert_equal(-1 * (BigDecimal(TEST_DEFAULT_PRICE) * BigDecimal(amendment_2_term) / BigDecimal(contract_term)).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(credit_stripe_price.unit_amount_decimal))
    end

    it 'syncs stacked amendments - both backdated and start on same day' do
      # initial order: 1yr contract, billed annually
      # amendment 1: starts 3 months ago
      # amendment 2: starts same day as first amendment
      contract_term = TEST_DEFAULT_CONTRACT_TERM
      initial_order_start_date = now_time - 4.months
      initial_order_end_date = initial_order_start_date + contract_term.months

      amendment_term = 9
      amendment_start_date = initial_order_start_date + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months
      # normalize the end_date so test doesn't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the first amendment to increase quantity (+3)
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 4
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      sf_order_amendment_1 = create_order_from_quote_data(amendment_data)

      # create the second amendment and increase quantity again (+1)
      sf_contract_2 = create_contract_from_order(sf_order_amendment_1)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract_2)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 5
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      _sf_order_amendment_2 = create_order_from_quote_data(amendment_data)

      # translate the orders (initial order and two amendments)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      # let's attempt to resync these amendments but should skip syncing since they're already synced
      Stripe::SubscriptionSchedule.expects(:save).never
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(3, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)
      third_phase = T.must(subscription_schedule.phases.third)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start 'now' (since it was a backdated amendment)
      # and have two products with total quantity of 4
      assert_equal(0, second_phase.start_date.to_i - first_phase.end_date)
      assert_equal(0, second_phase.end_date - third_phase.start_date.to_i)
      # second phase should have a second item with total quantity 4
      assert_equal(2, second_phase.items.count)

      assert_equal(1, second_phase.add_invoice_items.count)
      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(3, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_auto_archive'])
      assert_equal("true", prorated_price.metadata['salesforce_duplicate'])
      assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # third phase should start 'now' as well (since it was a backdated amendment)
      # and have three products with total quantity of 5
      assert_equal(0, third_phase.start_date.to_i - second_phase.end_date)
      assert_equal(0, second_phase.end_date - third_phase.start_date.to_i)
      # second phase should have a second item with total quantity 4
      assert_equal(3, third_phase.items.count)

      assert_equal(1, third_phase.add_invoice_items.count)
      prorated_item = T.unsafe(third_phase.add_invoice_items.first)
      assert_equal(1, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
    end

    it 'syncs three stacked backdated amendments with last being a termination order' do
      @user.disable_feature(FeatureFlags::SF_CACHING)
      @user.field_mappings['subscription_schedule'] = {
        'metadata.OrderNumber' => 'OrderNumber',
      }
      @user.save

      # initial order: 1yr contract, billed annually, started 3 months ago
      # amendment 1: started 2 months ago
      # amendment 2: started 1 month ago
      # amendment 3: started 3 days ago
      contract_term = TEST_DEFAULT_CONTRACT_TERM
      initial_order_start_date = now_time - 3.months - 3.days
      initial_order_end_date = initial_order_start_date + contract_term.months

      amendment_1_term = 11
      amendment_1_start_date = initial_order_start_date + (contract_term - amendment_1_term).months

      amendment_2_term = 10
      amendment_2_start_date = initial_order_start_date + (contract_term - amendment_2_term).months

      amendment_3_term = 9
      amendment_3_start_date = initial_order_start_date + (contract_term - amendment_3_term).months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the first amendment to increase quantity (+2)
      sf_contract_1 = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_1)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_1_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_1_term
      sf_order_amendment_1 = create_order_from_quote_data(amendment_quote)

      # create the second amendment to decrease quantity (-2)
      sf_contract_2 = create_contract_from_order(sf_order_amendment_1)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_2)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 1
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_2_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_2_term
      sf_order_amendment_2 = create_order_from_quote_data(amendment_quote)

      # translate the orders (initial order and two amendments)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      # create the third amendment to terminate the order
      sf_contract_3 = create_contract_from_order(sf_order_amendment_2)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_3)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_3_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_3_term
      sf_order_amendment_3 = create_order_from_quote_data(amendment_quote)

      # translate the termination order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment_3.Id)

      # fetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(3, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)
      third_phase = T.must(subscription_schedule.phases.third)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start 'now' (since it was a backdated amendment)
      # and have two products with total quantity of 2
      assert(second_phase.start_date.to_i - now_time.to_i < SECONDS_IN_DAY)
      assert_equal(0, second_phase.end_date - third_phase.start_date.to_i)
      # second phase should have a second item with a quantity of 1
      assert_equal(2, second_phase.items.count)
      second_phase_item_1 = T.must(second_phase.items.first)
      second_phase_item_2 = T.must(second_phase.items.second)
      assert_equal(1, second_phase_item_1.quantity)
      assert_equal(2, second_phase_item_2.quantity)

      # prorate the added items added since the amendment was backdated and missed a billing cycle
      assert_equal(1, second_phase.add_invoice_items.count)
      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(2, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_1_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # third phase should start 'now' (since it was a backdated amendment)
      assert_equal(0, third_phase.start_date.to_i - second_phase.end_date.to_i)
      assert_equal(0, third_phase.end_date.to_i - initial_order_end_date.to_i)
      assert_equal(1, third_phase.items.count)
      T.must(third_phase.items.detect {|i| i[:quantity] == 1 })

      # there should be one invoice items - credit item for decrease quantity
      assert_equal(1, third_phase.add_invoice_items.count)
      credit_item = T.must(third_phase.add_invoice_items.detect {|i| i[:quantity] == 2 })
      credit_stripe_price = Stripe::Price.retrieve(T.cast(credit_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', credit_stripe_price.type)
      assert_equal(-1 * (BigDecimal(TEST_DEFAULT_PRICE) * BigDecimal(amendment_2_term) / BigDecimal(contract_term)).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(credit_stripe_price.unit_amount_decimal))

      # ensure that metadata was remapped during termination
      assert_equal(sf_order.OrderNumber, subscription_schedule.metadata['OrderNumber'])
      # ensure that terminate metadata was added during termination
      amendment_opportunity = sf_get(sf_order_amendment_3["OpportunityId"])
      assert_equal(amendment_opportunity[SF_OPPORTUNITY_CLOSE_DATE], subscription_schedule.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::EFFECTIVE_TERMINATION_DATE)])
    end

    it 'syncs three stacked backdated amendments with quantity changes on different runs' do
      @user.disable_feature(FeatureFlags::SF_CACHING)
      # initial order: 1yr contract, billed annually, started 3 months ago
      # amendment 1: started 2 months ago
      # amendment 2: started 1 month ago
      # amendment 3: started 3 days ago
      contract_term = TEST_DEFAULT_CONTRACT_TERM
      initial_order_start_date = now_time - 3.months - 3.days
      initial_order_end_date = initial_order_start_date + contract_term.months

      amendment_1_term = 11
      amendment_1_start_date = initial_order_start_date + (contract_term - amendment_1_term).months

      amendment_2_term = 10
      amendment_2_start_date = initial_order_start_date + (contract_term - amendment_2_term).months

      amendment_3_term = 9
      amendment_3_start_date = initial_order_start_date + (contract_term - amendment_3_term).months

      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create the initial sf order
      sf_order = create_subscription_order(
        sf_product_id: sf_product_id,
        additional_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create the first amendment to increase quantity (+2)
      sf_contract_1 = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_1)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_1_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_1_term
      sf_order_amendment_1 = create_order_from_quote_data(amendment_quote)

      # translate the orders (initial order and first amendment)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # create the second amendment to increase quantity (+3)
      sf_contract_2 = create_contract_from_order(sf_order_amendment_1)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_2)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 6
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_2_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_2_term
      sf_order_amendment_2 = create_order_from_quote_data(amendment_quote)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment_2.Id)

      # create the third amendment to decrease the quantity (-5)
      sf_contract_3 = create_contract_from_order(sf_order_amendment_2)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract_3)
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 1
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_3_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_3_term
      sf_order_amendment_3 = create_order_from_quote_data(amendment_quote)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment_3.Id)

      # fetch the subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(4, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)
      third_phase = T.must(subscription_schedule.phases.third)
      fourth_phase = T.must(subscription_schedule.phases.fourth)

      # first phase should start at the backdated date
      assert_equal(0, first_phase.start_date - initial_order_start_date.to_i)
      assert_equal(0, first_phase.end_date - second_phase.start_date)
      # first phase should have an item with a quantity of 1 and no invoice items
      assert_equal(1, first_phase.items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item.quantity)
      assert_empty(first_phase.add_invoice_items)

      # second phase should start 'now' (since it was a backdated amendment)
      # and have two products with total quantity of 2
      assert(second_phase.start_date.to_i - now_time.to_i < SECONDS_IN_DAY)
      assert_equal(0, second_phase.end_date - third_phase.start_date.to_i)
      # second phase should have an item with a quantity of 3
      assert_equal(2, second_phase.items.count)
      second_phase_item_1 = T.must(second_phase.items.first)
      second_phase_item_2 = T.must(second_phase.items.second)
      assert_equal(1, second_phase_item_1.quantity)
      assert_equal(2, second_phase_item_2.quantity)

      # prorate the added items added since the amendment was backdated and missed a billing cycle
      assert_equal(1, second_phase.add_invoice_items.count)
      prorated_item = T.unsafe(second_phase.add_invoice_items.first)
      assert_equal(2, prorated_item.quantity)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_item.price, String), @user.stripe_credentials)
      assert_equal('one_time', prorated_price.type)
      assert_equal((TEST_DEFAULT_PRICE / (contract_term / BigDecimal(amendment_1_term))).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(prorated_price.unit_amount_decimal))
      assert_equal("true", prorated_price.metadata['salesforce_proration'])

      # third phase should start 'now' (since it was a backdated amendment)
      assert_equal(0, third_phase.start_date.to_i - second_phase.end_date.to_i)
      assert_equal(0, third_phase.end_date.to_i - fourth_phase.start_date.to_i)
      assert_equal(3, third_phase.items.count)
      T.must(third_phase.items.detect {|i| i[:quantity] == 1 })
      T.must(third_phase.items.detect {|i| i[:quantity] == 2 })
      T.must(third_phase.items.detect {|i| i[:quantity] == 3 })

      # fourth phase should start 'now' (since it was a backdated amendment)
      assert_equal(0, fourth_phase.end_date.to_i - initial_order_end_date.to_i)
      assert_equal(1, fourth_phase.items.count)
      T.must(fourth_phase.items.detect {|i| i[:quantity] == 1 })

      # there should be one invoice items - credit item for decrease quantity
      assert_equal(2, fourth_phase.add_invoice_items.count)
      credit_item_1 = T.must(fourth_phase.add_invoice_items.detect {|i| i[:quantity] == 2 })
      _credit_item_2 = T.must(fourth_phase.add_invoice_items.detect {|i| i[:quantity] == 3 })
      credit_stripe_price = Stripe::Price.retrieve(T.cast(credit_item_1.price, String), @user.stripe_credentials)
      assert_equal('one_time', credit_stripe_price.type)
      assert_equal(-1 * (BigDecimal(TEST_DEFAULT_PRICE) * BigDecimal(amendment_3_term) / BigDecimal(contract_term)).round(MAX_STRIPE_PRICE_PRECISION), BigDecimal(credit_stripe_price.unit_amount_decimal))
    end
  end

  describe 'metadata' do
    it 'pulls metadata from each order amendment to the phase of each subscription'
    it 'uses metadata on the original line item if an item is not removed'
    it 'uses the latest metadata on an order line represented in a previous subscription schedule phase'
  end

end
