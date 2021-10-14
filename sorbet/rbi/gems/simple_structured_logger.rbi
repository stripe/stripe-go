# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: strict
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/simple_structured_logger/all/simple_structured_logger.rbi
#
# simple_structured_logger-0.1.5

module SimpleStructuredLogger
  def log; end
  def self.included(klass); end
end
module SimpleStructuredLogger::Configuration
  def expand_context(&block); end
  def expand_log(&block); end
  extend SimpleStructuredLogger::Configuration
end
class SimpleStructuredLogger::Writer
  def debug(msg, opts = nil); end
  def default_tags; end
  def default_tags=(arg0); end
  def error(msg, opts = nil); end
  def generate_log(msg, opts); end
  def info(msg, opts = nil); end
  def initialize; end
  def logger; end
  def logger=(arg0); end
  def reset_context!; end
  def self.allocate; end
  def self.instance; end
  def self.new(*arg0); end
  def set_context(context); end
  def stringify_tags(additional_tags); end
  def warn(msg, opts = nil); end
  extend Singleton::SingletonClassMethods
  include Singleton
end
