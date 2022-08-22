# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class TranslatorTest < Critic::UnitTest
    before do
      @user = make_user
      @translator = make_translator(user: @user)
    end

    describe 'required and optional mapping' do
      before do
        @user.expects(:required_mappings).returns({
          "subscription_schedule" => {
            "start_date" => CPQ_QUOTE_SUBSCRIPTION_START_DATE,
            "subscription_iterations" => CPQ_QUOTE_SUBSCRIPTION_TERM,
          },
        })
      end

      it 'throws an error if required fields are missing' do
        # omit subscription term to intentionally create an error
        sf_order = create_mock_salesforce_order
        sf_order[CPQ_QUOTE_SUBSCRIPTION_START_DATE] = DateTime.now.to_s

        # TODO should redefine `assert_raises` to accept a dynamic class param
        exception = assert_raises(Integrations::Errors::MissingRequiredFields) do
          @translator.extract_salesforce_params!(sf_order, Stripe::SubscriptionSchedule)
        end

        exception = T.cast(exception, Integrations::Errors::MissingRequiredFields)

        assert_equal([CPQ_QUOTE_SUBSCRIPTION_TERM], exception.missing_salesforce_fields)
      end

      it 'extracts optional params if they exist' do
        @user.expects(:default_mappings).returns({
          "subscription_schedule" => {
            "extra" => 'Description',
          },
        })

        # omit subscription term to intentionally create an error
        sf_order = create_mock_salesforce_order
        sf_order[CPQ_QUOTE_SUBSCRIPTION_START_DATE] = DateTime.now.to_s
        sf_order[CPQ_QUOTE_SUBSCRIPTION_TERM] = 12
        sf_order['Description'] = 'a description'

        results = @translator.extract_salesforce_params!(sf_order, Stripe::SubscriptionSchedule)

        assert_equal({"start_date" => sf_order[CPQ_QUOTE_SUBSCRIPTION_START_DATE], "subscription_iterations" => 12, "extra" => "a description"}, results)
      end

      it 'gracefully handles missing optional fields' do
        @user.expects(:default_mappings).returns({
          "subscription_schedule" => {
            "extra" => 'Description',
          },
        })

        # omit subscription term to intentionally create an error
        sf_order = create_mock_salesforce_order
        sf_order[CPQ_QUOTE_SUBSCRIPTION_START_DATE] = DateTime.now.to_s
        sf_order[CPQ_QUOTE_SUBSCRIPTION_TERM] = 12

        results = @translator.extract_salesforce_params!(sf_order, Stripe::SubscriptionSchedule)

        assert_equal({"start_date" => sf_order[CPQ_QUOTE_SUBSCRIPTION_START_DATE], "subscription_iterations" => 12}, results)
      end
    end

    describe 'namespace custom field prefix' do
      it 'prefixes stripe ID field when the QA namespace is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::QA.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params["#{SalesforceNamespaceOptions::QA.serialize}__#{GENERIC_STRIPE_ID}"])
        end

        @translator.expects :create_user_success

        @translator.update_sf_stripe_id(customer, stripe_customer)
      end

      it 'prefixes stripe ID field when the standard is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::PRODUCTION.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params["#{SalesforceNamespaceOptions::PRODUCTION.serialize}__#{GENERIC_STRIPE_ID}"])
        end

        @translator.expects :create_user_success

        @translator.update_sf_stripe_id(customer, stripe_customer)
      end

      it 'prefixes stripe ID field when no namespace is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::NONE.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params[GENERIC_STRIPE_ID])
        end

        @translator.expects :create_user_success

        @translator.update_sf_stripe_id(customer, stripe_customer)
      end
    end

    describe 'generate_idempotency_key_with_credentials' do
      it 'generates a key' do
        sf_order = create_mock_salesforce_order

        creds = @translator.generate_idempotency_key_with_credentials(@user, sf_order)

        assert_equal(sf_order.Id, creds.delete(:idempotency_key))
        assert_equal(@user.stripe_credentials, creds)
      end

      it 'generates a key with an action' do
        sf_order = create_mock_salesforce_order

        creds = @translator.generate_idempotency_key_with_credentials(
          @user,
          sf_order,
          :finalize_invoice
        )

        assert_equal("#{sf_order.Id}-finalize_invoice", creds.delete(:idempotency_key))
        assert_equal(@user.stripe_credentials, creds)
      end
    end
  end
end
