# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  class OrderStructure < T::Struct
    const :initial, Restforce::SObject
    const :amendments, T::Array[Restforce::SObject], default: []
    const :termination, T.nilable(Restforce::SObject), default: nil
  end

  class PhaseItemStructure < T::Struct
    const :original_order_line_id, T.nilable(String)
    const :order_line, Restforce::SObject
    const :stripe_params, Hash
  end
end
