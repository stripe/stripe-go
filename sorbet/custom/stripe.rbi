# typed: true

# TODO I think this is true? Validate before pulling this into upstream
class Stripe::APIResource
  def save; end
end

class Stripe::Price
  sig { returns(Stripe::Price).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end

  sig { returns(Integer) }
  def unit_amount; end

  sig { returns(String) }
  def unit_amount_decimal; end

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

# these are NOT the exact same structure as a subscription item
# TODO report this as a bug, this is confusing
class Stripe::SubscriptionSchedulePhaseSubscriptionItem < Stripe::StripeObject
  sig { returns(String) }
  def price; end

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

class Stripe::SubscriptionScheduleSettings
  sig { returns(String) }
  def collection_method; end
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