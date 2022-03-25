# frozen_string_literal: true
# typed: false

# https://github.com/iloveitaly/simple_structured_logger/blob/master/lib/simple_structured_logger.rb
# TODO pull into open source gem
SimpleStructuredLogger.class_eval do
  def self.configure(&block)
    SimpleStructuredLogger::Configuration.instance_eval(&block)
  end
end

SimpleStructuredLogger::Configuration.class_eval do
  def expand_context(&block)
    if block.nil?
      @expand_context
    else
      @expand_context = block
    end
  end

  def expand_log(&block)
    if block.nil?
      @expand_log
    else
      @expand_log = block
    end
  end
end

SimpleStructuredLogger::Writer.class_eval do
  def level(level=nil)
    if level
      @logger.level = level
    else
      @logger.level
    end
  end

  def stringify_tags(additional_tags)
    additional_tags = additional_tags.dup

    if SimpleStructuredLogger::Configuration.expand_log
      additional_tags = SimpleStructuredLogger::Configuration.expand_log.call(additional_tags, self.default_tags)
    end

    @default_tags.merge(additional_tags).map {|k, v| "#{k}=#{v}" }.join(' ')
  end
end

# more simple log formatting
SimpleStructuredLogger::Writer.instance.logger.formatter = proc do |severity, _datetime, _progname, msg|
  "#{severity}: #{msg}\n"
end

SimpleStructuredLogger.configure do
  expand_log do |tags, default_tags|
    if tags[:salesforce_object] && tags[:salesforce_object].is_a?(Restforce::SObject)
      salesforce_object = tags.delete(:salesforce_object)
      tags[:salesforce_object_id] = salesforce_object.Id
      tags[:salesforce_object_type] = salesforce_object.sobject_type
    end

    if tags[:stripe_resource] && tags[:stripe_resource].respond_to?(:id)
      stripe_resource = tags.delete(:stripe_resource)
      tags[:stripe_resource_id] = stripe_resource.id
      tags[:stripe_resource_type] = stripe_resource.class.to_s
    end

    if tags[:metric]
      dimensions = default_tags.slice(:stripe_user_id, :livemode, :salesforce_id)
      Integrations::Metrics::Writer.instance.track_counter(tags[:metric], dimensions: dimensions)
    end

    tags
  end
end

module Integrations
  module Log
    include SimpleStructuredLogger

    # TODO handle metrics expansion
  end
end
