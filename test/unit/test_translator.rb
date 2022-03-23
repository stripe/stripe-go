# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class TranslatorTest < Critic::UnitTest
    before do
      @user = make_user
      @locker = Integrations::Locker.new(@user)

      @translator = StripeForce::Translate.new(
        @user,
        @locker
      )
    end

    describe 'namespace custom field prefix' do
      it 'prefixes stripe ID field when the QA namespace is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::QA.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: create_stripe_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params["#{SalesforceNamespaceOptions::QA.serialize}_#{GENERIC_STRIPE_ID}"])
        end

        @translator.update_sf_stripe_id(customer, stripe_customer)
      end

      it 'prefixes stripe ID field when the standard is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::PRODUCTION.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: create_stripe_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params["#{SalesforceNamespaceOptions::PRODUCTION.serialize}_#{GENERIC_STRIPE_ID}"])
        end

        @translator.update_sf_stripe_id(customer, stripe_customer)
      end

      it 'prefixes stripe ID field when no namespace is selected' do
        @user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::NONE.serialize

        customer = create_mock_salesforce_customer
        stripe_customer = Stripe::Customer.construct_from(id: create_stripe_id(:cus))

        @translator.sf.expects(:update!).with do |name, params|
          assert_equal(SF_ACCOUNT, name)
          assert_equal(stripe_customer.id, params[GENERIC_STRIPE_ID])
        end

        @translator.update_sf_stripe_id(customer, stripe_customer)

      end
    end
  end
end
