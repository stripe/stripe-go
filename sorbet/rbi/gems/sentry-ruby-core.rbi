# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: true
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/sentry-ruby-core/all/sentry-ruby-core.rbi
#
# sentry-ruby-core-4.8.3

module Sentry
  def self.add_breadcrumb(breadcrumb, **options); end
  def self.apply_patches(config); end
  def self.background_worker; end
  def self.background_worker=(arg0); end
  def self.capture_event(event); end
  def self.capture_exception(exception, **options, &block); end
  def self.capture_message(message, **options, &block); end
  def self.clone_hub_to_current_thread; end
  def self.configuration(*args, &block); end
  def self.configure_scope(&block); end
  def self.csp_report_uri; end
  def self.exception_locals_tp; end
  def self.get_current_client; end
  def self.get_current_hub; end
  def self.get_current_scope; end
  def self.get_main_hub; end
  def self.init(&block); end
  def self.initialized?; end
  def self.integrations; end
  def self.last_event_id; end
  def self.logger; end
  def self.register_integration(name, version); end
  def self.register_patch(&block); end
  def self.registered_patches; end
  def self.sdk_meta; end
  def self.send_event(*args, &block); end
  def self.set_context(*args, &block); end
  def self.set_extras(*args, &block); end
  def self.set_tags(*args, &block); end
  def self.set_user(*args, &block); end
  def self.start_transaction(**options); end
  def self.sys_command(command); end
  def self.utc_now; end
  def self.with_scope(&block); end
end
class Sentry::Error < StandardError
end
class Sentry::ExternalError < Sentry::Error
end
module Sentry::ArgumentCheckingHelper
  def check_argument_type!(argument, expected_type); end
end
module Sentry::LoggingHelper
  def log_debug(message); end
  def log_error(message, exception, debug: nil); end
  def log_info(message); end
  def log_warn(message); end
end
module Sentry::Utils
end
module Sentry::Utils::ExceptionCauseChain
  def self.exception_to_array(exception); end
end
module Sentry::CustomInspection
  def inspect; end
end
class Sentry::DSN
  def csp_report_uri; end
  def envelope_endpoint; end
  def host; end
  def initialize(dsn_string); end
  def path; end
  def port; end
  def project_id; end
  def public_key; end
  def scheme; end
  def secret_key; end
  def server; end
  def to_s; end
  def valid?; end
end
class Sentry::ReleaseDetector
  def self.detect_release(project_root:, running_on_heroku:); end
  def self.detect_release_from_capistrano(project_root); end
  def self.detect_release_from_env; end
  def self.detect_release_from_git; end
  def self.detect_release_from_heroku(running_on_heroku); end
end
class Sentry::Transport
  def discarded_events; end
  def encode(event); end
  def fetch_pending_client_report; end
  def generate_auth_header; end
  def get_item_type(event_hash); end
  def initialize(configuration); end
  def is_rate_limited?(item_type); end
  def last_client_report_sent; end
  def logger; end
  def rate_limits; end
  def record_lost_event(reason, item_type); end
  def send_data(data, options = nil); end
  def send_event(event); end
  include Sentry::LoggingHelper
end
class Sentry::Transport::Configuration
  def encoding; end
  def encoding=(arg0); end
  def faraday_builder; end
  def faraday_builder=(arg0); end
  def http_adapter; end
  def http_adapter=(arg0); end
  def initialize; end
  def open_timeout; end
  def open_timeout=(arg0); end
  def proxy; end
  def proxy=(arg0); end
  def ssl; end
  def ssl=(arg0); end
  def ssl_ca_file; end
  def ssl_ca_file=(arg0); end
  def ssl_verification; end
  def ssl_verification=(arg0); end
  def timeout; end
  def timeout=(arg0); end
  def transport_class; end
  def transport_class=(klass); end
end
class Sentry::LineCache
  def get_file_context(filename, lineno, context); end
  def getline(path, n); end
  def getlines(path); end
  def initialize; end
  def valid_path?(path); end
end
class Sentry::StacktraceBuilder
  def app_dirs_pattern; end
  def backtrace_cleanup_callback; end
  def build(backtrace:, &frame_callback); end
  def context_lines; end
  def convert_parsed_line_into_frame(line); end
  def initialize(project_root:, app_dirs_pattern:, linecache:, context_lines:, backtrace_cleanup_callback: nil); end
  def linecache; end
  def parse_backtrace_lines(backtrace); end
  def project_root; end
end
class Sentry::Configuration
  def app_dirs_pattern; end
  def app_dirs_pattern=(arg0); end
  def async; end
  def async=(value); end
  def background_worker_threads; end
  def background_worker_threads=(arg0); end
  def backtrace_cleanup_callback; end
  def backtrace_cleanup_callback=(arg0); end
  def before_breadcrumb; end
  def before_breadcrumb=(value); end
  def before_send; end
  def before_send=(value); end
  def breadcrumbs_logger; end
  def breadcrumbs_logger=(logger); end
  def capture_exception_frame_locals; end
  def capture_exception_frame_locals=(arg0); end
  def capture_in_environment?; end
  def check_callable!(name, value); end
  def context_lines; end
  def context_lines=(arg0); end
  def csp_report_uri; end
  def debug; end
  def debug=(arg0); end
  def detect_release; end
  def dsn; end
  def dsn=(value); end
  def enabled_environments; end
  def enabled_environments=(arg0); end
  def enabled_in_current_env?; end
  def environment; end
  def environment=(environment); end
  def environment_from_env; end
  def error_messages; end
  def errors; end
  def exception_class_allowed?(exc); end
  def exclude_loggers; end
  def exclude_loggers=(arg0); end
  def excluded_exception?(incoming_exception); end
  def excluded_exception_classes; end
  def excluded_exceptions; end
  def excluded_exceptions=(arg0); end
  def gem_specs; end
  def get_exception_class(x); end
  def init_dsn(dsn_string); end
  def initialize; end
  def inspect_exception_causes_for_exclusion; end
  def inspect_exception_causes_for_exclusion=(arg0); end
  def inspect_exception_causes_for_exclusion?; end
  def linecache; end
  def linecache=(arg0); end
  def logger; end
  def logger=(arg0); end
  def matches_exception?(excluded_exception_class, incoming_exception); end
  def max_breadcrumbs; end
  def max_breadcrumbs=(arg0); end
  def project_root; end
  def project_root=(arg0); end
  def propagate_traces; end
  def propagate_traces=(arg0); end
  def rack_env_whitelist; end
  def rack_env_whitelist=(arg0); end
  def release; end
  def release=(arg0); end
  def run_post_initialization_callbacks; end
  def running_on_heroku?; end
  def safe_const_get(x); end
  def sample_allowed?; end
  def sample_rate; end
  def sample_rate=(arg0); end
  def self.add_post_initialization_callback(&block); end
  def self.post_initialization_callbacks; end
  def send_client_reports; end
  def send_client_reports=(arg0); end
  def send_default_pii; end
  def send_default_pii=(arg0); end
  def send_modules; end
  def send_modules=(arg0); end
  def sending_allowed?; end
  def server=(value); end
  def server_name; end
  def server_name=(arg0); end
  def server_name_from_env; end
  def skip_rake_integration; end
  def skip_rake_integration=(arg0); end
  def stacktrace_builder; end
  def traces_sample_rate; end
  def traces_sample_rate=(arg0); end
  def traces_sampler; end
  def traces_sampler=(arg0); end
  def tracing_enabled?; end
  def transport; end
  def trusted_proxies; end
  def trusted_proxies=(arg0); end
  def valid?; end
  include Sentry::CustomInspection
  include Sentry::LoggingHelper
end
class Sentry::Logger < Logger
  def initialize(*arg0); end
end
class Sentry::ExceptionInterface < Sentry::Interface
  def initialize(values:); end
  def self.build(exception:, stacktrace_builder:); end
  def to_hash; end
end
class Sentry::RequestInterface < Sentry::Interface
  def cookies; end
  def cookies=(arg0); end
  def data; end
  def data=(arg0); end
  def encode_to_utf_8(value); end
  def env; end
  def env=(arg0); end
  def filter_and_format_env(env); end
  def filter_and_format_headers(env); end
  def headers; end
  def headers=(arg0); end
  def initialize(request:); end
  def is_server_protocol?(key, value, protocol_version); end
  def is_skippable_header?(key); end
  def method; end
  def method=(arg0); end
  def query_string; end
  def query_string=(arg0); end
  def read_data_from(request); end
  def self.build(env:); end
  def self.clean_env(env); end
  def url; end
  def url=(arg0); end
end
class Sentry::SingleExceptionInterface < Sentry::Interface
  def initialize(exception:, stacktrace: nil); end
  def module; end
  def self.build_with_stacktrace(exception:, stacktrace_builder:); end
  def stacktrace; end
  def thread_id; end
  def to_hash; end
  def type; end
  def value; end
  include Sentry::CustomInspection
end
class Sentry::StacktraceInterface
  def frames; end
  def initialize(frames:); end
  def inspect; end
  def to_hash; end
end
class Sentry::StacktraceInterface::Frame < Sentry::Interface
  def abs_path; end
  def abs_path=(arg0); end
  def compute_filename; end
  def context_line; end
  def context_line=(arg0); end
  def filename; end
  def filename=(arg0); end
  def function; end
  def function=(arg0); end
  def in_app; end
  def in_app=(arg0); end
  def initialize(project_root, line); end
  def lineno; end
  def lineno=(arg0); end
  def longest_load_path; end
  def module; end
  def module=(arg0); end
  def post_context; end
  def post_context=(arg0); end
  def pre_context; end
  def pre_context=(arg0); end
  def set_context(linecache, context_lines); end
  def to_hash(*args); end
  def to_s; end
  def under_project_root?; end
  def vars; end
  def vars=(arg0); end
end
class Sentry::ThreadsInterface
  def initialize(crashed: nil, stacktrace: nil); end
  def self.build(backtrace:, stacktrace_builder:, **options); end
  def to_hash; end
end
class Sentry::Interface
  def self.inherited(klass); end
  def self.registered; end
  def to_hash; end
end
class Sentry::Backtrace
  def ==(other); end
  def initialize(lines); end
  def inspect; end
  def lines; end
  def self.parse(backtrace, project_root, app_dirs_pattern, &backtrace_cleanup_callback); end
  def to_s; end
end
class Sentry::Backtrace::Line
  def ==(other); end
  def file; end
  def in_app; end
  def in_app_pattern; end
  def initialize(file, number, method, module_name, in_app_pattern); end
  def inspect; end
  def method; end
  def module_name; end
  def number; end
  def self.parse(unparsed_line, in_app_pattern); end
  def to_s; end
end
class Sentry::Utils::RealIp
  def calculate_ip; end
  def filter_trusted_proxy_addresses(ips); end
  def initialize(remote_addr: nil, client_ip: nil, real_ip: nil, forwarded_for: nil, trusted_proxies: nil); end
  def ip; end
  def ips_from(header); end
end
module Sentry::Utils::RequestId
  def self.read_from(env); end
end
class Sentry::Event
  def add_exception_interface(exception); end
  def add_request_interface(env); end
  def add_threads_interface(backtrace: nil, **options); end
  def backtrace; end
  def backtrace=(arg0); end
  def breadcrumbs; end
  def breadcrumbs=(arg0); end
  def calculate_real_ip_from_rack(env); end
  def configuration; end
  def contexts; end
  def contexts=(arg0); end
  def environment; end
  def environment=(arg0); end
  def event_id; end
  def event_id=(arg0); end
  def exception; end
  def extra; end
  def extra=(arg0); end
  def fingerprint; end
  def fingerprint=(arg0); end
  def initialize(configuration:, integration_meta: nil, message: nil); end
  def level; end
  def level=(new_level); end
  def message; end
  def message=(arg0); end
  def modules; end
  def modules=(arg0); end
  def platform; end
  def platform=(arg0); end
  def rack_env=(env); end
  def release; end
  def release=(arg0); end
  def request; end
  def sdk; end
  def sdk=(arg0); end
  def self.get_log_message(event_hash); end
  def self.get_message_from_exception(event_hash); end
  def serialize_attributes; end
  def server_name; end
  def server_name=(arg0); end
  def tags; end
  def tags=(arg0); end
  def threads; end
  def timestamp; end
  def timestamp=(time); end
  def to_hash; end
  def to_json_compatible; end
  def transaction; end
  def transaction=(arg0); end
  def type; end
  def user; end
  def user=(arg0); end
  include Sentry::CustomInspection
end
class Sentry::TransactionEvent < Sentry::Event
  def contexts; end
  def contexts=(arg0); end
  def environment; end
  def environment=(arg0); end
  def event_id; end
  def event_id=(arg0); end
  def extra; end
  def extra=(arg0); end
  def initialize(configuration:, integration_meta: nil, message: nil); end
  def level; end
  def modules; end
  def modules=(arg0); end
  def platform; end
  def platform=(arg0); end
  def release; end
  def release=(arg0); end
  def sdk; end
  def sdk=(arg0); end
  def server_name; end
  def server_name=(arg0); end
  def spans; end
  def spans=(arg0); end
  def start_timestamp; end
  def start_timestamp=(time); end
  def tags; end
  def tags=(arg0); end
  def timestamp; end
  def to_hash; end
  def transaction; end
  def transaction=(arg0); end
  def type; end
  def user; end
  def user=(arg0); end
end
class Sentry::Span
  def data; end
  def deep_dup; end
  def description; end
  def finish; end
  def get_trace_context; end
  def initialize(description: nil, op: nil, status: nil, trace_id: nil, parent_span_id: nil, sampled: nil, start_timestamp: nil, timestamp: nil); end
  def op; end
  def parent_span_id; end
  def sampled; end
  def set_data(key, value); end
  def set_description(description); end
  def set_http_status(status_code); end
  def set_op(op); end
  def set_status(status); end
  def set_tag(key, value); end
  def set_timestamp(timestamp); end
  def span_id; end
  def span_recorder; end
  def span_recorder=(arg0); end
  def start_child(**options); end
  def start_timestamp; end
  def status; end
  def tags; end
  def timestamp; end
  def to_hash; end
  def to_sentry_trace; end
  def trace_id; end
  def transaction; end
  def transaction=(arg0); end
  def with_child_span(**options, &block); end
end
class Sentry::Transaction < Sentry::Span
  def configuration; end
  def deep_dup; end
  def finish(hub: nil); end
  def generate_transaction_description; end
  def hub; end
  def init_span_recorder(limit = nil); end
  def initialize(hub:, name: nil, parent_sampled: nil, **options); end
  def logger; end
  def name; end
  def parent_sampled; end
  def self.from_sentry_trace(sentry_trace, hub: nil, **options); end
  def set_initial_sample_decision(sampling_context:); end
  def to_hash; end
  include Sentry::LoggingHelper
end
class Sentry::Transaction::SpanRecorder
  def add(span); end
  def initialize(max_length); end
  def max_length; end
  def spans; end
end
class Sentry::Breadcrumb
  def category; end
  def category=(arg0); end
  def data; end
  def data=(arg0); end
  def initialize(category: nil, data: nil, message: nil, timestamp: nil, level: nil, type: nil); end
  def level; end
  def level=(arg0); end
  def message; end
  def message=(msg); end
  def serialized_data; end
  def timestamp; end
  def timestamp=(arg0); end
  def to_hash; end
  def type; end
  def type=(arg0); end
end
class Sentry::BreadcrumbBuffer
  def buffer; end
  def buffer=(arg0); end
  def dup; end
  def each(&block); end
  def empty?; end
  def initialize(size = nil); end
  def members; end
  def peek; end
  def record(crumb); end
  def to_hash; end
  include Enumerable
end
class Sentry::Scope
  def add_breadcrumb(breadcrumb); end
  def add_event_processor(&block); end
  def apply_to_event(event, hint = nil); end
  def breadcrumbs; end
  def breadcrumbs=(arg0); end
  def clear; end
  def clear_breadcrumbs; end
  def contexts; end
  def contexts=(arg0); end
  def dup; end
  def event_processors; end
  def event_processors=(arg0); end
  def extra; end
  def extra=(arg0); end
  def fingerprint; end
  def fingerprint=(arg0); end
  def get_span; end
  def get_transaction; end
  def initialize(max_breadcrumbs: nil); end
  def level; end
  def level=(arg0); end
  def rack_env; end
  def rack_env=(arg0); end
  def self.os_context; end
  def self.runtime_context; end
  def set_context(key, value); end
  def set_contexts(contexts_hash); end
  def set_default_value; end
  def set_extra(key, value); end
  def set_extras(extras_hash); end
  def set_fingerprint(fingerprint); end
  def set_level(level); end
  def set_new_breadcrumb_buffer; end
  def set_rack_env(env); end
  def set_span(span); end
  def set_tag(key, value); end
  def set_tags(tags_hash); end
  def set_transaction_name(transaction_name); end
  def set_user(user_hash); end
  def span; end
  def span=(arg0); end
  def tags; end
  def tags=(arg0); end
  def transaction_name; end
  def transaction_names; end
  def transaction_names=(arg0); end
  def update_from_options(contexts: nil, extra: nil, tags: nil, user: nil, level: nil, fingerprint: nil); end
  def update_from_scope(scope); end
  def user; end
  def user=(arg0); end
  include Sentry::ArgumentCheckingHelper
end
class Sentry::Envelope
  def add_item(headers, payload); end
  def initialize(headers); end
  def to_s; end
end
class Sentry::DummyTransport < Sentry::Transport
  def events; end
  def events=(arg0); end
  def initialize(*arg0); end
  def send_event(event); end
end
class Sentry::HTTPTransport < Sentry::Transport
  def adapter; end
  def conn; end
  def faraday_opts; end
  def handle_rate_limited_response(headers); end
  def has_rate_limited_header?(headers); end
  def initialize(*args); end
  def parse_rate_limit_header(rate_limit_header); end
  def send_data(data); end
  def set_conn; end
  def should_compress?(data); end
  def ssl_configuration; end
end
class Sentry::Client
  def capture_event(event, scope, hint = nil); end
  def configuration; end
  def dispatch_async_event(async_block, event, hint); end
  def dispatch_background_event(event, hint); end
  def event_from_exception(exception, hint = nil); end
  def event_from_message(message, hint = nil, backtrace: nil); end
  def event_from_transaction(transaction); end
  def generate_sentry_trace(span); end
  def initialize(configuration); end
  def logger; end
  def send_event(event, hint = nil); end
  def transport; end
  include Sentry::LoggingHelper
end
class Sentry::Hub
  def add_breadcrumb(breadcrumb, hint: nil); end
  def bind_client(client); end
  def capture_event(event, **options, &block); end
  def capture_exception(exception, **options, &block); end
  def capture_message(message, **options, &block); end
  def clone; end
  def configuration; end
  def configure_scope(&block); end
  def current_client; end
  def current_layer; end
  def current_scope; end
  def initialize(client, scope); end
  def last_event_id; end
  def new_from_top; end
  def pop_scope; end
  def push_scope; end
  def start_transaction(transaction: nil, custom_sampling_context: nil, **options); end
  def with_background_worker_disabled(&block); end
  def with_scope(&block); end
  include Sentry::ArgumentCheckingHelper
end
class Sentry::Hub::Layer
  def client; end
  def client=(arg0); end
  def initialize(client, scope); end
  def scope; end
end
class Sentry::BackgroundWorker
  def _perform(&block); end
  def initialize(configuration); end
  def logger; end
  def max_queue; end
  def number_of_threads; end
  def perform(&block); end
  def shutdown; end
  def shutdown_timeout; end
  def shutdown_timeout=(arg0); end
  include Sentry::LoggingHelper
end
module Sentry::Rake
end
module Sentry::Rake::Application
  def display_error_message(ex); end
end
module Sentry::Rake::Task
  def execute(args = nil); end
end
module Rake
end
class Rake::Application
end
class Rake::Task
end
module Sentry::Rack
end
class Sentry::Rack::CaptureExceptions
  def call(env); end
  def capture_exception(exception); end
  def collect_exception(env); end
  def finish_transaction(transaction, status_code); end
  def initialize(app); end
  def start_transaction(env, scope); end
  def transaction_op; end
end
module Sentry::Net
end
module Sentry::Net::HTTP
  def do_finish; end
  def do_start; end
  def extract_request_info(req); end
  def finish_sentry_span; end
  def from_sentry_sdk?; end
  def record_sentry_breadcrumb(req, res); end
  def record_sentry_span(req, res); end
  def request(req, body = nil, &block); end
  def set_sentry_trace_header(req); end
  def start_sentry_span; end
end
module Sentry::Integrable
  def capture_exception(exception, **options, &block); end
  def capture_message(message, **options, &block); end
  def integration_name; end
  def register_integration(name:, version:); end
end
module Sentry::Breadcrumb::SentryLogger
  def add(*args, &block); end
  def add_breadcrumb(severity, message = nil, progname = nil); end
  def current_breadcrumbs; end
  def ignored_logger?(progname); end
end