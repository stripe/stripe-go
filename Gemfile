source 'https://rubygems.org/'
git_source(:github) { |repo| "https://github.com/#{repo}.git" }

ruby '2.6.8'

gem 'sorbet'
gem 'sorbet-runtime'

gem 'dotenv'
gem 'rails', '6.1.4.1'

# translator
gem 'money', '~> 6.13.8'
gem 'rest-client', '~> 2.0.2'
gem 'sentry-raven', '~> 2.13.0'
gem 'restforce'
gem 'stripe', '~> 5.38.0'
gem 'countries', '~> 3.0.1'

# https://github.com/NetSweet/netsuite/issues/363
gem 'savon', '2.11.1'

# https://github.com/resque/resque/issues/1436
# gem 'sinatra', '~> 2.0.5'

# resque
# gem 'resque', "~> 2.0.0"
# gem 'resque-scheduler', "~> 4.4.0"
# gem 'resque-retry', '~> 1.7.4'
# gem 'resque-sentry', '~> 1.2.0'
# gem 'resque-heroku-signals', '~> 2.0.0'

gem 'rake', '~> 13.0.3'

# redis
gem 'redis-namespace', '1.8.0'
gem 'redis', '~> 4.1.4'

group :test do
  gem 'bundler-audit', '~> 0.7.0.1', require: false
  gem 'brakeman', '~> 4.10', require: false

  gem 'minitest', '~> 5.14.1'
  gem 'minitest-ci', '~> 3.4.0'
  gem 'minitest-profile'
  gem 'minitest-reporters', '~> 1.4.2'
  gem 'minitest-rails', '~> 6.1.0'

  # feature test
  gem 'capybara', '~> 3.35.3'
  gem 'webdrivers', '~> 4.6.0'
  gem 'selenium-webdriver', '~> 3.142.7'
  gem 'capybara-screenshot', '~> 1.0.25'

  # Dependency of selenium, pinning to fix CVE
  gem 'rubyzip', '~> 1.3.0'

  gem 'mocha', '~> 1.9'
  gem 'rack-test', '~> 1.1.0'
  gem 'database_cleaner', '~> 1.7.0'

  # Dependencies to debug the Ruby tests
  gem 'ruby-debug-ide'
  gem 'debase'
end

group :development do
  gem 'pry', '~> 0.12.2'
  gem 'pry-rescue', '~> 1.5.2'
  gem 'pry-stack_explorer', '~> 0.4.9.3'
  gem 'pry-remote', '~> 0.1.8'
  gem 'pry-nav', '~> 0.3.0'
  gem 'pry-rails', '~> 0.3.9'
  gem 'pry-doc', '~> 1.1.0'

  gem 'derailed_benchmarks', '~> 1.8.1'

  gem 'web-console', '~> 4.1.0'
  gem 'better_errors', '~> 2.9.1'
  gem 'binding_of_caller', '~> 1.0.0'

  gem 'spring'
  gem 'listen', '>= 3.0.5', '< 3.2'
end
