# https://gist.github.com/cyberfox/44de7f5f70b5e8916425
require_relative 'config'
require 'sinatra'

configure do
  enable :sessions
end

helpers do
  def h(text)
    Rack::Utils.escape_html(text)
  end
end

use OmniAuth::Builder do
  provider :salesforce, ENV['SF_CONSUMER_KEY'], ENV['SF_CONSUMER_SECRET']
  provider :stripe, ENV["STRIPE_CLIENT_ID"], ENV["STRIPE_CLIENT_SECRET"], scope: 'read_write'
end

set :port, ENV.fetch("PORT")

OmniAuth.config.full_host = ENV.fetch("DOMAIN")

get '/' do
  <<-EOL
<p>StripeForce.</p>
<p><a href="/auth/salesforce">Connect.</a></p>
EOL
end

get '/auth/stripe/callback' do
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

get '/auth/salesforce/callback' do
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

  erb <<-EOL
<p>SalesForce connected. Let's login to Stripe.</p>
<p><a href="/auth/stripe">Connect to Stripe</a>
EOL
end

get '/auth/failure' do
  params[:message]
end