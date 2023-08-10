# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
    @user.enable_feature(StripeForce::Constants::FeatureFlags::UPDATE_PRODUCT_ON_SYNC, update: true)
  end

  it 'translates a non-recurring pricebook entry to a stripe price' do
    price_in_cents = 100_00

    sf_product_id = create_salesforce_product(static_id: true)
    sf_pricebook_entry_id = create_salesforce_price(sf_product_id: sf_product_id, price: price_in_cents)

    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_price_id)

    stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

    # TODO this will break when we create dynamic
    assert_equal(sf_pricebook_entry.Id, stripe_price.metadata['salesforce_pricebook_entry_id'])
    assert_match(sf_pricebook_entry.Id, stripe_price.metadata['salesforce_pricebook_entry_link'])

    assert_equal(price_in_cents, stripe_price.unit_amount)
    assert_equal('per_unit', stripe_price.billing_scheme)
    assert_equal('one_time', stripe_price.type)
    assert_equal('usd', stripe_price.currency)

    stripe_product = Stripe::Product.retrieve(T.cast(stripe_price.product, String), @user.stripe_credentials)
    assert_equal(sf_product_id, stripe_product.metadata['salesforce_product2_id'])
    assert_match(sf_product_id, stripe_product.metadata['salesforce_product2_link'])
  end

  it 'translates a recurring pricebook entry to a stripe price' do
    price_in_cents = 120_00

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price_in_cents, static_id: true)

    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_price_id)

    stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

    # TODO this will break when we create dynamic
    assert_equal(sf_pricebook_entry.Id, stripe_price.metadata['salesforce_pricebook_entry_id'])
    assert_match(sf_pricebook_entry.Id, stripe_price.metadata['salesforce_pricebook_entry_link'])

    assert_equal(price_in_cents, stripe_price.unit_amount)
    assert_equal('per_unit', stripe_price.billing_scheme)
    assert_equal('recurring', stripe_price.type)
    assert_equal('usd', stripe_price.currency)

    assert_equal('month', stripe_price.recurring.interval)
    assert_equal(1, stripe_price.recurring.interval_count)
    assert_equal('licensed', stripe_price.recurring.usage_type)

    stripe_product = Stripe::Product.retrieve(T.cast(stripe_price.product, String), @user.stripe_credentials)
    assert_equal(sf_product_id, stripe_product.metadata['salesforce_product2_id'])
    assert_match(sf_product_id, stripe_product.metadata['salesforce_product2_link'])
  end

  it 'translates a recurring metered pricebook entry to a stripe price' do
    @user.field_defaults = {
      "price" => {
        "recurring.aggregate_usage" => "last_during_period",
      },
    }
    @user.save

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(static_id: true)

    sf.update!(SF_PRODUCT, {
      SF_ID => sf_product_id,

      # this is transformed into a metered billing type
      CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS,
    })

    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_price_id)

    stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

    assert_equal('recurring', stripe_price.type)
    assert_equal('month', stripe_price.recurring.interval)
    assert_equal(1, stripe_price.recurring.interval_count)
    assert_equal('metered', stripe_price.recurring.usage_type)
    assert_equal('last_during_period', stripe_price.recurring.aggregate_usage)
  end

  it 'handles decimal prices' do
    # subcent decimal prices
    price_in_cents = 100_25.55

    sf_product_id = create_salesforce_product(static_id: true)
    sf_pricebook_entry_id = create_salesforce_price(sf_product_id: sf_product_id, price: price_in_cents)

    StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_price_id)

    stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)
    assert_nil(stripe_price.unit_amount)
    assert_equal('10025.55', stripe_price.unit_amount_decimal)
  end

  describe 'cpq details' do
    it 'sets the interval_count based on the billing frequency selected' do
      sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(static_id: true)

      sf.update!(SF_PRODUCT, {
        SF_ID => sf_product_id,
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::QUARTERLY.serialize,
      })

      StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

      sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
      stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_price_id)

      stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

      # quarterly frequency is 3 `interval_count` on a `month`
      assert_equal('recurring', stripe_price.type)
      assert_equal('month', stripe_price.recurring.interval)
      assert_equal(3, stripe_price.recurring.interval_count)
      assert_equal('licensed', stripe_price.recurring.usage_type)
    end
  end
end
