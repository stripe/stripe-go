# frozen_string_literal: true
# typed: ignore
require 'resque/failure/multiple'
require 'resque/failure/redis'

require './lib/resque_backtrace'

redis_uri = URI.parse(ENV.fetch("REDIS_URL"))

Resque.logger.level = Logger::DEBUG
Resque.redis = Redis.new(
  host: redis_uri.host,
  port: redis_uri.port,
  password: redis_uri.password,
  thread_safe: true,

  # stronger reconnection retry to avoid throwing errors when an intermittent connection failure occurs
  # https://github.com/redis/redis-rb#reconnections
  reconnect_attempts: 10,
  reconnect_delay: 1.5,
  reconnect_delay_max: 10.0,
)

Resque.redis.namespace = "resque:salesforce"

# sharing a database connection after the process forks can cause strange behavior:
# http://sequel.jeremyevans.net/rdoc/files/doc/fork_safety_rdoc.html
# to avoid this, we disconnect from the DB. Sequel will automatically reconnect for us.

Resque.before_fork do
  DB.disconnect
end

# NOTE `MultipleWithRetrySuppression` will NOT process `classes`
#       we need to route the error through the backtrace output and sentry THEN the retry handled

# https://github.com/lantins/resque-retry#failure-backend
# https://github.com/resque/resque/blob/master/lib/resque/failure/multiple.rb#L18
Resque::Failure::MultipleWithRetrySuppression.classes = [Resque::Failure::Redis]

Resque::Failure::Multiple.classes = [
  Resque::Failure::Backtrace,
  Resque::Failure::MultipleWithRetrySuppression,
]
Resque::Failure.backend = Resque::Failure::Multiple

# remove increment_stat implementation to avoid redis `incrby`
# this accounts for ~1% of total redis computation time
# original implementation: https://github.com/resque/resque/blob/adb633a0f6b98b1eb5a5a85bb36ebac9309978fd/lib/resque/data_store.rb#L319
# related PR: https://github.com/stripe/stripe-netsuite/pull/2185
Resque::DataStore::StatsAccess.class_eval do
  def increment_stat(stat, by=1); end
end
