# typed: true
# frozen_string_literal: true

# NOTE ensure logs are pushed to heroku in real time
# https://devcenter.heroku.com/articles/logging#writing-to-your-log
$stdout.sync = true

module StripeForce

end

module Integrations

end

require_relative './hash_diff'

Dir[File.join(File.dirname(__FILE__), "integrations/**/*.rb")].sort.each {|f| require f }

require_relative 'stripe-force/resque'
require_relative 'stripe-force/translate/translate'

Dir[File.join(File.dirname(__FILE__), "stripe-force/**/*.rb")].sort.each {|f| require f }
