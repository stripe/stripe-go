# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::BackendProratedAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'handles an initial order with a backend proration' do
    backend_prorated_term = 6
    contract_term = TEST_DEFAULT_CONTRACT_TERM + backend_prorated_term
    initial_start_date = now_time
    end_date = initial_start_date + contract_term.months

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      }
    )

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    # although there is just a single order, it is split into two pieces
    assert_equal(2, subscription_schedule.phases.count)

    first_phase = T.must(subscription_schedule.phases.first)
    second_phase = T.must(subscription_schedule.phases[1])

    # second phase proration none is very important!
    assert_equal("create_prorations", first_phase.proration_behavior)
    assert_equal("none", second_phase.proration_behavior)

    # first phase should start now and end at the end of the billing cycle
    assert_equal(0, first_phase.start_date - initial_start_date.to_i)
    assert_equal(0, first_phase.end_date - (initial_start_date + TEST_DEFAULT_CONTRACT_TERM.months).to_i)

    # second phase should start at the end of the billing cycle and end at the contract term
    assert_equal(0, second_phase.start_date - (initial_start_date + TEST_DEFAULT_CONTRACT_TERM.months).to_i)
    assert_equal(0, second_phase.end_date - end_date.to_i)

    # first phase should have one item with no quantity, since it is a metered product
    assert_equal(1, first_phase.items.count)
    assert_equal(0, first_phase.add_invoice_items.count)
    first_phase_item = T.must(first_phase.items.first)
    assert_equal(1, first_phase_item[:quantity])

    assert_equal(1, second_phase.items.count)
    assert_equal(1, second_phase.add_invoice_items.count)
    second_phase_item = T.must(second_phase.items.first)
    prorated_phase_item = T.must(second_phase.add_invoice_items.first)
    assert_equal(1, first_phase_item.quantity)
    assert_equal(1, prorated_phase_item.quantity)

    second_item_price = Stripe::Price.retrieve(T.cast(second_phase_item.price, String), @user.stripe_credentials)
    assert_equal(TEST_DEFAULT_PRICE, second_item_price.unit_amount_decimal.to_i)
    assert_equal("month", second_item_price.recurring.interval)
    assert_equal(12, second_item_price.recurring.interval_count)

    prorated_price = Stripe::Price.retrieve(T.cast(prorated_phase_item.price, String), @user.stripe_credentials)
    assert_equal(TEST_DEFAULT_PRICE / 2, prorated_price.unit_amount_decimal.to_i)
    assert_equal("one_time", prorated_price.type)
    assert_equal("true", prorated_price.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
  end

  describe 'amendments' do
    it 'backend prorated order amendment which starts on the same day as the proration phase' do
      contract_term = 18
      amendment_term = 6

      # 18 month contract, billed annually
      # amendment starts 12 months in, same day as the proration phase
      initial_start_date = now_time
      amendment_start_date = initial_start_date + (contract_term - amendment_term).months
      amendment_end_date = amendment_start_date + amendment_term.months

      sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      # create initial order
      sf_order = create_salesforce_order(
        sf_product_id: sf_product_id,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_start_date),
        }
      )

      # translate the initial order
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # create amendment order
      sf_contract = create_contract_from_order(sf_order)
      amendment_data = create_quote_data_from_contract_amendment(sf_contract)
      amendment_data["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

      calculate_and_save_cpq_quote(amendment_data)
      sf_order_amendment = create_order_from_quote_data(amendment_data)

      # translate the amendment order
      StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # although the initial order had two phases,
      # the proration phase from the initial order is combined with the prorated order
      assert_equal(2, subscription_schedule.phases.count)

      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases[1])

      assert_equal("create_prorations", first_phase.proration_behavior)
      # second phase proration none is very important!
      assert_equal("none", second_phase.proration_behavior)

      # first phase should start now and end at the end of the billing cycle
      # which is the same as the amendment start date
      assert_equal(0, first_phase.start_date - initial_start_date.to_i)
      assert_equal(0, first_phase.end_date - amendment_start_date.to_i)

      # second phase should start at the end of the billing cycle and end at the contract term
      assert_equal(0, second_phase.start_date - amendment_start_date.to_i)
      assert_equal(0, second_phase.end_date - amendment_end_date.to_i)

      # first phase should have one item
      assert_equal(1, first_phase.items.count)
      assert_equal(0, first_phase.add_invoice_items.count)
      first_phase_item = T.must(first_phase.items.first)
      assert_equal(1, first_phase_item[:quantity])

      # two add-hoc items: one for the amendment, one for the initial order
      assert_equal(2, second_phase.items.count)
      assert_equal(2, second_phase.add_invoice_items.count)
      second_phase_item = T.must(second_phase.items.first)
      second_phase_quantity_addition = T.must(second_phase.items[1])
      prorated_phase_item = T.must(second_phase.add_invoice_items.detect {|i| i[:metadata][StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::BACKEND_PRORATION)] != "true" })
      backend_proration_item = T.must(second_phase.add_invoice_items.detect {|i| i[:metadata][StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::BACKEND_PRORATION)] == "true" })
      assert_equal(1, second_phase_item.quantity)
      assert_equal(1, second_phase_quantity_addition.quantity)
      assert_equal(1, prorated_phase_item.quantity)
      assert_equal(1, backend_proration_item.quantity)

      second_item_price = Stripe::Price.retrieve(T.cast(second_phase_item.price, String), @user.stripe_credentials)
      assert_equal(TEST_DEFAULT_PRICE, second_item_price.unit_amount_decimal.to_i)
      assert_equal("month", second_item_price.recurring.interval)
      assert_equal(12, second_item_price.recurring.interval_count)

      second_phase_quantity_addition_price = Stripe::Price.retrieve(T.cast(second_phase_quantity_addition.price, String), @user.stripe_credentials)
      assert_equal(TEST_DEFAULT_PRICE, second_phase_quantity_addition_price.unit_amount_decimal.to_i)
      assert_equal("month", second_phase_quantity_addition_price.recurring.interval)
      assert_equal(12, second_phase_quantity_addition_price.recurring.interval_count)

      prorated_price = Stripe::Price.retrieve(T.cast(prorated_phase_item.price, String), @user.stripe_credentials)
      assert_equal(TEST_DEFAULT_PRICE / 2, prorated_price.unit_amount_decimal.to_i)
      assert_equal("one_time", prorated_price.type)
      assert_equal("true", prorated_price.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])

      backend_proration_price = Stripe::Price.retrieve(T.cast(backend_proration_item.price, String), @user.stripe_credentials)
      assert_equal(TEST_DEFAULT_PRICE / 2, backend_proration_price.unit_amount_decimal.to_i)
      assert_equal("one_time", backend_proration_price.type)
      assert_equal("true", backend_proration_price.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION)])
    end

    it 'backend prorated order amendment which starts after the proration phase start date'
    it 'backend prorated order amendment which starts before the proration phase start date'
    it 'multiple backend prorated order amendments'
    it 'multiple backend prorated order amendments which start after the proration phase start date'
    it 'multiple backend prorated order amendments which are then partially terminated'
    it 'multiple backend prorated order amendments which are then fully terminated'
  end
end
