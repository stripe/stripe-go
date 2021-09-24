# heroku run bundle exec ruby migration.rb

require 'sequel'
DB = Sequel.connect(ENV.fetch('DATABASE_URL'))

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
