# typed: strict

# class ActionController::Base
#   # should be auto included, but sorbet is acting up
#   #   ActionController::RequestForgeryProtection::ClassMethods
#   def self.protect_from_forgery(*args); end
# end

class ActiveSupport::StringInquirer
  def development?; end
end

class Restforce::SObject
  # TODO investigate why this isn't included by default on the type generation
  def [](arg); end
end