# frozen_string_literal: true
# typed: true

require_relative './_revenue_contract_validation_helper'

class Critic::RevRecContractCreation < Critic::RevenueContractValidationHelper
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @defaultSignedDate = "2022-12-31"
    @defaultTFC = 30
    @user = make_user(save: true)
    @user.enable_feature FeatureFlags::STRIPE_REVENUE_CONTRACT, update: true
    @user.field_defaults = {
      "subscription_item" => {
        "metadata.contract_tfc_duration" => @defaultTFC,
      },
      "subscription_schedule" => {
        "metadata.contract_cf_signed_date" => @defaultSignedDate,
      },
    }
    @user.save
  end

  it 'Create contract from a standard subscription order' do
    backdated_months = 2
    price = 10_00
    subscription_term = 12
    start_date = now_time - backdated_months.months

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price)

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,
      contact_email: "create_contract_standard_sub_order",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    revenue_contract = get_revenue_contract(contract_id)

    revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate)
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)
    assert_equal(1, phase.items.count)
    assert_equal(0, phase.add_invoice_items.count)
    phase_item = T.must(phase.items.first)

    assert_equal(1, revenue_contract.items.data.count)
    contract_item = T.must(revenue_contract.items.data.first)
    revenue_contract_validate_item(phase_item, contract_item, sf_pricebook_entry, 1, 12000, @defaultTFC)
    assert_equal(start_date.to_i, contract_item.period.start)

    end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
      Time.at(start_date) + subscription_term.months
    )
    assert_equal(end_date, contract_item.period.end)
  end

  it 'Create contract from a standard subscription order with override contract value' do
    @user.field_defaults = {
      "subscription_item" => {
        "metadata.item_contract_value" => "20000.52",
      },
      "subscription_schedule" => {
        "metadata.contract_cf_signed_date" => @defaultSignedDate,
      },
    }
    @user.save

    backdated_months = 2
    price = 10_00
    subscription_term = 12
    start_date = now_time - backdated_months.months

    # creating these directly so we have the IDs
    sf_product_id, sf_pricebook_entry_id = salesforce_recurring_product_with_price(price: price)

    sf_account_id = create_salesforce_account
    sf_order = create_salesforce_order(
      sf_product_id: sf_product_id,
      sf_account_id: sf_account_id,
      contact_email: "create_contract_override_sub_order",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => format_date_for_salesforce(start_date),
        CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,
      }
    )

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_entry_id)

    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    revenue_contract = get_revenue_contract(contract_id)

    revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate)
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)
    assert_equal(1, phase.items.count)
    assert_equal(0, phase.add_invoice_items.count)
    phase_item = T.must(phase.items.first)

    assert_equal(1, revenue_contract.items.data.count)
    contract_item = T.must(revenue_contract.items.data.first)
    revenue_contract_validate_item(phase_item, contract_item, sf_pricebook_entry, 1, 20000.52, nil)
    assert_equal(start_date.to_i, contract_item.period.start)

    end_date = StripeForce::Utilities::SalesforceUtil.datetime_to_unix_timestamp(
      Time.at(start_date) + subscription_term.months
    )
    assert_equal(end_date, contract_item.period.end)
  end

  it 'multiple line items with 1 skipped and 1 0 amount' do
    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_product_with_price
    sf_product_id_3, sf_pricebook_id_3 = salesforce_recurring_product_with_price
    sf_product_id_4, sf_pricebook_id_4 = salesforce_recurring_product_with_price(price: 0)

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(sf_account_id: sf_account_id,
                                       contact_email: "multiple_line_items_1_skip_1_zero",
                                       additional_quote_fields: {
                                         CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
                                         CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
                                       })

    # only CPQ fields can be customized within this special quote creation process

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_3)
    calculate_and_save_cpq_quote(quote_with_product)

    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_4)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    sf_line_items = sf_get_related(sf_order, SF_ORDER_ITEM)
    assert_equal(4, sf_line_items.size)
    second_line_item = sf_line_items.detect {|i| i.Product2Id == sf_product_id_2 }
    fourth_line_item = sf_line_items.detect {|i| i.Product2Id == sf_product_id_4 }

    sf.update!(SF_ORDER_ITEM, {
      SF_ID => second_line_item.Id,
      prefixed_stripe_field(ORDER_LINE_SKIP) => true,
    })

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    revenue_contract = get_revenue_contract(contract_id)

    revenue_contract_validate_basics(sf_order, subscription_schedule, revenue_contract, sf_account_id, @defaultSignedDate)
    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    # should have four, but one was excluded
    assert_equal(3, phase.items.count)
    assert_equal(0, phase.add_invoice_items.count)
    assert_equal(3, revenue_contract.items.data.count)

    phase_item_1 = T.must(phase.items[0])
    matching_contract_Item = T.must(revenue_contract.items.data.find {|x| x.price == phase_item_1.price })
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id_1)
    revenue_contract_validate_item(phase_item_1, matching_contract_Item, sf_pricebook_entry, 1, 144000, @defaultTFC)

    phase_item_2 = T.must(phase.items[1])
    matching_contract_Item = T.must(revenue_contract.items.data.find {|x| x.price == phase_item_2.price })
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id_3)
    revenue_contract_validate_item(phase_item_2, matching_contract_Item, sf_pricebook_entry, 1, 144000, @defaultTFC)

    phase_item_3 = T.must(phase.items[2])
    matching_contract_Item = T.must(revenue_contract.items.data.find {|x| x.price == phase_item_3.price })
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id_4)
    revenue_contract_validate_item(phase_item_3, matching_contract_Item, sf_pricebook_entry, 1, 0, @defaultTFC)

    assert_empty(phase.items.select do |i|
      refute_nil(i.metadata['salesforce_order_item_id'])
      i.metadata['salesforce_order_item_id'] == second_line_item.Id
    end)
  end

  it 'multiple line items with a metered price' do
    # price customization makes it easier to debug and ensure there aren't weird state bugs
    # across different price types
    monthly_price = 120_00
    one_time_price = 45_00
    metered_price = 55_00

    sf_product_id_1, sf_pricebook_id_1 = salesforce_recurring_product_with_price(price: monthly_price)
    sf_product_id_2, sf_pricebook_id_2 = salesforce_recurring_metered_produce_with_price(price_in_cents: metered_price)
    sf_product_id_3, sf_pricebook_id_3 = salesforce_standalone_product_with_price(price: one_time_price)

    sf_account_id = create_salesforce_account

    quote_id = create_salesforce_quote(
      sf_account_id: sf_account_id,
      contact_email: "multi_line_item_metered",
      additional_quote_fields: {
        CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
        CPQ_QUOTE_SUBSCRIPTION_TERM => 12.0,
      }
    )

    # set first product quantity to 5
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_1)
    quote_with_product["lineItems"].first["record"][CPQ_QUOTE_QUANTITY] = 5.0
    calculate_and_save_cpq_quote(quote_with_product)

    # recurring and arrears, should not have a quantity set when passed to stripe
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_2)
    calculate_and_save_cpq_quote(quote_with_product)

    # third product should be standalone and not recurring
    quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id_3)
    calculate_and_save_cpq_quote(quote_with_product)

    sf_order = create_order_from_cpq_quote(quote_id)

    SalesforceTranslateRecordJob.translate(@user, sf_order)

    sf_order.refresh
    stripe_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    contract_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_REVENUE_CONTRACT_ID)]
    subscription_schedule = Stripe::SubscriptionSchedule.retrieve(stripe_id, @user.stripe_credentials)
    revenue_contract = get_revenue_contract(contract_id)

    assert_equal(1, subscription_schedule.phases.count)
    phase = T.must(subscription_schedule.phases.first)

    assert_equal(2, phase.items.count)
    assert_equal(1, phase.add_invoice_items.count)

    # There is one less item in revenue contracts since it does not track metered billing items.
    assert_equal(2, revenue_contract.items.data.count)

    phase_item_with_five = T.must(phase.items.detect {|i| i.try(:quantity) == 5 })
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id_1)
    matching_contract_Item = T.must(revenue_contract.items.data.find {|x| x.price == phase_item_with_five.price })
    revenue_contract_validate_item(phase_item_with_five, matching_contract_Item, sf_pricebook_entry, 5, 720000, @defaultTFC)

    standalone_item = T.must(phase.add_invoice_items.first)
    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_pricebook_id_3)
    matching_contract_Item = T.must(revenue_contract.items.data.find {|x| x.price == standalone_item.price })
    revenue_contract_validate_item(standalone_item, matching_contract_Item, sf_pricebook_entry, 1, 4500, @defaultTFC)

    phase_item_with_metered = T.must(phase.items.detect {|i| !i.respond_to?(:quantity) })
    price_2 = Stripe::Price.retrieve(T.cast(phase_item_with_metered.price, String), @user.stripe_credentials)
    assert_equal(sf_pricebook_id_2, price_2.metadata['salesforce_pricebook_entry_id'], "pricebook entry does not exist, price may be created in error from an order line")
  end
end
