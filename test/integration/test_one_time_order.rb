# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OneTimeOrderTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'integrates a invoice order' do
    @user.update(field_mappings: {
      customer: {
        # if accounts are mapped to customer, there is no default email field
        "email": "Description",
      },
    })

    sf_account_id = create_salesforce_account(additional_fields: {
      # an email is required for creating an invoice without a payment method
      "Description" => create_random_email,
    })

    sf_product_id, _ = salesforce_standalone_product_with_price

    sf_order = create_salesforce_order(
      sf_account_id: sf_account_id,
      sf_product_id: sf_product_id
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
  end
end
