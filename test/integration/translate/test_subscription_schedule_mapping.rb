# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

# there are some 'special' fields within the subscription schedule mapping
class Critic::SubscriptionScheduleMappingTest < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  describe '#type' do
    it 'infers the type of a amendment from the databasen structure, not the type field' do
      # TODO set `Type = Something Crazy` on the order and check to make sure it's still determined to be new
      # we need to modify our types enum, and serialize that change in our sandbox setup definition, which is why this TODO is here
    end
  end

  describe 'days_until_due' do
    it 'maps an enum value from SBQQ__PaymentTerms__c' do
      @user.field_defaults['subscription_schedule'] = {
        'default_settings.collection_method' => 'send_invoice',
      }
      @user.field_mappings['customer'] = {
        'email' => 'Description',
      }
      @user.field_mappings['subscription_schedule'] = {
        'default_settings.invoice_settings.days_until_due' => 'SBQQ__PaymentTerm__c',
      }
      @user.save

      sf_account_id = create_salesforce_account(additional_fields: {'Description' => create_random_email})
      sf_order = create_subscription_order(sf_account_id: sf_account_id)
      assert_equal('Net 30', sf_order.SBQQ__PaymentTerm__c)

      StripeForce::Translate.perform_inline(@user, sf_order.Id)

      sf_order.refresh
      stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
      subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)

      assert_equal(30, subscription_schedule.default_settings.invoice_settings.days_until_due)
    end
  end
end
