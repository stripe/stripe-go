# frozen_string_literal: true
# typed: true

require_relative '../constants'
require_relative './metadata'
require_relative './utilities/stripe_util'
require_relative './utilities/salesforce_util'

require_relative './mapper'

class StripeForce::Translate
  extend T::Sig

  include Integrations::Log
  include Integrations::ErrorContext
  include Integrations::Utilities::Currency
  include Integrations::Utilities::StripeUtil

  include StripeForce::Constants
  include StripeForce::Utilities::StripeUtil
  include StripeForce::Utilities::SalesforceUtil

  # convenience method for console debugging
  def self.perform_inline(user, sf_id)
    locker = Integrations::Locker.new(user)

    translator = self.new(user, locker)

    sf_type = StripeForce::Utilities::SalesforceUtil.salesforce_type_from_id(user, sf_id)
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

  sig { returns(Integrations::Locker) }
  attr_reader :locker

  def sf
    @user.sf_client
  end

  sig { returns(CacheService) }
  def cache_service
    @cache_service ||= CacheService.new(@user)
  end

  def translate(sf_object)
    set_error_context(user: @user, integration_record: sf_object)

    # Cache Related Objects
    cache_service.cache_for_object(sf_object)

    Integrations::Metrics::Writer.instance.time(
      Integrations::Metrics::RESOURCE_TRANSLATION_TIME,
      dimensions: {salesforce_account_id: @user.salesforce_account_id, stripe_account_id: @user.stripe_account_id, livemode: @user.livemode, salesforce_object_type: sf_object.sobject_type}
    ) do
        catch_errors_with_salesforce_context(primary: sf_object) do
          case sf_object.sobject_type
          when SF_ORDER
            translate_order(sf_object)
          when SF_PRODUCT
            translate_product(sf_object)
          when SF_PRICEBOOK_ENTRY
            translate_pricebook(sf_object)
          when SF_ACCOUNT
            translate_account(sf_object)
          else
            raise "unsupported translation type #{sf_object.sobject_type}"
          end
        end
    end
  end

  sig { params(primary: T.nilable(Restforce::SObject), secondary: T.nilable(Restforce::SObject), block: T.proc.returns(T.untyped)).returns(T.untyped) }
  def catch_errors_with_salesforce_context(primary: nil, secondary: nil, &block)
    if primary && @origin_salesforce_object
      raise Integrations::Errors::ImpossibleInternalError.new("origin object already set, exiting")
    end

    if primary && secondary
      raise Integrations::Errors::ImpossibleInternalError.new("should not set primary and secondary salesforce context at the same time")
    end

    if !primary && !secondary
      raise Integrations::Errors::ImpossibleInternalError.new("either primary or secondary context should be passed")
    end

    # TODO once we rework the `create_user_failure` logic we may be able to stop using instance variable
    if secondary
      # TODO maybe pass content to sentry as well?
      original_secondary_object = @secondary_salesforce_object
      @secondary_salesforce_object = secondary

      if original_secondary_object
        log.debug "overwriting secondary object"
      end

      # if we are a secondary wrapper, do not use the same `rescue` logic below
      begin
        # the return result is often a Stripe object that we need to reference downstream, which is why we are careful to retain & return the return value of the incoming block
        result = yield
      end

      @secondary_salesforce_object = original_secondary_object

      return result
    end

    @origin_salesforce_object = primary

    # `@origin_salesforce_object` is used as the primary object by `create_user_failure`, so we only use that
    # if no secondary record is defined. The secondary context could be set within the `yield` which is why
    # we duplicate the selection logic.

    begin
      result = yield
    rescue StripeForce::Errors::RawUserError => e
      # this exception indicates an error that is safe to display to the user
      create_user_failure(
        salesforce_object: @secondary_salesforce_object || @origin_salesforce_object,
        message: e.message,
        error: e
      )

      # convert to standard UserError to make it clear that it's been reported and converted
      raise StripeForce::Errors::UserError.new(e.message)
    rescue Integrations::Errors::MissingRequiredFields => e
      create_user_failure(
        # in the case of missing fields, we know *exactly* which fields are missing
        salesforce_object: e.salesforce_object,
        message: "The following required fields are missing from this Salesforce record: #{e.missing_salesforce_fields.join(', ')}",
        error: e
      )

      raise
    rescue Restforce::ResponseError => e
      create_user_failure(
        salesforce_object: @secondary_salesforce_object || @origin_salesforce_object,
        message: e.message,
        error: e
      )

      raise
    rescue Stripe::StripeError => e
      # TODO probably remove this in the future, but I want to understand the error code details that are coming through
      log.warn 'stripe api error',
        code: e.code,
        message: e.message

      create_user_failure(
        salesforce_object: @secondary_salesforce_object || @origin_salesforce_object,
        message: "Stripe error occurred: #{e.message} #{e.request_id}",
        error: e
      )

      raise
    rescue => e
      if @user.feature_enabled?(FeatureFlags::CATCH_ALL_ERRORS) && e.class != Integrations::Errors::LockTimeout
        create_user_failure(
          salesforce_object: @secondary_salesforce_object || @origin_salesforce_object,
          message: "Translation error occurred",
          error: e
        )
      else
        # log this error but don't create a sync record
        log.error metric: transform_error_into_metric(e), stripe_account_id: @user.stripe_account_id, error_message: e.message, error_class: e.class
      end

      raise
    ensure
      @origin_salesforce_object = nil

      # if an exception is raised within a secondary context block, this will not be unset
      @secondary_salesforce_object = nil
    end

    result
  end

  def transform_error_into_metric(error)
    prefix = "translation.error"

    if error.nil?
      return prefix
    end

    suffix = ""
    error_class = error.class
    error_class_as_string = error_class.to_s

    if error_class_as_string.starts_with?("Stripe::StripeError")
      suffix = ".stripe"
    elsif error_class_as_string.starts_with?("Restforce")
      suffix = ".salesforce"
    elsif error_class == StripeForce::Errors::RawUserError || error_class == StripeForce::Errors::UserError || error_class == Integrations::Errors::MissingRequiredFields
      suffix = ".user"
    elsif error_class == Integrations::Errors::UnhandledEdgeCase || error_class == Integrations::Errors::TranslatorError
      suffix = ".system"
    end

    "#{prefix}#{suffix}"
  end

  sig { params(salesforce_object: Restforce::SObject, message: String, error: T.nilable(Exception)).void }
  def create_user_failure(salesforce_object:, message:, error: nil)
    if @origin_salesforce_object&.Id.blank?
      raise "origin salesforce object is blank, cannot record error"
    end

    log.error 'translation failed, creating sf sync record', {
      metric: transform_error_into_metric(error),
      stripe_account_id: @user.stripe_account_id,
      salesforce_account_id: @user.salesforce_account_id,
      livemode: @user.livemode,
      salesforce_object_id: salesforce_object.Id,
      salesforce_object_type: salesforce_object.sobject_type,
      error_message: message,
      error_class: error.class,
    }

    compound_external_id = generate_compound_external_id(@origin_salesforce_object, salesforce_object)

    log.debug 'creating sync record for failure', external_id: compound_external_id

    # interestingly enough, if the external ID field does not exist we'll get a NOT_FOUND response
    # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_upsert.htm

    sf_sync_record_id = backoff do
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
          SyncRecordFields::RESOLUTION_STATUS => SyncRecordResolutionStatuses::ERROR,
        }.transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field))
      )
    end

    log.debug 'sync record created', sync_record_id: sf_sync_record_id
  end

  sig { params(salesforce_object: Restforce::SObject, stripe_object: Stripe::APIResource).void }
  def create_user_success(salesforce_object:, stripe_object:)
    if @origin_salesforce_object&.Id.blank?
      raise "origin salesforce object is blank, cannot record success"
    end

    # TODO we want to just skip order lines
    unless @origin_salesforce_object.Id == salesforce_object.Id
      log.info 'skipping successful sync record creation in sf for child-object', {
        primary_salesforce_id: @origin_salesforce_object.Id,
        secondary_salesforce_id: salesforce_object.Id,
      }
      return
    end

    message = "Sync Successful, created Stripe #{stripe_object.class.to_s.demodulize.titleize} Object with ID: #{stripe_object.id}"

    log.info 'translation success', {
      metric: 'translation.success',
      stripe_account_id: @user.stripe_account_id,
      salesforce_account_id: @user.salesforce_account_id,
      livemode: @user.livemode,
      salesforce_object_id: salesforce_object.Id,
      salesforce_object_type: salesforce_object.sobject_type,
    }
    compound_external_id = generate_compound_external_id(@origin_salesforce_object, salesforce_object)

    log.debug 'creating sync record'

    # interestingly enough, if the external ID field does not exist we'll get a NOT_FOUND response
    # https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/dome_upsert.htm

    backoff do
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
          SyncRecordFields::RESOLUTION_STATUS => SyncRecordResolutionStatuses::SUCCESS,
        }.transform_keys(&:serialize).transform_keys(&method(:prefixed_stripe_field))
      )
    end
  end

  sig do
    params(salesforce_object: Restforce::SObject, message: String, error_class: T.nilable(T.class_of(Integrations::Errors::BaseIntegrationError)))
    # noreturn indicates an exception is thrown
    .returns(T.noreturn)
  end
  def throw_user_failure!(salesforce_object:, message:, error_class: nil)
    error_class ||= Integrations::Errors::UserError

    error = error_class.new(
      message,
      salesforce_object: salesforce_object
    )

    create_user_failure(
      salesforce_object: salesforce_object,
      message: message,
      error: error
    )

    raise error
  end

  sig { params(sf_object: T.untyped, stripe_object: Stripe::APIResource, additional_salesforce_updates: Hash).void }
  def update_sf_stripe_id(sf_object, stripe_object, additional_salesforce_updates: {})
    stripe_id_field = prefixed_stripe_field(GENERIC_STRIPE_ID)
    stripe_object_id = stripe_object.id
    if sf_object[stripe_id_field]
      if sf_object[stripe_id_field] == stripe_object_id
        log.info 'stripe id already exists in salesforce, no update',
          stripe_id: sf_object[stripe_id_field],
          field_name: stripe_id_field
        return
      end

      log.info 'stripe id is different than existing stripe id in salesforce, overwriting',
        old_stripe_id: sf_object[stripe_id_field],
        new_stripe_id: stripe_object_id,
        field_name: stripe_id_field
    end

    backoff do
      sf.update!(sf_object.sobject_type, {
        SF_ID => sf_object.Id,
        stripe_id_field => stripe_object_id,
      }.merge(additional_salesforce_updates))
    end

    create_user_success(
      salesforce_object: sf_object,
      stripe_object: stripe_object
    )

    # the cached object is missing the new Stripe ID field value, we could append but it's safer to invalidate and refetch when needed.
    cache_service.invalidate_cache_object(sf_object[SF_ID])

    log.info 'updated salesforce with stripe id',
      salesforce_object: sf_object,
      stripe_id: stripe_object.id
  end

  def construct_stripe_object(stripe_class:, salesforce_object:, additional_stripe_params: {})
    stripe_fields = StripeForce::Utilities::SalesforceUtil.extract_salesforce_params!(mapper, salesforce_object, stripe_class)

    stripe_object = stripe_class.construct_from(additional_stripe_params)

    # the keys (stripe references) in the resulting hash could be dot-paths, so let's assign them using the mapper
    mapper.assign_values_from_hash(stripe_object, stripe_fields)

    stripe_object.metadata = Metadata.stripe_metadata_for_sf_object(@user, salesforce_object)

    # Prices and Coupons require currency fields
    if [Stripe::Price, Stripe::Coupon].include?(stripe_object.class)
      stripe_object['currency'] = Integrations::Utilities::Currency.currency_for_sf_object(@user, salesforce_object)
    end

    apply_mapping(stripe_object, salesforce_object)

    sanitize(stripe_object)

    stripe_object
  end

  def update_stripe_object(stripe_class:, stripe_object_id:, sf_object:)
    # creates the stripe object from the sf_object
    stripe_object = construct_stripe_object(stripe_class: stripe_class, salesforce_object: sf_object)

    log.info 'updating stripe object', stripe_id: stripe_object_id, stripe_object_type: stripe_class

    # updates only the parameters passed in and will leave any parameters not passed in unchanged
    # NOTE the stripe_object should not contain it's id or the Stripe update call will fail
    #      since we are not allowed to update the id of an existing object
    # NOTE the individual keys in the stripe_object metadata field will be replaced and not the entire metadata field
    updated_stripe_object = catch_errors_with_salesforce_context(secondary: sf_object) do
      stripe_class.update(
        stripe_object_id,
        stripe_object.to_hash,
        @user.stripe_credentials
      )
    end

    log.info 'updated stripe object', stripe_id: stripe_object_id, stripe_object_type: stripe_class

    updated_stripe_object
  end

  def create_stripe_object(stripe_class, sf_object, additional_stripe_params: {})
    stripe_object = construct_stripe_object(stripe_class: stripe_class, salesforce_object: sf_object, additional_stripe_params: additional_stripe_params)

    if block_given?
      yield(stripe_object)
    end

    log.info 'creating stripe object', salesforce_object: sf_object, stripe_object_type: stripe_class

    # there's a decent chance this causes us issues down the road: we shouldn't be using `construct_from`
    # and then converting the finalized object into a parameters hash. However, without using `construct_from`
    # we'd have to pass the object type around when mapping which is equally as bad.

    created_stripe_object = catch_errors_with_salesforce_context(secondary: sf_object) do
      stripe_class.create(
        stripe_object.to_hash,
        @user.stripe_credentials
      )
    end

    log.info 'created stripe object', stripe_id: created_stripe_object.id

    created_stripe_object
  end

  sig { params(stripe_class: T.class_of(Stripe::APIResource), sf_object: Restforce::SObject).returns(T.nilable(Stripe::APIResource)) }
  def retrieve_from_stripe(stripe_class, sf_object)
    stripe_id = sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)]
    return if stripe_id.nil?

    # TODO need to cast to unsafe other sorbet thinks the code below is unreachable
    # by default tiers are not returned, but this is an important parameter for price comparison purposes
    additional_retrieve_params = {}
    if stripe_class == Stripe::Price
      additional_retrieve_params = {expand: %w{tiers}}
    end

    begin
      # TODO use application sync record here as well
      stripe_record = stripe_class.retrieve(
        {id: sf_object[prefixed_stripe_field(GENERIC_STRIPE_ID)]}.merge(additional_retrieve_params),
        @user.stripe_credentials
      )

      if stripe_record.deleted?
        log.error 'attempting to retrieve deleted Stripe object, raising error',
          deleted_id: stripe_record.id,
          deleted_resource_type: stripe_record.class.to_s

        throw_user_failure!(
          salesforce_object: @origin_salesforce_object || sf_object,
          message: "During translation we attempted to fetch a related #{stripe_record.class} record with ID #{stripe_record.id}, but found that it was deleted."
        )
      end

      # if this ID was provided to us by the user, the metadata does not exist in Stripe
      stripe_record_metadata = Metadata.stripe_metadata_for_sf_object(@user, sf_object)

      salesforce_regex = %r{https://.*?.my.salesforce.com}

      stripe_record_metadata_copy = stripe_record_metadata.deep_dup
      stripe_record_metadata_copy.each do |key, value|
        stripe_record_metadata_copy[key] = value.gsub(salesforce_regex, '')
      end

      metadata_from_stripe = stripe_record.metadata.to_h.deep_dup
      metadata_from_stripe.each do |key, value|
        metadata_from_stripe[key] = value.gsub(salesforce_regex, '')
      end

      # `to_h` on the StripeObject symbolizes all of the keys
      # <= is a special `includes` operator https://stackoverflow.com/questions/7584801/how-to-check-if-an-hash-is-completely-included-in-another-hash
      unless stripe_record_metadata_copy.symbolize_keys <= metadata_from_stripe
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

  # according to the salesforce documentation, if this field is non-nil ("empty") then it's a subscription item
  def recurring_item?(sf_object)
    if ![SF_ORDER_ITEM, SF_PRODUCT].include?(sf_object.sobject_type)
      raise "recurring definition is specified on the order item or product level only"
    end

    # pull field from the mapper. if field not nil, then is subscription item
    !mapper.extract_salesforce_object_field(sf_object, CPQ_QUOTE_SUBSCRIPTION_PRICING).nil?
  end

  # TODO allow for multiple records to be linked?
  sig { params(record_to_map: Stripe::APIResource, source_record: Restforce::SObject, compound_key: T.nilable(T::Boolean)).void }
  def apply_mapping(record_to_map, source_record, compound_key: false)
    mapper.apply_mapping(record_to_map, source_record, compound_key: compound_key)
  end

  sig { returns(StripeForce::Mapper) }
  def mapper
    @mapper ||= StripeForce::Mapper.new(@user, cache_service)
  end

  sig { params(stripe_record: Stripe::StripeObject).void }
  def sanitize(stripe_record)
    StripeForce::Sanitizer.sanitize(stripe_record)
  end
end
