require 'dotenv/load'

require 'rubygems'
require 'bundler'

Bundler.require(:default, :development)

Stripe.api_key = ENV['STRIPE_KEY']
Stripe.api_version = '2020-08-27'

# credentials from the CLI tool is available here: /Users/mbianco/.sfdx/*.json
# puts "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/authorize?response_type=code&client_id=#{ENV['SF_CONSUMER_KEY']}&redirect_uri=https://login.salesforce.com"
# oauth_code = "aPrxD8aRZGb_abrLry4Dl8fC.tyhGkTBCIWTjVKZ59yRhJssNw2ToltBW5zQYzM4JAE3ZeHDdA%3D%3D"
# "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com/services/oauth2/token?code=#{oauth_code}&grant_type=authorization_code&client_id=#{ENV['SF_CONSUMER_KEY']}&client_secret=#{ENV.fetch('SF_CONSUMER_SECRET')}&redirect_uri=https://login.salesforce.com"
# curl -X POST

def sf_endpoint
  "https://#{ENV.fetch('SF_INSTANCE')}.my.salesforce.com"
end

# https://help.salesforce.com/s/articleView?id=sf.remoteaccess_oauth_refresh_token_flow.htm&type=5
def refresh_token
  # POST /services/oauth2/token HTTP/1.1
  # Host: login.salesforce.com/
  # grant_type=refresh_token&
  # client_id=3MVG9lKcPoNINVBIPJjdw1J9LLM82HnFVVX19KY1uA5mu0QqEWhqKpoW3svG3XHrXDiCQjK1mdgAvhCscA9GE&client_secret=1955279925675241571&
  # refresh_token=your token here

  req = HTTPI::Request.new("#{sf_endpoint}/services/oauth2/token")

  req.body = {
    'client_id' => ENV.fetch('SF_CONSUMER_KEY'),
    'client_secret' => ENV.fetch('SF_CONSUMER_SECRET'),
    'refresh_token' => ENV.fetch('SF_REFRESH_TOKEN'),
    'grant_type' => 'refresh_token'
  }

  response = HTTPI.post(req)
  payload = JSON.parse(response.body)

  payload['access_token']
end


Restforce.configure do |config|
  config.log_level = :debug
  # config.log = true
end

# really? Can't set this on an instance or `configure` level?
Restforce.log = true

def salesforce_client
  Restforce.new(
    oauth_token: ENV.fetch('SF_ACCESS_TOKEN'),
    refresh_token: ENV.fetch('SF_REFRESH_TOKEN'),
    instance_url: sf_endpoint,
    client_id: ENV.fetch('SF_CONSUMER_KEY'),
    client_secret: ENV.fetch('SF_CONSUMER_SECRET'),

    # authentication_callback: Proc.new { |x| Rails.logger.debug x.to_s },

    # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_versions.htm
    # `http https://biancodevorg-dev-ed.my.salesforce.com/services/data`
    api_version: '52.0',

    log: true,
    log_level: :debug
  )
end

alias :sf :salesforce_client

Restforce::Data::Client.class_eval do

end

# Really? https://github.com/lsegal/yard/blob/9865620413d3519b561d87479e44f1b4fe782904/lib/yard/globals.rb#L20
# pry-doc loads yard and will clobber a global logger
# require 'yard'
include SimpleStructuredLogger
log = SimpleStructuredLogger::Writer.instance

def sf_metadata(sf_object)
  {
    salesforce_id: sf_object.Id,
    salesforce_url: "#{sf_endpoint}/#{sf_object.Id}"
  }
end

def example_sf_order
  # sf.find('Order', '8015e000000IJ1rAAG')

  # order with recurring item
  # sf.find('Order', '8015e000000IJDxAAO')

  sf.find('Order', '8015e000000IJF5AAO')
end

CPQ_QUOTE = 'SBQQ__Quote__c'.freeze
CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'.freeze
CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'.freeze
CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'.freeze
CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'.freeze
CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'.freeze
CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'.freeze

ORDER_STRIPE_ID = 'Stripe_Transaction_ID__c'.freeze
PRICE_BOOK_STRIPE_ID = 'Stripe_Price_ID__c'.freeze

def create_customer_from_sf_customer(sf_customer)
  # TODO pull contact information? Specifically the email?

  customer = Stripe::Customer.create({
    # TODO to be pulled from the contact
    email: "salesforce+#{Time.now.to_i}@stripe.com",

    name: sf_customer.Name,
    metadata: {
      salesforce_id: sf_customer.Id,
      salesforce_url: "#{sf_endpoint}/#{sf_customer.Id}"
    }
  })

  # TODO update ext ID in SF
end

def create_product_from_sf_product(sf_product)
  begin
    return Stripe::Product.retrieve(sf_product.Id)
  rescue Stripe::InvalidRequestError => e
    raise if e.code != 'resource_missing'
  end

  Stripe::Product.create(
    # TODO setting custom Ids may not be the best idea here
    id: sf_product.Id,
    name: sf_product.Name,
    description: sf_product.Description,
    metadata: sf_metadata(sf_product)
  )

  # TODO update SF ID
end

def create_price_for_order_item(sf_order_item)
  sf_product = sf.find('Product2', sf_order_item.Product2Id)
  product = create_product_from_sf_product(sf_product)

  sf_pricebook_entry = sf.find('PricebookEntry', sf_order_item.PricebookEntryId)

  # have we already pushed this to Stripe?
  stripe_price_id = sf_pricebook_entry[PRICE_BOOK_STRIPE_ID]
  if !stripe_price_id.nil?
    log.info 'price already pushed, retrieving from stripe'
    return Stripe::Price.retrieve(stripe_price_id)
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
  }.merge(optional_params))

  sf.update('PricebookEntry', Id => sf_pricebook_entry.Id, PRICE_BOOK_STRIPE_ID => price.id)

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

    })
  else
    stripe_transaction = Stripe::Invoice.create({
      customer: stripe_customer,
      metadata: sf_metadata(sf_order)
    })
  end

  log.info 'stripe txn created', stripe_resource_id: stripe_transaction.id

  sf.update('Order', 'Id' => sf_order.Id, ORDER_STRIPE_ID => stripe_transaction.id)

  stripe_transaction
end

def poll_orders
  updated_orders = sf.get_updated('Order', DateTime.now - 1, DateTime.now)["ids"]
  updated_orders.each do |sf_order_id|
    sf_order = sf.find('Order', sf_order_id)
    create_stripe_transaction_from_sf_order(sf_order)
  end
end

# sf.authenticate! will refresh oauth tokens

create_stripe_transaction_from_sf_order(example_sf_order)

# create_salesforce_order
