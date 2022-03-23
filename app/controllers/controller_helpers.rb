# frozen_string_literal: true
# typed: strict

module ControllerHelpers
  extend T::Sig
  include Kernel

  # https://github.com/stripe/stripe-salesforce/pull/171
  sig { params(raw_namespace: T.nilable(String)).returns(String) }
  protected def subdomain_namespace_from_param(raw_namespace)
    case raw_namespace
    when nil, ""
      "stripeConnector"
    when "c", "stripeConnector", "stripeConnectQA"
      raw_namespace
    else
      raise "unexpected namespace #{raw_namespace}"
    end
  end
end
