# typed: strict
require_relative 'config'

include SimpleStructuredLogger

loop do
  log.info 'initiating poll'

  begin
    StripeForce::User.each do |user|
      log.info 'polling user'
      StripeForce::OrderPoller.perform(user: user)
    end
  rescue => e
    puts "Worker Error"
    puts e.message
    puts e.backtrace
  end

  sleep(30)
end