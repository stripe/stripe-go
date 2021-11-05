# typed: true

class FunctionalTest < MiniTest::Spec
  include CommonHelpers
end

# class UnitTest < MiniTest::Spec
class UnitTest < ActiveSupport::TestCase
  include CommonHelpers
end
