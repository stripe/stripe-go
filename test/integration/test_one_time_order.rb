# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OneTimeOrderTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || DateTime.now.utc)

    @user = make_user(save: true)
  end

  it 'integrates an invoice order' do
    @user.update(field_mappings: {
      customer: {
        # if accounts are mapped to customer, there is no default email field
        "email": "Description",
      },
    })

    sf_account_id = create_salesforce_account(additional_fields: {
      # an email is required for creating an invoice without a payment method
      # and we mapped customer.email to Account.Description above
      "Description" => create_static_email(email: "test_email"),
    })

    sf_product_id, _ = salesforce_standalone_product_with_price
    start_date = now_time - 1.month

    sf_order = create_salesforce_order(
      sf_account_id: sf_account_id,
      sf_product_id: sf_product_id,
      contact_email: "one_time_invoice_order_0_2",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh
    stripe_invoice_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    invoice = Stripe::Invoice.retrieve(stripe_invoice_id, @user.stripe_credentials)
    customer = Stripe::Customer.retrieve(invoice.customer, @user.stripe_credentials)
    refute_empty(customer.email)
    assert_equal(1, invoice.lines.count)
    line = invoice.lines.first
    assert_equal("one_time", line.price.type)

    # check the one-time invoice period
    assert_equal(start_date.to_i, line.period.start)
    assert_equal((start_date + TEST_DEFAULT_CONTRACT_TERM.months).to_i, line.period.end)
  end
end
