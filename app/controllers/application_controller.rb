# frozen_string_literal: true
# typed: strict

require_relative './controller_helpers'

# NOTE API controller uses a different base
class ApplicationController < ActionController::Base
  include ControllerHelpers
  include Integrations::ErrorContext
  include Integrations::Log
  include StripeForce::Constants

  # Prevent CSRF attacks by raising an exception.
  # For APIs, you may want to use :null_session instead.
  protect_from_forgery with: :exception
end
