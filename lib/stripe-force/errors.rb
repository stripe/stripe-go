# frozen_string_literal: true
# typed: strict
module StripeForce
  module Errors
    # TODO remove this user error, there's alread one in the integration errors
    class UserError < Integrations::Errors::BaseIntegrationError; end
    class RawUserError < UserError; end
  end
end
