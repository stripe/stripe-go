# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_product(sf_product)
    catch_errors_with_salesforce_context(secondary: sf_product) do
      create_product_from_sf_product(sf_product)
    end
  end

  def create_product_from_sf_product(sf_product)
    log.info 'translating product', salesforce_object: sf_product

    if (stripe_product = retrieve_from_stripe(Stripe::Product, sf_product))
      return stripe_product
    end

    stripe_product = create_stripe_object(Stripe::Product, sf_product)

    update_sf_stripe_id(sf_product, stripe_product)

    stripe_product
  end
end
