# This file is autogenerated. Do not edit it by hand. Regenerate it with:
#   srb rbi gems

# typed: strict
#
# If you would like to make changes to this file, great! Please create the gem's shim here:
#
#   https://github.com/sorbet/sorbet-typed/new/master?filename=lib/omniauth-stripe/all/omniauth-stripe.rbi
#
# omniauth-stripe-0.1.0

module Omniauth
end
module Omniauth::Stripe
end
module OmniAuth
end
module OmniAuth::Strategies
end
class OmniAuth::Strategies::Stripe < OmniAuth::Strategies::OAuth2
  def callback_url; end
  def raw_info; end
  extend OmniAuth::Strategy::ClassMethods
end
