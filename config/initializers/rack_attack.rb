# frozen_string_literal: true
# typed: false

Rack::Attack.enabled = false

class Rack::Attack
  THROTTLED_CONTROLLERS = %w{
    customer_portal
    payments
  }

  throttle('customer_portal/ip', limit: 5, period: 1.minute) do |request|
    recognized_path = begin
      Rails.application.routes.recognize_path(request.path, method: request.request_method)
    rescue => e
      # the error we'd expect here in most cases is ActionController::RoutingError
      # however, with routes authenticated via devise, the request env is thrown away and will trigger
      # a unknown method on nil error when checking if the user is authenticated
      # related to: https://github.com/heartcombo/devise/issues/1670

      SimpleStructuredLogger::Writer.instance.warn 'unable to determine path controller', error: e.message
      nil
    end

    if THROTTLED_CONTROLLERS.include?(recognized_path&.dig(:controller)) && request.get?
      request.ip
    end
  end
end

ActiveSupport::Notifications.subscribe("throttle.rack_attack") do |_name, _start, _finish, _request_id, payload|
  # request object available in payload[:request]
  # we have access to the name and other data about the matched throttle
  #  request.env['rack.attack.matched'],
  #  request.env['rack.attack.match_type'],
  #  request.env['rack.attack.match_data'],
  #  request.env['rack.attack.match_discriminator']
  request = payload.fetch(:request)

  Raven.capture_message(
    "Rate Limit Exceeded",
    extra: {
      ip: request.ip,
      path: request.path,
      user_agent: request.user_agent,
    }
  )
end
