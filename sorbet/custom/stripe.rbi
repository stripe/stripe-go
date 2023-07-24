# typed: true

# TODO I think this is true? Validate before pulling this into upstream
class Stripe::APIResource
  def save; end
end

class Stripe::StripeObject
  sig {returns(T.any(Stripe::Subscription, Stripe::Customer, Stripe::Invoice, Stripe::InvoiceItem))}
  def object; end
end

class Stripe::InvoiceItem
  sig { returns(T.any(String, Stripe::Subscription))}
  def subscription; end
end

class Stripe::Price
  sig { returns(Stripe::Price).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end

  sig { returns(Stripe::Price).params(params: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.construct_from(params={}); end

  sig { returns(Integer) }
  def unit_amount; end

  sig { returns(String) }
  def unit_amount_decimal; end

  sig { params(arg: T.any(String, BigDecimal)).void }
  def unit_amount_decimal=(arg); end

  sig { returns(String) }
  def billing_scheme; end

  sig { returns(String) }
  def type; end

  sig { returns(String) }
  def currency; end

  sig { returns(String) }
  def tiers_mode; end

  sig { returns(Array) }
  def tiers; end

  sig { returns(T.any(Stripe::Product, String)) }
  def product; end

  # TODO maybe I should use array accessor instead of defining this? Need to look at existing pattern in typing
  def product=(arg); end

  def recurring; end

  def custom_unit_amount; end

  sig { returns(T.nilable(String))}
  def lookup_key; end

  sig { returns(T.nilable(String))}
  def nickname; end

  sig { returns(String)}
  def tax_behavior; end

  def transform_quantity; end

  sig { returns(T::Boolean)}
  def active; end

  sig { params(arg: T::Boolean).void}
  def active=(arg); end
end

class Stripe::Customer
  sig { returns(String) }
  def currency; end

  sig { returns(T.nilable(Stripe::Address)) }
  def address; end

  sig { returns(String) }
  def description; end

  sig { returns(String) }
  def email; end

  sig { returns(String) }
  def name; end

  sig { returns(T.nilable(String, Stripe::TestHelpers::TestClock))}
  def test_clock; end

  sig { returns(Stripe::Shipping) }
  def shipping; end

  sig { returns(T.nilable(String)) }
  def phone; end
end

class Stripe::Shipping
  sig { returns(T.nilable(Stripe::Address)) }
  def address; end

  sig { returns(T.nilable(String)) }
  def name; end

  sig { returns(T.nilable(String)) }
  def phone; end
end

class Stripe::Address
  sig { returns(T.nilable(String)) }
  def city; end

  sig { returns(T.nilable(String)) }
  def country; end

  sig { returns(T.nilable(String)) }
  def line1; end

  sig { returns(T.nilable(String)) }
  def line2; end

  sig { returns(T.nilable(String)) }
  def postal_code; end

  sig { returns(T.nilable(String)) }
  def state; end
end

class Stripe::SubscriptionSchedulePhaseInvoiceItem < Stripe::StripeObject
  sig { returns(T.any(String, Stripe::Price)) }
  def price; end

  sig { returns(Integer) }
  def quantity; end

  def period; end

  def metadata; end
end

# these are NOT the exact same structure as a subscription item
# TODO report this as a bug, this is confusing
class Stripe::SubscriptionSchedulePhaseSubscriptionItem < Stripe::StripeObject
  sig { returns(T.any(String, Stripe::Price)) }
  def price; end

  sig { params(arg: String).void }
  def price=(arg); end

  sig { returns(String) }
  def plan; end

  sig { returns(Integer) }
  def quantity; end

  sig { params(arg: Integer).void }
  def quantity=(arg); end

  sig { returns(T::Hash[T.any(String, Symbol), T.untyped]) }
  def metadata; end

  sig { returns(T::Array[Stripe::Discounts]) }
  def discounts; end
end

class Stripe::SubscriptionSchedulePhase < Stripe::StripeObject
  sig { returns(T::Array[Stripe::SubscriptionSchedulePhaseSubscriptionItem])}
  def items; end

  sig { returns(T::Array[Stripe::SubscriptionSchedulePhaseSubscriptionItem])}
  def add_invoice_items; end

  sig { returns(Integer)}
  def start_date; end

  sig { params(arg: Integer).void }
  def start_date=(arg); end

  sig { returns(String) }
  def proration_behavior; end

  sig { returns(Integer)}
  def end_date; end

  sig { params(arg: Integer).void}
  def end_date=(arg); end

  sig { returns(T::Hash[T.any(String, Symbol), T.untyped]) }
  def metadata; end

  sig { returns(T::Array[Stripe::Discounts]) }
  def discounts; end
end

class Stripe::InvoiceSettingRendering
  sig { returns(String) }
  def template; end

  sig { returns(Integer) }
  def template_version; end
end

class Stripe::SubscriptionScheduleInvoiceSettings
  sig {returns(Integer)}
  def days_until_due; end

  sig {returns(Stripe::InvoiceSettingRendering)}
  def rendering; end
end

class Stripe::SubscriptionScheduleSettings
  sig { returns(String) }
  def collection_method; end

  sig { returns(Stripe::SubscriptionScheduleInvoiceSettings)}
  def invoice_settings; end
end

class Stripe::SubscriptionSchedulePrebilling
  sig { returns(String) }
  def invoice; end

  sig { returns(Integer) }
  def period_start; end

  sig { returns(Integer) }
  def period_end; end
end

class Stripe::SubscriptionSchedule
  sig { returns(Stripe::SubscriptionScheduleSettings) }
  def default_settings; end

  sig { returns(Stripe::SubscriptionSchedulePrebilling)}
  def prebilling; end

  sig { returns(T.nilable(T.any(Stripe::Subscription, String)))}
  def subscription; end

  sig { returns(T.any(Stripe::Customer, String))}
  def customer; end

  sig { returns(T::Array[Stripe::SubscriptionSchedulePhase])}
  def phases; end

  sig { params(arg: T::Array[Stripe::SubscriptionSchedulePhase]).void }
  def phases=(arg); end

  sig { returns(String) }
  def status; end

  sig { returns(String) }
  def proration_behavior; end

  sig { params(arg: String).void }
  def proration_behavior=(arg); end

  sig { returns(Stripe::SubscriptionSchedule).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end
end

class Stripe::Subscription
  sig { returns(T::Boolean)}
  def cancel_at_period_end; end

  sig { returns(T.nilable(Integer)) }
  def cancel_at; end

  sig { returns(T.any(Stripe::Customer, String))}
  def customer; end

  sig { returns(T::Hash[T.any(String, Symbol), T.untyped])}
  def items; end

  sig { returns(String) }
  def status; end

  sig { returns(Integer) }
  def current_period_start; end

  sig { returns(Integer) }
  def current_period_end; end

  sig { returns(T::Hash[T.any(String, Symbol), T.untyped]) }
  def metadata; end

  sig { returns(Stripe::Subscription).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end
end

class Stripe::SubscriptionItem < Stripe::APIResource
  sig { params(arg:Hash).returns(Stripe::SubscriptionItem)}
  def self.construct_from(arg); end

  sig { returns(Integer)}
  def quantity; end

  sig { params(arg:Integer).void}
  def quantity=(arg); end

  sig { params(arg:Stripe::Discounts).returns(Stripe::Discounts) }
  def discounts=(arg); end
end

class Stripe::Invoice
  sig { returns(T::Boolean)}
  def auto_advance; end
  
  sig { returns(Integer)}
  def period_end; end

  sig { returns(Integer)}
  def period_start; end
end

class Stripe::InvoiceItem
  sig { returns(T.any(String, Stripe::Subscription))}
  def subscription; end
end

class Stripe::TestHelpers::TestClock
  sig { returns(Stripe::TestHelpers::TestClock).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end

  sig { returns(String)}
  def status; end

  sig { returns(Integer)}
  def frozen_time; end
end

class Stripe::Coupon
  sig { returns(String)}
  def id; end

  sig { returns(Integer)}
  def percent_off; end

  sig { returns(Integer)}
  def amount_off; end

  sig { returns(String)}
  def duration; end

  sig { returns(String)}
  def duration_in_months; end

  sig { returns(String)}
  def max_redemptions; end

  sig { returns(String)}
  def currency; end

  sig { returns(T::Boolean)}
  def valid; end
end

class Stripe::Discounts
  sig { returns(String)}
  def coupon; end

  sig { returns(String)}
  def discount; end
end
class Stripe::RevenueContract < Stripe::StripeObject
  sig { returns(String)}
  def id; end

  sig { returns(String)}
  def customer; end
  
  sig { returns(String)}
  def currency; end
  
  sig { returns(Integer)}
  def signed_at; end
  
  sig { returns(Integer)}
  def version; end
  
  sig { returns(T::Array[Stripe::RevenueContractBillingModel])}
  def billing_models; end
  
  sig { returns(Stripe::RevenueContractItemsListOjbect)}
  def items; end
end

class Stripe::RevenueContractBillingModel
  sig { returns(String)}
  def type; end

  sig { returns(String)}
  def subscription_schedule; end
end

class Stripe::RevenueContractItemsListOjbect < Stripe::StripeObject
  sig { returns(T::Array[Stripe::RevenueContractItem])}
  def data; end
  
  sig { returns(T::Boolean)}
  def has_more; end

  sig { returns(String)}
  def url; end
end

class Stripe::RevenueContractItem < Stripe::StripeObject
  sig { returns(String)}
  def id; end

  sig { returns(Integer)}
  def amount_subtotal; end
  
  sig { returns(String)}
  def price; end
  
  sig { returns(Integer)}
  def quantity; end
  
  sig { returns(T.nilable(Stripe::RevenueContractItemTerminationForConvenience))}
  def termination_for_convenience; end
  
  sig { returns(Stripe::RevenueContractItemPeriod)}
  def period; end
end

class Stripe::RevenueContractItemTerminationForConvenience
  sig { returns(Integer)}
  def expires_at; end
end

class Stripe::RevenueContractItemPeriod
  sig { returns(Integer)}
  def start; end

  sig { returns(Integer)}
  def end; end
end