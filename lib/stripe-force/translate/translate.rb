# frozen_string_literal: true
# typed: true

require_relative '../constants'
require_relative './utilities/metadata'
require_relative './utilities/stripe_util'
require_relative './utilities/salesforce_util'

require_relative './mapper'

class StripeForce::Translate
  extend T::Sig

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities::Currency

  include StripeForce::Constants
  include StripeForce::Utilities::Metadata
  include StripeForce::Utilities::StripeUtil
  include StripeForce::Utilities::SalesforceUtil

  # convenience method for console debugging
  def self.perform_inline(user, sf_id)
    locker = Integrations::Locker.new(user)

    translator = self.new(user, locker)

    sf_type = translator.salesforce_type_from_id(sf_id)
    sf_object = user.sf_client.find(sf_type, sf_id)

    result = translator.translate(sf_object)

    locker.clear_locked_resources

    result
  end

  def self.perform(user:, sf_object:, locker:)
    interactor = self.new(user, locker)
    interactor.translate(sf_object)
  end

  sig { params(user: StripeForce::User, locker: Integrations::Locker).void }
  def initialize(user, locker)
    @user = user
    @locker = locker

    set_error_context(user: user)
  end

  def sf
    @user.sf_client
  end

  def translate(sf_object)
    @origin_salesforce_object = sf_object

    set_error_context(user: @user, integration_record: sf_object)

    case sf_object.sobject_type
    when SF_ORDER
      translate_order(sf_object)
    when SF_PRODUCT
      translate_product(sf_object)
    when SF_PRICEBOOK_ENTRY
      create_price_from_pricebook(sf_object)
    when SF_ACCOUNT
      translate_account(sf_object)
    else
      raise 'unsupported translation type'
    end
  rescue Integrations::Errors::MissingRequiredFields => e
    create_user_failure(
      salesforce_object: e.salesforce_object,
      message: "The following required fields are missing from this Salesforce record: #{e.missing_salesforce_fields.join(', ')}",
    )

    raise
  rescue Restforce::ResponseError => e
    create_user_failure(
      # TODO can we indicate a more specific error object here?
      salesforce_object: @origin_salesforce_object,
      message: e.message
    )

    raise
  end

  # report user failures with the correct secondary SF object context
  def catch_stripe_api_errors(salesforce_object)
    yield
  rescue Stripe::InvalidRequestError => e
    # TODO probably remove this in the future, but I want to understand the error codes that are coming through
    log.warn 'stripe api error',
      code: e.code,
      message: e.message

    create_user_failure(
      salesforce_object: salesforce_object,
      message: "#{e.message} #{e.request_id}"
    )

    raise
  end

  def translate_order(sf_object)
    create_stripe_transaction_from_sf_order(sf_object)
  end

  def translate_product(sf_product)
    create_product_from_sf_product(sf_product)
  end

  def translate_account(sf_account)
    create_customer_from_sf_account(sf_account)
  end

  # TODO why is stripe_resource an argument here? Should we remove it?
  def create_user_failure(stripe_resource: nil, salesforce_object:, message:)
    if @origin_salesforce_object&.Id.blank?
      raise "origin salesforce object is blank, cannot record error"
    end

    log.error 'translation failed', {
      metric: 'error.user',
      stripe_resource: stripe_resource,
      secondary_salesforce_id: salesforce_object.Id,
      secondary_salesforce_type: salesforce_object.sobject_type,
      error_message: message,
    }

    compound_external_id = "#{@origin_salesforce_object.Id}-#{salesforce_object.Id}"

    log.debug 'creating sync record'

    # interestingly enough, if the external ID field does not exist we'll get a NOT_FOUND response
    # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_upsert.htm

    sf.upsert!(
      prefixed_stripe_field(SYNC_RECORD),
      prefixed_stripe_field(SyncRecordFields::COMPOUND_ID.serialize),
      {
        SyncRecordFields::COMPOUND_ID => compound_external_id,

        SyncRecordFields::PRIMARY_RECORD_ID => @origin_salesforce_object.Id,
        SyncRecordFields::PRIMARY_OBJECT_TYPE => @origin_salesforce_object.sobject_type,

        SyncRecordFields::SECONDARY_RECORD_ID => salesforce_object.Id,
        SyncRecordFields::SECONDARY_OBJECT_TYPE => salesforce_object.sobject_type,

        SyncRecordFields::RESOLUTION_MESSAGE => message,
        SyncRecordFields::RESOLUTION_STATUS => 'Error',
      }.transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field))
    )
  end

  # NOTE salesforce_object OR stripe_resource must be provided
  def throw_user_failure!(stripe_resource: nil, salesforce_object:, message:, error_class: nil)
    create_user_failure(
      stripe_resource: stripe_resource,
      salesforce_object: salesforce_object,
      message: message,
    )

    # TODO right now all attempt/audit log information is stored in salesforce
    #      it may make sense to change that at some point in the future

    error_class ||= Integrations::Errors::UserError

    raise error_class.new(
      message,

      stripe_resource: stripe_resource,
      salesforce_object: salesforce_object
    )
  end

  sig { params(sf_object: T.untyped, stripe_object: Stripe::APIResource, additional_salesforce_updates: Hash).void }
  def update_sf_stripe_id(sf_object, stripe_object, additional_salesforce_updates: {})
    stripe_id_field = prefixed_stripe_field(GENERIC_STRIPE_ID)

    if sf_object[stripe_id_field]
      log.info 'stripe id already exists on object, overwriting',
        old_stripe_id: sf_object[stripe_id_field],
        new_stripe_id: stripe_object.id,
        field_name: stripe_id_field
    end

    sf.update!(sf_object.sobject_type, {
      SF_ID => sf_object.Id,
      stripe_id_field => stripe_object.id,
    }.merge(additional_salesforce_updates))

    log.info 'updated with stripe id',
      salesforce_object: sf_object,
      stripe_id: stripe_object.id
  end

  def create_stripe_object(stripe_class, sf_object, additional_stripe_params: {}, skip_field_extraction: false)
    stripe_fields = if skip_field_extraction
      {}
    else
      extract_salesforce_params!(sf_object, stripe_class)
    end

    stripe_object = stripe_class.construct_from(additional_stripe_params)

    # the fields in the resulting hash could be dot-paths, so let's assign them using the mapper
    mapper.assign_values_from_hash(stripe_object, stripe_fields)

    stripe_object.metadata = stripe_metadata_for_sf_object(sf_object)

    apply_mapping(stripe_object, sf_object)

    if block_given?
      yield(stripe_object)
    end

    sanitize(stripe_object)

    log.info 'creating stripe object', salesforce_object: sf_object, stripe_object_type: stripe_class

    # there's a decent chance this causes us issues down the road: we shouldn't be using `construct_from`
    # and then converting the finalized object into a parameters hash. However, without using `construct_from`
    # we'd have to pass the object type around when mapping which is equally as bad.
    stripe_object.dirty!
    catch_stripe_api_errors(sf_object) do
      stripe_class.create(stripe_object.serialize_params, @user.stripe_credentials)
    end
  end

  sig { params(stripe_class: T.class_of(Stripe::APIResource), sf_object: Restforce::SObject).returns(T.nilable(Stripe::APIResource)) }
  def retrieve_from_stripe(stripe_class, sf_object)
    stripe_id = sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    return if stripe_id.nil?

    begin
      stripe_record = stripe_class.retrieve(sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

      # if this ID was provided to us by the user, the metadata does not exist in Stripe
      stripe_record_metadata = stripe_metadata_for_sf_object(sf_object)

      # `to_h` on the StripeObject symbolizes all of the keys
      # <= is a special `includes` operator https://stackoverflow.com/questions/7584801/how-to-check-if-an-hash-is-completely-included-in-another-hash
      unless stripe_record_metadata.symbolize_keys <= stripe_record.metadata.to_h
        stripe_record_metadata.each do |k, v|
          if stripe_record.metadata[k] != v
            log.warn 'overwriting metadata value',
              metadata_key: k,
              old_value: stripe_record.metadata[k],
              new_value: v
          end

          stripe_record.metadata[k] = v
        end

        log.info 'no metadata found on valid stripe record reference, adding'
        stripe_record.save
      end

      log.info 'existing stripe record found',
        salesforce_object: sf_object,
        stripe_resource: stripe_record

      stripe_record
    rescue Stripe::InvalidRequestError => e
      raise if e.code != 'resource_missing'
      nil
    end
  end

  def create_customer_from_sf_account(sf_account)
    log.info 'translating customer', salesforce_object: sf_account

    if (stripe_customer = retrieve_from_stripe(Stripe::Customer, sf_account))
      return stripe_customer
    end

    customer = create_stripe_object(Stripe::Customer, sf_account) do |generated_stripe_customer|
      # passing a partial shipping hash will trigger an error
      if !generated_stripe_customer.shipping.respond_to?(:address) || generated_stripe_customer.shipping.address.to_h.empty?
        log.info 'no address specified on shipping hash, removing'
        generated_stripe_customer.shipping = {}
      end
    end

    update_sf_stripe_id(sf_account, customer)

    customer
  end

  def create_product_from_sf_product(sf_product)
    log.info 'translating product', salesforce_object: sf_product

    if (stripe_product = retrieve_from_stripe(Stripe::Product, sf_product))
      return stripe_product
    end

    stripe_product = create_stripe_object(Stripe::Product, sf_product)

    update_sf_stripe_id(sf_product, stripe_product)

    stripe_product
  end

  def extract_tiered_price_params_from_pricebook_entry(sf_pricebook_entry)
    # TODO should be moved to the record service
    joining_records = sf.query("SELECT ConsumptionScheduleId FROM ProductConsumptionSchedule WHERE ProductId = '#{sf_pricebook_entry.Product2Id}'")

    if joining_records.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if joining_records.count > 1
      raise "should not be more than one"
    end

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

    # TODO should be moved to a helper in the record service
    consumption_rates = sf.query("SELECT Id FROM #{SF_CONSUMPTION_RATE} WHERE ConsumptionScheduleId = '#{consumption_schedule.Id}'").map {|o| sf.find(SF_CONSUMPTION_RATE, o.Id) }

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    tiers_mode = transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.Type)

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # for convenience, add a link to the consumption schedule in metadata
      "metadata" => stripe_metadata_for_sf_object(consumption_schedule),
    }
  end

  sig { params(raw_consumption_schedule_type: String).returns(String) }
  def transform_salesforce_consumption_schedule_type_to_tier_mode(raw_consumption_schedule_type)
    case raw_consumption_schedule_type
    when "Range"
      'volume'
    when "Slab"
      'graduated'
    else
      # should never happen
      raise "unexpected consumption schedule type #{raw_consumption_schedule_type}"
    end
  end

  sig { params(sf_consumption_rate: T.untyped).returns(Hash) }
  def transform_salesforce_consumption_rate_type_to_tier(sf_consumption_rate)
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

    up_to = if upper_bound.nil?
      'inf'
    else
      if Integrations::Utilities::StripeUtil.is_integer_value?(upper_bound)
        upper_bound.to_i
      else
        throw_user_failure!(
          salesforce_object: sf_consumption_rate,
          message: "Decimal value provided for tier bound. Ensure all tier bounds are integers."
        )
      end
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
    consumption_schedules = sf.query("SELECT Id FROM #{CPQ_CONSUMPTION_SCHEDULE} WHERE SBQQ__OrderItem__c = '#{sf_order_line[SF_ID]}'").map {|o| sf.find(CPQ_CONSUMPTION_SCHEDULE, o.Id) }

    if consumption_schedules.count.zero?
      return {}
    end

    # it does not look like this is programmatically enforced within CPQ, but should never happen
    if consumption_schedules.count > 1
      raise "should not be more than one"
    end

    # the CPQ consumption schedule is materially different from the standard consumption schedule, so we can't do a drop-in replacement

    consumption_schedule = consumption_schedules.first

    if consumption_schedule.IsDeleted
      log.warn 'consumption schedule is deleted, ignoring'
      return {}
    end

    # In our CPQ testing, "Tier" is the only valid picklist value, so we do not expect any other values
    if consumption_schedule.SBQQ__RatingMethod__c != "Tier"
      raise "unexpected rating method #{consumption_schedule.SBQQ__RatingMethod__c}"
    end

    # TODO should be moved to a helper in the record service
    consumption_rates = sf.query("SELECT Id FROM #{CPQ_CONSUMPTION_RATE} WHERE #{CPQ_CONSUMPTION_SCHEDULE} = '#{consumption_schedule.Id}'").map {|o| sf.find(CPQ_CONSUMPTION_RATE, o.Id) }

    pricing_tiers = consumption_rates.map do |consumption_rate|
      transform_salesforce_consumption_rate_type_to_tier(consumption_rate)
    end

    tiers_mode = transform_salesforce_consumption_schedule_type_to_tier_mode(consumption_schedule.SBQQ__Type__c)

    {
      "tiers" => pricing_tiers,
      "tiers_mode" => tiers_mode,
      "billing_scheme" => "tiered",

      # TODO dropping this because of key size limits: the CPQ object key size is greater than 40
      # for convenience, add a link to the consumption schedule in metadata
      # "metadata" => stripe_metadata_for_sf_object(consumption_schedule),
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
      currency: base_currency_iso(@user),

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
      # if non-tiered pricing fields are included Stripe will throw an error
      Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(stripe_price, :unit_amount_decimal)

      # TODO are there other pricing fields we should delete? It's unclear what the requirements are here?

      # API also requires pricing tiers to be sorted. Wat?
      # TODO https://jira.corp.stripe.com/browse/PLATINT-1573
      stripe_price.tiers.sort! do |a, b|
        # TODO can we use null instead of `inf`
        # a & b should never both be `inf`; this should be checked upstream
        if a.up_to == 'inf'
          1
        elsif b.up_to == 'inf'
          -1
        else
          a.up_to <=> b.up_to
        end
      end
    end

    # `recurring` could have been partially constructed by the mapper, which is why we use a symbol-based accessor
    # this avoids nil exception errors in various cases where a field has not been defined
    stripe_price[:recurring] ||= {}

    # by omitting the recurring params Stripe defaults to `type: one_time`
    # this param cannot be set directly
    is_recurring_item = recurring_item?(sf_object.sobject_type == SF_PRICEBOOK_ENTRY ? sf_product : sf_object)

    if is_recurring_item
      # NOTE on the price level, we are not concerned with the billing term
      #      that can be defined (as a default) on the product, but it only comes into play with the subscription schedule
      if sf_cpq_term_interval != 'month'
        raise 'only monthly terms are currently supported'
      end

      stripe_price.recurring[:usage_type] = transform_salesforce_billing_type_to_usage_type(stripe_price.recurring[:usage_type])

      # TODO it's possible that a custom mapping is defined for this value and it's an integer, we should support this case in the helper method
      # this represents how often the price is billed: i.e. if `interval` is month and `interval_count`
      # is 2, then this price is billed every two months.
      stripe_price.recurring[:interval_count] = transform_salesforce_billing_frequency_to_recurring_interval(stripe_price.recurring[:interval_count])

      # frequency: monthly or daily, defined on the CPQ
      stripe_price.recurring[:interval] = sf_cpq_term_interval
    else
      # TODO should we nil out any custom mapped recurring values? Let's wait and see if we get some errors
      # stripe_price.recurring = {}
    end

    # we should only transform licensed prices that are not customized
    if is_recurring_item && !is_tiered_price && sf_object.sobject_type == SF_ORDER_ITEM
      # TODO this is very naive: we need a better way of determining if the price field was customized
      # TODO we'll probably need some sort of feature flag for this as well
      order_line_price_key = mapper.mapping_key_for_record(Stripe::Price, sf_object)
      is_using_custom_price = !@user.field_defaults.dig(order_line_price_key, 'unit_amount_decimal').nil? ||
        (
          !@user.field_mappings.dig(order_line_price_key, 'unit_amount_decimal').nil? &&
          @user.field_mappings.dig(order_line_price_key, 'unit_amount_decimal') != @user.required_mappings.dig(order_line_price_key, 'unit_amount_decimal')
        )

      if !is_using_custom_price
        log.info 'custom price not used, adjusting unit_amount_decimal', sf_order_item_id: sf_object['Id']

        billing_frequency = billing_frequency_of_price_in_months(stripe_price)

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

  sig { params(raw_usage_type: T.nilable(String)).returns(String) }
  def transform_salesforce_billing_type_to_usage_type(raw_usage_type)
    # # https://help.salesforce.com/s/articleView?id=sf.cpq_order_product_fields.htm&type=5
    raw_usage_type ||= begin
      log.warn 'usage type not defined, defaulting to advanced (licensed)'
      CPQProductBillingTypeOptions::ADVANCE.serialize
    end

    case CPQProductBillingTypeOptions.try_deserialize(raw_usage_type)
    when CPQProductBillingTypeOptions::ADVANCE
      'licensed'
    when CPQProductBillingTypeOptions::ARREARS
      'metered'
    else
      raise "unexpected billing type"
    end
  end

  sig { params(raw_billing_frequency: T.nilable(String)).returns(Integer) }
  def transform_salesforce_billing_frequency_to_recurring_interval(raw_billing_frequency)
    raw_billing_frequency ||= begin
      log.warn 'interval_count not defined via mapping, using monthly fallback'
      CPQBillingFrequencyOptions::MONTHLY.serialize
    end

    # convert picklist description of frequency to month integers
    case CPQBillingFrequencyOptions.try_deserialize(raw_billing_frequency)
    when CPQBillingFrequencyOptions::MONTHLY
      1
    when CPQBillingFrequencyOptions::QUARTERLY
      3
    when CPQBillingFrequencyOptions::SEMIANNUAL
      6
    when CPQBillingFrequencyOptions::ANNUAL
      12
    else
      throw_user_failure!(
        # TODO including the exact line item reference, or specifying the context upstream, would be an improvement here
        salesforce_object: @origin_salesforce_object,
        message: "Unexpected billing frequency #{raw_billing_frequency}. Must use default CPQ billing frequencies."
      )
    end
  end

  # TODO this is defined globally in SF and needs to be pulled dynamically from our configuration
  def sf_cpq_term_interval
    # TODO move sanitization somewhere else
    @user.connector_settings[CONNECTOR_SETTING_CPQ_TERM_UNIT].downcase
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

  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  def create_price_for_order_item(sf_order_item)
    log.info 'translating price from a order line origin', salesforce_object: sf_order_item

    sf_product = sf.find(SF_PRODUCT, sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_order_item.PricebookEntryId)

    log.info 'product and pricebook references',
      salesforce_product_id: sf_product.Id,
      salesforce_pricebook_id: sf_pricebook_entry.Id

    # TODO should be able to use the pricebook entries for these checks

    # if the order line and pricebook entries are identical then we can use a pre-existing price
    # otherwise, we'll have to create a new price

    if pricebook_and_order_line_identical?(sf_pricebook_entry, sf_order_item, sf_product)
      log.info 'pricebook and product data is identical, attemping to use existing stripe price'
      existing_stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
      sf_target_for_stripe_price = sf_pricebook_entry
    else
      log.info 'pricebook and product data is different, creating new price from order line'
      sf_target_for_stripe_price = sf_order_item
    end

    if existing_stripe_price
      existing_stripe_price = T.cast(existing_stripe_price, Stripe::Price)
      generated_stripe_price = generate_price_params_from_sf_object(sf_pricebook_entry, sf_product)

      # this should never happen if our identical check is correct, unless the data in Salesforce is mutated over time
      if BigDecimal(existing_stripe_price.unit_amount_decimal.to_s) != BigDecimal(generated_stripe_price.unit_amount_decimal.to_s) ||
          existing_stripe_price.recurring.present? != generated_stripe_price.recurring.present? ||
          existing_stripe_price.recurring.interval != generated_stripe_price.recurring[:interval] ||
          existing_stripe_price.recurring.interval_count != generated_stripe_price.recurring[:interval_count]

        raise Integrations::Errors::UnhandledEdgeCase.new("expected generated prices to be equal, but they differed",
          metadata: HashDiff::Comparison.new(existing_stripe_price.to_hash, generated_stripe_price.to_hash).diff
        )
      end

      log.info 'using existing stripe price'
      return existing_stripe_price
    end

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

  # TODO this should be dynamic and pulled from the mapper
  # according to the salesforce documentation, if this field is non-nil ("empty") than it's a subscription item
  def recurring_item?(sf_object)
    if ![SF_ORDER_ITEM, SF_PRODUCT].include?(sf_object.sobject_type)
      raise "recurring definition is specified on the order item or product level only"
    end

    # TODO maybe pull this from the mapper in order to make this customizable?
    !sf_object[CPQ_QUOTE_SUBSCRIPTION_PRICING].nil?
  end

  # TODO we can probably remove this now that we have `unit_amount_decimal`
  def high_precision_float?(value)
    # unfortunately the UnitPrice seems to come back as a float, so we need to `to_s`
    # TODO determine if Restforce is doing this translation or if it is us, we want the string float if we can
    (BigDecimal(value.to_s) * 100.0).to_i != normalize_float_amount_for_stripe(value.to_s, @user)
  end

  # retrieves all line items, filters out skipped order lines
  def order_lines_from_order(sf_order)
    # TODO should move to cache service
    sf_order_items = sf.query("SELECT Id FROM #{SF_ORDER_ITEM} WHERE OrderID = '#{sf_order.Id}'").map(&:Id).map do |order_line_id|
      sf.find(SF_ORDER_ITEM, order_line_id)
    end

    sf_order_items.select do |sf_order_item|
      # never expect this to occur
      if sf_order_item.IsDeleted || !sf_order_item.SBQQ__Activated__c
        report_edge_case("order line is deleted or not activated")
      end

      should_keep = sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)].nil? ||
        !sf_order_item[prefixed_stripe_field(ORDER_LINE_SKIP)]

      if !should_keep
        log.info 'order line marked as skipped'
      end

      should_keep
    end
  end

  def phase_items_from_order_lines(sf_order_lines)
    invoice_items = []
    subscription_items = []

    sf_order_lines.map do |sf_order_item|
      price = create_price_for_order_item(sf_order_item)

      # could occur if a coupon is required for a negative amount, although this should probably be built into the `price` method instead
      next if price.nil?

      # a phase item is NOT a subscription item, but they are close, and right now the data mapper & revelant keys uses these across the codebase
      # I wonder if they will eventually converge in the public Stripe API over time?
      phase_item = Stripe::SubscriptionItem.construct_from({
        price: price.id,
        metadata: stripe_metadata_for_sf_object(sf_order_item),
      })

      phase_item_params = extract_salesforce_params!(sf_order_item, Stripe::SubscriptionItem)
      mapper.assign_values_from_hash(phase_item, phase_item_params)
      apply_mapping(phase_item, sf_order_item)

      # TODO add test case for this
      unless Integrations::Utilities::StripeUtil.is_integer_value?(phase_item.quantity)
        throw_user_failure!(
          salesforce_object: sf_order_item,
          message: "Quantity specified as a decimal value. Only integers are supported."
        )
      end

      # we know the quantity is not a float, so we can force it to an integer
      phase_item.quantity = phase_item.quantity.to_i

      # quantity cannot be specified if usage type is metered
      if price.recurring&.usage_type == 'metered'
        # setting quantity to null generates an empty string when `serialize_params` is called
        # which will throw an API error when this is passed to Stripe. This is why we need to use this hack
        Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
          phase_item,
          :quantity
        )
      end

      phase_item.dirty!

      # TODO this helper uses a hardcoded field to determine if the line is recurring or not
      if recurring_item?(sf_order_item)
        subscription_items << phase_item.serialize_params
      else
        # TODO metadata is currently not supported here https://jira.corp.stripe.com/browse/PLATINT-1609
        Integrations::Utilities::StripeUtil.delete_field_from_stripe_object(
          phase_item,
          :metadata
        )

        invoice_items << phase_item.serialize_params
      end
    end

    [invoice_items, subscription_items]
  end

  # TODO can we assume a consistent date format? What about TZs here?
  def salesforce_date_to_unix_timestamp(date_string)
    DateTime.parse(date_string).to_time.to_i
  end

  def billing_frequency_of_price_in_months(stripe_price)
    # for dev speed, let's assume everything is > monthly
    if %w{week day}.include?(stripe_price.recurring.interval)
      raise Integrations::Errors::UnhandledEdgeCase.new("unsupported price interval")
    end

    interval_in_months = case stripe_price.recurring.interval
    when 'month'
      1
    when 'year'
      12
    else
      raise Integrations::Errors::UnhandledEdgeCase.new("unexpected stripe pricing interval")
    end

    stripe_price.recurring.interval_count * interval_in_months
  end

  # service period and billing frequency are decoupled in CPQ
  # both values should be in months, but we want to support days in the future
  def determine_subscription_term_multiplier_for_billing_frequency(subscription_term, billing_frequency)
    # if specified iterations is less than the billing frequency of the stripe price then
    if subscription_term < billing_frequency
      # TODO we should probably create a invoice item price instead? Unsure of the best approach here, we would want the invoice item to be tied to a product?
      #      also, if we take this approach we should process all subscription items instead of just one
      log.info 'iterations is less than price billing frequency, creating a new price for the prorated amount', subscription_term: subscription_term, frequency: billing_frequency
      return 1
    end

    # TODO I expect this to happen in proration cases, i.e. 24mo subscription with a order amendment at mo 6
    if subscription_term % billing_frequency != 0
      throw_user_failure!(
        salesforce_object: @origin_salesforce_object,
        message: "Prorated order amendments are not yet supported"
      )
    end

    # TODO maybe this logic should be different if the term source is customized in salesforce?
    subscription_term / billing_frequency
  end

  def transform_iterations_by_billing_frequency(iterations, subscription_items)
    # this is terrible: the best way to figure out the recurrance schedule of a subscription
    # is to pull one of it's items and check the `recurring` hash. Why is this the case?
    # because of the level of field customization we allow, the values for the recurring hash
    # could be *anywhere* so the only alternative would be to regenerate the line items
    # or bubble up the recurrance through the line item helpers, which would mix responsibility
    # and make it more challenging to support multi-frequency subscriptions in the future

    price = Stripe::Price.retrieve(subscription_items.first[:price], @user.stripe_credentials)

    billing_frequency_in_months = billing_frequency_of_price_in_months(price)
    determine_subscription_term_multiplier_for_billing_frequency(iterations, billing_frequency_in_months)
  end

  # TODO How can we organize code to support CPQ & non-CPQ use-cases? how can this be abstracted away from the order?
  def create_stripe_transaction_from_sf_order(sf_order)
    log.info 'translating order', salesforce_object: sf_order

    if sf_order.Type != OrderTypeOptions::NEW.serialize
      raise Integrations::Errors::ImpossibleState.new("only new orders should be passed for transaction generation")
    end

    stripe_transaction = retrieve_from_stripe(Stripe::SubscriptionSchedule, sf_order)
    stripe_transaction ||= retrieve_from_stripe(Stripe::Invoice, sf_order)

    if !stripe_transaction.nil?
      log.info 'order already translated',
        stripe_transaction_id: stripe_transaction.id,
        salesforce_order_id: sf_order.Id
      return
    end

    sf_account = sf.find(SF_ACCOUNT, sf_order[SF_ORDER_ACCOUNT])

    stripe_customer = create_customer_from_sf_account(sf_account)

    sf_order_items = order_lines_from_order(sf_order)
    invoice_items, subscription_items = phase_items_from_order_lines(sf_order_items)
    is_recurring_order = !subscription_items.empty?

    order_update_params = {}

    if is_recurring_order
      log.info 'recurring items found, creating subscription schedule'

      subscription_params = extract_salesforce_params!(sf_order, Stripe::SubscriptionSchedule)

      # TODO should file an API papercut for this
      # when creating the subscription schedule the start_date must be specified on the heaer
      # when updating it, it is specified on the individual phase object
      subscription_params['start_date'] = salesforce_date_to_unix_timestamp(subscription_params['start_date'])


      # TODO this should really be done *before* generating the line items and therefore creating prices
      phase_iterations = transform_iterations_by_billing_frequency(
        # TODO is the restforce gem somehow formatting everything as a float? Or is this is the real value returned from SF?
        subscription_params.delete('iterations').to_i,
        subscription_items
      )

      # initial order, so there is only a single phase
      initial_phase = {
        add_invoice_items: invoice_items,
        items: subscription_items,
        iterations: phase_iterations,

        metadata: stripe_metadata_for_sf_object(sf_order),
      }

      # TODO add mapping support against the subscription schedule phase

      # TODO subs in SF must always have an end date
      stripe_transaction = Stripe::SubscriptionSchedule.construct_from({
        customer: stripe_customer.id,

        # TODO this should be specified in the defaults hash... we should create a defaults hash https://jira.corp.stripe.com/browse/PLATINT-1501
        end_behavior: 'cancel',

        # initial order will only ever contain a single phase
        phases: [initial_phase],

        metadata: stripe_metadata_for_sf_object(sf_order),
      })

      mapper.assign_values_from_hash(stripe_transaction, subscription_params)
      apply_mapping(stripe_transaction, sf_order)

      # TODO the idempotency key here is not perfect, need to refactor and use a job UID or something
      stripe_transaction.dirty!
      catch_stripe_api_errors(sf_order) do
        stripe_transaction = Stripe::SubscriptionSchedule.create(
          stripe_transaction.serialize_params,
          @user.stripe_credentials.merge(idempotency_key: sf_order[SF_ID])
        )
      end
    else
      log.info 'no recurring items found, creating a one-time invoice'

      # TODO there has got to be a way to include the lines on the invoice item create call
      invoice_items.each do |invoice_item_params|
        # TODO we should wrap these in `catch_stripe_api_errors`
        # TODO idempotency keys https://jira.corp.stripe.com/browse/PLATINT-1474
        Stripe::InvoiceItem.create({customer: stripe_customer}.merge(invoice_item_params), @user.stripe_credentials)
      end

      stripe_transaction = create_stripe_object(Stripe::Invoice, sf_order, additional_stripe_params: {
        customer: stripe_customer.id,
      })

      stripe_transaction.finalize_invoice

      order_update_params[prefixed_stripe_field(ORDER_INVOICE_PAYMENT_LINK)] = stripe_transaction.hosted_invoice_url
    end

    log.info 'stripe subscription or invoice created', stripe_resource_id: stripe_transaction.id

    update_sf_stripe_id(sf_order, stripe_transaction, additional_salesforce_updates: order_update_params)

    stripe_transaction
  end

  # param_mapping: { stripe_key_name => salesforce_field_name }
  sig { params(sf_record: Restforce::SObject, stripe_record_or_class: T.any(Class, Stripe::APIResource)).returns(Hash) }
  def extract_salesforce_params!(sf_record, stripe_record_or_class)
    stripe_mapping_key = mapper.mapping_key_for_record(stripe_record_or_class, sf_record)
    required_mappings = @user.required_mappings[stripe_mapping_key]

    if required_mappings.nil?
      raise "expected mappings for #{stripe_mapping_key} but they were nil"
    end

    # first, let's pull required mappings and check if there's anything missing
    required_data = mapper.build_dynamic_mapping_values(sf_record, required_mappings)

    missing_stripe_fields = required_mappings.select {|k, _v| required_data[k].nil? }

    if missing_stripe_fields.present?
      missing_salesforce_fields = missing_stripe_fields.keys.map {|k| required_mappings[k] }

      raise Integrations::Errors::MissingRequiredFields.new(
        salesforce_object: sf_record,
        missing_salesforce_fields: missing_salesforce_fields
      )
    end

    # then, let's extract optional fields and then merge them in
    default_mappings = @user.default_mappings[stripe_mapping_key]
    return required_data if default_mappings.blank?

    optional_data = mapper.build_dynamic_mapping_values(sf_record, default_mappings)

    required_data.merge(optional_data)
  end

  # TODO allow for multiple records to be linked?
  sig { params(record_to_map: Stripe::APIResource, source_record: Restforce::SObject, compound_key: T.nilable(T::Boolean)).void }
  def apply_mapping(record_to_map, source_record, compound_key: false)
    mapper.apply_mapping(record_to_map, source_record, compound_key: compound_key)
  end

  sig { returns(StripeForce::Mapper) }
  def mapper
    @mapper ||= StripeForce::Mapper.new(@user)
  end

  def sanitize(stripe_record)
    @sanitizer ||= StripeForce::Sanitizer.new(@user)
    @sanitizer.sanitize(stripe_record)
  end
end
