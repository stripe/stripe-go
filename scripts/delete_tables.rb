# Usage: RAILS_ENV=test be ruby scripts/delete_tables.rb

require File.expand_path('../config/environment', __dir__)

if !Rails.env.development? && !Rails.env.test?
  raise "do not run this!"
end

begin
  DB.drop_table :users
  DB.drop_table :poll_timestamps
rescue => e
  puts "error dropping tables"
end
