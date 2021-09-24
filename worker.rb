require_relative 'config'

include SimpleStructuredLogger

loop do
  begin
    StripeForce::User.each do |user|
      log.info 'polling user'
      StripeForce::OrderPoller.perform(user: user)
    end
  rescue => e
    puts "Error! #{e}"
  end

  sleep(60)
end