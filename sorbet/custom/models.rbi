# typed: strict
class StripeForce::User
  sig { returns(T::Array[Symbol]) }
  def feature_flags; end

  sig { returns(T::Hash[String, T::Hash[String, String]]) }
  def field_mappings; end

  sig { returns(T::Hash[String, T::Hash[String, T.nilable(T.any(String, Integer, T::Boolean))]]) }
  def field_defaults; end
end
