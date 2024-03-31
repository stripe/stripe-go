# frozen_string_literal: true
# typed: true
require 'sorbet-runtime'

class SessionsController < ApplicationController
  extend T::Sig
  include StateEncryptionAlgo

  # If you're using a strategy that POSTs during callback, you'll need to skip the authenticity token check for the callback action only.
  skip_before_action :verify_authenticity_token, only: [:create, :change_default_account_config, :delete_account_config]

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
      livemode = auth_hash["provider"] == 'stripelivemode'
      user_id = stripe_callback_handler(state, livemode)
      user = T.must(StripeForce::User.find(id: user_id))

      default_user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, is_default_account_config: true)
      if default_user
        # When we create a new user, we need to copy over the field mappings and the feature flags so they don't need to be reset
        user.field_mappings = default_user.field_mappings
        user.feature_flags = default_user.feature_flags
        user.field_defaults = default_user.field_defaults
        user.save
      end

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

  # We switch the is_default_account_config flag from the current (default) account to another provided account
  def change_default_account_config
    begin
      state = require_state
      account_configurations = StripeForce::User.where(salesforce_account_id: state.salesforce_account_id)
      livemode = params.require(:livemode) == 'live'

      salesforce_account_id = state.salesforce_account_id
      stripe_account_id = params.require(:stripe_account_id)
      new_default_user = T.must(StripeForce::User.find(salesforce_account_id: salesforce_account_id, stripe_account_id: stripe_account_id, livemode: livemode))

      current_default_stripe_account = account_configurations.detect {|config| config.is_default_account_config == true }
      current_default_stripe_account.is_default_account_config = false
      new_default_user.is_default_account_config = true

      current_default_stripe_account.save
      new_default_user.save
    rescue TypeError
      multi_account_failure_response('change_default_account_config', 'The provided Stripe Account ID is not associated with a User')
    rescue StateException => e
      multi_account_failure_response('passed_state_invalid', 'The provided State value is not valid. Error: ' + e.message)
    end
  end

  # We delete the account associated with the stripe account id given.
  def delete_account_config
    begin
      state = require_state
      current_user = T.must(StripeForce::User.find(id: state.user_id))

      livemode = params.require(:livemode) == 'live'
      stripe_account_id = params.require(:stripe_account_id)

      associated_users = StripeForce::User.where(salesforce_account_id: state.salesforce_account_id)
      user_to_be_removed = T.must(StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, stripe_account_id: stripe_account_id, livemode: livemode))

      if current_user.id == user_to_be_removed.id && associated_users.count > 1
        multi_account_failure_response('delete_account_config', "The current User can only be deleted if it is the only User in the org.")
      elsif user_to_be_removed.is_default_account_config == true && associated_users.count > 1
        multi_account_failure_response('delete_account_config', "Cannot delete the default User unless it is the only User in the org.")
      else
        # Keep this consistent with v2, where we soft delete instead of hard deleting
        user_to_be_removed.enabled = false
        user_to_be_removed.connector_settings[CONNECTOR_SETTING_POLLING_ENABLED] = false
        user_to_be_removed.save
      end
    rescue TypeError => _
      multi_account_failure_response('change_default_account', 'The provided Stripe Account ID is not associated with a User.')
    rescue StateException => e
      multi_account_failure_response('passed_state_invalid', 'The provided State value is not valid. Error: ' + e.message)
    end
  end

  # We get the accounts associated with the stripe account id given.
  def get_all_account_configs
    begin
      state = require_state
      # Get all users associated with the current Salesforce account id
      associated_users = StripeForce::User.where(salesforce_account_id: state.salesforce_account_id)

      if associated_users.empty?
        multi_account_failure_response('get_all_account_configs', 'No users are associated with the provided Salesforce Account ID')
      else
        render json: associated_users, status: :ok
      end
    rescue StateException => e
      multi_account_failure_response('passed_state_invalid', 'The provided State value is not valid. Error: ' + e.message)
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
      log.warn 'salesforce_callback_handler requested with empty sf_account_id',
               sf_account_id: sf_account_id
      head :not_found
      return
    end

    state.salesforce_account_id = sf_account_id if state.salesforce_account_id.nil?

    user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, is_default_account_config: true)
    log.default_tags[:sf_account_id] = sf_account_id

    unless user
      log.info 'creating new user'
      user = StripeForce::User.new(salesforce_account_id: sf_account_id, stripe_account_id: state.stripe_account_id, livemode: state.livemode)

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

    # TODO it seems possible for a user to auth against the wrong account, need to investigate further
    if !user.new? && user.salesforce_account_id != sf_account_id
      # this likely means a user is trying to auth to Salesforce Org A but most recently logged into Salesforce Org B
      raise "User already exists and Salesforce Account ID is not equal. This should never happen."
    end

    user.salesforce_account_id = sf_account_id
    user.salesforce_refresh_token = sf_refresh_token
    user.salesforce_instance_url = sf_instance_url
    user.salesforce_token = sf_token
    user.salesforce_instance_type = determine_instance_type(sf_instance_url, state)
    user.salesforce_namespace = state.salesforce_namespace
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
      log.warn 'stripe_callback_handler requested with empty stripe_account_id',
               stripe_account_id: stripe_account_id
      head :not_found
      return
    end

    log.default_tags[:stripe_account_id] = stripe_account_id

    # set values on the state
    # it's important we do this before the below since we use state.stripe_account_id below to create the user
    if stripe_auth['extra'] && stripe_auth['extra']['settings'] && stripe_auth['extra']['settings']['dashboard'] && stripe_auth['extra']['settings']['dashboard']['display_name']
      state.stripe_account_name = stripe_auth['extra']['settings']['dashboard']['display_name']
    end
    state.stripe_account_id = stripe_account_id if state.stripe_account_id.nil?
    state.primary_stripe_account_id = stripe_account_id if state.primary_stripe_account_id.nil?
    # this field is used in stripe_v2_callback_response(state) to postmessage
    state.livemode = livemode if state.stripe_account_livemode?

    # Either the user is:
    # -- PATH 1: making the first call to auth a Stripe acct in a new Salesforce org
    # -- PATH 2: authorizing a new user (MSA)
    # -- PATH 3: reauthing an existing user (pre MSA)
    # Logically, Path 1 & 3 are the same since all we do is update the stripe id on the existing user.
    default_user = StripeForce::User.find(id: state.user_id)
    if default_user.nil?
      # If state can't pick up the default user, ie. we are in PATH 2 or 3
      log.info 'could not find an existing user based off user_id so finding default user', salesforce_account_id: state.salesforce_account_id
      default_user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, is_default_account_config: true)
    end

    if default_user.nil?
      log.info "default user not picked up from state"
      return
    end

    if default_user.feature_enabled?(FeatureFlags::MULTI_STRIPE_ACCOUNT) && default_user.stripe_account_id != stripe_account_id
      # PATH 2
      log.info 'creating a new user', stripe_account_id: stripe_account_id, salesforce_account_id: state.salesforce_account_id, is_default_account_config: default_user.nil?
      new_user = StripeForce::User.new(salesforce_account_id: state.salesforce_account_id, stripe_account_id: stripe_account_id, livemode: livemode, is_default_account_config: default_user.nil?)
      log.info 'created a new user', id: new_user.id

      # created a new Stripe acct that is not the default_account_config
      if default_user.present? && new_user.stripe_account_id != default_user.stripe_account_id
        # When we create a new user, we need to copy over the field mappings and the feature flags so they don't need to be reset
        log.info "updating new user (#{new_user.id}) account config to match default user (#{default_user.id})"
        new_user.connector_settings = default_user.connector_settings
        new_user.feature_flags = default_user.feature_flags
        new_user.field_defaults = default_user.field_defaults
        new_user.field_mappings = default_user.field_mappings
        new_user.salesforce_object_prefix_mappings = default_user.salesforce_object_prefix_mappings
        new_user.salesforce_token = default_user.salesforce_token
        new_user.salesforce_refresh_token = default_user.salesforce_refresh_token
        new_user.salesforce_organization_key = default_user.salesforce_organization_key
        new_user.save
      end
      return new_user.id
    end

    if default_user.stripe_account_id && default_user.stripe_account_id != stripe_account_id
      Integrations::ErrorContext.report_edge_case("Stripe account ID already set, overwriting.")
    end

    # PATH 1 or 3
    log.info 'updating existing (default) user stripe account ID', user_id: default_user.id, stripe_account_id: stripe_account_id
    default_user.update(stripe_account_id: stripe_account_id, livemode: livemode)
    default_user.id
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
    window.opener.postMessage("connection_successful-stripe|#{state.stripe_account_id}|#{state.stripe_account_livemode}-#{state}", "#{postmessage_domain}")
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

  def multi_account_failure_response(operation, message)
    render json: {error: "Operation failed", operation: operation, message: message}, status: 400
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
    encrypted_state = nil
    if request.headers['X-Middleware-Authorization'].present?
      encrypted_state = request.headers['X-Middleware-Authorization']
    else
      encrypted_state = params.permit(:state)["state"].to_s
    end
    # puts encrypted_state unless Rails.env.production?
    state = restore_state(encrypted_state, defaults)
    if state.nil?
      raise StateException.new(encrypted_state)
    elsif state.user_id.nil?
      # Find user after first Stripe account setup
      user = StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, stripe_account_id: state.primary_stripe_account_id, livemode: state.primary_livemode, is_default_account_config: true)
      state.user_id = user.id unless user.nil?
      state.csac_id = user.id unless user.nil?
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
