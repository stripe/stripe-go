# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ConfigurationControllerTest < ApplicationIntegrationTest
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

  def authentication_headers
    {Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => @user.salesforce_account_id}
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

    @user = StripeForce::User[@user.id]

    assert_equal(@user.field_mappings, updated_field_mapping)
    assert_equal(@user.field_defaults, updated_field_defaults)

    assert_equal(@user.field_mappings, result["field_mappings"])
    assert_equal(@user.field_defaults, result["field_defaults"])
  end
end