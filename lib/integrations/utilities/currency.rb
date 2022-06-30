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

  sig { params(string_float_amount: String, user: StripeForce::User, as_decimal: T::Boolean).returns(T.any(Integer, BigDecimal, String)) }
  def normalize_float_amount_for_stripe(string_float_amount, user, as_decimal: false)
    normalize_float_amount_in_currency_for_stripe(
      Integrations::Utilities::Currency.base_currency_iso(user),
      string_float_amount,
      as_decimal: as_decimal
    )
  end

  sig { params(currency_iso: String, string_float_amount: String, as_decimal: T::Boolean).returns(T.any(Integer, BigDecimal)) }
  def normalize_float_amount_in_currency_for_stripe(currency_iso, string_float_amount, as_decimal: false)
    if zero_decimal_currency?(currency_iso.upcase)
      if as_decimal
        BigDecimal(string_float_amount)
      else
        string_float_amount.to_i
      end
    else
      # https://github.com/stripe/stripe-netsuite/issues/835
      float_amount = BigDecimal(string_float_amount)

      if as_decimal
        (float_amount * 100.0)
      else
        (float_amount * 100.0).to_i
      end
    end
  end

  sig { params(user: StripeForce::User).returns(String) }
  def base_currency_iso(user)
    currency = user.connector_settings['default_currency']

    # TODO determine currency from SF? Maybe cache default currency in Stripe?

    if currency.blank?
      raise Integrations::Errors::ImpossibleState.new("default currency should always be specified")
    end

    # stripe always expects lowercase currency
    currency.downcase
  end
  module_function :base_currency_iso

end
