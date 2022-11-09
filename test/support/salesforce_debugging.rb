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

  def is_order_amendment?(sf_order)
    order_with_amended_contract_query = @user.sf_client.query(
      # include `Type`, `OpportunityId` for debugging purposes
      <<~EOL
        SELECT Type, OpportunityId,
              Opportunity.SBQQ__AmendedContract__c
        FROM #{SF_ORDER}
        WHERE Id = '#{sf_order.Id}'
      EOL
    )

    if order_with_amended_contract_query.size.zero?
      raise Integrations::Errors::ImpossibleInternalError.new("query should never return an empty result")
    end

    if order_with_amended_contract_query.size > 1
      raise Integrations::Errors::ImpossibleInternalError.new("query should only return a single result")
    end

    order_with_amended_contract = order_with_amended_contract_query.first

    amended_contract_id = order_with_amended_contract.dig(SF_OPPORTUNITY, "SBQQ__AmendedContract__c")
    is_order_amendment = amended_contract_id.present?

    if !OrderTypeOptions.values.map(&:serialize).include?(order_with_amended_contract.Type)
      log.warn 'order type is not standard', order_type: order_with_amended_contract.Type
    end

    if is_order_amendment && order_with_amended_contract.Type == OrderTypeOptions::NEW.serialize
      Integrations::ErrorContext.report_edge_case("order is determined to be an amendment, but type is new")
    end

    is_order_amendment
  end

  def extract_initial_order_from_amendment(sf_order_amendment)
    # in the case of an amendment, the associated opportunity contains a reference to the contract
    # which contains a reference to the original quote, which references the orginal order (initial non-amendment order)
    initial_quote_query = @user.sf_client.query(
      # include `OpportunityId`, `SBQQ__AmendedContract__c` for debugging purposes
      <<~EOL
        SELECT OpportunityId, Opportunity.SBQQ__AmendedContract__c,
               Opportunity.SBQQ__AmendedContract__r.#{SF_CONTRACT_QUOTE_ID}
        FROM #{SF_ORDER}
        WHERE Id = '#{sf_order_amendment.Id}'
      EOL
    )

    if initial_quote_query.count.zero?
      raise Integrations::Errors::ImpossibleState.new("order amendments should always be associated with the initial quote")
    end

    # TODO this should never happen and should be removed
    if initial_quote_query.count > 1
      raise Integrations::Errors::ImpossibleState.new("exact ID match yields two records")
    end

    # the contract tied to the amended order has the ID of the quote of the original order
    # let's get that ID, then we can pull the order tied to that original quote
    sf_original_quote_id = initial_quote_query.first.dig(
      SF_OPPORTUNITY,
      # TODO should pull into a constant
      "SBQQ__AmendedContract__r",
      SF_CONTRACT_QUOTE_ID
    )

    if sf_original_quote_id.blank?
      log.warn "related records while retrieve missing quote",
        opportunity_id: initial_quote_query.first.dig("OpportunityId"),
        amended_contract: initial_quote_query.first.dig(SF_OPPORTUNITY, "SBQQ__AmendedContract__c")

      raise StripeForce::Errors::RawUserError.new("Could not find initial quote associated with order amendment")
    end

    log.info 'quote tied to initial order found', quote_id: sf_original_quote_id

    initial_order_query = @user.sf_client.query(
      <<~EOL
        SELECT Id
        FROM #{SF_ORDER}
        WHERE SBQQ__Quote__c = '#{sf_original_quote_id}'
      EOL
    )

    if initial_order_query.count.zero?
      raise Integrations::Errors::ImpossibleState.new("initial order should be associated with an initial quote")
    end

    # TODO this should never happen and should be removed
    if initial_order_query.count > 1
      raise Integrations::Errors::ImpossibleState.new("exact ID match yields two records")
    end

    initial_order_id = initial_order_query.first[SF_ID]

    log.info 'found initial order', initial_order_id: initial_order_id

    sf_get(initial_order_id)
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

    sf_type = StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_id)
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

    @user.sf_client.query("SELECT Id FROM #{related_object} WHERE #{relationship_field} = '#{source_object.Id}'").map do |o|
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
      StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, id_or_type_or_object)
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

  def wipe_object(sf_object_or_id)
    sf_object = if sf_object_or_id.is_a?(String)
      sf_get(sf_object_or_id)
    else
      sf_object_or_id
    end

    puts "updating\t#{sf_object.sobject_type}\t#{sf_object.Id}"
    @user.sf_client.update!(sf_object.sobject_type, {
      SF_ID => sf_object.Id,
      prefixed_stripe_field(GENERIC_STRIPE_ID) => "",
    })
  end
end
