# frozen_string_literal: true
# typed: false

Restforce.configure do |config|
  config.log_level = :debug
  # config.log = true
end

# really? Can't set this on an instance or `configure` level?
Restforce.log = ENV.fetch('SALESFORCE_LOG', 'false') == 'true'

Restforce::SObject.class_eval do
  def refresh
    ensure_id
    new_record = @client.find(sobject_type, self.Id)
    # TODO unsure if this is the right call
    deep_update(new_record.attrs)
  end
end
