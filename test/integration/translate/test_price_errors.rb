# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceReuse < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'gracefully fails when a metered price is specified without billing tiers' do
    @user.field_defaults = {"price" => {"billing_scheme" => "tiered"}}
    @user.save

    sf_order = create_subscription_order

    # before this test, this would throw a fatal exception without a nice user error
    exception = assert_raises(Stripe::InvalidRequestError) do
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
    end

    assert_match("tiered the parameter tiers must be set", exception.message)
  end
end
