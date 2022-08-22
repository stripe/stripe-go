# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::SyncRecords < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'sync records are updated to success even after initial failure' do
    subscription_term = 12
    start_date = now_time

    sf_order = create_salesforce_order(
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        # omit subscription term
      }
    )

    exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    sync_records = get_sync_records_by_primary_id(sf_order.Id)
    assert_equal(1, sync_records.count)
    sync_record = T.must(sync_records.first)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::ERROR.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    assert_match("The following required fields are missing from this Salesforce record: SBQQ__Quote__c.SBQQ__SubscriptionTerm__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])

    # fix the missing subscription term in order to test the sync record upsert & associated trigger
    sf_quote_id = sf_order[SF_ORDER_QUOTE]
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
    })

    # Retranslate
    SalesforceTranslateRecordJob.translate(@user, sf_order)

    # Assert Order Sync Record Success
    sync_records = get_sync_records_by_primary_id(sf_order.Id)
    assert_equal(1, sync_records.count)
    sync_record = T.must(sync_records.first)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
  end

  it 'updates secondary id failures when the primary succeeds' do
    @user.field_defaults = {
      "price" => {
        "billing_scheme" => "bad_value",
      },
    }
    @user.save

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price
    sf_order = create_subscription_order

    exception = assert_raises(StripeForce::Errors::UserError) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    sync_records = get_sync_records_by_primary_id(sf_order.Id)
    assert_equal(1, sync_records.count)
    sync_record = T.must(sync_records.first)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER_ITEM, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    refute_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::ERROR.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])

    # now let's fix the mapping and reprocess
    @user.update(field_defaults: {})

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    original_sync_record = sync_record

    # after success there would be two records: one for the order line, one for the order, both marked as success
    sync_records = get_sync_records_by_primary_id(sf_order.Id)
    assert_equal(2, sync_records.count)
    sync_records.each do |r|
      assert_equal(sf_order.Id, r[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
      assert_includes([SyncRecordResolutionStatuses::SUCCESS.serialize, SyncRecordResolutionStatuses::RESOLVED.serialize], r[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    end
  end
end
