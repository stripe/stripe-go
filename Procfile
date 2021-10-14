web: bundle exec puma
worker: env QUEUES=high,medium,low TERM_CHILD=1 INTERVAL=0.0001 RESQUE_PRE_SHUTDOWN_TIMEOUT=20 RESQUE_TERM_TIMEOUT=8 bundle exec rake resque:work
scheduler: env RESQUE_SCHEDULER_INTERVAL=1 bundle exec rake resque:scheduler
