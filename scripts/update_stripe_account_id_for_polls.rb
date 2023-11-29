# Note: This is ONLY being run right after running migration 004. It is to ensure that we populate our newly created stripe_account_id field

StripeForce::PollTimestamp.dataset.each do |poll_timestamp|
    user = StripeForce::User.where(salesforce_account_id: poll_timestamp.salesforce_account_id).first
    
    if user
      poll_timestamp.stripe_account_id = user.stripe_account_id
      if poll_timestamp.valid?
        poll_timestamp.save
      else
        puts "Validation failed for poll_timestamp with id #{poll_timestamp.id}"
      end
    else
      puts "No user found for salesforce_account_id #{poll_timestamp.salesforce_account_id}"
    end
  end