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

    describe '#salesforce_type_from_id' do
      it 'does not fetch custom sf obj prefix when in cache' do
        @user = make_user

        # sanity check to make sure prefix mappings are clean
        assert_equal({}, @user.salesforce_object_prefix_mappings)

        # setup
        sf_object_id = 'a257V000008itIeQAI'
        sf_object_prefix = 'a25'
        sf_object_name = 'SBQQ__Quote__c'

        @user.salesforce_object_prefix_mappings[sf_object_prefix] = sf_object_name
        @user.save

        # we don't expect to make an api call for the object info
        @user.expects(:save).never
        StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_object_id)
      end

      it 'attempts to fetch custom sf obj prefix when not in db' do
        @user = make_user(save: true)

        # sanity check to make sure prefix mappings are clean
        assert_equal({}, @user.salesforce_object_prefix_mappings)

        # setup
        sf_object_id = '6S97V000008itIeQAI'
        sf_object_prefix = '6S9'
        sf_object_name = "AIApplicationConfig"

        StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_object_id)

        # we expect to have saved the object info
        assert_equal({sf_object_prefix => sf_object_name}, @user.salesforce_object_prefix_mappings)
      end
    end
  end
end
