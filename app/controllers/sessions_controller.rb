# frozen_string_literal: true
# typed: true

class SessionsController < ApplicationController
  # If you're using a strategy that POSTs during callback, you'll need to skip the authenticity token check for the callback action only.
  skip_before_action :verify_authenticity_token, only: :create

  # TODO more gracefully handle oauth redirect errors
  # rescue_from OAuth2::Error do
  # end

  def root_action
    if Rails.env.production?
      head :ok
    else
      render_oauth_post_redirect('salesforce')
    end
  end

  # LWC "authorize" button hits this action
  def login_entrypoint
    oauth_type = params.require(:oauth_type)

    # the namespace can change depending on what SF environment we are in
    # we need to postMessage to the correct domain when auth is complete
    salesforce_namespace = subdomain_namespace_from_param(params.permit(:salesforceNamespace)["salesforceNamespace"])
    session[:salesforce_namespace] = salesforce_namespace

    render_oauth_post_redirect(oauth_type)
  end

  def salesforce_callback
    sf_auth = auth_hash
    raw_sf_account_url = sf_auth["uid"]

    # first ID in the URL is the organizational ID
    # ex: "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
    sf_account_id = raw_sf_account_url.match(%r{id/([^/]+)/})[1]

    user = StripeForce::User.find(salesforce_account_id: sf_account_id)

    log.default_tags[:sf_account_id] = sf_account_id

    if !user
      log.info 'creating new user'
      user = StripeForce::User.new(salesforce_account_id: sf_account_id)

      # TODO log name of user
      report_feature_usage("new user #{sf_account_id}")
    end

    log.info 'updating existing user'

    sf_credentials = sf_auth["credentials"]
    sf_refresh_token = sf_credentials['refresh_token']
    sf_instance_url = sf_credentials["instance_url"]
    sf_token = sf_credentials["token"]

    # TODO it seems possible for a user to auth against the wrong account, need to investigate further
    if !user.new? && user.salesforce_account_id != sf_account_id
      raise "user already exists and account ID is not equal, this should never happen"
    end

    user.salesforce_account_id = sf_account_id
    user.salesforce_refresh_token = sf_refresh_token
    user.salesforce_instance_url = sf_instance_url
    user.salesforce_token = sf_token

    user.name = sf_auth["extra"]["display_name"]
    user.email = sf_auth["extra"]["email"]

    user.save

    session[:user_id] = user.id

    # TODO before redirecting to Stripe, check if there's a valid Stripe connection

    postmessage_domain = build_postmessage_domain(user, session[:salesforce_namespace])

    render inline: <<-EOL
    <%= form_tag(omniauth_path('stripe'), method: 'post', id: 'js-submission') %>

    <script>
    window.opener.postMessage("stripeConnectionSuccessful", "#{postmessage_domain}")
    document.getElementById('js-submission').submit()
    </script>
    EOL
  end

  def stripe_callback
    user_id = session[:user_id]

    if user_id.blank?
      log.warn 'callback requested with empty user_id',
        user_id: user_id
      head :not_found
      return
    end

    user = StripeForce::User[user_id]

    if user.blank?
      report_edge_case("invalid user identifier", metadata: {user_id: user_id})
      head :not_found
      return
    end

    stripe_auth = auth_hash
    stripe_user_id = stripe_auth["uid"]

    log.info 'updating stripe account ID', user_id: user.id, stripe_account_id: stripe_user_id

    if user.stripe_account_id && user.stripe_account_id != stripe_user_id
      report_edge_case("stripe account ID already set, overwriting")
    end

    user.update(
      stripe_account_id: stripe_user_id,
      # TODO maybe public key?
      # stripe_refresh_token: stripe_auth[""]
    )

    postmessage_domain = build_postmessage_domain(user, session[:salesforce_namespace])

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">Great, your connected!</h1>
      <p>Your Stripe & Salesforce accounts are connected. You can safely close this window.</p>
      <p>Navigate to Salesforce to configure this connector.</p>
    </div>

    <script type="application/javascript">
    window.opener.postMessage("salesforceConnectionSuccessful", "#{postmessage_domain}")
    </script>
    EOL
  end

  def failure
    render inline: 'Authorization Failure'
  end

  private def build_postmessage_domain(user, raw_namespace)
    salesforce_namespace = subdomain_namespace_from_param(raw_namespace)
    iframe_domain = iframe_domain_from_user(user)
    "https://#{user.sf_subdomain}--#{salesforce_namespace}.#{iframe_domain}"
  end

  private def render_oauth_post_redirect(oauth_type)
    render locals: {oauth_type: oauth_type}, inline: <<-EOL
      <%= form_tag(omniauth_path(oauth_type), method: 'post', id: 'js-submission') %>

      <script>
      document.getElementById('js-submission').submit()
      </script>
    EOL
  end

  # TODO need to handle failure flows more cleanly
  protected def auth_failure
    params[:message]
  end

  protected def auth_hash
    request.env['omniauth.auth']
  end
end
