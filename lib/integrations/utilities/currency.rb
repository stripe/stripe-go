# typed: true
# frozen_string_literal: true

module Integrations::Utilities::Currency
  extend T::Sig
  include Kernel

  ZERO_DECIMAL_CURRENCIES = %w{
    BIF
    CLP
    DJF
    GNF
    JPY
    KMF
    KRW
    MGA
    PYG
    RWF
    VND
    VUV
    XAF
    XOF
    XPF
  }.freeze

  sig { params(currency_iso: String).returns(T::Boolean) }
  def zero_decimal_currency?(currency_iso)
    ZERO_DECIMAL_CURRENCIES.include?(currency_iso.upcase)
  end

  sig { params(string_float_amount: String, user: StripeForce::User).returns(Integer) }
  def normalize_float_amount_for_stripe(string_float_amount, user)
    normalize_float_amount_in_currency_for_stripe(base_currency_iso(user), string_float_amount)
  end

  sig { params(currency_iso: String, string_float_amount: String).returns(Integer) }
  def normalize_float_amount_in_currency_for_stripe(currency_iso, string_float_amount)
    if zero_decimal_currency?(currency_iso.upcase)
      string_float_amount.to_i
    else
      # https://github.com/stripe/stripe-netsuite/issues/835
      float_amount = BigDecimal(string_float_amount)
      (float_amount * 100.0).to_i
    end
  end

  sig { params(user: StripeForce::User).returns(String) }
  def base_currency_iso(user)
    # TODO should cache result from Stripe?
    # TODO determine currency from SF?
    'usd'
  end

end
