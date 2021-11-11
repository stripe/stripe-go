# frozen_string_literal: true
# typed: true

require_relative './controller'

module Api
  class AccountsController < Controller
    wrap_parameters false

    before_action(:set_error_context, :create_user_reference)

    def show
      render json: {
        salesforce_account_id: @user.salesforce_account_id,
        field_mappings: @user.field_mappings,
        field_defaults: @user.field_defaults,
        feature_flags: @user.feature_flags,
        connection_status: {
          salesforce: true,
          stripe: true,
        },
      }
    end

    def update
      safe_params = params.permit(field_defaults: {}, field_mappings: {})

      @user.update(
        field_defaults: safe_params[:field_defaults],
        field_mappings: safe_params[:field_mappings]
      )

      head :ok
    end

    private def create_user_reference
      salesforce_account_id = request.headers[SALESFORCE_ACCOUNT_ID_HEADER]

      if salesforce_account_id.blank?
        log.warn 'no salesforce account ID specified'
        head :not_found
      end

      @user = StripeForce::User.find(salesforce_account_id: salesforce_account_id)

      if @user.blank?
        log.warn 'invalid user ID specified'
        head :not_found
      end

      set_error_context(user: @user)
    end
  end
end
