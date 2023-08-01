# frozen_string_literal: true
# typed: true

require_relative './vcr_config'
require_relative '../test_helper'

class Critic::VCRTest < Critic::FunctionalTest
  include VCRConfig

  def setup
    super

    setup_vcr(name)
  end

  def teardown
    super

    teardown_vcr
  end

  def always_use_vcr
    # If we're manually enabling/disabling VCR: Clear current cassette to set up new one
    # Ignore error saying recording was not used
    teardown_vcr(skip_unused_interactions: true)
    setup_vcr(name, override_vcr_usage: true)
  end

  def never_use_vcr
    teardown_vcr(skip_unused_interactions: true)
    setup_vcr(name, override_vcr_usage: false)
  end

end
