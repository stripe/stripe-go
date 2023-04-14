# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::InvoiceRendering < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    @user.enable_feature(StripeForce::Constants::FeatureFlags::INVOICE_RENDERING_TEMPLATE)
  end

  it 'translates a sf order with an invoice rendering template' do
    invoice_rendering_template_id = create_invoice_rendering_template
    @user.field_defaults = {
      "subscription_schedule" => {
        "default_settings.invoice_settings.rendering.template" => invoice_rendering_template_id,
      },
    }
    @user.save

    # create and translate a Salesforce order
    sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price
    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # fetch the generated subscription and verify the invoice rendering template is set
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(invoice_rendering_template_id, subscription_schedule.default_settings.invoice_settings.rendering.template)

    # fetch invoice and verify the invoice rendering template is used
    sf_account = sf_get(sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)

    assert_equal(invoice_rendering_template_id, invoice.data.first.rendering.template)
  end

  it 'translates a sf order with an invoice rendering template for a specific version' do
    invoice_rendering_template_id = create_invoice_rendering_template
    @user.field_defaults = {
      "subscription_schedule" => {
        "default_settings.invoice_settings.rendering.template" => invoice_rendering_template_id,
        "default_settings.invoice_settings.rendering.template_version" => 1,
      },
    }
    @user.save

    # create and translate a Salesforce order
    sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price
    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # fetch the generated subscription and verify the invoice rendering template is set
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    assert_equal(invoice_rendering_template_id, subscription_schedule.default_settings.invoice_settings.rendering.template)

    # fetch invoice and verify the invoice rendering template is used
    sf_account = sf_get(sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    invoice = Stripe::Invoice.list({customer: stripe_customer_id}, @user.stripe_credentials)

    assert_equal(invoice_rendering_template_id, invoice.data.first.rendering.template)
    assert_equal(1, invoice.data.first.rendering.template_version)
  end

  it 'translates a subscription with an invalid invoice rendering template' do
    invoice_rendering_template_id = "invalid_id_does_not_exist"
    @user.field_defaults = {
      "subscription_schedule" => {
        "default_settings.invoice_settings.rendering.template" => invoice_rendering_template_id,
      },
    }
    @user.save

    # create and translate a Salesforce order
    sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price
    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,

      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    # Verify Stripe::Error is thrown
    template_does_not_exist_error = assert_raises(Stripe::InvalidRequestError) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end
    assert_match('no such invoice rendering template', template_does_not_exist_error.message.downcase)
  end
end
