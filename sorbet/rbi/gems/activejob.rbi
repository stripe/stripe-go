# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: strict
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/activejob/all/activejob.rbi
#
# activejob-6.1.4.1

module ActiveJob
  def self.gem_version; end
  def self.version; end
  extend ActiveSupport::Autoload
end
module ActiveJob::VERSION
end
class ActiveJob::Railtie < Rails::Railtie
end
module ActiveJob::Core
  def arguments; end
  def arguments=(arg0); end
  def arguments_serialized?; end
  def deserialize(job_data); end
  def deserialize_arguments(serialized_args); end
  def deserialize_arguments_if_needed; end
  def enqueued_at; end
  def enqueued_at=(arg0); end
  def exception_executions; end
  def exception_executions=(arg0); end
  def executions; end
  def executions=(arg0); end
  def initialize(*arguments); end
  def job_id; end
  def job_id=(arg0); end
  def locale; end
  def locale=(arg0); end
  def priority=(arg0); end
  def provider_job_id; end
  def provider_job_id=(arg0); end
  def queue_name=(arg0); end
  def scheduled_at; end
  def scheduled_at=(arg0); end
  def serialize; end
  def serialize_arguments(arguments); end
  def serialize_arguments_if_needed(arguments); end
  def serialized_arguments=(arg0); end
  def timezone; end
  def timezone=(arg0); end
  extend ActiveSupport::Concern
end
module ActiveJob::Core::ClassMethods
  def deserialize(job_data); end
  def set(options = nil); end
end
module ActiveJob::QueueAdapter
  extend ActiveSupport::Concern
end
module ActiveJob::QueueAdapter::ClassMethods
  def assign_adapter(adapter_name, queue_adapter); end
  def queue_adapter; end
  def queue_adapter=(name_or_adapter); end
  def queue_adapter?(object); end
  def queue_adapter_name; end
end
module ActiveJob::QueueName
  def queue_name; end
  extend ActiveSupport::Concern
end
module ActiveJob::QueueName::ClassMethods
  def default_queue_name; end
  def default_queue_name=(val); end
  def queue_as(part_name = nil, &block); end
  def queue_name_from_part(part_name); end
  def self.default_queue_name; end
  def self.default_queue_name=(val); end
end
module ActiveJob::QueuePriority
  def priority; end
  extend ActiveSupport::Concern
end
module ActiveJob::QueuePriority::ClassMethods
  def default_priority; end
  def default_priority=(val); end
  def queue_with_priority(priority = nil, &block); end
  def self.default_priority; end
  def self.default_priority=(val); end
end
class ActiveJob::DeserializationError < StandardError
  def initialize; end
end
class ActiveJob::SerializationError < ArgumentError
end
module ActiveJob::Arguments
  def convert_to_global_id_hash(argument); end
  def custom_serialized?(hash); end
  def deserialize(arguments); end
  def deserialize_argument(argument); end
  def deserialize_global_id(hash); end
  def deserialize_hash(serialized_hash); end
  def serialize(arguments); end
  def serialize_argument(argument); end
  def serialize_hash(argument); end
  def serialize_hash_key(key); end
  def serialize_indifferent_hash(indifferent_hash); end
  def serialized_global_id?(hash); end
  def transform_symbol_keys(hash, symbol_keys); end
  extend ActiveJob::Arguments
end
module ActiveJob::Enqueuing
  def enqueue(options = nil); end
  extend ActiveSupport::Concern
end
module ActiveJob::Enqueuing::ClassMethods
  def job_or_instantiate(*args); end
  def perform_later(*args); end
end
module ActiveJob::Execution
  def perform(*arg0); end
  def perform_now; end
  extend ActiveSupport::Concern
  include ActiveSupport::Rescuable
end
module ActiveJob::Execution::ClassMethods
  def execute(job_data); end
  def perform_now(*args); end
end
module ActiveJob::Callbacks
  def halted_callback_hook(_filter, name); end
  def self.__callbacks; end
  def self.__callbacks?; end
  def self._execute_callbacks; end
  def self._run_execute_callbacks(&block); end
  extend ActiveSupport::Concern
  include ActiveSupport::Callbacks
end
module ActiveJob::Callbacks::ClassMethods
  def after_enqueue(*filters, &blk); end
  def after_perform(*filters, &blk); end
  def around_enqueue(*filters, &blk); end
  def around_perform(*filters, &blk); end
  def before_enqueue(*filters, &blk); end
  def before_perform(*filters, &blk); end
  def inherited(klass); end
end
module ActiveJob::Exceptions
  def determine_delay(seconds_or_duration_or_algorithm:, executions:, jitter: nil); end
  def determine_jitter_for_delay(delay, jitter); end
  def executions_for(exceptions); end
  def retry_job(options = nil); end
  extend ActiveSupport::Concern
end
module ActiveJob::Exceptions::ClassMethods
  def discard_on(*exceptions); end
  def retry_on(*exceptions, wait: nil, attempts: nil, queue: nil, priority: nil, jitter: nil); end
end
class ActiveJob::LogSubscriber < ActiveSupport::LogSubscriber
end
module ActiveJob::Logging
  def logger_tagged_by_active_job?; end
  def tag_logger(*tags); end
  extend ActiveSupport::Concern
end
module ActiveJob::Instrumentation
  def halted_callback_hook(*arg0); end
  def instrument(operation, payload = nil, &block); end
  extend ActiveSupport::Concern
end
module ActiveJob::Timezones
  extend ActiveSupport::Concern
end
module ActiveJob::Translation
  extend ActiveSupport::Concern
end
module ActiveJob::QueueAdapters
  def self.lookup(name); end
  extend ActiveSupport::Autoload
end
class ActiveJob::QueueAdapters::AsyncAdapter
  def enqueue(job); end
  def enqueue_at(job, timestamp); end
  def immediate=(immediate); end
  def initialize(**executor_options); end
  def shutdown(wait: nil); end
end
class ActiveJob::QueueAdapters::AsyncAdapter::JobWrapper
  def initialize(job); end
  def perform; end
end
class ActiveJob::QueueAdapters::AsyncAdapter::Scheduler
  def enqueue(job, queue_name:); end
  def enqueue_at(job, timestamp, queue_name:); end
  def executor; end
  def immediate; end
  def immediate=(arg0); end
  def initialize(**options); end
  def shutdown(wait: nil); end
end
class ActiveJob::Base
  def __callbacks; end
  def __callbacks?; end
  def _enqueue_callbacks; end
  def _perform_callbacks; end
  def _run_enqueue_callbacks(&block); end
  def _run_perform_callbacks(&block); end
  def logger; end
  def logger=(val); end
  def queue_adapter(*args, &block); end
  def queue_name_prefix; end
  def queue_name_prefix=(arg0); end
  def queue_name_prefix?; end
  def rescue_handlers; end
  def rescue_handlers=(arg0); end
  def rescue_handlers?; end
  def self.__callbacks; end
  def self.__callbacks=(value); end
  def self.__callbacks?; end
  def self._enqueue_callbacks; end
  def self._enqueue_callbacks=(value); end
  def self._perform_callbacks; end
  def self._perform_callbacks=(value); end
  def self._queue_adapter; end
  def self._queue_adapter=(value); end
  def self._queue_adapter_name; end
  def self._queue_adapter_name=(value); end
  def self.log_arguments; end
  def self.log_arguments=(value); end
  def self.log_arguments?; end
  def self.logger; end
  def self.logger=(val); end
  def self.priority; end
  def self.priority=(value); end
  def self.priority?; end
  def self.queue_name; end
  def self.queue_name=(value); end
  def self.queue_name?; end
  def self.queue_name_delimiter; end
  def self.queue_name_delimiter=(value); end
  def self.queue_name_delimiter?; end
  def self.queue_name_prefix; end
  def self.queue_name_prefix=(value); end
  def self.queue_name_prefix?; end
  def self.rescue_handlers; end
  def self.rescue_handlers=(value); end
  def self.rescue_handlers?; end
  def self.retry_jitter; end
  def self.retry_jitter=(value); end
  def self.return_false_on_aborted_enqueue(*args, &block); end
  def self.return_false_on_aborted_enqueue=(*args, &block); end
  def self.skip_after_callbacks_if_terminated; end
  def self.skip_after_callbacks_if_terminated=(val); end
  extend ActiveJob::Callbacks::ClassMethods
  extend ActiveJob::Core::ClassMethods
  extend ActiveJob::Enqueuing::ClassMethods
  extend ActiveJob::Exceptions::ClassMethods
  extend ActiveJob::Execution::ClassMethods
  extend ActiveJob::QueueAdapter::ClassMethods
  extend ActiveJob::QueueName::ClassMethods
  extend ActiveJob::QueuePriority::ClassMethods
  extend ActiveSupport::Callbacks::ClassMethods
  extend ActiveSupport::DescendantsTracker
  extend ActiveSupport::Rescuable::ClassMethods
  include ActiveJob::Callbacks
  include ActiveJob::Core
  include ActiveJob::Enqueuing
  include ActiveJob::Exceptions
  include ActiveJob::Execution
  include ActiveJob::Instrumentation
  include ActiveJob::Logging
  include ActiveJob::QueueAdapter
  include ActiveJob::QueueName
  include ActiveJob::QueuePriority
  include ActiveJob::Timezones
  include ActiveJob::Translation
  include ActiveSupport::Callbacks
  include ActiveSupport::Rescuable
end
