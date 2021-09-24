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
    puts "Error! #{e}"
  end

  sleep(30)
end