# frozen_string_literal: true
# typed: true

require_relative 'log'

module Integrations
  module ErrorContext
    extend T::Sig
    include Kernel

    include Log

    def report_exception(exception)
      Sentry.capture_exception(exception)
    end

    def report_edge_case(message, stripe_resource: nil, integration_record: nil, metadata: nil)
      report_error(Integrations::Errors::UnhandledEdgeCase, message, stripe_resource: stripe_resource, integration_record: integration_record, metadata: metadata)
    end

    def report_feature_usage(message, stripe_resource: nil, integration_record: nil, metadata: nil)
      report_error(Integrations::Errors::FeatureUsage, message, stripe_resource: stripe_resource, integration_record: integration_record, metadata: metadata)
    end

    sig { params(error_class: T.class_of(Integrations::Errors::BaseIntegrationError), message: String, stripe_resource: T.nilable(Stripe::StripeObject), integration_record: T.nilable(T.untyped), metadata: T.nilable(Hash)).returns(T.nilable(T.any(Sentry::Event, T::Boolean))) }
    def report_error(error_class, message, stripe_resource: nil, integration_record: nil, metadata: nil)
      sentry_options = {tags: metadata&.delete(:tags)}.compact

      log.error message, {stripe_resource: stripe_resource, integration_record: integration_record}.merge(metadata || {})

      exception = error_class.new(
        message,

        # TODO add additional context
        # stripe_resource: stripe_resource,
        # integration_record: integration_record,
        # metadata: metadata
      )

      # stacktrace needs to be set on the exception to appear in sentry
      # https://jira.corp.stripe.com/browse/REPROD-60
      exception.set_backtrace(caller)

      Sentry.capture_exception(exception, sentry_options)
    end

    # TODO add an `append: ` option to avoid wiping out existing context
    sig { params(user: T.nilable(StripeForce::User), stripe_resource: T.nilable(Stripe::APIResource), integration_record: T.untyped, tags: T.untyped).void }
    def set_error_context(user: nil, stripe_resource: nil, integration_record: nil, **tags)
      # clear out all of the context
      Sentry.set_extras({})
      Sentry.set_user({})
      Sentry.set_tags({})

      # Sentry allows you to filter on some user context fields and all tags
      tags_context = {}

      # Sentry does not allow you to filter on `extra` data
      extra_context = {}

      user_context = {}

      if user
        user_context = {
          id: user.stripe_account_id,
          email: user.email,
          livemode: user.livemode,
          # sandbox: user.sandbox?,
          username: user.name,
        }

        tags_context['production'] = user.in_production?

        extra_context.merge!({
          'netsuite_email' => user.email,
          # 'netsuite_sandbox' => user.sandbox?,
        })
      end

      # useful for the dashboard, add the admin user to all logging and errors
      if @admin
        user_context[:admin_id] = @admin.id
      end

      if stripe_resource
        if stripe_resource.is_a?(Stripe::Event)
          extra_context['stripe_event_type'] = stripe_resource.type
          extra_context['stripe_event_id'] = stripe_resource.id

          tags_context['stripe_resource_id'] = (begin
                                                  stripe_resource.data.object.id
                                                rescue
                                                  nil
                                                end)
          tags_context['stripe_resource_type'] = (begin
                                                    stripe_resource.data.object.class.to_s
                                                  rescue
                                                    nil
                                                  end)
        else
          tags_context['stripe_resource_type'] = stripe_resource.class.to_s
          tags_context['stripe_resource_id'] = stripe_resource.id
        end
      end

      if integration_record
        extra_context['integration_record_type'] = integration_record.class.to_s
        extra_context['integration_record_id'] = integration_record.internal_id

        tags_context['integration_record_id'] = integration_record.internal_id
      end

      if !tags.empty?
        extra_context.merge!(tags)
      end

      Sentry.set_user(user_context)
      Sentry.set_tags(tags_context)
      Sentry.set_extras(extra_context)

      log.set_context({
        stripe_account_id: user&.stripe_account_id,
        salesforce_account_id: user&.salesforce_account_id,
        livemode: user&.livemode,

        stripe_resource_id: stripe_resource&.id,
        stripe_resource_type: stripe_resource&.class,

        integration_record_type: integration_record&.sobject_type,
        integration_record_id: integration_record&.Id,
      }.compact)

      # if `set_error_context` is run from a class method, we want to display the class name
      log.default_tags[:job] = if self.instance_of?(Class)
        self.to_s
      else
        self.class.to_s
      end

      if !tags.empty?
        log.default_tags.merge!(tags)
      end

      env_log_level = ENV['LOG_LEVEL']

      if !env_log_level.nil? && Logger::Severity.const_defined?(env_log_level)
        log.level(Logger::Severity.const_get(env_log_level))
      elsif user&.sandbox? && !user.feature_enabled?(:loud_sandbox_logging)
        log.level(Logger::Severity::WARN)
      else
        log.level(Logger::Severity::INFO)
      end
    end
  end
end
