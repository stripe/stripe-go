//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The data type of the field.
type V2ExtendObjectDefinitionVersionPropertiesItemsType string

// List of values that V2ExtendObjectDefinitionVersionPropertiesItemsType can take
const (
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeArray           V2ExtendObjectDefinitionVersionPropertiesItemsType = "array"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeBoolean         V2ExtendObjectDefinitionVersionPropertiesItemsType = "boolean"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeCurrency        V2ExtendObjectDefinitionVersionPropertiesItemsType = "currency"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeDatetime        V2ExtendObjectDefinitionVersionPropertiesItemsType = "datetime"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeDecimal         V2ExtendObjectDefinitionVersionPropertiesItemsType = "decimal"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeEnum            V2ExtendObjectDefinitionVersionPropertiesItemsType = "enum"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeInteger         V2ExtendObjectDefinitionVersionPropertiesItemsType = "integer"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeMonetaryAmount  V2ExtendObjectDefinitionVersionPropertiesItemsType = "monetary_amount"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeObjectReference V2ExtendObjectDefinitionVersionPropertiesItemsType = "object_reference"
	V2ExtendObjectDefinitionVersionPropertiesItemsTypeString          V2ExtendObjectDefinitionVersionPropertiesItemsType = "string"
)

// The data type of the field.
type V2ExtendObjectDefinitionVersionPropertiesType string

// List of values that V2ExtendObjectDefinitionVersionPropertiesType can take
const (
	V2ExtendObjectDefinitionVersionPropertiesTypeArray           V2ExtendObjectDefinitionVersionPropertiesType = "array"
	V2ExtendObjectDefinitionVersionPropertiesTypeBoolean         V2ExtendObjectDefinitionVersionPropertiesType = "boolean"
	V2ExtendObjectDefinitionVersionPropertiesTypeCurrency        V2ExtendObjectDefinitionVersionPropertiesType = "currency"
	V2ExtendObjectDefinitionVersionPropertiesTypeDatetime        V2ExtendObjectDefinitionVersionPropertiesType = "datetime"
	V2ExtendObjectDefinitionVersionPropertiesTypeDecimal         V2ExtendObjectDefinitionVersionPropertiesType = "decimal"
	V2ExtendObjectDefinitionVersionPropertiesTypeEnum            V2ExtendObjectDefinitionVersionPropertiesType = "enum"
	V2ExtendObjectDefinitionVersionPropertiesTypeInteger         V2ExtendObjectDefinitionVersionPropertiesType = "integer"
	V2ExtendObjectDefinitionVersionPropertiesTypeMonetaryAmount  V2ExtendObjectDefinitionVersionPropertiesType = "monetary_amount"
	V2ExtendObjectDefinitionVersionPropertiesTypeObjectReference V2ExtendObjectDefinitionVersionPropertiesType = "object_reference"
	V2ExtendObjectDefinitionVersionPropertiesTypeString          V2ExtendObjectDefinitionVersionPropertiesType = "string"
)

// The stable identity and API names of the parent object definition.
type V2ExtendObjectDefinitionVersionParentDetails struct {
	// The plural API name used in this definition's object record API paths.
	APINamePlural string `json:"api_name_plural"`
	// The singular API name used in this definition's object record tags.
	APINameSingular string `json:"api_name_singular"`
	// The ID of the parent object definition.
	ID string `json:"id"`
}

// The required scalar schema for each array item. Array values retain their supplied order, and
// order is significant. Values are never sorted. Empty arrays are valid unless min_items is positive.
type V2ExtendObjectDefinitionVersionPropertiesItemsItems struct{}

// The required scalar schema for each array item. Array values retain their supplied order, and
// order is significant. Values are never sorted. Empty arrays are valid unless min_items is positive.
type V2ExtendObjectDefinitionVersionPropertiesItems struct {
	// A required list of values for enum types.
	AllowedValues []string `json:"allowed_values,omitempty"`
	// An optional default value for the field.
	DefaultValue any `json:"default_value,omitempty"`
	// An optional human-readable description of the field.
	Description string `json:"description,omitempty"`
	// An optional exclusive maximum value for integer or decimal types.
	ExclusiveMaximum any `json:"exclusive_maximum,omitempty"`
	// An optional exclusive minimum value for integer or decimal types.
	ExclusiveMinimum any `json:"exclusive_minimum,omitempty"`
	// Whether or not the field is deprecated (default false).
	IsDeprecated bool `json:"is_deprecated,omitempty"`
	// The required scalar schema for each array item. Array values retain their supplied order, and
	// order is significant. Values are never sorted. Empty arrays are valid unless min_items is positive.
	Items *V2ExtendObjectDefinitionVersionPropertiesItemsItems `json:"items,omitempty"`
	// An optional maximum value for integer or decimal types.
	Maximum any `json:"maximum,omitempty"`
	// An optional maximum number of array items. It defaults to and cannot exceed 100.
	MaxItems int64 `json:"max_items,omitempty"`
	// An optional maximum length validation for string fields.
	MaxLength int64 `json:"max_length,omitempty"`
	// An optional minimum value for integer or decimal types.
	Minimum any `json:"minimum,omitempty"`
	// An optional minimum number of array items. Empty arrays are valid when this is absent or zero.
	MinItems int64 `json:"min_items,omitempty"`
	// An optional minimum length validation for string fields.
	MinLength int64 `json:"min_length,omitempty"`
	// The type referenced by object reference fields.
	ReferencedType string `json:"referenced_type,omitempty"`
	// Whether or not the field is required (default false).
	Required bool `json:"required,omitempty"`
	// An optional human-readable name for the field.
	Title string `json:"title,omitempty"`
	// The data type of the field.
	Type V2ExtendObjectDefinitionVersionPropertiesItemsType `json:"type"`
	// Whether array item values must be unique (default false). Duplicate values are allowed by default.
	UniqueItems bool `json:"unique_items,omitempty"`
}

// The schema of the object record fields defined by this object definition.
type V2ExtendObjectDefinitionVersionProperties struct {
	// A required list of values for enum types.
	AllowedValues []string `json:"allowed_values,omitempty"`
	// An optional default value for the field.
	DefaultValue any `json:"default_value,omitempty"`
	// An optional human-readable description of the field.
	Description string `json:"description,omitempty"`
	// An optional exclusive maximum value for integer or decimal types.
	ExclusiveMaximum any `json:"exclusive_maximum,omitempty"`
	// An optional exclusive minimum value for integer or decimal types.
	ExclusiveMinimum any `json:"exclusive_minimum,omitempty"`
	// Whether or not the field is deprecated (default false).
	IsDeprecated bool `json:"is_deprecated,omitempty"`
	// The required scalar schema for each array item. Array values retain their supplied order, and
	// order is significant. Values are never sorted. Empty arrays are valid unless min_items is positive.
	Items *V2ExtendObjectDefinitionVersionPropertiesItems `json:"items,omitempty"`
	// An optional maximum value for integer or decimal types.
	Maximum any `json:"maximum,omitempty"`
	// An optional maximum number of array items. It defaults to and cannot exceed 100.
	MaxItems int64 `json:"max_items,omitempty"`
	// An optional maximum length validation for string fields.
	MaxLength int64 `json:"max_length,omitempty"`
	// An optional minimum value for integer or decimal types.
	Minimum any `json:"minimum,omitempty"`
	// An optional minimum number of array items. Empty arrays are valid when this is absent or zero.
	MinItems int64 `json:"min_items,omitempty"`
	// An optional minimum length validation for string fields.
	MinLength int64 `json:"min_length,omitempty"`
	// The type referenced by object reference fields.
	ReferencedType string `json:"referenced_type,omitempty"`
	// Whether or not the field is required (default false).
	Required bool `json:"required,omitempty"`
	// An optional human-readable name for the field.
	Title string `json:"title,omitempty"`
	// The data type of the field.
	Type V2ExtendObjectDefinitionVersionPropertiesType `json:"type"`
	// Whether array item values must be unique (default false). Duplicate values are allowed by default.
	UniqueItems bool `json:"unique_items,omitempty"`
}

// Represents an immutable version of an object definition.
// It defines the fields of the object records created from it.
type V2ExtendObjectDefinitionVersion struct {
	APIResource
	// Time at which the object definition version was created.
	Created time.Time `json:"created"`
	// A description of what this version of the object definition represents.
	Description string `json:"description,omitempty"`
	// The human-readable plural name of this version.
	DisplayNamePlural string `json:"display_name_plural"`
	// The human-readable singular name of this version.
	DisplayNameSingular string `json:"display_name_singular"`
	// The complete version ID, consisting of the object definition ID, `@`, and a semantic version.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the parent object definition.
	ObjectDefinition string `json:"object_definition"`
	// The stable identity and API names of the parent object definition.
	ParentDetails *V2ExtendObjectDefinitionVersionParentDetails `json:"parent_details"`
	// The schema of the object record fields defined by this object definition.
	Properties map[string]*V2ExtendObjectDefinitionVersionProperties `json:"properties"`
	// The semantic version of this object definition version, in `major.minor.patch` format.
	Version string `json:"version"`
}
