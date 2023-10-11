# frozen_string_literal: true
# typed: true

require_relative '../test_helper'
require_relative './amendments/_lib'

class Critic::EvergreenOrders < Critic::OrderAmendmentFunctionalTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
    @user.field_defaults = {
      "customer" => {
        "email" => create_random_email,
      },
      "subscription" => {
        "days_until_due" => 30,
        "collection_method" => "send_invoice",
      },
      "subscription_schedule" => {
        "default_settings.invoice_settings.days_until_due" => 30,
        "default_settings.collection_method" => "send_invoice",
      },
    }

    @user.field_mappings = {
      "subscription" => {
        "metadata.example_field" => "OrderNumber",
      },
    }
    @user.save
  end

  it 'creates an evergreen order object' do
    sf_order = create_evergreen_salesforce_order(contact_email: "create_evergreen")

    # get the order items to check the subscription data inside
    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

    order_items.each do |order_item|
      assert_equal(1, order_item[CPQ_QUOTE_SUBSCRIPTION_TERM])
      assert_equal(CPQProductSubscriptionTypeOptions::EVERGREEN.serialize, order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE])
      assert_equal('Fixed Price', order_item[CPQ_QUOTE_SUBSCRIPTION_PRICING])
    end
  end

  it 'translates evergreen salesforce order to a stripe subscription' do
    # create two products to add to the order
    sf_product_id_1, _ = salesforce_evergreen_product_with_price
    sf_product_id_2, _ = salesforce_evergreen_product_with_price

    sf_account_id = create_salesforce_account

    subscription_start_date = now_time + 5.day

    quote_id = create_salesforce_quote(
      sf_account_id: sf_account_id,
      contact_email: "translate_evergreen",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(subscription_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # add both products to the sf quote
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    # translate salesforce order to subscription and find
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    sf_order = sf.find(SF_ORDER, sf_order.Id)
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    stripe_customer = stripe_get(subscription_schedule.customer)
    refute_nil(stripe_customer.test_clock)

    # there is no subscription attached to the schedule yet
    assert_nil(subscription_schedule.subscription)

    # advance clock to after Stripe subscription is created and released by schedule
    test_clock = advance_test_clock(stripe_customer, (subscription_start_date + 1.days).to_i)

    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    subscription = Stripe::Subscription.retrieve(subscription_schedule.released_subscription, @user.stripe_credentials)

    # top level subscription fields
    assert_match(sf_order.Id, subscription.metadata['salesforce_order_link'])
    assert_equal(subscription.metadata['salesforce_order_id'], sf_order.Id)

    assert_equal(subscription_schedule.default_settings.collection_method, subscription.collection_method)
    assert_equal(subscription_schedule.default_settings.invoice_settings.days_until_due, subscription.days_until_due)

    # subscription products
    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

    assert_equal(sf_product_id_1, order_items.first['Product2Id'])
    assert_equal(sf_product_id_2, order_items.second['Product2Id'])

    # Check subscription and invoice properties
    invoices = Stripe::Invoice.list({subscription: subscription, limit: 100}, @user.stripe_credentials)
    stripe_invoice = Stripe::Invoice.retrieve(invoices.data.first['id'], @user.stripe_credentials)

    # Start date
    assert_equal(Time.at(subscription.current_period_start).utc, subscription_start_date)
    assert_equal('active', subscription.status)
  end

  it 'amendment to scheduled evergreen stripe subscription changes stripe id on sf order' do
    subscription_start_date = now_time + 5.days

    sf_order = create_evergreen_salesforce_order(
      # need to set these fields explicitly to use translate
      contact_email: "change_stripe_id_on_sf_order",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(subscription_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # translate salesforce order to subscription and find
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    sf_order = sf.find(SF_ORDER, sf_order.Id)
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    # there is no subscription attached to the schedule yet
    assert_nil(subscription_schedule.subscription)

    stripe_customer = stripe_get(subscription_schedule.customer)
    refute_nil(stripe_customer.test_clock)

    test_clock = advance_test_clock(stripe_customer, (subscription_start_date + 1.days).to_i)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

    assert_not_nil(subscription_schedule.released_subscription)

    subscription = Stripe::Subscription.retrieve(T.must(subscription_schedule.released_subscription), @user.stripe_credentials)

    # start date
    assert_equal(Time.at(subscription.current_period_start).utc, subscription_start_date)

    amendment_start_date = subscription_start_date + 3.days

    sf_contract = create_contract_from_order(sf_order)
    sf_order.refresh

    amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

    # wipe out the product
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

    sf_order_amendment = create_order_from_quote_data(amendment_quote)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    # translate salesforce amendment
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # stripe id on sf order should have been updated to the subscription's
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    assert_equal(stripe_id, subscription.id)

    test_clock = advance_test_clock(stripe_customer, (amendment_start_date + 1.days).to_i)
    subscription = Stripe::Subscription.retrieve(subscription.id, @user.stripe_credentials)
    assert_equal('canceled', subscription.status)
  end

  it 'ignores attempt to resync evergreen order amendment' do
    @user.disable_feature FeatureFlags::SF_CACHING, update: true
    sf_order = create_evergreen_salesforce_order(
      # need to set these fields explicitly to use translate
      contact_email: "ignore_resync_attempt_3",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # translate salesforce order to subscription and find
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    sf_contract = create_contract_from_order(sf_order)
    # api precondition: initial orders have a nil contract ID
    sf_order.refresh
    assert_nil(sf_order.ContractId)

    # the contract should reference the initial order that was created
    assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

    amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

    # wipe out the product
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = now_time_formatted_for_salesforce
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

    sf_order_amendment = create_order_from_quote_data(amendment_quote)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    assert_nil(sf_order_amendment.Stripe_ID__c)
    # translate order amendment
    SalesforceTranslateRecordJob.translate(@user, sf_order_amendment)
    sf_order_amendment.refresh

    sf_order.refresh
    sf_order = sf.find(SF_ORDER, sf_order.Id)
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    assert_not_nil(sf_order_amendment.Stripe_ID__c)
    assert_equal('canceled', subscription.status)

    # sync it again
    SalesforceTranslateRecordJob.translate(@user, sf_order_amendment)
  end

  describe 'failure cases' do
    it 'raises user error when there are evergreen and renewable items on same order' do
      sf_account_id = create_salesforce_account

      sf_product_id_1, _ = salesforce_evergreen_product_with_price
      sf_product_id_2, _ = salesforce_recurring_product_with_price

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "raise_error_when_evergreen_and_renewable",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }
      )

      # add both products to the sf quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Salesforce orders with both Evergreen and Renewable items are not yet supported.", exception.message)
    end

    it 'raises user error when evergreen order has default subscription term not 1' do
      sf_product_id = create_salesforce_product(additional_fields: {
        # anything non-nil indicates subscription/recurring pricing
        CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',
        CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::EVERGREEN,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 2,
      })
      _ = create_salesforce_price(sf_product_id: sf_product_id)

      sf_account_id = create_salesforce_account

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,
        contact_email: "raise_error_when_evergreen_has_default_term",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
          CPQ_QUOTE_SUBSCRIPTION_TERM => 12,
        }
      )

      # add both products to the sf quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
      calculate_and_save_cpq_quote(quote_with_product)
      sf_order = create_order_from_cpq_quote(quote_id)

      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Evergreen Salesforce orders should have default subscription term equal to 1.", exception.message)
    end

    it 'raises error when attempt to amend subscription that has not started' do
      subscription_start_date = now_time + 5.day

      sf_order = create_evergreen_salesforce_order(
        # need to set these fields explicitly to use translate
        contact_email: "raise_error_when_amending_before_start_date",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(subscription_start_date),
          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }
      )

      amendment_start_date = subscription_start_date + 1.days

      # create amendment with one product to add
      sf_contract = create_contract_from_order(sf_order)
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

      sf_product_id, _ = salesforce_evergreen_product_with_price
      quote_with_product = add_product_to_cpq_quote(amendment_quote['record']['Id'], sf_product_id: sf_product_id)
      calculate_and_save_cpq_quote(quote_with_product)

      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_start_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      # translate salesforce amendment
      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Amending a Salesforce evergreen order that has not started is not yet supported.", exception.message)
    end
  end
end
