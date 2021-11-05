# frozen_string_literal: true
# typed: ignore
require_relative '../test_helper'

class SessionsControllerTest < ApplicationIntegrationTest
  before do
    OmniAuth.config.test_mode = true
  end

  after do
    Rails.application.env_config['omniauth.auth'] = nil
    OmniAuth.config.test_mode = false
  end

  def mock_omniauth_salesforce
    OmniAuth.config.mock_auth[:salesforce] = OmniAuth::AuthHash.new(
    {"provider"=>"salesforce",
      "uid"=>
       "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
      "info"=>
       {"name"=>"Michael Bianco",
        "email"=>"pbo+billing@stripe.com",
        "nickname"=>"User16270609962827892049",
        "first_name"=>"Michael",
        "last_name"=>"Bianco",
        "location"=>"",
        "description"=>"",
        "image"=>
         "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/T?oauth_token=00D5e000003V0C7!AQkAQN6sJSBXBPCNnKvJ2AH5Uewt61p302fkm157epQHfqed4rSW0rBHcAu7I1mq4cgEg9ySiK.cHRYtqPjy74OwadiGlmFT",
        "phone"=>"",
        "urls"=>
         {"enterprise"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/c/{version}/00D5e000003V0C7",
          "metadata"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/m/{version}/00D5e000003V0C7",
          "partner"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/u/{version}/00D5e000003V0C7",
          "rest"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/",
          "sobjects"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/sobjects/",
          "search"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/search/",
          "query"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/query/",
          "recent"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/recent/",
          "tooling_soap"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/T/{version}/00D5e000003V0C7",
          "tooling_rest"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/tooling/",
          "profile"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/0055e000005HBroAAG",
          "feeds"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feeds",
          "groups"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/groups",
          "users"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/users",
          "feed_items"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-items",
          "feed_elements"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-elements",
          "custom_domain"=>"https://biancodevorg-dev-ed.my.salesforce.com"}},
      "credentials"=>
       {"token"=>
         "00D5e000003V0C7!AQkAQN6sJSBXBPCNnKvJ2AH5Uewt61p302fkm157epQHfqed4rSW0rBHcAu7I1mq4cgEg9ySiK.cHRYtqPjy74OwadiGlmFT",
        "expires"=>false,
        "instance_url"=>"https://biancodevorg-dev-ed.my.salesforce.com",
        "refresh_token"=>
         "5Aep861ZBValxWWBUfOotNJy7CNwailpM5.3b2NwRCWebnUFXQdHGsw6Wm7icEbnV7iDes7T3nt_vb4J_hFSMAa"},
      "extra"=>
       {"id"=>
         "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
        "asserted_user"=>true,
        "user_id"=>"0055e000005HBroAAG",
        "organization_id"=>"00D5e000003V0C7EAK",
        "username"=>"mbianco+biancodevorg@stripe.com",
        "nick_name"=>"User16270609962827892049",
        "display_name"=>"Michael Bianco",
        "email"=>"pbo+billing@stripe.com",
        "email_verified"=>true,
        "first_name"=>"Michael",
        "last_name"=>"Bianco",
        "timezone"=>"America/Los_Angeles",
        "photos"=>
         {"picture"=>
           "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/F",
          "thumbnail"=>
           "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/T"},
        "addr_street"=>nil,
        "addr_city"=>nil,
        "addr_state"=>nil,
        "addr_country"=>"US",
        "addr_zip"=>nil,
        "mobile_phone"=>nil,
        "mobile_phone_verified"=>false,
        "is_lightning_login_user"=>false,
        "status"=>{"created_date"=>nil, "body"=>nil},
        "urls"=>
         {"enterprise"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/c/{version}/00D5e000003V0C7",
          "metadata"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/m/{version}/00D5e000003V0C7",
          "partner"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/u/{version}/00D5e000003V0C7",
          "rest"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/",
          "sobjects"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/sobjects/",
          "search"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/search/",
          "query"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/query/",
          "recent"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/recent/",
          "tooling_soap"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/T/{version}/00D5e000003V0C7",
          "tooling_rest"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/tooling/",
          "profile"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/0055e000005HBroAAG",
          "feeds"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feeds",
          "groups"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/groups",
          "users"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/users",
          "feed_items"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-items",
          "feed_elements"=>
           "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-elements",
          "custom_domain"=>"https://biancodevorg-dev-ed.my.salesforce.com"},
        "active"=>true,
        "user_type"=>"STANDARD",
        "language"=>"en_US",
        "locale"=>"en_US",
        "utcOffset"=>-28800000,
        "last_modified_date"=>"2021-09-24T20:27:45Z",
        "is_app_installed"=>true,
        "instance_url"=>"https://biancodevorg-dev-ed.my.salesforce.com",
        "pod"=>"https://biancodevorg-dev-ed.my.salesforce.com",
        "signature"=>"/2CACreG+0ijLnSzncsjoXNkx3/afPGJz/Myvc7tJ2s=",
        "issued_at"=>"1636050290161"}}

    Rails.application.env_config['omniauth.auth'] = OmniAuth.config.mock_auth[:salesforce]
  end

  def mock_omniauth_stripe; end

  it 'should redirect to the salesforce login page by default' do
    mock_omniauth_salesforce

    Admin.expects(:from_omniauth).never

    get admin_google_oauth2_omniauth_callback_url

    assert_response :redirect
    assert_equal(response.headers['Location'], login_url(mode: nil))
    assert_match(/Contact support to setup an account/, flash[:notice])
    assert_equal(true, Admin.all.empty?)
  end

  it 'should redirect the user to the page they originally tried to access' do
    user = make_user(save: true)
    attempt = create_attempt(user: user)

    Admin.create(
      email: admin_email,
      stripe_user_id: user.stripe_user_id
    )

    target_location = translation_attempts_url(attempt)

    get target_location

    # redirects to login and stores accessed location in a cookie
    assert_equal(response.headers["Location"], login_url(mode: nil))

    mock_omniauth_google2
    get admin_google_oauth2_omniauth_callback_url

    assert_equal(target_location, response.headers["Location"])
  end

  it 'should login a standard admin associated with a stripe account' do
    user = make_user(save: true)

    Admin.create(
      email: admin_email,
      stripe_user_id: user.stripe_user_id
    )

    mock_omniauth_google2

    get admin_google_oauth2_omniauth_callback_url

    assert_response :redirect

    assert_nil(flash[:notice])
    assert_equal(response.headers["Location"], dashboard_url)

    admin = Admin.find(email: admin_email)
    assert_equal(admin_uid, admin.uid)
    assert_equal('view', admin.role)
    assert_equal('google_oauth2', admin.provider)
  end

  it 'should allow root login if user on admin is nil but role is root' do
    admin = Admin.create(
      email: admin_email,
      stripe_user_id: nil,
      role: 'root',
      provider: 'google_oauth2',
      uid: admin_uid
    )

    assert_equal(true, admin.root?)
    assert_nil(admin.user)

    mock_omniauth_google2

    get admin_google_oauth2_omniauth_callback_url

    assert_response :redirect

    assert_nil(flash[:notice])
    assert_equal(response.headers["Location"], users_url(mode: nil))
  end

  it 'should allow root login when a first login attempt from a stripe email is detected' do
    stripe_email = 'netsuite@stripe.com'

    admin = Admin.create(
      email: stripe_email,
      role: 'root',
      provider: 'google_oauth2',
    )

    assert_nil(admin.uid)

    mock_omniauth_google2(email: stripe_email)

    get admin_google_oauth2_omniauth_callback_url

    assert_response :redirect

    assert_nil(flash[:notice])
    assert_equal(users_url(mode: nil), response.headers["Location"])
  end

  it 'should not allow root login when a first login attempt is attempted from a non-stripe email' do
    admin = Admin.create(
      email: admin_email,
      role: 'root',
      provider: 'google_oauth2',
    )

    assert_nil(admin.uid)

    mock_omniauth_google2

    get admin_google_oauth2_omniauth_callback_url

    assert_response :redirect

    assert_match("Your account setup is not complete", flash[:notice])
    assert_equal(login_url(mode: nil), response.headers["Location"])
  end
end
