# frozen_string_literal: true
# typed: true

class StripeForce::OrderPoller
  include SimpleStructuredLogger

  def self.perform(user:)
    interactor = self.new(user)
    interactor.perform
  end

  def initialize(user)
    @user = user
  end

  def perform
    log.info 'initiating order poll', user_id: @user.id

    sf = @user.sf_client

    updated_orders = sf.get_updated('Order', DateTime.now - 2.minutes, DateTime.now)
    # anything else but "ids" in the hash?
    updated_orders = updated_orders["ids"] if updated_orders.is_a?(Hash)

    # TODO updating the line item does NOT update the order

    updated_orders.each do |sf_order_id|
      log.info 'translating order', sf_order_id: sf_order_id

      sf_order = sf.find('Order', sf_order_id)
      StripeForce::Translate.perform(user: @user, sf_object: sf_order)
    end

    log.info 'order poll complete'
  end
end
