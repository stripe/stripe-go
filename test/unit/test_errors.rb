# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ErrorTest < Critic::UnitTest
  before do
    Sentry.configuration.enabled_environments.push Rails.env
  end

  after do
    Sentry.configuration.enabled_environments.delete Rails.env
  end

  it 'formats an error message when a stripe and salesforce record is passed' do
    stripe_resource = Stripe::Charge.construct_from(id: 'ch_xyz')
    salesforce_object = create_mock_salesforce_customer

    exception = Integrations::Errors::TranslatorError.new("logic error", stripe_resource: stripe_resource, salesforce_object: salesforce_object)

    assert_includes(exception.message, "logic error stripe_resource_id=ch_xyz salesforce_object_type=Account salesforce_object_id=#{salesforce_object.Id}")
  end

  it 'all exception types can be captured' do
    assert(Sentry.capture_message('a capture message'))
    assert(Sentry.capture_exception(StandardError.new('bla bla bla')))

    assert(Sentry.capture_exception(Stripe::InvalidRequestError.new("invalid", "param")))

    assert(Sentry.capture_exception(
      Integrations::Errors::TranslatorError.new('bla bla bla',
        stripe_resource: Stripe::Invoice.construct_from(id: stripe_create_id(:id)),
        salesforce_object: create_mock_salesforce_order,
        metadata: {flag: "blah"},
      )
    ))

    assert(Sentry.capture_exception(
      Integrations::Errors::TranslatorError.new('bla bla bla',
        stripe_resource: Stripe::Invoice.construct_from(id: stripe_create_id(:id)),
        salesforce_object: create_mock_salesforce_order,
      )
    ))
  end

  it 'assigns sentry exception levels' do
    skip("need to reimplement sentry error levels once we are closer to production")

    event = Sentry.capture_exception(StandardError.new("a great error"))
    assert_equal(:error, event.level)

    event = Sentry.capture_exception(Integrations::Errors::UserError.new("a great error"))
    refute_equal(:error, event.level)

    event = Sentry.capture_exception(Integrations::Errors::ImpossibleState.new("a bad state"))
    assert_equal(:error, event.level)
  end

  it 'extracts metadata into error context' do
    event = Sentry.capture_exception(Integrations::Errors::TranslatorError.new("a great error", metadata: {foo: 'bar'}))
    assert_equal("bar", event.to_json_compatible["extra"]["foo"])
  end

  it 'generates a custom fingerprint when record context is passed' do
    skip("fingerprint customization is not pulled in yet")

    error_message = "great error"
    exception = Integrations::Errors::TranslatorError.new(error_message,
      stripe_resource: Stripe::Invoice.construct_from(id: stripe_create_id(:id)),
      salesforce_object: create_mock_salesforce_order
    )

    event = Sentry.capture_exception(exception)

    assert_equal(error_message, event.to_json_compatible['fingerprint'].last)
  end

  it 'records non-exception errors with a backtrace' do
    translator = make_translator

    event = translator.report_edge_case("hello")

    # digging deep into how sentry works here: `to_json_compatible` gives us a representation
    # which ~matches what is sent to Sentry
    refute_nil(event.to_json_compatible["exception"]["values"].first["stacktrace"])
  end

  it 'handles errors without an exception' do
    # ensure `hint[:exception] == nil` doesn't cause an error
    Sentry.capture_message("test")
  end

  it 'drops the level of a reraised resque error' do
    skip("need to reimplement sentry error levels once we are closer to production")

    exception = RuntimeError.new("Additional error (Resque::TermException: SIGTERM) failure hooks for job")

    event = Sentry.capture_exception(exception)

    assert_equal(:info, event.level)
  end
end
