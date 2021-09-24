# https://gist.github.com/cyberfox/44de7f5f70b5e8916425

require 'dotenv/load'

require 'sinatra'
require 'omniauth-salesforce'

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
end

set :port, ENV.fetch("PORT")

OmniAuth.config.full_host = ENV.fetch("DOMAIN")

get '/' do
  <<-EOL
<p>StripeForce.</p>
<p><a href="/login">Connect.</a></p>
EOL
end

get '/login' do
  redirect to('/auth/salesforce')
end

get '/auth/:provider/callback' do
  puts env['omniauth.auth'].inspect
  session[:user] = env['omniauth.auth']['raw_info']
  session[:uid] = env['omniauth.auth']['uid']
  session[:logged_in] = true

  erb <<-EOHTML
<!DOCTYPE html>
<html>
  <head>
    <title>Welcome!</title>
  </head>
  <body>
    <h1>Hello <%= h env['omniauth.auth']['info']['name'] %>!</h1>
    <br/>
    <%= h env['omniauth.auth'].inspect %>
    <br/>
    <br/>
    From provider: #{env['omniauth.auth']['provider']} vs. #{params[:provider]}
  </body>
</html>
EOHTML
end

get '/auth/failure' do
  params[:message]
end