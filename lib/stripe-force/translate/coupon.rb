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
end
