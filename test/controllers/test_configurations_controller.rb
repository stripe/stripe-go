# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ConfigurationsControllerTest < ApplicationIntegrationTest
  describe '#translate' do
    before do
      @user = make_user(save: true)
    end

    it 'validates the input' do
      post api_translate_path, as: :json
      assert_response :not_found

      post api_translate_path, headers: authentication_headers
      assert_response :not_acceptable

      post api_translate_path, as: :json, headers: authentication_headers
      assert_response :bad_request

      post api_translate_path, as: :json, params: {object_type: 'invalid', object_ids: ["123"]}, headers: authentication_headers
      assert_response :bad_request

      post api_translate_path, as: :json, params: {object_type: 'Order', object_ids: 'not an array'}, headers: authentication_headers
      assert_response :bad_request
    end

    it 'queues jobs for translation' do
      number_of_orders = 5

      order_ids = number_of_orders.times.map do
        create_salesforce_id
      end

      SalesforceTranslateRecordJob.expects(:work).times(number_of_orders)

      post api_translate_path, as: :json, params: {object_type: 'Order', object_ids: order_ids}, headers: authentication_headers
      assert_response :success
    end

    it 'accepts account reference' do
      SalesforceTranslateRecordJob.expects(:work)

      post api_translate_path, as: :json, params: {object_type: SF_ACCOUNT, object_ids: [create_salesforce_id]}, headers: authentication_headers
      assert_response :success
    end

    it 'accepts product reference' do
      SalesforceTranslateRecordJob.expects(:work)

      post api_translate_path, as: :json, params: {object_type: SF_PRODUCT, object_ids: [create_salesforce_id]}, headers: authentication_headers
      assert_response :success
    end
  end

  describe '#post_install' do
    it 'rejects a invalid request' do
      post api_post_install_path, as: :json, headers: {SALESFORCE_KEY_HEADER => '123'}
      assert_response :not_found

      post api_post_install_path, as: :json
      assert_response :not_found

      post api_post_install_path, params: 'I am not json', headers: {SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'), SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id}
      assert_response :not_acceptable
    end

    it 'rejects a request with no organization api key' do
      post api_post_install_path, params: {}, as: :json, headers: {SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'), SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id}
      assert_response :bad_request
    end

    it 'creates a new user with a valid organization API key' do
      api_key = SecureRandom.alphanumeric(16)
      post api_post_install_path, params: {key: api_key}, as: :json, headers: {
        # not using `authentication_headers` since the user is not created
        SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'),
        SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id,
        SALESFORCE_INSTANCE_TYPE_HEADER => SFInstanceTypes::SANDBOX.serialize,
        SALESFORCE_PACKAGE_NAMESPACE_HEADER => "",
      }

      assert_equal(1, StripeForce::User.count)
      user = T.must(StripeForce::User.first)

      assert_equal(sf_account_id, user.salesforce_account_id)
      assert_equal(api_key, user.salesforce_organization_key)
      assert_equal(SalesforceNamespaceOptions::PRODUCTION.serialize, user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE])
      assert_equal(SFInstanceTypes::SANDBOX.serialize, user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE])
    end

    it 'persists the salesforce namespace when in QA' do
      api_key = SecureRandom.alphanumeric(16)
      post api_post_install_path, params: {key: api_key}, as: :json, headers: {
        # not using `authentication_headers` since the user is not created
        SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'),
        SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id,
        SALESFORCE_PACKAGE_NAMESPACE_HEADER => SalesforceNamespaceOptions::QA.serialize,
      }

      assert_equal(1, StripeForce::User.count)
      user = T.must(StripeForce::User.first)
      assert_equal(SalesforceNamespaceOptions::QA.serialize, user.connector_settings['salesforce_namespace'])
    end
  end

  describe 'configuration' do
    before do
      @user = make_user
      @user.field_defaults = {
        "customer" => {
          "metadata.from_salesforce" => true,
        },
      }

      @user.field_mappings = {
        "customer" => {
          "FromSalesforce" => "metadata.from_salesforce",
        },
      }
      @user.save
    end

    describe 'errors' do
      it 'throws a 404 if no user id is passed' do
        get api_configuration_path
        assert_response :not_found
      end

      it 'throws a 404 if invalid user is passed' do
        get api_configuration_path, headers: authentication_headers.merge({SALESFORCE_ACCOUNT_ID_HEADER => create_salesforce_id})
        assert_response :not_found
      end

      it 'throws a not accepted error if JSON is not passed' do
        put api_configuration_path, params: "i am not json", headers: authentication_headers
        assert_response :not_acceptable
      end
    end

    it 'returns user status JSON' do
      get api_configuration_path, headers: authentication_headers

      assert_response :success

      result = parsed_json

      refute_nil(result['default_mappings'])

      assert_equal(@user.salesforce_account_id, result["salesforce_account_id"])
      assert_equal(@user.field_mappings, result["field_mappings"])
      assert_equal(@user.field_defaults, result["field_defaults"])

      assert_equal(95, result['settings']['api_percentage_limit'])
      assert_equal(10_000, result['settings']['sync_record_retention'])
      assert_equal('USD', result['settings']['default_currency'])
      assert(result['settings']['sync_start_date'] > Time.now.to_i - 1)
    end

    it 'updates settings' do
      future_time = Time.now.to_i + 3600

      put api_configuration_path, params: {
        settings: {
          api_percentage_limit: 90,
          sync_start_date: future_time,
          sync_record_retention: 1_000,
          default_currency: 'EUR',
        },
      }, as: :json, headers: authentication_headers

      assert_response :success

      result = parsed_json

      @user = T.must(StripeForce::User[@user.id])

      assert_equal(90, result['settings']['api_percentage_limit'])
      assert_equal(1_000, result['settings']['sync_record_retention'])
      assert_equal('EUR', result['settings']['default_currency'])
      assert_equal(result['settings']['sync_start_date'], future_time)
    end

    it 'does not remove settings which are not present in the incoming hash' do
      @user.connector_settings['salesforce_namespace'] = SalesforceNamespaceOptions::QA.serialize
      @user.save

      put api_configuration_path, params: {
        settings: {
          default_currency: 'EUR',
        },
      }, as: :json, headers: authentication_headers

      assert_response :success

      result = parsed_json

      @user = T.must(StripeForce::User[@user.id])

      assert_equal('EUR', @user.connector_settings['default_currency'])
      assert_equal(SalesforceNamespaceOptions::QA.serialize, @user.connector_settings['salesforce_namespace'])
    end

    it 'updates mappings and defaults without settings' do
      updated_field_mapping = {
        "subscription_schedule" => {
          "Email" => "email",
        },
      }

      updated_field_defaults = {
        "customer" => {
          "phone" => "1231231234",
        },
      }

      put api_configuration_path, params: {
        field_mappings: updated_field_mapping,
        field_defaults: updated_field_defaults,
      }, as: :json, headers: authentication_headers

      assert_response :success

      result = parsed_json

      @user = T.must(StripeForce::User[@user.id])

      assert_equal(@user.field_mappings, updated_field_mapping)
      assert_equal(@user.field_defaults, updated_field_defaults)

      assert_equal(@user.field_mappings, result["field_mappings"])
      assert_equal(@user.field_defaults, result["field_defaults"])
    end
  end
end
