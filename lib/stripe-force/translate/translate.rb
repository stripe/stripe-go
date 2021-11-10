# frozen_string_literal: true
# typed: true

class StripeForce::Translate
  extend T::Sig

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities::Currency

  include StripeForce::Constants

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
    if sf_object.sobject_type == SF_ORDER
      translate_order(sf_object)
    else
      raise 'only order translation is supported right now'
    end
  end

  def translate_order(sf_object)
    log.info 'translating order'

    create_stripe_transaction_from_sf_order(sf_object)
  end

  def sf_metadata(sf_object)
    {
      salesforce_id: sf_object.Id,
      salesforce_url: "#{@user.sf_endpoint}/#{sf_object.Id}",
    }
  end

  def create_customer_from_sf_order(sf_order)
    # opportunity, instead of contact/customer, is used here since {customer, contact} pairs don't map to Stripe a customer`
    # each oppt can use a unique contact/customer pair, this feels like the best fit for now
    sf_quote = sf.find(CPQ_QUOTE, sf_order[CPQ_QUOTE])
    sf_opportunity = sf.find('Opportunity', sf_quote[CPQ_QUOTE_OPPORTUNITY])

    stripe_customer_id = sf_opportunity[OPPORTUNITY_STRIPE_ID]
    if !stripe_customer_id.nil?
      return Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    end

    sf_account = sf.find('Account', sf_order.AccountId)
    sf_primary_contact = sf.find('Contact', sf_quote[CPQ_QUOTE_PRIMARY_CONTACT])

    customer = Stripe::Customer.create({
      email: sf_primary_contact.Email,
      name: sf_account.Name,
      metadata: sf_metadata(sf_opportunity),
    }, @user.stripe_credentials)

    sf.update!('Opportunity', 'Id' => sf_opportunity.Id, OPPORTUNITY_STRIPE_ID => customer.id)

    customer
  end

  def create_product_from_sf_product(sf_product)
    begin
      return Stripe::Product.retrieve(sf_product.Id, @user.stripe_credentials)
    rescue Stripe::InvalidRequestError => e
      raise if e.code != 'resource_missing'
    end

    Stripe::Product.create({
      # TODO setting custom Ids may not be the best idea here
      id: sf_product.Id,
      name: sf_product.Name,
      description: sf_product.Description,
      metadata: sf_metadata(sf_product),
    }, @user.stripe_credentials)

    # TODO update SF ID
  end

  # TODO this is defined globally in SF and needs to be pulled dynamically
  def sf_cpq_term_interval
    'month'
  end

  # TODO what if the list price is updated in SF? We shouldn't probably create a new price object
  def create_price_for_order_item(sf_order_item)
    sf_product = sf.find(SF_PRODUCT, sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find('PricebookEntry', sf_order_item.PricebookEntryId)

    if sf_pricebook_entry.UnitPrice != sf_order_item.UnitPrice
      raise 'unit prices between pricebook and order item should not be different'
    end

    unless integer_quantity?(sf_order_item)
      # TODO need to think about taxes calculated in SF at this point
      raise 'float quantities are not yet supported'
    end

    if high_precision_float?(sf_order_item)
      # TODO this will probably require us to use `unit_price_decimal` on the price
      raise 'float prices with more than two decimals of precision are not supported'
    end

    # TODO shouldn't be `to_s`ing this here, need to determine if SF is sending this as a float or what
    unit_price_for_stripe = normalize_float_amount_for_stripe(sf_pricebook_entry.UnitPrice.to_s, @user)

    # have we already pushed this to Stripe?
    stripe_price_id = sf_pricebook_entry[PRICE_BOOK_STRIPE_ID]
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
        stripe_price = nil # rubocop:disable Lint/UselessAssignment
      end
    end

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we shoudl try to create a price for this in Stripe, otherwise we'll create a price and use that

    optional_params = {}

    if recurring_item?(sf_order_item)
      optional_params[:recurring] = {
        interval: sf_cpq_term_interval,
      }
    end

    log.info 'creating price in stripe', sf_resource: sf_pricebook_entry

    price = Stripe::Price.construct_from({
      # TODO there is an undocumented id param, but it doesn't seem to work
      # id: sf_pricebook_entry.Id,

      product: product,

      unit_amount: unit_price_for_stripe,

      # TODO most likely need to pass the order over? Can
      # TODO not seeing currency anywhere? This is only enabled on certain accounts
      currency: base_currency_iso(@user),

      # using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027

      # by omitting the recurring params it sets `type: one_time`
      # you cannot set this param directly

      # recurring: {interval: 'month'},

      metadata: sf_metadata(sf_pricebook_entry),
    }.merge(optional_params), @user.stripe_credentials)

    # TODO should probably figure out the best pattern here and wrap this
    apply_mapping(price, sf_pricebook_entry)
    price.dirty!
    price.save

    sf.update!('PricebookEntry', 'Id' => sf_pricebook_entry.Id, PRICE_BOOK_STRIPE_ID => price.id)

    price
  end

  def recurring_item?(sf_order_item)
    # TODO unsure why ChargeType is nil?
    !sf_order_item.SBQQ__ChargeType__c.nil? || !sf_order_item.SBQQ__SubscriptionType__c.nil?
  end

  def integer_quantity?(sf_order_item)
    # TODO this is naive, should investigate this logic further
    is_float_quantity = sf_order_item.Quantity.to_i == sf_order_item.Quantity.to_f

    if is_float_quantity
      log.warn 'user has float quantities defined on the line level'
    end

    is_float_quantity
  end

  def high_precision_float?(sf_order_item)
    # unfortunately the UnitPrice seems to come back as a float, so we need to `to_s`
    # TODO determine if Restforce is doing this translation or if it is us, we want the string float if we can
    (BigDecimal(sf_order_item.UnitPrice.to_s) * 100.0).to_i != normalize_float_amount_for_stripe(sf_order_item.UnitPrice.to_s, @user)
  end

  # TODO How can we organize code to support CPQ & non-CPQ use-cases? how can this be abstracted away from the order?
  def create_stripe_transaction_from_sf_order(sf_order)
    if sf_order.Status != 'Activated'
      log.info 'order is not activated, skipping'
      return
    end

    stripe_transaction_id = sf_order[ORDER_STRIPE_ID]
    if !stripe_transaction_id.nil?
      log.info 'order already translated', stripe_resource_id: stripe_transaction_id
      return
    end

    sf_quote = sf.find(CPQ_QUOTE, sf_order[CPQ_QUOTE])
    stripe_customer = create_customer_from_sf_order(sf_order)

    # https://github.com/Neuralab/WooCommerce-Salesforce-integration/blob/3af9ecf491f770ece9cdce993c21cee45f2647a0/includes/controllers/core/class-nwsi-salesforce-object-manager.php#L306
    # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
    # https://developer.salesforce.com/docs/atlas.en-us.api_placeorder.meta/api_placeorder/sforce_placeorder_rest_api_get_details_order.htm
    sf_order_with_lines = sf.api_get("commerce/sale/order/#{sf_order.Id}").body.first
    sf_order_items = sf_order_with_lines.OrderItems

    # the OrderItems from the commerce API is some sort of limited version
    sf_order_items = sf_order_items.map {|o| sf.find('OrderItem', o.Id) }

    sf_recurring_items, _sf_one_time_items = sf_order_items.partition {|i| recurring_item?(i) }
    is_recurring_order = !sf_recurring_items.empty?
    subscription_items = []

    stripe_prices = sf_order_items.map do |sf_order_item|
      price = create_price_for_order_item(sf_order_item)

      # if the quantity is specified as a float, we normalize it to 1 since Stripe does not support quantities
      quantity = integer_quantity?(sf_order_item) ? sf_order_item.Quantity.to_i : 1

      if !is_recurring_order
        Stripe::InvoiceItem.create({
          price: price,
          customer: stripe_customer,
          quantity: quantity,
        }, @user.stripe_credentials)
      else
        subscription_items << {
          price: price,
          quantity: quantity,
        }
      end
    end

    order_update_params = {}


    if is_recurring_order
      log.info 'recurring items found, creating subscription schedule'

      # TODO right now this just reports errors, but we should reference this for values in the future
      extract_salesforce_params!(sf_quote, {
        start_date: CPQ_QUOTE_SUBSCRIPTION_START_DATE,
        subscription_iterations: CPQ_QUOTE_SUBSCRIPTION_TERM,
      })

      # TODO subs in SF must always have an end date
      stripe_transaction = Stripe::SubscriptionSchedule.create({
        customer: stripe_customer,

        # TODO can we assume a consistent date format? What about TZs here?
        start_date: DateTime.parse(sf_quote[CPQ_QUOTE_SUBSCRIPTION_START_DATE]).to_time.to_i,
        end_behavior: 'cancel',

        phases: [
          {
            items: stripe_prices,
            # TODO we should ensure integer terms in SF
            # TODO is the restforce gem somehow formatting everything as a float?
            iterations: sf_quote[CPQ_QUOTE_SUBSCRIPTION_TERM].to_i,
          },
        ],

        # TODO mainly to avoid having to add a card right now, will need to map thi
        # collection_method: 'send_invoice',
        # days_until_due: 30

        metadata: sf_metadata(sf_order),
      }, @user.stripe_credentials)

      # TODO should we propogate the metadata down to the subscription?
    else
      log.info 'no recurring items found, creating a one-time invoice'

      stripe_transaction = Stripe::Invoice.create({
        customer: stripe_customer,
        metadata: sf_metadata(sf_order),

        collection_method: 'send_invoice',
        days_until_due: 30,
      }, @user.stripe_credentials)

      stripe_transaction.finalize_invoice

      order_update_params[ORDER_INVOICE_PAYMENT_LINK] = stripe_transaction.hosted_invoice_url
    end

    log.info 'stripe txn created', stripe_resource_id: stripe_transaction.id

    sf.update('Order', {'Id' => sf_order.Id, ORDER_STRIPE_ID => stripe_transaction.id}.merge(order_update_params))

    stripe_transaction
  end

  def extract_salesforce_params!(sf_record, param_mapping)
    result_params = param_mapping.transform_values do |sf_key|
      sf_record[sf_key]
    end

    missing_fields = result_params.select {|_k, v| v.nil? }

    if missing_fields.present?
      raise Integrations::Errors::MissingRequiredFields.new("missing required params #{missing_fields}")
    end

    result_params
  end

  # TODO allow for multiple records to be linked?
  def apply_mapping(record_to_map, source_record)
    @mapper ||= Integrations::Mapper.new(@user)
    @mapper.apply_mapping(record_to_map, source_record)
  end
end
