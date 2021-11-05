# frozen_string_literal: true
# typed: strict

ENV['RAILS_ENV'] ||= 'test'
require_relative "../config/environment"
require "rails/test_help"

# more information on how the default test runner operates:
# https://github.com/rails/rails/blob/6-1-stable/railties/lib/rails/test_unit/testing.rake
# https://github.com/rails/rails/blob/6-1-stable/railties/lib/rails/test_unit/runner.rb

module Critic

end

require "minitest/autorun"
require "minitest/rails"
require 'minitest/spec'
require 'minitest/profile'
require "mocha/minitest"
require 'pry-rescue/minitest' unless ENV['CI'] || ENV['EXT_DIR']

require_relative 'support/common_helpers'
Dir[File.join(File.dirname(__FILE__), "support/**/*.rb")].sort.each {|f| require f }

Minitest::Ci.clean = false if ENV['CI']
