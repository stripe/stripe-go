package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v86/form"
)

// colorModel is the discriminator type for the color union.
type colorModel string

const (
	colorModelRGB colorModel = "rgb"
	colorModelHSV colorModel = "hsv"
	colorModelHSL colorModel = "hsl"
)

// colorUnion is the struct-with-pointers pattern Go uses for discriminated
// unions on the response side. Only one variant field is populated at a time,
// based on the value of Model.
type colorUnion struct {
	Model colorModel  `json:"model"`
	RGB   *rgbVariant `json:"rgb_color,omitempty"`
	HSV   *hsvVariant `json:"hsv_color,omitempty"`
	HSL   *hslVariant `json:"hsl_color,omitempty"`
}

type rgbVariant struct {
	R int64 `json:"r"`
	G int64 `json:"g"`
	B int64 `json:"b"`
}

type hsvVariant struct {
	H int64 `json:"h"`
	S int64 `json:"s"`
	V int64 `json:"v"`
}

type hslVariant struct {
	H int64 `json:"h"`
	S int64 `json:"s"`
	L int64 `json:"l"`
}

// colorParams mirrors the form-encoded request side for a standalone
// discriminated union: a single flat struct with the discriminator field and
// all variant fields as optional pointers. The caller sets exactly one
// variant's fields alongside the discriminator.
type colorParams struct {
	Model *string `form:"model"`
	// RGB variant fields
	R *int64 `form:"r"`
	G *int64 `form:"g"`
	B *int64 `form:"b"`
	// HSV variant fields
	H *int64 `form:"h"`
	S *int64 `form:"s"`
	V *int64 `form:"v"`
}

// llamaParams demonstrates the inline union pattern: the discriminator lives
// at the parent level alongside base fields, and each variant brings its own
// nested struct.
type llamaParams struct {
	Params     `form:"*"`
	Name       *string           `form:"name"`
	Type       *string           `form:"type"`
	AlienLlama *alienLlamaParams `form:"alien_llama"`
	EarthLlama *earthLlamaParams `form:"earth_llama"`
	MagicLlama *magicLlamaParams `form:"magic_llama"`
}

type alienLlamaParams struct {
	Planet    *string `form:"planet"`
	Telepathic *bool  `form:"telepathic"`
}

type earthLlamaParams struct {
	Country *string `form:"country"`
}

type magicLlamaParams struct {
	Spell *string `form:"spell"`
}

// TestDiscriminatedUnion_RequestParams_RGBVariant verifies that a standalone
// discriminated union params struct encodes its discriminator and the chosen
// variant's fields, while omitting unset variant fields.
func TestDiscriminatedUnion_RequestParams_RGBVariant(t *testing.T) {
	model := "rgb"
	r := int64(255)
	g := int64(128)
	b := int64(0)

	params := &colorParams{
		Model: &model,
		R:     &r,
		G:     &g,
		B:     &b,
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	// Discriminator is encoded.
	assert.Equal(t, []string{"rgb"}, values["model"])

	// RGB variant fields are encoded.
	assert.Equal(t, []string{"255"}, values["r"])
	assert.Equal(t, []string{"128"}, values["g"])
	assert.Equal(t, []string{"0"}, values["b"])

	// HSV variant fields are absent (nil pointers omitted).
	assert.Nil(t, values["h"])
	assert.Nil(t, values["s"])
	assert.Nil(t, values["v"])
}

// TestDiscriminatedUnion_RequestParams_HSVVariant verifies that switching to
// the HSV variant encodes the correct discriminator and fields.
func TestDiscriminatedUnion_RequestParams_HSVVariant(t *testing.T) {
	model := "hsv"
	h := int64(180)
	s := int64(100)
	v := int64(50)

	params := &colorParams{
		Model: &model,
		H:     &h,
		S:     &s,
		V:     &v,
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"hsv"}, values["model"])
	assert.Equal(t, []string{"180"}, values["h"])
	assert.Equal(t, []string{"100"}, values["s"])
	assert.Equal(t, []string{"50"}, values["v"])

	// RGB variant fields are absent.
	assert.Nil(t, values["r"])
	assert.Nil(t, values["g"])
	assert.Nil(t, values["b"])
}

// TestDiscriminatedUnion_RequestParams_NilVariantFieldsOmitted verifies that
// when only a discriminator is set (no variant fields), only the discriminator
// appears in the encoded output.
func TestDiscriminatedUnion_RequestParams_NilVariantFieldsOmitted(t *testing.T) {
	model := "rgb"
	params := &colorParams{
		Model: &model,
		// All variant fields nil — none should appear.
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"rgb"}, values["model"])
	assert.Nil(t, values["r"])
	assert.Nil(t, values["g"])
	assert.Nil(t, values["b"])
	assert.Nil(t, values["h"])
	assert.Nil(t, values["s"])
	assert.Nil(t, values["v"])
	assert.Nil(t, values["l"])
}

// TestDiscriminatedUnion_RequestParams_UnsetDiscriminatorOmitted verifies that
// if neither discriminator nor variant fields are set, nothing is encoded.
func TestDiscriminatedUnion_RequestParams_UnsetDiscriminatorOmitted(t *testing.T) {
	params := &colorParams{}

	body := &form.Values{}
	form.AppendTo(body, params)

	assert.True(t, body.Empty(), "expected empty form values when all fields are nil")
}

// TestDiscriminatedUnion_InlineUnion_AlienLlama verifies that an inline union
// (discriminator at the parent level with variant structs as sibling fields)
// encodes the discriminator alongside the base fields and the chosen variant's
// nested struct.
func TestDiscriminatedUnion_InlineUnion_AlienLlama(t *testing.T) {
	llamaType := "alien_llama"
	name := "Cosmo"
	planet := "Mars"
	telepathic := true

	params := &llamaParams{
		Name: &name,
		Type: &llamaType,
		AlienLlama: &alienLlamaParams{
			Planet:     &planet,
			Telepathic: &telepathic,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	// Base field and discriminator are present.
	assert.Equal(t, []string{"Cosmo"}, values["name"])
	assert.Equal(t, []string{"alien_llama"}, values["type"])

	// Chosen variant's nested fields are encoded under the variant key.
	assert.Equal(t, []string{"Mars"}, values["alien_llama[planet]"])
	assert.Equal(t, []string{"true"}, values["alien_llama[telepathic]"])

	// Other variant structs are absent.
	assert.Nil(t, values["earth_llama[country]"])
	assert.Nil(t, values["magic_llama[spell]"])
}

// TestDiscriminatedUnion_InlineUnion_EarthLlama verifies the earth llama
// variant of the inline union.
func TestDiscriminatedUnion_InlineUnion_EarthLlama(t *testing.T) {
	llamaType := "earth_llama"
	name := "Llama Del Rey"
	country := "Peru"

	params := &llamaParams{
		Name: &name,
		Type: &llamaType,
		EarthLlama: &earthLlamaParams{
			Country: &country,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"Llama Del Rey"}, values["name"])
	assert.Equal(t, []string{"earth_llama"}, values["type"])
	assert.Equal(t, []string{"Peru"}, values["earth_llama[country]"])

	assert.Nil(t, values["alien_llama[planet]"])
	assert.Nil(t, values["magic_llama[spell]"])
}

// TestDiscriminatedUnion_Response_UnmarshalRGBVariant verifies that JSON for
// an RGB color union is unmarshaled with the discriminator and variant fields
// correctly populated, and other variant fields left nil.
func TestDiscriminatedUnion_Response_UnmarshalRGBVariant(t *testing.T) {
	data := []byte(`{"model":"rgb","rgb_color":{"r":255,"g":128,"b":0}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelRGB, c.Model)
	assert.NotNil(t, c.RGB)
	assert.Equal(t, int64(255), c.RGB.R)
	assert.Equal(t, int64(128), c.RGB.G)
	assert.Equal(t, int64(0), c.RGB.B)

	// Other variant fields remain nil.
	assert.Nil(t, c.HSV)
	assert.Nil(t, c.HSL)
}

// TestDiscriminatedUnion_Response_UnmarshalHSVVariant verifies the HSV
// discriminated union variant unmarshals correctly.
func TestDiscriminatedUnion_Response_UnmarshalHSVVariant(t *testing.T) {
	data := []byte(`{"model":"hsv","hsv_color":{"h":180,"s":100,"v":50}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelHSV, c.Model)
	assert.NotNil(t, c.HSV)
	assert.Equal(t, int64(180), c.HSV.H)
	assert.Equal(t, int64(100), c.HSV.S)
	assert.Equal(t, int64(50), c.HSV.V)

	assert.Nil(t, c.RGB)
	assert.Nil(t, c.HSL)
}

// TestDiscriminatedUnion_Response_UnmarshalHSLVariant verifies the HSL
// discriminated union variant unmarshals correctly.
func TestDiscriminatedUnion_Response_UnmarshalHSLVariant(t *testing.T) {
	data := []byte(`{"model":"hsl","hsl_color":{"h":240,"s":100,"l":50}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelHSL, c.Model)
	assert.NotNil(t, c.HSL)
	assert.Equal(t, int64(240), c.HSL.H)
	assert.Equal(t, int64(100), c.HSL.S)
	assert.Equal(t, int64(50), c.HSL.L)

	assert.Nil(t, c.RGB)
	assert.Nil(t, c.HSV)
}

// TestDiscriminatedUnion_Response_DiscriminatorOnly verifies that a JSON
// payload containing only the discriminator field (no variant data) deserializes
// without error and leaves all variant pointers nil.
func TestDiscriminatedUnion_Response_DiscriminatorOnly(t *testing.T) {
	data := []byte(`{"model":"rgb"}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelRGB, c.Model)
	assert.Nil(t, c.RGB)
	assert.Nil(t, c.HSV)
	assert.Nil(t, c.HSL)
}

// TestDiscriminatedUnion_Response_RoundTrip verifies that a discriminated union
// value can be marshaled to JSON and unmarshaled back without data loss.
func TestDiscriminatedUnion_Response_RoundTrip(t *testing.T) {
	original := colorUnion{
		Model: colorModelRGB,
		RGB: &rgbVariant{
			R: 100,
			G: 200,
			B: 50,
		},
	}

	data, err := json.Marshal(&original)
	assert.NoError(t, err)

	var roundTripped colorUnion
	err = json.Unmarshal(data, &roundTripped)
	assert.NoError(t, err)

	assert.Equal(t, original.Model, roundTripped.Model)
	assert.NotNil(t, roundTripped.RGB)
	assert.Equal(t, original.RGB.R, roundTripped.RGB.R)
	assert.Equal(t, original.RGB.G, roundTripped.RGB.G)
	assert.Equal(t, original.RGB.B, roundTripped.RGB.B)
	assert.Nil(t, roundTripped.HSV)
	assert.Nil(t, roundTripped.HSL)
}
