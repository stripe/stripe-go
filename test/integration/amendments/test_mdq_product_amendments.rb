# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::MDQAmendmentsTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
    @user.feature_enabled?(FeatureFlags::MDQ)
  end

  describe 'amending orders with mdq products' do
    it 'same day amend the current phase of salesforce order with a mdq product' do
      # setup
      contract_term = 24
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date # amendment quote start date must be the start date of the current phase
      amendment_term = contract_term

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_same_day_0_3",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})
      segmented_metered_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_metered_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # add a non-segmented metered product to quote
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_metered_produce_with_price
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # three products of which two products have two segments
      assert_equal(5, quote_with_product["lineItems"].count)

      # create the initial order
      sf_order = create_order_from_cpq_quote(quote_id)

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # segmented product creates sub schedule with two phases
      assert_equal(2, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(3, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)

      # You can't update a single segment of a segmented product or the following error will be thrown
      # Restforce::ErrorCode::FieldCustomValidationException: FIELD_CUSTOM_VALIDATION_EXCEPTION: Can't activate an order with partially segmented order items.
      amendment_data["lineItems"].each do |line_item|
        # let's specifically increase the quantity of the licensed mdq product
        if line_item["record"][CPQ_QUOTE_LINE_PRODUCT] == segmented_licensed_product
          line_item["record"][CPQ_QUOTE_QUANTITY] = 2
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      # assert only the affected order items are in the amendment order across the segments
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(2, order_lines.count)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh

      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # number of phases should increase since we are ending the current phase and adding a new phase
      assert_equal(3, subscription_schedule.phases.count)
      subscription_schedule.phases.each_with_index do |phase, index|
        if index > 0
          assert_equal(4, phase.items.count)
          assert_equal(0, phase.add_invoice_items.count)

          items = phase.items.filter {|i| i[:quantity].present? }
          # there should be two licensed products now
          assert_equal(2, items.count)
          assert_equal(1, T.must(items.first)[:quantity])
        end
      end

      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(StripeProrationBehavior::CREATE_PRORATIONS.serialize, first_phase.proration_behavior)
      assert_equal("1", first_phase.metadata["salesforce_segment_index"])
      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal(sf_order.Id, first_phase.metadata['salesforce_order_id'])

      second_phase = T.must(subscription_schedule.phases.second)
      assert_equal(StripeProrationBehavior::NONE.serialize, second_phase.proration_behavior)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)
      assert(second_phase.start_date.to_i - initial_order_start_date.to_i < 1.day.to_i)
      assert_equal("1", second_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, second_phase.metadata['salesforce_order_id'])

      third_phase = T.must(subscription_schedule.phases.third)
      assert_equal(StripeProrationBehavior::NONE.serialize, third_phase.proration_behavior)
      assert_equal(second_phase.end_date.to_i, third_phase.start_date.to_i)
      assert_equal(initial_order_end_date.to_i, third_phase.end_date.to_i)
      assert_equal("2", third_phase.metadata["salesforce_segment_index"])

      # Note the amendment order Id will override the initial order Id for future phases
      assert_equal(sf_order_amendment.Id, third_phase.metadata['salesforce_order_id'])
    end

    it 'amend a future phase of a salesforce order with a mdq product' do
      # setup
      contract_term = 36
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date + 1.year
      amendment_term = 24

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_future_day",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})
      segmented_metered_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_metered_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products and each have three segments
      assert_equal(6, quote_with_product["lineItems"].count)

      # create and sync the initial order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # initial order should create sub schedule with two phases - one phase for each segment
      assert_equal(3, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(2, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote
      sf_contract = create_contract_from_order(sf_order)

      # Workaround #1: Leverage the Amendment Start Date on the contract to determine the start of the segment you are editing
      # and to prevent the error "First Segment must fall between Quote Start and Quote End Date error"
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      # By setting the SBQQ__AmendmentStartDate__c on the contract, the amendment quote only has the current and future segments available for editing
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      assert_equal(4, amendment_data["lineItems"].count)

      amendment_data["lineItems"].each do |line_item|
        # let's specifically increase the quantity of metered mdq product in the second segment
        if line_item["record"][CPQ_QUOTE_LINE_PRODUCT] == segmented_metered_product && line_item["record"][CPQ_ORDER_ITEM_SEGMENT_INDEX] == 1
          line_item["record"][CPQ_QUOTE_QUANTITY] = 2
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # assert only the affected order items are in the amendment order across the segments
      sf_order_amendment = create_order_from_quote_data(amendment_data, activate_order: false)
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(1, order_lines.count)

      # Workaround #2: Null out the three MDQ fields on the Order Product BEFORE activating in order to activate a Salesforce Order with partially segmented order items.
      # When I null those fields out, the order activates just fine. Pushing through to contracts, renewals, further amendments, everything looks as it should and the Quote Lines follow MDQ segmentation.
      # otherwise this will error with: Restforce::ErrorCode::FieldCustomValidationException: FIELD_CUSTOM_VALIDATION_EXCEPTION: Can't activate an order with partially segmented order items.
      sf.update!(SF_ORDER_ITEM, {
        SF_ID => order_lines.first.Id,
        CPQ_ORDER_ITEM_SEGMENT_INDEX => nil,
        CPQ_ORDER_ITEM_SEGMENT_KEY => nil,
      })

      sf.update!(SF_ORDER,
        SF_ID => sf_order_amendment.Id,
        'Status' => 'Activated'
      )

      # sync the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # number of phases should not increase since we are amending a future phase
      assert_equal(3, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(StripeProrationBehavior::CREATE_PRORATIONS.serialize, first_phase.proration_behavior)
      assert_equal("1", first_phase.metadata["salesforce_segment_index"])
      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal(amendment_start_date.to_i, first_phase.end_date.to_i)
      assert_equal(sf_order.Id, first_phase.metadata['salesforce_order_id'])
      first_phase_items = first_phase.items
      assert_equal(2, first_phase_items.count)

      second_phase = T.must(subscription_schedule.phases.second)
      assert_equal(StripeProrationBehavior::NONE.serialize, second_phase.proration_behavior)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 2.year).to_i, second_phase.end_date.to_i)
      assert_equal("2", second_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, second_phase.metadata['salesforce_order_id'])
      # we increased quantity in this phase so this should be 3
      second_phase_items = second_phase.items
      assert_equal(3, second_phase_items.count)

      third_phase = T.must(subscription_schedule.phases.third)
      assert_equal(StripeProrationBehavior::NONE.serialize, second_phase.proration_behavior)
      assert_equal(second_phase.end_date.to_i, third_phase.start_date.to_i)
      assert_equal(initial_order_end_date.to_i, third_phase.end_date.to_i)
      assert_equal("3", third_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order.Id, third_phase.metadata['salesforce_order_id'])
      third_phase_items = third_phase.items
      assert_equal(2, third_phase_items.count)
    end

    it 'amend all future phases of a salesforce order with mdq product' do
      # setup
      contract_term = 36
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date + 1.year
      amendment_term = 24

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_all_future_phases",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})
      segmented_metered_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_metered_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products and each have three segments
      assert_equal(6, quote_with_product["lineItems"].count)

      # create and sync the initial order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # initial order should create sub schedule with two phases - one phase for each segment
      assert_equal(3, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(2, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote
      sf_contract = create_contract_from_order(sf_order)

      # Workaround #1: Leverage the Amendment Start Date on the contract to determine the start of the segment you are editing
      # and to prevent the error "First Segment must fall between Quote Start and Quote End Date error"
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      # By setting the SBQQ__AmendmentStartDate__c on the contract, the amendment quote only has the current and future segments available for editing
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      assert_equal(4, amendment_data["lineItems"].count)

      amendment_data["lineItems"].each do |line_item|
        # let's specifically increase the quantity of metered mdq product in the second segment
        if line_item["record"][CPQ_QUOTE_LINE_PRODUCT] == segmented_licensed_product
          line_item["record"][CPQ_QUOTE_QUANTITY] = 3
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # assert only the affected order items are in the amendment order across the segments
      sf_order_amendment = create_order_from_quote_data(amendment_data, activate_order: false)
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(2, order_lines.count)

      # Workaround #2: Null out the three MDQ fields on the Order Product BEFORE activating in order to activate a Salesforce Order with partially segmented order items.
      # When I null those fields out, the order activates just fine. Pushing through to contracts, renewals, further amendments, everything looks as it should and the Quote Lines follow MDQ segmentation.
      # otherwise this will error with: Restforce::ErrorCode::FieldCustomValidationException: FIELD_CUSTOM_VALIDATION_EXCEPTION: Can't activate an order with partially segmented order items.
      order_lines.each do |order_line|
        sf.update!(SF_ORDER_ITEM, {
          SF_ID => order_line.Id,
          CPQ_ORDER_ITEM_SEGMENT_INDEX => nil,
          CPQ_ORDER_ITEM_SEGMENT_KEY => nil,
        })
      end

      sf.update!(SF_ORDER,
        SF_ID => sf_order_amendment.Id,
        'Status' => 'Activated'
      )

      # sync the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # number of phases should not increase since we are amending a future phase
      assert_equal(3, subscription_schedule.phases.count)
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(StripeProrationBehavior::CREATE_PRORATIONS.serialize, first_phase.proration_behavior)
      assert_equal("1", first_phase.metadata["salesforce_segment_index"])
      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal(amendment_start_date.to_i, first_phase.end_date.to_i)
      assert_equal(sf_order.Id, first_phase.metadata['salesforce_order_id'])
      first_phase_items = first_phase.items
      assert_equal(2, first_phase_items.count)

      second_phase = T.must(subscription_schedule.phases.second)
      assert_equal(StripeProrationBehavior::NONE.serialize, second_phase.proration_behavior)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 2.year).to_i, second_phase.end_date.to_i)
      assert_equal("2", second_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, second_phase.metadata['salesforce_order_id'])
      # we increased quantity in this phase so this should be 4
      second_phase_items = second_phase.items
      assert_equal(4, second_phase_items.count)

      third_phase = T.must(subscription_schedule.phases.third)
      assert_equal(StripeProrationBehavior::NONE.serialize, second_phase.proration_behavior)
      assert_equal(second_phase.end_date.to_i, third_phase.start_date.to_i)
      assert_equal(initial_order_end_date.to_i, third_phase.end_date.to_i)
      assert_equal("3", third_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, third_phase.metadata['salesforce_order_id'])
      third_phase_items = third_phase.items
      assert_equal(4, third_phase_items.count)
    end

    it 'add a non-segmented product in a future phase of a salesforce order' do
      # setup
      contract_term = 12
      initial_order_start_date = now_time
      _initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date + 6.months
      amendment_term = 6

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_non_segmented_item_future_phase_1_2",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create and add a mdq licensed to quote
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      # one product with four segments
      assert_equal(4, quote_with_product["lineItems"].count)

      # create and sync the initial order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # initial order should create sub schedule with two phases - one phase for each segment
      assert_equal(4, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(1, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote
      sf_contract = create_contract_from_order(sf_order)

      # Workaround #1: By setting the SBQQ__AmendmentStartDate__c on the contract, the amendment quote only has the current and future segments available for editing
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      assert_equal(2, amendment_data["lineItems"].count)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term
      quote_id = calculate_and_save_cpq_quote(amendment_data)

      # let's add a non segmented product to the third phase
      # add a non segmented metered product to quote
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_metered_produce_with_price(price_in_cents: 18_00)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

      # assert only the affected order items are in the amendment order across the segments
      sf_order_amendment = create_order_from_quote_data(quote_with_product)
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(1, order_lines.count)
      assert_equal(sf_product_id, order_lines.first["Product2Id"])

      # sync the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # number of phases should not increase since we are amending a future phase
      assert_equal(4, subscription_schedule.phases.count)

      # sanity check first phase has not changed
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(StripeProrationBehavior::CREATE_PRORATIONS.serialize, first_phase.proration_behavior)
      assert_equal("1", first_phase.metadata["salesforce_segment_index"])
      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 3.months).to_i, first_phase.end_date.to_i)
      assert_equal(sf_order.Id, first_phase.metadata['salesforce_order_id'])
      first_phase_items = first_phase.items
      assert_equal(1, first_phase_items.count)

      third_phase = T.must(subscription_schedule.phases.third)
      assert_equal(StripeProrationBehavior::NONE.serialize, third_phase.proration_behavior)
      assert_equal((initial_order_start_date + 6.months).to_i, third_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 9.months).to_i, third_phase.end_date.to_i)
      assert_equal("3", third_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, third_phase.metadata['salesforce_order_id'])
      # verify the additional item was added
      assert_equal(2, third_phase.items.count)

      last_phase = T.must(subscription_schedule.phases.last)
      assert_equal("4", last_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order_amendment.Id, third_phase.metadata['salesforce_order_id'])
      # verify the additional item was added
      assert_equal(2, last_phase.items.count)

      # TODO verify invoice amounts
    end

    it 'remove a non-segmented product in a future phase of a salesforce order' do
      # setup
      contract_term = 12
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + 3.months
      amendment_term = 9

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_remove_non_segmented_item_4",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: 18_00)

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products of which one products has four segments
      assert_equal(5, quote_with_product["lineItems"].count)

      # create the initial order
      sf_order = create_order_from_cpq_quote(quote_id)

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # segmented product creates sub schedule with two phases
      assert_equal(4, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(2, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(2, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote to remove the non segmented product
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].each do |line_item|
        if line_item["record"][CPQ_ORDER_ITEM_SEGMENT_KEY].nil?
          line_item["record"][CPQ_QUOTE_QUANTITY] = 0
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      sf_order_amendment = create_order_from_quote_data(amendment_data)
      # sanity check the amendment touches only the non segmented line
      amendment_order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(1, amendment_order_lines.count)

      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # refetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(4, subscription_schedule.phases.count)
      subscription_schedule.phases.each_with_index do |phase, index|
        if index == 0
          assert_equal(2, phase.items.count)
          assert_equal(0, phase.add_invoice_items.count)

          items = phase.items.filter {|i| i[:quantity].present? }
          assert_equal(2, items.count)
        else
          assert_equal(1, phase.items.count)
          assert_equal(0, phase.add_invoice_items.count)

          items = phase.items.filter {|i| i[:quantity].present? }
          assert_equal(1, items.count)
        end
      end
    end
  end

  describe 'error handling' do
    it 'throws error if attempting to amend phase on non-segmented start date' do
      # setup
      contract_term = 48
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + 2.years + 1.day
      amendment_term = 24

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_error_non_segment",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products and each have three segments
      assert_equal(4, quote_with_product["lineItems"].count)

      # create and sync the initial order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # initial order should create sub schedule with two phases - one phase for each segment
      assert_equal(4, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(1, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote
      # Workaround #1: Leverage the Amendment Start Date on the contract to determine the start of the segment you are editing
      # and to prevent the error "First Segment must fall between Quote Start and Quote End Date error"
      sf_contract = create_contract_from_order(sf_order)
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      # By setting the SBQQ__AmendmentStartDate__c on the contract, the amendment quote only has the current and future segments available for editing
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      assert_equal(2, amendment_data["lineItems"].count)

      amendment_data["lineItems"].each do |line_item|
        # let's specifically increase the quantity of metered mdq product in the second segment
        if line_item["record"][CPQ_QUOTE_LINE_PRODUCT] == segmented_licensed_product
          line_item["record"][CPQ_QUOTE_QUANTITY] = 2
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date + 1.day)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # assert only the affected order items are in the amendment order across the segments
      sf_order_amendment = create_order_from_quote_data(amendment_data, activate_order: false)
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(2, order_lines.count)

      # Workaround #2: Null out the three MDQ fields on the Order Product BEFORE activating in order to activate a Salesforce Order with partially segmented order items.
      # When I null those fields out, the order activates just fine. Pushing through to contracts, renewals, further amendments, everything looks as it should and the Quote Lines follow MDQ segmentation.
      # otherwise this will error with: Restforce::ErrorCode::FieldCustomValidationException: FIELD_CUSTOM_VALIDATION_EXCEPTION: Can't activate an order with partially segmented order items.
      order_lines.each do |order_line|
        sf.update!(SF_ORDER_ITEM, {
          SF_ID => order_line.Id,
          CPQ_ORDER_ITEM_SEGMENT_INDEX => nil,
          CPQ_ORDER_ITEM_SEGMENT_KEY => nil,
        })
      end

      sf.update!(SF_ORDER,
        SF_ID => sf_order_amendment.Id,
        'Status' => 'Activated'
      )

      # sync the amendment order
      exception = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("Amendment order segment dates do not match initial order segment dates.", exception.message)
    end

    it 'throws error if contract has different date than amendment' do
      # setup
      contract_term = 36
      initial_order_start_date = now_time
      amendment_start_date = initial_order_start_date + 1.year
      amendment_term = 24

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_error_contract_date_2",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})
      # add mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      # two products and each have three segments
      assert_equal(3, quote_with_product["lineItems"].count)

      # create and sync the initial order
      sf_order = create_order_from_cpq_quote(quote_id)

      # create the amendment quote
      sf_contract = create_contract_from_order(sf_order)

      # Workaround #1: Leverage the Amendment Start Date on the contract to determine the start of the segment you are editing
      # and to prevent the error "First Segment must fall between Quote Start and Quote End Date error"
      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date - 1.day),
      })
      sf_contract.refresh

      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].each do |line_item|
        # let's specifically increase the quantity of metered mdq product in the second segment
        if line_item["record"][CPQ_ORDER_ITEM_SEGMENT_INDEX] == 1
          line_item["record"][CPQ_QUOTE_QUANTITY] = 2
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      # assert only the affected order items are in the amendment order across the segments
      sf_order_amendment = create_order_from_quote_data(amendment_data, activate_order: false)
      order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(1, order_lines.count)

      # Workaround #2: Null out the three MDQ fields on the Order Product BEFORE activating in order to activate a Salesforce Order with partially segmented order items.
      # When I null those fields out, the order activates just fine. Pushing through to contracts, renewals, further amendments, everything looks as it should and the Quote Lines follow MDQ segmentation.
      # otherwise this will error with: Restforce::ErrorCode::FieldCustomValidationException: FIELD_CUSTOM_VALIDATION_EXCEPTION: Can't activate an order with partially segmented order items.
      sf.update!(SF_ORDER_ITEM, {
        SF_ID => order_lines.first.Id,
        CPQ_ORDER_ITEM_SEGMENT_INDEX => nil,
        CPQ_ORDER_ITEM_SEGMENT_KEY => nil,
      })

      sf.update!(SF_ORDER,
        SF_ID => sf_order_amendment.Id,
        'Status' => 'Activated'
      )

      # sync the amendment order
      exception = assert_raises(StripeForce::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end
      assert_match("Start date on amendment order does not match the corresponding contract's amendment start date field.", exception.message)
    end
  end

  describe 'terminate orders with mdq' do
    it 'same day termination a salesforce order with a mdq product and quarterly segments' do
      # setup
      contract_term = 24
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date
      amendment_term = contract_term

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_same_day_termination_7",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})
      segmented_metered_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_metered_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products of which two products have four segments
      assert_equal(16, quote_with_product["lineItems"].count)

      # create the initial order
      sf_order = create_order_from_cpq_quote(quote_id)

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # segmented product creates sub schedule with two phases
      assert_equal(8, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(2, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # let's sanity check the first two phases have the right dates
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal("1", first_phase.metadata["salesforce_segment_index"])
      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 3.months).to_i, first_phase.end_date.to_i)
      assert_equal(sf_order.Id, first_phase.metadata['salesforce_order_id'])

      second_phase = T.must(subscription_schedule.phases.second)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)
      assert_equal((initial_order_start_date + 6.months).to_i, second_phase.end_date.to_i)
      assert_equal("2", second_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order.Id, second_phase.metadata['salesforce_order_id'])

      # ... bunch of middle phases each representing a quarter

      # check the last phase
      last_phase = T.must(subscription_schedule.phases.last)
      assert_equal((initial_order_end_date - 3.months).to_i, last_phase.start_date.to_i)
      assert_equal(initial_order_end_date.to_i, last_phase.end_date.to_i)
      assert_equal("8", last_phase.metadata["salesforce_segment_index"])
      assert_equal(sf_order.Id, second_phase.metadata['salesforce_order_id'])

      # create the amendment quote to terminate the entire order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].each do |line_item|
          line_item["record"][CPQ_QUOTE_QUANTITY] = 0
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # sanity check the amendment touches all the lines
      amendment_order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(16, amendment_order_lines.count)

      # since this is a full same day termination
      assert_equal(sf_order["TotalAmount"], -1 * sf_order_amendment["TotalAmount"])
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # refetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # ensure the subscription schedule was cancelled
      assert_equal(8, subscription_schedule.phases.count)
      assert_equal("canceled", subscription_schedule.status)
    end

    it 'partial termination of a salesforce order with mdq products' do
      # setup
      contract_term = 24
      initial_order_start_date = now_time
      _initial_order_end_date = initial_order_start_date + contract_term.months
      amendment_start_date = initial_order_start_date
      amendment_term = contract_term

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "test_mdq_partial_termination_3",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq licensed and metered product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})
      segmented_metered_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::QUARTER.serialize})

      # add both mdq products to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_metered_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # two products of which two products have four segments
      assert_equal(16, quote_with_product["lineItems"].count)

      # create the initial order
      sf_order = create_order_from_cpq_quote(quote_id)

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # segmented product creates sub schedule with two phases
      assert_equal(8, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(2, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        # assert the quantity is 1 for the non-metered product
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      # create the amendment quote to terminate one item from the order
      item_segment_key = quote_with_product["lineItems"].first["record"][CPQ_ORDER_ITEM_SEGMENT_KEY]

      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].each do |line_item|
        if line_item["record"][CPQ_ORDER_ITEM_SEGMENT_KEY] == item_segment_key
          line_item["record"][CPQ_QUOTE_QUANTITY] = 0
        end
      end
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      sf.update!(SF_CONTRACT, {
        SF_ID => sf_contract.Id,
        CONTRACT_AMENDMENT_START_DATE => format_date_for_salesforce(amendment_start_date),
      })
      sf_contract.refresh

      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # sanity check the amendment touches half the order items
      amendment_order_lines = sf_get_related(sf_order_amendment, SF_ORDER_ITEM)
      assert_equal(8, amendment_order_lines.count)
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      # refetch the subscription schedule
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # segmented product creates sub schedule with two phases
      assert_equal(9, subscription_schedule.phases.count)
      subscription_schedule.phases.each_with_index do |phase, index|
        if index == 0
          assert_equal(2, phase.items.count)
          next
        end
        assert_equal(1, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)
      end
    end

    it 'order item credit on partial termination' do end

    # TODO this is not currently support, but question is should we support this? This is essentially "shortening" a ramp deal.
    # it 'future termination of a salesforce order with a mdq product' do end
  end

  describe 'custom segments with mdq' do
    it 'sync a salesforce order amendment with a custom segment mdq product' do end
  end

  describe 'backdated amendment orders with mdq' do
     it 'sync a backdated amendment salesforce order with a mdq product' do end
  end
end