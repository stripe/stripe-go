# frozen_string_literal: true
# typed: true

require_relative './controller'

module Api
  class ConfigurationsController < Controller
    wrap_parameters false

    before_action(:set_error_context, :create_user_reference)
    skip_before_action(:create_user_reference, only: [:post_install])
    before_action(:ensure_json_content, only: [:update, :post_install, :translate])

    rescue_from ActionController::ParameterMissing do
      T.bind(self, Api::ConfigurationsController)

      log.error 'invalid input parameters'
      head :bad_request
    end

    def translate_all
      sf_record_type = params.require(:object_type)

      if ![SF_ORDER, SF_ACCOUNT, SF_PRODUCT, SF_PRICEBOOK_ENTRY].include?(sf_record_type)
        log.error 'invalid object type', object_type: sf_record_type
        head :bad_request
        return
      end

      head :ok
    end

    # TODO should really belong in a separate controller, but this is the only method that doesn't fit so we are stuffing it here
    def translate
      sf_record_type, sf_record_ids = params.require([:object_type, :object_ids])

      if ![SF_ORDER, SF_ACCOUNT, SF_PRODUCT, SF_PRICEBOOK_ENTRY].include?(sf_record_type)
        log.error 'invalid object type', object_type: sf_record_type
        head :bad_request
        return
      end

      if !sf_record_ids.is_a?(Array)
        log.error 'must send array of object IDs', object_ids: sf_record_ids
        head :bad_request
        return
      end

      sf_record_ids.each do |sf_record_id|
        # TODO validate that SF IDs look correct? Quick regex sanity check?
        log.info 'queuing object', salesforce_id: sf_record_id, salesforce_type: sf_record_type
        SalesforceTranslateRecordJob.work(@user, sf_record_type, sf_record_id)
      end

      head :ok
    end

    # called by salesforce after package install
    def post_install
      saleforce_key = request.headers[SALESFORCE_KEY_HEADER]

      if saleforce_key != ENV.fetch('SF_MANAGED_PACKAGE_API_KEY')
        log.error 'invalid managed package API key'
        head :not_found
        return
      end

      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
        nil
      end

      salesforce_organization_key = params.require(:key)

      if salesforce_organization_key.blank?
        log.error 'valid request, but org key is blank'
        head :bad_request
        return
      end

      user = StripeForce::User.find_or_new(salesforce_account_id: salesforce_account_id)

      set_error_context(user: user)

      # this can happen if two users are in the setup page at the same time
      if user.new?
        report_edge_case("updating api key for user, but is already set")
      end

      if request.headers[SALESFORCE_INSTANCE_TYPE_HEADER].blank? || request.headers[SALESFORCE_PACKAGE_NAMESPACE_HEADER].nil?
        report_edge_case("important headers are blank on post install", metadata: request.headers)
      end

      # these settings are very important, especially the organization key
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE] = salesforce_instance_type_from_headers(request.headers[SALESFORCE_INSTANCE_TYPE_HEADER])
      user.connector_settings[CONNECTOR_SETTING_SALESFORCE_NAMESPACE] = subdomain_namespace_from_param(request.headers[SALESFORCE_PACKAGE_NAMESPACE_HEADER])
      user.salesforce_organization_key = salesforce_organization_key

      user.save

      head :ok
    end

    def show
      render json: user_configuration_json
    end

    def update
      update_hash = params.permit(settings: {}, field_defaults: {}, field_mappings: {}).to_h

      if update_hash.keys != %w{field_defaults field_mappings} && update_hash.keys != ['settings']
        log.info 'invalid parameters passed', keys: update_hash.keys
        head :bad_request
        return
      end

      # `settings` is too general for our model, but unnecessary in the API schema
      # this is why `connector_settings` is used on the backend but `settings` is the frontend key
      if update_hash.key?("settings")
        # the existing settings should be merged, not replaced
        update_hash["connector_settings"] = @user.connector_settings.merge(update_hash.delete("settings"))
      end

      # top-level keys not passed by the mapper should be preserved, but all 2nd level keys should be removed
      if update_hash.key?('field_mappings')
        update_hash['field_mappings'] = @user.field_mappings.merge(update_hash['field_mappings'])
      end

      @user.update(update_hash)

      render json: user_configuration_json
    end

    private def create_user_reference
      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]
      salesforce_api_key = request.headers[SALESFORCE_KEY_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
        return
      end

      if salesforce_api_key.blank?
        log.warn 'no salesforce api key specified'
        head :not_found
        return
      end

      # DB enforces that SF org IDs must be unique
      @user = StripeForce::User.find(salesforce_account_id: salesforce_account_id)

      if @user.blank?
        log.warn 'invalid user ID specified', salesforce_account_id: salesforce_account_id
        head :not_found
        return
      end

      if @user.salesforce_organization_key != salesforce_api_key
        log.error 'api key does not match user'
        # TODO until the issue with the package is resolved, this needs to stay the way it is
        # head :not_found
        # return
      end

      set_error_context(user: @user)
    end

    private def user_configuration_json
      {
        salesforce_account_id: @user.salesforce_account_id,
        field_mappings: @user.field_mappings,
        field_defaults: @user.field_defaults,
        default_mappings: @user.default_mappings,
        required_mappings: @user.required_mappings,
        feature_flags: @user.feature_flags,
        connection_status: {
          # TODO use dynamic connection status
          salesforce: @user.salesforce_token.present?,
          stripe: @user.stripe_account_id.present?,

          last_synced: Time.now.to_i,
          stripe_account_id: @user.stripe_account_id,
        },
        settings: @user.connector_settings,
      }
    end

    private def ensure_json_content
      return if request.content_type == 'application/json' && request.accept.include?('application/json')

      # TODO ideally, we could just inspect the `format`, but it's broken: https://github.com/rails/rails/pull/14540/files
      # return if request.format == :json

      log.warn 'not json, rejecting request'
      head :not_acceptable
    end

    private def salesforce_instance_type_from_headers(raw_header)
      SFInstanceTypes.try_deserialize(raw_header)&.serialize
    end
  end
end
