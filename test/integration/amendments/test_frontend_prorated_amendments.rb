# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::BackendProratedAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

    @user = make_user(save: true)
    @user.enable_feature(StripeForce::Constants::FeatureFlags::FRONTEND_PRORATIONS, update: true)
  end

  it 'handles an initial order with a frontend proration' do
    # Creates an initial order that is 9 months with a yearly billed item

    frontend_prorated_term = 9
    contract_term = frontend_prorated_term
    order_start_date = now_time
    order_end_date = order_start_date + contract_term.months

    annual_price = 1200_00

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: annual_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      contact_email: "frontend_prorated",
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

    # this should be a single phase
    assert_equal(1, subscription_schedule.phases.count)

    phase = T.must(subscription_schedule.phases.first)

    # Starts today, ends in 9 months
    assert_equal(order_start_date.to_i, phase.start_date)
    assert_equal(order_end_date.to_i, phase.end_date)

    assert_equal(1, phase.items.count)
    sub_item = T.must(phase.items.first)
    stripe_price = stripe_get(sub_item.price)

    assert_equal("recurring", stripe_price.type)

    proration_multiplier = frontend_prorated_term.to_f / 12
    assert_equal(annual_price * proration_multiplier, stripe_price.unit_amount.to_f)
  end

  it 'does not translate a prorated item for a user without the feature enabled' do
    # Creates an initial order that is 9 months with a yearly billed item
    @user.disable_feature(StripeForce::Constants::FeatureFlags::FRONTEND_PRORATIONS, update: true)

    frontend_prorated_term = 9
    contract_term = frontend_prorated_term
    order_start_date = now_time
    order_end_date = order_start_date + contract_term.months

    annual_price = 1200_00

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: annual_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_order = create_salesforce_order(
      contact_email: "no_translate_prorated_item",
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

    # this should be a single phase
    assert_equal(1, subscription_schedule.phases.count)

    phase = T.must(subscription_schedule.phases.first)

    # Starts today, ends in 9 months
    assert_equal(order_start_date.to_i, phase.start_date)
    assert_equal(order_end_date.to_i, phase.end_date)

    assert_equal(1, phase.items.count)
    sub_item = T.must(phase.items.first)
    stripe_price = stripe_get(sub_item.price)

    assert_equal("recurring", stripe_price.type)

    assert_equal(annual_price, stripe_price.unit_amount.to_f)
  end

  it 'handles an initial order with a frontend proration order that also contains a non-recurring item' do
    # Creates an initial order that is 9 months with a yearly billed item

    frontend_prorated_term = 9
    contract_term = frontend_prorated_term
    order_start_date = now_time
    order_end_date = order_start_date + contract_term.months

    annual_price = 1200_00
    one_off_price = 100_00

    subscription_product_id, subscription_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: annual_price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    one_off_product_id, one_off_pricebook_entry_id = salesforce_standalone_product_with_price(price: one_off_price)

    sf_order = create_salesforce_order(
      contact_email: "proration_with_one_off",
      sf_product_id: subscription_product_id,

    )

    sf_pricebook_id = default_pricebook_id
    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id,
                                       contact_email: "proration_one_off_quote",
                                       additional_quote_fields: {
                                         CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
                                         CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
                                       })

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: subscription_product_id)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: one_off_product_id)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    # this should be a single phase
    assert_equal(1, subscription_schedule.phases.count)

    phase = T.must(subscription_schedule.phases.first)

    # Starts today, ends in 9 months
    assert_equal(order_start_date.to_i, phase.start_date)
    assert_equal(order_end_date.to_i, phase.end_date)

    assert_equal(1, phase.items.count)
    sub_item = T.must(phase.items.first)
    sub_stripe_price = stripe_get(sub_item.price)

    assert_equal("recurring", sub_stripe_price.type)

    proration_multiplier = frontend_prorated_term.to_f / 12
    assert_equal(annual_price * proration_multiplier, sub_stripe_price.unit_amount.to_f)

    assert_equal(1, phase.add_invoice_items.length)

    one_off_item = T.must(phase.add_invoice_items.first)
    one_off_stripe_price = stripe_get(one_off_item.price)

    assert_equal(one_off_price, one_off_stripe_price.unit_amount)
  end

end
