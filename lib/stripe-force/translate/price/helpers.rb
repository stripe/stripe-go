# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  module PriceHelpers
    include Kernel
    extend T::Sig

    sig { params(raw_consumption_schedule_type: String).returns(String) }
    def self.transform_salesforce_consumption_schedule_type_to_tier_mode(raw_consumption_schedule_type)
      case raw_consumption_schedule_type
      when "Range"
        'volume'
      when "Slab"
        'graduated'
      else
        # should never happen
        raise "unexpected consumption schedule type #{raw_consumption_schedule_type}"
      end
    end
  end
end
