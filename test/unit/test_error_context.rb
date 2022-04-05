# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ErrorContextTest < Critic::UnitTest
  extend T::Sig

  class ErrorContextTester
    include Integrations::ErrorContext

    def initialize(user: nil, admin: nil)
      @user = user
      @admin = admin
    end
  end

  sig { returns(ErrorContextTester) }
  def error_context; ErrorContextTester.new end

  it 'sets context with a user, stripe, netsuite reference' do
    user = make_user
    user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = SFInstanceTypes::PRODUCTION.serialize
    refute(user.sandbox?)

    sf_record = create_mock_salesforce_order
    stripe_resource = Stripe::Charge.construct_from(id: create_id(:ch))
    error_context.set_error_context(user: user, custom: 'unique_tag', integration_record: sf_record, stripe_resource: stripe_resource)

    # avoid testing configuration too deeply: if custom tags are set, the rest of the configuration should be correct as well
    refute_empty(Sentry.get_current_scope.tags)

    stdout, _ = capture_subprocess_io do
      error_context.log.info 'error with tagging'
    end

    ['unique_tag', sf_record.sobject_type, stripe_resource.id, stripe_resource.class.to_s].each do |str|
      assert_match(str, stdout)
    end
  end

  describe 'log level configuration' do
    before do
      @original_level = ENV['LOG_LEVEL']
    end

    after do
      ENV['LOG_LEVEL'] = @original_level
    end

    it 'log level is set to info by default' do
      error_context.set_error_context

      assert_equal(Logger::Severity::INFO, SimpleStructuredLogger.logger.level)

      stdout, _ = capture_subprocess_io do
        error_context.log.info 'log everything'
      end

      assert_match('log everything', stdout)
    end

    it 'allows configuration of the log level via ENV' do
      ENV['LOG_LEVEL'] = 'DEBUG'
      user = make_user
      error_context.set_error_context(user: user)

      assert_equal(Logger::Severity::DEBUG, SimpleStructuredLogger.logger.level)
    end

    it 'gracefully fails when an invalid log level is specified' do
      ENV['LOG_LEVEL'] = 'WACKY'

      user = make_user(sandbox: false)
      error_context.set_error_context(user: user)

      assert_equal(Logger::Severity::INFO, SimpleStructuredLogger.logger.level)
    end

    it 'decreases log level for sandbox users to reduce log noise' do
      previous_log_level =
        ENV['LOG_LEVEL'] = nil

      user = make_user(sandbox: true)
      user.disable_feature(:loud_sandbox_logging)
      assert(user.sandbox?)

      error_context.set_error_context(user: user)

      assert_equal(Logger::Severity::WARN, SimpleStructuredLogger.logger.level)

      stdout, _ = capture_io do
        error_context.log.debug 'everything'
        error_context.log.info 'something'
      end

      refute_match('everything', stdout)
      refute_match('something', stdout)
    end
  end
end
