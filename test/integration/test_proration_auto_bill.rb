# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

class Critic::ProrationAutoBillTranslation < Critic::VCRTest
  before do
    set_cassette_dir(__FILE__)
    Timecop.freeze(VCR.current_cassette.originally_recorded_at || now_time)

    @user = make_user(save: true)
  end

  it 'creates an invoice automatically for a invoice item' do
    @user.enable_feature FeatureFlags::AUTO_ADVANCE_PRORATION_INVOICE, update: true
    @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

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
        StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
      },
    }, @user.stripe_credentials)

    invoice_event = get_invoice_item_event(invoice_item.id)

    # we don't want the api keys to be in the event so we can mimic prod
    copied_event = Stripe::Event.construct_from(invoice_event.to_hash)
    assert_raises(Stripe::AuthenticationError) { copied_event.refresh }

    invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, copied_event)

    invoice = T.must(invoice)
    assert_equal(1, invoice.lines.count)
    assert_equal("true", invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])

    assert(invoice.auto_advance)
    assert_equal('draft', invoice.status)

    # let's advance the clock by an hour and test the invoice finalizes
    stripe_customer = stripe_get(subscription.customer)
    refute_nil(stripe_customer.test_clock)
    advance_test_clock(stripe_customer, (Time.now + 2.hour).to_i)

    invoice = Stripe::Invoice.retrieve(invoice.id, @user.stripe_credentials)
    assert_not_equal('draft', invoice.status)
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

      invoice_event = get_invoice_item_event(invoice_item.id)

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

      invoice_event = get_invoice_item_event(invoice_item.id)

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
          StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
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

    it 'is duplicate invoice item' do
      # @user.enable_feature FeatureFlags::AUTO_ADVANCE_PRORATION_INVOICE, update: true
      @user.enable_feature FeatureFlags::TEST_CLOCKS, update: true

      subscription = create_customer_with_subscription
      _, ad_hoc_price = create_price(additional_price_fields: {recurring: {}})

      # create an invoice item that looks similar to what we'll get from add_invoice_items
      start_timestamp = Time.now.to_i
      end_timestamp = Time.now.to_i + 1.day.to_i
      invoice_item_data = {
        customer: subscription.customer,
        subscription: subscription.id,
        price: ad_hoc_price.id,
        quantity: 1,
        period: {
          start: start_timestamp,
          end: end_timestamp,
        },
        metadata: {
          StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
          "salesforce_order_item_id" => "8025c00000DcVBNAAV",
        },
      }

      # create the invoice_item and invoice_item event
      invoice_item = Stripe::InvoiceItem.create(invoice_item_data, @user.stripe_credentials)
      invoice_item_event = get_invoice_item_event(invoice_item.id)

      # create the invoice
      invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_item_event)
      invoice = T.must(invoice)
      assert_equal(1, invoice.lines.count)
      assert_equal("true", invoice.metadata[StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION_INVOICE)])
      assert_equal('draft', invoice.status)

      # now let's pretend a duplicate invoice_item event is triggered from the same data
      duplicate_invoice_item = Stripe::InvoiceItem.create(invoice_item_data, @user.stripe_credentials)
      duplicate_invoice_item_event = get_invoice_item_event(duplicate_invoice_item.id)
      invoice = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, duplicate_invoice_item_event)
      assert_nil(invoice)

      # change the metadata and ensure an invoice is created
      invoice_item_data_2 = {
        customer: subscription.customer,
        subscription: subscription.id,
        price: ad_hoc_price.id,
        quantity: 1,
        period: {
          start: start_timestamp,
          end: end_timestamp,
        },
        metadata: {
          StripeForce::Translate::Metadata.metadata_key(@user, MetadataKeys::PRORATION) => "true",
          "salesforce_order_item_id" => "8025c00000DcVBNAAY", # this is a different ID than above
        },
      }
      invoice_item_2 = Stripe::InvoiceItem.create(invoice_item_data_2, @user.stripe_credentials)
      invoice_event_2 = get_invoice_item_event(invoice_item_2.id)

      invoice_2 = StripeForce::ProrationAutoBill.create_invoice_from_invoice_item_event(@user, invoice_event_2)
      assert_not_nil(invoice_2)
    end
  end
end
