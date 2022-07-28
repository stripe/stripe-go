# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_account(sf_account)
    create_customer_from_sf_account(sf_account)
  end

  def create_customer_from_sf_account(sf_account)
    log.info 'translating customer', salesforce_object: sf_account

    if (stripe_customer = retrieve_from_stripe(Stripe::Customer, sf_account))
      return stripe_customer
    end

    customer = create_stripe_object(Stripe::Customer, sf_account) do |generated_stripe_customer|
      # passing a partial shipping hash will trigger an error
      if !generated_stripe_customer.shipping.respond_to?(:address) || generated_stripe_customer.shipping.address.to_h.empty?
        log.info 'no address on shipping hash, removing'
        generated_stripe_customer.shipping = {}
      end
    end

    update_sf_stripe_id(sf_account, customer)

    customer
  end
end
