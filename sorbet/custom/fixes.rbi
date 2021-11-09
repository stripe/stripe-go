# typed: strict

class ActionController::Base
  # should be auto included, but sorbet is acting up
  #   ActionController::RequestForgeryProtection::ClassMethods
  def self.protect_from_forgery(*args); end
end

class StripeForce::User
  sig { returns(T::Array[String]) }
  def feature_flags; end
end