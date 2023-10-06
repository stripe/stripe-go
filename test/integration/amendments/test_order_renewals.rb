# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::OrderRenewalTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
    @user.enable_feature(FeatureFlags::TERMINATION_METADATA)
  end

  it 'creates a renewal salesforce order' do
    # setup
    sf_initial_order = create_subscription_order(contact_email: "create_renewal_order_3")
    # initial orders have a nil contract ID
    assert_nil(sf_initial_order.ContractId)

    sf_initial_contract = create_contract_from_order(sf_initial_order)
    sf_initial_order.refresh
    # sanity check that this sf order corresponds to this sf contract
    assert_equal(sf_initial_order.Id, sf_initial_contract.SBQQ__Order__c)
    # sanity check that the sf contract doesn't have "Renewal Quoted" set to true
    refute(sf_initial_contract.SBQQ__RenewalQuoted__c)

    # creates a renewal opportunity/quote from this contract
    sf_renewal_quote = create_renewal_quote_from_contract(sf_initial_contract)

    sf_initial_contract.refresh
    assert(sf_initial_contract.SBQQ__RenewalQuoted__c)

    # verify the renewal opportunity is the same in between the renewal quote and initial contract
    assert_equal(sf_renewal_quote.SBQQ__Opportunity2__c, sf_initial_contract.SBQQ__RenewalOpportunity__c)
    # verify renewal term matches the initial contract’s term
    assert_equal(sf_renewal_quote.SBQQ__RenewalTerm__c, sf_initial_contract.ContractTerm)
    # verify accounts is the same between renewal quote and initial contract
    assert_equal(sf_renewal_quote.SBQQ__Account__c, sf_initial_contract.AccountId)
    # verify renewal end date is the original order end date + sub term
    sf_renewal_quote_end_timestamp = DateTime.parse(sf_renewal_quote.SBQQ__EndDate__c).to_i
    assert_equal(sf_renewal_quote_end_timestamp, (DateTime.parse(sf_initial_contract.EndDate) + sf_initial_contract.ContractTerm.months).to_i)
  end

  it 'syncs a renewal salesforce order' do
    # create and sync the initial sf order
    sf_initial_order = create_subscription_order(contact_email: "sync_renewal_order_33")
    StripeForce::Translate.perform_inline(@user, sf_initial_order.Id)
    sf_initial_order.refresh

    sf_initial_contract = create_contract_from_order(sf_initial_order)
    # creates a renewal opportunity/quote from this contract and translates the renewal order
    sf_renewal_quote = create_renewal_quote_from_contract(sf_initial_contract)

    # TODO look into this since we would expect the subscription term to be populated here
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_renewal_quote.Id,
      CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
    })

    # create and sync the renewal order
    sf_renewal_order = create_order_from_cpq_quote(sf_renewal_quote.Id)
    StripeForce::Translate.perform_inline(@user, sf_renewal_order.Id)
    sf_renewal_order.refresh

    # fetch the initial order's subscription schedule
    initial_order_stripe_id = sf_initial_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    initial_order_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(initial_order_stripe_id, @user.stripe_credentials)
    assert_equal(1, initial_order_subscription_schedule.phases.count)

    # fetch the renewal order subscription schedule
    renewal_order_stripe_id = sf_renewal_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    renewal_order_subscription_schedule = Stripe::SubscriptionSchedule.retrieve(renewal_order_stripe_id, @user.stripe_credentials)
    assert_equal('not_started', renewal_order_subscription_schedule.status)
    assert_equal(1, renewal_order_subscription_schedule.phases.count)

    # the initial order and renewal order should have the same products
    initial_order_phase = T.must(initial_order_subscription_schedule.phases.first)
    renewal_order_phase = T.must(renewal_order_subscription_schedule.phases.first)

    # confirm that the renewal subscription schedule starts when the original subscription schedule ends
    assert_equal(initial_order_phase.end_date, renewal_order_phase.start_date)
    assert_equal(initial_order_phase.items.count, renewal_order_phase.items.count)
    assert_equal(T.must(initial_order_phase.items.first)[GENERIC_STRIPE_ID], T.must(renewal_order_phase.items.first)[GENERIC_STRIPE_ID])
  end

  it 'amend a renewal order' do
    # create the initial order
    sf_initial_order = create_subscription_order(contact_email: "amend_renewal_3")
    StripeForce::Translate.perform_inline(@user, sf_initial_order.Id)
    sf_initial_order.refresh

    # create and sync the renewal order
    sf_initial_contract = create_contract_from_order(sf_initial_order)
    # creates a renewal opportunity/quote from this contract and translates the renewal order
    sf_renewal_quote = create_renewal_quote_from_contract(sf_initial_contract)
    assert_equal(sf_renewal_quote[CPQ_QUOTE_TYPE], CPQQuoteTypeOptions::RENEWAL.serialize)
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_renewal_quote.Id,
      CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
    })

    sf_renewal_order = create_order_from_cpq_quote(sf_renewal_quote.Id)
    # even though this can be customized, let's check the default
    assert_equal(sf_renewal_order.Type, OrderTypeOptions::RENEWAL.serialize)

    StripeForce::Translate.perform_inline(@user, sf_renewal_order.Id)
    sf_renewal_order.refresh

    # create an amendment order to increase item quantity by 2
    sf_contract = create_contract_from_order(sf_renewal_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(now_time + 1.years + 2.month)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = TEST_DEFAULT_CONTRACT_TERM - 2
    sf_order_amendment = create_order_from_quote_data(amendment_data)

    # sync the amendment order
    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    sf_order_amendment.refresh
    stripe_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    # validate the subscription schedule
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal('not_started', subscription_schedule.status)
    assert_equal(2, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # first phase should have one item with a quantity of one
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_equal(1, first_phase_item.quantity)

    # first phase should start in year and end in 2 months
    assert_equal((now_time + 1.years).to_i, first_phase.start_date)
    assert_equal((now_time + 1.years + 2.months).to_i, first_phase.end_date)

    # second phase should start at the end date
    assert_equal(first_phase.end_date, second_phase.start_date)
    assert_equal((now_time + 2.years).to_i, second_phase.end_date)

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
  end

  it 'terminate a renewal order' do
    # create the initial order
    sf_initial_order = create_subscription_order(contact_email: "terminate_renewal_6")
    StripeForce::Translate.perform_inline(@user, sf_initial_order.Id)
    sf_initial_order.refresh

    # create and sync the renewal order
    sf_initial_contract = create_contract_from_order(sf_initial_order)
    # creates a renewal opportunity/quote from this contract and translates the renewal order
    sf_renewal_quote = create_renewal_quote_from_contract(sf_initial_contract)
    assert_equal(sf_renewal_quote[CPQ_QUOTE_TYPE], CPQQuoteTypeOptions::RENEWAL.serialize)
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_renewal_quote.Id,
      CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
    })

    sf_renewal_order = create_order_from_cpq_quote(sf_renewal_quote.Id)
    # even though this can be customized, let's check the default
    assert_equal(sf_renewal_order.Type, OrderTypeOptions::RENEWAL.serialize)

    StripeForce::Translate.perform_inline(@user, sf_renewal_order.Id)
    sf_renewal_order.refresh

    # create an amendment order to terminate the renewal order
    sf_contract = create_contract_from_order(sf_renewal_order)
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # wipe out the product
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(now_time + 1.years + 1.month)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = TEST_DEFAULT_CONTRACT_TERM - 1

    sf_order_amendment = create_order_from_quote_data(amendment_data)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    sf_order_amendment.refresh

    stripe_id = sf_order_amendment[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(1, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)
    assert_equal(1, first_phase.items.count)
    # the first phase should end when the amendment/termination order starts
    assert_equal((now_time + 1.years + 1.month).to_i, first_phase.end_date)
  end

  it 'throws an error when attempting to terminate a fake renewal order (order with revised product field set)' do
    # fake order we don't care about
    sf_order_0 = create_subscription_order(contact_email: "terminate_renewal_order_15")
    sf_order_items_0 = sf_get_related(sf_order_0, SF_ORDER_ITEM)

    # renewal order (just a "new order" but with the SBQQ__RevisedOrderProduct__c field set)
    sf_order = create_subscription_order(contact_email: "terminate_renewal_order_16")
    # add the Revised_Order_product field to the Order Item
    sf_order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    sf.update!(SF_ORDER_ITEM,
      SF_ID => sf_order_items[0].Id,
      # "EndDate" => format_date_for_salesforce(now_time_in_future + 364.days),
      "SBQQ__RevisedOrderProduct__c" => sf_order_items_0[0].Id,
    )

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_contract = create_contract_from_order(sf_order)
    sf_order.refresh

    # create an amendment order to terminate the renewal order
    amendment_end_date = now_time + 9.months
    amendment_data = create_quote_data_from_contract_amendment(sf_contract)
    # wipe out the product
    amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_end_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = 3

    sf_order_amendment = create_order_from_quote_data(amendment_data)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    exception = assert_raises(TypeError) do
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    end
    assert_match("Passed `nil` into T.must", exception.message)
  end
end
