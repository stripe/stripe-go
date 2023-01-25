# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::OrderAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
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

  # TODO pretty sure this is tested in the price reuse tests, check if any additional cases need to be tested here
  it 'customized prices on the line level' do

  end

  it 'supports adding one-off line items on a order amendment' do

  end

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
      assert_equal(365, (T.must(billing_cycle_dates[1]) - T.must(billing_cycle_dates[0])) / (60 * 60 * 24))

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

    it 'supports amending when there is a single billing cycle which has already passed ' do
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
    it 'test order amendments not coterminating with initial order' do
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

    it 'initial order should coterminate with amendment even though initial order day of month does not exist in amendment month' do
      # initial order: starts Sept 29, billed yearly
      # amendment: starts 5 months later, on Feb 28, and ends 7 months later (should co-terminate with initial order)

      # specifically pick an initial order day of month that does not exist in the amendment month
      initial_order_start_date = DateTime.new(2022, 9, 30).utc.beginning_of_day
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

    it 'initial order does not coterminate with amendment order but both end in same month' do
      # initial order: starts Sept 27, billed yearly => ends Sept 27, 2023
      # amendment: starts a day and 5 months later, on Feb 28, and ends 7 months later => Sept 28, 2023

      # specifically pick an initial order day of month that does not exist in the amendment month
      initial_order_start_date = DateTime.new(2022, 9, 27).utc.beginning_of_day
      amendment_start_date = initial_order_start_date + 1.day + 5.months
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
      assert_match("order amendments must coterminate with the initial order", contracts_not_coterminating_error.message.downcase)
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

  describe 'metadata' do
    it 'pulls metadata from each order amendment to the phase of each subscription'
    it 'uses metadata on the original line item if an item is not removed'
    it 'uses the latest metadata on an order line represented in a previous subscription schedule phase'
  end

end
