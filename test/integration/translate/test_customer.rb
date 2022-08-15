# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::CustomerTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'translates billing address' do
    sf_account_id = create_salesforce_account(additional_fields: {
      "BillingStreet" => "500 Westover Drive",
      "BillingCity" => "Sanford",
      "BillingState" => "North Carolina",
      "BillingPostalCode" => "27330",
      "BillingCountry" => "US",
    })

    StripeForce::Translate.perform_inline(@user, sf_account_id)

    sf_account = sf.find(SF_ACCOUNT, sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_customer_id)

    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    assert_equal(sf_account.Id, stripe_customer.metadata['salesforce_account_id'])

    assert_equal("500 Westover Drive", stripe_customer.address&.line1)
    assert_nil(stripe_customer.address&.line2)
    assert_equal("Sanford", stripe_customer.address&.city)
    assert_equal("North Carolina", stripe_customer.address&.state)
    assert_equal("27330", stripe_customer.address&.postal_code)
    assert_equal("US", stripe_customer.address&.country)

    assert_nil(stripe_customer.shipping)
  end

  it 'translates shipping address and standard fields' do
    sf_account_id = create_salesforce_account(additional_fields: {
      "Name" => "Sales force",
      "Phone" => "1231231234",
      "Description" => "A CRM company",

      "ShippingStreet" => "500 Westover Drive",
      "ShippingCity" => "Sanford",
      "ShippingState" => "North Carolina",
      "ShippingPostalCode" => "27330",
      "ShippingCountry" => "US",
    })

    StripeForce::Translate.perform_inline(@user, sf_account_id)

    sf_account = sf.find(SF_ACCOUNT, sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_customer_id)

    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    assert_equal(sf_account.Id, stripe_customer.metadata['salesforce_account_id'])
    assert_nil(stripe_customer.test_clock)

    assert_nil(stripe_customer.address)

    assert_equal("Sales force", stripe_customer.name)
    assert_equal("A CRM company", stripe_customer.description)

    assert_equal("1231231234", stripe_customer.phone)
    assert_equal("1231231234", stripe_customer.shipping.phone)

    assert_equal("500 Westover Drive", stripe_customer.shipping.address&.line1)
    assert_nil(stripe_customer.shipping.address&.line2)
    assert_equal("Sanford", stripe_customer.shipping.address&.city)
    assert_equal("North Carolina", stripe_customer.shipping.address&.state)
    assert_equal("27330", stripe_customer.shipping.address&.postal_code)
    assert_equal("US", stripe_customer.shipping.address&.country)
  end

  it 'attaches a test clock to a customer when enabled' do
    @user.enable_feature(FeatureFlags::TEST_CLOCKS)

    sf_account_id = create_salesforce_account

    StripeForce::Translate.perform_inline(@user, sf_account_id)

    sf_account = sf.find(SF_ACCOUNT, sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_customer_id)

    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    refute_nil(stripe_customer.test_clock)
  end
end
