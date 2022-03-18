# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::PriceTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
    @locker = Integrations::Locker.new(@user)
  end

  it 'translates a pricebook entry to a stripe price' do
    skip("right now the pricebook translation is tied to a order line AND a pricebook entry")

    sf_pricebook_entry_id = create_salesforce_price
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    # TODO this isn't going to work right now
    stripe_price = StripeForce::Translate.perform(user: @user, sf_object: sf_pricebook_entry, locker: @locker)

    # TODO this will break when we create dynamic
    refute_nil(stripe_price.metadata['salesforce_id'])
    refute_nil(stripe_price.metadata['salesforce_url'])

    # the ID cannot be specified in Stripe for a price
    refute_equal(stripe_price.id, sf_pricebook_entry_id)
    # TODO assert on price, recurrance, etc
  end
end
