# frozen_string_literal: true
# typed: true

require_relative './controller'

module Api
  class ConfigurationsController < Controller
    wrap_parameters false

    before_action(:set_error_context, :create_user_reference)
    before_action(:ensure_json_content, only: [:update, :post_install])

    # called by salesforce after package install
    def post_install
      # TODO valid incoming SF API key

      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
      end
    end

    def show
      render json: user_configuration_json
    end

    def update
      field_defaults, field_mappings = params.require([:field_defaults, :field_mappings])

      @user.update(
        field_defaults: field_defaults,
        field_mappings: field_mappings
      )

      render json: user_configuration_json
    end

    private def create_user_reference
      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
      end

      @user = StripeForce::User.find(salesforce_account_id: salesforce_account_id)

      if @user.blank?
        log.warn 'invalid user ID specified', salesforce_account_id: salesforce_account_id
        head :not_found
      end

      set_error_context(user: @user)
    end

    private def user_configuration_json
      {
        salesforce_account_id: @user.salesforce_account_id,
        field_mappings: @user.field_mappings,
        field_defaults: @user.field_defaults,
        default_mapping: {

        },
        required_mapping: {

        },
        feature_flags: @user.feature_flags,
        connection_status: {
          salesforce: true,
          stripe: true,
        },
      }
    end

    private def ensure_json_content
      return if request.format == :json
      head :not_acceptable
    end
  end
end
