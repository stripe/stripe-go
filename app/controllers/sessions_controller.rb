# frozen_string_literal: true
# typed: ignore
class SessionsController < ApplicationController
  # If you're using a strategy that POSTs during callback, you'll need to skip the authenticity token check for the callback action only.
  skip_before_action :verify_authenticity_token, only: :create

  def create
    # @user = User.find_or_create_from_auth_hash(auth_hash)
    # self.current_user = @user

    # <<~EOL
    #   <p>StripeForce.</p>
    #   <p><a href="/auth/salesforce">Connect.</a></p>
    # EOL

    redirect_to '/'
  end

  def stripe_callback
    # TODO check for missing user ID
    user = StripeForce::User[session[:user_id]]
    stripe_auth = env['omniauth.auth']

    user.update(
      stripe_account_id: stripe_auth["uid"],
      # stripe_refresh_token: stripe_auth[""]
    )

    erb <<-EOL
    Great, your orders will now be synced
    EOL
  end

  def salesforce_callback
    sf_auth = env['omniauth.auth']
    sf_account_id = sf_auth["uid"]

    user = StripeForce::User.find(salesforce_account_id: sf_account_id)

    if !user
      user = StripeForce::User.new(salesforce_account_id: sf_account_id)
    end

    sf_credentials = sf_auth["credentials"]
    sf_refresh_token = sf_credentials["refresh_token"]
    sf_instance_url = sf_credentials["instance_url"]
    sf_token = sf_credentials["token"]

    user.salesforce_account_id = sf_account_id,
                                 user.salesforce_refresh_token = sf_refresh_token,
                                 user.salesforce_instance_url = sf_instance_url,
                                 user.salesforce_token = sf_token,

                                 user.name = sf_auth["extra"]["display_name"],
                                 user.email = sf_auth["extra"]["email"]

    user.save

    session[:user_id] = user.id

    # puts env['omniauth.auth'].inspect

    erb <<~EOL
      <p>SalesForce connected. Let's login to Stripe.</p>
      <p><a href="/auth/stripe">Connect to Stripe</a>
    EOL
  end

  def auth_failure
    params[:message]
  end

  protected def auth_hash
    request.env['omniauth.auth']
  end
end
