# typed: true
# frozen_string_literal: true

module StripeForce::Utilities
    # TODO @keshavc: Demo class that should be removed / modified after 08/15/2023
    module DemoUtil
      extend T::Sig
      extend T::Helpers
      abstract!

      include Kernel
      include StripeForce::Constants
      include Integrations::Log
      include StripeForce::Utilities::SalesforceUtil

      TEST_DEFAULT_CONTRACT_TERM = 12
      TEST_DEFAULT_PRICE = 120_00
      TEST_DEFAULT_STANDALONE_PRICE = 100_00

      @@user = nil
      def self.user
        @@user
      end

      def self.user=(new_user)
        @@user = new_user
      end

      def sf
        @@user.sf_client
      end

      sig { returns(String) }
      def now_time_formatted_for_salesforce
        format_date_for_salesforce(now_time)
      end

      sig { returns(String) }
      def now_time_in_future_formatted_for_salesforce
        format_date_for_salesforce(now_time_in_future)
      end

      def now_time
        DateTime.now.utc.beginning_of_day
      end

      def now_time_in_future
        DateTime.new(2033, 7, 1)
      end

      def format_date_for_salesforce(date)
        date.strftime("%Y-%m-%d")
      end

      sig { params(prefix: T.nilable(String)).returns(String) }
      def create_salesforce_id(prefix: nil)
        (prefix || "") + SecureRandom.alphanumeric(prefix ? 15 : 18)
      end

      def create_mock_salesforce_order
        id = create_salesforce_id(prefix: "802")
        Restforce::SObject.new({
          "attributes" => {"type" => SF_ORDER, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER}/#{id}"},
          "Id" => id,
          "IsDeleted" => false,
          SF_LAST_MODIFIED_DATE => "2022-09-20T17:53:31.000+0000",
        })
      end

      def create_mock_salesforce_order_item(additional_attributes: {})
        id = create_salesforce_id(prefix: "0Mh")
        Restforce::SObject.new({
          "Id" => id,
          "IsDeleted" => false,
          "attributes" => {"type" => SF_ORDER_ITEM, "url" => "/services/data/v52.0/sobjects/#{SF_ORDER_ITEM}/#{id}"},
        }.merge(additional_attributes))
      end

      def create_mock_salesforce_customer
        # TODO should move all prefixes to an enum
        id = create_salesforce_id(prefix: "003")
        Restforce::SObject.new({
          "attributes" => {"type" => "Account", "url" => "/services/data/v52.0/sobjects/Account/#{id}"},
          "IsDeleted" => false,
          "Id" => id,
        })
      end

      sig { params(additional_fields: Hash).returns(String) }
      def create_salesforce_account(additional_fields: {})
        _ = sf.create!(SF_ACCOUNT, {
          Name: sf_randomized_name(SF_ACCOUNT),
        }.merge(additional_fields))
      end

      def create_salesforce_contact(contact_email:, static_id: true)
        _ = sf.create!(SF_CONTACT, {
          LastName: 'Bianco',
          Email: static_id && contact_email.present? ? create_static_email(email: contact_email) : create_random_email,
        })
      end

      def create_salesforce_opportunity(sf_account_id:, currency_iso_code: nil, additional_fields: {})
        currency_iso_code ||= @user.connector_settings['default_currency']

        if @user.is_multicurrency_org?
          additional_fields = additional_fields.merge({
            StripeForce::Constants::SF_CURRENCY_ISO_CODE => currency_iso_code,
          })
        end

        sf.create!(SF_OPPORTUNITY, {
          Name: sf_randomized_name(SF_OPPORTUNITY),
          "CloseDate": now_time_formatted_for_salesforce,
          StageName: "Closed/Won",
          "AccountId": sf_account_id,
        }.merge(additional_fields))
      end

      def create_salesforce_price(sf_product_id: nil, price: nil, currency_iso_code: nil, additional_fields: {})
        price ||= TEST_DEFAULT_PRICE
        sf_product_id ||= create_salesforce_product
        currency_iso_code ||= @user.connector_settings['default_currency']

        if @user.is_multicurrency_org?
          additional_fields = additional_fields.merge({
            StripeForce::Constants::SF_CURRENCY_ISO_CODE => currency_iso_code,
          })
        end

        _ = sf.create!(SF_PRICEBOOK_ENTRY, {
          "Pricebook2Id" => default_pricebook_id,
          "Product2Id" => sf_product_id,
          "IsActive" => true,
          "UnitPrice" => price / 100.0,
          'UseStandardPrice' => false,
        }.merge(additional_fields))
      end

      def default_pricebook_id
        # https://help.salesforce.com/s/articleView?id=000326219&type=1
        standard_pricebook = sf.query("Select #{SF_ID} from #{SF_PRICEBOOK} where IsStandard = true").first

        if !standard_pricebook
          raise "could not find standard pricebook"
        end

        standard_pricebook.Id
      end

      def create_salesforce_product(static_id: true, additional_fields: {})
        sf.create!(SF_PRODUCT, {
          "Name" => "SaaS Subscription",
          'IsActive' => true,
          "Description" => "Demo subscription to show bidirectional syncing",
          'ProductCode' => static_id ? sf_static_id : sf_randomized_id,
        }.merge(additional_fields))
      end

      def create_salesforce_stripe_coupon_quote_association(sf_quote_id:, sf_stripe_coupon_id:)
       sf_stripe_coupon_id ||= create_salesforce_stripe_coupon

       sf.create!(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON_ASSOCIATION), {
         "Quote__c" => sf_quote_id,
         QUOTE_SF_STRIPE_COUPON => sf_stripe_coupon_id,
       }.transform_keys(&method(:prefixed_stripe_field)))
      end

      def create_salesforce_stripe_coupon_quote_line_association(sf_quote_line_id:, sf_stripe_coupon_id:)
        sf_stripe_coupon_id ||= create_salesforce_stripe_coupon

        # Our association objects unfortunately have currency fields in multicurrency orgs, so we must pass
        #   the coupon's currency to the association here.
        # This is error checked in the UI with this PR: https://github.com/stripe/stripe-salesforce/pull/1006
        currency_field = {}
        if @user.is_multicurrency_org?
          sf_stripe_coupon = sf_get(sf_stripe_coupon_id)
          currency_field = {
            StripeForce::Constants::SF_CURRENCY_ISO_CODE => sf_stripe_coupon[StripeForce::Constants::SF_CURRENCY_ISO_CODE],
          }
        end

        sf.create!(prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION), {
          "Quote_Line__c" => sf_quote_line_id,
          QUOTE_SF_STRIPE_COUPON => sf_stripe_coupon_id,
        }.merge(currency_field).transform_keys(&method(:prefixed_stripe_field)))
      end

      def create_salesforce_stripe_coupon(additional_fields: {})
        # we want to ensure that either amount_off or percent_off is set and not both
        amount_off = nil
        if additional_fields[SalesforceStripeCouponFields::PERCENT_OFF].nil?
            amount_off = TEST_DEFAULT_PRICE / 2
        end

        sf.create!(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), {
          SalesforceStripeCouponFields::NAME => sf_randomized_name(QUOTE_SF_STRIPE_COUPON),
          SalesforceStripeCouponFields::AMOUNT_OFF => amount_off,
          SalesforceStripeCouponFields::DURATION => 'once',
        }.merge(additional_fields).transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field)))
      end

      def get_sf_order_coupon_from_quote_coupon_id(quote_coupon_id)
        order_coupons = sf.query("SELECT #{SF_ID} FROM #{prefixed_stripe_field(ORDER_SF_STRIPE_COUPON)} WHERE #{prefixed_stripe_field('Quote_Stripe_Coupon_Id__c')} = '#{quote_coupon_id}'")

        # sanity check the order coupon info is the same as the quote coupon
        sf.find(prefixed_stripe_field(ORDER_SF_STRIPE_COUPON), order_coupons.first.Id)
      end

      def get_salesforce_stripe_coupons_associated_to_quote_line(quote_line_id:)
        quote_line_associations = sf.query("Select #{SF_ID} from #{prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION)} where #{prefixed_stripe_field('Quote_Line__c')} = '#{quote_line_id}'")

        if !quote_line_associations
          raise "could not find any stripe coupon quote line associations related to this quote line"
        end

        # there could be multiple coupons associated with a quote line
        coupons = quote_line_associations.map do |quote_line_association|
            association = sf.find(prefixed_stripe_field(QUOTE_LINE_SF_STRIPE_COUPON_ASSOCIATION), quote_line_association.Id)
            coupon_id = association[prefixed_stripe_field('Quote_Stripe_Coupon__c')]

            # return the coupon object
            sf.find(prefixed_stripe_field(QUOTE_SF_STRIPE_COUPON), coupon_id)
        end
        coupons
      end

      def salesforce_recurring_metered_produce_with_price(price_in_cents: nil)
        salesforce_recurring_product_with_price(
          price: price_in_cents,
          additional_product_fields: {
            CPQ_PRODUCT_BILLING_TYPE => CPQProductBillingTypeOptions::ARREARS.serialize,
          }
        )
      end

      # TODO `price` should be `price_in_cents`
      def salesforce_recurring_product_with_price(price: nil, currency_iso_code: nil, static_id: true, additional_product_fields: {})
        # I don't fully understand how the subscription term on the price iteracts with the billing frequency,
        # but if the term is set to a value which is different than the billing frequency it seems to use the
        # subscription term value. i.e. a yearly billed product
        subscription_term = if additional_product_fields.key?(CPQ_QUOTE_BILLING_FREQUENCY) && additional_product_fields[CPQ_QUOTE_BILLING_FREQUENCY] != CPQBillingFrequencyOptions::MONTHLY.serialize
          # Defaults to annually (12) if not set
          nil
        else
          1
        end

        # blanking out the subscription type ensures it is a one-time product
        product_id = create_salesforce_product(static_id: static_id, additional_fields: {
          # anything non-nil indicates subscription/recurring pricing
          CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

          CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::RENEWABLE,

          CPQ_QUOTE_SUBSCRIPTION_TERM => subscription_term,

          # one month
          CPQ_QUOTE_BILLING_FREQUENCY => CPQBillingFrequencyOptions::MONTHLY.serialize,
        }.merge(additional_product_fields))

        pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price, currency_iso_code: currency_iso_code)
        [product_id, pricebook_entry_id]
      end

      def salesforce_standalone_product_with_price(price: TEST_DEFAULT_STANDALONE_PRICE)
        # blanking out the subscription type ensures it is a one-time product
        product_id = create_salesforce_product(additional_fields: {
          CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::ONE_TIME,
        })

        pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price)

        [product_id, pricebook_entry_id]
      end

      def salesforce_evergreen_product_with_price(price: nil, currency_iso_code: nil, additional_product_fields: {})
        # evergreen products all have subscription term set to 1
        product_id = create_salesforce_product(additional_fields: {
          # anything non-nil indicates subscription/recurring pricing
          CPQ_QUOTE_SUBSCRIPTION_PRICING => 'Fixed Price',

          CPQ_PRODUCT_SUBSCRIPTION_TYPE => CPQProductSubscriptionTypeOptions::EVERGREEN,

          CPQ_QUOTE_SUBSCRIPTION_TERM => SF_ORDER_DEFAULT_EVERGREEN_SUBSCRIPTION_TERM,
        }.merge(additional_product_fields))

        pricebook_entry_id = create_salesforce_price(sf_product_id: product_id, price: price, currency_iso_code: currency_iso_code)
        [product_id, pricebook_entry_id]
      end

      # returns the full object, not the ID
      def create_subscription_order(sf_product_id: nil, sf_account_id: nil, currency_iso_code: nil, contact_email: sf_randomized_id, additional_fields: {})
        create_salesforce_order(
          sf_product_id: sf_product_id,
          sf_account_id: sf_account_id,
          currency_iso_code: currency_iso_code,
          contact_email: contact_email,
          additional_quote_fields: {
            CPQ_QUOTE_SUBSCRIPTION_START_DATE => now_time_formatted_for_salesforce,
            # one year / 12 months
            CPQ_QUOTE_SUBSCRIPTION_TERM => TEST_DEFAULT_CONTRACT_TERM,
          }.merge(additional_fields)
        )
      end

      def add_product_to_cpq_quote(quote_id, sf_product_id:, sf_pricebook_id: nil)
        sf_pricebook_id ||= default_pricebook_id

        # this error indicates that the SF account disconnected from the steelbrick REST service
        # reauth via: https://brick-rest.steelbrick.com/oauth/auth
        # Restforce::ErrorCode::ApexError: APEX_ERROR: SBQQ.RestClient.RefreshTokenNilException: Invalid nil argument: OAuth Refresh Token

        # https://developer.salesforce.com/docs/atlas.en-us.cpq_api_dev.meta/cpq_api_dev/cpq_api_read_quote.htm
        # get CPQ version of the quote
        quote_data = JSON.parse(sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body)

        cpq_product_representation = JSON.parse(sf.patch("services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.ProductAPI.ProductLoader&uid=#{sf_product_id}", {
          context: {
            # currencyCode:
            pricebookId: sf_pricebook_id,
          }.to_json,
        }).body)

        # TODO how can I specify quantity?
        # https://gist.github.com/paustint/bd18bd281134a180e014829b49ed043a
        quote_with_product = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteProductAdder', {
          context: {
            "quote": quote_data,
            "products": [
              cpq_product_representation,
            ],
            "groupKey": 0,
            "ignoreCalculate": false,
          }.to_json,
        }).body)

        quote_with_product
      end

      def calculate_and_save_cpq_quote(quote_data)
        log.info 'calculate and save quote'

        # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_calculate_final.htm
        calculated_quote = JSON.parse(sf.patch('/services/apexrest/SBQQ/ServiceRouter?loader=SBQQ.QuoteAPI.QuoteCalculator', {
          # "context": quote_with_product.to_json
          # "context": saved_quote.to_json
          # "context": sf.get("services/apexrest/SBQQ/ServiceRouter?reader=SBQQ.QuoteAPI.QuoteReader&uid=#{quote_id}").body
          "context": {"quote" => quote_data}.to_json,
        }).body)

        # https://developer.salesforce.com/docs/atlas.en-us.cpq_dev_api.meta/cpq_dev_api/cpq_quote_api_save_final.htm
        saved_quote = JSON.parse(sf.post('/services/apexrest/SBQQ/ServiceRouter', {
          "saver": "SBQQ.QuoteAPI.QuoteSaver",
          "model": calculated_quote.to_json,
        }).body)

        # TODO @keshavc, remove uncomment this line once the demo is finished
        # refute(saved_quote["calculationRequired"])

        saved_quote.dig("record", "Id")
      end

      # returns full SF order object
      sig { params(sf_quote_id: String, additional_order_fields: Hash).returns(T.untyped) }
      def create_order_from_cpq_quote(sf_quote_id, additional_order_fields: {})
        # it looks like there is additional field validation triggered here when `ordered` is set to true
        # If you recieve a "FIELD_CUSTOM_VALIDATION_EXCEPTION: Required fields are missing: [PricebookEntryId]" here,
        #   this means that SF could not find a pricebook entry for the product you are passing in. This may be because of a
        #   currency mismatch on the opportunity and the pricebook entry you created for the product.
        sf.update!(CPQ_QUOTE, {
          SF_ID => sf_quote_id,
          CPQ_QUOTE_ORDERED => true,
        })

        # TODO this is silly, I should do a simple SOQL query and grab the result
        # https://salesforce.stackexchange.com/questions/251904/get-sales-order-line-on-rest-api
        # TODO note that looking in the UI is the easiest way to get these magic relational values
        related_orders = sf.get("/services/data/v52.0/sobjects/#{CPQ_QUOTE}/#{sf_quote_id}/SBQQ__Orders__r")
        sf_order = related_orders.body.first

        sf.update!(SF_ORDER,
          {
            SF_ID => sf_order.Id,
            'Status' => 'Activated',
          }.merge(additional_order_fields)
        )

        sf_order.refresh
        sf_order
      end

      def create_salesforce_quote(sf_pricebook_id: nil, sf_account_id:, currency_iso_code: nil, contact_email:, additional_quote_fields: {})
        sf_pricebook_id ||= default_pricebook_id
        opportunity_id = create_salesforce_opportunity(sf_account_id: sf_account_id, currency_iso_code: currency_iso_code)
        contact_id = create_salesforce_contact(contact_email: contact_email)

        # you can create a quote without *any* fields, which seems completely silly
        _ = sf.create!(CPQ_QUOTE, {
          CPQ_QUOTE_PRIMARY => true,
          CPQ_QUOTE_OPPORTUNITY => opportunity_id,
          CPQ_QUOTE_PRIMARY_CONTACT => contact_id,
          CPQ_QUOTE_PRICEBOOK => sf_pricebook_id,
        }.merge(additional_quote_fields))
      end

      # https://github.com/sseixas/CPQ-JS
      def create_salesforce_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, currency_iso_code: nil, contact_email: sf_randomized_id, additional_quote_fields: {}, additional_order_fields: {})
        if !sf_product_id
          sf_product_id, _ = salesforce_recurring_product_with_price(currency_iso_code: currency_iso_code)
        end

        sf_account_id ||= create_salesforce_account

        quote_id = create_salesforce_quote(sf_account_id: sf_account_id, currency_iso_code: currency_iso_code, contact_email: contact_email, additional_quote_fields: additional_quote_fields)

        quote_with_product = add_product_to_cpq_quote(quote_id, sf_product_id: sf_product_id)

        calculate_and_save_cpq_quote(quote_with_product)

        create_order_from_cpq_quote(quote_id, additional_order_fields: additional_order_fields)

        # contract_id = salesforce_client.create!('Contract', accountId: account_id)
        # order_id = salesforce_client.create!('Order', {Status: "Draft", EffectiveDate: "2021-09-21", AccountId: account_id, ContractId: contract_id})
      end

      def create_demo_evergreen_order(sf_product_id: nil, sf_account_id: nil, sf_pricebook_id: nil, currency_iso_code: nil, additional_quote_fields: {}, user: @user, additional_order_fields: {})
        log.info "CREATE_DEMO_EVERGREEN_ORDER"
        original_user = @user
        @user = user
        if !sf_product_id
          sf_product_id, _ = salesforce_evergreen_product_with_price(currency_iso_code: currency_iso_code)
        end
        create_salesforce_order(sf_product_id: sf_product_id, sf_account_id: sf_account_id, sf_pricebook_id: sf_pricebook_id, currency_iso_code: currency_iso_code, additional_quote_fields: additional_quote_fields, contact_email: sf_randomized_id, additional_order_fields: additional_order_fields)
        @user = original_user
      end

      def sf_randomized_name(sf_object_name)
        # node_identifier = ENV['CIRCLE_NODE_INDEX'] || ""
        node_identifier = ""
        "REST #{sf_object_name} #{node_identifier} #{now_time}"
      end

      def sf_randomized_id
        random_id = SecureRandom.alphanumeric(29)

        if ENV['CIRCLE_NODE_INDEX']
          random_id = "#{random_id}#{ENV['CIRCLE_NODE_INDEX']}"
        end

        random_id
      end

      sig { params(email: String).returns(String) }
      def create_static_email(email: "")
        "#{email}@example.com"
      end

      sig { returns(String) }
      def create_random_email
        "#{sf_randomized_id}@example.com"
      end

      def sf_static_id
        static_id = "EfxBsBKAsjaF0caCUQPWQYJ8h61G"

        # if ENV['CIRCLE_NODE_INDEX']
        #   static_id = "#{static_id}#{ENV['CIRCLE_NODE_INDEX']}"
        # end

        static_id
      end

      def sf_get(sf_id)
        sf_id = sf_id.
          # remove whitespace
          strip.
          # remove quotes, sometimes easy to pull in when copy/pasting
          gsub("'\"", '')

        sf_type = StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_id)
        _ = @user.sf_client.find(sf_type, sf_id)
      end
    end
end
