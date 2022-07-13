# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::BillingFrequencyTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  describe 'errors' do
    # TODO this will be supported in the future
    it 'throws an error when billing frequency of items do not match' do
      # one monthly and one annual price
      sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
      sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price(
        additional_product_fields: {
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        }
      )

      sf_account_id = create_salesforce_account

      quote_id = create_salesforce_quote(sf_account_id: sf_account_id, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      })

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      # Stripe should throw an error when this occurs and we should display it to the user
      assert_raises(Stripe::InvalidRequestError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      sync_record = get_sync_record_by_primary_id(sf_order.Id)
      assert_match('prices must have the same `interval`', sync_record[prefixed_stripe_field(SyncRecordFields::RESOLUTION_MESSAGE.serialize)])
    end
  end
end
