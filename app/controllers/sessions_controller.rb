# frozen_string_literal: true
# typed: true

class SessionsController < ApplicationController
  # If you're using a strategy that POSTs during callback, you'll need to skip the authenticity token check for the callback action only.
  skip_before_action :verify_authenticity_token, only: :create

  # rescue_from OAuth2::Error do
  #   redirect_to '/'
  # end

  def root_action
    if Rails.env.production?
      head :ok
    else
      redirect_to '/auth/salesforce'
    end
  end

  def salesforce_callback
    sf_auth = auth_hash
    raw_sf_account_url = sf_auth["uid"]

    # first ID in the URL is the organizational ID
    # ex: "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
    sf_account_id = raw_sf_account_url.match(%r{id/([^/]+)/})[1]

    user = StripeForce::User.find(salesforce_account_id: sf_account_id)

    if !user
      log.info 'creating new user', sf_account_id: sf_account_id
      user = StripeForce::User.new(salesforce_account_id: sf_account_id)

      report_feature_usage("new user #{sf_account_id}")
    end

    sf_credentials = sf_auth["credentials"]
    sf_refresh_token = sf_credentials['refresh_token']
    sf_instance_url = sf_credentials["instance_url"]
    sf_token = sf_credentials["token"]

    user.salesforce_account_id = sf_account_id
    user.salesforce_refresh_token = sf_refresh_token
    user.salesforce_instance_url = sf_instance_url
    user.salesforce_token = sf_token

    user.name = sf_auth["extra"]["display_name"]
    user.email = sf_auth["extra"]["email"]

    user.save

    session[:user_id] = user.id

    # TODO before redirecting to Stripe, check if there's a valid Stripe connection

    redirect_to '/auth/stripe'
  end

  def stripe_callback
    user_id = session[:user_id]

    if user_id.blank?
      log.warn 'callback requested with empty user id'
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

    user.update(
      stripe_account_id: stripe_user_id,
      # TODO maybe public key?
      # stripe_refresh_token: stripe_auth[""]
    )

    render inline: <<-EOL
    <div style="text-align:center; font-family:'Helvetica Neue', Arial, sans-serif;">
      <h1 style="margin-top: 30%;">Great, your connected!</h1>
      <p>Your Stripe & SalesForce accounts are connected. You can safely close this window.</p>
      <p>Navigate to SalesForce to configure this connector.</p>
    </div>
    <script type="application/javascript">
    window.opener.postMessage("connectionSuccessful", "https://#{user.sf_subdomain}--c.visualforce.com")
    </script>
    EOL
  end

  def auth_failure
    params[:message]
  end

  protected def auth_hash
    request.env['omniauth.auth']
  end
end
