# typed: true
# frozen_string_literal: true

require_relative './salesforce_debugging'

module Critic::CommonHelpers
  include Kernel
  extend T::Sig

  include Minitest::Assertions
  include StripeForce::Constants
  include SalesforceDebugging

  # NOTE this is a little dangerous: we are only doing this for `prefixed_stripe_field` right now
  include StripeForce::Utilities::SalesforceUtil

  def sf_instance_account_id
    ENV.fetch('SF_INSTANCE_ID')
  end

  def sf_caching_global_disabled
    # TODO should change to use a more specific name `DISABLE_SALESFORCE_CACHE_SERVICE`
    ENV['DISABLE_CACHE'] == 'true'
  end

  def get_sync_records_by_primary_id(primary_id)
    sync_record_results = sf.query("SELECT Id FROM #{prefixed_stripe_field(SYNC_RECORD)} WHERE #{prefixed_stripe_field(SyncRecordFields::PRIMARY_RECORD_ID.serialize)} = '#{primary_id}'")

    sync_record_results.map do |sync_record|
      sf.find(prefixed_stripe_field(SYNC_RECORD), sync_record.Id)
    end
  end

  # The get_sync_record_by_secondary_id does not make sense with the addition of Success Sync records, as that is the grouping ID.
  #  ie the translation of an Order would create multiple sync records, ie (Account, Product2, PricebookEntry as well as the Order itself).
  # This helper will fetch by the secondary id, ie the id of the actual object the sync record pertains to. Not the parent one that caused its creation.
  def get_sync_record_by_secondary_id(secondary_id)
    sync_record_results = sf.query("SELECT Id FROM #{prefixed_stripe_field(SYNC_RECORD)} WHERE #{prefixed_stripe_field(SyncRecordFields::SECONDARY_RECORD_ID.serialize)} = '#{secondary_id}'")

    assert_equal(1, sync_record_results.count)

    sync_record = sf.find(prefixed_stripe_field(SYNC_RECORD), sync_record_results.first.Id)
  end

  sig { params(user: T.nilable(StripeForce::User)).returns(StripeForce::Translate) }
  def make_translator(user: nil)
    user ||= make_user
    locker = Integrations::Locker.new(user)

    StripeForce::Translate.new(
      user,
      locker
    )
  end

  sig { params(sandbox: T::Boolean, save: T::Boolean, random_user_id: T::Boolean, livemode: T::Boolean).returns(StripeForce::User) }
  def make_user(sandbox: true, save: false, random_user_id: false, livemode: false)
    user = StripeForce::User.new(
      livemode: livemode,

      salesforce_account_id: sf_instance_account_id,
      salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
      salesforce_refresh_token: ENV['SF_REFRESH_TOKEN'],
      salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE_DOMAIN')}.my.salesforce.com",
      salesforce_organization_key: SecureRandom.alphanumeric(16),

      stripe_account_id: if random_user_id
        stripe_create_id("acct_")
      else
        ENV.fetch('STRIPE_ACCOUNT_ID')
      end
    )

    if !sf_caching_global_disabled
      log.debug 'enabling salesforce caching'
      user.enable_feature(FeatureFlags::SF_CACHING)
    end

    if user.salesforce_account_id == "00DHo0000007HgXMAU"
      # platform-integrations-eng+qapackage@stripe.com
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::QA.serialize
    elsif user.salesforce_account_id == "00D6e000000qb1uEAA"
      # platform-integrations-eng+prodpackage@stripe.com
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::PRODUCTION.serialize
    else
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = SalesforceNamespaceOptions::NONE.serialize
    end

    # MULTI_CURRENCY ORG OVERRIDES -------------------------------------------------------------------------
    # If you are adding an account here, make sure you have added 'GBP' in Setup -> Manage Currencies

    # Default to False
    user.connector_settings[CONNECTOR_SETTING_MULTICURRENCY_ENABLED] = false

    # mbianco+cpqmulticurrency@stripe.com
    if user.salesforce_account_id == "00D5f000006O9HAEA0"
      user.connector_settings[CONNECTOR_SETTING_MULTICURRENCY_ENABLED] = true
    end
    # brennen-multi-curr-scratch
    if user.salesforce_account_id == "00D52000000YKbHEAW"
      user.connector_settings[CONNECTOR_SETTING_MULTICURRENCY_ENABLED] = true
    end
    # ------------------------------------------------------------------------------------------------------

    # clocks won't be enabled in prod, so we want to mimic this
    user.disable_feature(FeatureFlags::TEST_CLOCKS)
    # disable this to prevent flooding SF with API calls while testing
    user.disable_feature(FeatureFlags::CATCH_ALL_ERRORS)

    if sandbox
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = SFInstanceTypes::SANDBOX.serialize
    else
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = SFInstanceTypes::PRODUCTION.serialize
    end

    user.save if save

    user
  end

  def stripe_create_id(prefix)
    # NOTE: The number after the underscore has significance for Stripe's internal routing.
    #   While we don't expect these IDs to be used for real API calls, we want to ensure
    #   they don't lead to unexpected behavior if they are.
    random_id = "_1" + SecureRandom.alphanumeric(29)

    if ENV['CIRCLE_NODE_INDEX']
      random_id = "#{random_id}#{ENV['CIRCLE_NODE_INDEX']}"
    end

    prefix.to_s + random_id
  end

  def sf_randomized_id
    random_id = SecureRandom.alphanumeric(29)

    if ENV['CIRCLE_NODE_INDEX']
      random_id = "#{random_id}#{ENV['CIRCLE_NODE_INDEX']}"
    end

    random_id
  end

  sig { returns(String) }
  def create_random_email
    "#{sf_randomized_id}@example.com"
  end

  def sf_randomized_name(sf_object_name)
    node_identifier = ENV['CIRCLE_NODE_INDEX'] || ""
    "REST #{sf_object_name} #{node_identifier} #{DateTime.now}"
  end

  # Helper to poll for an expected result
  sig do
    params(
      timeout: T.any(Numeric, ActiveSupport::Duration),   # Max time to attempt polling. In seconds if not defined as a Duration
      interval: T.any(Numeric, ActiveSupport::Duration),  # Time to wait before retrying. In seconds if not defined as a Duration
      block: Proc,                # Method that is attempted to start and after each interval. If a truthy value is returned, the
    )                             #     polling will stop and that value is returned.
    .void
  end
  def wait_until(timeout: 30.seconds, interval: 1.seconds, &block)
    condition_met = T.let(false, T::Boolean)

    Timeout.timeout(timeout) do
      condition_met = yield block
      until condition_met
        puts "Condition not met, waiting #{interval.seconds} seconds"
        sleep(interval.seconds)
        condition_met = yield block
      end
      condition_met
    end
  end

  def advance_test_clock(stripe_customer, advance_timestamp)
    test_clock = Stripe::TestHelpers::TestClock.retrieve(stripe_customer.test_clock, @user.stripe_credentials)
    test_clock.advance(frozen_time: advance_timestamp)

    # test clocks can take some time...
    wait_until(timeout: 5.minutes) do
      test_clock.refresh
      test_clock.status == 'ready'
    end

    test_clock
  end

  def sf
    @user.sf_client
  end

  def inline_job_processing!
    Resque.inline = true
  end

  def normal_job_processing!
    Resque.inline = false
  end

  def common_setup
    assert_equal(0, StripeForce::User.count)

    # https://github.com/resque/resque-scheduler/pull/602
    redis.redis.flushdb

    inline_job_processing!

    DatabaseCleaner.start

    KMSEncryptionTestHelpers.mock_encryption_fields(StripeForce::User)

    Integrations::Metrics::Writer.instance.timer.shutdown
    Integrations::Metrics::Writer.instance.queue.clear

    # output current test, useful for debugging which fail because of CI timeout limits
    T.bind(self, ActiveSupport::TestCase)
    puts "\n\n" + self.location
  end

  def common_teardown
    DatabaseCleaner.clean
  end

  def redis
    Resque.redis
  end

  def initial_poll_delta; 60 * 60 * 24 end

  def set_initial_poll_timestamp(sf_class)
    initial_poll = DateTime.now - initial_poll_delta

    poll_timestamp = StripeForce::PollTimestamp.build_with_user_and_record(
      @user,
      sf_class
    )

    poll_timestamp.last_polled_at = initial_poll
    poll_timestamp.save

    poll_timestamp
  end
end
