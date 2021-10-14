# frozen_string_literal: true
# typed: false
Rails.application.config.middleware.use OmniAuth::Builder do
  provider :google_oauth2, ENV["GOOGLE_OAUTH_CLIENT"], ENV["GOOGLE_OAUTH_SECRET"]
end