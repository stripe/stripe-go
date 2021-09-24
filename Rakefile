#!/usr/bin/env rake

# require_relative 'config/application'
# Rails.application.load_tasks

require 'resque/tasks'
require 'resque/scheduler/tasks'

namespace :resque do
  task :setup => :environment do

  end

  task :setup_schedule => :setup do
    # without `first_in` the job will only run after the first interval is complete
    # for these jobs, that's waiting too long, so we wait a short period of time instead
    Resque.schedule = {
      InitiatePollsJobs: {
        every: [
          '90s',
          { first_in: 10.seconds }
        ]
      },
      InitiateBatchJobs: {
        cron: '0 5 * * * America/Los_Angeles'
      },
    }
  end
  task :scheduler => :setup_schedule
end