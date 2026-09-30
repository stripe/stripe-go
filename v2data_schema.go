//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The data type of the column.
type V2DataSchemaColumnType string

// List of values that V2DataSchemaColumnType can take
const (
	V2DataSchemaColumnTypeBigint    V2DataSchemaColumnType = "bigint"
	V2DataSchemaColumnTypeBoolean   V2DataSchemaColumnType = "boolean"
	V2DataSchemaColumnTypeDate      V2DataSchemaColumnType = "date"
	V2DataSchemaColumnTypeDatetime  V2DataSchemaColumnType = "datetime"
	V2DataSchemaColumnTypeDecimal   V2DataSchemaColumnType = "decimal"
	V2DataSchemaColumnTypeDouble    V2DataSchemaColumnType = "double"
	V2DataSchemaColumnTypeInteger   V2DataSchemaColumnType = "integer"
	V2DataSchemaColumnTypeTimestamp V2DataSchemaColumnType = "timestamp"
	V2DataSchemaColumnTypeVarchar   V2DataSchemaColumnType = "varchar"
)

// The dataset the table belongs to.
type V2DataSchemaDataset string

// List of values that V2DataSchemaDataset can take
const (
	V2DataSchemaDatasetAnalytical V2DataSchemaDataset = "analytical"
)

// Columns in other schemas that reference this column as a foreign key.
type V2DataSchemaColumnForeignKeysFrom struct {
	// The name of the referenced column.
	Column string `json:"column"`
	// The identifier of the referenced schema.
	Schema string `json:"schema"`
}

// Columns in other schemas that this column references as a foreign key.
type V2DataSchemaColumnForeignKeysTo struct {
	// The name of the referenced column.
	Column string `json:"column"`
	// The identifier of the referenced schema.
	Schema string `json:"schema"`
}

// The columns of the table.
type V2DataSchemaColumn struct {
	// A description of what the column represents.
	Description string `json:"description"`
	// Columns in other schemas that reference this column as a foreign key.
	ForeignKeysFrom []*V2DataSchemaColumnForeignKeysFrom `json:"foreign_keys_from"`
	// Columns in other schemas that this column references as a foreign key.
	ForeignKeysTo []*V2DataSchemaColumnForeignKeysTo `json:"foreign_keys_to"`
	// Whether the column forms part of the table's primary key.
	IsPrimaryKey bool `json:"is_primary_key"`
	// The name of the column.
	Name string `json:"name"`
	// The data type of the column.
	Type V2DataSchemaColumnType `json:"type"`
}

// Reports relevant to this table.
type V2DataSchemaRelevantReport struct {
	// A description of the `Report`.
	Description string `json:"description"`
	// The unique identifier of the `Report`.
	ID string `json:"id"`
	// The human-readable name of the `Report`.
	Name string `json:"name"`
}

// The `Schema` resource describes the columns, types, and relationships of a table that
// can be queried.
type V2DataSchema struct {
	APIResource
	// The columns of the table.
	Columns []*V2DataSchemaColumn `json:"columns"`
	// The dataset the table belongs to.
	Dataset V2DataSchemaDataset `json:"dataset"`
	// A description of the table.
	Description string `json:"description"`
	// An extended, LLM-friendly description of the table, useful for query generation.
	ExtendedDescription string `json:"extended_description,omitempty"`
	// The unique identifier of the `Schema`.
	ID string `json:"id"`
	// Whether this `Schema` describes live mode data.
	Livemode bool `json:"livemode"`
	// The human-readable name of the table.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Time at which the table's schema was last refreshed.
	RefreshedAt time.Time `json:"refreshed_at"`
	// Reports relevant to this table.
	RelevantReports []*V2DataSchemaRelevantReport `json:"relevant_reports"`
}
