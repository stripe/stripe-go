# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ProductTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    @locker = Integrations::Locker.new(@user)
  end

  def sf_randomized_name(sf_object_name)
    "REST Product #{DateTime.now}"
  end

  def create_salesforce_product
    sf.create!(SF_PRODUCT,
      "Name" => sf_randomized_name(SF_PRODUCT),
      "Description" => "A great description",

      # this field is an enum/picklist in SF
      CPQ_PRODUCT_SUBSCRIPTION_TYPE => 'Renewable'
    )
  end

  it 'translates a subscription product' do
    sf_product_id = create_salesforce_product
    sf_product = sf.find(SF_PRODUCT, sf_product_id)

    stripe_product = StripeForce::Translate.perform(user: @user, sf_object: sf_product, locker: @locker)

    # TODO this will break when we create dynamic
    refute_nil(stripe_product.metadata['salesforce_id'])
    refute_nil(stripe_product.metadata['salesforce_url'])

    assert_equal(stripe_product.id, sf_product_id)
    assert_equal(stripe_product.description, sf_product.Description)
    assert_equal(stripe_product.name, sf_product.Name)
  end
end