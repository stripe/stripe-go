# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class StripeUtilTest < Critic::UnitTest
    before do
      @user = make_user
    end

    describe 'generate_idempotency_key_with_credentials' do
      it 'generates a key' do
        sf_order = create_mock_salesforce_order

        creds = StripeForce::Utilities::StripeUtil.generate_idempotency_key_with_credentials(@user, sf_order)

        refute_nil(sf_order.Id)
        ikey = creds.delete(:idempotency_key)
        timestamp = DateTime.parse(sf_order[SF_LAST_MODIFIED_DATE]).to_i
        assert_match(sf_order.Id, ikey)
        assert_match(timestamp.to_s, ikey)
        assert_equal(@user.stripe_credentials, creds)
      end

      it 'generates a key with an action' do
        sf_order = create_mock_salesforce_order

        creds = StripeForce::Utilities::StripeUtil.generate_idempotency_key_with_credentials(
          @user,
          sf_order,
          :finalize_invoice
        )

        ikey = creds.delete(:idempotency_key)
        assert_match(sf_order.Id.to_s, ikey)
        assert_match("finalize_invoice", ikey)
        assert_equal(@user.stripe_credentials, creds)
      end

      it 'generates a different key when mappings change' do
        sf_order = create_mock_salesforce_order

        cred_1 = StripeForce::Utilities::StripeUtil.generate_idempotency_key_with_credentials(@user, sf_order)
        key_1 = cred_1[:idempotency_key]

        # now let's update the mapping
        @user.field_mappings['price'] = {'some' => 'map'}

        cred_2 = StripeForce::Utilities::StripeUtil.generate_idempotency_key_with_credentials(@user, sf_order)
        key_2 = cred_2[:idempotency_key]

        refute_equal(key_2, key_1)

        # max idempotency key size
        assert(key_2.size < 255)
        assert(key_1.size < 255)
      end
    end

  end
end
