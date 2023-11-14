# frozen_string_literal: true
# typed: true
require 'sorbet-runtime'

class SessionsController < ApplicationController
  extend T::Sig
  include StateEncryptionAlgo

  # If you're using a strategy that POSTs during callback, you'll need to skip the authenticity token check for the callback action only.
  skip_before_action :verify_authenticity_token, only: :create

  # TODO more gracefully handle oauth redirect errors
  # rescue_from OAuth2::Error do
  # end

  def root_action
    if Rails.env.production?
      head :ok
    else
      auth_entrypoint
    end
  end

  # LWC "authorize" button hits this action
  def auth_entrypoint
    state = nil
    begin
      state = require_state
    rescue StateException
      # only happens with old code calling us
      state = create_state(oauth_version: 'v1',)
      extract_salesforce_metadata(state)
    end
    oauth_type = params.require(:oauth_type)

    session[:state] = state.to_s
    prompt = params[:prompt]

    render_oauth_post_redirect(oauth_type, state, prompt)
  end

  def auth_v2_entrypoint
    oauth_type = params.require(:oauth_type)
    begin
      state = require_state(oauth_version: 'v2')
      session[:state] = state.to_s
      prompt = params[:prompt]
      render_oauth_post_redirect(oauth_type, state, prompt)
    rescue StateException => e
      v2_failure_response(oauth_type, e.state, "Connection State is Invalid.")
    end
  end

  def salesforce_callback
    begin
      state = require_state

      user_id = salesforce_callback_handler(state)
      # user = T.must(StripeForce::User.find(id: user_id))
      state.user_id = user_id
      # puts state.data.to_json
      if state.oauth_version == 'v2'
        sf_v2_callback_response(state)
      else
        # because we need to persist the session for the redirect to Stripe used by v1 auth flow
        session[:state] = state.to_s
        sf_v1_callback_response(state)
      end
    rescue StateException => e
      v2_failure_response("salesforce", e.state, "Connection State is Invalid.")
    end
  end

  def stripe_callback
    begin
      state = require_state
      livemode = request.path.include? '/stripelivemode/'
      user_id = stripe_callback_handler(state, livemode)
      user = T.must(StripeForce::User.find(id: user_id))

      if state.oauth_version == 'v2'
        user.cache_stripe_v2_connection_status(connected: true, livemode: livemode)
        stripe_v2_callback_response(state)
      else
        user.cache_connection_status(StripeForce::Constants::Platforms::STRIPE, true)
        # this is for the old LWC, we need to redirect to the salesforce auth page
        stripe_v1_callback_response(state)
      end
    rescue StateException => e
      v2_failure_response("salesforce", e.state, "Connection State is Invalid.")
    end
  end

  sig { params(state: StateEncryptionAlgo::StripeOAuthState).void }
  private def extract_salesforce_metadata(state)
    # the namespace can change depending on what SF environment we are in
    # we need to postMessage to the correct domain when auth is complete
    if state.salesforce_namespace.nil?
      salesforce_namespace = subdomain_namespace_from_param(params.permit(:salesforceNamespace)["salesforceNamespace"])
      state.salesforce_namespace = salesforce_namespace
    end

    if state.salesforce_instance_type.nil?
      salesforce_instance_type = salesforce_instance_type_from_headers(request.headers[SALESFORCE_INSTANCE_TYPE_HEADER])
      salesforce_instance_type = params.permit(:instanceType)["instanceType"] unless params.permit(:instanceType)["instanceType"].nil?
      salesforce_instance_type = determine_instance_type(request.referer, nil) if salesforce_instance_type.nil?

      state.salesforce_instance_type = salesforce_instance_type
    end

    if state.salesforce_instance_subdomain.nil?
      uri = request.referer.match(%r{https://([^/]+)/})[1]
      uri_host = uri.split('.')&.first
      uri_host = uri_host[0..uri_host.length - 4] if uri_host && uri_host[(uri_host.length - 3)..] == "--#{state.salesforce_namespace}"
      state.salesforce_instance_subdomain = uri_host if uri_host
    end
  end

  private def salesforce_callback_handler(state)
    sf_auth = auth_hash
    raw_sf_account_url = sf_auth["uid"]

    # first ID in the URL is the organizational ID
    # ex: "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
    sf_account_id = raw_sf_account_url.match(%r{id/([^/]+)/})[1]

    if sf_account_id.blank?
      log.warn 'callback requested with empty sf_account_id',
               sf_account_id: sf_account_id
      head :not_found
      return
    end

    state.salesforce_account_id = sf_account_id if state.salesforce_account_id.nil?

    user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id)

    log.default_tags[:sf_account_id] = sf_account_id

    unless user
      log.info 'creating new user'
      user = StripeForce::User.new(salesforce_account_id: sf_account_id)

      Integrations::Metrics::Writer.instance.track_counter('user.new', dimensions: {livemode: user.livemode, sf_account_id: sf_account_id})
      Integrations::ErrorContext.report_feature_usage("New user. sf_account_id: #{sf_account_id}")
    end

    log.default_tags[:user_id] = user.id
    state.user_id = user.id

    log.info 'updating existing user'
    sf_credentials = sf_auth["credentials"]
    sf_refresh_token = sf_credentials['refresh_token']
    sf_instance_url = sf_credentials["instance_url"]
    sf_token = sf_credentials["token"]
    # puts sf_credentials.to_json

    # TODO it seems possible for a user to auth against the wrong account, need to investigate further
    if !user.new? && user.salesforce_account_id != sf_account_id
      # this likely means a user is trying to auth to Salesforce Org A but most recently logged into Salesforce Org B
      raise "user already exists and account ID is not equal, this should never happen"
    end
    user.salesforce_account_id = sf_account_id
    user.salesforce_refresh_token = sf_refresh_token
    user.salesforce_instance_url = sf_instance_url
    user.salesforce_token = sf_token

    user.salesforce_instance_type = determine_instance_type(sf_instance_url, state)
    user.salesforce_namespace = state.salesforce_namespace

    # puts user.to_json

    user.name = sf_auth["extra"]["display_name"]
    user.email = sf_auth["extra"]["email"]

    user.save

    user.cache_connection_status(StripeForce::Constants::Platforms::SALESFORCE, true)

    user.id
  end

  private def salesforce_instance_type_from_headers(raw_header)
    SFInstanceTypes.try_deserialize(raw_header)&.serialize
  end

  private def determine_instance_type(instance_url, state)
    # example urls
    # scratch: https://momentum-customization-3160-dev-ed--c.scratch.vf.force.com/
    # sandbox: https://momentum-customization-3160-dev-ed.cs88.my.salesforce.com/
    # sandbox: https://somesite--sandboxname.sandbox.my.salesforce.com/
    # production: https://momentum-customization-3160-dev-ed.my.salesforce.com/
    # production: https://somesite.my.salesforce.com/
    return state.salesforce_instance_type unless state.nil? || state.salesforce_instance_type.nil?
    if instance_url.include? ".scratch."
      SFInstanceTypes::SCRATCH_ORG.serialize
    elsif instance_url.include? ".sandbox."
      SFInstanceTypes::SANDBOX.serialize
    else
      SFInstanceTypes::PRODUCTION.serialize
    end
  end

  # pushing the logic to creating users records on to the post-install endpoint.
  private def stripe_callback_handler(state, livemode)
    stripe_auth = auth_hash
    stripe_account_id = stripe_auth["uid"]

    if stripe_account_id.blank?
      log.warn 'callback requested with empty stripe_account_id',
               stripe_account_id: stripe_account_id
      head :not_found
      return
    end

    state.stripe_account_id = stripe_account_id if state.stripe_account_id.nil?
    state.primary_stripe_account_id = stripe_account_id if state.primary_stripe_account_id.nil?

    unless state.user_id
      return nil
    end

    user = StripeForce::User.find(id: state.user_id)

    log.default_tags[:stripe_account_id] = stripe_account_id

    if user.blank?
      Integrations::ErrorContext.report_edge_case("invalid user identifier", metadata: {user_id: state.user_id})
      head :not_found
      return
    end

    log.default_tags[:user_id] = user.id

    log.info 'updating stripe account ID', user_id: user.id, stripe_account_id: stripe_account_id

    if user.stripe_account_id && user.stripe_account_id != stripe_account_id
      Integrations::ErrorContext.report_edge_case("stripe account ID already set, overwriting")
    end

    user.update(stripe_account_id: stripe_account_id, livemode: livemode)

    user.id
  end

  def sf_v1_callback_response(state)
    # TODO before redirecting to Stripe, check if there's a valid Stripe connection
    postmessage_domain = build_postmessage_domain_from_state(state)
    is_production_org = state.salesforce_instance_type == SFInstanceTypes::PRODUCTION.serialize
    omniauth_path_name = is_production_org ? 'stripelivemode' : 'stripetestmode'
    render locals: {omniauth_path_name: omniauth_path_name, state: state}, inline: <<-EOL
    <%= form_tag(omniauth_path(omniauth_path_name), method: 'post', id: 'js-submission') do %>
    <%= hidden_field_tag(:state, state) %>
    <% end %>

    <script>
    window.opener.postMessage("stripeConnectionSuccessful", "#{postmessage_domain}")
    document.getElementById('js-submission').submit()
    </script>
    EOL
  end

  def sf_v2_callback_response(state)
    postmessage_domain = build_postmessage_domain_from_state(state)

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">Great, you're connected!</h1>
      <p>Your Salesforce account is connected. You can safely close this window.</p>
      <p>Navigate to Salesforce to authenticate with Stripe.</p>
    </div>

    <script type="application/javascript">
    window.opener.postMessage("connection_successful-salesforce|#{state.salesforce_account_id}-#{state}", "#{postmessage_domain}")
    </script>
    EOL
  end

  def stripe_v1_callback_response(state)
    postmessage_domain = build_postmessage_domain_from_state(state)

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">Great, you're connected!</h1>
      <p>Your Stripe & Salesforce accounts are connected. You can safely close this window.</p>
      <p>Navigate to Salesforce to configure this connector.</p>
    </div>

    <script type="application/javascript">
    window.opener.postMessage("stripeConnectionSuccessful", "#{postmessage_domain}")
    window.opener.postMessage("salesforceConnectionSuccessful", "#{postmessage_domain}")
    </script>
    EOL
  end

  def stripe_v2_callback_response(state)
    postmessage_domain = build_postmessage_domain_from_state(state)

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">Great, you're connected!</h1>
      <p>Your Stripe is connected. You can safely close this window.</p>
      <p>Navigate to Salesforce to configure this connector.</p>
    </div>

    <script type="application/javascript">
    window.opener.postMessage("connection_successful-stripe|#{state.stripe_account_id}-#{state}", "#{postmessage_domain}")
    </script>
    EOL
  end

  # there's scenarios where we don't have proper information to build the postmessage domain
  # So let's allow any, just for failure messages.
  def v2_failure_response(system_id, state, message)
    postmessage_domain = "*"

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">We encountered an error!</h1>
      <p>#{message}</p>
      <p>Navigate to Salesforce to try again.</p>
    </div>

    <script type="application/javascript">
    window.opener.postMessage("connection_failed-#{system_id}-#{state}-#{message}", "#{postmessage_domain}")
    </script>
    EOL
  end

  def failure
    render inline: 'Authorization Failure'
  end

  private def render_oauth_post_redirect(oauth_type, state, prompt)
    auth_path = omniauth_path(oauth_type)
    auth_path = omniauth_path(oauth_type, prompt: prompt) if prompt
    render locals: {auth_path: auth_path, state: state.to_s}, inline: <<-EOL
      <%= form_tag(auth_path, method: 'post', id: 'js-submission') do %>
      <%= hidden_field_tag(:state, state) %>
      <% end %>

      <script>
      document.getElementById('js-submission').submit()
      </script>
    EOL
  end

  class StateException < StandardError
    attr_reader :state
    def initialize(state)
      @state = state
    end
  end

  protected def is_valid_state?(state)
    state.is_a? StateEncryptionAlgo::StripeOAuthState
  end

  sig { params(defaults: Hash).returns(StateEncryptionAlgo::StripeOAuthState) }
  protected def require_state(defaults={})
    encrypted_state = params.permit(:state)["state"].to_s
    # puts encrypted_state unless Rails.env.production?
    state = restore_state(encrypted_state, defaults)
    if state.nil?
      raise StateException.new(encrypted_state)
    elsif state.user_id.nil?
      user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id)
      state.user_id = user.id unless user.nil?
    end
    state
  end

  # TODO need to handle failure flows more cleanly
  protected def auth_failure
    params[:message]
  end

  protected def auth_hash
    request.env['omniauth.auth']
  end
end
