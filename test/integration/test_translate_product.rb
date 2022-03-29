# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ProductTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'translates a subscription product' do
    sf_product_id = create_salesforce_product

    stripe_product = StripeForce::Translate.perform_inline(@user, sf_product_id)

    sf_product = sf.find(SF_PRODUCT, sf_product_id)
    assert_match(sf_product_id, stripe_product.metadata['salesforce_product2_link'])
    assert_equal(stripe_product.metadata['salesforce_product2_id'], sf_product_id)

    refute_equal(stripe_product.id, sf_product_id)
    assert_equal(stripe_product.description, sf_product.Description)
    assert_equal(stripe_product.name, sf_product.Name)
  end
end
