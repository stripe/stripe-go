# frozen_string_literal: true
# typed: true
module Resque
  module Failure
    class Backtrace < Base
      def save
        # TODO modify once we have the locktime defined here
        # if exception.class != StripeSuite::Errors::LockTimeout
        if true
          bt = filter_backtrace(exception.backtrace)
          SimpleStructuredLogger::Writer.instance.error 'resque exception',
            exception_class: exception.class,
            exception_message: exception.message,
            exception_backtrace: bt.join("\n")
        end
      end

      private def filter_backtrace(backtrace)
        return [] if backtrace.nil?

        index = backtrace.index {|item| item.include?('/lib/resque/job.rb') }
        backtrace.first(index.to_i)
      end

    end
  end
end
