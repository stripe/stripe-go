# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::EvergreenOrderTest < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'creates an evergreen order object' do
    @user.field_defaults['customer'] = {
      'email' => create_random_email,
    }

    sf_account_id = create_salesforce_account

    sf_order = create_evergreen_salesforce_order(
      sf_account_id: sf_account_id
    )

    # get the order items to check the subscription data inside
    order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

    order_items.each do |order_item|
      assert_equal(1, order_item[CPQ_QUOTE_SUBSCRIPTION_TERM])
      assert_equal(CPQProductSubscriptionTypeOptions::EVERGREEN.serialize, order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE])
      assert_equal('Fixed Price', order_item[CPQ_QUOTE_SUBSCRIPTION_PRICING])
    end
  end

  describe 'temporary failure case test for evergreen orders' do
    it 'raises error when evergreen order is encountered' do
      @user.update(field_defaults: {
        customer: {
          # if accounts are mapped to customer, there is no default email field
          "email": create_random_email,
        },
      })

      sf_account_id = create_salesforce_account

      sf_order = create_evergreen_salesforce_order(
        sf_account_id: sf_account_id,

        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => 1,
        }
      )

      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Evergreen Salesforce Orders are not yet supported.", exception.message)
    end
  end

  describe 'failure cases' do
    it 'raises user error when there are evergreen and renewable items on same order' do
      @user.update(field_defaults: {
        customer: {
          # if accounts are mapped to customer, there is no default email field
          "email": create_random_email,
        },
      })

      sf_account_id = create_salesforce_account

      sf_product_id_1, _ = salesforce_evergreen_product_with_price
      sf_product_id_2, _ = salesforce_recurring_product_with_price

      quote_id = create_salesforce_quote(
        sf_account_id: sf_account_id,

        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => 1,
        }
      )

      # add both products to the sf quote
      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
      calculate_and_save_cpq_quote(quote_with_product)

      quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
      calculate_and_save_cpq_quote(quote_with_product)

      sf_order = create_order_from_cpq_quote(quote_id)

      order_items = sf_get_related(sf_order, SF_ORDER_ITEM)

      exception = assert_raises(Integrations::Errors::UserError) do
        SalesforceTranslateRecordJob.translate(@user, sf_order)
      end

      assert_match("Salesforce orders with both Evergreen and Renewable items are not yet supported.", exception.message)
    end
  end
end
