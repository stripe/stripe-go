# DO NOT RUN DIRECTLY
# hacky script to help migrate users between databases, leaving around in case we need it in the future

user = StripeForce::User.find(salesforce_instance_url: "https://stripe.my.salesforce.com")
Marshal.dump(user.to_hash)

# separate machine
user_hash = Marshal.load(load_str)
user_hash.delete(:id)
user_hash[:field_defaults] = JSON.parse(user_hash[:field_defaults])
user_hash[:field_mappings] = JSON.parse(user_hash[:field_mappings])
user_hash[:feature_flags] = JSON.parse(user_hash[:feature_flags])

salesforce_refresh_token = user_hash.delete(:salesforce_refresh_token)
salesforce_token = user_hash.delete(:salesforce_token)
connector_settings = JSON.parse(user_hash.delete(:connector_settings))
user_hash.delete(:stripe_refresh_token)

user = StripeForce::User.new(user_hash)
user.salesforce_token = salesforce_token
user.salesforce_refresh_token = salesforce_refresh_token
user.connector_settings = connector_settings
user.save