# frozen_string_literal: true
# typed: strict

require_relative '../test_helper'

module Critic::Unit
  class KMSEncryptionTest < Critic::UnitTest
    before do
      KMSEncryptionTestHelpers.restore_encryption_fields(StripeForce::User)
    end

    it "does not encrypt nil values on local fields" do
      user = make_user(save: true)

      Aws::KMS::Client.any_instance.expects(:encrypt).never

      user.set(
        salesforce_token: nil,
        salesforce_refresh_token: nil
      )

      user.save
    end

    it 'returns nil when encrypted field data is nil' do
      user = make_user
      user.set(salesforce_token: nil)

      assert_nil(user.encrypted_salesforce_token)
      assert_nil(user.salesforce_token)
    end

    it 'encrypts a fresh token using kms, and retrieves the unencrypted value using kms' do
      user = make_user
      user.set(salesforce_token: 'foo_token')

      # until the record is saved, the unencrypted value is stored locally in the record
      assert_equal('foo_token', user.salesforce_token)
      assert_nil(user.encrypted_salesforce_token)

      user.save

      # after save, the token is encrypted
      refute_nil(user.encrypted_salesforce_token)

      # grab a fresh model to ensure all kms-related cache on model is clear
      user = StripeForce::User[user.id]

      assert(!T.must(user).encrypted_salesforce_token.nil? && !T.must(T.must(user).encrypted_salesforce_token).empty?, "encrypted value should not be empty or nil")
      refute_equal('foo_token', T.must(user).encrypted_salesforce_token)
      assert_equal('foo_token', T.must(user).salesforce_token)
    end
  end
end
