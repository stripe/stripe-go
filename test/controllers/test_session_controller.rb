# frozen_string_literal: true
# typed: true
require_relative '../test_helper'
require_relative './helper'

class SessionsControllerTest < ApplicationIntegrationTest
  before do
    OmniAuth.config.test_mode = true
  end

  after do
    Rails.application.env_config['omniauth.auth'] = nil
    OmniAuth.config.test_mode = false
  end

  def mock_omniauth_salesforce(sf_instance_account_id)
    OmniAuth.config.mock_auth.delete :default

    OmniAuth.config.mock_auth[:stripetestmode] = OmniAuth::AuthHash.new({
      "uid" => ENV.fetch('STRIPE_ACCOUNT_ID'),
    })

    OmniAuth.config.mock_auth[:stripetestmodev2] = OmniAuth::AuthHash.new({
      "uid" => ENV.fetch('STRIPE_ACCOUNT_ID'),
    })

    OmniAuth.config.mock_auth[:stripelivemodev2] = OmniAuth::AuthHash.new({
      "uid" => ENV.fetch('STRIPE_ACCOUNT_ID'),
    })

    OmniAuth.config.mock_auth[:salesforce] = get_salesforce_auth_hash(sf_instance_account_id)
    OmniAuth.config.mock_auth[:salesforcev2] = get_salesforce_auth_hash(sf_instance_account_id)

    Rails.application.env_config['omniauth.auth'] = OmniAuth.config.mock_auth[:salesforce]
  end

  def mock_omniauth_stripe; end

  def create_post_install_user
    # no other fields are configured after post-install
    StripeForce::User.create(
      salesforce_account_id: sf_instance_account_id,
      salesforce_organization_key: SecureRandom.alphanumeric(16)
    )
  end

  def assert_post_redirect(path)
    assert_includes(response.body, path)
    assert_includes(response.body, "document.getElementById('js-submission').submit()")
  end

  def refute_post_redirect(path)
    refute_includes(response.body, path)
    refute_includes(response.body, "document.getElementById('js-submission').submit()")
  end

  describe 'initial login redirect' do
    it 'submits a POST request and persists namespace if a GET method is used to initiate oauth' do
      get omniauth_path(:salesforce, salesforceNamespace: "c")

      assert_response :success
      assert_post_redirect(omniauth_path(:salesforce))
      assert_equal("c", session[:salesforce_namespace])
    end

    it 'submits a POST request with a default namespace if one is not defined' do
      get omniauth_path(:salesforce)

      assert_response :success
      assert_post_redirect(omniauth_path(:salesforce))
      assert_equal("stripeConnector", session[:salesforce_namespace])
    end

    it 'fails if an invalid namespace is provided' do
      assert_raises RuntimeError, 'unexpected namespace: invalid' do
        get omniauth_path(:salesforce, salesforceNamespace: "invalid")

        assert_response :error
      end
    end
  end

  it 'creates a new user when it is authorized with salesforce' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_salesforce_callback_path

    assert_response :success
    assert_post_redirect(omniauth_path(:stripetestmode))

    user = T.must(StripeForce::User[user.id])

    assert_equal(1, StripeForce::User.count)

    refute_nil(user.salesforce_account_id)
    refute_nil(user.salesforce_refresh_token)
    refute_nil(user.salesforce_instance_url)
    refute_nil(user.salesforce_token)

    refute_nil(user.name)
    refute_nil(user.email)
  end

  it 'creates a new user when it is authorized with the salesforce v2 endpoint' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_v2_salesforce_callback_path

    assert_response :success
    refute_post_redirect(omniauth_path(:stripetestmode))

    user = T.must(StripeForce::User[user.id])

    assert_equal(1, StripeForce::User.count)

    refute_nil(user.salesforce_account_id)
    refute_nil(user.salesforce_refresh_token)
    refute_nil(user.salesforce_instance_url)
    refute_nil(user.salesforce_token)

    refute_nil(user.name)
    refute_nil(user.email)
  end

  it 'caches user as authenticated with salesforce post auth' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    get auth_salesforce_callback_path

    assert(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    assert_response :success
    assert_post_redirect(omniauth_path(:stripetestmode))

    user = T.must(StripeForce::User[user.id])

    assert_equal(1, StripeForce::User.count)

    refute_nil(user.salesforce_account_id)
    refute_nil(user.salesforce_refresh_token)
    refute_nil(user.salesforce_instance_url)
    refute_nil(user.salesforce_token)

    refute_nil(user.name)
    refute_nil(user.email)
  end

  it 'caches user as authenticated with salesforce v2 post auth' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    get auth_v2_salesforce_callback_path

    assert(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    assert_response :success
    refute_post_redirect(omniauth_path(:stripetestmode))

    user = T.must(StripeForce::User[user.id])

    assert_equal(1, StripeForce::User.count)

    refute_nil(user.salesforce_account_id)
    refute_nil(user.salesforce_refresh_token)
    refute_nil(user.salesforce_instance_url)
    refute_nil(user.salesforce_token)

    refute_nil(user.name)
    refute_nil(user.email)
  end

  it 'caches user as authenticated with stripe post auth' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path
    assert_response :success

    get auth_stripetestmode_callback_path
    assert_response :success

    assert(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

  end

  it 'caches user as authenticated with stripe v2 testmode post auth' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_v2_salesforce_callback_path
    assert_response :success

    get auth_v2_stripetestmode_callback_path
    assert_response :success

    assert(user.get_cached_stripe_v2_connection_status(false))
  end

  it 'caches user as authenticated with stripe v2 livemode post auth' do
    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_v2_salesforce_callback_path
    assert_response :success

    get auth_v2_stripelivemode_callback_path
    assert_response :success

    assert(user.get_cached_stripe_v2_connection_status(true))
  end

  it 'updates the stripe account id after the stripe account is authenticated' do
    # in order to set the session, which cannot be set via the test suite
    get omniauth_path(:salesforce, salesforceNamespace: "c")

    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_salesforce_callback_path
    assert_response :success

    get auth_stripetestmode_callback_path
    assert_response :success

    assert_equal(1, StripeForce::User.count)

    user = T.must(StripeForce::User[user.id])
    assert_equal(OmniAuth.config.mock_auth[:stripetestmode]["uid"], user.stripe_account_id)

    # ensures the correct namespace is used for the callback
    assert_includes(response.body, "--c")
  end

  it 'updates stripe account id with stripe v2 testmode post auth' do
    get omniauth_path(:salesforcev2, salesforceNamespace: "c")

    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_v2_salesforce_callback_path
    assert_response :success

    get auth_v2_stripetestmode_callback_path
    assert_response :success

    assert_equal(1, StripeForce::User.count)
    user = T.must(StripeForce::User[user.id])

    assert_equal(OmniAuth.config.mock_auth[:stripetestmodev2]["uid"], user.stripe_account_id)
  end

  it 'updates stripe account id with stripe v2 livemode post auth' do
    get omniauth_path(:salesforcev2, salesforceNamespace: "c")

    user = create_post_install_user

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_v2_salesforce_callback_path
    assert_response :success

    get auth_v2_stripetestmode_callback_path
    assert_response :success

    assert_equal(1, StripeForce::User.count)
    user = T.must(StripeForce::User[user.id])

    assert_equal(OmniAuth.config.mock_auth[:stripelivemodev2]["uid"], user.stripe_account_id)
  end

  # it 'updates an existing user with stripe oauth information if the account is already authorized with salesforce' do
  #   skip("not yet implemented")
  # end
end
