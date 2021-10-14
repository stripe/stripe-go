# NOTE ensure logs are pushed to heroku in real time
# https://devcenter.heroku.com/articles/logging#writing-to-your-log
$stdout.sync = true

module StripeForce

end

# TODO move this somewhere else
# tired of writing binding.pry...
module Kernel
  def bp
    Pry.start(binding.of_caller(1))
  end
end

Dir[File.join(File.dirname(__FILE__), "stripe-force/**/*.rb")].sort.each {|f| require f }