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
        if @user.feature_enabled?(FeatureFlags::UPDATE_PRODUCT_ON_SYNC)
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
end
