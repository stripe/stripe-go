# frozen_string_literal: true
# typed: strict
ENV['BUNDLE_GEMFILE'] ||= File.expand_path('../Gemfile', __dir__)

# https://github.com/rails/rails/commit/a0f18e60900fc45eb3524ab3cdfe57be430d6016
ENV['DEFAULT_TEST_EXCLUDE'] = 'test/support/*'

require "bundler/setup" # Set up gems listed in the Gemfile.
require "bootsnap/setup" # Speed up boot time by caching expensive operations.
