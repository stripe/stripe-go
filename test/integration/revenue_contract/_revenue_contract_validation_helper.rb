# typed: true
# frozen_string_literal: true

require_relative '../amendments/_lib'

class Critic::RevenueContractValidationHelper < Critic::OrderAmendmentFunctionalTest
  def revenue_contract_validate_basics(
    sf_order,
    subscription_schedule,
    revenue_contract,
    sf_account_id,
    signed_date,
    version: 1
  )
    customer = Stripe::Customer.retrieve(T.cast(subscription_schedule.customer, String), @user.stripe_credentials)

    # basic customer creation check
    refute_empty(customer.name)
    assert_nil(customer.email)
    assert_equal(customer.id, revenue_contract.customer)

    assert_match(sf_account_id, customer.metadata['salesforce_account_link'])
    assert_equal(customer.metadata['salesforce_account_id'], sf_account_id)

    # top level subscription fields
    assert_match(sf_order.Id, subscription_schedule.metadata['salesforce_order_link'])
    assert_equal(subscription_schedule.metadata['salesforce_order_id'], sf_order.Id)
    assert_equal(signed_date, subscription_schedule.metadata['contract_cf_signed_date'])

    # top level contract fields
    # These are only specific to creation right now, will need to update when we handle amendments\
    if subscription_schedule.status == "canceled"
      # TODO: Eventually this can result in uncollectible.
      assert_equal("void", revenue_contract.status)
      assert_equal(subscription_schedule.canceled_at, revenue_contract.status_transitions.voided_at)
    else
      assert_equal("signed", revenue_contract.status)
      assert_equal(DateTime.parse(signed_date).to_i, revenue_contract.status_transitions.signed_at)
    end

    assert_equal(version, revenue_contract.version)
    assert_equal(subscription_schedule.metadata['salesforce_order_link'], revenue_contract.metadata['salesforce_order_link'])
    assert_equal(subscription_schedule.metadata['salesforce_order_id'], revenue_contract.metadata['salesforce_order_id'])
    assert_equal(subscription_schedule.metadata['contract_cf_signed_date'], revenue_contract.metadata['contract_cf_signed_date'])
    assert_equal(subscription_schedule.metadata.count, revenue_contract.metadata.count)
  end

  def revenue_contract_validate_item(
    phase_item,
    contract_item,
    sf_pricebook_entry,
    quantity,
    amount,
    tfc
  )
    if !sf_pricebook_entry.nil?
      assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], phase_item.price)
      assert_equal(sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)], contract_item.price)
    end

    assert_equal(quantity, phase_item.quantity)
    assert_equal(quantity, contract_item.quantity)

    if !tfc.nil?
      assert_equal(tfc.to_s, phase_item.metadata['contract_tfc_duration'])
      assert_equal(tfc.to_s, contract_item.metadata['contract_tfc_duration'])
      expires_at = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
        Time.at(contract_item.period.start) + (tfc.to_i + 1).days
      )
      assert_equal(expires_at, contract_item.termination_for_convenience.expires_at)
    else
      assert_nil(contract_item.termination_for_convenience)
    end

    amount_subtotal = amount
    # TODO: Currently contract item does not support metadata update, we need to support this due to
    # amendments adding in extra metadata to the phase item but not to contract item.
    assert(phase_item.metadata.count >= contract_item.metadata.count)
    if !phase_item.metadata['item_contract_value'].nil?
      assert_equal(amount.to_s, phase_item.metadata['item_contract_value'])
      assert_equal(amount.to_s, contract_item.metadata['item_contract_value'])
      amount_subtotal = amount * 100
    end

    assert_equal(amount_subtotal, contract_item.amount_subtotal)
  end

  def revenue_contract_validate_item_period(
    contract_item,
    start_date,
    end_date
  )
    assert_equal(start_date.is_a?(Integer) ? start_date : start_date.to_time.to_i, contract_item.period.start)
    assert_equal(end_date.is_a?(Integer) ? end_date : end_date.to_time.to_i, contract_item.period.end)
  end
end
