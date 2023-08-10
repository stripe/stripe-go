# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OneTimeOrderTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    if !VCR.current_cassette.originally_recorded_at.nil?
      Timecop.freeze(VCR.current_cassette.originally_recorded_at)
    end

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
      "Description" => create_static_email(email: "test"),
    })

    sf_product_id, _ = salesforce_standalone_product_with_price

    # sf_order = create_salesforce_order(
    #   sf_account_id: sf_account_id,
    #   sf_product_id: sf_product_id,
    #   contact_email: "integrates_invoice_order"
    # )

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id,
                                       contact_email: "integrates_invoice_order"
                                      )

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)
    calculate_and_save_cpq_quote(quote_with_product)


    sf_order = create_order_from_cpq_quote(quote_id)

    # Timecop.freeze(Time.now + 1.minute)
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

  it 'allows the same pricebook entry to be used multiple times on one-time invoices' do

  end
end
