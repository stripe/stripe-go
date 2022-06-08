# frozen_string_literal: true
# typed: true

require_relative './common_helpers'

class ApplicationIntegrationTest < ActionDispatch::IntegrationTest
  include CommonHelpers

  def parsed_json
    JSON.parse(response.body)
  end

  def use_https_protocol
    https!
    app.default_url_options[:protocol] = :https
  end

  def setup
    common_setup
    use_https_protocol
  end

  def teardown
    common_teardown
  end

  def authentication_headers
    {
      SALESFORCE_ACCOUNT_ID_HEADER => @user.salesforce_account_id,
      SALESFORCE_PACKAGE_NAMESPACE_HEADER => "",
      SALESFORCE_KEY_HEADER => @user.salesforce_organization_key,
    }
  end

end
