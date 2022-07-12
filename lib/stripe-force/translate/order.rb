# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  class OrderStructure < T::Struct
    const :initial, Restforce::SObject
    const :amendments, T::Array[Restforce::SObject], default: []
    const :termination, T.nilable(Restforce::SObject), default: nil
  end

  class PhaseItemStructure < T::Struct
    extend T::Sig

    const :stripe_params, Hash

    # will be nil if this is a net-new line item
    const :original_order_line_id, T.nilable(String)

    # when we are creating a struct from a sub schedule phase item, we don't have the SF object handyu
    const :order_line, T.nilable(Restforce::SObject)

    # even in a metered billing scenario, which have no quantity in the subscriptio phase item,
    # this will have a value indicating that the subscripion item should be added to the order.
    prop :quantity, Integer

    # used when generating a phase item struct from the last active phases structure
    sig { params(stripe_params_hash: Hash).returns(PhaseItemStructure) }
    def self.new_from_created_phase_item(stripe_params_hash)
      quantity = if stripe_params_hash[:quantity].nil?
        log.info 'no quantity field, assuming metered billing at quantity 1'
        1
      else
        stripe_params_hash[:quantity]
      end

      # TODO I don't like relying on the metadata here; maybe we could regenerate the line items for the first order and use that representation instead?
      #      or maybe in the future we could use an internal sync record to pull references? Either way, this should change.
      original_order_id = stripe_params_hash[:metadata][:salesforce_order_item_id]

      self.new(
        stripe_params: stripe_params_hash,
        quantity: quantity,
        original_order_line_id: original_order_id
      )
    end

    sig { params(old_phase_item: PhaseItemStructure).void }
    def append_previous_phase_item(old_phase_item)
      self.quantity += old_phase_item.quantity

      # this could be possible if a user has multiple quantity of a metered billing item, which should never happen
      # if this happens, we'll need to report a user error about this.
      if self.quantity < 0
        raise Integrations::Errors::ImpossibleState.new("quantity should never be zero")
      end

      if stripe_params[:quantity].nil? && old_phase_item.stripe_params[:quantity].nil?
        log.debug 'item is metered, skipping quantity adjustment'
      else
        stripe_params[:quantity] += old_phase_item.stripe_params[:quantity]
      end
    end

    sig { returns(T::Boolean) }
    def is_terminated?
      quantity.zero?
    end
  end
end
