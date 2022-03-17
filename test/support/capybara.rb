# frozen_string_literal: true
# typed: strict
verbose_logging = ENV['FEATURE_TEST_LOGGING'] && ENV['FEATURE_TEST_LOGGING'] == 'true'

if verbose_logging
  # https://gist.github.com/twalpole/2c97a3d5837aabe747c28a52d4c5de47
  require "rack/handler/webrick"
  Capybara.server = :webrick, {Logger: WEBrick::Log.new($stdout), AccessLog: [[$stdout, WEBrick::AccessLog::COMMON_LOG_FORMAT]]}
  Webdrivers.logger.level = :DEBUG
else
  Capybara.server = :webrick
end

if !ENV['CI']
  # chromedriver is an approved stripe application, but it needs to be located within `~/stripe`
  # make sure stripe-netsuite is in this folder and you won't encounter any issues running tests
  Webdrivers.install_dir = File.expand_path('vendor/bundle/webdrivers', Rails.root)
end

Capybara.raise_server_errors = false
Capybara.run_server  = true
Capybara.server_port = ENV['RAILS_TEST_SERVER_PORT']
Capybara.default_max_wait_time = 2

# this binds the server to any interface
Capybara.server_host = '0.0.0.0'

# normally, `app_host` would be set, but `host!` in the rails system tests handle that setup

# original definition: https://github.com/teamcapybara/capybara/blob/ad5d347e43ddb5f90763751fa10c4a9fdf7939b3/lib/capybara/registrations/drivers.rb#L18
# https://gist.github.com/iloveitaly/d7e88e3a596caf2cfb690944d5fcd228
Capybara.register_driver :selenium_chrome_driver do |app|
  client = Selenium::WebDriver::Remote::Http::Default.new

  # recent capybara versions have these two separate properties
  client.open_timeout = 60 * 5
  client.read_timeout = 60 * 5

  logging_opts = if verbose_logging
    {
      # https://github.com/SeleniumHQ/selenium/blob/5e97ce3a52a38bdf10701afcd22e36e1ab04e0ea/rb/lib/selenium/webdriver/common/service.rb
      # https://github.com/SeleniumHQ/selenium/blob/5e97ce3a52a38bdf10701afcd22e36e1ab04e0ea/rb/lib/selenium/webdriver/chrome/service.rb#L36-L46
      log_path: "#{Rails.root}/log/capybara.log",
      verbose: true,
    }
  else
    {}
  end

  browser_options = %w{--ignore-certificate-errors --disable-site-isolation-trials}

  # https://github.com/spark-solutions/spree/blob/49b5215d9c6c49c8aa0d4adc592ebca0c8f84a8e/core/lib/spree/testing_support/capybara_config.rb
  if verbose_logging
    Selenium::WebDriver.logger.level = :info
  end

  browser_options += %w{no-sandbox disable-dev-shm-usage disable-popup-blocking disable-gpu window-size=1920,1080 --enable-features=NetworkService,NetworkServiceInProcess --disable-features=VizDisplayCompositor}

  # add `--headless` in your local environment to debug weird local-only chrome errors
  if ENV['CI'] || (ENV['CAPYBARA_HEADLESS'] && ENV['CAPYBARA_HEADLESS'] == 'true')
    browser_options << 'headless'
  end

  # https://github.com/teamcapybara/capybara/blob/master/lib/capybara/selenium/driver.rb
  Capybara::Selenium::Driver.new(app, {
    # NOTE this entire hash (excluding some "special" params) are passed off to the bridge
    # https://github.com/SeleniumHQ/selenium/blob/5e97ce3a52a38bdf10701afcd22e36e1ab04e0ea/rb/lib/selenium/webdriver/common/driver.rb#L301

    # `browser` is one of the 'special' options defined here:
    # https://github.com/teamcapybara/capybara/blob/4786328b84db0d7a43f0d85142d72fdf369639ee/lib/capybara/selenium/driver.rb#L14
    browser: :chrome,

    http_client: client,

    service: Selenium::WebDriver::Service.chrome(args: logging_opts),
    options: Selenium::WebDriver::Chrome::Options.new(
      args: browser_options,
      # TODO phased out in a recent version of capyabara
      # log_level: :info
    ),
  })
end

if ENV['CIRCLE_ARTIFACTS'].present?
  Capybara.save_path = ENV['CIRCLE_ARTIFACTS']
end

# https://github.com/thoughtbot/capybara-webkit/issues/143
# Capybara::Screenshot.register_driver(:selenium_chrome_driver) do |driver, path|
#   driver.save_screenshot(path)
# end

# try `:selenium_chrome` if you are running into local chrome issues
Capybara.default_driver = :selenium_chrome_driver
