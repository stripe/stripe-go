# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::CustomerTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

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

  it 'updates an existing stripe customer after the salesforce account info has changed' do
    @user.enable_feature(FeatureFlags::UPDATE_CUSTOMER_ON_ORDER_TRANSLATION)

    # setup
    sf_account_initial_name = "Test Account Name"
    sf_account_initial_description = "A CRM company"
    sf_account_updated_name = "Test Account Name 2"
    sf_account_updated_description = "A CRM company 2"

    sf_account_id = create_salesforce_account(additional_fields: {
      "Name" => sf_account_initial_name,
      "Phone" => "1231231234",
      "Description" => sf_account_initial_description,
      "ShippingStreet" => "500 Briar Forest Drive",
      "ShippingCity" => "Houston",
      "ShippingState" => "Texas",
      "ShippingPostalCode" => "77077",
      "ShippingCountry" => "US",
    })

    # translate sf account
    StripeForce::Translate.perform_inline(@user, sf_account_id)

    # confirm the sf account was translated
    sf_account = sf.find(SF_ACCOUNT, sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_customer_id)

    # retrieve the corresponding stripe customer
    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)

    # update the sf account info
    sf.update!(SF_ACCOUNT, {
      SF_ID => sf_account_id,
      "Name" => sf_account_updated_name,
      "Description" => sf_account_updated_description,
    })

    # sanity check that stripe customer info is the original info
    assert_equal(sf_account_id, stripe_customer.metadata['salesforce_account_id'])
    assert_equal(sf_account_initial_name, stripe_customer.name)
    assert_equal(sf_account_initial_description, stripe_customer.description)

    # retranslate and confirm stripe customer info is updated
    StripeForce::Translate.perform_inline(@user, sf_account_id)
    stripe_customer.refresh

    assert_equal(sf_account_updated_name, stripe_customer.name)
    assert_equal(sf_account_updated_description, stripe_customer.description)
  end

  it 'updates stripe customer on order translation only when feature is enabled' do
      # setup
      sf_account_initial_name = "Test Account Name"
      sf_account_initial_postal_code = "77077"
      sf_account_updated_name = "Test Account Name 2"
      sf_account_updated_postal_code = "77048"

      sf_account_id = create_salesforce_account(additional_fields: {
        "Name" => sf_account_initial_name,
        "Phone" => "1231231234",
        "Description" => "A CRM company",
        "ShippingStreet" => "500 Briar Forest Drive",
        "ShippingCity" => "Houston",
        "ShippingState" => "Texas",
        "ShippingPostalCode" => sf_account_initial_postal_code,
        "ShippingCountry" => "US",
      })

      # translate the initial sf account
      StripeForce::Translate.perform_inline(@user, sf_account_id)

      # confirm the sf account was translated
      sf_account = sf.find(SF_ACCOUNT, sf_account_id)
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_customer_id)

      # retrieve the corresponding stripe customer
      stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)

      # sanity check that stripe customer info is the original info
      assert_equal(sf_account_id, stripe_customer.metadata['salesforce_account_id'])
      assert_equal(sf_account_initial_name, stripe_customer.name)
      assert_equal(sf_account_initial_postal_code, T.must(stripe_customer.shipping.address).postal_code)

      # update the sf account info
      sf.update!(SF_ACCOUNT, {
        SF_ID => sf_account_id,
        "Name" => sf_account_updated_name,
        "ShippingPostalCode" => sf_account_updated_postal_code,
      })

      # create an sf order and translate the order
      sf_order_1 = create_salesforce_order(
        sf_account_id: sf_account_id,
        contact_email: "updates_stripe_customer_1",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
        }
      )
      SalesforceTranslateRecordJob.translate(@user, sf_order_1)
      stripe_customer.refresh

      # confirm that customer info was NOT updated on order translation
      # since feature_flag UPDATE_CUSTOMER_ON_ORDER_TRANSLATION isn't enabled
      assert_equal(sf_account_initial_name, stripe_customer.name)
      assert_equal(sf_account_initial_postal_code, T.must(stripe_customer.shipping.address).postal_code)

      # now enable feature flag UPDATE_CUSTOMER_ON_ORDER_TRANSLATION
      @user.enable_feature(FeatureFlags::UPDATE_CUSTOMER_ON_ORDER_TRANSLATION)
      @user.save

      # create another sf order and translate the order
      sf_order_2 = create_salesforce_order(
        sf_account_id: sf_account_id,
        contact_email: "updates_stripe_customer_2",
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
        }
      )
      SalesforceTranslateRecordJob.translate(@user, sf_order_2)
      stripe_customer.refresh

      # confirm stripe customer info was updated on order translation
      # since feature flag is enabled
      assert_equal(sf_account_updated_name, stripe_customer.name)
      assert_equal(sf_account_updated_postal_code, T.must(stripe_customer.shipping.address).postal_code)
  end

  it 'ensures customer metadata is merged on customer updates' do
      @user.enable_feature(FeatureFlags::UPDATE_CUSTOMER_ON_ORDER_TRANSLATION)

      # add a custom metadata field which maps to the sf account name
      @user.field_mappings = {
        "customer" => {
          "metadata.sf_mapped_metadata_field_1" => "Name",
          "metadata.sf_mapped_metadata_field_2" => "Description",
        },
      }

      # create the sf account
      sf_account_name = "Awesome Test Account Name"
      sf_account_description = "Awesome Test Account Description"
      sf_account_id = create_salesforce_account(additional_fields: {
        "Name" => sf_account_name,
        "Description" => sf_account_description,
      })

      # translate the sf account
      StripeForce::Translate.perform_inline(@user, sf_account_id)

      # confirm the sf account was translated
      sf_account = sf.find(SF_ACCOUNT, sf_account_id)
      stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_customer_id)

      # retrieve the corresponding stripe customer
      stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)

      # confirm the stripe customer metadata includes our custom metadata and the sf mapped metadata
      assert_equal(4, stripe_customer.metadata.keys.count)
      assert_equal(sf_account_id, stripe_customer.metadata['salesforce_account_id'])
      assert_match(sf_account_id, stripe_customer.metadata['salesforce_account_link'])
      assert_match(sf_account_name, stripe_customer.metadata['sf_mapped_metadata_field_1'])
      assert_match(sf_account_description, stripe_customer.metadata['sf_mapped_metadata_field_2'])

      # update the salesforce customer name
      updated_sf_account_name = "Incredible Test Account Name"
      sf.update!(SF_ACCOUNT, {
        SF_ID => sf_account_id,
        "Name" => updated_sf_account_name,
      })

      # add a custom metadata field on the stripe customer
      stripe_customer.metadata["stripe_custom_metadata_field"] = "stripe_custom_metadata_value"
      stripe_customer.save

      # retranslate the sf account
      StripeForce::Translate.perform_inline(@user, sf_account_id)
      stripe_customer.refresh

      # confirm the newly added stripe metadata field has been added
      # and the sf account name has been updated
      assert_equal(5, stripe_customer.metadata.keys.count)
      assert_equal(sf_account_id, stripe_customer.metadata['salesforce_account_id'])
      assert_match(sf_account_id, stripe_customer.metadata['salesforce_account_link'])
      assert_equal("stripe_custom_metadata_value", stripe_customer.metadata['stripe_custom_metadata_field'])
      assert_match(updated_sf_account_name, stripe_customer.metadata['sf_mapped_metadata_field_1'])
      assert_match(sf_account_description, stripe_customer.metadata['sf_mapped_metadata_field_2'])

      # update the user mappings to remove the 2nd mapping
      @user.field_mappings = {
        "customer" => {
          "metadata.sf_mapped_metadata_field_1" => "Name",
        },
      }

      # retranslate the sf account
      StripeForce::Translate.perform_inline(@user, sf_account_id)
      stripe_customer.refresh

      # confirm this key is not wiped
      # in order to remove a metadat field, you must set the value to empty
      assert_equal(5, stripe_customer.metadata.keys.count)
      assert_match(sf_account_description, stripe_customer.metadata['sf_mapped_metadata_field_2'])
  end

  it 'sync customer with invoice_settings.custom_fields' do
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
    @user.enable_feature FeatureFlags::UPDATE_CUSTOMER_ON_ORDER_TRANSLATION, update: true
    @user.field_defaults = {
      "customer" => {
        "invoice_settings.custom_fields.name" => 'key1',
        "invoice_settings.custom_fields.value" => 'value1',
      },
    }
    @user.save

    # create and sync the sf account
    sf_account_id = create_salesforce_account
    StripeForce::Translate.perform_inline(@user, sf_account_id)

    # confirm the sf account was translated
    sf_account = sf.find(SF_ACCOUNT, sf_account_id)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    refute_nil(stripe_customer_id)

    # retrieve the corresponding stripe customer and confirm the stripe customer invoice settings fields include our custom data
    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    assert_equal(1, T.must(stripe_customer.invoice_settings.custom_fields).count)
    invoice_settings_custom_fields = stripe_customer.invoice_settings.custom_fields
    assert_equal({"name": "key1", "value": "value1"}, T.must(invoice_settings_custom_fields).first.to_h)

    # let's create an Order with this account and ensure the corresponding Invoice has this data
    sf_product_id, _sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
      }
    )
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,
      contact_email: "invoice_settings_custom_fields_8",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
      }
    )
    StripeForce::Translate.perform_inline(@user, sf_order.Id)
    sf_order.refresh

    # get Stripe customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)
    advance_test_clock(stripe_customer, (now_time + 1.day).to_i)

    # sanity check the initial invoice
    invoices = Stripe::Invoice.list({customer: stripe_customer.id}, @user.stripe_credentials).data
    assert_equal(1, invoices.length)
    initial_invoice = invoices.first
    assert_equal(1, initial_invoice.custom_fields.count)
    assert_equal({"name": "key1", "value": "value1"}, T.must(initial_invoice.custom_fields).first.to_h)

    # clear the invoice settings and resync
    @user.field_defaults = {
      "customer" => {
        "invoice_settings.custom_fields.name" => 'key2',
        "invoice_settings.custom_fields.value" => 'value2',
      },
    }
    @user.save
    StripeForce::Translate.perform_inline(@user, sf_account_id)
    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    assert_equal(1, T.must(stripe_customer.invoice_settings.custom_fields).count)
    assert_equal({"name": "key2", "value": "value2"}, T.must(stripe_customer.invoice_settings.custom_fields).first.to_h)
  end
end
