# frozen_string_literal: true
# typed: true
# Usage: heroku run bundle exec ruby scripts/poll-timestamps.rb acct_

def user_from_script_argument
  user_id = ARGV[0]

  if user_id.nil?
    raise 'valid user id required'
  end

  if StripeForce::User.where(stripe_user_id: user_id).count > 1
    if ARGV[1] != 'true' && ARGV[1] != 'false'
      raise 'multiple modes for specified user, but livemode flag not defined'
    end

    mode = ARGV[1] == 'true'

    user = StripeSuite::User.find(stripe_user_id: user_id, livemode: mode)
  else
    user = StripeSuite::User.find(stripe_user_id: user_id)
  end

  if user.nil?
    raise "couldn't find user with ID: #{user_id}"
  end

  user
end

def set_initial_poll_timestamp(user, sf_class)
  poll_timestamp = StripeForce::PollTimestamp.by_user_and_record(
    user,
    sf_class
  )

  if poll_timestamp.present?
    puts "poll timestamp already existed"
    return
  end

  poll_timestamp = StripeForce::PollTimestamp.build_with_user_and_record(
    user,
    sf_class
  )

  poll_timestamp.last_polled_at = Time.now
  poll_timestamp.save
end

include StripeForce::Constants

user = user_from_script_argument

set_initial_poll_timestamp(
  user,
  SF_ORDER
)