# frozen_string_literal: true
# typed: true
class StripeForce::Translate
  include Integrations::Log
  include Integrations::ErrorContext
  include StripeForce::Constants

  def self.perform(user:, sf_object:, locker:)
    interactor = self.new(user, locker)
    interactor.translate(sf_object)
  end

  def initialize(user, locker)
    @user = user
    @locker = locker

    set_error_context(user: user)
  end

  def sf
    @user.sf_client
  end

  def translate(sf_object)
    # assume order
    translate_order(sf_object)
  end

  def translate_order(sf_object)
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
    sf_product = sf.find('Product2', sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find('PricebookEntry', sf_order_item.PricebookEntryId)

    # have we already pushed this to Stripe?
    stripe_price_id = sf_pricebook_entry[PRICE_BOOK_STRIPE_ID]
    if !stripe_price_id.nil?
      log.info 'price already pushed, retrieving from stripe', stripe_resource_id: stripe_price_id
      begin
        return Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)
      rescue => e
        # TODO report this
      end
    end

    if sf_pricebook_entry.UnitPrice != sf_order_item.UnitPrice
      raise 'unit prices should not be different'
    end

    # TODO the pricebook entry is really a 'suggested' price and it can be overwritten by the quote or order line
    # TODO if the list price is used, we shoudl try to create a price for this in Stripe, otherwise we'll create a price and use that

    unless integer_quantity?(sf_order_item)
      # TODO need to think about taxes calculated in SF at this point
      raise 'implement calculating net price on item'
    end

    # TODO check for float quantities

    optional_params = {}

    if recurring_item?(sf_order_item)
      optional_params[:recurring] = {
        interval: sf_cpq_term_interval,
      }
    end

    log.info 'creating price in stripe', sf_resource: sf_pricebook_entry

    price = Stripe::Price.create({
      # TODO there is an undocumented id param, but it doesn't seem to work
      # id: sf_pricebook_entry.Id,

      product: product,

      # TODO not seeing currency anywhere
      unit_amount_decimal: sf_pricebook_entry.UnitPrice,
      currency: 'usd',

      # using a `lookup_key` here would allow users to easily update prices
      # https://jira.corp.stripe.com/browse/RUN_COREMODELS-1027

      # by omitting the recurring params it sets `type: one_time`
      # you cannot set this param directly

      # recurring: {interval: 'month'},

      metadata: sf_metadata(sf_pricebook_entry),
    }.merge(optional_params), @user.stripe_credentials)

    sf.update!('PricebookEntry', 'Id' => sf_pricebook_entry.Id, PRICE_BOOK_STRIPE_ID => price.id)

    price
  end

  def recurring_item?(sf_order_item)
    # TODO unsure why ChargeType is nil?
    !sf_order_item.SBQQ__ChargeType__c.nil? || !sf_order_item.SBQQ__SubscriptionType__c.nil?
  end

  def integer_quantity?(sf_order_item)
    log.warn 'user has float quantities defined on the line level'

    # TODO this is naive
    sf_order_item.Quantity.to_i == sf_order_item.Quantity.to_f
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

    sf_recurring_items, sf_one_time_items = sf_order_items.partition {|i| recurring_item?(i) }
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

    sf_params = extract_salesforce_params!(sf_quote, {
      start_date: CPQ_QUOTE_SUBSCRIPTION_START_DATE,
      subscription_iterations: CPQ_QUOTE_SUBSCRIPTION_TERM,
    })

    if is_recurring_order
      log.info 'recurring items found, creating subscription schedule'

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
    result_params = param_mapping.each_with_object({}) do |(k, sf_key), h|
      h[k] = sf_record[sf_key]
    end

    missing_fields = result_params.select {|_k, v| v.nil? }

    if missing_fields.present?
      raise Integrations::Errors::MissingRequiredFields.new("missing required params #{missing_fields}")
    end

    result_params
  end
end
