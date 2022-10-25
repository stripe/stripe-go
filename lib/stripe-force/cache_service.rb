# frozen_string_literal: true
# typed: true

class CacheService
  extend T::Sig

  include StripeForce::Constants
  include StripeForce::Utilities::SalesforceUtil

  # Cache Structure:
  # {
  #   "Pricebook2": {
  #     "01s2F000004YDPVQA4" => {
  #       ...
  #     }
  #   }
  # }
  def initialize(user)
    @user = user
    @cache = {}
    @previously_invalidated = []
  end

  sig { params(sf_object: Restforce::SObject).void }
  def cache_for_object(sf_object)
    record_list = []

    # grab bulked catch response from SF
    if enabled?
      record_list = retrieve_related_objects_from_salesforce(sf_object)
    end

    cache_list_of_records(record_list)
  end

  def invalidate_cache_object(sf_object_id)
    log.info("Invalidating cached object", sf_object_id: sf_object_id)
    @cache[StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_object_id)] ||= {}
    @cache[StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(@user, sf_object_id)].delete(sf_object_id)

    @previously_invalidated.append(sf_object_id)
  end

  # ie
  # order_amendment_query = sf.query(
  #   <<~EOL
  #     SELECT Id FROM #{SF_ORDER}
  #     WHERE Opportunity.SBQQ__AmendedContract__c = '#{sf_contract_id}'
  #     ORDER BY SBQQ__Quote__r.SBQQ__StartDate__c ASC
  #   EOL
  # )
  #
  # ~~~~~~~~~~~~~~~~~~~~~~~~~ Turns into ~~~~~~~~~~~~~~~~~~~~~~~~~
  # Notice the __r -> __c swap.
  # We don't have dot path relationships like the __r represents.
  # We have to get the ID and go from there.
  #
  # order_amendments = search_cache_via_nested_field('Order', ['Opportunity', 'SBQQ__AmendedContract__c'], "800J0000001WZtpIAG")
  # ordered_order_amendments = sort_via_nested_fields(order_amendments, 'Order', ['SBQQ__Quote__c', 'SBQQ__StartDate__c'])
  def sort_via_nested_fields(sf_objects, sf_object_type, nested_fields, ascending)
    # TODO
  end

  # ie
  # order_amendment_query = sf.query(
  #   <<~EOL
  #     SELECT Id FROM #{SF_ORDER}
  #     WHERE Opportunity.SBQQ__AmendedContract__c = '#{sf_contract_id}'
  #   EOL
  # )
  #
  # search_cache_via_nested_field('Order', ['Opportunity', 'SBQQ__AmendedContract__c'], "800J0000001WZtpIAG")
  def search_cache_via_nested_field(sf_object_type, nested_fields, search_value)

    matched_objects = []

    if enabled?
      @cache[sf_object_type].keys.each do |id|
        sf_object = @cache[sf_object_type][id]
        field_value = sf_object
        loop do
          # If this field didn't exist on this object, move to the next
          if field_value.nil?
            break
          end

          # Go one level deeper
          field_value = field_value[nested_fields.shift]

          # If we are at the last field, this is the one to compare
          if nested_fields.empty?
            break
          end
        end

        if field_value == search_value
          matched_objects.append(sf_object)
        end
      end
    end

    if matched_objects.empty?
      missed_objects = backoff do
        @user.sf_client.query(
          <<~EOL
            SELECT Id FROM #{sf_object_type}
            WHERE #{nested_fields.join('.')} = '#{search_value}'
          EOL
        )
      end

      cache_list_of_records(missed_objects)
      matched_objects = missed_objects.map {|i| get_record_from_cache(sf_object_type, i[SF_ID]) }
      # TODO: log if we caught missed objs
    end

    matched_objects
  end

  # ex: get_relationship_from_cache(sf_order_id, 'OrderId', 'OrderLineItems')
  def get_related_record_from_cache(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)

    if enabled?
      cached_records = pull_related_objects_from_cache(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)

      if cached_records.count > 1
        raise Integrations::Errors::ImpossibleState.new("exact ID match yields two cached records")
      end

      return cached_records.first unless cached_records.empty?
    end

    cache_missed_relationships(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type).first
  end

  # ex: get_relationships_from_cache(sf_order_id, 'OrderId', 'OrderLineItems')
  def get_related_records_from_cache(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)

    if enabled?
      cached_records = pull_related_objects_from_cache(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)

      return cached_records unless cached_records.empty?
    end

    cache_missed_relationships(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)
  end

  # ex: get_record_from_cache("OrderLine", "0015e00000NNGYwAAP")
  sig { params(sf_record_type: String, sf_record_id: String).returns(T.untyped) }
  def get_record_from_cache(sf_record_type, sf_record_id)
    @cache[sf_record_type] ||= {}

    if enabled?
      cached_record = @cache[sf_record_type][sf_record_id]
      return cached_record if cached_record
    end

    # Log if we find but didnt have in cache
    missed_object = backoff { @user.sf_client.find(sf_record_type, sf_record_id) }

    # TODO; Cache missed object even when cache is disabled
    # https://jira.corp.stripe.com/browse/PLATINT-1878
    if !missed_object.nil? && enabled?
      metadata = {
        sf_record_type: sf_record_type,
        sf_record_id: sf_record_id,
      }
      message = "Missed cache for SF Object, but found when reaching out to Salesforce API"
      report_missing_cache(message, metadata, missing_record_ids: [sf_record_id])
      @cache[sf_record_type][sf_record_id] = missed_object
    end

    missed_object
  end

  sig { returns(T::Boolean) }
  private def enabled?
    @user.feature_enabled?(FeatureFlags::SF_CACHING)
  end

  private def cache_list_of_records(record_list)
    record_list.each do |record|
      @cache[record.sobject_type] ||= {}
      @cache[record.sobject_type][record.Id] = record
    end
  end

  private def generate_batch_service_url
    url = "/services/apexrest/batchApi"
    if @user.salesforce_namespace.present? && @user.salesforce_namespace != SalesforceNamespaceOptions::NONE.serialize
      url = "/services/apexrest/#{@user.salesforce_namespace}/batchApi"
    end

    url
  end

  private def retrieve_related_objects_from_salesforce(sf_object)

    args = case sf_object.sobject_type
    when SF_ORDER
      {order_ids: [sf_object.Id]}
    when SF_PRODUCT
      {product2_ids: [sf_object.Id]}
    when SF_PRICEBOOK_ENTRY
      {product2_ids: [sf_object.Product2Id]}
    else
      raise ArgumentError.new("unexpected salesforce object type when pulling cache #{sf_object.sobject_type}")
    end

    url = generate_batch_service_url

    response = backoff { @user.sf_client.post(url, args) }
    response.body
  end

  private def previously_invalidated?(sf_object_ids)
    sf_object_ids.each do |id|
      if @previously_invalidated.include? id
        return true
      end
    end
    false
  end


  private def report_missing_cache(message, metadata, missing_record_ids: [])
    Integrations::ErrorContext.report_edge_case(message, metadata: metadata)

    # Error out in tests, provide the meatdata for alerting us to what exactly is missing.
    if (ENV['CI'] || Rails.env.test?) && !previously_invalidated?(missing_record_ids)
      # TODO: Uncomment the line below once the batch service is updated to support tiered pricing
      # raise Integrations::Errors::TranslatorError.new(message, metadata: metadata)
    end
  end

  private def cache_missed_relationships(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)
    record_list = backoff do
      @user.sf_client.query(
        <<~EOL
          SELECT Id
          FROM #{related_sf_object_type}
          WHERE #{foreign_key_field} = '#{foreign_key_sf_object_id}'
        EOL
      )
    end

    record_list = record_list.map(&:Id).map do |related_object_id|
      # SOQL does not support *, so we must use find() to get all fields.
      # We can't pass a field list, as we need to make sure we have all
      # of the user specific custom fields incase they are used in mappings.
      backoff { @user.sf_client.find(related_sf_object_type, related_object_id) }
    end

    cache_list_of_records(record_list)

    if record_list.length > 0 && enabled?
      metadata = {
        related_sf_object_type: related_sf_object_type,
        foreign_key_field: foreign_key_field,
        root_object_id: foreign_key_sf_object_id,
      }
      message = "Missed cache for relational objects, but found when reaching out to Salesforce API"
      report_missing_cache(message, metadata, missing_record_ids: record_list.map(&:Id))
    end

    record_list
  end

  private def pull_related_objects_from_cache(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)
    if @cache[related_sf_object_type].nil?
      @cache[related_sf_object_type] = {}
      cache_missed_relationships(foreign_key_sf_object_id, foreign_key_field, related_sf_object_type)
    end

    # ie SELECT * FROM Opportunity WHERE SBQQ__AmendedContract__c = foreign_key_sf_object_id
    @cache[related_sf_object_type].select do |id|
      @cache[related_sf_object_type][id].send(foreign_key_field) == foreign_key_sf_object_id
    end.values
  end
end
