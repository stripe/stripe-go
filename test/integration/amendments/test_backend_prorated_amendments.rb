# frozen_string_literal: true
# typed: true

require_relative './_lib'

class Critic::BackendProratedAmendmentTranslation < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'throws an error if the subscription term is not divisible by billing frequency and is greater than one' do
    skip("this case will be handled in an upcoming release, right now this functionality is undefined")

    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(
      additional_product_fields: {
        CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::ANNUAL.serialize,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12,
      }
    )

    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_TERM => 13,
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
      }
    )

    exception = assert_raises(Integrations::Errors::UserError) do
      StripeForce::Translate.perform_inline(@user, sf_order.Id)
    end

    assert_match("Prorated order amendments are not yet supported", exception.message)
  end
end
