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

    # TODO should really belong in a separate controller, but this is the only method that doesn't fit so we are stuffing it here
    def translate
      sf_record_type, sf_record_ids = params.require([:object_type, :object_ids])

      if ![SF_ORDER, SF_ACCOUNT, SF_PRODUCT].include?(sf_record_type)
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
        # TODO validate that SF IDs look correct?
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

      unless user.new?
        report_edge_case("updating api key for user")
      end

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
      update_hash["connector_settings"] = update_hash.delete("settings") if update_hash.key?("settings")

      @user.update(update_hash)

      render json: user_configuration_json
    end

    private def create_user_reference
      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
        return
      end

      @user = StripeForce::User.find(salesforce_account_id: salesforce_account_id)

      # TODO validate incoming API key

      if @user.blank?
        log.warn 'invalid user ID specified', salesforce_account_id: salesforce_account_id
        head :not_found
        return
      end

      set_error_context(user: @user)
    end

    private def user_configuration_json
      {
        salesforce_account_id: @user.salesforce_account_id,
        field_mappings: @user.field_mappings,
        field_defaults: @user.field_defaults,
        default_mappings: @user.default_mappings,
        required_mapping: {},
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
  end
end
