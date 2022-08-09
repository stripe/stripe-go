# frozen_string_literal: true
# typed: strict
module StripeForce
  module Errors
    class UserError < Integrations::Errors::BaseIntegrationError; end
    class RawUserError < UserError; end
  end
end
