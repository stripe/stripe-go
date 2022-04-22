# frozen_string_literal: true
# typed: true

require_relative '../../test_helper'

class Critic::IntegrationOverrides < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  it 'integrates a subscription order when an invalid ID is entered in the stripe ID field' do
    sf_order = create_subscription_order

    sf.update!(sf_order.sobject_type, {
      SF_ID => sf_order.Id,
      # insert a random ID that does not represent a stripe object
      prefixed_stripe_field(GENERIC_STRIPE_ID) => SecureRandom.alphanumeric(16),
    })

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh

    assert_match(/^sub_sched_/, sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
  end

  it 'ignores a valid stripe ID of the wrong object type' do
    stripe_customer = Stripe::Customer.create({}, @user.stripe_credentials)

    sf_order = create_salesforce_order(additional_quote_fields: {
      CPQ_QUOTE_SUBSCRIPTION_START_DATE => DateTime.now.strftime("%Y-%m-%d"),
      CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
    })

    sf.update!(sf_order.sobject_type, {
      SF_ID => sf_order.Id,
      # insert a ID that represents a valid stripe object of the wrong type
      prefixed_stripe_field(GENERIC_STRIPE_ID) => stripe_customer.id,
    })

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh

    assert_match(/^sub_sched_/, sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)])
    schedule = Stripe::SubscriptionSchedule.retrieve(sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)
  end

  it 'accepts a valid stripe ID and does not create a additional object in Stripe' do
    stripe_customer = Stripe::Customer.create({}, @user.stripe_credentials)
    stripe_customer.metadata['custom'] = 'metadata'
    stripe_customer.save

    sf_account_id = create_salesforce_account(additional_fields: {
      prefixed_stripe_field(GENERIC_STRIPE_ID) => stripe_customer.id,
    })

    Stripe::Customer.expects(:create).never

    StripeForce::Translate.perform_inline(@user, sf_account_id)

    sf_customer = sf.find(SF_ACCOUNT, sf_account_id)
    assert_equal(stripe_customer.id, sf_customer[prefixed_stripe_field(GENERIC_STRIPE_ID)])

    stripe_customer.refresh
    assert_match(sf_account_id, stripe_customer.metadata['salesforce_account_link'])
    assert_equal(stripe_customer.metadata['salesforce_account_id'], sf_account_id)

    # ensure the existing metadata is not overwritten
    assert_equal('metadata', stripe_customer.metadata['custom'])
  end

end
