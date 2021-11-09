# frozen_string_literal: true
# typed: strict

class ApplicationController < ActionController::Base
  include Integrations::ErrorContext
  include Integrations::Log

  # Prevent CSRF attacks by raising an exception.
  # For APIs, you may want to use :null_session instead.
  protect_from_forgery with: :exception
end
