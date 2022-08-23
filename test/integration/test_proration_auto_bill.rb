# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ProrationAutoBillTranslation < Critic::FunctionalTest
  before do
    @user = make_user(save: true)
  end

  def get_invoice_event(invoice_item_id)
    # events can take some time to propogate
    events = T.let(nil, T.untyped)

    wait_until do
      events = Stripe::Event.list({
        object_id: invoice_item_id,
      }, @user.stripe_credentials)

      events.count >= 1
    end

    events.first
  end

  it 'creates an invoice automatically for a invoice item' do
    subscription = create_customer_with_subscription
    _, ad_hoc_price = create_price(additional_price_fields: {
      recurring: {},
    })
    # create an invoice item that looks similar to what we'll get from add_invoice_items
    invoice_item = Stripe::InvoiceItem.create({
      customer: subscription.customer,
      subscription: subscription.id,
      price: ad_hoc_price.id,
      quantity: 1,
      period: {
        start: Time.now.to_i,
        end: Time.now.to_i + 1.day.to_i,
      },
      metadata: {
        StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
      },
    }, @user.stripe_credentials)

    invoice_event = get_invoice_event(invoice_item.id)

    # we don't want the api keys to be in the event so we can mimic prod
    copied_event = Stripe::Event.construct_from(invoice_event.to_hash)
    assert_raises(Stripe::AuthenticationError) { copied_event.refresh }

    invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, copied_event)

    invoice = T.must(invoice)
    assert_equal(1, invoice.lines.count)
    assert_equal("true", invoice.metadata[StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
  end

  describe 'skip conditions' do
    it 'not tied to a subscription' do
      customer = create_customer
      _, ad_hoc_price = create_price(additional_price_fields: {
        recurring: {},
      })

      invoice_item = Stripe::InvoiceItem.create({
        # no subscription!
        customer: customer.id,
        price: ad_hoc_price.id,
        quantity: 1,
        period: {
          start: Time.now.to_i,
          end: Time.now.to_i + 1.day.to_i,
        },
      }, @user.stripe_credentials)

      invoice_event = get_invoice_event(invoice_item.id)

      invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event)

      assert_nil(invoice)
    end

    it 'does not contain the magic metadata key' do
      subscription = create_customer_with_subscription
      _, ad_hoc_price = create_price(additional_price_fields: {
        recurring: {},
      })
      # create an invoice item that looks similar to what we'll get from add_invoice_items
      invoice_item = Stripe::InvoiceItem.create({
        customer: subscription.customer,
        subscription: subscription.id,
        price: ad_hoc_price.id,
        quantity: 1,
        period: {
          start: Time.now.to_i,
          end: Time.now.to_i + 1.day.to_i,
        },
      }, @user.stripe_credentials)

      invoice_event = get_invoice_event(invoice_item.id)

      invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event)

      assert_nil(invoice)
    end

    it 'is already invoiced' do
      subscription = create_customer_with_subscription
      _, ad_hoc_price = create_price(additional_price_fields: {
        recurring: {},
      })

      # create an invoice item that looks similar to what we'll get from add_invoice_items
      invoice_item = Stripe::InvoiceItem.create({
        customer: subscription.customer,
        subscription: subscription.id,
        price: ad_hoc_price.id,
        quantity: 1,
        period: {
          start: Time.now.to_i,
          end: Time.now.to_i + 1.day.to_i,
        },
        metadata: {
          StripeForce::Utilities::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
        },
      }, @user.stripe_credentials)

      events = Stripe::Event.list({
        object_id: invoice_item.id,
      }, @user.stripe_credentials)

      assert_equal(1, events.count)
      invoice_event = events.first

      # pretend another system invoiced this
      out_of_band_invoice = Stripe::Invoice.create({
        customer: subscription.customer,
        subscription: subscription.id,
      }, @user.stripe_credentials)


      invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event)

      assert_nil(invoice)
    end
  end
end
