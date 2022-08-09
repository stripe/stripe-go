# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module OrderHelpers
    extend T::Sig

    include Kernel
    include StripeForce::Constants
    extend SimpleStructuredLogger

    sig { params(raw_days_until_due: T.any(String, Integer)).returns(Integer) }
    def self.transform_payment_terms_to_days_until_due(raw_days_until_due)
      if raw_days_until_due.is_a?(Integer)
        return raw_days_until_due
      end

      # TODO it is possible for users to customize the options here, we may need to use regex extraction or something at some point
      case raw_days_until_due
      when "Net 15"
        15
      when "Net 30"
        30
      when "Net 45"
        45
      when "Net 60"
        60
      when "Net 90"
        90
      else
        raise StripeForce::Errors::RawUserError.new("unexpected days_until_due option #{raw_days_until_due}")
      end
    end
  end
end
