# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_account(sf_account)
    catch_errors_with_salesforce_context(secondary: sf_account) do
      create_customer_from_sf_account(sf_account)
    end
  end

  def create_customer_from_sf_account(sf_account)
    log.info 'translating customer', salesforce_object: sf_account

    if (stripe_customer = retrieve_from_stripe(Stripe::Customer, sf_account))
      return stripe_customer
    end

    customer = create_stripe_object(Stripe::Customer, sf_account) do |generated_stripe_customer|
      if @user.feature_enabled?(FeatureFlags::TEST_CLOCKS) && !@user.livemode
        log.debug 'adding test clock to customer'

        test_clock = Stripe::TestHelpers::TestClock.create({
          frozen_time: Time.now.to_i,
        }, @user.stripe_credentials)

        generated_stripe_customer.test_clock = test_clock.id
      end
    end

    update_sf_stripe_id(sf_account, customer)

    customer
  end
end
