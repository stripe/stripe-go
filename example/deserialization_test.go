package example

import (
	"encoding/json"
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	assert "github.com/stretchr/testify/require"
)

type APIResource = stripe.APIResource
type ListMeta = stripe.ListMeta

var parseID = stripe.ParseID

func ParseID(data []byte) (string, bool) {
	return string(data), false
}

func TestStubWithIdDirect(t *testing.T) {
	input := []byte("{\"some_ref\":{\"id\":\"xyz\"}}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "xyz", resource.SomeRef.ID)
}

func TestStubWithIdExpanded(t *testing.T) {
	input := []byte("{\"some_expandable\":{\"id\":\"xyz\"}}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "xyz", resource.SomeExpandable.ID)
}

func TestStubWithIdArrayExpanded(t *testing.T) {
	input := []byte("{\"some_expanded_array\":[{\"id\":\"xyz\"}]}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "xyz", resource.SomeExpandedArray[0].ID)
}

func TestStubWithIdInArray(t *testing.T) {
	input := []byte("{\"some_ref_array\":[{\"id\":\"xyz\"}]}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "xyz", resource.SomeRefArray[0].ID)
}

func TestStubWithIdInListObject(t *testing.T) {
	input := []byte("{\"some_list_object\":{\"data\":[{\"id\":\"xyz\"}]}}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "xyz", resource.SomeListObject.Data[0].ID)
}

func TestEmptyObjectDirect(t *testing.T) {
	input := []byte("{\"some_ref\":{}}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", resource.SomeRef.ID)
}

func TestEmptyObjectArray(t *testing.T) {
	input := []byte("{\"some_ref_array\":[{}]}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", resource.SomeRefArray[0].ID)
}

func TestEmptyObjectExpanded(t *testing.T) {
	input := []byte("{\"some_expandable\":{}}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", resource.SomeExpandable.ID)
}

func TestEmptyObjectArrayExpanded(t *testing.T) {
	input := []byte("{\"some_expanded_array\":[{}]}")
	var resource MyResource
	err := json.Unmarshal(input, &resource)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", resource.SomeExpandedArray[0].ID)
}

type MyResourceSomeLiteral string

// List of values that MyResourceSomeLiteral can take
const (
	MyResourceSomeLiteralLiteral MyResourceSomeLiteral = "literal"
)

type MyResourceSomeEnum string

// List of values that MyResourceSomeEnum can take
const (
	MyResourceSomeEnumFoo MyResourceSomeEnum = "foo"
	MyResourceSomeEnumBar MyResourceSomeEnum = "bar"
)

type MyResourceSomePolymorphicGroupType string

// List of values that MyResourceSomePolymorphicGroupType can take
const (
	MyResourceSomePolymorphicGroupTypeMyResource           MyResourceSomePolymorphicGroupType = "MyResource"
	MyResourceSomePolymorphicGroupTypeMyResourceSomeObject MyResourceSomePolymorphicGroupType = "MyResourceSomeObject"
)

type MyResourceSomeObject struct {
	SomeString string `json:"some_string"`
}
type MyResource struct {
	ID                   string                         `json:"id"`
	SomeBoolean          bool                           `json:"some_boolean"`
	SomeDatetime         int64                          `json:"some_datetime"`
	SomeDecimalString    float64                        `json:"some_decimal_string,string"`
	SomeEnum             MyResourceSomeEnum             `json:"some_enum"`
	SomeExpandable       *MyResource                    `json:"some_expandable"`
	SomeExpandedArray    []*MyResource                  `json:"some_expanded_array"`
	SomeInteger          int64                          `json:"some_integer"`
	SomeListObject       *MyResourceList                `json:"some_list_object"`
	SomeLiteral          MyResourceSomeLiteral          `json:"some_literal"`
	SomeLonginteger      int64                          `json:"some_longinteger"`
	SomeMap              map[string]string              `json:"some_map"`
	SomeNullable         string                         `json:"some_nullable"`
	SomeNumber           float64                        `json:"some_number"`
	SomeObject           *MyResourceSomeObject          `json:"some_object"`
	SomePolymorphicGroup MyResourceSomePolymorphicGroup `json:"some_polymorphic_group"`
	SomeRef              *MyResource                    `json:"some_ref"`
	SomeRefArray         []*MyResource                  `json:"some_ref_array"`
	SomeString           string                         `json:"some_string"`
	SomeStringArray      []string                       `json:"some_string_array"`
}
type MyResourceSomePolymorphicGroup struct {
	ID   string                             `json:"id"`
	Type MyResourceSomePolymorphicGroupType `json:"object"`

	MyResource           *MyResource           `json:"-"`
	MyResourceSomeObject *MyResourceSomeObject `json:"-"`
}

// UnmarshalJSON handles deserialization of a MyResourceSomePolymorphicGroup.
// This custom unmarshaling is needed because the specific type of
// MyResourceSomePolymorphicGroup it refers to is specified in the JSON
func (m *MyResourceSomePolymorphicGroup) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		m.ID = id
		return nil
	}

	type myResourceSomePolymorphicGroup MyResourceSomePolymorphicGroup
	var v myResourceSomePolymorphicGroup
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*m = MyResourceSomePolymorphicGroup(v)
	var err error

	switch m.Type {
	case MyResourceSomePolymorphicGroupTypeMyResource:
		err = json.Unmarshal(data, &m.MyResource)
	case MyResourceSomePolymorphicGroupTypeMyResourceSomeObject:
		err = json.Unmarshal(data, &m.MyResourceSomeObject)
	}
	return err
}

// MyResourceList is a list of MyResources as retrieved from a list endpoint.
type MyResourceList struct {
	APIResource
	ListMeta
	Data []*MyResource `json:"data"`
}
