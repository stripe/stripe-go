# typed: true
# frozen_string_literal: true

module Integrations
  module Utilities
    include Kernel

    def fail_if_dying_worker!
      raise Integrations::Errors::DyingWorkerError if Resque.heroku_will_terminate?
    end

    def feature?(flag)
      @user.feature_enabled?(flag)
    end

    autoload :Stripe, 'integrations/utilities/stripe'
    autoload :Currency, 'integrations/utilities/currency'
  end
end
