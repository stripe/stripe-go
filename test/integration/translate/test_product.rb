# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::ProductTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  describe 'price reuse' do
    it 'uses the same product if it translated twice' do
      sf_product_id = create_salesforce_product

      StripeForce::Translate.perform_inline(@user, sf_product_id)

      sf_product = sf.find(SF_PRODUCT, sf_product_id)
      stripe_product_id = sf_product[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_product_id)

      stripe_product = Stripe::Product.retrieve(stripe_product_id, @user.stripe_credentials)

      Stripe::Product.expects(:create).never

      StripeForce::Translate.perform_inline(@user, sf_product_id)

      sf_product.refresh
      assert_equal(stripe_product.id, sf_product[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    end
  end
end
