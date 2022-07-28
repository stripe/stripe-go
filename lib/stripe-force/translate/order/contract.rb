# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # "structure" isn't the best name here; meant to indicate that this is an internal structure to represent a contract
  class ContractStructure < T::Struct
    const :initial, Restforce::SObject
    const :amendments, T::Array[Restforce::SObject], default: []

    # this is not determined when the contract is created, should we still include it here? Let's see where this refactor takes u
    # const :termination, T.nilable(Restforce::SObject), default: nil
  end
end
