# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def translate_pricebook(sf_pricebook_entry)
    locker.lock_salesforce_record(sf_pricebook_entry)

    catch_errors_with_salesforce_context(secondary: sf_pricebook_entry) do
      create_price_from_pricebook(sf_pricebook_entry)
    end
  end

  def create_price_from_pricebook(sf_pricebook_entry)
    stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
    return stripe_price if stripe_price

    sf_product = cache_service.get_record_from_cache(SF_PRODUCT, sf_pricebook_entry.Product2Id)
    stripe_product = translate_product(sf_product)

    stripe_price = create_price_from_sf_object(sf_pricebook_entry, sf_product, stripe_product)

    # TODO remove once negative line items are supported
    return if !stripe_price

    update_sf_stripe_id(sf_pricebook_entry, stripe_price)

    stripe_price
  end

  def pricebook_and_order_line_identical?(sf_pricebook_entry, sf_order_item, sf_product)
    log.info 'comparing price with order line',
      sf_pricebook_entry_id: sf_pricebook_entry.Id,
      sf_order_line_id: sf_order_item.Id,
      sf_product_id: sf_product.Id

    # generate the full params that would be sent to the price object and compare them
    # if they aren't *exactly* the same, this will trigger a new price to be created.
    pricebook_params = generate_price_params_from_sf_object(sf_pricebook_entry, sf_product)
    order_item_params = generate_price_params_from_sf_object(sf_order_item, sf_product)

    PriceHelpers.price_billing_amounts_equal?(pricebook_params, order_item_params)
  end

  # main entry point to creating prices from line items from the order logic
  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  # TODO remove nilable return once negative line items are supported https://jira.corp.stripe.com/browse/PLATINT-1483
  sig { params(sf_order_item: T.untyped).returns(T.nilable(Stripe::Price)) }
  def create_price_for_order_item(sf_order_item)
    locker.lock_salesforce_record(sf_order_item)

    log.info 'translating price from a order line', salesforce_object: sf_order_item

    sf_product = cache_service.get_record_from_cache(SF_PRODUCT, sf_order_item.Product2Id)
    product = translate_product(sf_product)

    sf_pricebook_entry = cache_service.get_record_from_cache(SF_PRICEBOOK_ENTRY, sf_order_item.PricebookEntryId)

    # if the order line and pricebook entries are identical then we can reuse the price
    # linked to the pricebook entry. Otherwise, we'll have to create a new price.

    if pricebook_and_order_line_identical?(sf_pricebook_entry, sf_order_item, sf_product)
      existing_stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
      sf_target_for_stripe_price = sf_pricebook_entry

      log.info 'pricebook and order line data is identical, creating or reusing pricebook price',
        pricebook_id: sf_pricebook_entry.Id,
        product_id: sf_product.Id,
        existing_pricebook_price: existing_stripe_price&.id
    else
      sf_target_for_stripe_price = sf_order_item
      existing_stripe_price = retrieve_from_stripe(Stripe::Price, sf_order_item)

      log.info 'pricebook and product data is different, creating or reusing order line price',
        pricebook_id: sf_pricebook_entry.Id,
        product_id: sf_product.Id
    end

    if existing_stripe_price
      existing_stripe_price = T.cast(existing_stripe_price, Stripe::Price)
      generated_stripe_price = generate_price_params_from_sf_object(sf_target_for_stripe_price, sf_product)

      # this should never happen if our identical check is correct, unless the data in Salesforce is mutated over time
      # leave this here as an extra safety check until this has backed in production and our test suite has expanded
      if !PriceHelpers.price_billing_amounts_equal?(existing_stripe_price, generated_stripe_price)
        raise StripeForce::Errors::RawUserError.new("Found a corresponding Stripe price for this Salesforce price, but the price data has changed. Please unlink the Stripe Id from the Salesforce price and sync again.")
      end

      log.info 'reusing existing stripe price', existing_price_id: existing_stripe_price.id
      return existing_stripe_price
    end

    log.info 'existing Stripe price not found, creating new', sf_target_id: sf_target_for_stripe_price

    stripe_price = create_price_from_sf_object(sf_target_for_stripe_price, sf_product, product, sf_order_item)

    # TODO remove once negative line items are supported
    return if !stripe_price

    update_sf_stripe_id(sf_target_for_stripe_price, stripe_price)

    stripe_price
  end

  # this method does NOT check for a pre-existing Stripe object linked to the incoming Salesforce records
  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject, stripe_product: Stripe::Product, sf_order_item: T.nilable(Restforce::SObject)).returns(T.nilable(Stripe::Price)) }
  def create_price_from_sf_object(sf_object, sf_product, stripe_product, sf_order_item=nil)
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise Integrations::Errors::ImpossibleInternalError.new("Stripe price can only be created from an order item or pricebook entry.")
    end

    log.info 'creating price', salesforce_object: sf_object

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we should try to create a price for this in Stripe, otherwise we'll create a price and use that
    stripe_price = generate_price_params_from_sf_object(sf_object, sf_product)

    # TODO these need to be partitioned and created as discounts
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1483
    # amount can be nil if tiered pricing is in place, in which case negative values are not possible
    if !stripe_price.unit_amount_decimal.nil? && stripe_price.unit_amount_decimal < 0
      log.error "negative line item encountered, skipping"
      return
    end

    stripe_price.product = stripe_product.id
    # to_h vs to_hash is important here: to_h returns a empty hash if NilClass
    stripe_price.metadata = (stripe_price['metadata'].to_h || {}).merge(Metadata.stripe_metadata_for_sf_object(@user, sf_object))

    # Is this a subscription price? We only care about frontend proration for recurring prices
    is_evergreen_order_item = sf_order_item && sf_order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE] == CPQProductSubscriptionTypeOptions::EVERGREEN.serialize
    if sf_order_item && !stripe_price["recurring"].to_h.empty? && !is_evergreen_order_item
      sf_order = cache_service.get_record_from_cache(SF_ORDER, sf_order_item["OrderId"])

      # Get the term and billing frequency to calculate proration
      subscription_term = StripeForce::Utilities::SalesforceUtil.extract_subscription_term_from_order!(@mapper, sf_order)
      billing_frequency = StripeForce::Utilities::StripeUtil.billing_frequency_of_price_in_months(stripe_price)

      _is_order_backend_prorated, is_order_frontend_prorated = OrderHelpers.prorated_order?(
        subscription_term: subscription_term,
        billing_frequency: billing_frequency
      )

      if is_order_frontend_prorated && @user.feature_enabled?(FeatureFlags::FRONTEND_PRORATIONS)
        log.info 'prorating price due to frontend prorated order', subscription_term: subscription_term, billing_frequency: billing_frequency, sf_order: sf_order.Id

        # frontend prorated Order (ie 9 month subscription on an annually billed item)
        effective_subscription_term = StripeForce::Utilities::SalesforceUtil.determine_quote_line_subscription_term(mapper, sf_order_item, sf_order)
        prorated_billing_amount = OrderAmendment.calculate_prorated_billing_amount(
          stripe_price: stripe_price,
          subscription_term: subscription_term,
          product_subscription_term: effective_subscription_term,
          billing_frequency: billing_frequency,
          days_prorating: 0
        )

        stripe_price.unit_amount_decimal = prorated_billing_amount.round(MAX_STRIPE_PRICE_PRECISION).to_s("F")
        stripe_price.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::FRONTEND_PRORATION)] = true
      end
    end


    # considered mapping SF pricebook ID to `lookup_key` but it's not *exactly* an external id and more presents a identifier
    # for an externall-used price so the "latest price" for a specific price-type can be used, probably in a website form or something
    # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027

    # mapping is applied within `generate_price_params_from_sf_object`
    sanitize(stripe_price)

    new_price = catch_errors_with_salesforce_context(secondary: sf_object) do
      Stripe::Price.create(
        stripe_price.to_hash,
        @user.stripe_credentials
      )
    end

    log.info 'price created', price_id: new_price.id
    new_price
  end

  sig { params(sf_pricebook_entry: Restforce::SObject).returns(Hash) }
  def extract_tiered_price_params_from_pricebook_entry(sf_pricebook_entry)
    joining_records = cache_service.get_related_records_from_cache(
      sf_pricebook_entry['Product2Id'],
      :ProductId,
      SF_PRODUCT_CONSUMPTION_SCHEDULE
    )

    if joining_records.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if joining_records.count > 1
      raise StripeForce::Errors::RawUserError.new("More than one consumption schedule is linked to a pricebook. There should only be one.", salesforce_object: sf_pricebook_entry['Product2Id'])
    end

    consumption_schedule = cache_service.get_record_from_cache(SF_CONSUMPTION_SCHEDULE, joining_records.first.ConsumptionScheduleId)

    if consumption_schedule.IsDeleted
      log.warn 'consumption schedule is deleted, ignoring'
      return {}
    end

    if !consumption_schedule.IsActive
      log.warn 'consumption schedule is not active, ignoring'
      return {}
    end

    # In our CPQ testing, "Tier" is the only valid picklist value, so we do not expect any other values
    if consumption_schedule.RatingMethod != "Tier"
      raise "unexpected rating method #{consumption_schedule.RatingMethod}"
    end

    consumption_rates = cache_service.get_related_records_from_cache(
      consumption_schedule[SF_ID],
      :ConsumptionScheduleId,
      SF_CONSUMPTION_RATE,
    )

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    tiers_mode = PriceHelpers.transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.Type)

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # for convenience, add a link to the consumption schedule in metadata
      "metadata" => Metadata.stripe_metadata_for_sf_object(@user, consumption_schedule),
    }
  end

  sig { params(sf_consumption_rate: T.untyped).returns(Hash) }
  def transform_salesforce_consumption_rate_type_to_tier(sf_consumption_rate)
    # consumption rate & schedule on the order line vs product are slightly different
    if sf_consumption_rate.sobject_type == SF_CONSUMPTION_RATE
      upper_bound = sf_consumption_rate.UpperBound
      pricing_method = sf_consumption_rate.PricingMethod
      price = sf_consumption_rate.Price
    elsif sf_consumption_rate.sobject_type == CPQ_CONSUMPTION_RATE
      upper_bound = sf_consumption_rate.SBQQ__UpperBound__c
      pricing_method = sf_consumption_rate.SBQQ__PricingMethod__c
      price = sf_consumption_rate.SBQQ__Price__c
    else
      raise "unexpected object type #{sf_consumption_rate.sobject_type}"
    end

    if !upper_bound.nil? && !Integrations::Utilities::StripeUtil.is_integer_value?(upper_bound)
      throw_user_failure!(
        salesforce_object: sf_consumption_rate,
        message: "Decimal value provided for upper tier bound. Consumption rate tier bounds should be integers."
      )
    end

    up_to = if upper_bound.nil?
      'inf'
    else
      upper_bound.to_i
    end

    pricing_key = case pricing_method
    when "PerUnit"
      'unit_amount_decimal'
    when 'FlatFee'
      'flat_amount_decimal'
    else
      # this should never happen unless CPQ changes
      raise StripeForce::Errors::RawUserError.new("Unexpected pricing method found on consumption rate: #{pricing_method}", salesforce_object: sf_consumption_rate)
    end

    stripe_currency = Integrations::Utilities::Currency.currency_for_sf_object(@user, sf_consumption_rate)
    {
      'up_to' => up_to,
      pricing_key => normalize_float_amount_for_stripe(stripe_currency, price.to_s, @user, as_decimal: true),
    }
  end

  sig { params(sf_order_line: Restforce::SObject).returns(Hash) }
  def extract_tiered_price_params_from_order_line(sf_order_line)
    consumption_schedules = cache_service.get_related_records_from_cache(
      sf_order_line[SF_ID],
      :SBQQ__OrderItem__c,
      CPQ_CONSUMPTION_SCHEDULE
    )

    if consumption_schedules.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if consumption_schedules.count > 1
      raise "should not be more than one consumption schedule associated with an order line"
    end

    # from here on out, we know there is exactly one consumption schedule
    # the CPQ consumption schedule is materially different from the standard consumption schedule, so we can't do a drop-in replacement
    consumption_schedule = consumption_schedules.first

    if consumption_schedule.IsDeleted
      log.warn 'consumption schedule is deleted, ignoring'
      return {}
    end

    # unlike pricebook consumption schedule, this consumption schedule does not have a IsActive field

    # in our CPQ testing, "Tier" is the only valid picklist value, so we do not expect any other values
    if consumption_schedule.SBQQ__RatingMethod__c != "Tier"
      raise "unexpected rating method #{consumption_schedule.SBQQ__RatingMethod__c}"
    end

    consumption_rates = cache_service.get_related_records_from_cache(
      consumption_schedule.Id,
      CPQ_CONSUMPTION_SCHEDULE.to_sym,
      CPQ_CONSUMPTION_RATE
    )

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    # TODO can we extract out this hardcoded field here?
    tiers_mode = PriceHelpers.transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.SBQQ__Type__c)

    log.info 'consumption schedule found, configuring as tiered price', consumption_schedule_id: consumption_schedule.Id

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # for convenience, add a link to the consumption schedule in metadata
      "metadata" => Metadata.stripe_metadata_for_sf_object(@user, consumption_schedule),
    }
  end

  # sf_product is required because for some fields the required value is pulled from the product
  # TODO this should be modified to be more flexible and pull via the mapper instead https://jira.corp.stripe.com/browse/PLATINT-1485
  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject).returns(Stripe::Price) }
  def generate_price_params_from_sf_object(sf_object, sf_product)
    # this should never happen, but provides self-documentation and extra test guards
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise ArgumentError.new("price can only be created from an order line or pricebook entry")
    end

    # TODO the tiered pricing logic should be extracted out to a separate method

    # generate these params first since it is hard to merge these
    tiered_pricing_params = if sf_object.sobject_type == SF_PRICEBOOK_ENTRY
      extract_tiered_price_params_from_pricebook_entry(sf_object)
    else
      extract_tiered_price_params_from_order_line(sf_object)
    end

    # NOTE not using `tiered_price?` helper here since this method
    #      constructs the params that the helper uses
    is_tiered_price = !tiered_pricing_params.empty?

    # omitting price param here, this should be defined upstream
    stripe_price = Stripe::Price.construct_from({
      # TODO most likely need to pass the order over?
      currency: Integrations::Utilities::Currency.currency_for_sf_object(@user, sf_object),

      # TODO using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027
    }.merge(tiered_pricing_params))

    # the params are extracted here *and* we apply custom mappings because this enables the user to setup custom mappings for
    # *everything* we sent to the price API including UnitPrice and other fields which go through custom transformations
    price_params = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, sf_object, Stripe::Price)
    mapper.assign_values_from_hash(stripe_price, price_params)

    if sf_object.sobject_type == SF_PRICEBOOK_ENTRY
      apply_mapping(stripe_price, sf_object)
    else
      # https://jira.corp.stripe.com/browse/PLATINT-1486
      apply_mapping(stripe_price, sf_object, compound_key: true)

      # indicate that this price was created as a duplicate for avoid Stripe API errors
      stripe_price[:metadata] ||= {}
      stripe_price[:metadata][Metadata.metadata_key(@user, "auto_archive")] = true
    end

    # although we are passing the amount as a decimal, the decimal amount still represents cents
    # to_s is used here to (a) satisfy typing requirements and (b) ensure BigDecimal can parse the float properly
    stripe_currency = Integrations::Utilities::Currency.currency_for_sf_object(@user, sf_object)
    stripe_price.unit_amount_decimal = T.cast(normalize_float_amount_for_stripe(stripe_currency, stripe_price.unit_amount_decimal.to_s, @user, as_decimal: true), BigDecimal)

    # TODO validate billing frequency and subscription term

    # if tiered pricing is set up, then we ignore any non-tiered pricing configuration
    if is_tiered_price
      stripe_price = PriceHelpers.sanitize_price_tier_params(stripe_price)
    end

    # `recurring` could have been partially constructed by the mapper, which is why we use a symbol-based accessor
    # this avoids nil exception errors in various cases where a field has not been defined
    stripe_price[:recurring] ||= {}

    # by omitting the recurring params Stripe defaults to `type: one_time`
    # this param cannot be set directly
    is_recurring_item = recurring_item?(sf_object.sobject_type == SF_PRICEBOOK_ENTRY ? sf_product : sf_object)

    if is_recurring_item
      # TODO move downcase sanitization upstream
      sf_cpq_term_interval = @user.connector_settings[CONNECTOR_SETTING_CPQ_TERM_UNIT].downcase

      # NOTE on the price level, we are not concerned with the billing term
      #      that can be defined (as a default) on the product, but it only comes into play with the subscription schedule
      if sf_cpq_term_interval != 'month'
        raise 'only monthly terms are currently supported'
      end

      stripe_price.recurring[:usage_type] = PriceHelpers.transform_salesforce_billing_type_to_usage_type(stripe_price.recurring[:usage_type])

      # wipe the `recurring.aggregate_usage` field (if set via the data mapper)
      # since you can only specify `aggregate_usage` for plans with `usage_type=metered`
      # https://jira.corp.stripe.com/browse/PLATINT-1956
      if stripe_price.recurring[:usage_type] == 'licensed'
        log.info "wiping field aggregate_usage since usage usage type is licensed"
        stripe_price.recurring[:aggregate_usage] = nil
      end

      # TODO it's possible that a custom mapping is defined for this value and it's an integer, we should support this case in the helper method
      # this represents how often the price is billed:
      # i.e. if `interval` is month and `interval_count` is 2, then this price is billed every two months.
      stripe_price.recurring[:interval_count] = PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(stripe_price.recurring[:interval_count])

      # frequency: monthly or daily, defined on the CPQ
      stripe_price.recurring[:interval] = sf_cpq_term_interval
    else
      # Remove out any custom mapped recurring params if this is a one-time
      stripe_price[:recurring] = {}
    end

    # prices are only transformed when they are tied to order lines, we trust pricebook prices as is
    # only licensed prices should be customized, tiered + metered billing prices should be left as is
    is_recurring_licensed_price_from_order_line = is_recurring_item &&
      !is_tiered_price &&
      sf_object.sobject_type == SF_ORDER_ITEM

    is_evergreen_item = is_recurring_licensed_price_from_order_line && sf_object[CPQ_PRODUCT_SUBSCRIPTION_TYPE] == CPQProductSubscriptionTypeOptions::EVERGREEN.serialize
    if is_recurring_licensed_price_from_order_line && !PriceHelpers.using_custom_order_line_price_field?(@user) && !is_evergreen_item
      # the formula for calculating the adjusted price is:
      #   billing price = order line unit price / quantity / (subscription term / billing frequency)

      log.info 'custom price not used, adjusting unit_amount_decimal', sf_order_item_id: sf_object.Id

      sf_order = cache_service.get_record_from_cache(SF_ORDER, sf_object['OrderId'])

      billing_frequency = StripeForce::Utilities::StripeUtil.billing_frequency_of_price_in_months(stripe_price)
      price_multiplier = calculate_price_multiplier(mapper, sf_order, sf_object, billing_frequency)
      stripe_price.unit_amount_decimal = T.cast(stripe_price.unit_amount_decimal, BigDecimal) / price_multiplier
    end

    stripe_price
  end

  sig { params(mapper: StripeForce::Mapper, sf_order: Restforce::SObject, sf_order_item: Restforce::SObject, billing_frequency: Integer).returns(BigDecimal) }
  def calculate_price_multiplier(mapper, sf_order, sf_order_item, billing_frequency)
    quote_subscription_term = StripeForce::Utilities::SalesforceUtil.extract_subscription_term_from_order!(mapper, sf_order)
    effective_subscription_term = StripeForce::Utilities::SalesforceUtil.determine_quote_line_subscription_term(mapper, sf_order_item, sf_order)
    cpq_price_multiplier = sf_order_item[CPQ_PRORATE_MULTIPLIER]

    sf_order_end_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_end_date_from_order(mapper, sf_order)
    is_evergreen_order_item = sf_order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE] == CPQProductSubscriptionTypeOptions::EVERGREEN.serialize
    if @user.feature_enabled?(FeatureFlags::NON_ANNIVERSARY_AMENDMENTS) && !sf_order_end_date.nil? && !is_evergreen_order_item
      sf_order_start_date = StripeForce::Utilities::SalesforceUtil.extract_subscription_start_date_from_order(mapper, sf_order)

      # calculate the number of days to prorate
      days = StripeForce::Utilities::SalesforceUtil.calculate_days_to_prorate(
        sf_order_start_date: sf_order_start_date,
        sf_order_end_date: T.must(sf_order_end_date),
        sf_order_subscription_term: quote_subscription_term)

      # if there is a partial month due to a non-anniversary amendment
      # we calculate the price multiplier differently depending on the CPQ Subscription Prorate Precision setting
      if days > 0
        if @user.feature_enabled?(FeatureFlags::DAY_PRORATIONS) || @user.connector_settings[CONNECTOR_SETTING_CPQ_PRORATE_PRECISION] == 'month+day'
          # calculate the price multiplier for when CPQ Subscription Prorate Precision = 'Month + Day'
          log.info 'using \'monthly + daily\' price multiplier calculations', sf_order_id: sf_order.Id, sf_order_item_id: sf_order_item.Id, days: days
          calculated_price_multiplier = StripeForce::Utilities::SalesforceUtil.calculate_month_plus_day_price_multiplier(whole_months: quote_subscription_term, partial_month_days: days, product_subscription_term: effective_subscription_term)
        else
          # calculate the price multiplier for when CPQ Subscription Prorate Precision = 'Month'
          # therefore cpq treats the partial month as a whole month so we add one to the provided subscription term
          log.info 'using \'monthly\' price multiplier calculations', sf_order_id: sf_order.Id, sf_order_item_id: sf_order_item.Id
          quote_subscription_term += 1
          calculated_price_multiplier = BigDecimal(T.must(quote_subscription_term)) / BigDecimal(billing_frequency)
        end

        validate_price_multipliers(calculated_price_multiplier, cpq_price_multiplier, true)
        return calculated_price_multiplier
      end
    end

    # TODO should we adjust based on the quantity? Most likely, let's wait until tests fail
    price_multiplier = BigDecimal(T.must(quote_subscription_term)) / BigDecimal(billing_frequency)

    # TODO should test this further with proration amendments
    # For MDQ orders, the quote subscription term is not the effective subscription term
    if !validate_price_multipliers(price_multiplier, cpq_price_multiplier, false) && quote_subscription_term != effective_subscription_term
      price_multiplier = BigDecimal(T.must(effective_subscription_term)) / BigDecimal(billing_frequency)
    end

    price_multiplier
  end

  sig { params(calculated_price_multiplier: BigDecimal, cpq_price_multiplier: Float, throw_error: T.nilable(T::Boolean)).returns(T::Boolean) }
  def validate_price_multipliers(calculated_price_multiplier, cpq_price_multiplier, throw_error)
    # check that the calculated price multiplier is equal to the cpq provided price multiplier
    # throw an error if they are not equal
    # note: we do not use the cpq_price_multiplier since the prorate multiplier field that you see on CPQ objects is rounded
    if !cpq_price_multiplier.nil?
      cpq_price_multiplier = cpq_price_multiplier.to_d
      threshold = 0.0000000001
      if (calculated_price_multiplier - cpq_price_multiplier).abs > threshold
        log.error 'calculated price multipler does not equal CPQ price multiplier', calculated_price_multiplier: calculated_price_multiplier, cpq_price_multiplier: cpq_price_multiplier
        if throw_error
          raise Integrations::Errors::TranslatorError.new("calculated price multiplier differs from cpq price multiplier")
        end
        return false
      end
    end

    true
  end
end
