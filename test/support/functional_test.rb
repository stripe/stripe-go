# typed: true

class ActiveSupport::TestCase
  # our code is not thread safe, so we have to ensure tests are not run in multiple threads
  parallelize(workers: 1)
end

# ActiveSupport::TestCase gives us the ability to run tests by line number

class Critic::FunctionalTest < ActiveSupport::TestCase
  include CommonHelpers
end

class Critic::UnitTest < ActiveSupport::TestCase
  include CommonHelpers
end
