# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class SalesforceUtilTest < Critic::UnitTest
    describe '#backoff' do
      it 'catch and retry known SF error' do
        count = 0
        assert_nothing_raised do
          backoff do
            # raise the error only once
            if count == 0
              count += 1
              raise Restforce::ServerError.new("test")
            end
          end
        end
        assert_equal(1, count)
      end

      it 'known SF error is raised after max retries' do
        count = 0
        assert_raise(Restforce::ServerError) do
            backoff do
              count += 1
              raise Restforce::ServerError.new("test")
            end
        end
        assert_equal(MAX_SF_RETRY_ATTEMPTS, count)
      end

      it 'unknown Salesforce error is raised' do
        count = 0
        assert_raise(RuntimeError) do
          backoff do
            if count == 0
              count += 1
              raise "Unknown error"
            end
          end
        end
        assert_equal(1, count)
      end
    end
  end
end
