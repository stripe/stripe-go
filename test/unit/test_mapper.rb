# frozen_string_literal: true
# typed: ignore

require_relative '../test_helper'

module Critic::Unit
  class MapperTest < Critic::UnitTest
    before do
      @user = make_user(random_user_id: true)
      @mapper = Integrations::Mapper.new(@user)
    end

    it 'properly saves the annotator hashes' do
      annotation_table = {
        "customer" => {
          "subsidiary_id" => 1,
          "custentity_invoiceform_id" => 104,
          "terms_id" => 2,
          "account_number" => 123,
        },
        "sales_order" => {
          "account_id" => 3,
        },
      }

      @user.field_defaults = annotation_table
      @user.field_mappings = annotation_table
      @user.save

      @user = T.must(StripeForce::User.find(id: @user.id))

      assert_equal(annotation_table, @user.field_defaults)
      assert_equal(annotation_table, @user.field_mappings)
    end

    it 'assigns field fields' do
      @user.field_defaults = {
        "customer" => {
          "address.line1" => '123 Great Street',
          "metadata.from_salesforce" => true,
          "email" => 'global@email.com',
          "currency" => 'usd',
        },
        "subscription_schedule" => {
          "default_settings.collection_method" => "send_invoice",
        },
        "price" => {
          "metadata.tax_id" => 123,
        },
      }

      stripe_customer = Stripe::Customer.new
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.new
      stripe_price = Stripe::Price.new

      # these aren't used in the mapping logic, but are here to make sure passing them doesn't throw an exception
      sf_customer = create_mock_salesforce_customer
      sf_order = create_mock_salesforce_order

      @mapper.apply_mapping(stripe_customer, sf_customer)
      @mapper.apply_mapping(stripe_subscription_schedule, sf_order)
      @mapper.apply_mapping(stripe_price)

      assert_equal("123 Great Street", T.must(stripe_customer.address).line1)
      assert(stripe_customer.metadata["from_salesforce"])
      assert_equal("global@email.com", stripe_customer.email)
      assert_equal("usd", stripe_customer.currency)

      assert_equal("send_invoice", stripe_subscription_schedule.default_settings.collection_method)

      assert_equal(123, stripe_price.metadata["tax_id"])
    end

    it 'assigns field values from the dynamic user table and stripe resource metadata' do
      skip("need to finalize metadata work")

      @user.field_mappings = {
        "customer" => {
          "ns_alt_name" => "alt_name",
          "ns_invoice_form_id" => "custentity_invoiceform_id",
          "ns_another_field" => "custentity_string_custom",
          "ns_not_applicable" => "custentity_not_used",
          'ns_another_body_field' => 'custbodynounderscore',
          'ns_is_person' => 'is_person',
        },
        "subscription_schedule" => {
          'ns_body_field' => 'custbody_field',
          'ns_location' => 'location_id',
        },
      }

      stripe_customer = Stripe::Customer.new
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.new

      sf_customer = create_mock_salesforce_customer
      sf_order = create_mock_salesforce_order

      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      stripe_customer.metadata = {
        "ns_alt_name" => 'Mike',
        "ns_invoice_form_id" => "104",
        "ns_another_field" => "Kyle",
        'ns_another_body_field' => 'body field without underscore',
        'ns_is_person' => 'True',
      }

      stripe_charge = Stripe::Charge.construct_from(
        id: create_id(:ch),
        metadata: {
          'ns_body_field' => 'body field with underscore',
          "ns_location" => "2",
        }
      )

      @annotator.annotate(ns_customer, stripe_customer)
      @annotator.annotate(ns_credit_memo, stripe_charge)

      assert_equal('Mike', ns_customer.alt_name)
      assert_equal("104", ns_customer.custom_field_list.custentity_invoiceform.value.internal_id)
      assert_equal("Kyle", ns_customer.custom_field_list.custentity_string_custom.value)
      assert_equal("body field without underscore", ns_customer.custom_field_list.custbodynounderscore.value)
      refute_respond_to(ns_customer.custom_field_list, :custentity_not_used)
      assert(ns_customer.is_person)

      assert_equal("body field with underscore", ns_credit_memo.custom_field_list.custbody_field.value)
      assert_equal("2", ns_credit_memo.location.internal_id)
    end

    it 'defaults to the static table when metadata for the dynamic table value is not defined' do
      skip("should use metadata v2 syntax here")

      @user.field_defaults = {
        "customer" => {
          "alt_name" => "Mike",
        },
      }

      @user.field_mappings = {
        "customer" => {
          "ns_alt_name" => "alt_name",
        },
      }

      sf_customer = create_mock_salesforce_customer
      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      stripe_customer.metadata = {}

      @mapper.apply_mapping(stripe_customer, sf_customer)

      assert_equal('Mike', sf_customer.alt_name)

      sf_customer = create_mock_salesforce_customer
      stripe_customer.metadata = {"ns_alt_name" => "Kyle"}

      @mapper.apply_mapping(stripe_customer, sf_customer)

      assert_equal('Kyle', sf_customer.alt_name)
    end

    it 'does not annotate a user with blank annotation tables' do
      sf_customer = create_mock_salesforce_customer
      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      metadata = {
        "ns_alt_name" => 'Mike',
        "ns_invoice_form_id" => 104,
      }
      stripe_customer.metadata = metadata

      @mapper.apply_mapping(stripe_customer, sf_customer)

      refute_equal(metadata, stripe_customer.metadata)
    end

    # https://github.com/stripe/stripe-netsuite/issues/825
    it 'handles associated object annotation' do
      skip("dot syntax is not working yet")

      @user.field_mappings['invoice'] = {
        'metadata.random_thing' => 'custbody_random',
        'metadata.random_float' => 'custbody_random_float',
        'id' => 'custbody_invoiceid',
        'amount' => 'tran_id',

        # ensure a bunch of invalid mappings are testing
        'date' => '',
        'invalid_field' => 'other_ref_num',
        'invalid.field' => 'other_ref_num',
        'currency' => 2,
        'status' => nil,

        'subscription.metadata.order_number' => 'custbodystripe_subscription_number',
        'subscription.id' => 'custbody_subscriptionid',
      }

      @user.field_mappings['invoice_item'] = {
        'metadata.foo' => 'description',
      }

      subscription_id = create_id(:sub)
      invoice_line_id = create_id(:ii)

      stripe_invoice = Stripe::Invoice.construct_from({
        id: create_id(:in),
        subscription: subscription_id,
        date: Time.now.to_i,
        currency: 'usd',
        amount: 100_00,
        status: 'paid',
        metadata: {
          random_thing: 'STRING',
          random_float: 13.0,
        },
        lines: {object: 'list', data: [
          {
            # there are many different fields that contain variants of the same obj reference
            # due to us being on an old API version
            # https://jira.corp.stripe.com/browse/REPROD-888
            id: invoice_line_id,
            unique_id: invoice_line_id,
            invoice_item: invoice_line_id,

            object: "line_item",

            metadata: {foo: "bar"},
          },
        ],},
      })

      stripe_subscription = Stripe::Subscription.construct_from(
        id: subscription_id,
        metadata: {
          order_number: 'NUMBER',
        }
      )

      Stripe::Subscription
        .expects(:retrieve)
        .at_least_once
        .with(stripe_invoice.subscription, @user.stripe_client_credentials)
        .returns(stripe_subscription)

      Stripe::InvoiceItem
        .expects(:retrieve)
        .at_least_once
        .with(invoice_line_id, anything)
        .returns(stripe_invoice.lines.first)

      ns_invoice = NetSuite::Records::Invoice.new

      @annotator.annotate(ns_invoice, stripe_invoice)

      assert_equal(5, ns_invoice.custom_field_list.custom_fields.count, ns_invoice.custom_field_list.custom_fields)
      assert_equal(stripe_subscription.metadata['order_number'], ns_invoice.custom_field_list.custbodystripe_subscription_number.value)
      assert_equal(stripe_invoice.metadata['random_thing'], ns_invoice.custom_field_list.custbody_random.value)
      assert_equal(stripe_subscription.id, ns_invoice.custom_field_list.custbody_subscriptionid.value)
      assert_equal(stripe_invoice.id, ns_invoice.custom_field_list.custbody_invoiceid.value)
      assert_equal(stripe_invoice.amount, ns_invoice.tran_id)
      assert_nil(ns_invoice.other_ref_num)

      # annotate invoice lines, which are handled in a specialized way due to our old API version
      ns_line = ns_invoice.item_list.sublist_class.new
      @annotator.annotate(ns_line, stripe_invoice.lines.first)

      assert_equal("bar", ns_line.description)
    end

    it 'allows payment intent metadata to be pulled via a charge' do
      skip("dot syntax is not working yet")

      payment_intent_id = create_id(:pi)

      payment_intent = Stripe::PaymentIntent.construct_from(
        id: payment_intent_id,
        metadata: {
          foo: 'bar',
        }
      )

      Stripe::PaymentIntent.expects(:retrieve).once.returns(payment_intent)

      charge = Stripe::Charge.construct_from(
        id: create_id(:ch),
        payment_intent: payment_intent_id,
        metadata: {}
      )

      @user.field_mappings['customer_payment'] = {
        'payment_intent.metadata.foo' => 'custbody_foo',
      }

      ns_customer_payment = NetSuite::Records::CustomerPayment.new

      @annotator.annotate(ns_customer_payment, charge)

      assert_equal('bar', ns_customer_payment.custom_field_list.custbody_foo.value)
    end
  end
end
