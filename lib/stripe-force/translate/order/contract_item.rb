# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # `Structure` meant to indicate that this is an internal representation of the contract item in SF
  class ContractItemStructure < T::Struct
    extend T::Sig
    include Integrations::Log

    const :stripe_params, Hash

    # will be nil if this is a net-new line item
    const :revised_order_line_id, T.nilable(String)

    # when we are creating a struct from a sub schedule phase item, we don't have the SF object handy
    # we should not rely on this
    const :order_line, T.nilable(Restforce::SObject)

    const :order_line_id, String

    # even in a metered billing scenario, which have no quantity in the subscriptio phase item,
    # this will have a value indicating that the subscripion item should be added to the order.
    prop :quantity, Integer

    sig { params(old_phase_item: ContractItemStructure).void }
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
