# frozen_string_literal: true
# typed: false
# TODO consider extracting into a separate gem?

module KMSEncryption
    # define `kms_encryption_context(field = nil)` and `kms_encryption_key(field = nil)` on the model

    # rubocop:disable Lint/AssignmentInCondition
    def self.included(klass)
      klass.send(:extend, ClassMethods)
      klass.send(:include, InstanceMethods)
    end

    module ClassMethods
      def kms_encrypted_fields
        @kms_encrypted_fields ||= []
      end

      def kms_encrypted_field(field)
        if !kms_encrypted_fields.include?(field)
          kms_encrypted_fields << field
        end

        attr_reader field

        define_method(field.to_s) do
          kms_get_field(field)
        end

        define_method("#{field}=") do |value|
          kms_set_field(field, value)
        end
      end
    end

    module InstanceMethods

      # NOTE this is a sequel-specific binding which causes KMS to be called syncronously
      def before_save
        self.kms_encrypt_updated_fields
        super
      end

      def kms_encrypted_fields_changed?
        self.class.kms_encrypted_fields {|f| kms_get_unencrypted_value(f) }.compact.any?
      end

      def kms_get_unencrypted_value(field)
        instance_variable_get("@unencrypted_#{field}")
      end

      protected def kms_set_field(field, value)
        self.instance_variable_set("@unencrypted_#{field}", value)
        self.instance_variable_set("@#{field}", nil)
      end

      protected def kms_get_field(field)
        # look for locally set unencrypted value first
        if unencrypted_value = kms_get_unencrypted_value(field)
          return unencrypted_value
        end

        if cached_unencrypted_value = instance_variable_get("@#{field}")
          return cached_unencrypted_value
        end

        # if there is no unencrypted value
        return if self.send("encrypted_#{field}").nil?

        decoded_encrypted_value = Base64.decode64(self.send("encrypted_#{field}"))

        unencrypted_value = self.kms_client.decrypt({
          ciphertext_blob: decoded_encrypted_value,
          encryption_context: self.kms_encryption_context(field),
        }).plaintext

        instance_variable_set("@#{field}", unencrypted_value)

        unencrypted_value
      end

      protected def kms_encrypt_updated_fields
        self.class.kms_encrypted_fields.each do |field|
          if unencrypted_value = kms_get_unencrypted_value(field)
            # skip encrypting nil values
            if unencrypted_value.nil? || unencrypted_value.empty?
              self.send("encrypted_#{field}=", nil)
              next
            end

            kms_response = self.kms_client.encrypt({
              key_id: self.kms_encryption_key(field),
              plaintext: unencrypted_value,
              encryption_context: self.kms_encryption_context(field),
            })

            encoded_encrypted_field = Base64.encode64(kms_response.ciphertext_blob.to_s)

            self.send(:"encrypted_#{field}=", encoded_encrypted_field)
          end
        end
      end

      protected def kms_client
        @kms_client ||= Aws::KMS::Client.new(
          region: ENV['AWS_REGION'],
          access_key_id: ENV['AWS_KEY'],
          secret_access_key: ENV['AWS_SECRET_KEY'],
          retry_limit: 6
        )
      end
    end
  # rubocop:enable Lint/AssignmentInCondition
end
