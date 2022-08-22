# frozen_string_literal: true
# typed: true

require_relative './stripe_factory'

# ActiveSupport::TestCase gives us the ability to run tests by line number
class Critic::UnitTest < ActiveSupport::TestCase
  include Critic::CommonHelpers
  include Critic::StripeFactory
  include Critic::SalesforceFactory

  def setup
    common_setup
  end

  def teardown
    common_teardown
  end
end
