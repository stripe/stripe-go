# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceReuse < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
  end

  it 'gracefully fails when a metered price is specified without billing tiers' do
    @user.field_defaults = {"price_order_item" => {"billing_scheme" => "tiered"}}
    @user.save

    sf_order = create_subscription_order(contact_email: "gracefully_fails_2")

    # before this test, this would throw a fatal exception without a nice user error
    exception = assert_raises(Stripe::InvalidRequestError) do
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
    end

    assert_match("tiered the parameter tiers must be set", exception.message)
  end

  it 'translation does not fail if aggregate_usage field is set on a licensed price' do
      @user.field_defaults = {"price" => {"recurring.aggregate_usage" => "last_during_period"}}
      @user.save

      sf_order = create_subscription_order(contact_email: "translation_does_not_fail_1")

      StripeForce::Translate.perform_inline(@user, sf_order.Id)
  end
end
