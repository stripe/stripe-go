# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class LockerTest < Critic::UnitTest
    before do
      @user = make_user
    end

    it 'refreshes a user lock' do
      past_expiration_time = (Time.now.to_i - 1_000).to_f

      locker = Integrations::Locker.new(@user)
      locker.lock_on_user do
        Resque.redis.set(locker.user_key, past_expiration_time)

        locker.lock_stripe_resource(Stripe::Charge.construct_from(id: stripe_create_id(:ch)))
        locker.refresh_user_lock

        updated_lock_time = Resque.redis.get(locker.user_key).to_f

        refute_equal(past_expiration_time, updated_lock_time)
        assert(updated_lock_time > past_expiration_time)
        assert(updated_lock_time < (Time.now.to_i + Integrations::Locker::LOCK_EXPIRATION_TIME + 1))
      end

      assert_empty(Resque.redis.keys)
    end

    it 'removes the user lock across expiration windows' do
      locker = Integrations::Locker.new(@user)
      locker.stubs(:generate_expiration).
        # first expiration date is in the past
        returns(Time.now.to_f - 1_000).then.
        # subsequent date is in the future
        returns(Time.now.to_f + 1_000)

      locker.lock_on_user do
        locker.lock_stripe_resource(Stripe::Charge.construct_from(id: stripe_create_id(:ch)))
        locker.refresh_user_lock
      end

      assert_empty(redis.keys)
    end

    it 'locks and refreshes a poll job' do
      locker = Integrations::Locker.new(@user)
      locker.lock_on_poll_job(StripeForce::OrderPoller)

      keys = redis.keys
      key = keys.first

      assert_equal(1, keys.count)
      assert_match(/StripeForce::OrderPoller/, key)
      assert_match(/#{@user.stripe_account_id}/, key)

      # a second locker (i.e. another process) on the class should fail
      second_locker = Integrations::Locker.new(@user)
      assert_raises(Integrations::Errors::LockTimeout) do
        second_locker.lock_on_poll_job(StripeForce::OrderPoller)
      end

      # grabbing the lock using the same locker should refresh the lock
      initial_expiration = redis.get(key)
      locker.lock_on_poll_job(StripeForce::OrderPoller)
      assert(initial_expiration < redis.get(key))
    end

    describe 'stripe resource locking' do
      it 'locks on a stripe resource' do
        charge_id = stripe_create_id(:ch)
        charge = Stripe::Charge.construct_from(id: charge_id)

        locker = Integrations::Locker.new(@user)
        locker.lock_stripe_resource(charge)

        keys = Resque.redis.keys

        assert_equal(1, keys.size)
        key = keys.first
        assert_match(/#{@user.stripe_account_id}/, key)
        assert_match(/#{charge_id}/, key)

        locker.clear_locked_resources(except: charge)

        assert_equal(1, Resque.redis.keys.size)

        locker.clear_locked_resources

        assert_empty(Resque.redis.keys)
      end

      # NOTE translator shouldn't have to worry about this
      it 'does not throw an exception if a lock on a resource is requested twice on the same locker' do
        charge_id = stripe_create_id(:ch)
        charge = Stripe::Charge.construct_from(id: charge_id)

        locker = Integrations::Locker.new(@user)
        locker.lock_stripe_resource(charge)
        locker.lock_stripe_resource(charge)
      end

      it 'refreshes a stripe resource lock if it is already locked' do
        past_expiration_time = (Time.now.to_i - 1_000).to_f

        charge_id = stripe_create_id(:ch)
        charge = Stripe::Charge.construct_from(id: charge_id)

        locker = Integrations::Locker.new(@user)

        charge_lock_key = locker.send(:generate_stripe_resource_lock_key, charge)
        Resque.redis.set(charge_lock_key, past_expiration_time)

        locker.lock_stripe_resource(charge)

        updated_lock_time = Resque.redis.get(charge_lock_key).to_f

        refute_equal(past_expiration_time, updated_lock_time)
        assert(updated_lock_time > past_expiration_time)
        assert(updated_lock_time < (Time.now.to_i + Integrations::Locker::LOCK_EXPIRATION_TIME + 1))
      end

      it 'allows a custom expiration window to be specified' do
        custom_expiration_time = 60 * 60 * 24

        charge_id = stripe_create_id(:ch)
        charge = Stripe::Charge.construct_from(id: charge_id)

        locker = Integrations::Locker.new(@user)

        charge_lock_key = locker.send(:generate_stripe_resource_lock_key, charge)

        locker.lock_stripe_resource(charge, expiration_time: custom_expiration_time)

        lock_expiration = Resque.redis.get(charge_lock_key).to_f

        refute_nil(lock_expiration)
        assert(lock_expiration > (Time.now.to_i + Integrations::Locker::LOCK_EXPIRATION_TIME))
        assert(lock_expiration < (Time.now.to_i + custom_expiration_time + 1))
      end
    end

    describe 'salesforce record locking' do
      it 'throws an exception if a lock is already on a record' do
        sf_order = create_mock_salesforce_order

        locker = Integrations::Locker.new(@user)
        locker.lock_salesforce_record(sf_order)

        locker_2 = Integrations::Locker.new(@user)
        assert_raises(Integrations::Errors::LockTimeout) { locker_2.lock_salesforce_record(sf_order) }
      end

      it 'does not throw an exception if a lock on a resource is requested twice on the same locker' do
        sf_order = create_mock_salesforce_order

        locker = Integrations::Locker.new(@user)
        locker.lock_salesforce_record(sf_order)
        locker.lock_salesforce_record(sf_order)
      end

      it 'locks on a salesforce record' do
        sf_order = create_mock_salesforce_order

        charge_id = stripe_create_id(:ch)
        charge = Stripe::Charge.construct_from(id: charge_id)

        locker = Integrations::Locker.new(@user)
        locker.lock_salesforce_record(sf_order)
        locker.lock_stripe_resource(charge)

        keys = Resque.redis.keys

        assert_equal(2, keys.size)
        key = keys.detect {|k| k =~ /-record-/ }
        refute_match(/#{@user.stripe_account_id}/, key)
        assert_match(/#{@user.salesforce_account_id}/, key)
        assert_match(/#{sf_order.sobject_type}/, key)
        assert_match(/#{sf_order.Id}/, key)

        locker.clear_locked_resources(except: sf_order)

        assert_equal(1, Resque.redis.keys.size)
        key = Resque.redis.keys.first
        assert_match(/#{sf_order.Id}/, key)
        assert_match(/#{sf_order.sobject_type}/, key)

        locker.clear_locked_resources

        assert_empty(Resque.redis.keys)
      end
    end
  end
end
