# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ConfigurationsControllerTest < ApplicationIntegrationTest
  def authentication_headers
    {Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => @user.salesforce_account_id}
  end

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

      post api_translate_path, as: :json, params: { object_type: 'invalid', object_ids: ["123"] }, headers: authentication_headers
      assert_response :bad_request

      post api_translate_path, as: :json, params: { object_type: 'Order', object_ids: 'not an array' }, headers: authentication_headers
      assert_response :bad_request
    end

    it 'queues jobs for translation' do
      number_of_orders = 5

      order_ids = number_of_orders.times.map do
        create_salesforce_id
      end

      SalesforceTranslateRecordJob.expects(:work).times(number_of_orders)

      post api_translate_path, as: :json, params: { object_type: 'Order', object_ids: order_ids }, headers: authentication_headers
      assert_response :success
    end
  end

  describe '#post_install' do
    it 'rejects a invalid request' do
      post api_post_install_path, as: :json, headers: { Api::Controller::SALESFORCE_KEY_HEADER => '123'}
      assert_response :not_found

      post api_post_install_path, as: :json
      assert_response :not_found

      post api_post_install_path, params: 'I am not json', headers: { Api::Controller::SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'), Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id}
      assert_response :not_acceptable
    end

    it 'rejects a request with no organization api key' do
      post api_post_install_path, params: {  }, as: :json, headers: { Api::Controller::SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'), Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id}
      assert_response :bad_request
    end

    it 'creates a new user with a valid organizatiopn API key' do
      api_key = SecureRandom.alphanumeric(16)
      post api_post_install_path, params: { key: api_key }, as: :json, headers: { Api::Controller::SALESFORCE_KEY_HEADER => ENV.fetch('SF_MANAGED_PACKAGE_API_KEY'), Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => sf_account_id}

      assert_equal(1, StripeForce::User.count)
      user = T.must(StripeForce::User.first)

      assert_equal(sf_account_id, user.salesforce_account_id)
      assert_equal(api_key, user.salesforce_organization_key)
    end
  end

  describe 'configuration' do
    before do
      @user = make_user
      @user.field_defaults = {
        "customer" =>{
          "metadata.from_salesforce" => true
        }
      }

      @user.field_mappings = {
        "customer" => {
          "FromSalesforce" => "metadata.from_salesforce"
        }
      }
      @user.save
    end

    describe 'errors' do
      it 'throws a 404 if no user id is passed' do
        get api_configuration_path
        assert_response :not_found
      end

      it 'throws a 404 if invalid user is passed' do
        get api_configuration_path, headers: {Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => create_salesforce_id}
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
      assert_equal(@user.salesforce_account_id, result["salesforce_account_id"])
      assert_equal(@user.field_mappings, result["field_mappings"])
      assert_equal(@user.field_defaults, result["field_defaults"])
    end

    it 'updates user status JSON' do
      updated_field_mapping = {
        "subscription_schedule" => {
          "Email" => "email"
        }
      }

      updated_field_defaults = {
        "customer" => {
          "phone" => "1231231234"
        }
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