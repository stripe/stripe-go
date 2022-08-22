# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::SubscriptionTermTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  # TODO but what exactly is this testing? Document this more clearly.
  it 'integrates an order with annual billing with a valid subscription term' do
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      }
    )

    # api preconditions:
    # * proration multiplier should be 1 since the order & product term match
    # * the line item price should be exactly what is specified on the product
    sf_order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, sf_order_items.count)

    sf_order_item = sf_order_items.first
    assert_equal(1, sf_order_item['SBQQ__ProrateMultiplier__c'])
    assert_equal(120, sf_order_item['UnitPrice'].to_i)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    start_date = now_time
    end_date = start_date + 1.year

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    # end date should only be one year although all prices are billing yearly
    assert_equal(0, phase.start_date - start_date.to_i)
    assert_equal(0, phase.end_date - end_date.to_i)

    assert_equal(1, phase.items.count)

    item = T.must(phase.items.first)
    price = Stripe::Price.retrieve(T.cast(item.price, String), @user.stripe_credentials)
    assert_equal(12, price.recurring.interval_count)
    assert_equal('month', price.recurring.interval)
  end

  # in this scenario a new price will be created for each line item, the pricebook entry price can never be used
  it 'translates an order without a subscription term or frequency on the product but contains a custom price and billing frequency' do
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => nil,
        CPQ_QUOTE_SUBSCRIPTION_TERM => nil,
      }
    )

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(
      sf_account_id: sf_account_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        # the term of the quote does not automatically propogate down to line items
        # if a term is not specified on the product, then none is defined on the order item
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12,
      }
    )

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
    assert_nil(quote_with_product["lineItems"].first["record"]["SBQQ__DefaultSubscriptionTerm__c"])
    quote_with_product["lineItems"].first["record"]["SBQQ__BillingFrequency__c"] = 'Monthly'
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    # api preconditions: test our assumptions about the proration multipler when nothing is specified on the line level
    sf_order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, sf_order_items.count)

    # TODO document this in our technical document
    # if no subscription term is specified on the product, it looks like it is assumed to match the subscription term of the order
    sf_order_item = sf_order_items.first
    assert_equal(1, sf_order_item['SBQQ__ProrateMultiplier__c'])
    assert_nil(sf_order_item['SBQQ__UnproratedNetPrice__c'])

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    start_date = now_time
    end_date = start_date + 1.year

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    # end date should only be one year although all prices are billing yearly
    assert_equal(0, phase.start_date - start_date.to_i)
    assert_equal(0, phase.end_date - end_date.to_i)

    assert_equal(1, phase.items.count)

    item = T.must(phase.items.first)
    price = Stripe::Price.retrieve(T.cast(item.price, String), @user.stripe_credentials)
    assert_equal(1, price.recurring.interval_count)
    assert_equal('month', price.recurring.interval)
    assert_equal("1000", price.unit_amount_decimal)
  end
end
