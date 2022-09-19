# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::BaseJobTest < Critic::UnitTest
  extend T::Sig

  it 'picks the correct user when there are multiple tied to the same stripe account' do
    user_1 = make_user(random_user_id: false)
    user_1.salesforce_account_id = create_salesforce_id
    user_1.enabled = false
    user_1.save

    user_2 = make_user(random_user_id: false, save: true)

    assert_equal(user_1.stripe_account_id, user_2.stripe_account_id)
    refute_equal(user_1.salesforce_account_id, user_2.salesforce_account_id)

    # making the assumption that the credential check runs before the enabled check
    SalesforceTranslateRecordJob.expects(:valid_system_credentials!).with do |selected_user|
      assert_equal(selected_user.id, user_1.id)
    end

    SalesforceTranslateRecordJob.work(user_1, create_salesforce_id)
  end

  it 'handles invalid object id' do
    user = make_user(save: true)

    exception = assert_raises(ArgumentError) do
      SalesforceTranslateRecordJob.work(user, 'a_very_bad_id')
    end

    assert_match('unknown object type', exception.message)
  end
end
