# frozen_string_literal: true
# typed: true

module Resque
  module Failure
    class Backtrace < Base
      def save
        if exception.class != Integrations::Errors::LockTimeout
          bt = filter_backtrace(exception.backtrace)
          Integrations::Log.log.error 'resque exception',
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
