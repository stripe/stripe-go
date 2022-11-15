# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_coupon(sf_coupon)
    locker.lock_salesforce_record(sf_coupon)

    log.info 'translating coupon', salesforce_object: sf_coupon

    coupon = create_stripe_object(Stripe::Coupon, sf_coupon)

    update_sf_stripe_id(sf_coupon, coupon)

    coupon
  end

  def self.get_salesforce_stripe_coupons_associated_to_order_line(sf_client:, sf_order_line_id:)
    order_item_associations = sf_client.query("Select Id from #{SF_STRIPE_COUPON_ORDER_ITEM_ASSOCIATION} where Order_Item__c = '#{sf_order_line_id}'")

    if !order_item_associations || order_item_associations.size == 0
      log.info "no stripe coupon order line associations related to this order line", salesforce_object: sf_order_line_id
      return
    end

    # there could be multiple coupons associated with a single order line
    coupons = order_item_associations.map do |order_item_association|
        association = sf_client.find(SF_STRIPE_COUPON_ORDER_ITEM_ASSOCIATION, order_item_association.Id)
        coupon = sf_client.query("Select Id from #{SF_STRIPE_COUPON_SERIALIZED} where Id = '#{association.Stripe_Coupon__c}'")

        # return the coupon object
        sf_client.find(SF_STRIPE_COUPON_SERIALIZED, coupon.first.Id)
    end

    coupons
  end
end
