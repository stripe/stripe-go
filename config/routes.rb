# frozen_string_literal: true
# typed: true

require 'resque/server'
require 'resque/scheduler/server'

Rails.application.routes.draw do
  T.bind(self, ActionDispatch::Routing::Mapper)

  default_url_options protocol: :https

  root to: 'sessions#root_action'
  get '/auth/salesforce/callback', to: 'sessions#salesforce_callback'
  get '/auth/stripe/callback', to: 'sessions#stripe_callback'

  namespace :v1, module: 'api', as: 'api', constraints: {format: 'json'} do
    resource :configuration, only: [:show, :update]
    post 'post-install' => 'configurations#post_install'
    post 'translate' => 'configurations#translate'
  end

  # TODO need basic auth
  # authenticated :admin, ->(u) { u.root? } do
  #   mount Resque::Server.new => "/resque-monitor", as: :resque_web
  # end
end
