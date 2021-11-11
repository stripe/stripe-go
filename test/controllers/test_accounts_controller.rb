# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::AccountsControllerTest < ApplicationIntegrationTest
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

  it 'throws a 404 if no user id is passed' do
    get api_accounts_path
    assert_response :not_found
  end

  it 'throws a 404 if invalid user is passed' do
    get api_accounts_path, headers: {Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => create_salesforce_id}
    assert_response :not_found
  end

  it 'returns user status JSON' do
    get api_accounts_path, headers: {Api::Controller::SALESFORCE_ACCOUNT_ID_HEADER => @user.salesforce_account_id}
    assert_response :success

    result = parsed_json
    assert_equal(@user.salesforce_account_id, result["salesforce_account_id"])
    assert_equal(@user.field_mappings, result["field_mappings"])
    assert_equal(@user.field_defaults, result["field_defaults"])
  end
end