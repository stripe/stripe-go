# frozen_string_literal: true
# typed: true


module VCRConfig

  # Each test worker should only configure VCR once
  @@configured = false

  def setup_vcr(test_name, override_vcr_usage: nil)
    if !vcr_enabled?(override_vcr_usage)
      VCR.turn_off!
      return
    end

    VCR.turn_on!

    # need this line to make sure that the Stripe client is properly being mocked on all requests
    # https://github.com/stripe/stripe-ruby/blob/af6fb06cd69e066adbdf7c2d9e67e80c3d870529/lib/stripe/stripe_client.rb#L381-L383
    Thread.current[:stripe_client__internal_use_only] = nil

    if !@@configured
      configure
    end

    # example: test with name 'test_0002_searches for a matching item' will turn into 'searches for a matching item'
    # we want to eliminate the number at the beginning so that if we add a test in the middle of a test file we
    # don't have to regenerate all tests after the added test
    matched_name = test_name.match(/test_[0-9]*_(.*)/)[1]
    VCR.insert_cassette matched_name
  end

  def teardown_vcr(skip_unused_interactions: false)
    # Ejects cassette if any exists
    VCR.eject_cassette(skip_no_unused_interactions_assertion: skip_unused_interactions)
    VCR.turn_off!
    WebMock.allow_net_connect!
  end

  def configure
    VCR.configure do |c|
      c.allow_http_connections_when_no_cassette = false
      c.cassette_library_dir = 'test/vcr_cassettes'
      c.default_cassette_options = {
        allow_unused_http_interactions: false,
        match_requests_on: [:method, :salesforce_uri, :salesforce_body],
      }
      c.filter_sensitive_data('<STRIPE_API_KEY>') do |interaction|
        filter_stripe_api_key(interaction)
      end
      c.filter_sensitive_data('<STRIPE_MERCHANT_ID>') do |interaction|
        filter_stripe_merchant_id(interaction)
      end
      c.filter_sensitive_data('<SALESFORCE_ACCESS_TOKEN>') do |interaction|
        filter_salesforce_auth_header(interaction)
      end
      c.filter_sensitive_data('<HEROKU_API_KEY>') { ENV['HEROKU_API_KEY'] }
      c.hook_into :webmock
      c.ignore_localhost = true
      c.register_request_matcher :salesforce_uri do |request1, request2|
        register_uri_matcher(request1, request2)
      end
      c.register_request_matcher :salesforce_body do |request1, request2|
        register_body_matcher(request1, request2)
      end
    end

    @@configured = true
  end

  def set_cassette_dir(filename)
    # example: test file test/functional/core/test_refunds.rb will turn into test/vcr_cassettes/functional/core/test_refunds/
    # each test will have its own VCR recording file within this folder
    file_regex = filename.match(%r{.*test/(.*)/([^/]*)\.rb})
    test_folder_path = file_regex[1]
    test_file_base_name = file_regex[2]

    cassette_name = "test/vcr_cassettes/#{test_folder_path}/#{test_file_base_name}"
    VCR.configuration.cassette_library_dir = cassette_name
  end

  def vcr_enabled?(override_vcr_usage=nil)
    # VCR may be manually enabled or disabled for a test file
    return true if override_vcr_usage == true
    return false if override_vcr_usage == false

    # we want to run live tests in CI's nightly build to make sure that tests haven't gone stale
    is_nightly_build = (ENV['NIGHTLY'].nil? || ENV['NIGHTLY'] == "true")
    !(!!ENV['CI'] && is_nightly_build)
  end

  def register_uri_matcher(request1, request2)
    uri1 = request1.uri
    uri2 = request2.uri

    return true if uri1 == uri2

    if uri1.include?('stripe.com') && uri2.include?('stripe.com')
      # if we are making requests to the Stripe API, sometimes the request will have an ID or params that we want to ignore
      #   GET https://api.stripe.com/v1/skus/new_sku_1661993849
      #   GET https://api.stripe.com/v1/payouts?arrival_date%5Bgte%5D=1665562166&limit=100

      request_regex = %r{https://api.stripe.com/v./([^/]*)[/?].*}
      request1_match = uri1.match(request_regex)
      request2_match = uri2.match(request_regex)

      # we want to see whether both requests matched the Stripe request formatting and whether they were both requesting the same API (such as SKU from the above request)
      !request1_match.nil? && !request2_match.nil? && request1_match[1] == request2_match[1]
    elsif uri1.include?('salesforce.com') && uri2.include?('salesforce.com')
      # if this request is to Salesforce everything except for the local subdomain should match
      salesforce_uri_regex = %r{https:\/\/[\w\-]*\.my.salesforce\.com\/.*}
      uri1.sub(salesforce_uri_regex, '') ==
        uri2.sub(salesforce_uri_regex, '')
    else
      # if this request isn't to Stripe then the URIs should exactly match
      false
    end
  end

  def register_body_matcher(request1, request2)
    body1 = request1.body
    body2 = request2.body
    uri = request1.uri # we only need to keep track of 1 uri since at this point we know the URIs are the same

    return true if body1 == body2 # if the bodies are already the same then we have a match
    return false if body1.empty? || body2.empty? # if one of the bodies is empty but the bodies are not equal to each other then we do not have a match

    if uri.include?('stripe.com')
      # if we are making requests to the Stripe API, as long as the method and URI are the same we do not care about the body for now
      # today the bodies differ for things like the email address in tests so we want to ignore those discrepencies
      # we may need to make this check more specific in the future
      return true
    end

    # if this request isn't to Stripe then the bodies should exactly match
    body1.deep_eql?(body2)
  end

  def filter_stripe_api_key(interaction)
    auth_header = interaction.request.headers["Authorization"]

    return nil if auth_header.nil?

    # when playing back a recording, VCR will try to replace the placeholder with the real secret but that is not necessary for our tests
    return nil if auth_header.first == "Bearer <STRIPE_API_KEY>"

    matches = auth_header.first.match(/Bearer (sk_.*)/)

    return nil if matches.nil?

    # index 0 is the whole matching string and 1 is the first capture group (the group starting with sk_)
    matches[1]
  end

  def filter_stripe_merchant_id(interaction)
    account_header = interaction.request.headers["Stripe-Account"]

    return nil if account_header.nil?

    # when playing back a recording, VCR will try to replace the placeholder with the real secret but that is not necessary for our tests
    return nil if account_header.first == "<STRIPE_MERCHANT_ID>"

    account_header.first.match(/(.*)/)
  end

  def filter_salesforce_auth_header(interaction)
    auth_header = interaction.request.headers["Authorization"]

    return nil if auth_header.nil?

    # when playing back a recording, VCR will try to replace the placeholder with the real secret but that is not necessary for our tests
    return nil if auth_header.first == "OAuth <SALESFORCE_ACCESS_TOKEN>"

    matches = auth_header.first.match(/OAuth (.*)/)

    return nil if matches.nil?

    # index 0 is the whole matching string and 1 is the first capture group (the group containing the token)
    matches[1]
  end
end
