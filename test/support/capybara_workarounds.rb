# frozen_string_literal: true
# typed: false

module Critic
  module CapybaraWorkarounds
    # include Capybara::DSL
    # include Minitest::Assertions

    # sig { returns(T.untyped) }
    # def page; end

    # https://robots.thoughtbot.com/automatically-wait-for-ajax-with-capybara
    def wait_for_ajax(wait_time=nil)
      # NOTE Capybara.default_max_wait_time is 2s, we need to extend this
      #      quite a bit to account for sync NS communication during testing.
      #      Most NetSuite communication is done server side, but some AJAX calls
      #      can trigger NS communication which is why the wait time here is so high

      # NOTE When working with iframes (especially nested frames), be sure to call
      #       wait_for_ajax from within the within_frame block.

      # TODO should eliminate explicit wait times from the feature tests!
      #      we should use an indicator of why we are waiting instead `wait_for_ajax(:netsuite)`
      #      which can adjust the time for us.

      wait_time ||= 60
      max_time = Time.now.to_i + wait_time

      loop do
        if finished_loading? && aura_finished_loading?
          break
        elsif Time.now.to_i < max_time
          puts "WARNING: Testing harness could not detect page load completion! This will most likely cause a feature test error."
          break
        else
          sleep(0.2)
        end
      end

      # TODO we should remove all instances where a block is passed
      yield if block_given?
    end

    WAIT_FOR_AURA_SCRIPT = <<~EOL
      done = arguments[0];
      if (document.getElementById('auraAppcacheProgress')) {
          var waitForXHRs = function() {
              if (window.$A && !window.$A.clientService.inFlightXHRs()) {
                  done();
              } else {
                  setTimeout(waitForXHRs, 100);
              }
          }
          setTimeout(waitForXHRs, 0);
      } else {
          done();
      }
    EOL

    def aura_finished_loading?
      page.evaluate_async_script(WAIT_FOR_AURA_SCRIPT)
    end

    # http://stackoverflow.com/questions/1012140/delaying-a-jquery-script-until-everything-else-has-loaded
    # http://stackoverflow.com/questions/7490461/check-if-the-jquery-page-load-events-fired-already
    def finished_loading?
      page.evaluate_script('document.readyState === "complete"') == true
    end

    # NOTE locally, we are getting random errors that have no recent references on the internet
    #      this is primarily caused by the chrome version set by "Managed Software Center" combined
    #      with chromedriver versions that don't match up perfectly. Overriding the `visit` call and
    #      forcing a retry if the page fails to load works surprisingly well but should NOT
    #      be used on CI

    # if !ENV['CI']
    #   def visit(*args)
    #     begin
    #       super
    #     rescue Selenium::WebDriver::Error::WebDriverError => e
    #       if e.message.include?('received Inspector.detached event') || e.message.include?('Unable to receive message from renderer')
    #         puts "ERROR: received strange chrome error, attempting retry!"
    #         super
    #       else
    #         raise
    #       end
    #     end
    #   end
    # end

    # we can't detect all javascript execution: if the JS does not make a visible modification to the page
    # sometimes the only way to know if a non-AJAX JS is complete is wait
    # for the md5 of the current page state to 'settle'
    def wait_for_javascript_execution
      max_wait = Time.now.to_i + 30
      initial_page = page_uid
      previous_page_state = T.let(nil, T.untyped)

      while Time.now.to_i < max_wait
        current_page_state = page_uid

        if initial_page != current_page_state && !previous_page_state.nil? && previous_page_state == current_page_state
          puts "page state changed and stayed the same for 2 seconds, page JS execution is complete"
          break
        end

        puts "waiting for javascript execution to complete..."
        previous_page_state = current_page_state

        sleep 2
      end
    end

    # after ajax has loaded, javascript may still execute
    # use this helper the first time you are checking for `true`
    # on a page to give the page a little bit more time to finish executing
    def initial_assert_exists
      wait_for_ajax

      if block_given?
        using_wait_time(130) do
          assert(yield)
        end
      end

      proposed_uid = page_uid

      # in some cases, the page_uid can change after we assert our match
      # let's ensure the page content doesn't change for the next couple seconds
      # to avoid flakey tests. Yes, this is incredibly hacky, but I don't believe
      # there's not anything else we can do to detect JS execution
      5.times do
        sleep(5)
        break if proposed_uid == page_uid
      end

      @initial_assert_complete = page_uid
    end

    # if we are asserting that something *doesn't* exist, we want a short wait time
    # this should *not* be used as the first assertion after a page load
    def refute_exists(wait: 2)
      if @initial_assert_complete.nil? || page_uid != @initial_assert_complete
        raise "you must make an initial assert first before refuting any elements to ensure capybara has finished loading"
      end

      using_wait_time(wait) do
        refute(yield)
      end
    end

    def page_uid
      Digest::MD5.hexdigest(page.body)
    end

    def safely_fill_in(field_name, value)
      5.times do
        # NOTE for some weird reason the letter "3" stopped working via send_keys
        #      avoid values with a "3"
        value.to_s.chars.each do |piece|
          find_field(field_name).send_keys(piece)
          sleep(0.1)
        end

        if find_field(field_name).value.gsub(/[^0-9]/, '') != value.to_s
          puts "error inputting #{value} into #{field_name}, retrying"
          # https://stackoverflow.com/questions/17119870/deleting-content-from-text-field-with-capybara
          fill_in(field_name, with: '', fill_options: {clear: :backspace})
        else
          break
        end
      end
    end

    def fill_and_submit_plaid_ach
      within_frame 'plaid-link-iframe-1' do
        # continue button on the privacy panel
        aggressive_click('#aut-continue-button')

        # Use an institution that does not require extra authentication
        aggressive_click('*[data-institution="nfcu"]')

        fill_in 'username', with: 'user_good'
        fill_in 'password', with: 'pass_good'

        aggressive_click("button[type=submit]")

        assert(page.has_text?('Your accounts', wait: 20) || page.has_text?('Select account', wait: 20))

        # checking vs savings
        aggressive_click('.ListItem:first-child')

        # continue
        aggressive_click('#aut-continue-button')
      end
    end

    def fill_and_submit_stripe_checkout(card, button_name: 'Pay with Card', has_zip: false, has_billing: false, has_email: false)
      wait_for_ajax

      click_on button_name

      within_frame('stripe_checkout_app') do
        if has_billing
          find_field('Name').send_keys "Michael Bianco"
          find_field('Address').send_keys "123 Great Street"
          find_field('City').send_keys "Downingtown"
          find_field('ZIP').send_keys "19382"

          if has_email
            find_field('Email').send_keys "#{Time.now.to_i}@example.com"
          end

          # the selection changes based on IP; force US for consistency
          find('#billing-country').find(:option, 'United States').select_option

          find('button[type="submit"]').click
        end

        safely_fill_in('Card number', card)

        safely_fill_in('MM / YY', '0123')
        find_field('CVC').send_keys('125')

        if has_zip && !has_billing
          find_field('ZIP Code').send_keys '80110'
        end

        # sometimes, chrome doesn't click the button OR let us know!
        aggressive_click('#submitButton')
      end
    end

    def fill_stripe_elements(card='4242424242424242')
      wait_for_ajax do
        within_frame(:css, "iframe[src*='js.stripe.com']") do
          safely_fill_in('cardnumber', card)
          safely_fill_in('exp-date', '1222')

          find_field('cvc').send_keys '124'
          find_field('postal').send_keys '80110'
        end
      end
    end

    def confirm_sca_challenge
      within_sca_challenge_frame do
        aggressive_click "#test-source-authorize-3ds"
      end
    end

    def fail_sca_challenge
      within_sca_challenge_frame do
        aggressive_click "#test-source-fail-3ds"
      end
    end

    def within_sca_challenge_frame(&block)
      wait_for_ajax

      # `within_frame` locates the frame using `find`
      # https://github.com/teamcapybara/capybara/blob/4d880c237c6f0d8c907a6a6469233677ea3ea8a4/lib/capybara/session.rb#L893
      # when debugging, you can use `find` to poke around at various Stripe iframes

      aggressive_within_frame do
        # <iframe frameborder="0" allowtransparency="true" scrolling="no" name="__privateStripeFrame4" allowpaymentrequest="true" src="https://js.stripe.com/v3/authorize-with-url-inner-960a9b61ff03b2fcf169e0e12be07627.html#url=https%3A%2F%2Fhooks.stripe.com%2Fredirect%2Fauthenticate%2Fsrc_0GuPDMnx3du3xKk7XvvaFvKK%3Fclient_secret%3Dsrc_client_secret_iQfjaz9QyzVlsO6L4elYWFG2&amp;locale=en&amp;intentId=pi_0GuPDLnx3du3xKk7m21RdNRE&amp;source=src_0GuPDMnx3du3xKk7XvvaFvKK&amp;origin=http%3A%2F%2Flvh.me%3A8081&amp;referrer=http%3A%2F%2Flvh.me%3A8081%2Fcustomers%2Ftest%2Facct_TEST_NETSUITE_ACCOUNT%2F14481950%2F1592228820&amp;controllerId=__privateStripeController1" cd_frame_id_="2d24de75bf2eadf6324dd0e1164fa1c1"></iframe>
        within_frame(:css, 'iframe[name^=__privateStripeFrame]') do
          # <iframe id="challengeFrame" name="__stripeJSChallengeFrame" frameborder="0" height="100%" width="100%" class="AuthorizeWithUr lApp-content" src="https://hooks.stripe.com/redirect/authenticate/src_G59WIxJv8SK3zw?client_secret=src_client_secret_G59WqR3Qgc95EsImaKDn 0fd9"></iframe>
          within_frame("__stripeJSChallengeFrame") do
            # <iframe class="FullscreenFrame" name="acsFrame"></iframe>
            within_frame("acsFrame") do
              yield
            end
          end
        end
      end
    end

    # we've ran into flakeyness on tests which need to switch into a iframe-within-an-iframe
    # this (hackily) solves the issue consistently. Right now, this is only necessary when
    # clicking SCA challenge buttons. If a frame-within-a-frame can't be found, the `frame` logic
    # looks to kick the user back to previous parent frame, which moves the context away from the
    # first node in the frame tree you are trying to access. The easiest fix for this is just to
    # attempt to retry a block containing the nested frame invocations.
    def aggressive_within_frame(&block)
      count = 0

      begin
        yield
      rescue Capybara::ElementNotFound => e
        # the error message is different, but always contains these two strings
        if count < 3 && e.message.include?('Unable to find') && e.message.include?('frame')
          puts "ERROR: could not find frame, trying again"
          count += 1
          sleep(count * count)
          retry
        else
          raise
        end
      end
    end

    # sometimes, chrome doesn't click the button OR let us know!
    # pretty certain this is caused by javascript UI frameworks
    def aggressive_click(identifier)
      3.times do
        begin
          find(identifier).click
        rescue Capybara::ElementNotFound
          puts "element couldn't be clicked, definitely not successful"
        rescue Selenium::WebDriver::Error::StaleElementReferenceError,
               Selenium::WebDriver::Error::WebDriverError,
               Selenium::WebDriver::Error::UnknownError
          puts "stale element, or unknown error, could indicate that a previous click was successful"
        else
          begin
            break if !find(identifier, wait: 2).visible?
          rescue Selenium::WebDriver::Error::StaleElementReferenceError
            puts "stale element error when testing element visibility, retrying"
          rescue Selenium::WebDriver::Error::NoSuchWindowError
            puts "could not find window, may have successfully clicked: #{identifier}"
          rescue Capybara::ElementNotFound
            puts "element could not be found, may have successfully clicked: #{identifier}"
            break
          end
        end

        puts "did not successfully click the target button, trying direct JS: #{identifier}"

        begin
          # https://medium.com/@yuliaoletskaya/capybara-inconsistent-click-behavior-and-flickering-tests-f50b5fae8ab2
          page.execute_script("document.querySelector('#{identifier}').click()")
        rescue Selenium::WebDriver::Error::JavascriptError
          # noop
        end
      end
    end
  end
end
