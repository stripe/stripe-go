require 'resque/server'
require 'resque/scheduler/server'

Rails.application.routes.draw do
  default_url_options protocol: :https

  root to: redirect('/auth/salesforce')
  post '/auth/:provider/callback', to: 'sessions#create'

  # TODO need basic auth
  # authenticated :admin, ->(u) { u.root? } do
  #   mount Resque::Server.new => "/resque-monitor", as: :resque_web
  # end
end
