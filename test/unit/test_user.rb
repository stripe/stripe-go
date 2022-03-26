# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class UserTest < Critic::UnitTest
    before do
      @user = make_user
    end

    describe 'host selection' do
      it 'uses a test host when not in production' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = SFInstanceTypes::SANDBOX.serialize

        assert_equal('test.salesforce.com', @user.sf_client.options[:host])
      end

      it 'uses a standard host when in production' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = SFInstanceTypes::PRODUCTION.serialize

        assert_equal('login.salesforce.com', @user.sf_client.options[:host])
      end
    end
  end
end
