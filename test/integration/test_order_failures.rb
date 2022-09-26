# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::OrderFailureTest < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'does not allow a quantity greater than 1 on metered billing items' do
    sf_product_id, sf_pricebook_id = salesforce_recurring_metered_produce_with_price
    sf_account_id ||= create_salesforce_account
    quote_id = create_salesforce_quote(
      sf_account_id: sf_account_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        # one year / 12 months
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      }
    )

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

    # bump quantity to 2, which should trigger an error
    quote_with_product["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 2

    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    exception = assert_raises(Integrations::Errors::UserError) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    assert_match("Order lines with a price configured for metered billing cannot have a quantity greater than 1.", exception.message)
  end

  it 'creates a sync failure when the start date does not exist' do
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
        # intentional ommission of the start date
      }
    )

    exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    sync_record = get_sync_record_by_secondary_id(sf_order.Id)

    assert_match("The following required fields are missing from this Salesforce record: SBQQ__Quote__c.SBQQ__StartDate__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])
  end

  # it looks like there are field validations in place to protect against the term being nil'd out on the order line
  it 'creates a user error when the subscription term does not exist' do
    sf_order = create_salesforce_order(additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      # omit subscription term
    })

    exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    sync_record = get_sync_record_by_secondary_id(sf_order.Id)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_RECORD_ID.serialize)])
    assert_match(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::COMPOUND_ID.serialize)])
    assert_match("The following required fields are missing from this Salesforce record: SBQQ__Quote__c.SBQQ__SubscriptionTerm__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])

    assert_match(@user.salesforce_instance_url, sync_record[prefixed_stripe_field('Primary_Record__c')])
    assert_match(@user.salesforce_instance_url, sync_record[prefixed_stripe_field('Secondary_Record__c')])
  end

  it 'throws an error when a float is specified for a quantity' do
    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price
    sf_account_id ||= create_salesforce_account

    quote_id = create_salesforce_quote(
      sf_account_id: sf_account_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        # one year / 12 months
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      }
    )

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

    # bump quantity to a float value
    quote_with_product["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = "1.5"

    calculate_and_save_cpq_quote(quote_with_product)
    sf_order = create_order_from_cpq_quote(quote_id)

    exception = assert_raises(Integrations::Errors::UserError) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    assert_match("Quantity specified as a decimal value. Only integers are supported", exception.message)
  end

  it 'it creates an error sync record when the customer has been deleted' do

    sf_account_id = create_salesforce_account

    sf_order = create_subscription_order(sf_account_id: sf_account_id)
    sf_account = @user.sf_client.find(SF_ACCOUNT, sf_account_id)

    SalesforceTranslateRecordJob.translate(@user, sf_account)

    sf_account.refresh

    stripe_customer = Stripe::Customer.retrieve(sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials).delete
    stripe_customer.refresh

    assert(true, stripe_customer.deleted?)

    exception = assert_raises(Integrations::Errors::UserError) do
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
    end

    assert_match("During translation we attempted to fetch a related Stripe::Customer record with ID #{stripe_customer.id}, but found that it was deleted.", exception.message)

    # Sync Records
    sync_records = get_sync_records_by_primary_id(sf_order.Id)
    assert_equal(1, sync_records.length)

    assert_equal(SF_ORDER, sync_records.first[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SyncRecordResolutionStatuses::ERROR.serialize, sync_records.first[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
  end
end
