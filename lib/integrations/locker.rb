# frozen_string_literal: true
# typed: true

# modified version of Redis::Lock
# https://github.com/nateware/redis-objects/blob/master/lib/redis/lock.rb

module Integrations
  class UserLocker
    extend T::Sig
    include Log

    # expiration should exceed NS timeouts configured in `init_netsuite_client`
    LOCK_EXPIRATION_TIME = 60 * 5
    NETSUITE_DEFAULT_SUITECLOUD_CONNECTIONS = 10

    # TODO https://github.com/stripe/stripe-netsuite/issues/902
    sig { params(user: StripeForce::User).returns(String) }
    def self.generate_netsuite_account_lock_key(user)
      "user-#{user.salesforce_account_id}"
    end

    attr_accessor :user, :user_key

    sig { params(user: StripeForce::User).void }
    def initialize(user)
      @user = user

      # TODO should be renamed to something more generic
      @locked_resource_keys = []
    end

    # TODO https://github.com/stripe/stripe-netsuite/issues/1516

    # https://github.com/stripe/stripe-netsuite/issues/903
    sig { params(expiration_time: T.nilable(Integer)).void }
    def refresh_user_lock(expiration_time: nil)
      if !user_key.nil?
        # NOTE setting @expiration is critical! `lock_on_user` looks at this variable after the passed block is yield'd
        @expiration = generate_expiration(expiration_time: expiration_time)
        redis.set(user_key, @expiration)
      end
    end

    # `except` must be an array of Stripe or NetSuite record instances, or an empty array
    # this allows the payout transactions to be backfilled without loosing the payout lock
    sig { params(except: T.untyped).void }
    def clear_locked_resources(except: [])
      if !except.is_a?(Array)
        except = [except]
      end

      keys_to_clear = @locked_resource_keys - except.map do |r|
        if r.class.to_s.start_with?('Stripe::')
          generate_stripe_resource_lock_key(r)
        elsif r.class.to_s.start_with?('NetSuite::')
          generate_netsuite_record_lock_key(r)
        else
          raise "unsupported exception key #{r}"
        end
      end

      keys_to_clear.each {|k| redis.del(k) }

      @locked_resource_keys -= keys_to_clear
    end

    # https://github.com/stripe/stripe-netsuite/issues/1152
    sig { params(job: Class).void }
    def lock_on_poll_job(job)
      job_key = "#{@user.stripe_account_id}-#{job}"
      acquire_or_refresh_resource_lock(job_key)
    end

    sig { params(stripe_resource: Stripe::APIResource, expiration_time: T.nilable(Integer)).void }
    def lock_stripe_resource(stripe_resource, expiration_time: nil)
      resource_lock_key = generate_stripe_resource_lock_key(stripe_resource)
      acquire_or_refresh_resource_lock(resource_lock_key, expiration_time)
    end

    def lock_netsuite_record(ns_record)
      record_lock_key = generate_netsuite_record_lock_key(ns_record)

      # TODO we should probably refresh the lock here
      if @locked_resource_keys.include?(record_lock_key)
        log.info 'record is already locked', key: record_lock_key, ns_record: ns_record
        return
      end

      acquire_lock(record_lock_key)

      @locked_resource_keys << record_lock_key
    end

    def lock_on_user(&block)
      key = generate_user_lock_key(user)
      @expiration = acquire_lock(key)

      begin
        @user_key = key

        log.info 'lock aquired', lock: @user_key

        yield
      ensure
        # We need to be careful when cleaning up the lock key.  If we took a really long
        # time for some reason, and the lock expired, someone else may have it, and
        # it's not safe for us to remove it.  Check how much time has passed since we
        # wrote the lock key and only delete it if it hasn't expired (or we're not using
        # lock expiration)
        if @expiration > Time.now.to_f
          redis.del(@user_key)
        else
          log.warn 'not deleting key', key: @user_key, expiration: @expiration
        end

        clear_locked_resources

        @user_key = nil
      end
    end

    sig { params(key: String).returns(T::Boolean) }
    def key_expired?(key)
      redis.get(key).to_f < Time.now.to_f
    end

    private

    def generate_netsuite_record_lock_key(ns_record)
      CacheKeyService.netsuite_record_cache_key(
        user: @user,
        record: ns_record
      )
    end

    sig { params(stripe_resource: Stripe::APIResource).returns(String) }
    def generate_stripe_resource_lock_key(stripe_resource)
      resource_lock_key = stripe_resource.id

      # plans and coupons don't have unique IDs across accounts
      # we must prefix them with the classname to enforce uniqueness
      if [Stripe::Plan, Stripe::Coupon].include?(stripe_resource.class)
        resource_lock_key = "#{stripe_resource.class}-#{resource_lock_key}"
      end

      "resource-#{@user.stripe_account_id}-#{resource_lock_key}"
    end

    # TODO we should really namespace our access to redis
    def redis
      Resque.redis
    end

    sig { params(resource_lock_key: String, expiration_time: T.nilable(Integer)).void }
    def acquire_or_refresh_resource_lock(resource_lock_key, expiration_time=nil)
      if @locked_resource_keys.include?(resource_lock_key)
        log.info 'resource is already locked, refreshing', key: resource_lock_key
        redis.set(resource_lock_key, generate_expiration(expiration_time: expiration_time))
        return
      end

      acquire_lock(resource_lock_key)
      @locked_resource_keys << resource_lock_key
    end

    sig { params(key: String).returns(Float) }
    def acquire_lock(key)
      expiration = generate_expiration

      # Use the expiration as the value of the lock.
      if redis.setnx(key, expiration)
        return expiration
      end

      # Lock is being held, now check if expired
      # See "Handling Deadlocks" section on http://redis.io/commands/setnx

      if key_expired?(key)
        # If it's expired, use GETSET to update it.
        old_expiration = redis.getset(key, expiration).to_f

        # Since GETSET returns the old value of the lock, if the old expiration
        # is still in the past, we know no one else has expired the locked
        # and we now have it.
        if old_expiration < Time.now.to_f
          return expiration
        end
      end

      log.debug 'could not acquire lock', key: key
      raise Integrations::Errors::LockTimeout.new("lock #{key} not available")
    end

    # https://github.com/stripe/stripe-netsuite/issues/751
    sig { params(user: StripeForce::User).returns(String) }
    def generate_user_lock_key(user)
      base_key = self.class.generate_netsuite_account_lock_key(user)

      # NOTE usage should be 1 less than the desired amount because of base_key
      concurrent_connection_limit = user.concurrent_connections - 1

      possible_keys = [base_key]
      possible_keys += (1..concurrent_connection_limit).map do |n|
        "#{base_key}-#{n}"
      end

      possible_keys.shuffle!

      available_key = possible_keys.detect {|k| redis.get(k).nil? }

      if !available_key.nil?
        return available_key
      end

      possible_keys.detect {|k| key_expired?(k) } ||
        # pick a random key to attempt to lock on
        T.cast(possible_keys.sample, String)
    end

    sig { params(expiration_time: T.nilable(Integer)).returns(Float) }
    def generate_expiration(expiration_time: nil)
      # ensure a minimum expiration time is in place
      expiration_time ||= LOCK_EXPIRATION_TIME
      expiration_time = [expiration_time, LOCK_EXPIRATION_TIME].max

      (Time.now + expiration_time).to_f
    end
  end
end
