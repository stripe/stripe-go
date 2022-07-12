# frozen_string_literal: true
source 'https://rubygems.org'
git_source(:github) {|repo| "https://github.com/#{repo}.git" }

ruby '2.7.6'

gem 'sorbet', '~> 0.5.10151', group: :development
gem 'sorbet-runtime', '~> 0.5.10151', require: true
gem 'sorbet-rails', '~> 0.7.33'

# https://github.com/ruby/irb/issues/43
gem 'reline', '~> 0.3.1'

gem 'dotenv-rails', '2.7.6', groups: [:development, :test]
gem 'rails', '~> 6.1.6'
gem 'lograge', '~> 0.12'

group :production do
  gem 'puma', '~> 5.6.4'
end

# sentry
gem "sentry-ruby", "~> 5.3.1"
gem "sentry-rails", "~> 5.3.1"
gem "sentry-resque", "~> 5.3.1"
gem 'simple_structured_logger', '~> 1.0.2'

# resque
gem 'resque', '~> 2.2.0'
gem 'resque-scheduler', "~> 4.5.0"
gem 'resque-retry', '~> 1.7.6'
gem 'resque-heroku-signals', '~> 2.2.0'
gem 'redis', '~> 4.5.1'

# database
gem 'pg', '~> 1.3.5'
gem 'sequel', '5.58.0'

# auth
# TODO hack to get around https://github.com/realdoug/omniauth-salesforce/issues/31
gem 'omniauth-rails_csrf_protection', '~> 1.0.1'
gem 'omniauth-salesforce', github: 'accel-com/omniauth-salesforce'
gem 'omniauth-stripe'
gem 'rack-attack', '~> 6.6.1'

# translation
gem 'restforce', '~> 5.3.0'
gem 'stripe', '~> 6.5.0'
gem 'rest-client', '~> 2.1.0'

# Reduces boot times through caching; required in config/boot.rb
gem 'bootsnap', '>= 1.9.1', require: false

gem 'omniauth', '~> 2.1.0'

# CVE-2019-13117 https://github.com/sparklemotion/nokogiri/issues/1943
gem 'nokogiri', '>= 1.13.0'

# https://groups.google.com/g/rubyonrails-security/c/ce9PhUANQ6s
gem 'rails-html-sanitizer', '~> 1.4.3'

group :test do
  gem 'bundler-audit', '~> 0.9.1', require: false
  gem 'brakeman', '~> 5.2.3', require: false

  gem 'minitest', '~> 5.16.2'
  gem 'minitest-ci', '~> 3.4.0'
  gem 'minitest-profile'
  gem 'minitest-reporters', '~> 1.5.0'
  gem 'minitest-rails', '~> 6.1.0'

  # feature test
  gem 'capybara', '~> 3.37.1'
  gem 'webdrivers', '~> 5.0.0'
  gem 'selenium-webdriver', '~> 4.3.0'
  gem 'capybara-screenshot', '~> 1.0.26'

  gem 'mocha', '~> 1.14'
  gem 'rack-test', '~> 2.0.2'
  gem 'database_cleaner', '~> 1.7.0'
end

group :test, :development do
  gem 'pry', '~> 0.14.1'
  gem 'pry-stack_explorer', '~> 0.6.1'
  gem 'pry-nav', '~> 1.0.0'
  gem 'pry-rails', '~> 0.3.9'
  gem 'pry-doc', '~> 1.3.0'
  # https://github.com/SudhagarS/pry-state
  gem 'binding_of_caller', '~> 1.0.0'
end

group :development do
  gem 'pry-rescue', '~> 1.5.2'
  gem 'pry-remote', '~> 0.1.8'

  gem 'better_errors', '~> 2.9.1'

  gem 'listen'
  gem 'spring'

  # lock to an old version to align with pay-server
  gem 'rubocop', '0.89.1'
  gem 'rubocop-daemon', require: false
  gem 'rubocop-minitest', require: false
end
