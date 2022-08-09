# typed: true
# Description: autogen sorbet typed for sequel models
# Usage: bundle exec ruby scripts/sequel-types.rb

require File.expand_path('../config/environment', __dir__)

# https://github.com/chanzuckerberg/sorbet-rails/blob/master/lib/sorbet-rails/model_plugins/active_record_attribute.rb
class SequelSorbetPlugin
  extend T::Sig

  class ColumnType < T::Struct
    extend T::Sig

    # const :base_type, T.any(T.class_of(Class), String)
    const :base_type, T.untyped
    const :nilable, T.nilable(T::Boolean)
    const :array_type, T.nilable(T::Boolean)
    const :json_type, T.nilable(T::Boolean)

    sig { returns(String) }
    def to_s
      type = base_type.to_s
      # A nullable array column should be T.nilable(T::Array[column_type]) not T::Array[T.nilable(column_type)]
      type = "T::Array[#{type}]" if array_type
      type = "T.any(T::Array[T.untyped], T::Hash[String, T.untyped])" if json_type
      type = "T.nilable(#{type})" if nilable
      type
    end
  end

  def initialize(model_class)
    @model_class = model_class
  end

  sig { params(root: Parlour::RbiGenerator::Namespace).void }
  def generate(root)
    return if !defined?(Sequel::Model) && @model_class.superclass != Sequel::Model

    attribute_module_name = "#{@model_class}::GeneratedAttributeMethods"
    attribute_module_rbi = root.create_module(attribute_module_name)

    model_class_rbi = root.create_class(@model_class.to_s, superclass: Sequel::Model.to_s)
    model_class_rbi.create_include(attribute_module_name)

    table_schema = @model_class.db.schema(@model_class.table_name)

    model_class_rbi.create_method(
      "self.find",
      parameters: [
        Parlour::RbiGenerator::Parameter.new("value", type: "T::Hash[T.untyped, T.untyped]")
      ],
      return_type: "T.nilable(#{@model_class})"
    )

    model_class_rbi.create_method(
      "self.[]",
      parameters: [
        Parlour::RbiGenerator::Parameter.new("value", type: "Integer")
      ],
      return_type: "T.nilable(#{@model_class})"
    )

    model_class_rbi.create_method(
      "self.first",
      return_type: "T.nilable(#{@model_class})"
    )

    model_class_rbi.create_method(
      "self.last",
      return_type: "T.nilable(#{@model_class})"
    )

    @model_class.columns.sort.each do |column_name|
      # => {:oid=>23, :db_type=>"integer", :default=>nil, :allow_null=>false, :primary_key=>true, :type=>:integer, :auto_increment=>true, :ruby_default=>nil}
      column_schema = table_schema.detect {|c| c[0] == column_name }[1]
      ruby_column_type = sequel_to_ruby_type(column_schema[:type] || column_schema[:db_type])

      column_type = ColumnType.new(
        base_type: ruby_column_type,
        nilable: column_schema[:allow_null],
        array_type: false,
        json_type: column_schema[:db_type] == "jsonb"
      )

      attribute_module_rbi.create_method(
        column_name.to_s,
        return_type: column_type.to_s,
      )

      attribute_module_rbi.create_method(
        "#{column_name}=",
        parameters: [
          Parlour::RbiGenerator::Parameter.new("value", type: value_type_for_attr_writer(column_type)),
        ],
        return_type: nil,
      )
    end

    if @model_class.included_modules.include?(KMSEncryption)
      @model_class.kms_encrypted_fields.sort.each do |column_name|
          attribute_module_rbi.create_method(
            column_name.to_s,
            return_type: 'T.nilable(String)',
          )

          attribute_module_rbi.create_method(
            "#{column_name}=",
            parameters: [
              Parlour::RbiGenerator::Parameter.new("value", type: 'String'),
            ],
            return_type: nil,
          )
      end
    end
  end

  # sig { params(sequel_type: Symbol).returns(Class) }
  def sequel_to_ruby_type(sequel_type)
    case sequel_type
    when :integer
      Integer
    when :string
      String
    when :boolean
      T::Boolean
    when :datetime
      DateTime
    when 'jsonb'
      nil
    else
      raise 'unsupported type'
    end
  end

  sig { params(column_type: ColumnType).returns(String) }
  def value_type_for_attr_writer(column_type)
    assignable_time_supertypes = [DateTime, Date, Time].map(&:to_s)

    type = column_type.base_type

    if type.is_a?(Class)
      if type == DateTime
        type = "T.any(#{assignable_time_supertypes.join(', ')})"
      elsif type < Numeric
        type = "T.any(Numeric)"
      elsif type == String
        type = "T.any(String, Symbol)"
      end
    end

    ColumnType.new(
      base_type: type,
      nilable: column_type.nilable,
      array_type: column_type.array_type,
      json_type: column_type.json_type
    ).to_s
  end
end

# https://github.com/chanzuckerberg/sorbet-rails/blob/8337c1cea41490b0c21fb67153129912eeb3504a/lib/sorbet-rails/model_rbi_formatter.rb
# https://github.com/chanzuckerberg/sorbet-rails/blob/master/lib/sorbet-rails/tasks/rails_rbi.rake

Sequel::Model.descendants.each do |sequel_model|
  puts "Generating #{sequel_model}..."

  generator = Parlour::RbiGenerator.new(break_params: 3)

  SequelSorbetPlugin.new(sequel_model).generate(generator.root)

  rbi = <<~MESSAGE
    # This is an autogenerated file for dynamic methods in #{sequel_model}
    # Please rerun bundle exec scripts/sequel-types.rb to regenerate
  MESSAGE

  rbi += generator.rbi

  file_path = Rails.root.join(
    "sorbet",
    "rails-rbi",
    "sequel",
    "#{sequel_model.name.underscore}.rbi",
  )

  FileUtils.mkdir_p(File.dirname(file_path))
  File.write(file_path, rbi)
end
