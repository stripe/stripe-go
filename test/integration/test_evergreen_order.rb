# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::EvergreenOrderTest < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  # for creating evergreen orders specifically. evergreen subscriptions do not have an end date.
  def create_evergreen_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, currency_iso_code: nil, additional_quote_fields: {})
    if !sf_product_id
      sf_product_id, _ = salesforce_evergreen_product_with_price(currency_iso_code: currency_iso_code)
    end

    # sf_product_id should point to an evergreen product so will not be nil inside create sf order
    create_salesforce_order(sf_product_id: sf_product_id, sf_account_id: sf_account_id, sf_pricebook_id: sf_pricebook_id, currency_iso_code: currency_iso_code, additional_quote_fields: additional_quote_fields)
  end

  it 'creates an evergreen order object' do
    @user.field_defaults['customer'] = {
      'email' => create_random_email,
    }

    sf_account_id = create_salesforce_account

    sf_order = create_evergreen_salesforce_order(
      sf_account_id: sf_account_id
    )

    # get the quote to check the subscription data inside
    sf_quote_id = sf_order[CPQ_QUOTE]

    # TODO: use better way to get quote object, i.e. sf_quote = sf_get(sf_quote_id)
    # https://jira.corp.stripe.com/browse/PLATINT-2542
    quote_data = JSON.parse(sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{sf_quote_id}").body)
    sf_quote = quote_data['lineItems'].first['record']['SBQQ__Product__r']

    assert_equal(1, sf_quote[CPQ_QUOTE_SUBSCRIPTION_TERM])
    assert_equal('Evergreen', sf_quote[CPQ_PRODUCT_SUBSCRIPTION_TYPE])
    assert_equal('Fixed Price', sf_quote[CPQ_QUOTE_SUBSCRIPTION_PRICING])
  end
end
