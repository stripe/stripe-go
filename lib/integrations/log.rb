# typed: true
# frozen_string_literal: true

SimpleStructuredLogger::Writer.class_eval do
  def level(level=nil)
    if level
      @logger.level = level
    else
      @logger.level
    end
  end
end

# more simple log formatting
SimpleStructuredLogger::Writer.instance.logger.formatter = proc do |severity, _datetime, _progname, msg|
  "#{severity}: #{msg}\n"
end

module Integrations
  module Log
    include SimpleStructuredLogger

    # TODO handle metrics expansion
  end
end
