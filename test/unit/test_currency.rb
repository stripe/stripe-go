# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class CurrencyTest < Critic::UnitTest
    it 'pulls currency from user' do
      user = make_user

      currency = Integrations::Utilities::Currency.base_currency_iso(user)

      assert_equal('usd', currency)

      user.connector_settings['default_currency'] = 'AUD'

      currency = Integrations::Utilities::Currency.base_currency_iso(user)

      assert_equal('aud', currency)
    end
  end
end
