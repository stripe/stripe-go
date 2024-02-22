# frozen_string_literal: true
# typed: true
require 'openssl'
require 'base64'
require 'json'
require 'sorbet-runtime'

# noinspection SpellCheckingInspection
DEFAULT_STATE_ENCRYPTION_KEY = '98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a'
DEFAULT_STATE_SIGNING_KEY = '98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a'
JWT_HEADER = {'typ': 'JWT', 'alg': 'HS256'}

# Implementation guided from https://stackoverflow.com/questions/8489486/encryption-and-decryption-algorithm-in-rails
# Pattern taken from kms_encryption
# Provides a mechanism for Salesforce and Ruby to share trusted state information
module StateEncryptionAlgo
  extend T::Sig

  def self.included(klass)
    super
    klass.send(:include, InstanceMethods)
  end

  sig { params(state: StateEncryptionAlgo::StripeOAuthState, defaults: Hash).returns(StateEncryptionAlgo::StripeOAuthState) }
  def self.apply_defaults(state, defaults)
    defaults.each_pair {|k, p| state.send("#{k}=", p) unless state.send(k) == p } unless state.nil?
    state
  end

  module InstanceMethods
    extend T::Sig

    sig { params(encrypted_state: T.nilable(String), defaults: Hash).returns(T.nilable(StateEncryptionAlgo::StripeOAuthState)) }
    def restore_state(encrypted_state, defaults={})
      state = StripeOAuthState.from_encrypted_state(encrypted_state)
      StateEncryptionAlgo.apply_defaults(state, defaults) unless state.nil?
    end

    sig { params(defaults: Hash).returns(StateEncryptionAlgo::StripeOAuthState) }
    def create_state(defaults={})
      state = StripeOAuthState.create
      StateEncryptionAlgo.apply_defaults(state, defaults)
    end
  end

  class StripeOAuthState
    extend T::Sig

    attr_accessor :data, :changed

    def initialize(value)
      if value.is_a? Hash
        @data = value
        @externally_generated = true
        @changed = false
      else
        # Dunno why setting it like a normal hash w/ strings wouldn't work but w/e...
        @data = {}
        @data["oauth_version"] = 'v1'
        @data["org_id"] = nil
        @data["user_id"] = nil
        @data["csrf"] = value
        @externally_generated = false
        @changed = true
      end
    end

    sig { params(field: Symbol, key: T.nilable(Symbol)).void }
    def self.json_field(field, key=nil)
      key = field if key.nil?
      define_method(field.to_s) do
        data_hash = instance_variable_get("@data")
        data_hash[key.to_s]
      end
      define_method("#{field}=") do |value|
        instance_variable_set("@changed", true)
        data_hash = instance_variable_get("@data")
        data_hash[key.to_s] = value
        instance_variable_set("@data", data_hash)
      end
      define_method("#{field}?") do
        data_hash = instance_variable_get("@data")
        s_key = key.to_s
        data_hash.key?(s_key) && data_hash[s_key].nil? == false
      end
    end

    json_field :oauth_version
    # Maps to line ~15 in StripeOAuthState.cls
    json_field :salesforce_account_id, :org_id
    json_field :salesforce_namespace, :sf_ns
    json_field :salesforce_instance_type, :sf_it
    json_field :salesforce_instance_subdomain, :sf_sd
    json_field :user_id, :uid
    json_field :csac_id, :csac
    json_field :csrf
    json_field :stripe_account_id, :st_id
    json_field :stripe_account_name, :st_na
    json_field :stripe_account_livemode, :st_lm
    json_field :primary_stripe_account_id, :pst_id
    json_field :primary_stripe_account_livemode, :pst_lm

    def livemode
      self.stripe_account_livemode == "live" || false
    end

    def livemode?
      self.stripe_account_livemode.nil? == false
    end

    def livemode=(value)
      if value == true
        self.stripe_account_livemode = "live"
      else
        self.stripe_account_livemode = "test"
      end
    end

    def primary_livemode
      self.primary_stripe_account_livemode == "live" || false
    end

    def primary_livemode?
      self.primary_stripe_account_livemode.nil? == false
    end

    def primary_livemode=(value)
      if value == true
        self.primary_stripe_account_livemode = "live"
      else
        self.primary_stripe_account_livemode = "test"
      end
    end

    sig { returns(T.nilable(String)) }
    def to_s
      str = JSON.dump(@data)
      StripeOAuthState.encrypt_state(str) || "null"
    end

    sig { returns(T.nilable(String)) }
    def to_s!
      StripeOAuthState.encrypt_state!(JSON.dump(@data)) || "null"
    end

    sig { returns(T::Boolean) }
    def is_externally_generated?
      @externally_generated
    end

    sig { returns(T::Boolean) }
    def is_locally_generated?
      !is_externally_generated?
    end

    sig { params(encrypted_state: T.nilable(String)).returns(T.nilable(StripeOAuthState)) }
    def self.from_encrypted_state(encrypted_state)
      begin
        from_encrypted_state!(encrypted_state)
      rescue => e
        Sentry.capture_exception(e)
        nil
      end
    end

    sig { returns(StripeOAuthState) }
    def self.create
      crypto = cipher
      crypto.encrypt
      StripeOAuthState.new(to_hex(crypto.random_key))
    end

    sig { params(encrypted_state: T.nilable(String)).returns(T.nilable(StripeOAuthState)) }
    def self.from_encrypted_state!(encrypted_state)
      result = decrypt_state!(encrypted_state)
      if result.nil?
        return nil
      end
      data = JSON.parse(result)
      StripeOAuthState.new(data)
    end

    sig { params(str: String).returns(T.nilable(String)) }
    def self.to_sanitized_base64(str)
      b64 = Base64.strict_encode64(str)
      b64.gsub!(/=+$/, '')
      b64.gsub!('+', '-')
      b64.gsub!(%r{/}, '_')
      b64
    end

    sig { params(b64: String).returns(T.nilable(String)) }
    def self.from_sanitized_base64(b64)
      b64.gsub!('-', '+')
      b64.gsub!('_', '/')
      b64 += '=' * (4 - b64.length % 4)
      Base64.decode64(b64)
    end

    sig { params(state: T.nilable(String)).returns(T.nilable(String)) }
    def self.decrypt_state!(state)
      return nil if state.nil?
      parts = state.split('.')
      input = parts[1]
      hash = from_sanitized_base64(T.must(parts[2]))
      # puts "input: #{input}"
      # puts "hash: #{hash}"
      signature = OpenSSL::HMAC.digest('sha256', hasher_key, parts[1])
      # puts signature
      if signature != hash
        return nil
      end
      from_sanitized_base64(T.must(input))
    end

    sig { params(state: T.nilable(String)).returns(T.nilable(String)) }
    def self.decrypt_state(state)
      return nil if state.nil?
      begin
        decrypt_state!(state)
      rescue => e
        Sentry.capture_exception(e)
        nil
      end
    end

    sig { params(state: T.nilable(String)).returns(T.nilable(String)) }
    def self.encrypt_state!(state)
      # puts "state: #{state}"
      return nil if state.nil?
      # c = cipher.encrypt
      # iv = c.random_iv
      # c.key = cipher_key
      # data = c.update(state.to_s) + c.final
      # to_hex(iv.to_s.concat(data.to_s))
      header_b64 = to_sanitized_base64(JSON.dump(JWT_HEADER))
      # puts "header_b64: #{header_b64}"
      return nil if header_b64.nil?
      payload_b64 = to_sanitized_base64(state.to_s)
      # puts "payload_b64: #{payload_b64}"
      return nil if payload_b64.nil?
      signature = OpenSSL::HMAC.digest('sha256', hasher_key, payload_b64)
      # puts "signature: #{signature}"
      return nil if signature.nil?
      sig_b64 = to_sanitized_base64(signature)
      # puts "sanitized sig: #{sig_b64}"
      return nil if sig_b64.nil?
      full = "#{header_b64.strip}.#{payload_b64.strip}.#{sig_b64.strip}"
      # logger = Rails.logger
      # puts "full: #{full}"
      full
    end

    sig { params(state: T.nilable(String)).returns(T.nilable(String)) }
    def self.encrypt_state(state)
      return nil if state.nil?
      begin
        encrypt_state!(state)
      rescue => e
        # puts e
        Sentry.capture_exception(e)
        nil
      end
    end

    sig { params(str: String).returns(String) }
    def self.to_hex(str)
      hex = ""
      str.chars.each do |c|
        hex += c.ord.to_i.to_s(16).rjust(2, "0")
      end
      hex
    end

    sig { returns(OpenSSL::Cipher) }
    def self.cipher
      OpenSSL::Cipher.new('aes-256-cbc')
    end

    sig { returns(T.nilable(String)) }
    def self.cipher_key
      key = ENV.fetch("OAUTH_STATE_ENCRYPTION_KEY") { DEFAULT_STATE_ENCRYPTION_KEY }
      [key].pack('H*')
    end

    sig { returns(T.nilable(String)) }
    def self.hasher_key
      key = ENV.fetch("OAUTH_STATE_SIGNING_KEY") { DEFAULT_STATE_SIGNING_KEY }
      # puts "key: #{key}"
      [key].pack('H*')
    end
  end
end

if __FILE__ == $PROGRAM_NAME
  #   DEFAULT_STATE_SIGNING_KEY = '1f143aa0518119ff7b4510fb9685c6c9e983ad5552e1b2793ddc4910d4d611fb'
  #   n = StateEncryptionAlgo::StripeOAuthState.create
  #   n.data = JSON.parse('{
  #   "primary_stripe_account_id": null,
  #   "salesforce_instance_subdomain": "nosoftware-saas-4417-dev-ed",
  #   "salesforce_instance_type": "SCRATCH_ORG",
  #   "salesforce_namespace": "c",
  #   "csrf": "9189b1f1ce51e68672f5b5f4dd51c9861f1e5d261fa3baafc34cc93a92be3adf",
  #   "org_id": "00D7d000009NqzkEAC",
  #   "oauth_version": "v2",
  #   "user_id": 7,
  #   "stripe_account_id": "acct_15uapDIsgf92XbAO"
  # }')
  # puts n.data
  # list = []
  # (0..10).each do |i|
  #   n.data["user_id"] = i
  #   list[i] = n.to_s
  #   puts list[i]
  # end
  # puts ""
  # puts ""
  # puts ""
  # puts list.to_json
  # puts ""
  # puts ""
  # puts ""
  #
  # token = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzYWxlc2ZvcmNlX2luc3RhbmNlX3R5cGUiOiJTQ1JBVENIX09SRyIsInNhbGVzZm9yY2VfbmFtZXNwYWNlIjoiYyIsImNzcmYiOiI3YmQ5NDYwNjY5OGNmZDVjZjY0M2YxZjMxNWUyMGM4MDI2OWQ5MTM1Y2MzNmQ1NDI1NzI0MTQ4YzE5YjQ1MmYxIiwib3JnX2lkIjoiMDBERFMwMDAwMDFJR0QzMkFPIn0.4d9eRiMLhVcpqEqs-5Kt-AyQdWs0oeQCtYSJTSF6f2A'
  # token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcmltYXJ5X3N0cmlwZV9hY2NvdW50X2lkIjpudWxsLCJzYWxlc2ZvcmNlX2luc3RhbmNlX3N1YmRvbWFpbiI6Im5vc29mdHdhcmUtc2Fhcy00NDE3LWRldi1lZCIsInNhbGVzZm9yY2VfaW5zdGFuY2VfdHlwZSI6IlNDUkFUQ0hfT1JHIiwic2FsZXNmb3JjZV9uYW1lc3BhY2UiOiJjIiwiY3NyZiI6IjkxNDdmMDcyNTRmZTY3Y2M4YzYzMTQ4YjBlNGQ4ZDdhZWNmM2NjNGYyMjRjZjA4MzFjMjFmODAwOWM2NzEzNjUiLCJvcmdfaWQiOiIwMEQ3ZDAwMDAwOU5xemtFQUMiLCJvYXV0aF92ZXJzaW9uIjoidjIiLCJ1c2VyX2lkIjo3LCJzdHJpcGVfYWNjb3VudF9pZCI6ImFjY3RfMTV1YXBESXNnZjkyWGJBTyJ9.KZHp8DOApyZe27ii0bufUs-5Kr5EsgPfrmCN830OY2Y'
  # token = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdHJpcGVfYWNjb3VudF9pZCI6ImFjY3RfMTV1YXBESXNnZjkyWGJBTyIsInVzZXJfaWQiOjcsIm9hdXRoX3ZlcnNpb24iOiJ2MiIsIm9yZ19pZCI6IjAwRDdkMDAwMDA5TnF6a0VBQyIsImNzcmYiOiI5MTQ3ZjA3MjU0ZmU2N2NjOGM2MzE0OGIwZTRkOGQ3YWVjZjNjYzRmMjI0Y2YwODMxYzIxZjgwMDljNjcxMzY1Iiwic2FsZXNmb3JjZV9uYW1lc3BhY2UiOiJjIiwic2FsZXNmb3JjZV9pbnN0YW5jZV90eXBlIjoiU0NSQVRDSF9PUkciLCJzYWxlc2ZvcmNlX2luc3RhbmNlX3N1YmRvbWFpbiI6Im5vc29mdHdhcmUtc2Fhcy00NDE3LWRldi1lZCIsInByaW1hcnlfc3RyaXBlX2FjY291bnRfaWQiOm51bGx9.6JWXiw6T3YCQfmdfKYM5t0CUGEO6hGoxzmiy4RshBHo'
  # s = StateEncryptionAlgo::StripeOAuthState.from_encrypted_state!(token)
  # puts s.data unless s.nil?
  # puts s.to_s unless s.nil?
  # puts token

  # unless s.nil?
  #   s.salesforce_instance_type = 'SANDBOX'
  #   s.salesforce_namespace = 'c'
  #
  #   puts "v2 with no user id: #{s}"
  #   s.user_id = 4
  #   puts "v2 with user id: #{s}"
  #
  #   puts s.salesforce_account_id
  #   puts s.csrf
  #   puts s.user_id
  #   puts s.oauth_version
  #   puts s.data.to_json
  #
  #   # s.user_id = 1234
  #   puts s.user_id
  #   puts s.to_s
  # end
  #
  # n = StateEncryptionAlgo::StripeOAuthState.create
  # unless n.nil?
  #   puts "csrf: #{n.csrf}"
  #   puts n.data
  #   puts "v1 with no info set: #{n}"
  # end
end
