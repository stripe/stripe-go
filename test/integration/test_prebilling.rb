# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::Prebilling < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'prebills one year of two year subscription' do
    # Prebill one year of an annually billed two year subscription
    prebill_iterations = 1

    @user.field_defaults = {
      "subscription_schedule" => {
        "prebilling.iterations" => prebill_iterations,
      },
    }
    @user.save

    price = 1000_00
    subscription_term = 24
    start_date = now_time + 3.months

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      price: price,
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh

    sf_product = sf.find(SF_PRODUCT, sf_product_id)
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

    # basic customer creation check
    # without a 1:1 contact relationship, the email is nil
    refute_empty(customer.name)
    assert_nil(customer.email)

    assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
    assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)

    # top level subscription fields
    assert_match(sf_order.Id, subscription_schedule.metadata['salesforce_order_link'])
    assert_equal(subscription_schedule.metadata['salesforce_order_id'], sf_order.Id)

    # line-level subscription phase data
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)
    assert_match(sf_order.Id, phase.metadata['salesforce_order_link'])
    assert_equal(phase.metadata['salesforce_order_id'], sf_order.Id)
    # NOTE iterations does not exist on the phase! https://jira.corp.stripe.com/browse/PLATINT-1479
    # TODO I have no idea why the math requires rounding here. This doesn't make any sense. https://jira.corp.stripe.com/browse/PLATINT-1480
    phase_iterations = ((phase.end_date - phase.start_date) / 1.month.to_f).round
    assert_equal(subscription_term, phase_iterations)

    assert_equal(1, phase.items.count)
    phase_item = T.must(phase.items.first)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.plan)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.price)
    assert_equal(1, phase_item.quantity)

    # Ensure a prebilling invoice has been created
    refute_nil(subscription_schedule.prebilling)

    assert_equal(start_date.to_i, subscription_schedule.prebilling.period_start)
    assert_equal((start_date + prebill_iterations.years).to_i, subscription_schedule.prebilling.period_end)

    invoice = Stripe::Invoice.retrieve(subscription_schedule.prebilling.invoice, @user.stripe_credentials)

    # One line per pre-billed period
    assert_equal(prebill_iterations, invoice.lines.data.length)
    assert_equal(price * prebill_iterations, invoice.total)

    line = invoice.lines.data.last
    assert_equal(1, line.quantity)
    assert_equal(price, line.amount)

    # lets ensure a duplicate subscription schedule is not created
    Stripe::SubscriptionSchedule.expects(:create).never

    SalesforceTranslateRecordJob.translate(@user, sf_order)
  end

  it 'prebills six months of a year subscription' do
    # Prebill 6 months of a monthly billed year subscription
    prebill_iterations = 6

    @user.field_defaults = {
      "subscription_schedule" => {
        "prebilling.iterations" => prebill_iterations,
      },
    }
    @user.save

    price = 10_00
    subscription_term = 12
    start_date = now_time + 3.months

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price)

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    sf_product = sf.find(SF_PRODUCT, sf_product_id)
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

    # basic customer creation check
    # without a 1:1 contact relationship, the email is nil
    refute_empty(customer.name)
    assert_nil(customer.email)

    assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
    assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)

    # top level subscription fields
    assert_match(sf_order.Id, subscription_schedule.metadata['salesforce_order_link'])
    assert_equal(subscription_schedule.metadata['salesforce_order_id'], sf_order.Id)

    # line-level subscription phase data
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)
    assert_match(sf_order.Id, phase.metadata['salesforce_order_link'])
    assert_equal(phase.metadata['salesforce_order_id'], sf_order.Id)
    # NOTE iterations does not exist on the phase! https://jira.corp.stripe.com/browse/PLATINT-1479
    # TODO I have no idea why the math requires rounding here. This doesn't make any sense. https://jira.corp.stripe.com/browse/PLATINT-1480
    phase_iterations = ((phase.end_date - phase.start_date) / 1.month.to_f).round
    assert_equal(subscription_term, phase_iterations)

    assert_equal(1, phase.items.count)
    phase_item = T.must(phase.items.first)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.plan)
    assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.price)
    assert_equal(1, phase_item.quantity)

    # Ensure a prebilling invoice has been created
    refute_nil(subscription_schedule.prebilling)

    assert_equal(start_date.to_i, subscription_schedule.prebilling.period_start)
    assert_equal((start_date + prebill_iterations.months).to_i, subscription_schedule.prebilling.period_end)

    invoice = Stripe::Invoice.retrieve(subscription_schedule.prebilling.invoice, @user.stripe_credentials)

    # One line per pre-billed period
    assert_equal(prebill_iterations, invoice.lines.data.length)
    assert_equal(price * prebill_iterations, invoice.total)

    line = invoice.lines.data.last
    assert_equal(1, line.quantity)
    assert_equal(price, line.amount)

    # lets ensure a duplicate subscription schedule is not created
    Stripe::SubscriptionSchedule.expects(:create).never

    SalesforceTranslateRecordJob.translate(@user, sf_order)
  end

  it 'has a float iteration value' do
     # Prebill one year of an annually billed two year subscription
     prebill_iterations = 1.0

     @user.field_defaults = {
       "subscription_schedule" => {
         "prebilling.iterations" => prebill_iterations,
       },
     }
     @user.save

     price = 1000_00
     subscription_term = 24
     start_date = now_time + 3.months

     # creating these directly so we have the IDs
     sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
       price: price,
       additional_product_fields: {
         CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
       }
     )

     sf_account_id = create_salesforce_account
     sf_order = create_salesforce_order(
       sf_product_id: sf_product_id,
       sf_account_id: sf_account_id,

       additional_quote_fields: {
         CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
         CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
       }
     )

     SalesforceTranslateRecordJob.translate(@user, sf_order)

     sf_order.refresh

     sf_product = sf.find(SF_PRODUCT, sf_product_id)
     sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

     stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
     subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
     customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

     # basic customer creation check
     # without a 1:1 contact relationship, the email is nil
     refute_empty(customer.name)
     assert_nil(customer.email)

     assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
     assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)

     # top level subscription fields
     assert_match(sf_order.Id, subscription_schedule.metadata['salesforce_order_link'])
     assert_equal(subscription_schedule.metadata['salesforce_order_id'], sf_order.Id)

     # line-level subscription phase data
     assert_equal(1, subscription_schedule.phases.count)
     phase = T.must(subscription_schedule.phases.first)
     assert_match(sf_order.Id, phase.metadata['salesforce_order_link'])
     assert_equal(phase.metadata['salesforce_order_id'], sf_order.Id)
     # NOTE iterations does not exist on the phase! https://jira.corp.stripe.com/browse/PLATINT-1479
     # TODO I have no idea why the math requires rounding here. This doesn't make any sense. https://jira.corp.stripe.com/browse/PLATINT-1480
     phase_iterations = ((phase.end_date - phase.start_date) / 1.month.to_f).round
     assert_equal(subscription_term, phase_iterations)

     assert_equal(1, phase.items.count)
     phase_item = T.must(phase.items.first)
     assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.plan)
     assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.price)
     assert_equal(1, phase_item.quantity)

     # Ensure a prebilling invoice has been created
     refute_nil(subscription_schedule.prebilling)

     assert_equal(start_date.to_i, subscription_schedule.prebilling.period_start)
     assert_equal((start_date + prebill_iterations.to_i.years).to_i, subscription_schedule.prebilling.period_end)

     invoice = Stripe::Invoice.retrieve(subscription_schedule.prebilling.invoice, @user.stripe_credentials)

     # One line per pre-billed period
     assert_equal(prebill_iterations, invoice.lines.data.length)
     assert_equal(price * prebill_iterations, invoice.total)

     line = invoice.lines.data.last
     assert_equal(1, line.quantity)
     assert_equal(price, line.amount)

     # lets ensure a duplicate subscription schedule is not created
     Stripe::SubscriptionSchedule.expects(:create).never

     SalesforceTranslateRecordJob.translate(@user, sf_order)
  end

  it 'prebilling iterations = 0 does not result in a Stripe API error' do
    # setup
    prebill_iterations = 0.0
    @user.field_defaults = {
      "subscription_schedule" => {
        "prebilling.iterations" => prebill_iterations,
      },
    }
    @user.save

    # translate the Salesforce order
    sf_order = create_salesforce_order(
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # confirm prebilling was not set
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_nil(subscription_schedule.prebilling)
  end
end
