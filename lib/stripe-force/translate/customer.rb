# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_account(sf_account)
    locker.lock_salesforce_record(sf_account)

    catch_errors_with_salesforce_context(secondary: sf_account) do
      existing_stripe_customer = retrieve_from_stripe(Stripe::Customer, sf_account)
      if existing_stripe_customer.blank?
        # creates a new stripe customer
        create_customer_from_sf_account(sf_account)
      else
        # check if the feature flag to update an existing customer is enabled
        if @user.feature_enabled?(FeatureFlags::UPDATE_CUSTOMER_ON_ORDER_TRANSLATION)
          update_stripe_object(stripe_class: Stripe::Customer, stripe_object_id: existing_stripe_customer.id, sf_object: sf_account)
        else
          existing_stripe_customer
        end
      end
    end
  end

  def create_customer_from_sf_account(sf_account)
    log.info 'translating customer', salesforce_object: sf_account

    customer, response = create_stripe_object(Stripe::Customer, sf_account) do |generated_stripe_customer|
      if @user.feature_enabled?(FeatureFlags::TEST_CLOCKS) && !@user.livemode
        log.debug 'adding test clock to customer'

        test_clock = Stripe::TestHelpers::TestClock.create({
          frozen_time: Time.now.to_i,
        }, @user.stripe_credentials)

        generated_stripe_customer.test_clock = test_clock.id
      end
    end

    update_sf_stripe_id(sf_account, customer, stripe_response: response)

    customer
  end
end
