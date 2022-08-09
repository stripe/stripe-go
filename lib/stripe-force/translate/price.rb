# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  def create_price_from_pricebook(sf_pricebook_entry)
    stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
    return stripe_price if stripe_price

    sf_product = sf.find(SF_PRODUCT, sf_pricebook_entry.Product2Id)
    stripe_product = create_product_from_sf_product(sf_product)

    stripe_price = create_price_from_sf_object(sf_pricebook_entry, sf_product, stripe_product)

    # TODO remove once negative line items are supported
    return if !stripe_price

    update_sf_stripe_id(sf_pricebook_entry, stripe_price)

    stripe_price
  end

  # TODO maybe look at "can edit price" boolean on the product? But what about custom mappings?
  # TODO https://jira.corp.stripe.com/browse/PLATINT-1485
  def pricebook_and_order_line_identical?(sf_pricebook_entry, sf_order_item, sf_product)
    # generate the full params that would be sent to the price object and compare them
    # if they aren't *exactly* the same, this will trigger a new price to be created
    # we'll probably need to do something a bit smarter here in the future but it's not obvious what
    # additional logic we should include here at the moment.
    pricebook_params = generate_price_params_from_sf_object(sf_pricebook_entry, sf_product).to_hash
    order_item_params = generate_price_params_from_sf_object(sf_order_item, sf_product).to_hash

    # metadata *could* have important financial data for downstream systems
    # however, most of the time users will not pull this information directly from prices
    # (they may use products instead) and a users mappings could change over time, so we don't
    # want to factor it in to our equality test
    pricebook_params.delete(:metadata)
    order_item_params.delete(:metadata)

    # creating prices unnecessarily can be problematic for users: many prices without clarity about why they exist
    # for this reason, let's log the diff in debug mode so we can easily determine exactly WHY the new price was created
    if pricebook_params != order_item_params
      log.debug 'price diff', diff: HashDiff::Comparison.new(pricebook_params, order_item_params).diff
    end

    pricebook_params == order_item_params
  end

  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  def create_price_for_order_item(sf_order_item)
    log.info 'translating price from a order line origin', salesforce_object: sf_order_item

    sf_product = sf.find(SF_PRODUCT, sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_order_item.PricebookEntryId)

    # TODO should be able to use the pricebook entries for these checks

    # if the order line and pricebook entries are identical then we can use a pre-existing price
    # otherwise, we'll have to create a new price

    if pricebook_and_order_line_identical?(sf_pricebook_entry, sf_order_item, sf_product)
      existing_stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
      sf_target_for_stripe_price = sf_pricebook_entry

      log.info 'pricebook and product data is identical, attemping to use existing stripe price',
        pricebook_id: sf_pricebook_entry.Id,
        product_id: sf_product.Id,
        existing_price: existing_stripe_price&.id
    else
      log.info 'pricebook and product data is different, creating new price from order line',
        pricebook_id: sf_pricebook_entry.Id,
        product_id: sf_product.Id

      sf_target_for_stripe_price = sf_order_item
    end

    if existing_stripe_price
      existing_stripe_price = T.cast(existing_stripe_price, Stripe::Price)
      generated_stripe_price = generate_price_params_from_sf_object(sf_pricebook_entry, sf_product)

      # this should never happen if our identical check is correct, unless the data in Salesforce is mutated over time
      if PriceHelpers.price_billing_amounts_equal?(existing_stripe_price, generated_stripe_price)
        raise Integrations::Errors::UnhandledEdgeCase.new("expected generated prices to be equal, but they differed")
      end

      log.info 'using existing stripe price'
      return existing_stripe_price
    end

    log.info 'existing price not found, creating new price'
    stripe_price = create_price_from_sf_object(sf_target_for_stripe_price, sf_product, product)

    # TODO remove once negative line items are supported
    return if !stripe_price

    update_sf_stripe_id(sf_target_for_stripe_price, stripe_price)

    stripe_price
  end

  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject, stripe_product: Stripe::Product).returns(T.nilable(Stripe::Price)) }
  def create_price_from_sf_object(sf_object, sf_product, stripe_product)
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise "price can only be created from an order line or pricebook entry"
    end

    log.info 'creating price', salesforce_object: sf_object

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we shoudl try to create a price for this in Stripe, otherwise we'll create a price and use that
    stripe_price = generate_price_params_from_sf_object(sf_object, sf_product)

    # TODO these need to be partitioned and created as discounts
    # TODO https://jira.corp.stripe.com/browse/PLATINT-1483
    # amount can be nil if tiered pricing is in place, in which case negative values are not possible
    if !stripe_price.unit_amount_decimal.nil? && stripe_price.unit_amount_decimal < 0
      log.error "negative line item encountered, skipping"
      return
    end

    stripe_price.product = stripe_product.id
    stripe_price.metadata = (stripe_price['metadata'].to_h || {}).merge(stripe_metadata_for_sf_object(sf_object))

    # considered mapping SF pricebook ID to `lookup_key` but it's not *exactly* an external id and more presents a identifier
    # for an externall-used price so the "latest price" for a specific price-type can be used, probably in a website form or something
    # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027

    sanitize(stripe_price)

    catch_stripe_api_errors(sf_object) do
      Stripe::Price.create(stripe_price.to_hash, @user.stripe_credentials)
    end
  end

  sig { params(sf_pricebook_entry: Restforce::SObject).returns(Hash) }
  def extract_tiered_price_params_from_pricebook_entry(sf_pricebook_entry)
    # TODO use cache_service
    joining_records = sf.query(
      <<~EOL
        SELECT ConsumptionScheduleId
        FROM ProductConsumptionSchedule
        WHERE ProductId = '#{sf_pricebook_entry['Product2Id']}'
      EOL
    )

    if joining_records.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if joining_records.count > 1
      raise "should not be more than one consumption schedule linked to a pricebook"
    end

    # TODO use cache_service
    consumption_schedule = sf.find(SF_CONSUMPTION_SCHEDULE, joining_records.first.ConsumptionScheduleId)

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

    # TODO use cache_service
    consumption_rates = sf.query(
      <<~EOL
        SELECT Id
        FROM #{SF_CONSUMPTION_RATE}
        WHERE ConsumptionScheduleId = '#{consumption_schedule.Id}'
      EOL
    ).map {|o| sf.find(SF_CONSUMPTION_RATE, o.Id) }

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    tiers_mode = PriceHelpers.transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.Type)

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # for convenience, add a link to the consumption schedule in metadata
      "metadata" => stripe_metadata_for_sf_object(consumption_schedule),
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
        message: "Decimal value provided for tier bound. Ensure all tier bounds are integers."
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
      # this should never happen
      raise "unexpected pricing method #{pricing_method}"
    end

    {
      'up_to' => up_to,
      pricing_key => normalize_float_amount_for_stripe(price.to_s, @user, as_decimal: true),
    }
  end

  sig { params(sf_order_line: Restforce::SObject).returns(Hash) }
  def extract_tiered_price_params_from_order_line(sf_order_line)
    # TODO use cache_service
    consumption_schedules = sf.query(
      <<~EOL
        SELECT Id
        FROM #{CPQ_CONSUMPTION_SCHEDULE}
        WHERE SBQQ__OrderItem__c = '#{sf_order_line[SF_ID]}'
      EOL
    ).map {|o| sf.find(CPQ_CONSUMPTION_SCHEDULE, o.Id) }

    if consumption_schedules.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if consumption_schedules.count > 1
      raise "should not be more than one consumption schedule associated with an order line"
    end

    # the CPQ consumption schedule is materially different from the standard consumption schedule, so we can't do a drop-in replacement

    consumption_schedule = consumption_schedules.first

    if consumption_schedule.IsDeleted
      log.warn 'consumption schedule is deleted, ignoring'
      return {}
    end

    # unlike pricebook consumption schedule, this consumption schedule does not have a IsActive field

    # In our CPQ testing, "Tier" is the only valid picklist value, so we do not expect any other values
    if consumption_schedule.SBQQ__RatingMethod__c != "Tier"
      raise "unexpected rating method #{consumption_schedule.SBQQ__RatingMethod__c}"
    end

    # TODO use cache_service
    consumption_rates = sf.query(
      <<~EOL
        SELECT Id
        FROM #{CPQ_CONSUMPTION_RATE}
        WHERE #{CPQ_CONSUMPTION_SCHEDULE} = '#{consumption_schedule.Id}'
      EOL
    ).map {|o| sf.find(CPQ_CONSUMPTION_RATE, o.Id) }

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    tiers_mode = PriceHelpers.transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.SBQQ__Type__c)

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # for convenience, add a link to the consumption schedule in metadata
      "metadata" => stripe_metadata_for_sf_object(consumption_schedule),
    }
  end

  # sf_product is required because for some fields the required value is pulled from the product
  # TODO this should be modified to be more flexible and pull via the mapper instead https://jira.corp.stripe.com/browse/PLATINT-1485
  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject).returns(Stripe::Price) }
  def generate_price_params_from_sf_object(sf_object, sf_product)
    # this should never happen, but provides self-documentation and extra test guards
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise "price can only be created from an order line or pricebook entry"
    end

    # TODO the tiered pricing logic should be extracted out to a separate method

    # generate these params first since it is hard to merge these
    tiered_pricing_params = if sf_object.sobject_type == SF_PRICEBOOK_ENTRY
      extract_tiered_price_params_from_pricebook_entry(sf_object)
    else
      extract_tiered_price_params_from_order_line(sf_object)
    end

    is_tiered_price = !tiered_pricing_params.empty?

    # omitting price param here, this should be defined upstream
    stripe_price = Stripe::Price.construct_from({
      # TODO most likely need to pass the order over? Can
      # TODO not seeing currency anywhere? This is only enab  led on certain accounts
      currency: Integrations::Utilities::Currency.base_currency_iso(@user),

      # TODO using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027
    }.merge(tiered_pricing_params))

    # the params are extracted here *and* we apply custom mappings because this enables the user to setup custom mappings for
    # *everything* we sent to the price API including UnitPrice and other fields which go through custom transformations
    price_params = extract_salesforce_params!(sf_object, Stripe::Price)
    mapper.assign_values_from_hash(stripe_price, price_params)

    if sf_object.sobject_type == SF_PRICEBOOK_ENTRY
      apply_mapping(stripe_price, sf_object)
    else
      # https://jira.corp.stripe.com/browse/PLATINT-1486
      apply_mapping(stripe_price, sf_object, compound_key: true)
    end

    # although we are passing the amount as a decimal, the decimal amount still represents cents
    # to_s is used here to (a) satisfy typing requirements and (b) ensure BigDecimal can parse the float properly
    stripe_price.unit_amount_decimal = normalize_float_amount_for_stripe(stripe_price.unit_amount_decimal.to_s, @user, as_decimal: true)

    # TODO validate billing frequency and subscription term

    # if tiered pricing is set up, then we ignore any non-tiered pricing configuration
    if is_tiered_price
      is_tiered_price = PriceHelpers.sanitize_price_tier_params(stripe_price)
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

      # TODO it's possible that a custom mapping is defined for this value and it's an integer, we should support this case in the helper method
      # this represents how often the price is billed: i.e. if `interval` is month and `interval_count`
      # is 2, then this price is billed every two months.
      stripe_price.recurring[:interval_count] = PriceHelpers.transform_salesforce_billing_frequency_to_recurring_interval(stripe_price.recurring[:interval_count])

      # frequency: monthly or daily, defined on the CPQ
      stripe_price.recurring[:interval] = sf_cpq_term_interval
    else
      # TODO should we nil out any custom mapped recurring values? Let's wait and see if we get some errors
      # stripe_price.recurring = {}
    end

    # we should only transform licensed prices that are not customized
    if is_recurring_item && !is_tiered_price && sf_object.sobject_type == SF_ORDER_ITEM
      if !PriceHelpers.using_custom_order_line_price_field?(@user)
        log.info 'custom price not used, adjusting unit_amount_decimal', sf_order_item_id: sf_object['Id']

        billing_frequency = StripeForce::Utilities::StripeUtil.billing_frequency_of_price_in_months(stripe_price)

        # TODO this is a hack and we should really extract this through the extract params
        quote_subscription_term_path = 'Order.SBQQ__Quote__c.SBQQ__SubscriptionTerm__c'
        quote_subscription_term = mapper.extract_key_path_for_record(sf_object, quote_subscription_term_path)

        if quote_subscription_term.nil?
          raise Integrations::Errors::MissingRequiredFields.new(
            salesforce_object: sf_object,
            missing_salesforce_fields: [quote_subscription_term_path]
          )
        end

        # it's looking like these values are never really aligned and we should ignore the line item
        if sf_object['SBQQ__SubscriptionTerm__c'] == quote_subscription_term
          report_edge_case("subscription term on quote matches line item")
        end

        # TODO need to stop hardcoding this value!
        # NOTE in most cases, this value should be the same as `sf_object['SBQQ__ProrateMultiplier__c']` if the user has product-level subscription term enabled
        price_multiplier = determine_subscription_term_multiplier_for_billing_frequency(quote_subscription_term, billing_frequency)

        stripe_price.unit_amount_decimal = stripe_price.unit_amount_decimal / price_multiplier
      end
    end

    stripe_price
  end
end
