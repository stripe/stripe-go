# typed: true
# frozen_string_literal: true

require 'sequel/plugins/defaults_setter'

# TODO submit a PR for this upstream
Sequel::Plugins::DefaultsSetter::ClassMethods.class_eval do
  # Parse the cached database schema for this model and set the default values appropriately.
  def set_default_values
    T.bind(self, Sequel::Plugins::DefaultsSetter::ClassMethods)

    h = {}
    if @db_schema
      @db_schema.each do |k, v|
        if v[:callable_default]
          h[k] = v[:callable_default]
        elsif !v[:ruby_default].nil?
          h[k] = convert_default_value(v[:ruby_default])
        elsif v[:db_type] == 'jsonb' && v[:default]
          # extract JSON to use as the initial default value
          h[k] = v[:default].match(/'(.*)'::jsonb/)[1]
        end
      end
    end
    @default_values = h.merge!(@default_values || Sequel::OPTS)
  end
end

is_test_environment = ENV['RAILS_ENV'] && ENV['RAILS_ENV'] == 'test'

url = if is_test_environment
  ENV.fetch('TEST_DATABASE_URL')
else
  ENV.fetch('DATABASE_URL')
end

DB = Sequel.connect(url)
