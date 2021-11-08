# ActiveSupport::TestCase gives us the ability to run tests by line number
class Critic::UnitTest < ActiveSupport::TestCase
  include CommonHelpers

  def setup
    common_setup
  end

  def teardown
    common_teardown
  end
end
