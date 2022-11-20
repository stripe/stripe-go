# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_coupon(sf_coupon)
    locker.lock_salesforce_record(sf_coupon)

    log.info 'translating coupon', salesforce_object: sf_coupon

    catch_errors_with_salesforce_context(secondary: sf_coupon) do
      create_coupon_from_sf_coupon(sf_coupon)
    end
  end

  def create_coupon_from_sf_coupon(sf_coupon)
    # check if the original sf coupon has been translated prior and exists in Stripe
    # this is the coupon that the serialized coupon on the order/order item was copied from

    # TODO hook up to cache service
    original_sf_coupon = sf.find(SF_STRIPE_COUPON, sf_coupon.Original_Stripe_Coupon_Beta_Id__c)
    existing_stripe_coupon = retrieve_from_stripe(Stripe::Coupon, original_sf_coupon)
    if existing_stripe_coupon
      existing_stripe_coupon = T.cast(existing_stripe_coupon, Stripe::Coupon)
      generated_stripe_coupon = construct_stripe_object(stripe_class: Stripe::Coupon, salesforce_object: original_sf_coupon)

      # this should never happen unless the coupon data in Salesforce is mutated
      # if so, we want to create a new Stripe object and write back the new id
      if coupons_are_equal?(existing_stripe_coupon: existing_stripe_coupon, generated_stripe_coupon: generated_stripe_coupon)
        log.info 'using existing stripe coupon', existing_stripe_coupon_id: existing_stripe_coupon.id
        return existing_stripe_coupon
      end
    end

    # create a new stripe coupon
    stripe_coupon = create_stripe_object(Stripe::Coupon, original_sf_coupon)

    # update the sf coupon id on the original coupon
    update_sf_stripe_id(original_sf_coupon, stripe_coupon)

    stripe_coupon
  end

  def coupons_are_equal?(existing_stripe_coupon:, generated_stripe_coupon:)
      fields_to_check_if_both_are_set = %i{
        name
        amount_off
        percent_off
        duration
        duration_in_months
        currency
      }

      simple_field_check_passed = fields_to_check_if_both_are_set.all? do |field_sym|
        existing_stripe_coupon[field_sym].blank? == generated_stripe_coupon[field_sym].blank? && existing_stripe_coupon[field_sym] == generated_stripe_coupon[field_sym]
      end

      if !simple_field_check_passed
        log.info 'coupons are not equal, simple field comparison failed', diff: HashDiff::Comparison.new(existing_stripe_coupon.to_hash, generated_stripe_coupon.to_hash).diff
      end
      simple_field_check_passed
  end

  # https://site-admin.stripe.com/docs/api/subscription_schedules/object#subscription_schedule_object-phases-discounts
  def discounts_from_sf_order_item(sf_order_item:)
    sf_coupons = StripeForce::Translate.get_salesforce_stripe_coupons_associated_to_order_line(sf_client: @user.sf_client, sf_order_line_id: sf_order_item.Id)
    if !sf_coupons || sf_coupons.empty?
      return
    end

    log.info 'found coupons for sf order item', salesforce_object: sf_order_item

    # this is the format the stripe api expects
    sf_coupons.map do |sf_coupon|
      coupon = translate_coupon(sf_coupon)
      {coupon: coupon.id}
    end
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
