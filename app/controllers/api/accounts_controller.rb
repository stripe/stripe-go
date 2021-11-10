# frozen_string_literal: true
# typed: true

require_relative './controller'

module Api
  class AccountsController < Controller
    wrap_parameters false

    def show
      render json: {
        salesforce_account_id: user.salesforce_account_id,
        field_mappings: user.field_mappings,
        field_defaults: user.field_defaults,
        feature_flags: user.feature_flags,
        connection_status: {
          salesforce: true,
          stripe: true,
        },
      }
    end

    def update
      safe_params = params.permit(field_defaults: {}, field_mappings: {})

      user.update(
        field_defaults: safe_params[:field_defaults],
        field_mappings: safe_params[:field_mappings]
      )

      head :ok
    end

    def user
      T.must(StripeForce::User.first)
    end
  end
end
