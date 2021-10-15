# TODO add typing

require 'resque/scheduler/tasks'

namespace :resque do
  task setup_schedule: :setup do
    # without `first_in` the job will only run after the first interval is complete
    # for these jobs, that's waiting too long, so we wait a short period of time instead
    Resque.schedule = {
      'StripeForce::InitiatePollsJobs': {
        every: [
          '90s',
          {first_in: 10.seconds},
        ],
      },
    }
  end

  task scheduler: :setup_schedule
end
