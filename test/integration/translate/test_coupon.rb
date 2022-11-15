# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::CouponTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'translate an SF coupon with percent off set' do
    # setup
    COUPON_NAME = '100% off coupon'
    COUPON_PERCENT_OFF = 100
    COUPON_MAX_REDEMPTIONS = 5

    # create the SF Stripe coupon
    sf_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => COUPON_NAME,
      SalesforceStripeCouponFields::PERCENT_OFF => COUPON_PERCENT_OFF,
    })

    # translate the SF coupon
    StripeForce::Translate.perform_inline(@user, sf_coupon_id)

    # confirm the Stripe ID was written back into the SF coupon obj
    sf_coupon = sf.find(SF_STRIPE_COUPON, sf_coupon_id)
    stripe_coupon_id = sf_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    assert(stripe_coupon_id)

    # retrieve the created Stripe coupon
    stripe_coupon = Stripe::Coupon.retrieve(stripe_coupon_id, @user.stripe_credentials)

    # compare the created Stripe coupon with the SF coupon
    assert_equal(COUPON_NAME, stripe_coupon.name)
    assert_equal(COUPON_PERCENT_OFF, stripe_coupon.percent_off)
    assert_equal('once', stripe_coupon.duration)
    assert_nil(stripe_coupon.currency)
    assert_equal(sf_coupon.Id, stripe_coupon.metadata['salesforce_stripe_coupon_beta_id'])
    assert_match(sf_coupon.Id, stripe_coupon.metadata['salesforce_stripe_coupon_beta_link'])
  end

  it 'verify multiple SF coupons are attached to a quote line' do
    # setup
    PRODUCT_PRICE = 100
    sf_account_id = create_salesforce_account
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(price: PRODUCT_PRICE)

    # create a SF CPQ quote
    sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    # create a sf quote with a product
    quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
    calculate_and_save_cpq_quote(quote_with_product)

    # retrieve the quote line
    quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
    assert_equal(1, quote_lines.size)
    quote_line_id = quote_lines.first.Id

    # create a coupon and attach to the quote line
    sf_coupon_id_1 = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => '25% off coupon',
      SalesforceStripeCouponFields::PERCENT_OFF => 25,
    })

    sf_coupon_id_2 = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => '$10 off coupon',
      SalesforceStripeCouponFields::AMOUNT_OFF => 10,
    })

    # create the association object to map the coupon to the quote line
    create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_coupon_id_1)
    create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_coupon_id_2)

    # note: the quote line does not have a reference to the stripe coupon quote line association object
    # so we query for the association objects that have a reference to this quote line
    associated_coupons = get_salesforce_stripe_coupons_associated_to_quote_line(quote_line_id: quote_line_id)
    assert(2, associated_coupons.size)

    coupon_1 = associated_coupons.first.Name_c == '25% off coupon' ? associated_coupons[0] : associated_coupons[1]
    assert(25, coupon_1.Percent_Off__c)
  end

  it 'coupons are copied to order lines when a quote is ordered' do
    # setup
    PRODUCT_PRICE = 100
    sf_account_id = create_salesforce_account
    sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(price: PRODUCT_PRICE)

    # create a SF CPQ quote
    sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    # create a quote with a product
    quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
    sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

    # retrieve the quote line
    quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
    assert_equal(1, quote_lines.size)
    quote_line_id = quote_lines.first.Id

    # create a coupon and attach to the quote line
    sf_stripe_coupon = create_salesforce_stripe_coupon(additional_fields: {
      SalesforceStripeCouponFields::NAME => 'Special 50% off coupon',
      SalesforceStripeCouponFields::PERCENT_OFF => 50,
    })

    # create the association object to map the coupon to the quote line
    create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_stripe_coupon)

    sf_order = create_order_from_cpq_quote(sf_quote_id)

    # query for the association objects that have a reference to this order line
    sf_order_line_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, sf_order_line_items.count)

    sf_order_item_id = sf_order_line_items.first.Id
    associated_coupons = StripeForce::Translate.get_salesforce_stripe_coupons_associated_to_order_line(sf_client: @user.sf_client, sf_order_line_id: sf_order_item_id)
    assert_equal(1, associated_coupons.size)

    order_line_coupon = associated_coupons.first
    assert_equal('Special 50% off coupon', order_line_coupon.Name__c)
    assert_equal(50, order_line_coupon.Percent_Off__c)
  end

  it 'translate order with coupon' do
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1952
  end
end
