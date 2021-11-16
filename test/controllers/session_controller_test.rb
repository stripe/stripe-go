# frozen_string_literal: true
# typed: true
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
      {"provider" => "salesforce",
       "uid" =>
         "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
       "info" =>
         {"name" => "Michael Bianco",
          "email" => "pbo+billing@stripe.com",
          "nickname" => "User16270609962827892049",
          "first_name" => "Michael",
          "last_name" => "Bianco",
          "location" => "",
          "description" => "",
          "image" =>
           "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/T?oauth_token=00D5e000003V0C7!AQkAQN6sJSBXBPCNnKvJ2AH5Uewt61p302fkm157epQHfqed4rSW0rBHcAu7I1mq4cgEg9ySiK.cHRYtqPjy74OwadiGlmFT",
          "phone" => "",
          "urls" =>
           {"enterprise" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/c/{version}/00D5e000003V0C7",
            "metadata" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/m/{version}/00D5e000003V0C7",
            "partner" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/u/{version}/00D5e000003V0C7",
            "rest" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/",
            "sobjects" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/sobjects/",
            "search" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/search/",
            "query" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/query/",
            "recent" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/recent/",
            "tooling_soap" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/T/{version}/00D5e000003V0C7",
            "tooling_rest" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/tooling/",
            "profile" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/0055e000005HBroAAG",
            "feeds" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feeds",
            "groups" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/groups",
            "users" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/users",
            "feed_items" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-items",
            "feed_elements" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-elements",
            "custom_domain" => "https://biancodevorg-dev-ed.my.salesforce.com",},},
       "credentials" =>
         {"token" =>
           "00D5e000003V0C7!AQkAQN6sJSBXBPCNnKvJ2AH5Uewt61p302fkm157epQHfqed4rSW0rBHcAu7I1mq4cgEg9ySiK.cHRYtqPjy74OwadiGlmFT",
          "expires" => false,
          "instance_url" => "https://biancodevorg-dev-ed.my.salesforce.com",
          "refresh_token" =>
           "5Aep861ZBValxWWBUfOotNJy7CNwailpM5.3b2NwRCWebnUFXQdHGsw6Wm7icEbnV7iDes7T3nt_vb4J_hFSMAa",},
       "extra" =>
         {"id" =>
           "https://login.salesforce.com/id/00D5e000003V0C7EAK/0055e000005HBroAAG",
          "asserted_user" => true,
          "user_id" => "0055e000005HBroAAG",
          "organization_id" => "00D5e000003V0C7EAK",
          "username" => "mbianco+biancodevorg@stripe.com",
          "nick_name" => "User16270609962827892049",
          "display_name" => "Michael Bianco",
          "email" => "pbo+billing@stripe.com",
          "email_verified" => true,
          "first_name" => "Michael",
          "last_name" => "Bianco",
          "timezone" => "America/Los_Angeles",
          "photos" =>
           {"picture" =>
             "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/F",
            "thumbnail" =>
             "https://biancodevorg-dev-ed--c.documentforce.com/profilephoto/005/T",},
          "addr_street" => nil,
          "addr_city" => nil,
          "addr_state" => nil,
          "addr_country" => "US",
          "addr_zip" => nil,
          "mobile_phone" => nil,
          "mobile_phone_verified" => false,
          "is_lightning_login_user" => false,
          "status" => {"created_date" => nil, "body" => nil},
          "urls" =>
           {"enterprise" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/c/{version}/00D5e000003V0C7",
            "metadata" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/m/{version}/00D5e000003V0C7",
            "partner" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/u/{version}/00D5e000003V0C7",
            "rest" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/",
            "sobjects" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/sobjects/",
            "search" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/search/",
            "query" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/query/",
            "recent" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/recent/",
            "tooling_soap" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/Soap/T/{version}/00D5e000003V0C7",
            "tooling_rest" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/tooling/",
            "profile" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/0055e000005HBroAAG",
            "feeds" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feeds",
            "groups" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/groups",
            "users" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/users",
            "feed_items" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-items",
            "feed_elements" =>
             "https://biancodevorg-dev-ed.my.salesforce.com/services/data/v{version}/chatter/feed-elements",
            "custom_domain" => "https://biancodevorg-dev-ed.my.salesforce.com",},
          "active" => true,
          "user_type" => "STANDARD",
          "language" => "en_US",
          "locale" => "en_US",
          "utcOffset" => -28800000,
          "last_modified_date" => "2021-09-24T20:27:45Z",
          "is_app_installed" => true,
          "instance_url" => "https://biancodevorg-dev-ed.my.salesforce.com",
          "pod" => "https://biancodevorg-dev-ed.my.salesforce.com",
          "signature" => "/2CACreG+0ijLnSzncsjoXNkx3/afPGJz/Myvc7tJ2s=",
          "issued_at" => "1636050290161",},})

    Rails.application.env_config['omniauth.auth'] = OmniAuth.config.mock_auth[:salesforce]
  end

  def mock_omniauth_stripe; end

  def create_post_install_user
    # no other fields are configured after post-install
    StripeForce::User.create(
      salesforce_account_id: sf_account_id,
      salesforce_organization_key: SecureRandom.alphanumeric(16)
    )
  end

  # the account
  it 'creates a new user when it is authorized with salesforce' do
    user = create_post_install_user

    mock_omniauth_salesforce

    get auth_salesforce_callback_path

    assert_response :redirect

    assert_match(%r{/auth/stripe$}, response.headers['Location'])

    user = StripeForce::User[user.id]

    assert_equal(1, StripeForce::User.count)

    refute_nil(user.salesforce_account_id)
    refute_nil(user.salesforce_refresh_token)
    refute_nil(user.salesforce_instance_url)
    refute_nil(user.salesforce_token)

    refute_nil(user.name)
    refute_nil(user.email)
  end

  it 'updates an existing user with stripe oauth information if the account is already authorized with salesforce' do

  end
end
