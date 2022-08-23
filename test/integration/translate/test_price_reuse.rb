# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceReuse < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'uses the same tiered price when there are no customizations on the order line level' do
    sf_product_id, sf_pricebook_entry_id = create_recurring_per_unit_tiered_price

    # first, translate the price twice and ensure the same ID is used, then we'll test the order line
    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_price_id)

    stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

    Stripe::Price.expects(:create).never

    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry.refresh
    assert_equal(stripe_price.id, sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)])

    sf_order = create_subscription_order(sf_product_id: sf_product_id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_id)

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal(1, subscription_schedule.phases.count)
    assert_equal(1, subscription_schedule.phases.first&.items&.count)

    item = subscription_schedule.phases.first&.items&.first
    assert_equal(stripe_price_id, item&.price)
  end
end
