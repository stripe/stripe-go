# typed: true

# class ActionController::Base
#   # should be auto included, but sorbet is acting up
#   #   ActionController::RequestForgeryProtection::ClassMethods
#   def self.protect_from_forgery(*args); end
# end

class ActiveSupport::StringInquirer
  def development?; end
  def production?; end
end

# TODO for some reason this class is excluded from the autogen's rbis
#      hopefully we can remove this hack in the future
class ActionDispatch::SystemTestCase
  include ::Capybara::DSL
  include ::Capybara::Minitest::Assertions
  include ::ActionDispatch::SystemTesting::TestHelpers::SetupAndTeardown
  include ::ActionDispatch::SystemTesting::TestHelpers::ScreenshotHelper
  include ::Minitest::Reportable
  include ::Minitest::Assertions
end

class ApplicationIntegrationTest
  # include ActionController::Metal
  # sig { returns(ActionDispatch::TestResponse) }
  def response; end

  sig { returns(ActionDispatch::Request::Session) }
  def session; end

  def https!(flag = nil); end

  include Minitest::Assertions
end

class Restforce::SObject
  # TODO investigate why this isn't included by default on the type generation
  def [](arg); end
end

class Array
  # pretty certain delete works for anything, not just symbols or strings, but the original def was just symbols
  sig { params(obj: T.any(String, Symbol)).returns(T.nilable(T.any(String, Symbol)))}
  def delete(obj); end
end

# TODO look into why this isn't autogen'd at all via sorbet
class StripeForce::User
  def self.find_or_new(*args); end
end

class Restforce::SObject
  # this is inherited from Mash, I don't know why sorbet isn't picking this up
  def initialize(source_hash = nil, client = nil, default = nil, &blk); end
end

module GeneratedUrlHelpers
  def omniauth_path(*args); end
end