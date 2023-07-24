# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_product(sf_product)
    locker.lock_salesforce_record(sf_product)

    catch_errors_with_salesforce_context(secondary: sf_product) do
      existing_stripe_product = retrieve_from_stripe(Stripe::Product, sf_product)
      if existing_stripe_product.blank?
        create_product_from_sf_product(sf_product)
      else
        # determine if we should update the Stripe product or not
        stripe_objects_are_similar = stripe_products_are_equal?(existing_stripe_object: existing_stripe_product, sf_object: sf_product)
        if @user.feature_enabled?(FeatureFlags::UPDATE_PRODUCT_ON_SYNC) && stripe_objects_are_similar
          update_stripe_object(stripe_class: Stripe::Product, stripe_object_id: existing_stripe_product.id, sf_object: sf_product)
        else
          existing_stripe_product
        end
      end
    ensure
      locker.release_salesforce_record_lock(sf_product)
    end
  end

  def create_product_from_sf_product(sf_product)
    log.info 'translating product', salesforce_object: sf_product

    stripe_product = create_stripe_object(Stripe::Product, sf_product)

    update_sf_stripe_id(sf_product, stripe_product)

    stripe_product
  end

  sig { params(existing_stripe_object: Stripe::StripeObject, sf_object: Restforce::SObject).returns(T::Boolean) }
  def stripe_products_are_equal?(existing_stripe_object:, sf_object:)
    # we want to enable users to update some fields on the Stripe object (e.g, metadata) without triggering a new Stripe object creation
    # but we also want to make sure that the Salesforce object wasn't mutated such that it would be a different object
    fields_to_check_if_both_are_set = %i{
      default_price
    }

    new_stripe_object = construct_stripe_object(stripe_class: Stripe::Product, salesforce_object: sf_object)

    simple_field_check_passed = fields_to_check_if_both_are_set.all? do |field_sym|
      existing_stripe_object[field_sym].blank? == new_stripe_object[field_sym].blank? && existing_stripe_object[field_sym] == new_stripe_object[field_sym]
    end

    if !simple_field_check_passed
      log.info 'stripe products are not equal, simple field comparison failed', diff: HashDiff::Comparison.new(existing_stripe_object.to_hash, new_stripe_object.to_hash).diff
    end

    simple_field_check_passed
  end
end
