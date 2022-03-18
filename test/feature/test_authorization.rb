# frozen_string_literal: true
# typed: ignore

require_relative '../test_helper'

class TestPackageAuthorizationTest < ApplicationSystemTest
  before do
    if ENV['CI']
      skip("cannot run on CI yet")
    end
  end

  def switch_to_stripe_lightning_application_context
    page.switch_to_frame(page.send(:_find_frame, 0))
  end

  def visit_stripe_app
    visit "https://appiphony15-dev-ed.lightning.force.com/lightning/n/Setup"
    initial_assert_exists
  end

  def login_to_salesforce
    visit "https://appiphony15-dev-ed.my.salesforce.com"
    fill_in "Username", with: ENV.fetch('SF_USERNAME')
    fill_in "Password", with: ENV.fetch('SF_PASSWORD')
    click_on "Log In"
  end

  def click_data_mapper_dropdown(name)
    # TODO first we need to click the area for the picker to render
    # TODO needs to be dynamic with name
    find(:xpath, "//tr[@data-index='6']//c-field-picker").click
  end

  def select_data_mapper_dropdown_item(name)
    find(:xpath, "//span[@title='#{name}']").click
  end

  def click_sidebar_link(name)
    find(:xpath, "//li[@class='stripe-navigation-item']/*[contains(text(), '#{name}')]").click
  end

  # TODO document capybara debugging tricks
  # find_all :button
  # find(...).base['innerHTML']

  describe 'initial launch' do
    it 'initial authorization and signup flow' do
      login_to_salesforce

      visit_stripe_app

      switch_to_stripe_lightning_application_context

      # TODO once we wipe the connection data completely, we'll need to mock oauth login
      # mock_omniauth_salesforce

      click_link_or_button "Continue"
      initial_assert_exists

      click_link_or_button "Get Started"
      initial_assert_exists

      # map the Stripe customer name field to the "Account Description" field in salesforce
      click_data_mapper_dropdown "Name"
      select_data_mapper_dropdown_item "Account Description"
      click_link_or_button "Next"

      # TODO once we've tunneled to our database service we can assert on the database state change
    end
  end

  describe 'second launch' do
    it 'data mapper' do
      login_to_salesforce

      visit_stripe_app

      switch_to_stripe_lightning_application_context

      click_sidebar_link "Data Mapping"

      # map the Stripe customer field to the Salesforce account description
      click_data_mapper_dropdown "Name"
      select_data_mapper_dropdown_item "Account Description"
    end

    it 'authorization status' do

    end

    it 'settings page' do

    end
  end

end
