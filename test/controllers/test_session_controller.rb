# frozen_string_literal: true
# typed: true
require_relative '../test_helper'
require_relative './helper'

class SessionsControllerTest < ApplicationIntegrationTest
  REFERER_INSTANCE_URL = "https://momentum-customization-3160-dev-ed.my.salesforce.com/"

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

    OmniAuth.config.mock_auth[:stripelivemode] = OmniAuth::AuthHash.new({
      "uid" => ENV.fetch('STRIPE_ACCOUNT_ID'),
    })

    OmniAuth.config.mock_auth[:salesforce] = get_salesforce_auth_hash(sf_instance_account_id)

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
      get omniauth_path(:salesforce, salesforceNamespace: "c"), headers: {'HTTP_REFERER' => REFERER_INSTANCE_URL}

      assert_response :success
      assert_post_redirect(omniauth_path(:salesforce))

      state = StateEncryptionAlgo::StripeOAuthState.from_encrypted_state(session[:state])
      assert_equal("c", T.must(state).salesforce_namespace)
    end

    it 'submits a POST request with a default namespace if one is not defined' do
      get omniauth_path(:salesforce), headers: {'HTTP_REFERER' => REFERER_INSTANCE_URL}

      assert_response :success
      assert_post_redirect(omniauth_path(:salesforce))

      state = StateEncryptionAlgo::StripeOAuthState.from_encrypted_state(session[:state])
      assert_equal("stripeConnector", T.must(state).salesforce_namespace)
    end

    it 'fails if an invalid namespace is provided' do
      assert_raises Integrations::Errors::ImpossibleInternalError, 'Unexpected namespace: invalid' do
        get omniauth_path(:salesforce, salesforceNamespace: "invalid")

        assert_response :error
      end
    end
  end

  it 'creates a new user when it is authorized with salesforce' do
    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_salesforce_callback_path, params: {state: state}

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
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_salesforce_callback_path, params: {state: state}

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
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    get auth_salesforce_callback_path, params: {state: state}

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
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

    get auth_salesforce_callback_path, params: {state: state}

    assert_response :success
    refute_post_redirect(omniauth_path(:stripetestmode))

    user = T.must(StripeForce::User[user.id])
    assert(user.get_cached_connection_status(StripeForce::Constants::Platforms::SALESFORCE))

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
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripetestmode_callback_path, params: {state: state}
    assert_response :success
    user = T.must(StripeForce::User[user.id])
    assert(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))
  end

  it 'caches user as authenticated with stripe v2 testmode post auth' do
    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripetestmode_callback_path, params: {state: state}
    assert_response :success

    user = T.must(StripeForce::User[user.id])
    assert(user.get_cached_stripe_v2_connection_status(false))
  end

  it 'caches user as authenticated with stripe v2 livemode post auth' do
    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id})
    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripelivemode_callback_path, params: {state: state}
    assert_response :success

    user = T.must(StripeForce::User[user.id])
    assert(user.get_cached_stripe_v2_connection_status(true))
  end

  it 'updates the stripe account id after the stripe account is authenticated' do
    # in order to set the session, which cannot be set via the test suite
    get omniauth_path(:salesforce, salesforceNamespace: "c"), headers: {'HTTP_REFERER' => REFERER_INSTANCE_URL}

    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id, salesforce_namespace: "c"})

    mock_omniauth_salesforce(sf_instance_account_id)

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripetestmode_callback_path, params: {state: state}
    assert_response :success

    assert_equal(1, StripeForce::User.count)

    user = T.must(StripeForce::User[user.id])
    assert_equal(OmniAuth.config.mock_auth[:stripetestmode]["uid"], user.stripe_account_id)

    # ensures the correct namespace is used for the callback
    assert_includes(response.body, "--c")
  end

  it 'updates stripe account id with stripe v2 testmode post auth' do
    get omniauth_v2_path(:salesforce, salesforceNamespace: "c"), headers: {'HTTP_REFERER' => REFERER_INSTANCE_URL}

    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id, salesforce_namespace: "c"})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripetestmode_callback_path, params: {state: state}
    assert_response :success

    assert_equal(1, StripeForce::User.count)
    user = T.must(StripeForce::User[user.id])

    assert_equal(OmniAuth.config.mock_auth[:stripetestmode]["uid"], user.stripe_account_id)
  end

  it 'updates stripe account id with stripe v2 livemode post auth' do
    get omniauth_v2_path(:salesforce, salesforceNamespace: "c"), headers: {'HTTP_REFERER' => REFERER_INSTANCE_URL}

    user = create_post_install_user
    state = StateEncryptionAlgo::StripeOAuthState.create
    StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v2', user_id: user.id, salesforce_namespace: "c"})

    mock_omniauth_salesforce(sf_instance_account_id)

    assert_nil(user.get_cached_connection_status(StripeForce::Constants::Platforms::STRIPE))

    get auth_salesforce_callback_path, params: {state: state}
    assert_response :success

    get auth_stripelivemode_callback_path, params: {state: state}
    assert_response :success

    assert_equal(1, StripeForce::User.count)
    user = T.must(StripeForce::User[user.id])

    assert_equal(OmniAuth.config.mock_auth[:stripelivemode]["uid"], user.stripe_account_id)
  end

  describe 'change default account configuration' do
    it 'changes the default account configuration' do
      user = create_post_install_user
      user2 = create_post_install_user
      state = StateEncryptionAlgo::StripeOAuthState.create
      StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id, salesforce_account_id: user.salesforce_account_id})

      user.stripe_account_id = 'test_account_1'
      user.livemode = true
      user.is_default_account_config = true
      user2.stripe_account_id = 'test_account_2'
      user2.livemode = true
      user2.is_default_account_config = false
      user.save
      user2.save

      ActionController::Base.any_instance.stubs(:verify_authenticity_token).returns(true)

      post set_default_config_path('test_account_2', 'true'), params: {state: state, livemode: 'true', stripe_account_id: 'test_account_2'}

      user.refresh
      user2.refresh

      assert_response :success
      assert(user2.is_default_account_config)
      refute(user.is_default_account_config)
    end
  end

  describe 'get all configurations' do
    it 'gets all the account configurations' do
      user = create_post_install_user
      user2 = create_post_install_user
      state = StateEncryptionAlgo::StripeOAuthState.create
      StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id, salesforce_account_id: user.salesforce_account_id})

      get accounts_path, params: {state: state}

      assert_response :success
      assert_equal(2, JSON.parse(response.body).length)
    end
  end

  describe 'delete account configuration' do
    it 'deletes the account configuration' do
      user = create_post_install_user
      state = StateEncryptionAlgo::StripeOAuthState.create
      StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id, salesforce_account_id: user.salesforce_account_id})

      user.stripe_account_id = 'test_account_1'
      user.livemode = true
      user.save

      ActionController::Base.any_instance.stubs(:verify_authenticity_token).returns(true)

      delete delete_account_config_path('test_account_1', 'true'), params: {state: state, livemode: 'true', stripe_account_id: 'test_account_id'}

      assert_response :success
      assert_nil StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, stripe_account_id: 'test_account_id', livemode: true)
    end

    it 'does not allow deleting default user if it is not the only account in the org' do
      user = create_post_install_user
      state = StateEncryptionAlgo::StripeOAuthState.create
      StateEncryptionAlgo.apply_defaults(state, {oauth_version: 'v1', user_id: user.id, salesforce_account_id: user.salesforce_account_id})

      user.stripe_account_id = 'test_account_1'
      user.livemode = true
      user.is_default_account_config = true
      user.save

      # Create another user in the same salesforce account.
      other_user = StripeForce::User.create(salesforce_account_id: user.salesforce_account_id, stripe_account_id: 'other_account', livemode: true, is_default_account_config: false)

      # Attempt to delete the default user
      ActionController::Base.any_instance.stubs(:verify_authenticity_token).returns(true)

      delete delete_account_config_path('test_account_1', 'true'), params: {state: state, livemode: 'true', stripe_account_id: 'test_account_1'}

      assert_response 400 # Expecting 400 Bad Request response, adjust as necessary
      assert_equal 'The current (default) User can only be deleted if it is the only account in the org.', JSON.parse(response.body)["message"]

      # Check to confirm that the default user is still present in the database
      refute_nil StripeForce::User.find(salesforce_account_id: state.salesforce_account_id, stripe_account_id: 'test_account_1', livemode: true)
    end
  end

  # it 'updates an existing user with stripe oauth information if the account is already authorized with salesforce' do
  #   skip("not yet implemented")
  # end
end
