# typed: true

class ActiveSupport::TestCase
  # our code is not thread safe, so we have to ensure tests are not run in multiple threads
  parallelize(workers: 1)
end
