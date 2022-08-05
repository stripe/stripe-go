# frozen_string_literal: true
# typed: true

require_relative '../constants'
require_relative './utilities/metadata'
require_relative './utilities/stripe_util'
require_relative './utilities/salesforce_util'

require_relative './mapper'

class StripeForce::Translate
  extend T::Sig

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities::Currency

  include StripeForce::Constants
  include StripeForce::Utilities::Metadata
  include StripeForce::Utilities::StripeUtil
  include StripeForce::Utilities::SalesforceUtil

  # convenience method for console debugging
  def self.perform_inline(user, sf_id)
    locker = Integrations::Locker.new(user)

    translator = self.new(user, locker)

    sf_type = translator.salesforce_type_from_id(sf_id)
    sf_object = user.sf_client.find(sf_type, sf_id)

    result = translator.translate(sf_object)

    locker.clear_locked_resources

    result
  end

  def self.perform(user:, sf_object:, locker:)
    interactor = self.new(user, locker)
    interactor.translate(sf_object)
  end

  sig { params(user: StripeForce::User, locker: Integrations::Locker).void }
  def initialize(user, locker)
    @user = user
    @locker = locker

    set_error_context(user: user)
  end

  def sf
    @user.sf_client
  end

  def translate(sf_object)
    @origin_salesforce_object = sf_object

    set_error_context(user: @user, integration_record: sf_object)

    case sf_object.sobject_type
    when SF_ORDER
      translate_order(sf_object)
    when SF_PRODUCT
      translate_product(sf_object)
    when SF_PRICEBOOK_ENTRY
      create_price_from_pricebook(sf_object)
    when SF_ACCOUNT
      translate_account(sf_object)
    else
      raise "unsupported translation type #{sf_object.sobject_type}"
    end
  rescue Integrations::Errors::MissingRequiredFields => e
    create_user_failure(
      salesforce_object: e.salesforce_object,
      message: "The following required fields are missing from this Salesforce record: #{e.missing_salesforce_fields.join(', ')}",
    )

    raise
  rescue Restforce::ResponseError => e
    create_user_failure(
      # TODO can we indicate a more specific error object here?
      salesforce_object: @origin_salesforce_object,
      message: e.message
    )

    raise
  end

  # report user failures with the correct secondary SF object context
  def catch_stripe_api_errors(salesforce_object)
    yield
  rescue Stripe::InvalidRequestError => e
    # TODO probably remove this in the future, but I want to understand the error codes that are coming through
    log.warn 'stripe api error',
      code: e.code,
      message: e.message

    create_user_failure(
      salesforce_object: salesforce_object,
      message: "#{e.message} #{e.request_id}"
    )

    raise
  end

  sig { params(salesforce_object: Restforce::SObject, message: String).void }
  def create_user_failure(salesforce_object:, message:)
    if @origin_salesforce_object&.Id.blank?
      raise "origin salesforce object is blank, cannot record error"
    end

    log.error 'translation failed', {
      metric: 'error.user',
      secondary_salesforce_id: salesforce_object.Id,
      secondary_salesforce_type: salesforce_object.sobject_type,
      error_message: message,
    }

    compound_external_id = "#{@origin_salesforce_object.Id}-#{salesforce_object.Id}"

    log.debug 'creating sync record'

    # interestingly enough, if the external ID field does not exist we'll get a NOT_FOUND response
    # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_upsert.htm

    sf.upsert!(
      prefixed_stripe_field(SYNC_RECORD),
      prefixed_stripe_field(SyncRecordFields::COMPOUND_ID.serialize),
      {
        SyncRecordFields::COMPOUND_ID => compound_external_id,

        SyncRecordFields::PRIMARY_RECORD_ID => @origin_salesforce_object.Id,
        SyncRecordFields::PRIMARY_OBJECT_TYPE => @origin_salesforce_object.sobject_type,

        SyncRecordFields::SECONDARY_RECORD_ID => salesforce_object.Id,
        SyncRecordFields::SECONDARY_OBJECT_TYPE => salesforce_object.sobject_type,

        SyncRecordFields::RESOLUTION_MESSAGE => message,
        SyncRecordFields::RESOLUTION_STATUS => 'Error',
      }.transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field))
    )
  end

  sig do
    params(salesforce_object: Restforce::SObject, message: String, error_class: T.nilable(T.class_of(Integrations::Errors::BaseIntegrationError)))
    # noreturn indicates an exception is thrown
    .returns(T.noreturn)
  end
  def throw_user_failure!(salesforce_object:, message:, error_class: nil)
    create_user_failure(
      salesforce_object: salesforce_object,
      message: message,
    )

    # TODO right now all attempt/audit log information is stored in salesforce
    #      if the user chooses to add record history to the sync record
    #      it may make sense to change that at some point in the future

    error_class ||= Integrations::Errors::UserError

    raise error_class.new(
      message,
      salesforce_object: salesforce_object
    )
  end

  sig { params(sf_object: T.untyped, stripe_object: Stripe::APIResource, additional_salesforce_updates: Hash).void }
  def update_sf_stripe_id(sf_object, stripe_object, additional_salesforce_updates: {})
    stripe_id_field = prefixed_stripe_field(GENERIC_STRIPE_ID)

    if sf_object[stripe_id_field]
      log.info 'stripe id already exists on object, overwriting',
        old_stripe_id: sf_object[stripe_id_field],
        new_stripe_id: stripe_object.id,
        field_name: stripe_id_field
    end

    sf.update!(sf_object.sobject_type, {
      SF_ID => sf_object.Id,
      stripe_id_field => stripe_object.id,
    }.merge(additional_salesforce_updates))

    log.info 'updated with stripe id',
      salesforce_object: sf_object,
      stripe_id: stripe_object.id
  end

  def create_stripe_object(stripe_class, sf_object, additional_stripe_params: {}, skip_field_extraction: false)
    stripe_fields = if skip_field_extraction
      {}
    else
      extract_salesforce_params!(sf_object, stripe_class)
    end

    stripe_object = stripe_class.construct_from(additional_stripe_params)

    # the fields in the resulting hash could be dot-paths, so let's assign them using the mapper
    mapper.assign_values_from_hash(stripe_object, stripe_fields)

    stripe_object.metadata = stripe_metadata_for_sf_object(sf_object)

    apply_mapping(stripe_object, sf_object)

    if block_given?
      yield(stripe_object)
    end

    sanitize(stripe_object)

    log.info 'creating stripe object', salesforce_object: sf_object, stripe_object_type: stripe_class

    # there's a decent chance this causes us issues down the road: we shouldn't be using `construct_from`
    # and then converting the finalized object into a parameters hash. However, without using `construct_from`
    # we'd have to pass the object type around when mapping which is equally as bad.

    catch_stripe_api_errors(sf_object) do
      stripe_class.create(stripe_object.to_hash, @user.stripe_credentials)
    end
  end

  sig { params(stripe_class: T.class_of(Stripe::APIResource), sf_object: Restforce::SObject).returns(T.nilable(Stripe::APIResource)) }
  def retrieve_from_stripe(stripe_class, sf_object)
    stripe_id = sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    return if stripe_id.nil?

    begin
      stripe_record = stripe_class.retrieve(sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)], @user.stripe_credentials)

      # if this ID was provided to us by the user, the metadata does not exist in Stripe
      stripe_record_metadata = stripe_metadata_for_sf_object(sf_object)

      # `to_h` on the StripeObject symbolizes all of the keys
      # <= is a special `includes` operator https://stackoverflow.com/questions/7584801/how-to-check-if-an-hash-is-completely-included-in-another-hash
      unless stripe_record_metadata.symbolize_keys <= stripe_record.metadata.to_h
        stripe_record_metadata.each do |k, v|
          if stripe_record.metadata[k] != v
            log.warn 'overwriting metadata value',
              metadata_key: k,
              old_value: stripe_record.metadata[k],
              new_value: v
          end

          stripe_record.metadata[k] = v
        end

        log.info 'no metadata found on valid stripe record reference, adding'
        stripe_record.save
      end

      log.info 'existing stripe record found',
        salesforce_object: sf_object,
        stripe_resource: stripe_record

      stripe_record
    rescue Stripe::InvalidRequestError => e
      raise if e.code != 'resource_missing'
      nil
    end
  end

  # TODO this should be dynamic and pulled from the mapper
  # according to the salesforce documentation, if this field is non-nil ("empty") than it's a subscription item
  def recurring_item?(sf_object)
    if ![SF_ORDER_ITEM, SF_PRODUCT].include?(sf_object.sobject_type)
      raise "recurring definition is specified on the order item or product level only"
    end

    # TODO maybe pull this from the mapper in order to make this customizable?
    !sf_object[CPQ_QUOTE_SUBSCRIPTION_PRICING].nil?
  end

  # service period and billing frequency are decoupled in CPQ
  # both values should be in months, but we want to support days in the future
  def determine_subscription_term_multiplier_for_billing_frequency(subscription_term, billing_frequency)
    # if specified iterations is less than the billing frequency of the stripe price then
    if subscription_term < billing_frequency
      # TODO we should probably create a invoice item price instead? Unsure of the best approach here, we would want the invoice item to be tied to a product?
      #      also, if we take this approach we should process all subscription items instead of just one
      log.info 'iterations is less than price billing frequency, creating a new price for the prorated amount', subscription_term: subscription_term, frequency: billing_frequency
      return 1
    end

    # TODO I expect this to happen in proration cases, i.e. 24mo subscription with a order amendment at mo 6
    if subscription_term % billing_frequency != 0
      throw_user_failure!(
        salesforce_object: @origin_salesforce_object,
        message: "Prorated order amendments are not yet supported"
      )
    end

    # TODO maybe this logic should be different if the term source is customized in salesforce?
    subscription_term / billing_frequency
  end

  sig { params(iterations: Integer, subscription_price_id: String).returns(Integer) }
  def transform_iterations_by_billing_frequency(iterations, subscription_price_id)
    # this is terrible: the best way to figure out the recurrance schedule of a subscription
    # is to pull one of it's items and check the `recurring` hash.
    #
    # Why is this the case?
    #
    # because of the level of field customization we allow, the values for the recurring hash
    # could be *anywhere* so the only alternative would be to regenerate the line items
    # or bubble up the recurrance through the line item helpers, which would mix responsibility
    # and make it more challenging to support multi-frequency subscriptions in the future

    price = Stripe::Price.retrieve(subscription_price_id, @user.stripe_credentials)

    billing_frequency_in_months = StripeForce::Utilities::StripeUtil.billing_frequency_of_price_in_months(price)
    determine_subscription_term_multiplier_for_billing_frequency(iterations, billing_frequency_in_months)
  end

  # param_mapping: { stripe_key_name => salesforce_field_name }
  sig { params(sf_record: Restforce::SObject, stripe_record_or_class: T.any(Class, Stripe::APIResource)).returns(Hash) }
  def extract_salesforce_params!(sf_record, stripe_record_or_class)
    stripe_mapping_key = StripeForce::Mapper.mapping_key_for_record(stripe_record_or_class, sf_record)
    required_mappings = @user.required_mappings[stripe_mapping_key]

    if required_mappings.nil?
      raise "expected mappings for #{stripe_mapping_key} but they were nil"
    end

    # first, let's pull required mappings and check if there's anything missing
    required_data = mapper.build_dynamic_mapping_values(sf_record, required_mappings)

    missing_stripe_fields = required_mappings.select {|k, _v| required_data[k].nil? }

    if missing_stripe_fields.present?
      missing_salesforce_fields = missing_stripe_fields.keys.map {|k| required_mappings[k] }

      raise Integrations::Errors::MissingRequiredFields.new(
        salesforce_object: sf_record,
        missing_salesforce_fields: missing_salesforce_fields
      )
    end

    # then, let's extract optional fields and then merge them in
    default_mappings = @user.default_mappings[stripe_mapping_key]
    return required_data if default_mappings.blank?

    optional_data = mapper.build_dynamic_mapping_values(sf_record, default_mappings)

    required_data.merge(optional_data)
  end

  # TODO allow for multiple records to be linked?
  sig { params(record_to_map: Stripe::APIResource, source_record: Restforce::SObject, compound_key: T.nilable(T::Boolean)).void }
  def apply_mapping(record_to_map, source_record, compound_key: false)
    mapper.apply_mapping(record_to_map, source_record, compound_key: compound_key)
  end

  sig { returns(StripeForce::Mapper) }
  def mapper
    @mapper ||= StripeForce::Mapper.new(@user)
  end

  sig { params(stripe_record: Stripe::StripeObject).void }
  def sanitize(stripe_record)
    @sanitizer ||= StripeForce::Sanitizer.new(@user)
    @sanitizer.sanitize(stripe_record)
  end
end