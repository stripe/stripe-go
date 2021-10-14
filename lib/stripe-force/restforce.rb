Restforce.configure do |config|
  config.log_level = :debug
  # config.log = true
end


# really? Can't set this on an instance or `configure` level?
Restforce.log = ENV.fetch('SALESFORCE_LOG', 'false') == 'true'
