# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::TieredPricingTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  def create_salesforce_consumption_schedule(additional_fields={})
    sf.create!(SF_CONSUMPTION_SCHEDULE, {
      'Name' => sf_randomized_name(SF_CONSUMPTION_SCHEDULE),
      'Type' => 'Range',
      'RatingMethod' => 'Tier',
      'SBQQ__Category__c' => 'Rates',
      'BillingTerm' => 12,
      'BillingTermUnit' => "Month",
      # cannot create the consumption schedule as activated
      # 'IsActive' => true,
    }.merge(additional_fields))
  end

  def create_salesforce_consumption_rate(consumption_schedule_id, additional_fields={})
    @consumption_schedule_processing_order ||= Hash.new(0)
    @consumption_schedule_processing_order[consumption_schedule_id] += 1

    sf.create!(SF_CONSUMPTION_RATE, {
      "ConsumptionScheduleId" => consumption_schedule_id,
      "ProcessingOrder" => @consumption_schedule_processing_order[consumption_schedule_id],
      "PricingMethod" => "PerUnit",
      "LowerBound" => 1,
      "UpperBound" => 2,
      "Price" => 10,
    }.merge(additional_fields))
  end

  def activate_and_link_consumption_schedule(consumption_schedule_id, product_id)
    sf.update!(SF_CONSUMPTION_SCHEDULE, {
      SF_ID => consumption_schedule_id,
      'IsActive' => true,
    })

    # joining record, must be an activated consumption schedule
    sf.create!('ProductConsumptionSchedule', {
      'ConsumptionScheduleId' => consumption_schedule_id,
      'ProductId' => product_id,
    })
  end

  it 'uses order line consumption schedule information when it differs' do
    # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
    product_id, pricebook_entry_id = salesforce_recurring_product_with_price

    consumption_schedule_id = create_salesforce_consumption_schedule

    # first tier is priced at per-seat pricing, the rest are not
    # this test also tested >2 tiered pricing
    consumption_rate_id_1 = create_salesforce_consumption_rate(consumption_schedule_id)
    consumption_rate_id_2 = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 2,
      "UpperBound" => 10,
      "Price" => 20,
    })
    consumption_rate_id_3 = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 10,
      "UpperBound" => nil,
      "Price" => 30,
    })

    activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

    sf_order = create_subscription_order(sf_product_id: product_id)

    # it seems as though SF lets you edit the CPQ consumption rates after the order is activated, so we can just grab and edit directly
    cpq_consumption_rates = sf.query("SELECT Id FROM #{CPQ_CONSUMPTION_RATE} WHERE SBQQ__OrderItemConsumptionSchedule__r.SBQQ__OrderItem__r.OrderId = '#{sf_order.Id}' ORDER BY SBQQ__LowerBound__c ASC").map {|o| sf_get(o.Id) }

    # lets take the second rate and update the price
    sf.update!(CPQ_CONSUMPTION_RATE, {
      SF_ID => cpq_consumption_rates[1].Id,
      "SBQQ__Price__c" => 50,
    })

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.must(subscription_schedule.phases.first&.items&.first&.price)
    price = Stripe::Price.retrieve({
      id: price_id,
      # the `tiers` field is not included in the API unless it is expanded
      expand: %w{tiers},
    }, @user.stripe_credentials)

    # TODO had to remove the consumption schedule metadata due to char size limits
    # assert_equal(consumption_schedule_id, price.metadata['salesforce_consumption_schedule_id'])
    # TODO maybe assert nil on pricebook link since it isn't used for consumption schedule calculation

    # standard pricing fields should be null
    assert_nil(price.unit_amount)
    assert_nil(price.unit_amount_decimal)

    assert_equal('tiered', price.billing_scheme)
    # key change in this test
    assert_equal('volume', price.tiers_mode)

    first_tier = price.tiers.detect {|t| t.up_to == 2 }
    second_tier = price.tiers.detect {|t| t.up_to == 10 }

    # although `inf` is sent to the API, a null value is returned for `up_to`? Consistency?!
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
    third_tier = price.tiers.detect {|t| t.up_to.nil? }

    assert_equal("1000", first_tier.unit_amount_decimal)
    assert_equal(2, first_tier.up_to)

    assert_equal("5000", second_tier.unit_amount_decimal)
    assert_equal(10, second_tier.up_to)

    assert_equal("3000", third_tier.unit_amount_decimal)
    assert_nil(third_tier.up_to)

    # checks that price was created from order line and not price
    refute_nil(price.metadata['salesforce_order_item_id'])
  end

  # TODO can we define a consumption schedule without an upper bound? Is the price level then used in that case?
  it 'supports recurring, flat fee, volume, tiered pricing' do
    # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
    product_id, pricebook_entry_id = salesforce_recurring_product_with_price

    consumption_schedule_id = create_salesforce_consumption_schedule

    # first tier is priced at per-seat pricing, the rest are not
    # this test also tested >2 tiered pricing
    consumption_rate_id_1 = create_salesforce_consumption_rate(consumption_schedule_id)
    consumption_rate_id_2 = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 2,
      "UpperBound" => 10,
      "Price" => 20,
      "PricingMethod" => "FlatFee",
    })
    consumption_rate_id_3 = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 10,
      "UpperBound" => nil,
      "Price" => 30,
      "PricingMethod" => "FlatFee",
    })

    activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

    sf_order = create_subscription_order(sf_product_id: product_id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.must(subscription_schedule.phases.first&.items&.first&.price)
    price = Stripe::Price.retrieve({
      id: price_id,
      # the `tiers` field is not included in the API unless it is expanded
      expand: %w{tiers},
    }, @user.stripe_credentials.merge)

    assert_equal(consumption_schedule_id, price.metadata['salesforce_consumption_schedule_id'])
    # TODO maybe assert nil on pricebook link since it isn't used for consumption schedule calculation

    # standard pricing fields should be null
    assert_nil(price.unit_amount)
    assert_nil(price.unit_amount_decimal)

    assert_equal('tiered', price.billing_scheme)
    # key change in this test
    assert_equal('volume', price.tiers_mode)

    first_tier = price.tiers.detect {|t| t.up_to == 2 }
    second_tier = price.tiers.detect {|t| t.up_to == 10 }

    # although `inf` is sent to the API, a null value is returned for `up_to`? Consistency?!
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
    third_tier = price.tiers.detect {|t| t.up_to.nil? }

    assert_nil(first_tier.flat_amount_decimal)
    assert_equal("1000", first_tier.unit_amount_decimal)
    assert_equal(2, first_tier.up_to)

    assert_nil(second_tier.unit_amount_decimal)
    assert_equal("2000", second_tier.flat_amount_decimal)
    assert_equal(10, second_tier.up_to)

    assert_nil(third_tier.unit_amount_decimal)
    assert_equal("3000", third_tier.flat_amount_decimal)
    assert_nil(third_tier.up_to)
  end

  it 'supports recurring, per-unit, graduated, tiered pricing' do
    # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
    product_id, pricebook_entry_id = salesforce_recurring_product_with_price

    consumption_schedule_id = create_salesforce_consumption_schedule({
      # this indicates graduated pricing
      "Type" => "Slab",
    })

    consumption_rate_id = create_salesforce_consumption_rate(consumption_schedule_id)
    consumption_rate_id = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 2,
      "UpperBound" => nil,
      "Price" => 20,
    })

    activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

    sf_order = create_subscription_order(sf_product_id: product_id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.must(subscription_schedule.phases.first&.items&.first&.price)
    price = Stripe::Price.retrieve({
      id: price_id,
      # the `tiers` field is not included in the API unless it is expanded
      expand: %w{tiers},
    }, @user.stripe_credentials.merge)

    assert_equal(consumption_schedule_id, price.metadata['salesforce_consumption_schedule_id'])
    # TODO maybe assert nil on pricebook link since it isn't used for consumption schedule calculation

    # standard pricing fields should be null
    assert_nil(price.unit_amount)
    assert_nil(price.unit_amount_decimal)

    assert_equal('tiered', price.billing_scheme)
    # key change in this test
    assert_equal('graduated', price.tiers_mode)

    first_tier = price.tiers.detect {|t| !t.up_to.nil? }

    # although `inf` is sent to the API, a null value is returned for `up_to`? Consistency?!
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
    second_tier = price.tiers.detect {|t| t.up_to.nil? }

    assert_equal("1000", first_tier.unit_amount_decimal)
    assert_equal(2, first_tier.up_to)

    assert_equal("2000", second_tier.unit_amount_decimal)
    assert_nil(second_tier.up_to)
  end

  it 'supports non-recurring tiered pricing' do

  end

  # TODO two unbounded rates?

  it 'supports recurring, per unit, volume, tiered pricing' do
    # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
    product_id, pricebook_entry_id = salesforce_recurring_product_with_price

    consumption_schedule_id = create_salesforce_consumption_schedule
    consumption_rate_id = create_salesforce_consumption_rate(consumption_schedule_id)
    consumption_rate_id = create_salesforce_consumption_rate(consumption_schedule_id, {
      "LowerBound" => 2,
      "UpperBound" => nil,
      "Price" => 20,
    })

    activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

    sf_order = create_subscription_order(sf_product_id: product_id)

    StripeForce::Translate.perform_inline(@user, sf_order.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    price_id = T.must(subscription_schedule.phases.first&.items&.first&.price)
    price = Stripe::Price.retrieve({
      id: price_id,
      # the `tiers` field is not included in the API unless it is expanded
      expand: %w{tiers},
    }, @user.stripe_credentials.merge)

    assert_equal(consumption_schedule_id, price.metadata['salesforce_consumption_schedule_id'])
    # TODO maybe assert nil on pricebook link since it isn't used for consumption schedule calculation

    # standard pricing fields should be null
    assert_nil(price.unit_amount)
    assert_nil(price.unit_amount_decimal)

    assert_equal('tiered', price.billing_scheme)
    assert_equal('volume', price.tiers_mode)

    first_tier = price.tiers.detect {|t| !t.up_to.nil? }

    # although `inf` is sent to the API, a null value is returned for `up_to`? Consistency?!
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
    second_tier = price.tiers.detect {|t| t.up_to.nil? }

    assert_equal("1000", first_tier.unit_amount_decimal)
    assert_equal(2, first_tier.up_to)

    assert_equal("2000", second_tier.unit_amount_decimal)
    assert_nil(second_tier.up_to)
  end

  it 'throws an error if an infinite tier is not specified' do
    # although the pricebook entry isn't used on our end when a consumption schedule is in place, it is still required by CPQ to go through the quoting process
    product_id, pricebook_entry_id = salesforce_recurring_product_with_price

    consumption_schedule_id = create_salesforce_consumption_schedule
    consumption_rate_id = create_salesforce_consumption_rate(consumption_schedule_id)

    activate_and_link_consumption_schedule(consumption_schedule_id, product_id)

    sf_order = create_subscription_order(sf_product_id: product_id)

    exception = assert_raises(Stripe::InvalidRequestError) do
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
    end

    assert_match("The tiers array must include a catch all tier", exception.message)

    # unrelated tests: ensure secondary record type is not the top-level order
    sync_record = get_sync_record_by_primary_id(sf_order.Id)
    assert_equal(SF_PRICEBOOK_ENTRY, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])
  end
end
