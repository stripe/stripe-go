# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class PriceHelpersTest < Critic::UnitTest
    describe '#using_custom_order_line_price_field' do
      it 'detects if a custom pricing field is not used' do
        user = make_user
        refute(StripeForce::Translate::PriceHelpers.using_custom_order_line_price_field?(user))
      end

      it 'detects if a custom pricing is used' do
        user = make_user
        user.field_mappings = {
          'price_order_item' => {
            'unit_amount_decimal' => 'SomethingCustomField',
          },
        }

        assert(StripeForce::Translate::PriceHelpers.using_custom_order_line_price_field?(user))
      end
    end
  end
end
