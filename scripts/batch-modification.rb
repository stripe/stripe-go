StripeForce::User.all.each do |user|
  if user.connector_settings['filters']['Order'].nil?
    puts "Updating order filter"
    user.connector_settings['filters']['Order'] = "Status = 'Activated'"
    user.save(columns: [:connector_settings])
  else
    puts "Skipping order filter:\t#{user.id}"
  end
end

StripeForce::User.where(stripe_account_id: 'acct_15uapDIsgf92XbAO').all.each do |user|
  puts "User\t#{user.salesforce_instance_url}\t#{user.id}"
  user.update(enabled: false)
end

StripeForce::User.all.each do |user|
  user.enable_feature StripeForce::Constants::FeatureFlags::CATCH_ALL_ERRORS, update: true
end

StripeForce::User.all.each do |user|
  user.enable_feature StripeForce::Constants::FeatureFlags::SF_CACHING, update: true
end