# frozen_string_literal: true
source 'https://rubygems.org'
git_source(:github) {|repo| "https://github.com/#{repo}.git" }

ruby '2.6.8'

gem 'sorbet', '~> 0.5.9363', group: :development
# gem 'sorbet-runtime', '~> 0.5.9318', require: true
gem 'sorbet-rails', '~> 0.7.27'

gem 'dotenv-rails', '2.7.6', groups: [:development, :test]
gem 'rails', '~> 6.1.4', '>= 6.1.4.1'
gem 'lograge'

group :production do
  gem 'puma', '~> 5.0'
end

# sentry
gem "sentry-ruby", "~> 4.8.0"
gem "sentry-rails", "~> 4.8.0"
gem "sentry-resque", "~> 4.8.0"
gem 'simple_structured_logger'

# resque
gem 'resque', '~> 2.2.0'
gem 'resque-scheduler', "~> 4.5.0"
gem 'resque-retry', '~> 1.7.6'
gem 'resque-heroku-signals', '~> 2.2.0'
gem 'redis', '~> 4.5.0'

# database
gem 'pg', '~> 1.2.3'
gem 'sequel', '5.51.0'

# auth
gem 'omniauth-salesforce'
gem 'omniauth-stripe'
gem 'rack-attack', '~> 6.5.0'

# translation
gem 'restforce', '~> 5.2.0'
gem 'stripe', '~> 5.41.0'
gem 'rest-client', '~> 2.1.0'

# Reduces boot times through caching; required in config/boot.rb
gem 'bootsnap', '>= 1.9.1', require: false

# TODO: Update omniauth above 2.X
# Can't upgrade because of https://github.com/realdoug/omniauth-salesforce/issues/31
# gem 'omniauth', '~> 2.0.4'

# CVE-2019-13117 https://github.com/sparklemotion/nokogiri/issues/1943
gem 'nokogiri', '>= 1.12.5'

group :test do
  gem 'bundler-audit', '~> 0.9.0.1', require: false
  gem 'brakeman', '~> 5.1.2', require: false

  gem 'minitest', '~> 5.14.4'
  gem 'minitest-ci', '~> 3.4.0'
  gem 'minitest-profile'
  gem 'minitest-reporters', '~> 1.4.2'
  gem 'minitest-rails', '~> 6.1.0'

  # feature test
  gem 'capybara', '~> 3.36.0'
  gem 'webdrivers', '~> 4.7.0'
  gem 'selenium-webdriver', '~> 3.142.7'
  gem 'capybara-screenshot', '~> 1.0.25'

  gem 'mocha', '~> 1.9'
  gem 'rack-test', '~> 1.1.0'
  gem 'database_cleaner', '~> 1.7.0'
end

group :test, :development do
  gem 'pry', '~> 0.14.1'
  gem 'pry-stack_explorer', '~> 0.6.1'
  gem 'pry-nav', '~> 1.0.0'
  gem 'pry-rails', '~> 0.3.9'
  gem 'pry-doc', '~> 1.2.0'
  # https://github.com/SudhagarS/pry-state
  gem 'binding_of_caller', '~> 1.0.0'
end

group :development do
  gem 'pry-rescue', '~> 1.5.2'
  gem 'pry-remote', '~> 0.1.8'

  gem 'better_errors', '~> 2.9.1'

  gem 'listen'
  gem 'spring'
  gem 'rubocop-daemon', require: false

  # lock to an old version to align with pay-server
  gem 'rubocop', '0.89.1'
  gem 'rubocop-minitest', require: false
end
