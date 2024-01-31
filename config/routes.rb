# frozen_string_literal: true
# typed: true

require 'resque/server'
require 'resque/scheduler/server'

Rails.application.routes.draw do
  T.bind(self, ActionDispatch::Routing::Mapper)

  default_url_options protocol: :https

  root to: 'sessions#root_action'
  get '/auth/:oauth_type', to: 'sessions#auth_entrypoint', as: :omniauth
  get '/auth/v2/:oauth_type', to: 'sessions#auth_v2_entrypoint', as: :omniauth_v2
  get '/auth/salesforce/callback', to: 'sessions#salesforce_callback'
  get '/auth/salesforcesandbox/callback', to: 'sessions#salesforce_callback'
  get '/auth/stripelivemode/callback', to: 'sessions#stripe_callback'
  get '/auth/stripetestmode/callback', to: 'sessions#stripe_callback'
  get '/accounts', to: 'sessions#get_all_account_configs'

  post '/stripe-webhooks' => 'stripe_webhook#stripe_webhook'
  post '/accounts/:stripe_account_id/:livemode/set_default', to: 'sessions#change_default_account_config', as: :set_default_config

  delete '/accounts/:stripe_account_id/:livemode', to: 'sessions#delete_account_config', as: :delete_account_config

  namespace :v1, module: 'api', as: 'api', constraints: {format: 'json'} do
    resource :configuration, only: [:show, :update]
    post 'post-install' => 'configurations#post_install'
    post 'translate' => 'configurations#translate'
    post 'translate_all' => 'configurations#translate_all'
    get 'connection_statuses' => 'configurations#connection_statuses'
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
