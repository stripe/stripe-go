# frozen_string_literal: true
# typed: false
module KMSEncryptionTestHelpers
    @encryption_fields_mocked = false

    def self.mock_encryption_fields(klass)
      return if @encryption_fields_mocked

      klass.class_eval do
        self.kms_encrypted_fields.each do |field|
          define_method(field) do
            self.send("encrypted_#{field}")
          end

          define_method("#{field}=") do |password|
            self.send("encrypted_#{field}=", password)
          end
        end
      end

      @encryption_fields_mocked = true
    end

    def self.restore_encryption_fields(klass)
      return if !@encryption_fields_mocked

      klass.class_eval do
        self.kms_encrypted_fields.each do |field|
          self.kms_encrypted_field(field)
        end
      end

      @encryption_fields_mocked = false
    end

end
