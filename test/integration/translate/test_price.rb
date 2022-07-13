# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  # TODO complex price map testing from cloudflare

  describe 'price reuse' do
    it 'uses the same price object when the price object is translated twice' do
      _, sf_pricebook_entry_id = salesforce_recurring_product_with_price

      StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

      sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)
      stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_price_id)

      stripe_price = Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)

      Stripe::Price.expects(:create).never

      StripeForce::Translate.perform_inline(@user, sf_pricebook_entry_id)

      sf_pricebook_entry.refresh
      assert_equal(stripe_price.id, sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end

    it 'uses the same price object if there are no price customizations on the line level' do
      sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price

      # translate pricebook so price is created so we can assert that no new prices are created
      StripeForce::Translate.perform_inline(@user, sf_pricebook_id)
      sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id)
      stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_price_id)

      Stripe::Price.expects(:create).never

      sf_account_id = create_salesforce_account

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        }
      )

      # set unit price to differ from
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_id)

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(1, subscription_schedule.phases.count)
      assert_equal(1, subscription_schedule.phases.first&.items&.count)

      item = subscription_schedule.phases.first&.items&.first
      assert_equal(stripe_price_id, item&.price)
    end

    it 'does not use the same price object if there customizations on the line level' do
      price_in_cents = 120_00
      sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
        price: price_in_cents,
        additional_product_fields: {
          # CPQ prevents users from editing the line price if this is not defined
          'SBQQ__PriceEditable__c' => true,
        }
      )

      StripeForce::Translate.perform_inline(@user, sf_pricebook_id)

      sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id)
      stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_price_id)

      sf_account_id = create_salesforce_account

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        }
      )

      # set unit price to differ from the standard price
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      quote_with_product["lineItems"].first["record"]["SBQQ__ListPrice__c"] = 150.0
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      # cpq preconditions: total amount is the net amount of the contract over the life of the subscription
      assert_equal(150 * 12, sf_order.TotalAmount.to_i)

      # TODO test the order line price preconditions
      sf_order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
      assert_equal(1, sf_order_items.size)
      sf_order_item = sf_order_items.first

      # proration multiplier is order term / product term
      assert_equal(12, sf_order_item['SBQQ__ProrateMultiplier__c'])
      # list price is the original price before modification
      assert_equal(120, sf_order_item['ListPrice'].to_i)
      assert_equal(150, sf_order_item['SBQQ__QuotedListPrice__c'].to_i)
      assert_equal(150 * 12, sf_order_item['UnitPrice'].to_i)
      assert_equal(sf_order_item['UnitPrice'], sf_order_item['TotalPrice'])

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_id)

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(1, subscription_schedule.phases.count)
      assert_equal(1, subscription_schedule.phases.first&.items&.count)

      item = T.must(subscription_schedule.phases.first&.items&.first)
      refute_equal(stripe_price_id, item.price)

      # is the customized price being used?
      new_stripe_price = Stripe::Price.retrieve(item.price, @user.stripe_credentials)
      assert_equal(150_00, new_stripe_price.unit_amount)

      order_lines = sf.query("SELECT Id FROM OrderItem WHERE OrderId = '#{sf_order.Id}'")
      assert_equal(1, order_lines.count)

      sf_order_item = sf.find(SF_ORDER_ITEM, order_lines.first.Id)
      assert_equal(new_stripe_price.id, sf_order_item[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end
  end

  it 'translates a non-recurring pricebook entry to a stripe price' do
    price_in_cents = 100_00

    sf_product_id = create_salesforce_product
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

  it 'handles a recurring pricebook entry to a stripe price' do
    price_in_cents = 120_00

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price_in_cents)

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

  it 'handles a recurring metered pricebook entry to a stripe price' do
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price

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
  end

  it 'handles decimal prices' do
    # subcent decimal prices
    price_in_cents = 100_25.55

    sf_product_id = create_salesforce_product
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
      sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price

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

    it 'handles daily CPQ terms'
  end
end
