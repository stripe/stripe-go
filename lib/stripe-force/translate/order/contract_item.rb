# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  # `Structure` meant to indicate that this is an internal representation of the contract item in SF
  class ContractItemStructure < T::Struct
    extend T::Sig

    include Integrations::Log
    include StripeForce::Constants

    const :stripe_params, Hash

    # will be nil if this is a net-new line item
    const :revised_order_line_id, T.nilable(String)

    # when we are creating a struct from a sub schedule phase item, we don't have the SF object handy
    # we should not rely on this
    const :order_line, Restforce::SObject
    const :order_line_id, String

    # not const so we can mutate it (`reduce_quantity`, etc)
    prop :quantity, Integer

    # TODO maybe consider including a boolean for this instead?
    # is_recurring: recurring_item?(sf_order_item),

    sig { params(sf_order_line: Restforce::SObject, stripe_params: Hash).returns(ContractItemStructure) }
    def self.from_order_line_and_params(sf_order_line, stripe_params)
      self.new(
        stripe_params: stripe_params,

        order_line: sf_order_line,
        order_line_id: sf_order_line.Id,

        revised_order_line_id: sf_order_line[SF_ORDER_ITEM_REVISED_ORDER_PRODUCT],
        quantity: stripe_params[:quantity]
      )
    end

    def reduce_quantity
      if self.quantity < 0
        raise "termination lines should never be reduced"
      end

      self.quantity -= 1
    end

    sig { returns(T::Boolean) }
    def termination?
      self.quantity < 0
    end

    sig { returns(T::Boolean) }
    def fully_terminated?
      quantity.zero?
    end

    # "new" meaning not-revised
    sig { returns(T::Boolean) }
    def new_order_line?
      revised_order_line_id.blank?
    end
  end
end
