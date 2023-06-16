
package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)
func TestStubWithIdDirect(t *testing.T) {
        input := []byte("{\"direct\":{\"id\":\"xyz\"}}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "xyz", resource.Direct.ID)
      }

func TestStubWithIdExpanded(t *testing.T) {
        input := []byte("{\"expanded\":{\"id\":\"xyz\"}}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "xyz", resource.Expanded.ID)
      }

func TestStubWithIdArrayExpanded(t *testing.T) {
        input := []byte("{\"array_expanded\":[{\"id\":\"xyz\"}]}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "xyz", resource.ArrayExpanded[0].ID)
      }

func TestStubWithIdInArray(t *testing.T) {
        input := []byte("{\"in_array\":[{\"id\":\"xyz\"}]}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "xyz", resource.InArray[0].ID)
      }

func TestStubWithIdInListObject(t *testing.T) {
        input := []byte("{\"in_list_object\":{\"data\":[{\"id\":\"xyz\"}]}}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "xyz", resource.InListObject.Data[0].ID)
      }

func TestEmptyObjectDirect(t *testing.T) {
        input := []byte("{\"direct\":{}}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "", resource.Direct.ID)
      }

func TestEmptyObjectArray(t *testing.T) {
        input := []byte("{\"in_array\":[{}]}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "", resource.InArray[0].ID)
      }

func TestEmptyObjectExpanded(t *testing.T) {
        input := []byte("{\"expanded\":{}}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "", resource.Expanded.ID)
      }

func TestEmptyObjectArrayExpanded(t *testing.T) {
        input := []byte("{\"array_expanded\":[{}]}")
        var resource MyTopLevelResource
        err := json.Unmarshal(input, &resource)
        assert.Equal(t, nil, err)
        assert.Equal(t, "", resource.ArrayExpanded[0].ID)
      }

type MyReferencedResource struct {
  ID string `json:"id"`
  SomeInt int64 `json:"some_int"`
}
// MyReferencedResourceList is a list of MyReferencedResources as retrieved from a list endpoint.
type MyReferencedResourceList struct {
  APIResource
  ListMeta
  Data []*MyReferencedResource `json:"data"`
}

type MyTopLevelResource struct {
  ArrayExpanded []*MyReferencedResource `json:"array_expanded"`
  Direct *MyReferencedResource `json:"direct"`
  Expanded *MyReferencedResource `json:"expanded"`
  InArray []*MyReferencedResource `json:"in_array"`
  InListObject *MyReferencedResourceList `json:"in_list_object"`
}
