# frozen_string_literal: true
# typed: true

# TODO move this somewhere else
# tired of writing binding.pry...
module Kernel
  def bp
    Pry.start(binding.of_caller(1))
  end
end

module SalesforceDebugging
  include StripeForce::Utilities::SalesforceUtil
  include StripeForce::Constants

  def stripe_get(stripe_id)
    stripe_id = stripe_id.
      # remove whitespace
      strip.
      # remove quotes, sometimes easy to pull in when copy/pasting
      gsub("'\"", '')

    stripe_class = Integrations::Utilities::StripeUtil.stripe_class_from_id(stripe_id, raise_on_missing: true)
    return if !stripe_class

    stripe_class.retrieve(stripe_id, @user.stripe_credentials)
  end

  # https://stackoverflow.com/questions/9742913/validating-a-salesforce-id
  def looks_like_sf_id?(some_string)
    some_string =~ /^([a-zA-Z0-9]{18}|[a-zA-Z0-9]{15})$/
  end

  def sf_get(sf_id)
    sf_id = sf_id.
      # remove whitespace
      strip.
      # remove quotes, sometimes easy to pull in when copy/pasting
      gsub("'\"", '')

    sf_type = salesforce_type_from_id(sf_id)
    sf_object = @user.sf_client.find(sf_type, sf_id)
  end

  def sf_get_related(source_object_or_id, related_object)
    source_object = if source_object_or_id.is_a?(String)
      sf_get(source_object_or_id)
    else
      source_object_or_id
    end

    # if a custom object, `Id` is not appended
    relationship_field = if source_object.sobject_type.end_with?('__c')
      source_object.sobject_type
    else
      "#{source_object.sobject_type}Id"
    end

    @user.sf_client.query("SELECT Id FROM #{related_object} WHERE relationship_field = '#{source_object.Id}'").map do |o|
      @user.sf_client.find(related_object, o.Id)
    end
  end

  def sf_get_recent(object_type)
    order_id = @user.sf_client.query("SELECT Id FROM #{object_type} ORDER BY CreatedDate DESC LIMIT 1").first.Id
    puts order_id
    sf_get(order_id)
  end

  def stripe_jump(id_or_object)
    id_or_object = id_or_object.id if id_or_object.is_a?(Stripe::APIResource)
    stripe_url = "https://dashboard.stripe.com/#{@user.stripe_account_id}/test/id/#{id_or_object}?account=acct_1JdEinJMQw1guJjW"

    if Rails.env.production?
      puts stripe_url
    else
      `open "#{stripe_url}"`
    end
  end

  def sf_jump(id_or_object)
    sf = @user.sf_client if !sf
    @user_info = sf.user_info

    id_or_object = id_or_object[SF_ID] if id_or_object.is_a?(Restforce::SObject)
    sf_url = "#{@user_info['urls']['custom_domain']}/#{id_or_object}"

    if Rails.env.production?
      puts sf_url
    else
      `open "#{sf_url}"`
    end
  end

  def sf_object_manager(id_or_type_or_object)
    if id_or_type_or_object.is_a?(Restforce::SObject)
      id_or_type_or_object = id_or_type_or_object[SF_ID]
    end

    sf_type = if looks_like_sf_id?(id_or_type_or_object)
      salesforce_type_from_id(id_or_type_or_object)
    else
      id_or_type_or_object
    end

    @user_info = @user.sf_client.user_info

    sf_url = "#{@user_info['urls']['custom_domain']}/lightning/setup/ObjectManager/#{sf_type}/FieldsAndRelationships/view"

    if Rails.env.production?
      puts sf_url
    else
      `open "#{sf_url}"`
    end
  end

  def salesforce_instance_type
    result = @user.sf_client.query("SELECT InstanceName, IsSandbox, TrialExpirationDate FROM Organization").first

    if !!result['IsSandbox'] && result['TrialExpirationDate'].nil?
      SFInstanceTypes::SANDBOX
    elsif !!result['IsSandbox']
      SFInstanceTypes::SCRATCH_ORG
    elsif !result['TrialExpirationDate'].nil?
      SFInstanceTypes::TRIAL
    else
      SFInstanceTypes::PRODUCTION
    end
  end

  def wipe_account_and_orders(sf_account_or_id)
    sf_account = if sf_account_or_id.is_a?(String)
      sf_get(sf_account_or_id)
    else
      sf_account_or_id
    end

    if sf_account.sobject_type == SF_ORDER
      sf_account = sf_get(sf_account.AccountId)
    end

    if sf_account.sobject_type != SF_ACCOUNT
      puts "invalid object passed: account or order required"
      return
    end

    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    stripe_customer.delete

    @user.sf_client.query("SELECT Id FROM Order WHERE AccountId = '#{sf_account.Id}'").each do |sf_order|
      puts "Removing Stripe ID from Order: #{sf_order.Id}"

      @user.sf_client.update!(SF_ORDER, {
        SF_ID => sf_order.Id,
        prefixed_stripe_field(GENERIC_STRIPE_ID) => nil,
      })
    end

    puts "Removing Stripe ID from Account: #{sf_account.Id}"
    @user.sf_client.update!(SF_ACCOUNT, {
      SF_ID => sf_account.Id,
      prefixed_stripe_field(GENERIC_STRIPE_ID) => nil,
    })
  end

  def wipe_product(product_id)
    puts "updating product\t#{product_id}"
    @user.sf_client.update!(SF_PRODUCT, {
      SF_ID => product_id,
      prefixed_stripe_field(GENERIC_STRIPE_ID) => "",
    })

    @user.sf_client.query("SELECT Id FROM #{SF_PRICEBOOK_ENTRY} WHERE Product2Id = '#{product_id}'").each do |pricebook_entry|
      puts "wiping pricebook #{pricebook_entry.Id}"
      @user.sf_client.update!(SF_PRICEBOOK_ENTRY, {
        SF_ID => pricebook_entry.Id,
        prefixed_stripe_field(GENERIC_STRIPE_ID) => "",
      })
    end
  end
end
