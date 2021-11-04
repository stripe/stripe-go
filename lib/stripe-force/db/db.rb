# typed: true
# frozen_string_literal: true

is_test_environment = ENV['RAILS_ENV'] && ENV['RAILS_ENV'] == 'test' && !ENV['TEST_DATABASE_URL'].nil?

url = if is_test_environment
  ENV.fetch('TEST_DATABASE_URL')
else
  ENV.fetch('DATABASE_URL')
end

DB = Sequel.connect(url)
