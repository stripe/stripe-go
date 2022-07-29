# frozen_string_literal: true
# typed: true

require_relative '../test_helper'

module Critic::Unit
  class MapperTest < Critic::UnitTest
    before do
      @user = make_user(random_user_id: true)
      @mapper = StripeForce::Mapper.new(@user)
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
          "metadata.Subscription Type" => "great_type",
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
      assert_equal("great_type", stripe_subscription_schedule.metadata['Subscription Type'])

      assert_equal(123, stripe_price.metadata["tax_id"])
    end

    it 'handles dynamic field mappings' do
      @user.field_mappings = {
        "customer" => {
          "metadata.custom_field" => "CustomField__c",
          "phone" => "Phone",
        },
        "subscription_schedule" => {
          'metadata.field_custom' => 'CustomField__c',
          'metadata.second_custom' => 'CustomField__c',
          'default_settings.automatic_tax' => 'TaxBoolean__c',
        },
      }

      stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))
      stripe_subscription_schedule = Stripe::SubscriptionSchedule.construct_from(id: stripe_create_id(:sub_sched))

      # set some previous metadata; the mapper should be additive
      stripe_subscription_schedule.metadata = {
        "salesforce_order_id" => "801780000006retAAA",
        "salesforce_order_link" => "https://cloudflare--stg2.my.salesforce.com/801780000006retAAA",
      }

      sf_customer = create_mock_salesforce_customer
      sf_customer['CustomField__c'] = 'a custom value'
      sf_customer['Phone'] = '4847534128'

      sf_order = create_mock_salesforce_order
      sf_order['CustomField__c'] = 'another custom value'
      sf_order['TaxBoolean__c'] = true

      @mapper.apply_mapping(stripe_customer, sf_customer)
      @mapper.apply_mapping(stripe_subscription_schedule, sf_order)

      assert_equal('a custom value', stripe_customer.metadata.custom_field)
      assert_equal("4847534128", stripe_customer.phone)

      assert_equal("another custom value", stripe_subscription_schedule.metadata.field_custom)
      assert_equal("another custom value", stripe_subscription_schedule.metadata.second_custom)
      assert_equal(sf_order['TaxBoolean__c'], stripe_subscription_schedule.default_settings.automatic_tax)

      refute_nil(stripe_subscription_schedule.metadata["salesforce_order_id"])
      refute_nil(stripe_subscription_schedule.metadata["salesforce_order_link"])
    end

    it 'defaults to the static table when metadata for the dynamic table value is not defined' do
      @user.field_defaults = {
        "customer" => {
          "metadata.example" => "Mike",
        },
      }

      @user.field_mappings = {
        "customer" => {
          "metadata.example" => "Alt_Name__c",
        },
      }

      sf_customer = create_mock_salesforce_customer
      stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))
      stripe_customer.metadata = {}

      @mapper.apply_mapping(stripe_customer, sf_customer)

      assert_equal('Mike', stripe_customer.metadata['example'])

      sf_customer = create_mock_salesforce_customer
      sf_customer['Alt_Name__c'] = "Kyle"

      @mapper.apply_mapping(stripe_customer, sf_customer)

      assert_equal('Kyle', stripe_customer.metadata['example'])
    end

    it 'does not annotate a user with blank annotation tables' do
      sf_customer = create_mock_salesforce_customer
      stripe_customer = Stripe::Customer.construct_from(id: stripe_create_id(:cus))
      metadata = {
        "ns_alt_name" => 'Mike',
        "ns_invoice_form_id" => 104,
      }
      stripe_customer.metadata = metadata

      @mapper.apply_mapping(stripe_customer, sf_customer)

      refute_equal(metadata, stripe_customer.metadata)
    end

    it 'supports dot path syntax when extracting data from salesforce' do
      # create a customer with a parent reference
      parent_id = create_salesforce_account
      child_id = create_salesforce_account

      sf.update!(SF_ACCOUNT, {
        SF_ID => parent_id,
        "Description" => "parent description",
      })

      sf.update!(SF_ACCOUNT, {
        SF_ID => child_id,
        "ParentId" => parent_id,
      })

      @user.field_mappings['customer'] = {
        'metadata.parent_description' => 'Parent.Description',
      }
      @user.save

      stripe_customer = Stripe::Customer.new
      sf_customer = sf_get(child_id)

      @mapper.apply_mapping(stripe_customer, sf_customer)

      assert_equal("parent description", stripe_customer.metadata['parent_description'])
    end

    it 'generates compound ids' do
      stripe_record = Stripe::Price.new
      sf_order_item = create_mock_salesforce_order_item

      key = StripeForce::Mapper.mapping_key_for_record(stripe_record, sf_order_item)

      assert_equal('price_order_item', key)
    end
  end
end
