# typed: true
# frozen_string_literal: true

Sentry.init do |config|
  config.logger = Integrations::Log.log
  config.dsn = ENV.fetch('SENTRY_DSN')
  config.traces_sample_rate = 0.5
  config.breadcrumbs_logger = [:sentry_logger, :http_logger]
  config.capture_exception_frame_locals = true
  config.enabled_environments = %w{production staging}
  config.excluded_exceptions = []

  # config.background_worker_threads = 0

  # `DYNO` is formatted as `worker.12`, `scheduler.1`, etc
  config.server_name = ENV.fetch('DYNO')[/[^.]+/, 0] if ENV['DYNO']
end
