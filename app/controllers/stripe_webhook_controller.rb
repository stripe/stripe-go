# frozen_string_literal: true
# typed: true
class StripeWebhookController < ApplicationController
  skip_forgery_protection
  include StripeForce::Utilities::DemoUtil

  WEBHOOK_SECRET = ENV.fetch("STRIPE_WEBHOOK_SECRET")
  STRIPEFORCE_CI_ACCOUNT = "acct_1MHBTOC9fP1FVBtd"

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
      Sentry.capture_exception(e, level: 'error')

      # this should not be possible as HTTP body is parsed based on content-type
      # and will fail upstream if in valid payload is send
      render status: :bad_request, plain: "Invalid JSON in payload"
      return
    rescue => e
      Sentry.capture_exception(e, level: 'error')

      render status: :bad_request, plain: "Invalid payload"
      return
    end

    stripe_account_id = event.account
    livemode = event.livemode
    event_id = event.id
    event_type = event.type

    # ignore webhooks and don't error for StripeForce CI account
    if stripe_account_id == STRIPEFORCE_CI_ACCOUNT
      render plain: "ignoring webhook for StripeForce CI account"
      return
    end

    Integrations::Metrics::Writer.instance.track_counter('webhook.event.received', dimensions: {livemode: livemode, stripe_account_id: stripe_account_id})

    user = find_user(stripe_account_id, livemode)
    if user.nil?
      log.info "no user found for webhook"
      render status: :not_found, plain: "user not found"
      return
    end

    set_error_context(user: user, stripe_resource: event)

    if !user.enabled
      log.info "would have responded to event, but account is disabled"
      render plain: "would have responded to event, but account is disabled"
      return
    end

    if livemode != user.livemode
      log.info "user exists, but does not match event livemode", expected_livemode: livemode
      render plain: "user exists, but does not match mode"
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

    if event_type == 'customer.subscription.created' && user.feature_enabled?(FeatureFlags::BIDIRECTIONAL_SYNC_DEMO)
      StripeForce::Utilities::DemoUtil.user = user
      subscription_start_date = event.data.object['current_period_start'] ? Time.at(event.data.object['current_period_start']) : now_time
      # TODO: check to make sure that there isn't a SF order ID already in item metadata
      create_demo_evergreen_order(user: user, additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(subscription_start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
      }, additional_order_fields: {GENERIC_STRIPE_ID => event.data.object['id']})
      render plain: "Successfully processed event to create demo evergreen Salesforce order #{event_id}"
      return
    end

    if event_type != 'invoiceitem.created'
      head :ok
      return
    end

    StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(user, event)

    log.info "successfully processed event"
    Integrations::Metrics::Writer.instance.track_counter('webhook.event.processed', dimensions: {livemode: livemode, stripe_account_id: stripe_account_id})

    render plain: "Successfully processed event #{event_id}"
  end

  def find_user(stripe_account_id, livemode)
    user = nil

    unless stripe_account_id.nil?
      user = StripeForce::User.find(stripe_account_id: stripe_account_id, livemode: livemode)

      # if a user does not exist for the specified livemode, search for the user in a different livemode
      # if one exists, return 200 to webhook originator to avoid resending this webhook
      if !user
        user = StripeForce::User.find(stripe_account_id: stripe_account_id)
      end
    end
    user
  end
end
