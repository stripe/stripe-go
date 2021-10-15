# frozen_string_literal: true
# typed: true

class StripeForce::OrderPoller
  def self.perform(user:)
    interactor = self.new(user)
    interactor.perform
  end

  def initialize(user)
    @user = user
  end

  def perform
    sf = @user.sf_client
    one_minute = 1.0 / 24 / 60

    updated_orders = sf.get_updated('Order', DateTime.now - one_minute, DateTime.now)
    # anything else but "ids" in the hash?
    updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

    # TODO updating the line item does NOT update the order

    updated_orders.each do |sf_order_id|
      sf_order = sf.find('Order', sf_order_id)
      StripeForce::Translate.perform(user: @user, sf_object: sf_order)
    end
  end
end
