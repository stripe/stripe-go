# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    # Only run on Multi-Currency Enabled CI Accounts
    unless @user.is_multicurrency_org?
      skip("Skipping multicurrency test on non-multicurrency org")
    end
  end

  describe 'success cases' do
    it 'translates a standard order in GBP' do
      currency_iso_code = 'GBP'

      sf_order = create_subscription_order(currency_iso_code: 'GBP')

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    end

    it 'translates a standard order in GBP, and then one in USD, for different customers but the same merchant' do
      currency_iso_code = 'GBP'

      sf_order = create_subscription_order(currency_iso_code: currency_iso_code)

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])

      currency_iso_code = 'USD'

      sf_order = create_subscription_order(currency_iso_code: currency_iso_code)

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    end

    it 'successfully creates sequential subscriptions in different currencies for a single customer' do
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      # Create Customer
      sf_account = create_salesforce_account

      order_one_start_date = DateTime.now
      contract_term = 12

      order_two_start_date = order_one_start_date + contract_term.months

      # Make a subscription in GBP
      currency_iso_code = 'GBP'

      sf_gbp_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        currency_iso_code: currency_iso_code,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order_one = create_subscription_order(sf_account_id: sf_account, sf_product_id: sf_gbp_product_id, currency_iso_code: currency_iso_code, additional_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => order_one_start_date,
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      })

      SalesforceTranslateRecordJob.translate(@user, sf_order_one)

      sf_order_one.refresh
      stripe_id = sf_order_one[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order_one.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])

      # Move the clock forward, as you cannot schedule a sub in a different currency until the old sub ends
      # https://jira.corp.stripe.com/browse/RUN_RB-5865
      stripe_customer = stripe_get(sf_get(sf_account)[GENERIC_STRIPE_ID])
      advance_test_clock(stripe_customer, (order_two_start_date).to_i)

      # Now for second Subscription, begining the same day that the other Subscription ends
      # Make a subscription in USD
      currency_iso_code = 'USD'

      sf_usd_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(
        currency_iso_code: currency_iso_code,
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_order_two = create_subscription_order(sf_account_id: sf_account, sf_product_id: sf_usd_product_id, currency_iso_code: currency_iso_code, additional_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => order_two_start_date,
        CPQ_QUOTE_SUBSCRIPTION_TERM => contract_term,
      })

      SalesforceTranslateRecordJob.translate(@user, sf_order_two)

      sf_order_two.refresh
      stripe_id = sf_order_two[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order_two.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    end
  end

  describe 'failure cases' do
    it 'fails to create concurrent subscriptions in multiple currencies for a single customer' do
      # Create Customer
      sf_account = create_salesforce_account

      # Make a subscription in GBP
      currency_iso_code = 'GBP'

      sf_order = create_subscription_order(sf_account_id: sf_account, currency_iso_code: currency_iso_code)

      SalesforceTranslateRecordJob.translate(@user, sf_order)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

      invoices = Stripe::Invoice.list({subscription: subscription_schedule.subscription}, @user.stripe_credentials)

      assert_equal(1, invoices.count)
      invoice = invoices.first

      # Make sure our currencies match
      assert_equal(currency_iso_code.downcase, invoice.currency.downcase)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.currency)
      assert_equal(currency_iso_code.downcase, invoice.lines.first.price.currency)

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])


      # Now for the failure, attempt to create another concurrent subscription but in USD
      currency_iso_code = 'USD'

      sf_order = create_subscription_order(sf_account_id: sf_account, currency_iso_code: currency_iso_code)

      begin
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      rescue Stripe::InvalidRequestError => e
        assert_equal("The currency of the selected prices (usd) does not match the customer currency (gbp).", e.message)
      end

      # Sync Records
      sync_records = get_sync_records_by_primary_id(sf_order.Id)
      assert_equal(1, sync_records.length)

      assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
      assert_equal(SyncRecordResolutionStatuses::ERROR.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    end
  end

end
