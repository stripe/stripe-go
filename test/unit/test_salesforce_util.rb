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

    describe '#calculate_days_to_prorate' do
      it 'calculate days to prorate' do
        # the amendment order end date should always be equal to the initial order end date since we
        # require amendment orders to coterminate with the initial order
        order_end_date = Time.new(2024, 1, 1)

        order_start_date = Time.new(2023, 1, 2)
        order_subscription_term = 11
        days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(sf_order_start_date: order_start_date, sf_order_end_date: order_end_date, sf_order_subscription_term: order_subscription_term)
        assert_equal(30, days)

        order_start_date = Time.new(2023, 1, 30)
        order_subscription_term = 11
        days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(sf_order_start_date: order_start_date, sf_order_end_date: order_end_date, sf_order_subscription_term: order_subscription_term)
        assert_equal(2, days)

        order_start_date = Time.new(2023, 2, 27)
        order_subscription_term = 10
        days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(sf_order_start_date: order_start_date, sf_order_end_date: order_end_date, sf_order_subscription_term: order_subscription_term)
        assert_equal(5, days)
      end

      it 'calculate days to prorate in leap year' do
        # note: 2024 is a leap year
        order_end_date = Time.new(2024, 2, 28)

        order_start_date = Time.new(2023, 3, 30)
        order_subscription_term = 10
        days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(sf_order_start_date: order_start_date, sf_order_end_date: order_end_date, sf_order_subscription_term: order_subscription_term)
        assert_equal(29, days)
      end

      # TODO test whole month eom cases
      # https://help.salesforce.com/s/articleView?id=sf.cpq_whole_month_calc_guidelines.htm&type=5
      it 'calculate days to prorate at eom' do
        order_end_date = Time.new(2024, 3, 15)

        # in this example, June 30 is the EOM, so when we add 9.months, it should end March 31 (eom)
        order_start_date = Time.new(2023, 12, 31)
        order_subscription_term = 2
        # days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(sf_order_start_date: order_start_date, sf_order_end_date: order_end_date, sf_order_subscription_term: order_subscription_term)
        # assert_equal(14, days)
      end
    end
  end
end
