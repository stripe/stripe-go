# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class SanitizerTest < Critic::UnitTest
    it 'customers' do
      original_description = "Interface provides a soft place to land. The company is a leading global producer and seller of modular carpet, also known as carpet tile. The tiles and rolls, used in offices and many institutional facilities, are sold under brand names Interface, GlasBac, and FLOR. Interface's other offerings include an antimicrobial chemical, Intersept, which it blends in its carpet as well as licenses for use in air filters, and the TacTiles carpet tile installation system. Core markets are the Americas, Europe, and the Asia/Pacific region; the Americas represent more than 55% of sales."

      customer = Stripe::Customer.construct_from(
        description: original_description,
        shipping: {
          phone: '1 (519) 888-4567 ext. 32110',
          address: {
            country: "US",
          },
        }
      )

      StripeForce::Sanitizer.sanitize(customer)

      refute_equal(original_description, customer.description)
      assert_equal(350, customer.description.length)
      assert_equal(20, customer.shipping.phone.length)
    end

    describe 'prices' do
      it 'truncates long price precision' do
        price = Stripe::Price.construct_from({
          unit_amount_decimal: 0.1234567890123,
        })

        StripeForce::Sanitizer.sanitize(price)
        assert_equal(BigDecimal("0.123456789012"), price.unit_amount_decimal)

        price_2 = Stripe::Price.construct_from({
          "recurring": {
            "usage_type": "licensed",
            "interval": "month",
            "interval_count": "12",
          },
          "metadata": {
            "salesforce_auto_archive": "true",
            "salesforce_order_item_link": "https://octopusdeploy.my.salesforce.com/8027V00000N0r5ZQAR",
            "salesforce_order_item_id": "8027V00000N0r5ZQAR",
          },
          "product": "prod_MON5VwSfaruSHL",
          "currency": "usd",
          "unit_amount_decimal": BigDecimal("273000.000000000000043680000000000000006989"),
        })

        StripeForce::Sanitizer.sanitize(price_2)
        assert_equal(BigDecimal("273000.0"), price_2.unit_amount_decimal)
      end
    end
  end
end
