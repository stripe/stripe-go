# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class SalesforceContextTest < Critic::UnitTest
    before do
      @translator = make_translator
    end

    it 'allows the primary and secondary context to be set' do
      order = create_mock_salesforce_order
      account = create_mock_salesforce_customer

      @translator.expects(:create_user_failure).with do |args|
        assert_equal(account, args[:salesforce_object])
        assert_equal("some error", args[:message])
      end

      # also tests the raw => standard user error conversion
      exception = assert_raises(StripeForce::Errors::UserError) do
        @translator.catch_errors_with_salesforce_context(primary: order) do
          @translator.catch_errors_with_salesforce_context(secondary: account) do
            raise StripeForce::Errors::RawUserError.new("some error")
          end
        end
      end

      assert_equal("some error", exception.message)
    end

    describe 'allows the secondary object to be set more than once' do
      it 'it uses the last context specified' do
        order = create_mock_salesforce_order
        account = create_mock_salesforce_customer
        account2 = create_mock_salesforce_customer

        @translator.expects(:create_user_failure).with do |args|
          assert_equal(account2, args[:salesforce_object])
          assert_equal("some error", args[:message])
        end

        assert_raises(StripeForce::Errors::UserError) do
          @translator.catch_errors_with_salesforce_context(primary: order) do
            @translator.catch_errors_with_salesforce_context(secondary: account) do
              @translator.catch_errors_with_salesforce_context(secondary: account2) do
                raise StripeForce::Errors::RawUserError.new("some error")
              end
            end
          end
        end
      end

      it 'properly restores the previous context' do
        order = create_mock_salesforce_order
        account = create_mock_salesforce_customer
        account2 = create_mock_salesforce_customer

        @translator.expects(:create_user_failure).with do |args|
          assert_equal(account, args[:salesforce_object])
          assert_equal("some error", args[:message])
        end

        assert_raises(StripeForce::Errors::UserError) do
          @translator.catch_errors_with_salesforce_context(primary: order) do
            @translator.catch_errors_with_salesforce_context(secondary: account) do
              @translator.catch_errors_with_salesforce_context(secondary: account2) do
                # noop
              end

              raise StripeForce::Errors::RawUserError.new("some error")
            end
          end
        end
      end
    end

    it 'does not allow primary context to be set twice' do
      order = create_mock_salesforce_order
      account = create_mock_salesforce_customer

      exception = assert_raises(Integrations::Errors::ImpossibleInternalError) do
        @translator.catch_errors_with_salesforce_context(primary: order) do
          @translator.catch_errors_with_salesforce_context(primary: account) do
            # noop
          end
        end
      end

      assert_equal("origin object already set, exiting", exception.message)
    end

    it 'handles errors without secondary context' do
      order = create_mock_salesforce_order

      @translator.expects(:create_user_failure).with do |args|
        assert_equal(order, args[:salesforce_object])
        assert_equal("some error", args[:message])
      end

      exception = assert_raises(StripeForce::Errors::UserError) do
        @translator.catch_errors_with_salesforce_context(primary: order) do
          raise StripeForce::Errors::RawUserError.new("some error")
        end
      end
    end

  end
end
