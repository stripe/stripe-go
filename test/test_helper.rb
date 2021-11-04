# frozen_string_literal: true
# typed: strict

ENV['RAILS_ENV'] ||= 'test'
require_relative "../config/environment"
require "rails/test_help"

# more information on how the default test runner operates:
# https://github.com/rails/rails/blob/6-1-stable/railties/lib/rails/test_unit/testing.rake
# https://github.com/rails/rails/blob/6-1-stable/railties/lib/rails/test_unit/runner.rb

require "minitest/autorun"
require "minitest/rails"
require 'minitest/spec'
require 'minitest/profile'
require "mocha/minitest"
require 'pry-rescue/minitest' unless ENV['CI'] || ENV['EXT_DIR']

Dir[File.join(File.dirname(__FILE__), "test/support/**/*.rb")].sort.each {|f| require f }

Minitest::Ci.clean = false if ENV['CI']

class ActiveSupport::TestCase
  # Run tests in parallel with specified workers
  parallelize(workers: :number_of_processors)
end
