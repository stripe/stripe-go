# This is an autogenerated file for dynamic methods in StripeForce::PollTimestamp
# Please rerun bundle exec scripts/sequel-types.rb to regenerate
# typed: strong
module StripeForce::PollTimestamp::GeneratedAttributeMethods
  sig { returns(DateTime) }
  def created_at; end

  sig { params(value: T.any(DateTime, Date, Time)).void }
  def created_at=(value); end

  sig { returns(Integer) }
  def id; end

  sig { params(value: T.any(Numeric)).void }
  def id=(value); end

  sig { returns(T.nilable(String)) }
  def integration_record_type; end

  sig { params(value: T.nilable(T.any(String, Symbol))).void }
  def integration_record_type=(value); end

  sig { returns(DateTime) }
  def last_polled_at; end

  sig { params(value: T.any(DateTime, Date, Time)).void }
  def last_polled_at=(value); end

  sig { returns(T.nilable(T::Boolean)) }
  def livemode; end

  sig { params(value: T.nilable(T::Boolean)).void }
  def livemode=(value); end

  sig { returns(String) }
  def stripe_account_id; end

  sig { params(value: T.any(String, Symbol)).void }
  def stripe_account_id=(value); end

  sig { returns(DateTime) }
  def updated_at; end

  sig { params(value: T.any(DateTime, Date, Time)).void }
  def updated_at=(value); end
end

class StripeForce::PollTimestamp
  include StripeForce::PollTimestamp::GeneratedAttributeMethods

  sig { params(value: Integer).returns(T.nilable(StripeForce::PollTimestamp)) }
  def self.[](value); end

  sig { returns(T.nilable(StripeForce::PollTimestamp)) }
  def self.first; end

  sig { returns(T.nilable(StripeForce::PollTimestamp)) }
  def self.last; end
end
