# typed: true
# frozen_string_literal: true

# ActiveSupport::TestCase gives us the ability to run tests by line number
class Critic::FunctionalTest < ActiveSupport::TestCase
  include CommonHelpers

  def setup
    common_setup
  end

  def teardown
    common_teardown
  end
end
