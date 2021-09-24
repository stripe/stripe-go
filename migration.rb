require_relative 'config'

DB.create_table :users do
  primary_key :id

  String :salesforce_account_id, unique: true, null: false
  String :salesforce_refresh_token, null: false
  String :salesforce_instance_url, null: false
  String :salesforce_token, null: false

  String :stripe_account_id
  String :stripe_refresh_token
  String :stripe_public_token

  TrueClass :enabled, default: true

  String :name
  String :email
end
