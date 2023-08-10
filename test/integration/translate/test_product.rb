# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::ProductTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
  end

  describe 'product reuse' do
    it 'uses the same product if it translated twice' do
      sf_product_id, _ = salesforce_recurring_product_with_price

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

    it 'updates the Stripe product if it translated again and has changed' do
      @user.enable_feature(FeatureFlags::UPDATE_PRODUCT_ON_SYNC)
      @user.save

      @user.field_defaults = {
        "product" => {
          "tax_code" => "txcd_99999999",
          "shippable" => true,
        },
      }
      @user.save

      sf_product_id, _ = salesforce_recurring_product_with_price(price: 1000, currency_iso_code: nil, additional_product_fields: {
        "Name" => "Gold Subscription",
        "Description" => "Best subscription ever.",
      })

      StripeForce::Translate.perform_inline(@user, sf_product_id)
      sf_product = sf.find(SF_PRODUCT, sf_product_id)
      stripe_product_id = sf_product[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      refute_nil(stripe_product_id)

      # fetch the Stripe product and assert the product was created correctly
      stripe_product = Stripe::Product.retrieve(stripe_product_id, @user.stripe_credentials)
      assert_equal(sf_product["Name"], stripe_product.name)
      assert_equal(sf_product["Description"], stripe_product.description)
      assert_equal("txcd_99999999", stripe_product.tax_code)
      assert(stripe_product.shippable)

      # update the related mappings and Salesforce product
      @user.field_defaults = {
        "product" => {
          "statement_descriptor" => "Gold Sub",
          "metadata.metadata_key_1" => "metadata_value_1",
        },
      }
      @user.save

      @user.sf_client.update!(SF_PRODUCT, {
        SF_ID => sf_product_id,
        "Name" => "Gold Subscription 2",
      })
      sf_product.refresh

      # resync the Stripe product and assert the product fields were updated
      Stripe::Product.expects(:create).never
      stripe_product = StripeForce::Translate.perform_inline(@user, sf_product_id)
      assert_equal(stripe_product.id, sf_product[prefixed_stripe_field(GENERIC_STRIPE_ID)])
      assert_equal(sf_product["Name"], stripe_product.name)
      assert_equal(sf_product["Description"], stripe_product.description)
      assert_equal("Gold Sub", stripe_product.statement_descriptor)
      assert_equal("metadata_value_1", stripe_product.metadata["metadata_key_1"])

      # TODO we should enable users to unset values
      # ArgumentError: You cannot set tax_code to an empty string. We interpret empty strings as nil in requests. You may set (object).tax_code = nil to delete the property.
      # assert_empty(stripe_product.tax_code)
    end
  end
end
