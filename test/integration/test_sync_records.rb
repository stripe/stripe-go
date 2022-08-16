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
      }
    )

    exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
      SalesforceTranslateRecordJob.translate(@user, sf_order)
    end

    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)

    sync_record = get_sync_record_by_secondary_id(order_items.first.Id)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER_ITEM, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::ERROR.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
    assert_match("The following required fields are missing from this Salesforce record: Order.SBQQ__Quote__c.SBQQ__SubscriptionTerm__c", sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])


    # fix the missing subscription term in order to test the sync record upsert & associated trigger
    sf_quote_id = sf_order[SF_ORDER_QUOTE]
    sf.update!(CPQ_QUOTE, {
      SF_ID => sf_quote_id,
      CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
    })

    sf_order = create_order_from_cpq_quote(sf_quote_id)

    # Retranslate
    SalesforceTranslateRecordJob.translate(@user, sf_order)

    # Assert Order Sync Record Success
    sync_record = get_sync_record_by_secondary_id((sf_order.Id))

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::SUCCESS.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])

    # Assert Order Item Sync Record Success
    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(1, order_items.size)

    sync_record = get_sync_record_by_secondary_id(order_items.first.Id)

    assert_equal(SF_ORDER, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize)])
    assert_equal(SF_ORDER_ITEM, sync_record[prefixed_stripe_field(SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize)])

    assert_equal(sf_order.Id, sync_record[prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)])
    assert_equal(SyncRecordResolutionStatuses::RESOLVED.serialize, sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_STATUS.serialize)])
  end
end
