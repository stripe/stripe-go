# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class SanitizerTest < Critic::UnitTest
    before do
      @user = make_user
      @sanitizer = StripeForce::Sanitizer.new(@user)
    end

    it 'customers' do
      original_description = "Interface provides a soft place to land. The company is a leading global producer and seller of modular carpet, also known as carpet tile. The tiles and rolls, used in offices and many institutional facilities, are sold under brand names Interface, GlasBac, and FLOR. Interface's other offerings include an antimicrobial chemical, Intersept, which it blends in its carpet as well as licenses for use in air filters, and the TacTiles carpet tile installation system. Core markets are the Americas, Europe, and the Asia/Pacific region; the Americas represent more than 55% of sales."

      customer = Stripe::Customer.construct_from(
        description: original_description
      )

      @sanitizer.sanitize(customer)

      refute_equal(original_description, customer.description)
      assert_equal(350, customer.description.length)
    end

    it 'prices' do
      price = Stripe::Price.construct_from({
        unit_amount_decimal: 0.1234567890123,
      })

      @sanitizer.sanitize(price)

      assert_equal(BigDecimal("0.123456789012"), price.unit_amount_decimal)
    end
  end
end
