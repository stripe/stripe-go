# frozen_string_literal: true
# typed: true
class StripeWebhookController < ApplicationController
  skip_forgery_protection

  WEBHOOK_SECRET = ENV.fetch("STRIPE_WEBHOOK_SECRET")

  def stripe_webhook
    payload = request.body.read.freeze

    begin
      event = Stripe::Webhook.construct_event(
        payload, request.env['HTTP_STRIPE_SIGNATURE'], WEBHOOK_SECRET
      )
    rescue Stripe::SignatureVerificationError
      log.warn 'Invalid webhook signature received'
      render status: :bad_request, plain: "Invalid request"
      return
    rescue JSON::ParserError => e
      Sentry.capture_exception(e)

      # this should not be possible as HTTP body is parsed based on content-type
      # and will fail upstream if in valid payload is send
      render status: :bad_request, plain: "Invalid JSON in payload"
      return
    rescue => e
      Sentry.capture_exception(e)

      render status: :bad_request, plain: "Invalid payload"
      return
    end

    stripe_account_id = event.account
    livemode = event.livemode
    event_id = event.id

    # TODO this should be pulled out into a helper method
    unless stripe_account_id.nil?
      user = StripeForce::User.find(stripe_account_id: stripe_account_id, livemode: livemode)

      # if a user does not exist for the specified livemode, search a user in a different livemode
      # if one exists, return 200 to webhook originator to avoid resending this webhook
      if !user
        user = StripeForce::User.find(stripe_account_id: stripe_account_id)
      end
    end

    if user.nil?
      log.info "no user found for webhook"
      render status: :not_found, plain: "SuiteSync: user not found"
      return
    end

    set_error_context(user: user, stripe_resource: event)

    if !user.enabled
      log.info "would have responded to event, but account is disabled"
      render plain: "SuiteSync: would have responded to event, but SuiteSync account is disabled"
      return
    end

    if livemode != user.livemode
      log.info "user exists, but does not match event livemode", expected_livemode: livemode
      render plain: "SuiteSync: user exists, but does not match SuiteSync mode"
      return
    end

    if user.feature_enabled?(FeatureFlags::IGNORE_WEBHOOKS)
      log.info 'ignoring webhook'
      render plain: "user exists, but webhook ignored"
      return
    end

    if user.feature_enabled?(FeatureFlags::REJECT_WEBHOOKS)
      log.info 'rejecting webhook'
      render status: :service_unavailable, plain: "user exists, but webhook rejected"
      return
    end

    # TODO here's where we will do the work in the future
    # EventTranslationJob.work(event_id, user.stripe_user_id, livemode)

    log.info "successfully processed event", metric: 'event.received'

    render plain: "Successfully processed event #{event_id}"
  end
end
