# typed: true
# frozen_string_literal: true

require_relative './salesforce_factory'

module CommonHelpers
  include StripeForce::Constants
  include Critic::SalesforceFactory
  include Kernel

  def sf_account_id
    ENV.fetch('SF_INSTANCE_ID')
  end

  def make_user(sandbox: false, save: false, random_user_id: false)
    user = StripeForce::User.new(
      livemode: false,

      salesforce_account_id: sf_account_id,
      salesforce_token: ENV.fetch('SF_ACCESS_TOKEN'),
      salesforce_refresh_token: ENV['SF_REFRESH_TOKEN'],
      salesforce_instance_url: "https://#{ENV.fetch('SF_INSTANCE_DOMAIN')}.my.salesforce.com",

      stripe_account_id: if random_user_id
        create_id("acct_")
      else
        ENV.fetch('STRIPE_ACCOUNT_ID')
      end
    )

    # TODO major hack until we figure out what we are doing with sandboxes
    user.instance_variable_set('@sandbox', sandbox)

    user.save if save

    user
  end

  def create_id(prefix)
    # NOTE: The number after the underscore has significance for Stripe's internal routing.
    #   While we don't expect these IDs to be used for real API calls, we want to ensure
    #   they don't lead to unexpected behavior if they are.
    random_id = "_1" + SecureRandom.alphanumeric(29)

    if ENV['CIRCLE_NODE_INDEX']
      random_id = "#{random_id}#{ENV['CIRCLE_NODE_INDEX']}"
    end

    prefix.to_s + random_id
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
    # https://github.com/resque/resque-scheduler/pull/602
    redis.redis.flushdb

    inline_job_processing!

    DatabaseCleaner.start

    # KMSEncryptionTestHelpers.mock_encryption_fields(StripeSuite::User)

    Integrations::Metrics::Writer.instance.timer.shutdown
    Integrations::Metrics::Writer.instance.queue.clear

    # output current test, useful for debugging which fail because of CI timeout limits
    T.bind(self, ActiveSupport::TestCase)
    puts self.location
  end

  def common_teardown
    DatabaseCleaner.clean
  end

  def redis
    Resque.redis
  end
end
