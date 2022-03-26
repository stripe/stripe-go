# typed: true

class Stripe::Price
  sig { returns(Stripe::Price).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end

  sig { returns(Integer) }
  def unit_amount; end
end

class Stripe::Customer
  sig { returns(String) }
  def currency; end
end

# these are NOT the exact same structure as a subscription item
# TODO report this as a bug, this is confusing
class Stripe::SubscriptionSchedulePhaseSubscriptionItem
  sig { returns(String) }
  def price; end

  sig { returns(String) }
  def plan; end

  sig { returns(Integer) }
  def quantity; end
end

class Stripe::SubscriptionSchedulePhase
  sig { returns(T::Array[Stripe::SubscriptionSchedulePhaseSubscriptionItem])}
  def items; end

  sig { returns(T::Array[Stripe::SubscriptionSchedulePhaseSubscriptionItem])}
  def add_invoice_items; end

  sig { returns(Integer)}
  def start_date; end

  sig { returns(Integer)}
  def end_date; end
end

class Stripe::SubscriptionScheduleSettings
  sig { returns(String) }
  def collection_method; end
end

class Stripe::SubscriptionSchedule
  sig { returns(Stripe::SubscriptionScheduleSettings) }
  def default_settings; end

  sig { returns(T.any(Stripe::Customer, String))}
  def customer; end

  # TODO can this be nil in specific situations?
  sig { returns(T.any(Stripe::Subscription, String))}
  def subscription; end

  sig { returns(T::Array[Stripe::SubscriptionSchedulePhase])}
  def phases; end

  sig { returns(Stripe::SubscriptionSchedule).params(id: T.any(String, T::Hash[Symbol, T.untyped]), opts: T.nilable(T::Hash[Symbol, T.untyped])) }
  def self.retrieve(id, opts={}); end
end