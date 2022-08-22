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
      customer = create_customer
      payment_method = Stripe::PaymentMethod.attach(
        'pm_card_visa',
        {customer: customer.id},
        @user.stripe_credentials
      )
      customer.invoice_settings.default_payment_method = payment_method.id
      customer.save
      customer
    end

    def create_customer
      email = create_random_email
      description = "Sample customer for Salesforce"

      customer = Stripe::Customer.create({
        description: description,
        email: email,
      }, @user.stripe_credentials)
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
  end
end
