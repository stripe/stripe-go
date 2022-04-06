# frozen_string_literal: true
# typed: true

require_relative './utilities/metadata'
require_relative './utilities/stripe_util'
require_relative './utilities/salesforce_util'

class StripeForce::Translate
  extend T::Sig

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities::Currency

  include StripeForce::Constants
  include StripeForce::Utilities::Metadata
  include StripeForce::Utilities::StripeUtil
  include StripeForce::Utilities::SalesforceUtil

  def self.salesforce_type_from_id(sf_id)
    case sf_id
    when /^01s/
      SF_PRICEBOOK
    when /^01t/
      SF_PRODUCT
    when /^01u/
      SF_PRICEBOOK_ENTRY
    when /^001/
      SF_ACCOUNT
    when /^003/
      SF_CONTACT
    when /^801/
      SF_ORDER
    when /^802/
      SF_ORDER_ITEM
    else
      # https://help.salesforce.com/s/articleView?id=000325244&type=1
      raise "unknown object type #{sf_id}"
    end
  end

  # convenience method for console debugging
  def self.sf_get(user, sf_id)
    sf_type = salesforce_type_from_id(sf_id)
    user.sf_client.find(sf_type, sf_id)
  end

  # convenience method for console debugging
  def self.perform_inline(user, sf_id)
    sf_type = salesforce_type_from_id(sf_id)
    sf_object = user.sf_client.find(sf_type, sf_id)
    locker = Integrations::Locker.new(user)

    result = self.perform(
      user: user,
      locker: locker,
      sf_object: sf_object
    )

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
  rescue Stripe::InvalidRequestError => e
    # TODO probably remove this in the future, but I want to understand the error codes that are coming through
    log.warn 'stripe error',
      code: e.code,
      message: e.message

    create_user_failure(
      salesforce_object: @origin_salesforce_object,
      message: e.message
    )

    raise
  rescue Restforce::ResponseError => e
    create_user_failure(
      salesforce_object: @origin_salesforce_object,
      message: e.message
    )

    raise
  end

  def translate_order(sf_object)
    log.info 'translating order', salesforce_object: sf_object

    create_stripe_transaction_from_sf_order(sf_object)
  end

  def translate_product(sf_product)
    create_product_from_sf_product(sf_product)
  end

  def translate_account(sf_account)
    create_customer_from_sf_account(sf_account)
  end

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

  # NOTE ns_record OR ns_class must be provided
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
      log.info 'stripe id already exists on object, overwritting',
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

    log.info 'creating stripe object', salesforce_object: sf_object, stripe_object_type: stripe_class

    # there's a decent chance this causes us issues down the road: we shouldn't be using `construct_from`
    # and then converting the finalized object into a parameters hash. However, without using `construct_from`
    # we'd have to pass the object type around when mapping which is equally as bad.
    stripe_object.dirty!
    stripe_class.create(stripe_object.serialize_params, @user.stripe_credentials)
  end

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

  # sf_product is required because for some fields the required value is pulled from the product
  # TODO this should be modified to be more flexible and pull via the mapper instead https://jira.corp.stripe.com/browse/PLATINT-1485
  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject).returns(Stripe::Price) }
  def generate_price_params_from_sf_object(sf_object, sf_product)
    # this should never happen, but provides self-documentation and extra test guards
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise "price can only be created from an order line or pricebook entry"
    end

    # omitting price param here, this should be defined upstream
    stripe_price = Stripe::Price.construct_from({
      # TODO most likely need to pass the order over? Can
      # TODO not seeing currency anywhere? This is only enab  led on certain accounts
      currency: base_currency_iso(@user),

      # TODO using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027
    })

    # the params are extracted here *and* we apply custom mappings because this enables the user to setup custom mappings for
    # *everything* we sent to the price API including UnitPrice and other fields which go through custom transformations
    price_params = extract_salesforce_params!(sf_object, Stripe::Price)
    mapper.assign_values_from_hash(stripe_price, price_params)

    # TODO this is not good: ideally we should have the product lookup defined in the mapping, but right now we don't allow arrays as the value of a mapping
    #      for now, let's hack this and just map the object twice from different perspectives
    #      https://jira.corp.stripe.com/browse/PLATINT-1486
    if sf_object.sobject_type == SF_PRICEBOOK_ENTRY
      mapper.assign_values_from_hash(stripe_price, mapper.build_dynamic_mapping_values(sf_product, {
        "recurring.interval_count" => CPQ_QUOTE_BILLING_FREQUENCY,
        "recurring.usage_type" => CPQ_PRODUCT_BILLING_TYPE,
      }))
    end

    # TODO we need to document this: the price is mapped to two unique objects
    apply_mapping(stripe_price, sf_object)

    # although we are passing the amount as a decimal, the decimal amount still represents cents
    # to_s is used here to (a) satisfy typing requirements and (b) ensure BigDecimal can parse the float properly
    stripe_price.unit_amount_decimal = normalize_float_amount_for_stripe(stripe_price.unit_amount_decimal.to_s, @user, as_decimal: true)

    # by omitting the recurring params it sets `type: one_time`
    # this param cannot be set directly
    if recurring_item?(sf_object.sobject_type == SF_PRICEBOOK_ENTRY ? sf_product : sf_object)
      # NOTE on the price level, we are not concerned with the billing term
      #      that can be defined (as a default) on the product, but it only comes into play with the subscription schedule
      if sf_cpq_term_interval != 'month'
        raise 'unsupported global term uniq'
      end

      # this can occur if interval_count and usage_type are both nil and are falling back on defaults
      if !stripe_price.respond_to?(:recurring)
        stripe_price.recurring = {
          interval_count: nil,
          usage_type: nil,
        }
      end

      # the "Billing Type" field mapped here by default is often empty, we use a nil value to default
      if !stripe_price.recurring.respond_to?(:usage_type)
        stripe_price.recurring.usage_type = nil
      end

      stripe_price.recurring.usage_type = transform_salesforce_billing_type_to_usage_type(stripe_price.recurring.usage_type)

      # TODO it's possible that a custom mapping is defined for this value and it's an integer, we should support this case in the helper method
      # this represents how often the price is billed: i.e. if `interval` is month and `interval_count`
      # is 2, then this price is billed every two months.
      stripe_price.recurring.interval_count = transform_salesforce_billing_frequency_to_recurring_interval(stripe_price.recurring.interval_count)

      # frequency: monthly or daily, defined on the CPQ
      stripe_price.recurring.interval = sf_cpq_term_interval
    else
      # TODO should we nil out any custom mapped recurring values? Let's wait and see if we get some errors
      # stripe_price.recurring = {}
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
      raise "unexpected billing frequency #{raw_billing_frequency}"
    end
  end

  # TODO this is defined globally in SF and needs to be pulled dynamically from our configuration
  def sf_cpq_term_interval
    # TODO move sanitization somewhere else
    @user.connector_settings[CONNECTOR_SETTING_CPQ_TERM_UNIT].downcase
  end

  # TODO maybe look at "can edit price" boolean on the product? But what about custom mappings?
  # TODO https://jira.corp.stripe.com/browse/PLATINT-1485
  def pricebook_and_order_line_identical?(sf_pricebook, sf_order_line, sf_product)
    # generate the full params that would be sent to the price object and compare them
    # if they aren't *exactly* the same, this will trigger a new price to be created
    # we'll probably need to do something a bit smarter here in the future but it's not obvious what
    # additional logic we should include here at the moment.
    generate_price_params_from_sf_object(sf_pricebook, sf_product).serialize_params !=
      generate_price_params_from_sf_object(sf_order_line, sf_product).serialize_params
  end

  def create_price_from_pricebook(sf_pricebook_entry)
    stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
    return stripe_price if stripe_price

    sf_product = sf.find(SF_PRODUCT, sf_pricebook_entry.Product2Id)
    stripe_product = create_product_from_sf_product(sf_product)

    stripe_price = create_price_from_sf_object(sf_pricebook_entry, sf_product, stripe_product)

    update_sf_stripe_id(sf_pricebook_entry, stripe_price)

    stripe_price
  end

  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  def create_price_for_order_item(sf_order_item)
    # TODO add test case for this
    unless integer_quantity?(sf_order_item)
      throw_user_failure!(
        salesforce_object: sf_order_item,
        message: "Quantity specified as a decimal value. Only integers are supported."
      )
    end

    log.info 'translating price via order line', salesforce_object: sf_order_item

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
      log.info 'pricebook and product data is different'
      sf_target_for_stripe_price = sf_order_item
    else
      log.info 'pricebook and product data is identical, attemping to use existing stripe price'
      existing_stripe_price = retrieve_from_stripe(Stripe::Price, sf_pricebook_entry)
      sf_target_for_stripe_price = sf_pricebook_entry
    end

    if existing_stripe_price
      generated_stripe_price = generate_price_params_from_sf_object(sf_pricebook_entry, sf_product)

      # this should never happen if our identical check is correct, unless the data in Salesforce is mutated over time
      if  BigDecimal(existing_stripe_price.unit_amount_decimal.to_s) != BigDecimal(generated_stripe_price.unit_amount_decimal.to_s) ||
          existing_stripe_price.recurring.interval != generated_stripe_price.recurring.interval ||
          existing_stripe_price.recurring.interval_count != generated_stripe_price.recurring.interval_count
        raise "expected generated prices to be equal, but they differed"
      end

      log.info 'using existing stripe price'
      return existing_stripe_price
    end

    stripe_price = create_price_from_sf_object(sf_target_for_stripe_price, sf_product, product)

    update_sf_stripe_id(sf_target_for_stripe_price, stripe_price)

    stripe_price
  end

  sig { params(sf_object: Restforce::SObject, sf_product: Restforce::SObject, stripe_product: Stripe::Product).returns(Stripe::Price) }
  def create_price_from_sf_object(sf_object, sf_product, stripe_product)
    if ![SF_ORDER_ITEM, SF_PRICEBOOK_ENTRY].include?(sf_object.sobject_type)
      raise "price can only be created from an order line or pricebook entry"
    end

    log.info 'creating price', salesforce_object: sf_object

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we shoudl try to create a price for this in Stripe, otherwise we'll create a price and use that
    stripe_price = generate_price_params_from_sf_object(sf_object, sf_product)

    # TODO https://jira.corp.stripe.com/browse/PLATINT-1483
    if stripe_price.unit_amount_decimal <= 0
      raise "negative prices should never be passed for item creation"
    end

    stripe_price.product = stripe_product.id
    stripe_price.metadata = stripe_metadata_for_sf_object(sf_object)

    # considered mapping SF pricebook ID to `lookup_key` but it's not *exactly* an external id and more presents a identifier
    # for an externall-used price so the "latest price" for a specific price-type can be used, probably in a website form or something
    # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027

    stripe_price.dirty!
    Stripe::Price.create(stripe_price.serialize_params, @user.stripe_credentials)
  end

  # according to the salesforce documentation, if this field is non-nil ("empty") than it's a subscription item
  def recurring_item?(sf_object)
    if ![SF_ORDER_ITEM, SF_PRODUCT].include?(sf_object.sobject_type)
      raise "recurring definition is specified on the order item or product level only"
    end

    # TODO maybe pull this from the mapper in order to make this customizable?
    !sf_object[CPQ_QUOTE_SUBSCRIPTION_PRICING].nil?
  end

  def integer_quantity?(sf_order_item)
    # TODO this is naive, should investigate this logic further
    is_integer_quantity = is_integer_value?(sf_order_item.Quantity)

    if !is_integer_quantity
      log.warn 'user has float quantities defined on the line level'
    end

    is_integer_quantity
  end

  def is_integer_value?(value)
    value.to_i == value.to_f
  end

  # TODO we can probably remove this now that we have `unit_amount_decimal`
  def high_precision_float?(value)
    # unfortunately the UnitPrice seems to come back as a float, so we need to `to_s`
    # TODO determine if Restforce is doing this translation or if it is us, we want the string float if we can
    (BigDecimal(value.to_s) * 100.0).to_i != normalize_float_amount_for_stripe(value.to_s, @user)
  end

  # TODO How can we organize code to support CPQ & non-CPQ use-cases? how can this be abstracted away from the order?
  def create_stripe_transaction_from_sf_order(sf_order)
    if sf_order.Status != OrderStatusOptions::ACTIVATED.serialize
      log.info 'order is not activated, skipping'
      return
    end

    if sf_order.Type != OrderTypeOptions::NEW.serialize
      raise "only initial orders are supported right now"
    end

    stripe_transaction_id = sf_order[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_transaction = stripe_object_from_id(stripe_transaction_id)

    if !stripe_transaction.nil? && [Stripe::SubscriptionSchedule, Stripe::Invoice].include?(stripe_transaction.class)
      log.info 'order already translated',
        stripe_resource_id: stripe_transaction_id,
        salesforce_order_id: sf_order.Id
      return
    end

    if !stripe_transaction.nil?
      log.info 'order contains incorrect stripe object reference, overwriting',
        stripe_resource_id: stripe_transaction_id,
        salesforce_order_id: sf_order.Id
    end

    sf_quote = sf.find(CPQ_QUOTE, sf_order[CPQ_QUOTE])
    sf_account = sf.find(SF_ACCOUNT, sf_order['AccountId'])

    stripe_customer = create_customer_from_sf_account(sf_account)

    # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api

    # TODO should move to cache service
    sf_order_items = sf.query("SELECT Id FROM OrderItem WHERE OrderID = '#{sf_order.Id}'").map(&:Id).map do |order_line_id|
      sf.find(SF_ORDER_ITEM, order_line_id)
    end

    invoice_items = []
    subscription_items = []

    sf_order_items.map do |sf_order_item|
      # TODO these need to be partitioned and created as discounts
      # TODO should respect customized price mapping here
      next if sf_order_item.UnitPrice <= 0

      price = create_price_for_order_item(sf_order_item)

      # TODO we need to document this or just prevent this from happening
      # if the quantity is specified as a float, we normalize it to 1 since Stripe does not support quantities
      quantity = integer_quantity?(sf_order_item) ? sf_order_item.Quantity.to_i : 1

      if recurring_item?(sf_order_item)
        subscription_items << {
          price: price,
          # quantity cannot be specified if usage type is metered
          quantity: price.recurring.usage_type == 'metered' ? nil : quantity,
        }
      else
        invoice_items << {
          price: price,
          quantity: quantity,
        }
      end
    end

    is_recurring_order = !subscription_items.empty?

    order_update_params = {}

    if is_recurring_order
      log.info 'recurring items found, creating subscription schedule'

      # TODO right now this just reports errors, but we should reference this for values in the future
      extract_salesforce_params!(sf_quote, Stripe::SubscriptionSchedule)

      # TODO subs in SF must always have an end date
      stripe_transaction = Stripe::SubscriptionSchedule.create({
        customer: stripe_customer,

        # TODO can we assume a consistent date format? What about TZs here?
        # TODO extract this out to a separate method to format a date for Stripe
        start_date: DateTime.parse(sf_quote[CPQ_QUOTE_SUBSCRIPTION_START_DATE]).to_time.to_i,

        # TODO this should be specified in the defaults hash... we should create a defaults hash
        end_behavior: 'cancel',

        # TODO we need to handle multiple items here
        phases: [
          # this is the initial order, so there is only a single phase
          # TODO if metadata is added to the subscription schedule phase items, we should add this to the data mapper
          {
            add_invoice_items: invoice_items,
            items: subscription_items,
            # TODO we should ensure integer terms in SF
            # TODO is the restforce gem somehow formatting everything as a float?
            iterations: sf_quote[CPQ_QUOTE_SUBSCRIPTION_TERM].to_i,
          },
        ],

        metadata: stripe_metadata_for_sf_object(sf_order),
      }, @user.stripe_credentials)

      # TODO should we propogate the metadata down to the subscription?

      # TODO should we conditionally do this based on user config?
      # TODO shluld we conditionally do this depending on if the user already has a payment method setup?

      # stripe_checkout_session = Stripe::Checkout::Session.create({
      #   # TODO these will need to specified by the customer
      #   success_url: 'https://example.com/success',
      #   cancel_url: 'https://example.com/cancel',

      #   customer: stripe_customer,
      #   mode: 'setup',
      #   payment_method_types: ['card'],
      # }, @user.stripe_credentials)

      # order_update_params[prefixed_stripe_field(ORDER_SUBSCRIPTION_PAYMENT_LINK)] = stripe_checkout_session.url
    else
      log.info 'no recurring items found, creating a one-time invoice'

      invoice_items.each do |invoice_item_params|
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
  def extract_salesforce_params!(sf_record, stripe_record)
    stripe_mapping_key = mapper.mapping_key_for_record(stripe_record)
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
  sig { params(record_to_map: Stripe::APIResource, source_record: Restforce::SObject).void }
  def apply_mapping(record_to_map, source_record)
    mapper.apply_mapping(record_to_map, source_record)
  end

  sig { returns(Integrations::Mapper) }
  def mapper
    @mapper ||= Integrations::Mapper.new(@user)
  end
end
