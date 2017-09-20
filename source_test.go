package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestSourceObjectParams_AppendTo(t *testing.T) {
	// Test to make sure that TypeData makes it to the root object level of
	// encoding
	{
		params := &SourceObjectParams{
			Type: "source_type",
			TypeData: map[string]string{
				"foo": "bar",
			},
		}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bar"}, body.Get("source_type[foo]"))
	}
}
