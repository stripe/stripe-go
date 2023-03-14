# typed: true
# frozen_string_literal: true

module Critic
  module StripeFactory
    extend T::Sig
    extend T::Helpers
    abstract!

    include Kernel
    include StripeForce::Constants
    include CommonHelpers

    sig { params(invoice_item_id: String).returns(Stripe::Event) }
    def get_invoice_item_event(invoice_item_id)
      # events can take some time to propogate
      events = T.let(nil, T.untyped)

      wait_until do
        events = Stripe::Event.list({
          object_id: invoice_item_id,
          type: 'invoiceitem.created',
        }, @user.stripe_credentials)

        events.count >= 1
      end

      assert_equal(1, events.count, "more than one event when we expected one")

      events.first
    end

    sig { params(type: String, obj: T.nilable(Stripe::StripeObject)).returns(Stripe::Event) }
    def create_event(type, obj=nil)
      obj ||= Stripe::Charge.construct_from(
        id: stripe_create_id(:ch)
      )

      Stripe::Event.construct_from({
        "id" => stripe_create_id(:evt),
        "created" => Time.now.getutc.to_i,
        "livemode" => false,
        "type" => type,
        "data" => {
          "object" => JSON.parse(obj.to_json),
        },
        "object" => "event",
        "pending_webhooks" => 0,
        "account" => stripe_create_id(:acct),
        "request" => stripe_create_id(:iar),
      })
    end

    def create_customer_with_subscription
      customer = create_customer_with_card
      product, price = create_price

      Stripe::Subscription.create({
        customer: customer.id,
        items: [
          {
            price: price.id,
          },
        ],
      }, @user.stripe_credentials)
    end

    def create_customer_with_card
      if @user.feature_enabled?(FeatureFlags::TEST_CLOCKS) && !@user.livemode
        test_clock = Stripe::TestHelpers::TestClock.create({
          frozen_time: Time.now.to_i,
        }, @user.stripe_credentials)
      end
      customer = create_customer(additional_fields: {test_clock: test_clock})

      payment_method = Stripe::PaymentMethod.attach(
        'pm_card_visa',
        {customer: customer.id},
        @user.stripe_credentials
      )
      customer.invoice_settings.default_payment_method = payment_method.id
      customer.save
      customer
    end

    def create_customer(additional_fields: {})
      email = create_random_email
      description = "Sample customer for Salesforce"

      customer = Stripe::Customer.create({
        description: description,
        email: email,
      }.merge(additional_fields), @user.stripe_credentials)
    end

    def create_price(additional_price_fields: {})
      stripe_product = Stripe::Product.create({
        name: sf_randomized_name("Product"),
      }, @user.stripe_credentials)

      stripe_price = Stripe::Price.create({
        product: stripe_product.id,
        currency: 'usd',
        unit_amount: 10_00,
        recurring: {
          interval: 'month',
        },
      }.merge(additional_price_fields), @user.stripe_credentials)

      [stripe_product, stripe_price]
    end

    # Creating an Invoice Rendering Templates is not yet supported
    # You can use the CreateInvoiceRenderingTemplate excelsior task (excl_NV8KZJ1238xCtc) to create a test template
    def create_invoice_rendering_template
       if ENV['CI']
        return "inrtem_1MlJnCC9fP1FVBtd3ts6JiPC"
       end

       # StripeForceDemo acct
       "inrtem_1Mk83SIsgf92XbAOfQinG8BA"
    end
  end
end
