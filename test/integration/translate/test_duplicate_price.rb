# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::DuplicatePriceTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'does not duplicate a price if no duplicate ID exists' do
    # also implicitly tests that no Stripe API calls are made in this operation
    price_1 = stripe_create_id(:price)
    price_2 = stripe_create_id(:price)
    order_line = create_mock_salesforce_order_item

    phase_items = [
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: price_1,
          quantity: 1,
        },
      ),
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: price_2,
          quantity: 1,
        },
      ),
    ]

    new_phase = StripeForce::Translate::OrderHelpers.ensure_unique_phase_item_prices(
      @user,
      phase_items
    )

    new_phase.each_with_index do |phase, i|
      assert_equal(phase.stripe_params, T.must(phase_items[i]).stripe_params)
      # assert_empty(diff)
    end
  end

  it 'properly duplicates a recurring price' do
    _, sf_pricebook_entry_id = salesforce_recurring_product_with_price

    recurring_price = StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    assert_equal(Stripe::Price, recurring_price.class)

    order_line = create_mock_salesforce_order_item
    phase_items = [
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: recurring_price.id,
          quantity: 1,
        },
      ),
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 2,
        order_line_id: order_line.Id,
        stripe_params: {
          price: recurring_price.id,
          quantity: 2,
        },
      ),
    ]

    new_phase = StripeForce::Translate::OrderHelpers.ensure_unique_phase_item_prices(
      @user,
      phase_items
    )

    assert_equal(2, phase_items.map(&:stripe_params).map {|h| h[:price] }.uniq.count)
    new_item = T.must(phase_items.detect {|i| i.stripe_params[:price] != recurring_price.id })
    new_price = Stripe::Price.retrieve(new_item.stripe_params[:price], @user.stripe_credentials)

    # let's make this test more paranoid
    assert(recurring_price.active)
    assert_equal(recurring_price.billing_scheme, new_price.billing_scheme)
    assert_equal(recurring_price.currency, new_price.currency)
    assert_equal(recurring_price.custom_unit_amount, new_price.custom_unit_amount)
    assert_equal(recurring_price.lookup_key, new_price.lookup_key)
    assert_equal(recurring_price.metadata, new_price.metadata)
    assert_equal(recurring_price.nickname, new_price.nickname)
    assert_equal(recurring_price.product, new_price.product)
    assert_equal(recurring_price.recurring, new_price.recurring)
    assert_equal(recurring_price.tax_behavior, new_price.tax_behavior)
    assert_equal(recurring_price.tiers_mode, new_price.tiers_mode)
    assert_equal(recurring_price.transform_quantity, new_price.transform_quantity)
    assert_equal(recurring_price.type, new_price.type)
    assert_equal(recurring_price.unit_amount, new_price.unit_amount)
    assert_equal(recurring_price.unit_amount_decimal, new_price.unit_amount_decimal)
  end

  it 'duplicates a one-time price' do
    sf_product_id, sf_pricebook_entry_id = salesforce_standalone_product_with_price

    one_time_price = StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)
    order_line = create_mock_salesforce_order_item

    phase_items = [
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: one_time_price.id,
          quantity: 1,
        },
      ),
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: one_time_price.id,
          quantity: 1,
        },
      ),
    ]

    new_phase_items = StripeForce::Translate::OrderHelpers.ensure_unique_phase_item_prices(
      @user,
      phase_items
    )

    assert_equal(2, new_phase_items.map(&:stripe_params).map {|h| h[:price] }.uniq.count)
    new_item = T.must(new_phase_items.detect {|i| i.stripe_params[:price] != one_time_price.id })
    new_price = Stripe::Price.retrieve(new_item.stripe_params[:price], @user.stripe_credentials)

    assert_equal("one_time", new_price.type)
  end

  it 'duplicates a metered price' do
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_metered_produce_with_price

    metered_price = StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)
    order_line = create_mock_salesforce_order_item

    phase_items = [
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: metered_price.id,
          quantity: 1,
        },
      ),
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: metered_price.id,
          quantity: 1,
        },
      ),
    ]

    new_phase_items = StripeForce::Translate::OrderHelpers.ensure_unique_phase_item_prices(
      @user,
      phase_items
    )

    assert_equal(2, new_phase_items.map(&:stripe_params).map {|h| h[:price] }.uniq.count)
    new_item = T.must(new_phase_items.detect {|i| i.stripe_params[:price] != metered_price.id })
    new_price = Stripe::Price.retrieve(new_item.stripe_params[:price], @user.stripe_credentials)

    assert_equal("metered", new_price.recurring.usage_type)
    assert_equal(metered_price.recurring, new_price.recurring)
  end

  it 'duplicates a tiered price' do
    sf_product_id, sf_pricebook_entry_id = create_recurring_per_unit_tiered_price

    tiered_price = StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)
    tiered_price = Stripe::Price.retrieve({id: tiered_price.id, expand: %w{tiers}}, @user.stripe_credentials)
    order_line = create_mock_salesforce_order_item

    phase_items = [
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: tiered_price.id,
          quantity: 1,
        },
      ),
      StripeForce::Translate::ContractItemStructure.new(
        quantity: 1,
        order_line_id: order_line.Id,
        stripe_params: {
          price: tiered_price.id,
          quantity: 1,
        },
      ),
    ]

    new_phase_items = StripeForce::Translate::OrderHelpers.ensure_unique_phase_item_prices(
      @user,
      phase_items
    )

    assert_equal(2, new_phase_items.map(&:stripe_params).map {|h| h[:price] }.uniq.count)
    new_item = T.must(new_phase_items.detect {|i| i.stripe_params[:price] != tiered_price.id })
    new_price = Stripe::Price.retrieve({id: new_item.stripe_params[:price], expand: %w{tiers}}, @user.stripe_credentials)

    assert_equal("licensed", new_price.recurring.usage_type)
    assert_equal("tiered", new_price.billing_scheme)
    assert_equal(tiered_price.tiers, new_price.tiers)
  end
end
