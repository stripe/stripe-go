# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    # Only run on Multi-Currency Enabled CI Accounts
    # If you are running this locally on a multi-currency scratch org, please go into make_user
    #   and add your scratch org ID similar to brennen's
    log.info "Multi-currency test user details", is_multicurrency_org: @user.is_multicurrency_org?, sf_account_id: @user.salesforce_account_id
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

    it 'translates an order with a coupon in a currency other than USD' do
      @user.enable_feature(FeatureFlags::COUPONS)

      # add a custom metadata field mapping
      @user.field_mappings['coupon'] = {
        'metadata.sf_coupon_mapped_metadata_field' => prefixed_stripe_field('Name__c'),
        'metadata.sf_order_mapped_metadata_field' => prefixed_stripe_field('Order_Item__c.OrderId'),
      }
      @user.save

      currency_iso_code = 'GBP'

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(currency_iso_code: currency_iso_code)

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, currency_iso_code: currency_iso_code, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # create a quote with a product
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
      sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # retrieve the quote line
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      assert_equal(1, quote_lines.size)
      quote_line_id = quote_lines.first.Id

      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '£10 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 10,
        SalesforceStripeCouponFields::CURRENCY_ISO_CODE => currency_iso_code,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)

      # the first phase item should have two coupons
      first_phase_item = T.must(first_phase.items.first)
      discounts = first_phase_item.discounts
      assert_equal(1, discounts.count)

      # retrieve the two stripe coupons
      stripe_coupon = Stripe::Coupon.retrieve(T.must(discounts.first).coupon, @user.stripe_credentials)

      # sanity check both stripe coupons have the right data
      assert_equal(10 * 100, stripe_coupon.amount_off)
      assert_equal(currency_iso_code.downcase, stripe_coupon.currency)
      assert_equal("once", stripe_coupon.duration)

      sf_amount_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_amount_off_coupon_id)
      assert_equal(sf_amount_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])
      assert_equal('£10 off coupon', stripe_coupon.metadata['sf_coupon_mapped_metadata_field'])
      assert_equal(sf_order.Id, stripe_coupon.metadata['sf_order_mapped_metadata_field'])

      # confirm the stripe coupon ids were written back to the quote coupons in salesforce
      sf_amount_off_coupon = sf_get(sf_amount_off_coupon_id)
      assert_equal(stripe_coupon.id, sf_amount_off_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end

    it 'allows association of a percent off coupon with an quote in a different currency' do
      @user.enable_feature(FeatureFlags::COUPONS)

      # add a custom metadata field mapping
      @user.field_mappings['coupon'] = {
        'metadata.sf_coupon_mapped_metadata_field' => prefixed_stripe_field('Name__c'),
        'metadata.sf_order_mapped_metadata_field' => prefixed_stripe_field('Order_Item__c.OrderId'),
      }
      @user.save

      quote_currency_iso_code = 'GBP'
      coupon_currency_iso_code = 'USD'

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(currency_iso_code: quote_currency_iso_code)

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, currency_iso_code: quote_currency_iso_code, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # create a quote with a product
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
      sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # retrieve the quote line
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      assert_equal(1, quote_lines.size)
      quote_line_id = quote_lines.first.Id

      sf_percent_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '20% off coupon',
        SalesforceStripeCouponFields::PERCENT_OFF => 20,
        SalesforceStripeCouponFields::CURRENCY_ISO_CODE => coupon_currency_iso_code,
      })

      # create the quote line coupon association object to map the coupons to the quote line
      create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_percent_off_coupon_id)

      # create and translate the SF order
      sf_order = create_order_from_cpq_quote(sf_quote_id)
      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      # fetch the stripe subscription schedule
      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
      assert_equal(1, subscription_schedule.phases.count)

      # first phase should have one item
      first_phase = T.must(subscription_schedule.phases.first)
      assert_equal(1, first_phase.items.count)

      # the first phase item should have two coupons
      first_phase_item = T.must(first_phase.items.first)
      discounts = first_phase_item.discounts
      assert_equal(1, discounts.count)

      # retrieve the two stripe coupons
      stripe_coupon = Stripe::Coupon.retrieve(T.must(discounts.first).coupon, @user.stripe_credentials)

      assert_equal(20, stripe_coupon.percent_off)
      assert_equal("once", stripe_coupon.duration)

      sf_percent_off_order_coupon = get_sf_order_coupon_from_quote_coupon_id(sf_percent_off_coupon_id)
      assert_equal(sf_percent_off_order_coupon.Id, stripe_coupon.metadata['salesforce_order_stripe_coupon_id'])
      assert_equal('20% off coupon', stripe_coupon.metadata['sf_coupon_mapped_metadata_field'])
      assert_equal(sf_order.Id, stripe_coupon.metadata['sf_order_mapped_metadata_field'])

      sf_percent_off_coupon = sf_get(sf_percent_off_coupon_id)
      assert_equal(stripe_coupon.id, sf_percent_off_coupon[prefixed_stripe_field(GENERIC_STRIPE_ID)])
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

    it 'fails to create associate a coupon with an quote in a different currency' do
      @user.enable_feature(FeatureFlags::COUPONS)

      # add a custom metadata field mapping
      @user.field_mappings['coupon'] = {
        'metadata.sf_coupon_mapped_metadata_field' => prefixed_stripe_field('Name__c'),
        'metadata.sf_order_mapped_metadata_field' => prefixed_stripe_field('Order_Item__c.OrderId'),
      }
      @user.save

      quote_currency_iso_code = 'GBP'
      coupon_currency_iso_code = 'USD'

      # setup
      sf_account_id = create_salesforce_account
      sf_product_id, _sf_pricebook_id = salesforce_recurring_product_with_price(currency_iso_code: quote_currency_iso_code)

      # create a CPQ quote
      sf_quote_id = create_salesforce_quote(sf_account_id: sf_account_id, currency_iso_code: quote_currency_iso_code, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      })

      # create a quote with a product
      quote_with_product = add_product_to_cpq_quote(sf_quote_id, sf_product_id: sf_product_id)
      sf_quote_id = calculate_and_save_cpq_quote(quote_with_product)

      # retrieve the quote line
      quote_lines = sf_get_related(sf_quote_id, CPQ_QUOTE_LINE)
      assert_equal(1, quote_lines.size)
      quote_line_id = quote_lines.first.Id

      sf_amount_off_coupon_id = create_salesforce_stripe_coupon(additional_fields: {
        SalesforceStripeCouponFields::NAME => '£10 off coupon',
        SalesforceStripeCouponFields::AMOUNT_OFF => 10,
        SalesforceStripeCouponFields::CURRENCY_ISO_CODE => coupon_currency_iso_code,
      })

      exception = assert_raises(Restforce::ErrorCode::FieldCustomValidationException) do
        # create the quote line coupon association object to map the coupons to the quote line
        create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id: quote_line_id, sf_stripe_coupon_id: sf_amount_off_coupon_id)
      end

      assert_equal(
        "FIELD_CUSTOM_VALIDATION_EXCEPTION: The Currency of the Quote (GBP) does not match the Currency of the Quote Stripe Coupon (USD).",
        exception.message
      )
    end
  end

end
