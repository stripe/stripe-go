# typed: true

# TODO I think this is true? Validate before pulling this into upstream
class Stripe::APIResource
  def save; end
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

# TODO I swear I added some of these in via the netsuite connector
class Stripe::Customer
  sig { returns(String) }
  def currency; end

  sig { returns(String) }
  def description; end

  sig { returns(T.nilable(String))}
  def test_clock; end

  # TODO should add proper subhash typing
  def shipping; end

  sig { returns(T.nilable(String)) }
  def phone; end
end

# TODO the address subhash typings need to be cleaned up
class Stripe::Address
  sig { returns(T.nilable(String)) }
  def phone; end
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

  sig { returns(Integer)}
  def end_date; end

  sig { params(arg: Integer).void}
  def end_date=(arg); end

  sig { returns(T::Hash[T.any(String, Symbol), T.untyped]) }
  def metadata; end
end

class Stripe::SubscriptionScheduleInvoiceSettings
  sig {returns(Integer)}
  def days_until_due; end
end

class Stripe::SubscriptionScheduleSettings
  sig { returns(String) }
  def collection_method; end

  sig { returns(Stripe::SubscriptionScheduleInvoiceSettings)}
  def invoice_settings; end
end

class Stripe::SubscriptionSchedule
  sig { returns(Stripe::SubscriptionScheduleSettings) }
  def default_settings; end

  sig { returns(T.any(Stripe::Subscription, String))}
  def subscription; end

  sig { returns(T.any(Stripe::Customer, String))}
  def customer; end

  # TODO can this be nil in specific situations?
  sig { returns(T.any(Stripe::Subscription, String))}
  def subscription; end

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

class Stripe::SubscriptionItem < Stripe::APIResource
  sig { params(arg:Hash).returns(Stripe::SubscriptionItem)}
  def self.construct_from(arg); end

  sig { returns(Integer)}
  def quantity; end

  sig { params(arg:Integer).void}
  def quantity=(arg); end
end