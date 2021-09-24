class StripeForce::Translate
  include SimpleStructuredLogger

  CPQ_QUOTE = 'SBQQ__Quote__c'.freeze
  CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'.freeze
  CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'.freeze
  CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'.freeze
  CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'.freeze
  CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'.freeze
  CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'.freeze

  ORDER_STRIPE_ID = 'Stripe_Transaction_ID__c'.freeze
  PRICE_BOOK_STRIPE_ID = 'Stripe_Price_ID__c'.freeze

  def self.perform(user:, sf_object:)
    interactor = self.new(user)
    interactor.translate(sf_object)
  end

  def initialize(user)
    @user = user
    # init_error_context
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
      salesforce_url: "#{@user.sf_endpoint}/#{sf_object.Id}"
    }
  end

  def create_customer_from_sf_customer(sf_customer)
    # TODO pull contact information? Specifically the email?

    customer = Stripe::Customer.create({
      # TODO to be pulled from the contact
      email: "salesforce+#{Time.now.to_i}@stripe.com",

      name: sf_customer.Name,
      metadata: sf_metadata(sf_customer)
    }, @user.stripe_credentials)

    # TODO update ext ID in SF
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
      metadata: sf_metadata(sf_product)
    }, @user.stripe_credentials)

    # TODO update SF ID
  end

  def create_price_for_order_item(sf_order_item)
    sf_product = sf.find('Product2', sf_order_item.Product2Id)
    product = create_product_from_sf_product(sf_product)

    sf_pricebook_entry = sf.find('PricebookEntry', sf_order_item.PricebookEntryId)

    # have we already pushed this to Stripe?
    stripe_price_id = sf_pricebook_entry[PRICE_BOOK_STRIPE_ID]
    if !stripe_price_id.nil?
      log.info 'price already pushed, retrieving from stripe', stripe_resource_id: stripe_price_id
      return Stripe::Price.retrieve(stripe_price_id, @user.stripe_credentials)
    end

    if sf_pricebook_entry.UnitPrice != sf_order_item.UnitPrice
      raise 'unit prices should not be different'
    end

    # TODO check for float quantities

    optional_params = {}

    if recurring_item?(sf_order_item)
      optional_params[:recurring] = {
        # TODO needs to dynamically pull from the recurring schedule
        interval: 'month'
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

      metadata: sf_metadata(sf_pricebook_entry)
    }.merge(optional_params), @user.stripe_credentials)

    sf.update('PricebookEntry', 'Id' => sf_pricebook_entry.Id, PRICE_BOOK_STRIPE_ID => price.id)

    price
  end

  def recurring_item?(sf_order_item)
    # TODO unsure why ChargeType is nil?
    !sf_order_item.SBQQ__ChargeType__c.nil? || !sf_order_item.SBQQ__SubscriptionType__c.nil?
  end

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

    # TODO should use opportunity instead here since {customer, contact} pairs don't map to Stripe
    sf_account_id = sf_order.AccountId
    sf_account = sf.find('Account', sf_account_id)
    stripe_customer = create_customer_from_sf_customer(sf_account)

    # https://github.com/Neuralab/WooCommerce-Salesforce-integration/blob/3af9ecf491f770ece9cdce993c21cee45f2647a0/includes/controllers/core/class-nwsi-salesforce-object-manager.php#L306
    # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
    # https://developer.salesforce.com/docs/atlas.en-us.api_placeorder.meta/api_placeorder/sforce_placeorder_rest_api_get_details_order.htm
    sf_order = sf.api_get("commerce/sale/order/#{sf_order.Id}").body.first
    sf_order_items = sf_order.OrderItems

    # the OrderItems from the commerce API is some sort of limited version
    sf_order_items = sf_order_items.map { |o| sf.find('OrderItem', o.Id) }

    sf_recurring_items, sf_one_time_items = sf_order_items.partition { |i| recurring_item?(i) }
    is_recurring_order = !sf_recurring_items.empty?
    subscription_items = []

    stripe_prices = sf_order_items.map do |sf_order_item|
      price = create_price_for_order_item(sf_order_item)

      if !is_recurring_order
        Stripe::InvoiceItem.create(
          customer: stripe_customer,
          price: price,
          # TODO does SF really allow floats for quantity? Really?
          quantity: sf_order_item.Quantity.to_i
        )
      else
        subscription_items << {
          price: price
        }
      end
    end

    if is_recurring_order
      stripe_transaction = Stripe::Subscription.create({
        customer: stripe_customer,
        items: stripe_prices,
        metadata: sf_metadata(sf_order),

        # TODO mainly to avoid having to add a card right now, will need to map thi
        collection_method: 'send_invoice',
        days_until_due: 30

      }, @user.stripe_credentials)
    else
      stripe_transaction = Stripe::Invoice.create({
        customer: stripe_customer,
        metadata: sf_metadata(sf_order)
      }, @user.stripe_credentials)
    end

    log.info 'stripe txn created', stripe_resource_id: stripe_transaction.id

    sf.update('Order', 'Id' => sf_order.Id, ORDER_STRIPE_ID => stripe_transaction.id)

    stripe_transaction
  end

end