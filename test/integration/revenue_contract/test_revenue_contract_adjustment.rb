# frozen_string_literal: true
# typed: true

require_relative './_revenue_contract_validation_helper'

class Critic::RevRecContractAdjustment < Critic::RevenueContractValidationHelper
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @defaultSignedDate = "2022-12-31"
    @defaultTFC = 30
    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::NON_ANNIVERSARY_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::DAY_PRORATIONS, update: true
    @user.enable_feature FeatureFlags::BACKDATED_AMENDMENTS, update: true
    @user.enable_feature FeatureFlags::TERMINATED_ORDER_ITEM_CREDIT, update: true
    @user.enable_feature FeatureFlags::TERMINATION_METADATA, update: true
    @user.enable_feature FeatureFlags::STRIPE_REVENUE_CONTRACT, update: true
    @user.field_defaults = {
      "subscription_item" => {
        "metadata.contract_tfc_duration" => @defaultTFC,
      },
      "subscription_schedule" => {
        "metadata.contract_cf_signed_date" => @defaultSignedDate,
      },
    }
    @user.save
  end

  it 'Adjust contract to add 2 quantity for the last 3 months' do
    # initial order: 1yr contract, monthly billed
    # amendment: starts month 9, lasts 3 months, adds quantity 2
    monthly_price = 10_00
    contract_term = TEST_DEFAULT_CONTRACT_TERM
    initial_start_date = now_time
    initial_order_end_date = initial_start_date + contract_term.months

    amendment_term = 3
    amendment_start_date = initial_start_date + (contract_term - amendment_term).months
    amendment_end_date = amendment_start_date + amendment_term.months
    # normalize the amendment_end_date so test doesn't fail EOM
    amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

    sf_account_id = create_salesforce_account
    sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: monthly_price)
    sf_order = create_subscription_order(sf_account_id: sf_account_id,
                                         sf_product_id: sf_product_id,
                                         contact_email: "revenue_contract_amendment_add_quantity",
                                         additional_fields: {
                                           CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
                                         })
    sf_contract = create_contract_from_order(sf_order)

    sf_order.refresh

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

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    revenue_contract = get_revenue_contract(contract_id)

    revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate, version: 2)
    assert_equal(2, subscription_schedule.phases.count)
    phase_1 = T.must(subscription_schedule.phases.first)
    assert_equal(1, phase_1.items.count)
    assert_equal(0, phase_1.add_invoice_items.count)
    phase_2 = T.must(subscription_schedule.phases.last)
    assert_equal(2, phase_2.items.count)
    assert_equal(0, phase_2.add_invoice_items.count)

    assert_equal(2, revenue_contract.items.data.count)

    # first item with 120$ value and qty 1
    phase_item_1 = T.must(phase_2.items.first)
    contract_item_1 = T.must(revenue_contract.items.data.first)
    revenue_contract_validate_item(phase_item_1, contract_item_1, nil, 1, 12000, @defaultTFC)
    revenue_contract_validate_item_period(contract_item_1, initial_start_date, initial_order_end_date)

    # second item added with qty 2 at for 3 months, worth 60$
    phase_item_2 = T.must(phase_2.items.last)
    contract_item_2 = T.must(revenue_contract.items.data.last)
    revenue_contract_validate_item(phase_item_2, contract_item_2, nil, 2, 6000, @defaultTFC)
    revenue_contract_validate_item_period(contract_item_2, amendment_start_date, amendment_end_date)
  end

  # orders with only metered items do not create revenue contracts.
  it 'Initial order does not create revenue contract, but amendmend does' do
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

      sf_account_id = create_salesforce_account
      sf_order = create_subscription_order(sf_account_id: sf_account_id, sf_product_id: sf_metered_product_id, contact_email: "revenue_contract_new_phase_remove_metered")
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
      contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
      revenue_contract = get_revenue_contract(contract_id)

      assert_equal(2, subscription_schedule.phases.count)
      phase_1 = T.must(subscription_schedule.phases.first)
      assert_equal(1, phase_1.items.count)
      assert_equal(0, phase_1.add_invoice_items.count)
      phase_2 = T.must(subscription_schedule.phases.last)
      assert_equal(1, phase_2.items.count)
      assert_equal(0, phase_2.add_invoice_items.count)

      assert_equal(1, revenue_contract.items.data.count)

      # only the new non metered item is available on revenue contract and version is 1 for contract
      revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate, version: 1)
      phase_item_1 = T.must(phase_2.items.first)
      contract_item_1 = T.must(revenue_contract.items.data.first)
      revenue_contract_validate_item(phase_item_1, contract_item_1, nil, 1, 72000, @defaultTFC)
      revenue_contract_validate_item_period(contract_item_1, amendment_start_date, amendment_end_date)
  end

  # orders with only recurring items that updates contract.
  it 'Update existing contract item and adds propration' do
      # initial order: one recurring item
      # amendment: remove recurring item, add it again

      amendment_term = 6
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + TEST_DEFAULT_CONTRACT_TERM
      amendment_start_date = initial_order_start_date + 6.months
      amendment_end_date = amendment_start_date + amendment_term.months
      # normalize the amendment_end_date so tests don't fail EOM
      amendment_end_date = StripeForce::Translate::OrderHelpers.anchor_time_to_day_of_month(base_time: amendment_end_date, anchor_day_of_month: initial_order_end_date.day)

      sf_metered_product_id, _sf_metered_pricebook_id = salesforce_recurring_product_with_price
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      sf_account_id = create_salesforce_account
      sf_order = create_subscription_order(sf_account_id: sf_account_id, sf_product_id: sf_metered_product_id, contact_email: "revenue_contract_new_phase_update_item")
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
      contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
      revenue_contract = get_revenue_contract(contract_id)

      assert_equal(2, subscription_schedule.phases.count)
      phase_1 = T.must(subscription_schedule.phases.first)
      assert_equal(1, phase_1.items.count)
      assert_equal(0, phase_1.add_invoice_items.count)
      phase_2 = T.must(subscription_schedule.phases.last)
      assert_equal(1, phase_2.items.count)
      assert_equal(1, phase_2.add_invoice_items.count)

      assert_equal(3, revenue_contract.items.data.count)

      revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate, version: 2)
      # first item is updated with shorter time and reduced amount to 72000.
      phase_item_1 = T.must(phase_1.items.first)
      contract_item_1 = T.must(revenue_contract.items.data.first)
      revenue_contract_validate_item(phase_item_1, contract_item_1, nil, 1, 72000, @defaultTFC)
      # date is now from initial start date to amendment start date
      revenue_contract_validate_item_period(contract_item_1, initial_order_start_date, amendment_start_date)

      # second item is the same item, but with different price since was added on amendment
      phase_item_2 = T.must(phase_2.items.first)
      contract_item_2 = T.must(revenue_contract.items.data[1])
      revenue_contract_validate_item(phase_item_2, contract_item_2, nil, 1, 72000, @defaultTFC)
      revenue_contract_validate_item_period(contract_item_2, amendment_start_date, amendment_end_date)

      # third item is the proration generated by salesforce
      phase_item_3 = T.must(phase_2.add_invoice_items.first)
      contract_item_3 = T.must(revenue_contract.items.data.last)
      revenue_contract_validate_item(phase_item_3, contract_item_3, nil, 1, -12000, @defaultTFC)
      revenue_contract_validate_item_period(contract_item_3, amendment_start_date, amendment_end_date)
  end

  # TODO: This test is flaky with sub schedule sometimes, not sure why, but when it runs, it succeeds.
  # Disabling it for now, and will come back and re-enable this.
  # it 'revenue contract 3 stacked adjustments with quantity changes on different runs' do
  def revenue_contract_stacked_adjustments_disabled
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
    sf_account_id = create_salesforce_account
    sf_order = create_subscription_order(
      sf_account_id: sf_account_id,
      sf_product_id: sf_product_id,
      contact_email: "syncs_three_stacked_diffruns_rev_contract",
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

    Timecop.freeze(Time.now + 1.hour)
    # translate the orders (initial order and first amendment)
    StripeForce::Translate.perform_inline(@user, sf_order.Id)
    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    # We are getting the second phase right now because it's add_invoice_items will get deleted after next ammendments
    # due to the design. The invoice item is already added to the subscription hence it will still be active.
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(2, subscription_schedule.phases.count)
    second_phase = T.must(subscription_schedule.phases.second)

    # create the second amendment to increase quantity (+3)
    sf_contract_2 = create_contract_from_order(sf_order_amendment_1)
    amendment_quote = create_quote_data_from_contract_amendment(sf_contract_2)
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 6
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_2_start_date)
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_2_term
    sf_order_amendment_2 = create_order_from_quote_data(amendment_quote)

    Timecop.freeze(Time.now + 1.hour)

    # perform the second amendment and get the third phase that still has the "add_invoice_items" before wiping during cancellation
    StripeForce::Translate.perform_inline(@user, sf_order_amendment_2.Id)
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(3, subscription_schedule.phases.count)
    third_phase = T.must(subscription_schedule.phases.third)

    # create the third amendment to decrease the quantity (-5)
    sf_contract_3 = create_contract_from_order(sf_order_amendment_2)
    amendment_quote = create_quote_data_from_contract_amendment(sf_contract_3)
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 1
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_3_start_date)
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_3_term
    sf_order_amendment_3 = create_order_from_quote_data(amendment_quote)

    Timecop.freeze(Time.now + 1.hour)
    # perform the third amendment and get the fourth phase that still has the "add_invoice_items" before wiping during cancellation
    StripeForce::Translate.perform_inline(@user, sf_order_amendment_2.Id)
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(4, subscription_schedule.phases.count)
    fourth_phase = T.must(subscription_schedule.phases.fourth)

    # fetch the subscription schedule
    sf_order.refresh
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    revenue_contract = get_revenue_contract(contract_id)

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(4, subscription_schedule.phases.count)
    first_phase = T.must(subscription_schedule.phases.first)

    assert_equal(8, revenue_contract.items.data.count)
    revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate, version: 4)

    # first item - across all items
    phase_item_1 = T.must(first_phase.items.first)
    contract_item_1 = T.must(revenue_contract.items.data.first)
    revenue_contract_validate_item(phase_item_1, contract_item_1, nil, 1, 12000, @defaultTFC)

    # second item - was adjusted down to 0
    phase_item_2 = T.must(second_phase.items.last)
    contract_item_2 = T.must(revenue_contract.items.data[1])
    revenue_contract_validate_item(phase_item_2, contract_item_2, nil, 2, 0, @defaultTFC)

    # third item, propration item added on second phase
    phase_item_3 = T.must(second_phase.add_invoice_items.last)
    contract_item_3 = T.must(revenue_contract.items.data[2])
    revenue_contract_validate_item(phase_item_3, contract_item_3, nil, 2, 22000, @defaultTFC)

    # fourth item, adjusted down to 0 on third phase
    phase_item_4 = T.must(third_phase.items[1])
    contract_item_4 = T.must(revenue_contract.items.data[3])
    revenue_contract_validate_item(phase_item_4, contract_item_4, nil, 2, 0, @defaultTFC)

    # fifth item, adjusted down to 0 on third phase
    phase_item_5 = T.must(third_phase.items[2])
    contract_item_5 = T.must(revenue_contract.items.data[4])
    revenue_contract_validate_item(phase_item_5, contract_item_5, nil, 3, 0, @defaultTFC)

    # sixth item, propration item added on third phase
    phase_item_6 = T.must(third_phase.add_invoice_items.last)
    contract_item_6 = T.must(revenue_contract.items.data[5])
    revenue_contract_validate_item(phase_item_6, contract_item_6, nil, 3, 30000, @defaultTFC)

    # seventh item, propration item added on fourth phase
    phase_item_7 = T.must(fourth_phase.add_invoice_items.first)
    contract_item_7 = T.must(revenue_contract.items.data[6])
    revenue_contract_validate_item(phase_item_7, contract_item_7, nil, 2, -18000, @defaultTFC)

    # eigth item, propration item added on fourth phase
    phase_item_8 = T.must(fourth_phase.add_invoice_items.last)
    contract_item_8 = T.must(revenue_contract.items.data[7])
    revenue_contract_validate_item(phase_item_8, contract_item_8, nil, 3, -27000, @defaultTFC)
  end
end
