# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::ProductTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
    @user.enable_feature(FeatureFlags::COUPONS)
  end

  describe 'mdq products' do
    it 'sync a salesforce order with a mdq licensed product' do
      # setup
      contract_term = 36
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "order_with_mdq_licensed_product_0_1",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(initial_order_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
        }
      )

      # create a mdq product
      segmented_product_id = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})
      # add mdq product to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_product_id)

      numnber_of_segments = contract_term / MONTHS_IN_YEAR
      assert_equal(numnber_of_segments, quote_with_product["lineItems"].count)

      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # create and translate order
      sf_order = create_order_from_cpq_quote(quote_id)
      sf_order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(sf_order.Id, subscription_schedule.metadata['salesforce_order_id'])

      assert_equal(3, subscription_schedule.phases.count)

      subscription_schedule.phases.each_with_index do |phase, index|
        assert_equal(1, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)
        item = T.must(phase.items.first)
        assert_equal(1, item[:quantity])

        # check that the same order item was added to the correct phase
        order_item = T.must(sf_order_items.detect {|i| i.Id == item.metadata['salesforce_order_item_id'] })
        assert_equal(order_item[CPQ_ORDER_ITEM_SEGMENT_INDEX], index + 1)
      end

      # sanity check the phase durations
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)

      assert_equal(first_phase.start_date.to_i, initial_order_start_date.to_i)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)

      third_phase = T.must(subscription_schedule.phases.third)
      assert_equal(third_phase.end_date.to_i, initial_order_end_date.utc.beginning_of_day.to_i)
    end

    it 'sync a salesforce order with two mdq products and non-segmented product' do
      # setup
      contract_term = 24
      initial_order_start_date = now_time
      initial_order_end_date = initial_order_start_date + contract_term.months

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "order_with_mdq_and_non_mdq_product_3",
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
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # add a non segmented metered product to quote
      sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_metered_produce_with_price
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # three products and two products with two segments
      assert_equal(5, quote_with_product["lineItems"].count)

      # create and translate order
      sf_order = create_order_from_cpq_quote(quote_id)

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(2, subscription_schedule.phases.count)
      subscription_schedule.phases.each do |phase|
        assert_equal(3, phase.items.count)
        assert_equal(0, phase.add_invoice_items.count)

        items = phase.items.filter {|i| i[:quantity].present? }
        assert_equal(1, items.count)
        assert_equal(1, T.must(items.first)[:quantity])
      end

      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)

      assert_equal(initial_order_start_date.to_i, first_phase.start_date.to_i)
      assert_equal(first_phase.end_date.to_i, second_phase.start_date.to_i)

      assert_equal((initial_order_start_date + 1.year).to_i, second_phase.start_date.to_i)
      assert_equal(second_phase.end_date.to_i, initial_order_end_date.to_i)
    end

    it 'sync a salesforce order with prebilling and mdq products' do
      # prebill one year of an annually billed two year subscription
      start_date = now_time + 3.months
      subscription_term = 24
      end_date = start_date + subscription_term.months
      prebill_iterations = 3

      @user.field_defaults = {
        "subscription_schedule" => {
          "prebilling.iterations" => prebill_iterations,
        },
      }
      @user.save

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "order_with_mdq_licensed_product_and_prebilling_2",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
        }
      )

      # create a mdq licensed product for $120 every quarter
      segmented_licensed_product = create_salesforce_cpq_segmented_product(
        additional_product_fields: {CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::QUARTERLY.serialize},
        additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add product to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # one product with two segments
      assert_equal(2, quote_with_product["lineItems"].count)

      # create and translate order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # ensure a prebilling invoice has been created
      first_phase = T.must(subscription_schedule.phases.first)
      second_phase = T.must(subscription_schedule.phases.second)

      # we are essentially prebilling for the first phase / segment
      assert_equal(start_date.to_i, first_phase.start_date.to_i)
      assert_equal(first_phase.start_date.to_i, subscription_schedule.prebilling.period_start)

      assert_equal((start_date + 1.year).to_i, second_phase.start_date.to_i)
      # prebilling 3 quarters
      assert_equal((start_date + (3 * 3).months).utc.beginning_of_day.to_i, subscription_schedule.prebilling.period_end)
      assert_equal(end_date.to_i, second_phase.end_date.to_i)

      invoice = Stripe::Invoice.retrieve(subscription_schedule.prebilling.invoice, @user.stripe_credentials)
      # one line per pre-billed period
      assert_equal(prebill_iterations, invoice.lines.data.length)
      assert_equal((TEST_DEFAULT_PRICE / 4) * prebill_iterations, invoice.total)
    end

    it 'sync a salesforce order with a mdq with variable discounts' do
      # prebill one year of an annually billed two year subscription
      start_date = now_time
      subscription_term = 24

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "order_with_mdq_licensed_product_and_discounts_3",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
        }
      )

      # create a mdq licensed product
      segmented_licensed_product = create_salesforce_cpq_segmented_product(additional_price_dimension_fields: {CPQ_PRICE_DIMENSION_TYPE => CPQPriceDimensionTypeOptions::YEAR.serialize})

      # add product to quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: segmented_licensed_product)
      calculate_and_save_cpq_quote(quote_with_product)

      # one product with two segments
      assert_equal(2, quote_with_product["lineItems"].count)

      # retrieve the quote line
      quote_lines = sf_get_related(quote_id, CPQ_QUOTE_LINE)
      assert_equal(2, quote_lines.size)
      quote_line_id = quote_lines.first.Id

      # create two coupons and attach one to the quote and the other to the quote line
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => 'Twenty-five percent off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 25,
      })
      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$10 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 10,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
      # create the quote coupon association object to map the coupons to the quote
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # create and translate order
      sf_order = create_order_from_cpq_quote(quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
      sf_order.refresh

      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(2, subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)

      # the first phase item should have one coupon
      first_phase_item = T.must(first_phase.items.first)
      discounts = first_phase_item.discounts
      assert_equal(1, discounts.count)

      # retrieve the stripe coupon on the quote line
      sf_percent_off_coupon_id = Stripe::Coupon.retrieve(T.must(discounts.first).coupon, @user.stripe_credentials)
      # sanity check the stripe coupon has the right data
      assert_equal(25, sf_percent_off_coupon_id.percent_off)
      assert_equal("once", sf_percent_off_coupon_id.duration)

      # retrieve the stripe coupon on the quote
      second_phase = T.must(subscription_schedule.phases.second)
      assert_equal(1, second_phase.items.count)

      # fetch invoice and verify final amount due
      sf_account = sf_get(sf_account_id)
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)
      assert_equal(1, invoice.data.count)

      # expected price should equal $120 with 25% off and $10 off coupons applied
      assert_equal((((TEST_DEFAULT_PRICE / 100) * 12 * 0.75) - 10) * 100, invoice.data.first.total)
    end

    describe 'amending orders with mdq products' do
      it 'amend the current segment of a salesforce order with a mdq product' do end
      it 'amend a future segment of a salesforce order with a mdq product' do end
      it 'termimnate a salesforce order with a mdq product' do end
    end
  end
end
