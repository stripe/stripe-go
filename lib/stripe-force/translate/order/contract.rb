# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # "structure" isn't the best name here; meant to indicate that this is an internal structure to represent a contract
  class ContractStructure < T::Struct
    const :initial, Restforce::SObject
    const :amendments, T::Array[Restforce::SObject], default: []
  end
end
