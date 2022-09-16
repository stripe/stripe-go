# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class OrderHelpersTest < Critic::UnitTest
    describe '#transform_payment_terms_to_days_until_due' do
      it 'allows for an integer input' do
        assert_equal(10, StripeForce::Translate::OrderHelpers.transform_payment_terms_to_days_until_due(10))
      end

      it 'allows for a string integer' do
        assert_equal(10, StripeForce::Translate::OrderHelpers.transform_payment_terms_to_days_until_due(" 10 "))
      end

      it 'throws an error with a float string' do
        assert_raises(StripeForce::Errors::RawUserError) do
          StripeForce::Translate::OrderHelpers.transform_payment_terms_to_days_until_due("10.123")
        end
      end

      it 'allows CPQ enum values' do
        assert_equal(90, StripeForce::Translate::OrderHelpers.transform_payment_terms_to_days_until_due("Net 90"))
      end
    end
  end
end
