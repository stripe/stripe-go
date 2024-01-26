# frozen_string_literal: true
# typed: ignore

require_relative '../../test_helper'

class Critic::AccountPollerTest < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || DateTime.now.utc)

    @user = make_user(save: true)
    inline_job_processing!
  end

  it 'does not poll if polling is disabled' do
    @user.enable_feature(FeatureFlags::ACCOUNT_POLLING)
    @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = false
    @user.save

    locker = Integrations::Locker.new(@user)

    StripeForce::AccountPoller.perform(user: @user, locker: locker)
    assert_equal(0, StripeForce::PollTimestamp.count)
  end

  it 'polls when feature flag ACCOUNT_POLLING is enabled' do
    # enable account polling
    @user.enable_feature(FeatureFlags::ACCOUNT_POLLING)
    @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = true
    @user.save

    # set up the intiial poll timestamp to enable account polling
    initial_poll_timestamp = set_initial_poll_timestamp(SF_ACCOUNT).last_polled_at.utc

    # create an sf account and translate
    sf_account_id = create_salesforce_account
    StripeForce::Translate.perform_inline(@user, sf_account_id)

    locker = Integrations::Locker.new(@user)

    one_min_in_future = VCR.current_cassette.originally_recorded_at ? VCR.current_cassette.originally_recorded_at + 60.seconds : DateTime.now.utc + 60.seconds
    Timecop.freeze(one_min_in_future)

    # kick off account poll job for this user
    StripeForce::AccountPoller.perform(user: @user, locker: locker)

    # get the poll timestamp for this user's account
    poll_timestamp = StripeForce::PollTimestamp.by_user_and_record(
      @user,
      SF_ACCOUNT
    )
    # confirm we have polled again since the initial poll timestamp
    assert(poll_timestamp.last_polled_at - initial_poll_timestamp > initial_poll_delta)
    assert_equal(SF_ACCOUNT, poll_timestamp.integration_record_type)

    # disable feature ACCOUNT_POLLING
    @user.disable_feature(FeatureFlags::ACCOUNT_POLLING)

    # assert that only order polling will occur (not account polling since feature flag is disabled)
    StripeForce::InitiatePollsJobs.expects(:queue_poll_job_for_user).once do |args|
      assert_equal(StripeForce::OrderPoller, args[:poller_job])
    end

    # kick off account poll job for this user
    StripeForce::InitiatePollsJobs.queue_polls_for_user(@user)
  end

  it 'polls for default account only' do
    # enable account polling
    @user.enable_feature(FeatureFlags::ACCOUNT_POLLING)
    @user.enable_feature(FeatureFlags::MULTI_STRIPE_ACCOUNT)
    @user.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = true
    @user.is_default_account_config = false
    @user.save

    locker = Integrations::Locker.new(@user)
    StripeForce::AccountPoller.perform(user: @user, locker: locker)
    assert_equal(0, StripeForce::PollTimestamp.count)
  end
end
