# frozen_string_literal: true
# typed: true

require 'resque/server'
require 'resque/scheduler/server'

Rails.application.routes.draw do
  T.bind(self, ActionDispatch::Routing::Mapper)

  default_url_options protocol: :https

  root to: 'sessions#root_action'
  get '/auth/:oauth_type', to: 'sessions#login_entrypoint', as: :omniauth
  get '/auth/salesforce/callback', to: 'sessions#salesforce_callback'
  get '/auth/salesforcesandbox/callback', to: 'sessions#salesforce_callback'
  get '/auth/stripe/callback', to: 'sessions#stripe_callback'

  post '/stripe-webhooks' => 'stripe_webhook#stripe_webhook'

  namespace :v1, module: 'api', as: 'api', constraints: {format: 'json'} do
    resource :configuration, only: [:show, :update]
    post 'post-install' => 'configurations#post_install'
    post 'translate' => 'configurations#translate'
    post 'translate_all' => 'configurations#translate_all'
  end

  scope :monitoring do
    if Rails.env.production?
      Resque::Server.use(Rack::Auth::Basic) do |_user, password|
        password == ENV.fetch("RESQUE_MONITOR_PASSWORD")
      end
    end

    mount Resque::Server.new => "/resque-monitor", as: :resque_web
  end
end
