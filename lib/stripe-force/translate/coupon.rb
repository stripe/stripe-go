# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_coupon(order_sf_coupon)
    log.info 'translating coupon', salesforce_object: order_sf_coupon

    catch_errors_with_salesforce_context(secondary: order_sf_coupon) do
      create_coupon_from_sf_coupon(order_sf_coupon)
    end
  end

  def create_coupon_from_sf_coupon(order_sf_coupon)
    # check if the quote sf coupon has been translated prior and exists in Stripe
    # the quote coupon is the coupon that the order / order item coupon was copied from
    quote_sf_coupon = sf.find(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), order_sf_coupon[prefixed_stripe_field("Quote_Stripe_Coupon_Id__c")])
    existing_stripe_coupon = retrieve_from_stripe(Stripe::Coupon, quote_sf_coupon)
    if existing_stripe_coupon
      existing_stripe_coupon = T.cast(existing_stripe_coupon, Stripe::Coupon)
      generated_stripe_coupon = construct_stripe_object(stripe_class: Stripe::Coupon, salesforce_object: quote_sf_coupon)

      if coupons_are_equal?(existing_stripe_coupon: existing_stripe_coupon, generated_stripe_coupon: generated_stripe_coupon)
        log.info 'reusing existing stripe coupon', existing_stripe_coupon_id: existing_stripe_coupon.id
        return existing_stripe_coupon
      else
        # this should never happen unless the coupon data in Salesforce is mutated
        # if so, we will log and continue to create a new stripe coupon
        log.info 'salesforce and stripe coupon differ. creating a new stripe coupon', existing_stripe_coupon_id: existing_stripe_coupon.id
      end
    end

    # create a new stripe coupon
    stripe_coupon = create_stripe_object(Stripe::Coupon, order_sf_coupon) do |stripe_coupon_data|
      sanitize_stripe_coupon(stripe_coupon_data)
    end

    # update the sf order coupon stripe id
    update_sf_stripe_id(order_sf_coupon, stripe_coupon)
    # also update the quote sf coupon stripe id so the same coupon object can be reused in another quote
    update_sf_stripe_id(quote_sf_coupon, stripe_coupon)

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

      generated_stripe_coupon = sanitize_stripe_coupon(generated_stripe_coupon)

      simple_field_check_passed = fields_to_check_if_both_are_set.all? do |field_sym|
        existing_stripe_coupon[field_sym].blank? == generated_stripe_coupon[field_sym].blank? && existing_stripe_coupon[field_sym] == generated_stripe_coupon[field_sym]
      end

      if !simple_field_check_passed
        log.info 'coupons are not equal, simple field comparison failed', diff: HashDiff::Comparison.new(existing_stripe_coupon.to_hash, generated_stripe_coupon.to_hash).diff
      end
      simple_field_check_passed
  end

  def sanitize_stripe_coupon(stripe_coupon)
    # Stripe expects amount_off to be specified in cents
    if !stripe_coupon["amount_off"].nil?
      stripe_coupon["amount_off"] = normalize_float_amount_for_stripe(stripe_coupon["amount_off"].to_s, @user)
      stripe_coupon["currency"] = 'usd'
    end

    # Prevents Stripe API error 'Stripe::InvalidRequestError: Invalid integer: 1.0'
    # since Salesforce stores these values as decimals (1.0) but Stripe expects an integer
    if !stripe_coupon["duration_in_months"].nil?
      stripe_coupon["duration_in_months"] = stripe_coupon["duration_in_months"].to_i
    end

    if !stripe_coupon["max_redemptions"].nil?
      stripe_coupon["max_redemptions"] = stripe_coupon["max_redemptions"].to_i
    end

    stripe_coupon
  end

  # Coupon can be related to an order or an order item. Either way, Stripe expects these coupons to be specified as discounts:
  # https://site-admin.stripe.com/docs/api/subscription_schedules/object#subscription_schedule_object-phases-discounts
  def stripe_discounts_for_sf_object(sf_object:)
    sf_coupons = get_salesforce_stripe_coupons_associated_to_sf_object(sf_client: @user.sf_client, sf_object: sf_object)

    if sf_coupons.nil?
      return
    end

    log.info 'found coupons for sf object', salesforce_object: sf_object

    sf_coupons.map do |sf_coupon|
      coupon = translate_coupon(sf_coupon)
      {coupon: coupon.id}
    end
  end

  def get_salesforce_stripe_coupons_associated_to_sf_object(sf_client:, sf_object:)
    catch_errors_with_salesforce_context(secondary: sf_object) do
      # coupons can either be related to an order or order item
      source_sf_record_type = sf_object.sobject_type
      if source_sf_record_type == SF_ORDER
        lookup_field = 'Order__c'
      elsif source_sf_record_type == SF_ORDER_ITEM
        lookup_field = 'Order_Item__c'
      else
        # this should never happen since coupons can only be tied to an order or order item
        raise Integrations::Errors::ImpossibleState.new("unsupported sf object type for coupons: #{source_sf_record_type}")
      end

      # check if there are any coupon associations to this order or order item
      coupon_data = sf_client.query("Select #{SF_ID} from #{prefixed_stripe_field(ORDER_SF_STRIPE_COUPON)} where #{prefixed_stripe_field(lookup_field)} = '#{sf_object.Id}'")

      # there could be multiple coupons associated with a single order or order line
      coupons = coupon_data.map do |data|
        sf_client.find(prefixed_stripe_field(ORDER_SF_STRIPE_COUPON), data.Id)
      end

      coupons
    end
  end
end
