# typed: strict

class ActionController::Base
  # should be auto included, but sorbet is acting up
  #   ActionController::RequestForgeryProtection::ClassMethods
  def self.protect_from_forgery(*args); end
end

class ActionController::API
  include AbstractController::Rendering
  def render(*args, &block); end
end

class StripeForce::User
  sig { returns(T::Array[String]) }
  def feature_flags; end

  sig { returns(T::Hash[String, T::Hash[String, String]]) }
  def field_mappings; end

  sig { returns(T::Hash[String, T::Hash[String, T.nilable(T.any(String, Integer, T::Boolean))]]) }
  def field_defaults; end
end

class Restforce::SObject
  # TODO investigate why this isn't included by default on the type generation
  def [](arg); end
end