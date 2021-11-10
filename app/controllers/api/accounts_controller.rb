# frozen_string_literal: true
# typed: true

require_relative './controller'

module Api
  class AccountsController < Controller
    def show
      user = T.must(StripeForce::User.first)

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

    def update; end
  end
end
