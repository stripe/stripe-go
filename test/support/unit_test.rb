# frozen_string_literal: true
# typed: true

# ActiveSupport::TestCase gives us the ability to run tests by line number
class Critic::UnitTest < ActiveSupport::TestCase
  include CommonHelpers
  include Critic::StripeFactory

  def setup
    common_setup
  end

  def teardown
    common_teardown
  end
end
