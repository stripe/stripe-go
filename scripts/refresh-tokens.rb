# typed: true
# Description: the oauth tokens can easily break if you are testing prod
#              if you reauth locally and then run this script, .env will be updated

require File.expand_path('../config/environment', __dir__)

unless Rails.env.development?
  puts "Should only be run locally"
  abort
end

username = ARGV[0]

if !username
  puts "No username provided, finding local user reference"
  user = StripeForce::User.find(salesforce_account_id: ENV.fetch('SF_INSTANCE_ID'))
end

if !user && username
  puts "Username specified, assuming cpq alias"
elsif !user && !username
  puts "No local user found, finding default username"
  username = `cd sfdx && sfdx config:get defaultusername --json | jq -r '.result[0].value'`.strip
  puts "Found default username '#{username}'"
end

# optionally pass the alias/name you want to refresh from
if username
  auth_list = JSON.parse(`cd sfdx && sfdx auth:list --json`)
  auth_info = auth_list['result'].detect { |a| a['username'] == username || a['alias'] == username }

  target_url = if auth_info['isScratchOrg']
    "https://test.salesforce.com"
  else
    "https://login.salesforce.com"
  end

  access_token = if auth_info['isScratchOrg']
    auth_info['accessToken']
  else
     `SF_URL=#{target_url} SF_USERNAME=#{username} SF_CONSUMER_KEY=#{ENV['SF_CONSUMER_KEY']} SF_JWT_PRIVATE_KEY_PATH=./sfdx/jwt-cert/private_key.pem node ./sfdx/bin/jwt-generator/index.js`.strip
  end

  # TODO refresh tokens are available in the ~/.sfdx token even though they do not show up in `auth:list`

  instance_id = auth_info['orgId']
  instance_domain = auth_info['instanceUrl'].match(/https:\/\/(.*)\.my.salesforce.com/)[1]

  shell_vars = <<~EOL
    SF_INSTANCE_ID=#{instance_id}
    SF_INSTANCE_DOMAIN=#{instance_domain}
    SF_ACCESS_TOKEN="#{access_token}"
    SF_REFRESH_TOKEN=" "
  EOL
else
  user = T.must(user)
  access_token = user.salesforce_access_token

  shell_vars = <<~EOL
    SF_REFRESH_TOKEN="#{user.salesforce_refresh_token}"
    SF_ACCESS_TOKEN="#{access_token}"
  EOL
end

puts "# http https://#{instance_domain}.my.salesforce.com 'Authorization: OAuth #{access_token}'"
puts shell_vars

if Rails.env.development?
  shell_vars.split("\n").each do |var|
    name, value = var.split("=", 2)
    `sed -i '' -e 's/^#{name}.*/#{name}=#{value}/' '.env'`
  end
end