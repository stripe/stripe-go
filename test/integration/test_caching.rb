# frozen_string_literal: true
# typed: true

require_relative './amendments/_lib'

class Critic::CachingTest < Critic::OrderAmendmentFunctionalTest
  before do
    @user = make_user
  end

  it 'disables the cache and verifies batch service is not called' do
    @user = make_user
    @user.disable_feature(FeatureFlags::SF_CACHING)

    cache_service = CacheService.new(@user)

    cache_service.expects(:retrieve_related_objects_from_salesforce).never

    cache_service.cache_for_object(Restforce::SObject.new({}))
  end

  it 'invalidates an object in the cache' do
    @user.enable_feature(FeatureFlags::SF_CACHING)
    sf_order = create_salesforce_order

    cache_service = CacheService.new(@user)
    cache_service.cache_for_object(sf_order)

    cached_order = cache_service.get_record_from_cache(SF_ORDER, sf_order.Id)
    refute_nil(cached_order)

    cache_service.invalidate_cache_object(sf_order.Id)

    # TODO why do we expect an exception here? The test suite will invalidate cache objects, that should not trigger a cache miss
    exception = assert_raises(Integrations::Errors::TranslatorError) do
      cached_order = cache_service.get_record_from_cache(SF_ORDER, sf_order.Id)
    end

    assert_match("Missed cache for SF Object", exception.message)
  end

  it 'caches records related to an amendment order, then utilizes the cache for translation' do
    skip("cache is disabled") if sf_caching_global_disabled

    @user.enable_feature(FeatureFlags::SF_CACHING)

    # initial order: 1yr contract, monthly billed
    # amendment: starts month 9, lasts 3 months, adds quantity 2

    monthly_price = 10_00
    contract_term = 12
    amendment_term = 3
    start_date = now_time + (contract_term - amendment_term).months
    end_date = start_date + amendment_term.months
    initial_start_date = now_time

    sf_product_id, sf_pricebook_id = salesforce_recurring_product_with_price(price: monthly_price)
    sf_order = create_subscription_order(sf_product_id: sf_product_id)
    sf_contract = create_contract_from_order(sf_order)

    sf_order.refresh

    amendment_data = create_quote_data_from_contract_amendment(sf_contract)

    # increase quantity by 2
    amendment_data["lineItems"].first["record"]["SBQQ__Quantity__c"] = 3

    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_START_DATE] = format_date_for_salesforce(start_date)
    amendment_data["record"][CPQ_QUOTE_SUBSCRIPTION_TERM] = amendment_term

    sf_order_amendment = create_order_from_quote_data(amendment_data)

    sf_order_amendment_contract = create_contract_from_order(sf_order_amendment)

    sf_order_amendment.refresh

    cache_service = CacheService.new(@user)

    cache_service.cache_for_object(sf_order_amendment)

    # Asserting the cache has the following and that it does not reach out to SF to get them:

    @user.expects(:sf_client).never

    #   The order amendment itself
    cached_amendment_order = cache_service.get_record_from_cache(SF_ORDER, sf_order_amendment.Id)

    #   The order amendment order item
    cached_order_item = cache_service.get_record_from_cache(SF_ORDER_ITEM, cached_amendment_order['OrderItems'].first)

    #   The amendment quote
    cached_amendment_quote = cache_service.get_record_from_cache(CPQ_QUOTE, cached_amendment_order['Quotes'].first)

    #   The original order
    cached_order = cache_service.get_record_from_cache(SF_ORDER, sf_order.Id)

    #   The original contract
    cached_contract = cache_service.get_record_from_cache(SF_CONTRACT, sf_contract.Id)
  end
end
