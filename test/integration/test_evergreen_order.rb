# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::EvergreenOrderTest < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
  end

  it 'creates an evergreen order object' do
    @user.field_defaults['customer'] = {
      'email' => create_random_email,
    }

    sf_order = create_evergreen_salesforce_order

    # get the order items to check the subscription data inside
    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

    order_items.each do |order_item|
      assert_equal(1, order_item[CPQ_QUOTE_SUBSCRIPTION_TERM])
      assert_equal(CPQProductSubscriptionTypeOptions::EVERGREEN.serialize, order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE])
      assert_equal('Fixed Price', order_item[CPQ_QUOTE_SUBSCRIPTION_PRICING])
    end
  end

  describe 'failure cases' do
    it 'raises user error when there are evergreen and renewable items on same order' do
      @user.field_defaults = {
        "customer" => {
          "email" => create_random_email,
        },
        "subscription" => {
          "days_until_due" => 30,
          "collection_method" => "send_invoice",
        },
      }
      @user.save

      sf_account_id = create_salesforce_account

      sf_product_id_1, _ = salesforce_evergreen_product_with_price
      sf_product_id_2, _ = salesforce_recurring_product_with_price

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,

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

      order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Salesforce orders with both Evergreen and Renewable items are not yet supported.", exception.message)
    end
  end

  it 'translates evergreen salesforce order to a stripe subscription' do
    @user.field_defaults = {
      "customer" => {
        "email" => create_random_email,
      },
      "subscription" => {
        "days_until_due" => 30,
        "collection_method" => "send_invoice",
      },
    }

    @user.field_mappings = {
      "subscription" => {
        "metadata.example_field" => "OrderNumber",
      },
    }
    @user.save

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_evergreen_product_with_price

    sf_account_id = create_salesforce_account

    subscription_start_date = format_date_for_salesforce(now_time)

    sf_order = create_evergreen_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => subscription_start_date,
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # translate salesforce order to subscription and find
    StripeForce::Translate.perform_inline(@user, sf_order.Id)
    sf_order.refresh

    sf_order = sf.find(SF_ORDER, sf_order.Id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(T.cast(subscription['customer'], String), @user.stripe_credentials)

    assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
    assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)

    # top level subscription fields
    assert_match(sf_order.Id, subscription.metadata['salesforce_order_link'])
    assert_equal(subscription.metadata['salesforce_order_id'], sf_order.Id)
    assert_equal(subscription.metadata['example_field'], sf_order.OrderNumber)

    # Check subscription and invoice properties
    stripe_invoice = Stripe::Invoice.retrieve(subscription['latest_invoice'], @user.stripe_credentials)

    # Start date
    assert_match(Time.at(subscription['current_period_start']).utc.strftime('%Y-%m-%d'), subscription_start_date)

    # Amount due per invoice
    assert_equal(stripe_invoice['amount_due'] / 100.0, sf_order['TotalAmount'])
  end

  it 'creates an invoice every pay period on the stripe subscription' do
    @user.field_defaults = {
      "customer" => {
        "email" => create_random_email,
      },
      "subscription" => {
        "days_until_due" => 30,
        "collection_method" => "send_invoice",
      },
    }
    @user.save

    sf_order = create_evergreen_salesforce_order(
      # need to set these fields explicitly to use translate
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # translate salesforce order to subscription and find
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # get Stripe customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)

    # get stripe subscription
    sf_order = sf.find(SF_ORDER, sf_order.Id)
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    # one invoice should be created for first pay period
    invoices = Stripe::Invoice.list({subscription: subscription, limit: 100}, @user.stripe_credentials)
    invoice_1 = Stripe::Invoice.retrieve(invoices.data.first['id'], @user.stripe_credentials)
    assert_equal(1, invoices.data.length)

    # Fast forward to start of next pay period to see another invoice get created
    test_clock = advance_test_clock(stripe_customer, (Time.at(subscription['current_period_end'])).to_i)

    # another invoice should have been created for second pay period
    invoices = Stripe::Invoice.list({subscription: subscription, limit: 100}, @user.stripe_credentials)
    assert_equal(2, invoices.data.length)

    # Ensure new invoices still have correct properties
    invoice_2 = Stripe::Invoice.retrieve(invoices.data.first['id'], @user.stripe_credentials)
    assert_equal(invoice_2['amount_due'] / 100.0, sf_order['TotalAmount'])
    assert_equal(invoice_1['period_end'], invoice_2['period_start'])

    # Fast forward more to ensure more invoices are created
    test_clock = advance_test_clock(stripe_customer, (Time.at(subscription['current_period_end']) + 2.month).to_i)
    invoices = Stripe::Invoice.list({subscription: subscription, limit: 100}, @user.stripe_credentials)
    invoice_3 = Stripe::Invoice.retrieve(invoices.data.second['id'], @user.stripe_credentials)
    invoice_4 = Stripe::Invoice.retrieve(invoices.data.first['id'], @user.stripe_credentials)

    assert_equal(4, invoices.data.length)
    assert_equal(invoice_4['amount_due'] / 100.0, sf_order['TotalAmount'])
    assert_equal(invoice_2['period_end'], invoice_3['period_start'])
    assert_equal(invoice_3['period_end'], invoice_4['period_start'])
  end
end
