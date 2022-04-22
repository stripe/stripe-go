# frozen_string_literal: true
# typed: true

module SalesforceDebugging
  include StripeForce::Utilities::SalesforceUtil
  include StripeForce::Constants

  def sf_get(sf_id)
    sf_type = salesforce_type_from_id(sf_id)
    sf_object = @user.sf_client.find(sf_type, sf_id)
  end

  def stripe_jump(id_or_object)
    id_or_object = id_or_object.id if id_or_object.is_a?(Stripe::APIResource)
    "https://dashboard.stripe.com/#{@user.stripe_account_id}/test/id/#{id_or_object}?account=acct_1JdEinJMQw1guJjW"
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

  def sf_object_manager(id_or_object)
    id_or_object = id_or_object[SF_ID] if id_or_object.is_a?(Restforce::SObject)
    sf_type = salesforce_type_from_id(id_or_object)

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

  def wipe_account_and_orders(sf_account)
    stripe_customer_id = sf_account[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    stripe_customer = Stripe::Customer.retrieve(stripe_customer_id, @user.stripe_credentials)
    stripe_customer.delete

    @user.sf_client.query("SELECT Id FROM Order WHERE AccountId = '#{sf_account.Id}'").each do |sf_order|
      puts "Removing Stripe ID From: #{sf_order.Id}"

      @user.sf_client.update!(SF_ORDER, {
        SF_ID => sf_order.Id,
        prefixed_stripe_field(GENERIC_STRIPE_ID) => nil,
      })
    end

    puts "Removing Stripe ID From: #{sf_account.Id}"
    @user.sf_client.update!(SF_ACCOUNT, {
      SF_ID => sf_account.Id,
      prefixed_stripe_field(GENERIC_STRIPE_ID) => nil,
    })
  end

end
