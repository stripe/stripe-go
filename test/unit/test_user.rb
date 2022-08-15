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

    describe '#stripe_credentials' do
      # minitest doesn't have around hooks :(
      before(:each) do
        # doing this so that it will not be equivalent to STRIPE_API_KEY, which
        # is needed to keep these tests meaningful
        # there's another test where STRIPE_TEST_API_KEY needs to be a real key
        StripeForce::User.send(:remove_const, :SF_STRIPE_TESTMODE_API_KEY)
        StripeForce::User.const_set(:SF_STRIPE_TESTMODE_API_KEY, 'thisisafakekeyfortests')
      end

      after(:each) do
        StripeForce::User.send(:remove_const, :SF_STRIPE_TESTMODE_API_KEY)
        StripeForce::User.const_set(:SF_STRIPE_TESTMODE_API_KEY, ENV.fetch('STRIPE_TEST_API_KEY'))
      end

      def make_user_for_client_credentials_tests(overrides={})
        user = make_user(overrides)
        user
      end

      it "returns platform keys if in the flag" do
        user = make_user(livemode: true)

        credentials = user.stripe_credentials

        assert(credentials.key?(:api_key))
        assert(credentials.key?(:stripe_account))
        assert_equal(
          user.stripe_account_id,
          credentials.fetch(:stripe_account),
        )
      end

      it "when enabled returns platform livemode keys for livemode" do
        user = make_user(livemode: true)

        credentials = user.stripe_credentials

        refute_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY
        )
        assert_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          credentials.fetch(:api_key),
        )
      end

      it "returns platform testmode keys for testmode" do
        user = make_user(livemode: false)

        credentials = user.stripe_credentials

        refute_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY
        )
        assert_equal(
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY,
          credentials.fetch(:api_key),
        )
      end

      it 'can force livemode keys' do
        user = make_user(livemode: false)

        credentials = user.stripe_credentials(forced_livemode: true)

        refute_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY
        )
        assert_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          credentials.fetch(:api_key),
        )
      end

      it 'can force testmode keys' do
        user = make_user(livemode: true)

        credentials = user.stripe_credentials(forced_livemode: false)

        refute_equal(
          StripeForce::User::SF_STRIPE_LIVEMODE_API_KEY,
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY
        )
        assert_equal(
          StripeForce::User::SF_STRIPE_TESTMODE_API_KEY,
          credentials.fetch(:api_key),
        )
      end

      it 'does not specify an api version by default' do
        user = make_user

        credentials = user.stripe_credentials

        assert_nil(credentials[:stripe_version])
      end
    end
  end
end
