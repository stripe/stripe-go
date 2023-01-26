# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: true
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/pry-remote/all/pry-remote.rbi
#
# pry-remote-0.1.8

module PryRemote
end
class PryRemote::InputProxy < Struct
  def completion_proc=(val); end
  def input; end
  def input=(_); end
  def readline(prompt); end
  def readline_arity; end
  def self.[](*arg0); end
  def self.inspect; end
  def self.members; end
  def self.new(*arg0); end
end
class PryRemote::IOUndumpedProxy
  def <<(data); end
  def completion_proc; end
  def completion_proc=(val); end
  def initialize(obj); end
  def print(*objs); end
  def printf(*args); end
  def puts(*lines); end
  def readline(prompt); end
  def tty?; end
  def write(data); end
  include DRb::DRbUndumped
end
class PryRemote::Client < Struct
  def input; end
  def input=(_); end
  def input_proxy; end
  def kill; end
  def output; end
  def output=(_); end
  def self.[](*arg0); end
  def self.inspect; end
  def self.members; end
  def self.new(*arg0); end
  def stderr; end
  def stderr=(_); end
  def stdout; end
  def stdout=(_); end
  def thread; end
  def thread=(_); end
  def wait; end
end
class PryRemote::Server
  def client; end
  def host; end
  def initialize(object, host = nil, port = nil, options = nil); end
  def object; end
  def port; end
  def run; end
  def self.run(object, host = nil, port = nil, options = nil); end
  def setup; end
  def teardown; end
  def uri; end
end
class PryRemote::CLI
  def capture; end
  def capture?; end
  def host; end
  def initialize(args = nil); end
  def port; end
  def run(input = nil, output = nil); end
  def uri; end
  def wait; end
  def wait?; end
end
class Object < BasicObject
  def pry_remote(host = nil, port = nil, options = nil); end
  def remote_pry(host = nil, port = nil, options = nil); end
end