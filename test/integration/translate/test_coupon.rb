# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::CouponTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    @user.enable_feature(FeatureFlags::COUPONS)
  end

  describe 'sfdx' do
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

    it 'coupons are copied over when an sf quote is ordered' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

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
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      sf_quote_line_coupon = get_salesforce_stripe_coupons_associated_to_quote_line(quote_line_id: quote_line_id)
      assert_equal(1, sf_quote_line_coupon.size)
      assert_equal('25% off coupon', sf_quote_line_coupon.first[prefixed_stripe_field('Name__c')])

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      sf_order_lines = sf.query("SELECT Id FROM OrderItem WHERE OrderId = '#{sf_order.Id}'")
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # verify coupon order association was created

      # check if there are any coupon associations to this order or order item
      order_associations = @user.sf_client.query("Select Id from #{prefixed_stripe_field(SF_STRIPE_COUPON_ORDER_ASSOCIATION)} where #{prefixed_stripe_field('Order__c')} = '#{sf_order.Id}'")
      order_item_associations = @user.sf_client.query("Select Id from #{prefixed_stripe_field(SF_STRIPE_COUPON_ORDER_ITEM_ASSOCIATION)} where #{prefixed_stripe_field('Order_Item__c')} = '#{sf_order_lines.first.Id}'")

      assert_equal(order_associations.size, order_item_associations.size)
      assert_equal(1, order_associations.size)
    end
  end

  describe 'order translation with coupons' do
    it 'translate sf order with multiple coupons on an order line' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

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

      # sanity check the stripe coupons have the right data
      stripe_amount_off_coupon = stripe_coupon_1.name == '$10 off coupon' ? stripe_coupon_1 : stripe_coupon_2
      assert_equal(10, stripe_amount_off_coupon.amount_off)
      assert_equal("usd", stripe_amount_off_coupon.currency)
      assert_equal("once", stripe_amount_off_coupon.duration)
      assert_equal(sf_amount_off_coupon_id, stripe_amount_off_coupon.metadata['salesforce_stripe_coupon_id'])

      stripe_percent_off_coupon = stripe_coupon_1.name == '25% off coupon' ? stripe_coupon_1 : stripe_coupon_2
      assert_equal(25, stripe_percent_off_coupon.percent_off)
      assert_equal("once", stripe_percent_off_coupon.duration)
      assert_equal(sf_percent_off_coupon_id, stripe_percent_off_coupon.metadata['salesforce_stripe_coupon_id'])

      # confirm the stripe coupon ids were written back to the original coupons in salesforce
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

      # create two coupons
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '25% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 25,
      })
      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '$50 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 50,
      })

      # create the quote line coupon association object to map the coupon to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
      # create the quote coupon association object to map the coupon to the quote
      create_salesforce_stripe_coupon_quote_association(sf_quote_id: sf_quote_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

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
      stripe_phase_item_coupon = Stripe::Coupon.retrieve(T.must(phase_item_discount.first).coupon, @user.stripe_credentials)

      # sanity check the stripe coupons have the right data
      assert_equal(50, stripe_phase_coupon.amount_off)
      assert_equal("usd", stripe_phase_coupon.currency)
      assert_equal("once", stripe_phase_coupon.duration)
      assert_equal(sf_amount_off_coupon_id, stripe_phase_coupon.metadata['salesforce_stripe_coupon_id'])

      assert_equal(25, stripe_phase_item_coupon.percent_off)
      assert_equal("once", stripe_phase_item_coupon.duration)
      assert_equal(sf_percent_off_coupon_id, stripe_phase_item_coupon.metadata['salesforce_stripe_coupon_id'])

      # fetch invoice and verify final amount due
      sf_account = sf_get(sf_account_id)
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)
      assert_equal(1, invoice.data.count)

      # expected price should equal $120 (original price) with 25%-off and $50-off coupons applied
      assert_includes([(TEST_DEFAULT_PRICE * 0.75 - 50).to_i, ((TEST_DEFAULT_PRICE - 50) * 0.75).to_i], invoice.data.first.amount_due)
    end
  end

  describe 'coupon reuse' do
    it 'uses the same stripe coupon if translated twice' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id_1, _sf_pricebook_id = salesforce_recurring_product_with_price
      sf_product_id_2, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
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
      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '55% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 55,
      })
      # create the quote line coupon association to map the coupon to the quote lines
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines.first.Id, sf_stripe_coupon_id: sf_percent_off_coupon_id)
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_lines[1].Id, sf_stripe_coupon_id: sf_percent_off_coupon_id)

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
    end

    it 'creates new stripe coupon if sf coupon is modified' do
      # setup
      sf_account_id = create_salesforce_account
      sf_product_id_1, _sf_pricebook_id = salesforce_recurring_product_with_price

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
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
      sf.update!(prefixed_stripe_field(SF_STRIPE_COUPON), {
        SalesforceStripeCouponFields::ID => sf_percent_off_coupon_id,
        SalesforceStripeCouponFields::NAME => 'Special coupon',
      }.transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field)))

      # create another sf quote
      sf_quote_id_2 = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
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

  it 'stripe invoice final due amount reflects coupons' do
    # setup
    sf_account_id = create_salesforce_account
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price

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
    # note: coupons are not applied in any given order so could be either final amount
    assert_includes([(TEST_DEFAULT_PRICE * 0.75 - 30).to_i, ((TEST_DEFAULT_PRICE - 30) * 0.75).to_i], invoice.data.first.amount_due)
  end
end
