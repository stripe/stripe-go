# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

# tests for assumptions we are making about the ruby bindings
module Critic::Unit
  class StripeBindingTest < Critic::UnitTest
    it 'deeply copies' do
      obj = Stripe::StripeObject.construct_from(
        fun: "field"
      )

      obj2 = Integrations::Utilities::StripeUtil.deep_copy(obj)
      obj2.fun = "nomore"

      refute_equal(obj2.fun, obj.fun)

      # the reason we are doing this is...
      obj3 = obj.deep_dup
      obj3.fun = "nomore"

      assert_equal(obj3.fun, obj.fun)
    end

    it 'gracefully handles hash accessor syntax' do
      price = Stripe::Price.new

      # default functionality that we intentionally avoid in some places
      assert_raises(NoMethodError) do
        price.recurring
      end

      assert_nil(price[:recurring])

      price[:recurring] ||= {}
      assert_equal(Stripe::StripeObject, price.recurring.class)

      price[:recurring] ||= 10
      assert_equal(Stripe::StripeObject, price.recurring.class)

      assert_raises(NoMethodError) do
        price.recurring.interval_count
      end

      assert_nil(price.recurring[:interval_count])

      price.recurring.interval_count = 10
      assert_equal(10, price.recurring.interval_count)

      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
        price,
        :recurring
      )

      assert_nil(price[:recurring])
    end
  end
end
