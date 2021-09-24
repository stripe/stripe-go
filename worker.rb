require_relative 'config'

loop do
  begin
    StripeForce::User.each do |user|
      StripeForce::OrderPoller.perform(user: user)
    end
  rescue => e
    puts "Error! #{e}"
    raise
  end

  sleep(60)
end