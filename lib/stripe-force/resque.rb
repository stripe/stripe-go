# frozen_string_literal: true
# typed: ignore
require 'resque/failure/multiple'
require 'resque/failure/redis'

require './lib/resque_backtrace'

Resque.logger.level = Logger::DEBUG

Resque.redis = Redis.new(
  url: ENV.fetch("REDIS_URL"),

  thread_safe: true,

  ssl_params: {verify_mode: OpenSSL::SSL::VERIFY_NONE},

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

# `redis.exists?` is not available < redis 4.2.0
# however, there's a bug in redis-namespace that reports that exists? is available when it is not:
#   https://github.com/resque/redis-namespace/blob/6f1bb31bbddd2efc9e16c53759d986b29206a12e/lib/redis/namespace.rb#L71-L72
# the implementation of `redis_key_exists?` is overwritten here to avoid using `exists?`
#   https://github.com/lantins/resque-retry/blob/f53624e3961a0f33493a0a9f37b78f79ff573438/lib/resque/failure/multiple_with_retry_suppression.rb#L151
# the root fix here is update to redis, or fix the bug in redis-namespace, but we can't do that without removing sidekiq
Resque::Failure::MultipleWithRetrySuppression.class_eval do
  def redis_key_exists?(key)
    ![false, 0].include?(Resque.redis.exists(key) || false)
  end
end
