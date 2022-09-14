# frozen_string_literal: true
# typed: true
require_relative '../test_helper'

module Critic::Unit
  class OrderAmendmentTest < Critic::UnitTest
    before do
      @user = make_user
    end

    describe '#delete_past_phases' do
      it 'deletes phases that have already ended' do
        one_day = 60 * 60 * 24

        customer = create_customer
        customer_id = customer.id

        now_timestamp = now_time.to_i

        phases = [
          Stripe::StripeObject.construct_from({
            end_date: now_timestamp - one_day,
          }),
          Stripe::StripeObject.construct_from({
            end_date: now_timestamp + one_day,
          }),
        ]

        updated_phases = StripeForce::Translate::OrderAmendment.delete_past_phases(@user, customer_id, phases)

        assert_equal(2, phases.count)
        assert_equal(1, updated_phases.count)

        assert_equal(now_timestamp + one_day, T.must(updated_phases.first).end_date)
      end
    end

    describe '#determine_current_time' do
      it 'uses the virtualized time when a test clock exists' do
        one_day = 60 * 60 * 24
        one_day_ago = now_time.to_i - one_day

        test_clock = Stripe::TestHelpers::TestClock.create({
          frozen_time: one_day_ago,
        }, @user.stripe_credentials)

        stripe_customer = create_customer(additional_fields: {test_clock: test_clock})

        current_time = StripeForce::Translate::OrderAmendment.determine_current_time(@user, stripe_customer.id)

        # allow for a 5s variance in case API response is slow
        assert(current_time - one_day_ago <= 5)
      end

      it 'uses the real current time when no test clock exists' do
        stripe_customer = create_customer
        current_time = StripeForce::Translate::OrderAmendment.determine_current_time(@user, stripe_customer.id)
        assert(Time.now.to_i - current_time <= 5)
      end
    end
  end
end
