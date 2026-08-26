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
// unions on the response side: the discriminator plus one nullable pointer per
// variant. Nothing in the generated code or the decoder enforces that exactly
// one variant is set — the API guarantees it. See
// TestDiscriminatedUnion_Response_MultipleVariantsPopulate for what the decoder
// actually does when that guarantee is violated.
//
// Field names, json tags, and field ordering here are copied from real
// generated output — sdk-codegen's committed integration snapshot, type
// V2TestLlamaColor. Keeping them verbatim is what makes drift between these
// tests and the generator detectable; only the type names are shortened.
type colorUnion struct {
	Hsl   *hslVariant `json:"hsl,omitempty"`
	Hsv   *hsvVariant `json:"hsv,omitempty"`
	Model colorModel  `json:"model"`
	Rgb   *rgbVariant `json:"rgb,omitempty"`
}

type rgbVariant struct {
	B int64 `json:"b"`
	G int64 `json:"g"`
	R int64 `json:"r"`
}

type hsvVariant struct {
	H int64 `json:"h"`
	S int64 `json:"s"`
	V int64 `json:"v"`
}

type hslVariant struct {
	H int64 `json:"h"`
	L int64 `json:"l"`
	S int64 `json:"s"`
}

// colorParams mirrors the form-encoded request side for a standalone
// discriminated union: a wrapper struct with the discriminator field and one
// nullable pointer per variant. The caller sets exactly one variant pointer
// alongside the discriminator. Generated params carry both form and json tags
// (V2TestLlamaColorParams); only the form tags affect encoding here, but the
// json tags are kept so the fixture matches what codegen emits.
type colorParams struct {
	Hsl   *hslVariantParams `form:"hsl" json:"hsl,omitempty"`
	Hsv   *hsvVariantParams `form:"hsv" json:"hsv,omitempty"`
	Model *string           `form:"model" json:"model"`
	Rgb   *rgbVariantParams `form:"rgb" json:"rgb,omitempty"`
}

type rgbVariantParams struct {
	B *int64 `form:"b" json:"b"`
	G *int64 `form:"g" json:"g"`
	R *int64 `form:"r" json:"r"`
}

type hsvVariantParams struct {
	H *int64 `form:"h" json:"h"`
	S *int64 `form:"s" json:"s"`
	V *int64 `form:"v" json:"v"`
}

type hslVariantParams struct {
	H *int64 `form:"h" json:"h"`
	L *int64 `form:"l" json:"l"`
	S *int64 `form:"s" json:"s"`
}

// llamaParams demonstrates the inline union pattern: the discriminator lives
// at the parent level alongside base fields, and each variant brings its own
// nested struct. Mirrors generated V2TestLlamaCreateParams.
type llamaParams struct {
	Params     `form:"*"`
	AlienLlama *alienLlamaParams `form:"alien_llama" json:"alien_llama,omitempty"`
	EarthLlama *earthLlamaParams `form:"earth_llama" json:"earth_llama,omitempty"`
	MagicLlama *magicLlamaParams `form:"magic_llama" json:"magic_llama,omitempty"`
	Name       *string           `form:"name" json:"name"`
	Type       *string           `form:"type" json:"type"`
}

type alienLlamaParams struct {
	Planet     *string `form:"planet" json:"planet"`
	Telepathic *bool   `form:"telepathic" json:"telepathic"`
}

type earthLlamaParams struct {
	Country *string `form:"country" json:"country"`
}

type magicLlamaParams struct {
	Spell *string `form:"spell" json:"spell"`
}

// TestDiscriminatedUnion_RequestParams_RGBVariant verifies that a standalone
// discriminated union params struct encodes its discriminator and the chosen
// variant's nested fields, while omitting unset variant structs.
func TestDiscriminatedUnion_RequestParams_RGBVariant(t *testing.T) {
	model := "rgb"
	r := int64(255)
	g := int64(128)
	b := int64(0)

	params := &colorParams{
		Model: &model,
		Rgb: &rgbVariantParams{
			R: &r,
			G: &g,
			B: &b,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	// Discriminator is encoded.
	assert.Equal(t, []string{"rgb"}, values["model"])

	// RGB variant fields are encoded under the variant key.
	assert.Equal(t, []string{"255"}, values["rgb[r]"])
	assert.Equal(t, []string{"128"}, values["rgb[g]"])
	assert.Equal(t, []string{"0"}, values["rgb[b]"])

	// Other variant structs are absent.
	assert.Nil(t, values["hsv[h]"])
	assert.Nil(t, values["hsv[s]"])
	assert.Nil(t, values["hsv[v]"])
	assert.Nil(t, values["hsl[h]"])
	assert.Nil(t, values["hsl[s]"])
	assert.Nil(t, values["hsl[l]"])
}

// TestDiscriminatedUnion_RequestParams_HSVVariant verifies that switching to
// the HSV variant encodes the correct discriminator and nested fields.
func TestDiscriminatedUnion_RequestParams_HSVVariant(t *testing.T) {
	model := "hsv"
	h := int64(180)
	s := int64(100)
	v := int64(50)

	params := &colorParams{
		Model: &model,
		Hsv: &hsvVariantParams{
			H: &h,
			S: &s,
			V: &v,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"hsv"}, values["model"])
	assert.Equal(t, []string{"180"}, values["hsv[h]"])
	assert.Equal(t, []string{"100"}, values["hsv[s]"])
	assert.Equal(t, []string{"50"}, values["hsv[v]"])

	// RGB and HSL variant structs are absent.
	assert.Nil(t, values["rgb[r]"])
	assert.Nil(t, values["rgb[g]"])
	assert.Nil(t, values["rgb[b]"])
	assert.Nil(t, values["hsl[h]"])
	assert.Nil(t, values["hsl[s]"])
	assert.Nil(t, values["hsl[l]"])
}

// TestDiscriminatedUnion_RequestParams_HSLVariant verifies the third variant of
// the standalone union. Three variants is the point: with only two, a bug that
// encoded "the other one" instead of "the named one" would still pass.
func TestDiscriminatedUnion_RequestParams_HSLVariant(t *testing.T) {
	model := "hsl"
	h := int64(240)
	s := int64(100)
	l := int64(50)

	params := &colorParams{
		Model: &model,
		Hsl: &hslVariantParams{
			H: &h,
			S: &s,
			L: &l,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"hsl"}, values["model"])
	assert.Equal(t, []string{"240"}, values["hsl[h]"])
	assert.Equal(t, []string{"100"}, values["hsl[s]"])
	assert.Equal(t, []string{"50"}, values["hsl[l]"])

	// RGB and HSV variant structs are absent.
	assert.Nil(t, values["rgb[r]"])
	assert.Nil(t, values["rgb[g]"])
	assert.Nil(t, values["rgb[b]"])
	assert.Nil(t, values["hsv[h]"])
	assert.Nil(t, values["hsv[s]"])
	assert.Nil(t, values["hsv[v]"])
}

// TestDiscriminatedUnion_RequestParams_NilVariantFieldsOmitted verifies that
// when only a discriminator is set (no variant structs), only the discriminator
// appears in the encoded output.
func TestDiscriminatedUnion_RequestParams_NilVariantFieldsOmitted(t *testing.T) {
	model := "rgb"
	params := &colorParams{
		Model: &model,
		// All variant struct pointers nil — none should appear.
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"rgb"}, values["model"])
	assert.Nil(t, values["rgb[r]"])
	assert.Nil(t, values["rgb[g]"])
	assert.Nil(t, values["rgb[b]"])
	assert.Nil(t, values["hsv[h]"])
	assert.Nil(t, values["hsv[s]"])
	assert.Nil(t, values["hsv[v]"])
	assert.Nil(t, values["hsl[h]"])
	assert.Nil(t, values["hsl[s]"])
	assert.Nil(t, values["hsl[l]"])
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

// TestDiscriminatedUnion_InlineUnion_MagicLlama verifies the third inline union
// variant. Three variants is the point: with only two, a bug that encoded "the
// other one" instead of "the named one" would still pass.
func TestDiscriminatedUnion_InlineUnion_MagicLlama(t *testing.T) {
	llamaType := "magic_llama"
	name := "Llamadeus"
	spell := "levitate"

	params := &llamaParams{
		Name: &name,
		Type: &llamaType,
		MagicLlama: &magicLlamaParams{
			Spell: &spell,
		},
	}

	body := &form.Values{}
	form.AppendTo(body, params)
	values := body.ToValues()

	assert.Equal(t, []string{"Llamadeus"}, values["name"])
	assert.Equal(t, []string{"magic_llama"}, values["type"])
	assert.Equal(t, []string{"levitate"}, values["magic_llama[spell]"])

	assert.Nil(t, values["alien_llama[planet]"])
	assert.Nil(t, values["earth_llama[country]"])
}

// TestDiscriminatedUnion_Response_UnmarshalRGBVariant verifies that JSON for
// an RGB color union is unmarshaled with the discriminator and variant fields
// correctly populated, and other variant fields left nil.
func TestDiscriminatedUnion_Response_UnmarshalRGBVariant(t *testing.T) {
	data := []byte(`{"model":"rgb","rgb":{"r":255,"g":128,"b":0}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelRGB, c.Model)
	assert.NotNil(t, c.Rgb)
	assert.Equal(t, int64(255), c.Rgb.R)
	assert.Equal(t, int64(128), c.Rgb.G)
	assert.Equal(t, int64(0), c.Rgb.B)

	// Other variant fields remain nil.
	assert.Nil(t, c.Hsv)
	assert.Nil(t, c.Hsl)
}

// TestDiscriminatedUnion_Response_UnmarshalHSVVariant verifies the HSV
// discriminated union variant unmarshals correctly.
func TestDiscriminatedUnion_Response_UnmarshalHSVVariant(t *testing.T) {
	data := []byte(`{"model":"hsv","hsv":{"h":180,"s":100,"v":50}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelHSV, c.Model)
	assert.NotNil(t, c.Hsv)
	assert.Equal(t, int64(180), c.Hsv.H)
	assert.Equal(t, int64(100), c.Hsv.S)
	assert.Equal(t, int64(50), c.Hsv.V)

	assert.Nil(t, c.Rgb)
	assert.Nil(t, c.Hsl)
}

// TestDiscriminatedUnion_Response_UnmarshalHSLVariant verifies the HSL
// discriminated union variant unmarshals correctly.
func TestDiscriminatedUnion_Response_UnmarshalHSLVariant(t *testing.T) {
	data := []byte(`{"model":"hsl","hsl":{"h":240,"s":100,"l":50}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelHSL, c.Model)
	assert.NotNil(t, c.Hsl)
	assert.Equal(t, int64(240), c.Hsl.H)
	assert.Equal(t, int64(100), c.Hsl.S)
	assert.Equal(t, int64(50), c.Hsl.L)

	assert.Nil(t, c.Rgb)
	assert.Nil(t, c.Hsv)
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
	assert.Nil(t, c.Rgb)
	assert.Nil(t, c.Hsv)
	assert.Nil(t, c.Hsl)
}

// TestDiscriminatedUnion_Response_UnknownDiscriminator verifies that a variant
// the API adds after this release still decodes: the discriminator is preserved
// verbatim rather than rejected, and every known variant pointer stays nil.
// Generated discriminator types are plain strings, so there is no enum
// validation to fail.
func TestDiscriminatedUnion_Response_UnknownDiscriminator(t *testing.T) {
	data := []byte(`{"model":"cmyk","cmyk":{"c":1,"m":0,"y":0,"k":0}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModel("cmyk"), c.Model)
	assert.Nil(t, c.Rgb)
	assert.Nil(t, c.Hsv)
	assert.Nil(t, c.Hsl)
}

// TestDiscriminatedUnion_Response_MultipleVariantsPopulate pins what the decoder
// does when the API's one-variant guarantee is violated: it populates every
// variant present and does not consult the discriminator. Nothing in the
// generated shape can enforce exclusivity, so a caller that trusts Model over
// the pointers is relying on the API, not on the SDK.
func TestDiscriminatedUnion_Response_MultipleVariantsPopulate(t *testing.T) {
	data := []byte(`{"model":"rgb","rgb":{"r":255,"g":0,"b":0},"hsv":{"h":180,"s":100,"v":50}}`)

	var c colorUnion
	err := json.Unmarshal(data, &c)
	assert.NoError(t, err)

	assert.Equal(t, colorModelRGB, c.Model)
	assert.NotNil(t, c.Rgb)
	assert.NotNil(t, c.Hsv, "the decoder populates a variant the discriminator does not name")
	assert.Nil(t, c.Hsl)
}

// llamaType is the discriminator enum for the inline union on a resource.
type llamaType string

const (
	llamaTypeAlienLlama llamaType = "alien_llama"
	llamaTypeEarthLlama llamaType = "earth_llama"
	llamaTypeMagicLlama llamaType = "magic_llama"
)

// llamaResource demonstrates the response-side inline union pattern: the
// discriminator and variant pointers are sibling fields alongside base resource
// fields. Only json tags are used (no form tags on response types).
type llamaResource struct {
	Name       string              `json:"name"`
	Type       llamaType           `json:"type"`
	AlienLlama *alienLlamaResource `json:"alien_llama,omitempty"`
	EarthLlama *earthLlamaResource `json:"earth_llama,omitempty"`
	MagicLlama *magicLlamaResource `json:"magic_llama,omitempty"`
}

type alienLlamaResource struct {
	Planet     string `json:"planet"`
	Telepathic bool   `json:"telepathic"`
}

type earthLlamaResource struct {
	Country string `json:"country"`
}

type magicLlamaResource struct {
	Spell string `json:"spell"`
}

// TestDiscriminatedUnion_ResponseInline_AlienLlama verifies that a response-side
// inline discriminated union (discriminator + variant payload as sibling fields
// on the resource struct) unmarshals correctly.
func TestDiscriminatedUnion_ResponseInline_AlienLlama(t *testing.T) {
	data := []byte(`{"name":"Cosmo","type":"alien_llama","alien_llama":{"planet":"Mars","telepathic":true}}`)

	var l llamaResource
	err := json.Unmarshal(data, &l)
	assert.NoError(t, err)

	assert.Equal(t, "Cosmo", l.Name)
	assert.Equal(t, llamaTypeAlienLlama, l.Type)
	assert.NotNil(t, l.AlienLlama)
	assert.Equal(t, "Mars", l.AlienLlama.Planet)
	assert.Equal(t, true, l.AlienLlama.Telepathic)

	// Other variants are nil.
	assert.Nil(t, l.EarthLlama)
	assert.Nil(t, l.MagicLlama)
}

// TestDiscriminatedUnion_ResponseInline_EarthLlama verifies the earth llama
// variant of the response-side inline union.
func TestDiscriminatedUnion_ResponseInline_EarthLlama(t *testing.T) {
	data := []byte(`{"name":"Llama Del Rey","type":"earth_llama","earth_llama":{"country":"Peru"}}`)

	var l llamaResource
	err := json.Unmarshal(data, &l)
	assert.NoError(t, err)

	assert.Equal(t, "Llama Del Rey", l.Name)
	assert.Equal(t, llamaTypeEarthLlama, l.Type)
	assert.NotNil(t, l.EarthLlama)
	assert.Equal(t, "Peru", l.EarthLlama.Country)

	assert.Nil(t, l.AlienLlama)
	assert.Nil(t, l.MagicLlama)
}

// TestDiscriminatedUnion_ResponseInline_DiscriminatorOnly verifies that a
// response with only the discriminator (no variant payload) leaves all variant
// pointers nil.
func TestDiscriminatedUnion_ResponseInline_DiscriminatorOnly(t *testing.T) {
	data := []byte(`{"name":"Mystery","type":"magic_llama"}`)

	var l llamaResource
	err := json.Unmarshal(data, &l)
	assert.NoError(t, err)

	assert.Equal(t, "Mystery", l.Name)
	assert.Equal(t, llamaTypeMagicLlama, l.Type)
	assert.Nil(t, l.AlienLlama)
	assert.Nil(t, l.EarthLlama)
	assert.Nil(t, l.MagicLlama)
}

// TestDiscriminatedUnion_ResponseInline_UnknownDiscriminator verifies forward
// compatibility for the inline shape: a variant added after this release keeps
// its discriminator value and leaves every known variant pointer nil. Base
// fields alongside the union still populate.
func TestDiscriminatedUnion_ResponseInline_UnknownDiscriminator(t *testing.T) {
	data := []byte(`{"name":"Robo","type":"robot_llama","robot_llama":{"firmware":"2.1"}}`)

	var l llamaResource
	err := json.Unmarshal(data, &l)
	assert.NoError(t, err)

	assert.Equal(t, "Robo", l.Name)
	assert.Equal(t, llamaType("robot_llama"), l.Type)
	assert.Nil(t, l.AlienLlama)
	assert.Nil(t, l.EarthLlama)
	assert.Nil(t, l.MagicLlama)
}
