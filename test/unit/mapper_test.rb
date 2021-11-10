# frozen_string_literal: true
# typed: true
require_relative '../_lib'

module Critic::Unit
  class NetSuiteAnnotatorTest < Critic::Test
    before do
      @user = make_user(random_user_id: true)
      @user.disable_feature(:annotator_v2)

      @annotator = StripeSuite::Annotator.new(@user)
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

      @user.netsuite_static_annotations = annotation_table
      @user.netsuite_dynamic_annotations = annotation_table
      @user.save

      @user = T.must(StripeSuite::User.find(id: @user.id))

      assert_equal(annotation_table, @user.netsuite_static_annotations)
      assert_equal(annotation_table, @user.netsuite_dynamic_annotations)
    end

    it 'assigns field values from static user table' do
      @user.netsuite_static_annotations = {
        "customer" => {
          "subsidiary_id" => 1,
          "custentity_invoiceform_id" => 104,
          "terms_id" => 2,
          "account_number" => 123,
        },
        "sales_order" => {
          "account_id" => 3,
        },
        "invoice" => {
          "due_date_date" => 1493681421,
        },
      }

      ns_customer = NetSuite::Records::Customer.new
      ns_sales_order = NetSuite::Records::SalesOrder.new
      ns_invoice = NetSuite::Records::Invoice.new

      # static annotations don't use stripe data; some legacy annotations do
      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      stripe_customer.metadata = {}

      @annotator.annotate(ns_customer, stripe_customer)
      @annotator.annotate(ns_sales_order, stripe_customer)
      @annotator.annotate(ns_invoice)

      assert_equal(104, ns_customer.custom_field_list.custentity_invoiceform.value.internal_id)
      assert_equal(1, ns_customer.subsidiary.internal_id)
      assert_equal(123, ns_customer.account_number)

      assert_equal(3, ns_sales_order.account.internal_id)

      assert_equal("2017-05-01T00:00:00-07:00", ns_invoice.due_date)
    end

    it 'assigns field values from the dynamic user table and stripe resource metadata' do
      @user.netsuite_dynamic_annotations = {
        "customer" => {
          "ns_alt_name" => "alt_name",
          "ns_invoice_form_id" => "custentity_invoiceform_id",
          "ns_another_field" => "custentity_string_custom",
          "ns_not_applicable" => "custentity_not_used",
          'ns_another_body_field' => 'custbodynounderscore',
          'ns_is_person' => 'is_person',
        },
        "credit_memo" => {
          'ns_body_field' => 'custbody_field',
          'ns_location' => 'location_id',
        },
      }

      ns_customer = NetSuite::Records::Customer.new
      ns_credit_memo = NetSuite::Records::CreditMemo.new

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
      assert_equal(false, ns_customer.custom_field_list.respond_to?(:custentity_not_used))
      assert_equal(true, ns_customer.is_person)

      assert_equal("body field with underscore", ns_credit_memo.custom_field_list.custbody_field.value)
      assert_equal("2", ns_credit_memo.location.internal_id)
    end

    it 'defaults to the static table when metadata for the dynamic table value is not defined' do
      @user.netsuite_static_annotations = {
        "customer" => {
          "alt_name" => "Mike",
        },
      }

      @user.netsuite_dynamic_annotations = {
        "customer" => {
          "ns_alt_name" => "alt_name",
        },
      }

      ns_customer = NetSuite::Records::Customer.new
      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      stripe_customer.metadata = {}

      @annotator.annotate(ns_customer, stripe_customer)

      assert_equal('Mike', ns_customer.alt_name)

      ns_customer = NetSuite::Records::Customer.new
      stripe_customer.metadata = {"ns_alt_name" => "Kyle"}

      @annotator.annotate(ns_customer, stripe_customer)

      assert_equal('Kyle', ns_customer.alt_name)
    end

    it 'does not annotate a user with blank annotation tables' do
      ns_record = NetSuite::Records::Customer.new
      stripe_customer = Stripe::Customer.construct_from(id: create_id(:cus))
      stripe_customer.metadata = {
        "ns_alt_name" => 'Mike',
        "ns_invoice_form_id" => 104,
      }

      @annotator.annotate(ns_record, stripe_customer)

      refute_equal('Mike', ns_record.alt_name)
      assert_equal(false, ns_record.custom_field_list.respond_to?(:custentity_invoiceform))
    end

    it "handles stripe resources without metadata" do
      stripe_customer = Stripe::Customer.construct_from(
        deleted: true,
        id: create_id(:cus)
      )
      ns_record = NetSuite::Records::Customer.new

      @annotator.annotate(ns_record, stripe_customer)
    end

    # TODO this should be generalized https://github.com/stripe/stripe-netsuite/issues/257
    # TODO this spec tests two approaches: we should eliminate the `cash_back_list_department_id` approach
    #      and use sublist item annotations instead
    it 'handles deposit sublist annotations' do
      @user.netsuite_static_annotations['deposit_cash_back'] = {
        'department_id' => 1000,
        'klass_id' => nil,
      }

      ns_deposit = NetSuite::Records::Deposit.new
      ns_cash_back_item = ns_deposit.cash_back_list.sublist_class.new({
        amount: 100,
        klass: {internal_id: 123},
      })

      @annotator.annotate(ns_cash_back_item)

      assert_equal(1000, ns_cash_back_item.department.internal_id)
      assert_nil(ns_cash_back_item.klass.internal_id)
    end

    describe 'feature: annotator_v2' do
      before { @user.enable_feature(:annotator_v2) }

      # https://github.com/stripe/stripe-netsuite/issues/825
      it 'handles associated object annotation' do
        @user.netsuite_dynamic_annotations['invoice'] = {
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

        @user.netsuite_dynamic_annotations['invoice_item'] = {
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
          status: StripeInvoiceStatus::PAID.serialize,
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

        @user.netsuite_dynamic_annotations['customer_payment'] = {
          'payment_intent.metadata.foo' => 'custbody_foo',
        }

        ns_customer_payment = NetSuite::Records::CustomerPayment.new

        @annotator.annotate(ns_customer_payment, charge)

        assert_equal(ns_customer_payment.custom_field_list.custbody_foo.value, 'bar')
      end
    end

    # https://github.com/stripe/stripe-netsuite/issues/257
    it 'handles annotating sublists that do not have unique sublist item classes' do
      @user.netsuite_static_annotations['service_sale_item'] = {
        'subsidiary_id' => 102,
      }

      ns_item = NetSuite::Records::ServiceSaleItem.new

      @annotator.annotate(ns_item)

      assert_equal(102, ns_item.subsidiary_list.record_ref.first.internal_id.to_i)
    end

    # https://jira.corp.stripe.com/browse/DATAIO-170
    it 'gracefully fails with an error when a custom field is provided, but no custom field exists on the record' do
      @user.netsuite_static_annotations['deposit_other'] = {
        "custentity_cseg1" => 39,
        "location_id" => 1,
      }

      # this record type does not contain a custom field list, this is very rare
      ns_deposit_other = NetSuite::Records::DepositOther.new
      refute(ns_deposit_other.respond_to?(:custom_field_list))

      @annotator.annotate(ns_deposit_other)

      assert_equal(1, ns_deposit_other.location.internal_id)
    end

    describe 'compound annotations' do
      before { @user.enable_feature(:annotator_v2) }

      it 'uses the same field defaults/static annotations key' do
        @user.netsuite_static_annotations['credit_memo'] = {
          "location_id" => 1,
        }

        ns_memo = NetSuite::Records::CreditMemo.new
        stripe_credit_note = Stripe::CreditNote.construct_from(id: create_id(:cn))

        @annotator.annotate(ns_memo, stripe_credit_note)

        assert_equal(1, ns_memo.location.internal_id)
      end

      it 'uses a different dynamic annotations key' do
        @user.netsuite_dynamic_annotations = {
          "credit_memo_credit_note" => {
            "memo" => "memo",
            "metadata.custom_data" => 'other_ref_num',
          },

          "credit_memo": {
            "memo" => "email",
            "metadata.custom_data" => 'vat_reg_num',
          },
        }

        ns_memo = NetSuite::Records::CreditMemo.new
        stripe_credit_note = Stripe::CreditNote.construct_from(
          id: create_id(:cn),
          memo: 'amazing memo',
          metadata: {
            custom_data: 'custom data',
          }
        )

        @annotator.annotate(ns_memo, stripe_credit_note)

        assert_match("amazing memo", ns_memo.memo)
        assert_equal('custom data', ns_memo.other_ref_num)

        assert_nil(ns_memo.email)
        assert_nil(ns_memo.vat_reg_num)
      end
    end
  end
end
