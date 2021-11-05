# frozen_string_literal: true
# typed: true

class ApplicationIntegrationTest < ActionDispatch::IntegrationTest
  include CommonHelpers

  def setup
    common_setup
  end

  def teardown
    common_teardown
  end
end
