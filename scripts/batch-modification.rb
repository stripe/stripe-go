StripeForce::User.all.each do |user|
  if user.connector_settings['filters']['Order'].nil?
    puts "Updating order filter"
    user.connector_settings['filters']['Order'] = "Status = 'Activated'"
    user.save(columns: [:connector_settings])
  else
    puts "Skipping order filter:\t#{user.id}"
  end
end