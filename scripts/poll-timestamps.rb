def set_initial_poll_timestamp(sf_class)
    poll_timestamp = StripeForce::PollTimestamp.build_with_user_and_record(
      @user,
      sf_class
    )

    poll_timestamp.last_polled_at = Time.now
    poll_timestamp.save
  end
