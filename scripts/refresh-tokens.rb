require File.expand_path('../config/environment', __dir__)

user = StripeForce::User.find(salesforce_account_id: ENV.fetch('SF_INSTANCE_ID'))


shell_vars = <<~EOL
  SF_REFRESH_TOKEN="#{user.salesforce_refresh_token}"
  SF_ACCESS_TOKEN="#{user.salesforce_token}"
EOL

puts shell_vars

shell_vars.split("\n").each do |var|
  name, value = var.split("=", 2)
  `sed -i '' -e 's/#{name}.*/#{name}=#{value}/' '.env'`
end
