# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::CouponTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || DateTime.now.utc)

    @user = make_user(save: true)
    @user.enable_feature(FeatureFlags::COUPONS)
  end

  describe 'salesforce trigger logic' do
    it 'validation error thrown if SF Stripe coupon is created with both Amount_Off__c and Percent_Off__c set' do
      exception = assert_raises(Restforce::ErrorCode::FieldCustomValidationException) do
        create_salesforce_stripe_coupon(additional_fields: {
          SalesforceStripeCouponFields::NAME => 'Invalid coupon',
          SalesforceStripeCouponFields::AMOUNT_OFF => 10,
          SalesforceStripeCouponFields::PERCENT_OFF => 25,
        })
      end
      assert_match("FIELD_CUSTOM_VALIDATION_EXCEPTION: Received both Percent_Off__c and Amount_Off__c parameters.", exception.message)
    end

    it 'order and order line coupons are created when an sf quote is ordered' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "order_and_order_line_when_quote_ordered_1", additional_quote_fields: {
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

      # create two coupons and attach both to the quote line
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
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # confirm the quote line coupon was created
      sf_quote_line_coupon = get_salesforce_stripe_coupons_associated_to_quote_line(quote_line_id: quote_line_id)
      assert_equal(1, sf_quote_line_coupon.size)
      assert_equal('Twenty-five percent off coupon', sf_quote_line_coupon.first[prefixed_stripe_field('Name__c')])

      # create and translate the sf order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # verify the coupon order or order item associations were created
      order_coupons = sf.query("SELECT #{SF_ID} FROM #{prefixed_stripe_field(ORDER_SF_STRIPE_COUPON)} WHERE #{prefixed_stripe_field('Order__c')} = '#{sf_order.Id}'")
      assert_equal(1, order_coupons.count)

      # sanity check the order coupon info is the same as the quote coupon
      order_coupon = sf.find(prefixed_stripe_field(ORDER_SF_STRIPE_COUPON), order_coupons.first.Id)
      assert_equal('$10 off coupon', order_coupon[prefixed_stripe_field('Name__c')])
      assert_equal(10, order_coupon[prefixed_stripe_field('Amount_Off__c')])
      assert_equal(sf_amount_off_coupon_id, order_coupon[prefixed_stripe_field('Quote_Stripe_Coupon_Id__c')])
      assert_nil(order_coupon[prefixed_stripe_field('Percent_Off__c')])
      assert_nil(order_coupon[prefixed_stripe_field('Max_Redemptions__c')])
    end
  end

  describe 'order translation with coupons' do
    it 'translate sf order with multiple coupons on an order line' do
      # add a custom metadata field mapping
      @user.field_mappings['coupon'] = {
        'metadata.sf_coupon_mapped_metadata_field' => prefixed_stripe_field('Name__c'),
        'metadata.sf_order_mapped_metadata_field' => prefixed_stripe_field('Order_Item__c.OrderId'),
      }
      @user.save

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "translate_multiple_coupons_order_line", additional_quote_fields: {
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

      # create two coupons and attach both to the quote line
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
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)

      # the first phase item should have two coupons
      first_phase_item = T.must(first_phase.items.first)
      discounts = first_phase_item.discounts
      assert_equal(2, discounts.count)

      # retrieve the two stripe coupons
      stripe_coupon_1 = Stripe::Coupon.retrieve(T.must(discounts.first).coupon, @user.stripe_credentials)
      stripe_coupon_2 = Stripe::Coupon.retrieve(T.must(discounts[1]).coupon, @user.stripe_credentials)

      # sanity check both stripe coupons have the right data
      stripe_amount_off_coupon = stripe_coupon_1.name == '$10 off coupon' ? stripe_coupon_1 : stripe_coupon_2
      assert_equal(10 * 100, stripe_amount_off_coupon.amount_off)
      assert_equal("usd", stripe_amount_off_coupon.currency)
      assert_equal("once", stripe_amount_off_coupon.duration)

      sf_amount_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_amount_off_coupon_id)
      assert_equal(sf_amount_off_order_coupon.Id, stripe_amount_off_coupon.metadata['salesforce_order_stripe_coupon_id'])
      assert_equal('$10 off coupon', stripe_amount_off_coupon.metadata['sf_coupon_mapped_metadata_field'])
      assert_equal(sf_order.Id, stripe_amount_off_coupon.metadata['sf_order_mapped_metadata_field'])

      stripe_percent_off_coupon = stripe_coupon_1.name == '25% off coupon' ? stripe_coupon_1 : stripe_coupon_2
      assert_equal(25, stripe_percent_off_coupon.percent_off)
      assert_equal("once", stripe_percent_off_coupon.duration)

      sf_percent_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_percent_off_coupon_id)
      assert_equal(sf_percent_off_order_coupon.Id, stripe_percent_off_coupon.metadata['salesforce_order_stripe_coupon_id'])
      assert_equal('25% off coupon', stripe_percent_off_coupon.metadata['sf_coupon_mapped_metadata_field'])
      assert_equal(sf_order.Id, stripe_percent_off_coupon.metadata['sf_order_mapped_metadata_field'])

      # confirm the stripe coupon ids were written back to the quote coupons in salesforce
      sf_amount_off_coupon = sf_get(sf_amount_off_coupon_id)
      assert_equal(stripe_amount_off_coupon.id, sf_amount_off_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)])

      sf_percent_off_coupon = sf_get(sf_percent_off_coupon_id)
      assert_equal(stripe_percent_off_coupon.id, sf_percent_off_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end

    it 'translate an sf order with coupons on both the order and order items' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "translate_coupons_on_order_and_items", additional_quote_fields: {
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

      # create two coupons
      sf_percent_off_quote_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '25% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 25,
      })
      sf_amount_off_quote_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$50 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 50,
      })

      # create the quote line coupon association object to map the coupon to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_quote_coupon_id)
      # create the quote coupon association object to map the coupon to the quote
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_quote_id, sf_stripe_coupon_id: sf_amount_off_quote_coupon_id)

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # the first phase should have a coupon
      first_phase = T.must(subscription_schedule.phases.first)
      phase_discount = first_phase.discounts
      assert_equal(1, phase_discount.count)

      # the first phase item should have one coupons
      first_phase_item = T.must(first_phase.items.first)
      phase_item_discount = first_phase_item.discounts
      assert_equal(1, phase_item_discount.count)

      # retrieve the two stripe coupons
      stripe_phase_coupon = Stripe::Coupon.retrieve(T.must(phase_discount.first).coupon, @user.stripe_credentials)
      # sanity check the stripe coupons have the right data
      assert_equal(50 * 100, stripe_phase_coupon.amount_off)
      assert_equal("usd", stripe_phase_coupon.currency)
      assert_equal("once", stripe_phase_coupon.duration)
      sf_amount_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_amount_off_quote_coupon_id)
      assert_equal(sf_amount_off_order_coupon.Id, stripe_phase_coupon.metadata['salesforce_order_stripe_coupon_id'])

      stripe_phase_item_coupon = Stripe::Coupon.retrieve(T.must(phase_item_discount.first).coupon, @user.stripe_credentials)
      assert_equal(25, stripe_phase_item_coupon.percent_off)
      assert_equal("once", stripe_phase_item_coupon.duration)
      sf_percent_off_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_percent_off_quote_coupon_id)
      assert_equal(sf_percent_off_coupon.Id, stripe_phase_item_coupon.metadata['salesforce_order_stripe_coupon_id'])

      # fetch invoice and verify final amount due
      sf_account = sf_get(sf_account_id)
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)
      assert_equal(1, invoice.data.count)

      # expected price should equal $120 (original price) with 25%-off order line and then $50-off order coupon applied
      # subscription item discounts are applied before subscription discounts
      assert_equal(((TEST_DEFAULT_PRICE * 0.75) - (50 * 100)).to_i, invoice.data.first.amount_due)
    end

    it 'translate sf order with duration repeating coupon on an order line' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "translate_order_with_duration", additional_quote_fields: {
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

      # create a quote stripe coupon and attach to the quote line
      sf_quote_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => 'Holiday 25.55 coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 25.55,
        SalesforceStripeCouponFields::DURATION => 'repeating',
        SalesforceStripeCouponFields::DURATION_IN_MONTHS => 2,
        SalesforceStripeCouponFields::MAX_REDEMPTIONS => 3,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_quote_coupon_id)

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)

      # the first phase item should have two coupons
      first_phase_item = T.must(first_phase.items.first)
      discounts = first_phase_item.discounts
      assert_equal(1, discounts.count)

      # retrieve the two stripe coupons
      stripe_coupon = Stripe::Coupon.retrieve(T.must(discounts.first).coupon, @user.stripe_credentials)
      assert_in_delta(25.55, stripe_coupon.percent_off)
      assert_equal("repeating", stripe_coupon.duration)
      assert_equal(2, stripe_coupon.duration_in_months)
      assert_equal(3, stripe_coupon.max_redemptions)

      # confirm the stripe coupon metdata has the salesforce order coupon id
      sf_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_quote_coupon_id)
      assert_equal(sf_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])

      # confirm the stripe coupon ids were written back to the quote coupon in salesforce
      sf_quote_coupon = sf_get(sf_quote_coupon_id)
      assert_equal(stripe_coupon.id, sf_quote_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end
  end

  describe 'coupon reuse' do
    it 'uses the same quote stripe coupon if translated twice' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id_1, _sf_pricebook_id = salesforce_recurring_product_with_price
      sf_product_id_2, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "uses_same_quote_coupon", additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # add both products to the sf quote
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id_2)
      calculate_and_save_cpq_quote(quote_with_product)

      # retrieve the sf quote lines
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      assert_equal(2, quote_lines.size)

      # create an sf coupon
      sf_amount_off_coupon = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$55 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 55,
        SalesforceStripeCouponFields::DURATION => 'forever',
      })
      # create the quote line coupon association to map the coupon to the quote lines
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines.first.Id, sf_stripe_coupon_id: sf_amount_off_coupon)
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines[1].Id, sf_stripe_coupon_id: sf_amount_off_coupon)

      # create and translate the order
      sf_order = create_order_from_cpq_quote(sf_quote_id)

      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      # first phase should have two items
      first_phase = T.must(subscription_schedule.phases.first)

      phase_item_1 = T.must(first_phase.items.first)
      phase_item_1_discount = T.must(phase_item_1.discounts.first)

      phase_item_2 = T.must(first_phase.items[1])
      phase_item_2_discount = T.must(phase_item_2.discounts.first)

      # the coupon ids on the phase items should be equal
      assert_equal(phase_item_1_discount.coupon, phase_item_2_discount.coupon)
      stripe_coupon = Stripe::Coupon.retrieve(T.must(phase_item_1_discount).coupon, @user.stripe_credentials)
      assert_equal('$55 off coupon', stripe_coupon.name)
      assert_equal(55 * 100, stripe_coupon.amount_off)
      assert(stripe_coupon.valid)
    end

    it 'creates a new stripe coupon if the sf quote coupon is modified' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id_1, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "create_new_stripe_coupon", additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      })

      # add product to the sf quote
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      # create an sf coupon
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '55% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 55,
      })

      # create the quote line coupon association to map the coupon to the quote lines
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines.first.Id, sf_stripe_coupon_id: sf_percent_off_coupon_id)

      # create and translate the order
      sf_order_1 = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order_1.Id)
      sf_order_1.refresh

      # now update the coupon in salesforce and attempt to translate a new order using this coupon
      sf.update!(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), {
        SF_ID => sf_percent_off_coupon_id,
        prefixed_stripe_field('Name__c') => 'Special coupon',
      })

      # create another sf quote
      sf_quote_id_2 = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "create_new_stripe_coupon_2", additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })
      # add a product to the sf quote
      quote_with_product = add_product_to_cpq_quote(sf_quote_id_2, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      # create the quote line coupon association to map the coupon to the quote lines
      quote_lines = sf_get_related(sf_quote_id_2, CPQ_QUOTE_LINE)
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines.first.Id, sf_stripe_coupon_id: sf_percent_off_coupon_id)

      # create and translate the order
      sf_order_2 = create_order_from_cpq_quote(sf_quote_id_2)
      StripeForce::Translate.perform_inline(@user, sf_order_2.Id)
      sf_order_2.refresh

      # sanity check these are two different orders
      assert_not_equal(sf_order_1.Id, sf_order_2.Id)

      # fetch the stripe subscription schedules
      subscription_schedule_1 = Stripe::SubscriptionSchedule.retrieve(sf_order_1[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
      first_phase_1 = T.must(subscription_schedule_1.phases.first)
      phase_item_1 = T.must(first_phase_1.items.first)
      phase_item_1_discount = T.must(phase_item_1.discounts.first)

      subscription_schedule_2 = Stripe::SubscriptionSchedule.retrieve(sf_order_2[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
      first_phase_2 = T.must(subscription_schedule_2.phases.first)
      phase_item_2 = T.must(first_phase_2.items.first)
      phase_item_2_discount = T.must(phase_item_2.discounts.first)

      # the coupons on the phase items should not be equal
      assert_not_equal(phase_item_1_discount.coupon, phase_item_2_discount.coupon)
    end
  end

  it 'stripe invoice final due amount takes into account coupons' do
    # setup
    sf_account_id = create_salesforce_account
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

    # create a CPQ quote
    sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, contact_email: "stripe_invoice_final_due_amount", additional_quote_fields: {
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

    # create two coupons and attach both to the quote line
    sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => '25% off coupon',
      SalesforceStripeCouponFields::PERCENT_OFF => 25,
    })
    sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => '$30 off coupon',
      SalesforceStripeCouponFields::AMOUNT_OFF => 30,
    })

    # create the quote line coupon association object to map the coupons to the quote line
    create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
    create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

    # create and translate the SF order
    sf_order = create_order_from_cpq_quote(sf_quote_id)
    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    # fetch invoice and verify final amount due
    sf_account = sf_get(sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)
    assert_equal(1, invoice.data.count)

    # expected price should equal $120 with 25% off and $30 off coupons applied
    # note: order line coupons are not applied in any given order so could be either final amount
    assert_includes([(TEST_DEFAULT_PRICE * 0.75 - (30 * 100)).to_i, ((TEST_DEFAULT_PRICE - (30 * 100)) * 0.75).to_i], invoice.data.first.amount_due)
  end
end
