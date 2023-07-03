# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::EvergreenOrderTest < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true
    @user.field_defaults = {
      "customer" => {
        "email" => create_random_email,
      },
      "subscription" => {
        "days_until_due" => 30,
        "collection_method" => "send_invoice",
      },
    }
    @user.save
  end

  it 'cancels an evergreen subscription now' do
    # creates evergreen order and cancels immediately

    sf_order = create_evergreen_salesforce_order(
      # need to set these fields explicitly to use translate
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    sf_contract = create_contract_from_order(sf_order)
    # api precondition: initial orders have a nil contract ID
    sf_order.refresh
    assert_nil(sf_order.ContractId)

    # the contract should reference the initial order that was created
    assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

    amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

    # wipe out the product
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = now_time_formatted_for_salesforce
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

    sf_order_amendment = create_order_from_quote_data(amendment_quote)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

    canceled_subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal("canceled", canceled_subscription.status)
    assert_equal(sf_order.Id, canceled_subscription.metadata['salesforce_order_id'])
  end

  it 'cancels an evergreen subscription in the future' do
    # set cancel time for future and amend, then move test clock to after that and ensure termination
    current_time = now_time

    sf_order = create_evergreen_salesforce_order(
      # need to set these fields explicitly to use translate
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(current_time),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }
    )

    # translate salesforce order to subscription and find
    SalesforceTranslateRecordJob.translate(@user, sf_order)
    sf_order.refresh

    # get Stripe customer
    sf_account = sf_get(sf_order['AccountId'])
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = stripe_get(stripe_customer_id)
    refute_nil(stripe_customer.test_clock)

    # get stripe subscription
    sf_order = sf.find(SF_ORDER, sf_order.Id)
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    sf_contract = create_contract_from_order(sf_order)
    # api precondition: initial orders have a nil contract ID
    sf_order.refresh
    assert_nil(sf_order.ContractId)

    # the contract should reference the initial order that was created
    assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

    current_time = Time.at(StripeForce::Translate::OrderAmendment.determine_current_time(@user, stripe_customer.id))

    order_end_date = format_date_for_salesforce(current_time + 3.day)
    amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

    # wipe out the product
    amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = order_end_date
    amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

    sf_order_amendment = create_order_from_quote_data(amendment_quote)
    assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

    StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
    sf_order.refresh

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    assert_equal("active", subscription.status)
    assert_equal(Time.parse(order_end_date).to_i, subscription.cancel_at)

    # Fast forward few days to observe cancelation
    test_clock = advance_test_clock(stripe_customer, (current_time + 5.day).to_i)

    canceled_subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

    # if the cancelation was in the second billing period or beyond,
    # will "past_due" instead of "canceled" because invoices are not being paid
    assert_equal("canceled", canceled_subscription.status)
    assert_equal(sf_order.Id, canceled_subscription.metadata['salesforce_order_id'])
  end

  describe 'failure cases' do
    it 'salesforce raises error when attempt to amend evergreen order with non zero quantity' do
      sf_order = create_evergreen_salesforce_order(
        # need to set these fields explicitly to use translate
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      # api precondition: initial orders have a nil contract ID
      sf_order.refresh
      assert_nil(sf_order.ContractId)

      # the contract should reference the initial order that was created
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      amendment_end_date = now_time
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

      # attempt to set amendment quantity > 1
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 3
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_end_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      exception = assert_raises(Restforce::ErrorCode::ApexError) do
        create_order_from_quote_data(amendment_quote)
      end

      assert_equal("APEX_ERROR: SBQQ.ValidationException: [\"The quantity field must be set to 0 for evergreen subscriptions.\"]\n\n(System Code)", exception.message)
    end

    it 'processes the first amendment and throws user error beyond that' do
      # creates evergreen order and cancels immediately

      sf_order = create_evergreen_salesforce_order(
        # need to set these fields explicitly to use translate
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }
      )

      sf_contract = create_contract_from_order(sf_order)
      # api precondition: initial orders have a nil contract ID
      sf_order.refresh
      assert_nil(sf_order.ContractId)

      # the contract should reference the initial order that was created
      assert_equal(sf_order[SF_ID], sf_contract[SF_CONTRACT_ORDER_ID])

      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

      # wipe out the product
      amendment_quote["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = now_time_formatted_for_salesforce
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      amendment_quote_1 = create_quote_data_from_contract_amendment(sf_contract)

      # wipe out the product
      amendment_quote_1["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_quote_1["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(now_time + 20.day)
      amendment_quote_1["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      sf_order_amendment_1 = create_order_from_quote_data(amendment_quote_1)
      assert_equal(sf_order_amendment_1.Type, OrderTypeOptions::AMENDMENT.serialize)

      amendment_quote_2 = create_quote_data_from_contract_amendment(sf_contract)

      # wipe out the product
      amendment_quote_2["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 0
      amendment_quote_2["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(now_time + 10.day)
      amendment_quote_2["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      sf_order_amendment_2 = create_order_from_quote_data(amendment_quote_2)
      assert_equal(sf_order_amendment_2.Type, OrderTypeOptions::AMENDMENT.serialize)

      exception = assert_raises(Integrations::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment_2.Id)
      end

      assert_match("The first cancelation amendment to an evergreen order has been processed and the Stripe subscription canceled. Multiple amendments are not supported.", exception.message)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]

      canceled_subscription = Stripe::Subscription.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal("canceled", canceled_subscription.status)
      assert_equal(sf_order.Id, canceled_subscription.metadata['salesforce_order_id'])
    end

    it 'not yet supported user error when try add product to evergreen order through amendment' do
      sf_order = create_evergreen_salesforce_order(
        # need to set these fields explicitly to use translate
        additional_quote_fields: {
          CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(now_time),
          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }
      )

      # create contract for amendment
      sf_contract = create_contract_from_order(sf_order)
      sf_order.refresh

      amendment_end_date = now_time
      amendment_quote = create_quote_data_from_contract_amendment(sf_contract)

      sf_product_id, _ = salesforce_evergreen_product_with_price
      quote_with_product = add_product_to_cpq_quote(amendment_quote['record']['Id'], sf_product_id: sf_product_id)
      calculate_and_save_cpq_quote(quote_with_product)

      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(amendment_end_date)
      amendment_quote["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM

      sf_order_amendment = create_order_from_quote_data(amendment_quote)
      assert_equal(sf_order_amendment.Type, OrderTypeOptions::AMENDMENT.serialize)

      exception = assert_raises(Integrations::Errors::UserError) do
        StripeForce::Translate.perform_inline(@user, sf_order_amendment.Id)
      end

      assert_match("Non-cancelation amendment to evergreen order is not supported.", exception.message)
    end
  end
end
