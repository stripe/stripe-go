# typed: true
# frozen_string_literal: true

module Integrations
  module Utilities
    include Kernel

    def fail_if_dying_worker!
      raise Integrations::Errors::DyingWorkerError if Resque.heroku_will_terminate?
    end

    autoload :StripeUtil, 'integrations/utilities/stripe_util'
    autoload :Currency, 'integrations/utilities/currency'
  end
end
