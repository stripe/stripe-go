# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceReuse < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
    @user.enable_feature(StripeForce::Constants::FeatureFlags::UPDATE_PRODUCT_ON_SYNC, update: true)
  end

  describe 'pricebook' do
    it 'uses the same tiered pricebook price when there are no customizations on the order line level' do
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

      sf_order = create_subscription_order(sf_product_id: sf_product_id, contact_email: "same_tired_pricebook_no_customizations")

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

    it 'uses the same licensed pricebook price object when the price object is translated twice' do
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

    it 'uses the same pricebook price object if there are no price customizations on the order line' do
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
        contact_email: "uses_the_same_pricebook_price_object",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        }
      )

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

    it 'does not use the same pricebook price object if there customizations on the order line level' do
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
        contact_email: "not_use_same_pricebook_object",
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
      new_stripe_price = Stripe::Price.retrieve(T.cast(item.price, String), @user.stripe_credentials)
      assert_equal(150_00, new_stripe_price.unit_amount)

      order_lines = sf.query("SELECT Id FROM OrderItem WHERE OrderId = '#{sf_order.Id}'")
      assert_equal(1, order_lines.count)

      sf_order_item = sf.find(SF_ORDER_ITEM, order_lines.first.Id)
      assert_equal(new_stripe_price.id, sf_order_item[prefixed_stripe_field(GENERIC_STRIPE_ID)])

      # the price for the order item should be archived
      refute(new_stripe_price.active)
    end
  end

  describe 'order line' do
    # https://jira.corp.stripe.com/browse/PLATINT-1833
    it 'uses the same order line price when the order line is translated twice' do
      translator = make_translator(user: @user)

      price_in_cents = 120_00
      adjusted_price = price_in_cents + 20_00

      sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(
        price: price_in_cents,
        additional_product_fields: {
          # CPQ prevents users from editing the line price if this is not defined
          'SBQQ__PriceEditable__c' => true,
        }
      )

      sf_account_id = create_salesforce_account
      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "use_same_order_line_price",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        }
      )

      # set unit price to differ from the standard price so the pricebook entry is not used
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      quote_with_product["lineItems"].first["record"]["SBQQ__ListPrice__c"] = adjusted_price / 100.0
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)
      sf_lines = sf_get_related sf_order, SF_ORDER_ITEM

      assert_equal(1, sf_lines.count)
      sf_line = sf_lines[0]

      translator.cache_service.cache_for_object(sf_order)

      # although stripe price is only created once, we expect two upsert calls to the same sync record
      translator.expects(:create_user_success).twice

      price = translator.create_price_for_order_item(sf_line)
      price = T.must(price)

      # internal API does not have strong gaurentees, let's make sure something hasn't changed
      assert_equal(Stripe::Price, price.class)
      # refresh the line item to pull the latest Stripe ID field value, which should contain our new price
      sf_line.refresh

      # when we translate this again, we shouldn't create a new price object
      Stripe::Price.expects(:create).never
      price_2 = translator.create_price_for_order_item(sf_line)
      price_2 = T.must(price_2)

      # the price objects should be completely identical
      assert_equal(price.id, price_2.id)
      assert_equal(price, price_2)

      # ensure the pricebook entry price is not used
      assert_equal(adjusted_price.to_s, price.unit_amount_decimal, "the adjusted order line price is not being used")
    end
  end
end
