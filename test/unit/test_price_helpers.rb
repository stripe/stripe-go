# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class PriceHelpersTest < Critic::UnitTest
    describe '#pricing_tiers_equal?' do
      # https://jira.corp.stripe.com/browse/PLATINT-1817
      it 'handles different ordering of tiers' do
        price_1 = Stripe::Price.construct_from({
          id: "price_1LalAnCM3YxTUEmtSIWT580H",
          object: "price",
          active: false,
          billing_scheme: "tiered",
          created: 1661453581,
          currency: "usd",
          custom_unit_amount: nil,
          livemode: false,
          lookup_key: nil,
          metadata: {},
          nickname: nil,
          product: "prod_LfIYek0p06aEF8",
          recurring: {aggregate_usage: nil,
                      interval: "month",
                      interval_count: 1,
                      trial_period_days: nil,
                      usage_type: "metered",},
          tax_behavior: "unspecified",
          tiers: [{flat_amount: 0,
                   flat_amount_decimal: "0",
                   unit_amount: nil,
                   unit_amount_decimal: nil,
                   up_to: 11,},
                  {flat_amount: nil,
                   flat_amount_decimal: nil,
                   unit_amount: 2537,
                   unit_amount_decimal: "2537",
                   up_to: nil,},],
          tiers_mode: "graduated",
          transform_quantity: nil,
          type: "recurring",
          unit_amount: nil,
          unit_amount_decimal: nil,
        })

        price_2 = Stripe::Price.construct_from({
          billing_scheme: "tiered",
          currency: "usd",
          product: "prod_LfIYek0p06aEF8",
          recurring: {aggregate_usage: nil,
                      interval: "month",
                      interval_count: 1,
                      trial_period_days: nil,
                      usage_type: "metered",},
          tax_behavior: "unspecified",
          tiers: [
            # NOTE tier ordering is flipped!
            {flat_amount: nil,
             flat_amount_decimal: nil,
             unit_amount: 2537,
             unit_amount_decimal: "2537",
             up_to: nil,},

            {flat_amount: 0,
             flat_amount_decimal: "0",
             unit_amount: nil,
             unit_amount_decimal: nil,
             up_to: 11,},
          ],
          tiers_mode: "graduated",
          transform_quantity: nil,
          type: "recurring",
          unit_amount: nil,
          unit_amount_decimal: nil,
        })

        assert(StripeForce::Translate::PriceHelpers.pricing_tiers_equal?(price_1, price_2))
      end
    end

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
