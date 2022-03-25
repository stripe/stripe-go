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
  end

  def translate_order(sf_object)
    log.info 'translating order'

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

    compound_external_id = "#{@origin_salesforce_object.Id}-#{salesforce_object.Id}"

    sf.create!(prefixed_stripe_field(SYNC_RECORD), {
      SyncRecordFields::COMPOUND_ID.serialize => compound_external_id,

      SyncRecordFields::PRIMARY_RECORD_ID.serialize => @origin_salesforce_object.Id,
      SyncRecordFields::PRIMARY_OBJECT_TYPE.serialize => @origin_salesforce_object.sobject_type,

      SyncRecordFields::SECONDARY_RECORD_ID.serialize => salesforce_object.Id,
      SyncRecordFields::SECONDARY_OBJECT_TYPE.serialize => salesforce_object.sobject_type,

      SyncRecordFields::RESOLUTION_MESSAGE.serialize => message,
      SyncRecordFields::RESOLUTION_STATUS.serialize => 'Error',
    }.transform_keys!(&method(:prefixed_stripe_field)))

    log.error 'translation failed', {
      metric: 'error.user',
      stripe_resource: stripe_resource,
      secondary_salesforce_id: salesforce_object.Id,
      secondary_salesforce_type: salesforce_object.sobject_type,
      error_message: message,
    }
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

  def create_stripe_object(stripe_class, sf_object, additional_stripe_params: {})
    stripe_key = stripe_class.to_s.demodulize.underscore
    default_mappings = @user.default_mappings[stripe_key]

    if default_mappings.nil?
      raise 'expected mappings but they were nil'
    end

    stripe_fields = extract_salesforce_params!(sf_object, default_mappings)

    stripe_object = stripe_class.construct_from(additional_stripe_params.merge(stripe_fields))
    stripe_object.metadata = stripe_metadata_for_sf_object(sf_object)

    apply_mapping(stripe_object, sf_object)

    log.info 'creating stripe object', salesforce_object: sf_object

    # there's a decent chance this causes us issues down the road: we shouldn't be using `construct_from`
    # and then converting the finalized object into a parameters hash. However, without using `construct_from`
    # we'd have to pass the object type around when mapping which is equally as bad.
    stripe_object.dirty!
    stripe_class.create(stripe_object.serialize_params, @user.stripe_credentials)
  end

  def retrieve_from_stripe(stripe_class, sf_object)
    begin
      stripe_class.retrieve(sf_object.Id, @user.stripe_credentials)
    rescue Stripe::InvalidRequestError => e
      raise if e.code != 'resource_missing'
      nil
    end
  end

  def create_customer_from_sf_account(sf_account)
    if (stripe_customer = retrieve_from_stripe(Stripe::Customer, sf_account))
      return stripe_customer
    end

    customer = create_stripe_object(Stripe::Customer, sf_account)

    update_sf_stripe_id(sf_account, customer)

    customer
  end

  def create_product_from_sf_product(sf_product)
    if (stripe_product = retrieve_from_stripe(Stripe::Product, sf_product))
      return stripe_product
    end

    stripe_product = create_stripe_object(Stripe::Product, sf_product)

    update_sf_stripe_id(sf_product, stripe_product)

    stripe_product
  end

  # TODO this is defined globally in SF and needs to be pulled dynamically from our configuration
  def sf_cpq_term_interval
    @user.connector_settings['cpq_term_interval']
  end

  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  def create_price_for_order_item(sf_order_item)
    sf_product = sf.find(SF_PRODUCT, sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find(SF_PRICEBOOK_ENTRY, sf_order_item.PricebookEntryId)

    if sf_order_item[CPQ_PRODUCT_SUBSCRIPTION_TERM] != sf_order_item[CPQ_PRODUCT_SUBSCRIPTION_TERM]
      raise 'subscription term differs between pricebook and order item'
    end

    # if sf_pricebook_entry.UnitPrice != sf_order_item.UnitPrice
    #   raise 'unit prices between pricebook and order item should not be different'
    # end

    # TODO should be able to use the pricebook entries for these checks

    unless integer_quantity?(sf_order_item)
      # TODO need to think about taxes calculated in SF at this point
      raise 'float quantities are not yet supported'
    end

    if high_precision_float?(sf_order_item)
      # TODO this will probably require us to use `unit_price_decimal` on the price
      raise 'float prices with more than two decimals of precision are not supported'
    end

    if metered_billing?(sf_order_item)
      raise "metered billing not yet supported"
    end

    # TODO detect billing type
    # https://help.salesforce.com/s/articleView?id=sf.cpq_order_product_fields.htm&type=5
    # if billing type = 'metered/areers'

    # TODO shouldn't be `to_s`ing this here, need to determine if SF is sending this as a float or what
    unit_price_for_stripe = normalize_float_amount_for_stripe(sf_pricebook_entry.UnitPrice.to_s, @user)

    # have we already pushed this to Stripe?
    stripe_price_id = sf_pricebook_entry[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    if stripe_price_id.present?
      log.info 'price already pushed, retrieving from stripe', stripe_resource_id: stripe_price_id

      stripe_price = begin
        Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)
      rescue => e
        report_exception(e)
        nil
      end

      # TODO if line item price is different than the stripe price, we need to create a new price
      # TODO we probably shouldn't overwrite the price on the SF object in this case, need to think this through
      if stripe_price && stripe_price.unit_amount != unit_price_for_stripe
        report_edge_case("specified price is different than stripe price", integration_record: sf_order_item, stripe_resource: stripe_price)
        stripe_price = nil # rubocop:disable Lint/UselessAssignment:
      end
    end

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we shoudl try to create a price for this in Stripe, otherwise we'll create a price and use that

    optional_price_params = {}

    # by omitting the recurring params it sets `type: one_time`
    # this param cannot be set directly
    if recurring_item?(sf_order_item)
      recurring_price_params = extract_salesforce_params!(sf_order_item, {
        'interval_count' => CPQ_PRODUCT_SUBSCRIPTION_TERM,
      })

      interval_count = recurring_price_params['interval_count']

      # I don't think this can actually occur, but let's keep an eye out
      if !is_integer_value?(interval_count)
        raise 'intervals need to be floats'
      end

      optional_price_params[:recurring] = {
        interval: sf_cpq_term_interval,
        # TODO maybe we should move the formatting stuff into the `extract_salesforce_params!`
        interval_count: interval_count.to_i,
      }
    end

    price = create_stripe_object(Stripe::Price, sf_pricebook_entry, additional_stripe_params: {
      product: product.id,

      unit_amount: unit_price_for_stripe,

      # TODO most likely need to pass the order over? Can
      # TODO not seeing currency anywhere? This is only enab  led on certain accounts
      currency: base_currency_iso(@user),

      # TODO using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027
    }.merge(optional_price_params))

    update_sf_stripe_id(sf_pricebook_entry, price)

    price
  end

  def metered_billing?(sf_order_item)
    sf_order_item[CPQ_PRODUCT_BILLING_TYPE] == CPQProductBillingTypeOptions::ARREARS.serialize
  end

  def recurring_item?(sf_order_item)
    # TODO unsure why ChargeType is nil? do we actually need to check this?
    !sf_order_item[CPQ_PRODUCT_CHARGE_TYPE].nil? || !sf_order_item[CPQ_PRODUCT_SUBSCRIPTION_TYPE].nil?
  end

  def integer_quantity?(sf_order_item)
    # TODO this is naive, should investigate this logic further
    is_integer_quantity = is_integer_value?(sf_order_item.Quantity)

    if is_integer_quantity
      log.warn 'user has float quantities defined on the line level'
    end

    is_integer_quantity
  end

  def is_integer_value?(value)
    value.to_i == value.to_f
  end

  def high_precision_float?(sf_order_item)
    # unfortunately the UnitPrice seems to come back as a float, so we need to `to_s`
    # TODO determine if Restforce is doing this translation or if it is us, we want the string float if we can
    (BigDecimal(sf_order_item.UnitPrice.to_s) * 100.0).to_i != normalize_float_amount_for_stripe(sf_order_item.UnitPrice.to_s, @user)
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

    # TODO should use batch service for this instead
    # https://github.com/Neuralab/WooCommerce-Salesforce-integration/blob/3af9ecf491f770ece9cdce993c21cee45f2647a0/includes/controllers/core/class-nwsi-salesforce-object-manager.php#L306
    # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
    # https://developer.salesforce.com/docs/atlas.en-us.api_placeorder.meta/api_placeorder/sforce_placeorder_rest_api_get_details_order.htm
    sf_order_with_lines = sf.api_get("commerce/sale/order/#{sf_order.Id}").body.first
    sf_order_items = sf_order_with_lines.OrderItems

    # the OrderItems from the commerce API is some sort of limited version
    sf_order_items = sf_order_items.map {|o| sf.find('OrderItem', o.Id) }

    invoice_items = []
    subscription_items = []

    sf_order_items.map do |sf_order_item|
      price = create_price_for_order_item(sf_order_item)

      # TODO we need to document this or just prevent this from happening
      # if the quantity is specified as a float, we normalize it to 1 since Stripe does not support quantities
      quantity = integer_quantity?(sf_order_item) ? sf_order_item.Quantity.to_i : 1

      if recurring_item?(sf_order_item)
        subscription_items << {
          price: price,
          quantity: quantity,
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
      # TODO at minimum we should pull this hash from the default_mappings
      extract_salesforce_params!(sf_quote, {
        start_date: CPQ_QUOTE_SUBSCRIPTION_START_DATE,
        subscription_iterations: CPQ_QUOTE_SUBSCRIPTION_TERM,
      })

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

  # TODO maybe build this into the mapper?
  # param_mapping: { stripe_key_name => salesforce_field_name }
  def extract_salesforce_params!(sf_record, param_mapping)
    result_params = param_mapping.transform_values do |sf_key|
      sf_record[sf_key]
    end

    missing_stripe_fields = result_params.select {|_k, v| v.nil? }

    if missing_stripe_fields.present?
      missing_salesforce_fields = missing_stripe_fields.keys.map {|k| param_mapping[k] }

      raise Integrations::Errors::MissingRequiredFields.new(
        salesforce_object: sf_record,
        missing_salesforce_fields: missing_salesforce_fields
      )
    end

    result_params
  end

  # TODO allow for multiple records to be linked?
  sig { params(record_to_map: Stripe::StripeObject, source_record: T.untyped).void }
  def apply_mapping(record_to_map, source_record)
    @mapper ||= Integrations::Mapper.new(@user)
    @mapper.apply_mapping(record_to_map, source_record)
  end
end
