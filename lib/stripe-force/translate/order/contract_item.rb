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
    const :from_renewal_order, T::Boolean

    # not const so we can mutate it (`reduce_quantity`, etc)
    prop :quantity, Integer

    # the delta in quantity for order amendments
    prop :reduced_by, Integer

    sig { params(sf_order_line: Restforce::SObject, stripe_params: Hash, from_renewal_order: T::Boolean).returns(ContractItemStructure) }
    def self.from_order_line_and_params(sf_order_line, stripe_params, from_renewal_order: false)
      self.new(
        stripe_params: stripe_params,
        order_line: sf_order_line,
        order_line_id: sf_order_line.Id,
        from_renewal_order: from_renewal_order,
        revised_order_line_id: sf_order_line[SF_ORDER_ITEM_REVISED_ORDER_PRODUCT],
        quantity: stripe_params[:quantity],
        reduced_by: 0
      )
    end

    sig { params(user: T.nilable(StripeForce::User)).returns(Stripe::Price) }
    def price(user=nil)
      if @price.nil? && user.nil?
        raise Integrations::Errors::ImpossibleInternalError.new("query should never return an empty result")
      end

      # sorbet doesn't like instance memoized instance variables
      @price ||= Stripe::Price.retrieve(self.stripe_params[:price], T.unsafe(user).stripe_credentials)
      @price
    end

    def reduce_quantity
      if self.quantity < 0
        raise "termination lines should never be reduced"
      end

      self.reduced_by += 1
      self.quantity -= 1
    end

    sig { returns(T::Boolean) }
    def termination?
      self.quantity < 0
    end

    sig { returns(T::Boolean) }
    def fully_terminated?
      self.quantity.zero?
    end

    # "new" meaning not-revised from previous order
    sig { returns(T::Boolean) }
    def new_order_line?
      self.revised_order_line_id.blank? || self.from_renewal_order
    end

    sig { params(sf_order: Restforce::SObject).returns(T::Boolean) }
    def from_order?(sf_order)
      self.order_line['OrderId'] == sf_order.Id
    end

    sig { returns(T::Boolean) }
    def is_mdq_segment?
      self.order_line[CPQ_ORDER_ITEM_SEGMENT_KEY].present?
    end

    # this starts at 1 index
    sig { returns(T.nilable(Integer)) }
    def mdq_segment_index
      self.order_line[CPQ_ORDER_ITEM_SEGMENT_INDEX]
    end

    sig { returns(T.nilable(String)) }
    def mdq_dimension_type
      self.order_line[CPQ_ORDER_ITEM_PRICE_DIMENSION_TYPE]
    end
  end
end
