# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

# there are some metadata special cases that require some specific tests
class Critic::MapperIntegrationTests < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'allows the quantity field to be customized' do
    @user.field_mappings['subscription_item'] = {
      'quantity' => 'Description',
    }

    sf_order = create_subscription_order

    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)

    order_item = order_items.first
    sf.update!(SF_ORDER_ITEM, {
      SF_ID => order_item.Id,
      "Description" => "10",
    })

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    item = T.must(subscription_schedule.phases.first&.items&.first)
    assert_equal(10, item.quantity)
  end

  it 'allows the price field to be customized' do
    @user.field_mappings['price_order_item'] = {
      'unit_amount_decimal' => 'Description',
    }

    sf_order = create_subscription_order

    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)
    order_item = order_items.first

    sf.update!(SF_ORDER_ITEM, {
      SF_ID => order_item.Id,
      "Description" => "200.121212",
    })

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.cast(subscription_schedule.phases.first&.items&.first&.price, String)
    price = Stripe::Price.retrieve(price_id, @user.stripe_credentials)
    assert_equal("20012.1212", price.unit_amount_decimal)

    order_item.refresh
    assert_equal(price.id, order_item[prefixed_stripe_field(GENERIC_STRIPE_ID)])

    pricebook_entry = sf_get(order_item.PricebookEntryId)
    assert_nil(pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)])
  end

  it 'allows the price field to be customized when mapped from an order line' do
    # we need to map both since the order line will be used if it is different than the pricebook mapping
    @user.field_mappings['price'] = {
      'unit_amount_decimal' => 'Product2.Description',
    }

    @user.field_mappings['price_order_item'] = {
      'unit_amount_decimal' => 'Description',
    }

    sf_order = create_subscription_order

    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)
    order_item = order_items.first

    sf.update!(SF_ORDER_ITEM, {
      SF_ID => order_item.Id,
      "Description" => "200.121212",
    })

    product = sf_get(order_item.Product2Id)

    sf.update!(SF_PRODUCT, {
      SF_ID => product.Id,
      "Description" => "200.121212",
    })

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.cast(subscription_schedule.phases.first&.items&.first&.price, String)
    price = Stripe::Price.retrieve(price_id, @user.stripe_credentials)
    assert_equal("20012.1212", price.unit_amount_decimal)

    pricebook_entry = sf_get(order_item.PricebookEntryId)
    assert_equal(price.id, pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)])
  end

  # https://jira.corp.stripe.com/browse/PLATINT-1421
  it 'supports metadata on the subscription phase item level' do
    @user.field_mappings['subscription_item'] = {
      'metadata.special_description' => 'Description',
    }

    sf_order = create_subscription_order

    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)

    order_item = order_items.first
    sf.update!(SF_ORDER_ITEM, {
      SF_ID => order_item.Id,
      "Description" => "special field",
    })

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    item = T.must(subscription_schedule.phases.first&.items&.first)
    assert_equal('special field', item.metadata['special_description'])
  end
end
