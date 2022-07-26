# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # TODO rename to contract
  class OrderStructure < T::Struct
    const :initial, Restforce::SObject
    const :amendments, T::Array[Restforce::SObject], default: []
    const :termination, T.nilable(Restforce::SObject), default: nil
  end
end
