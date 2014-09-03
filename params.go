package stripe

import (
	"net/url"
	"reflect"
	"strconv"
)

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	Start, End string
	Limit      uint64
	Filters    Filters
}

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	Expand      []string
	Meta        map[string]string
	AccessToken string
}

// getTag gets the tagName tag from an element in a struct. It takes a struct
// and a fieldName (string), and returns the tag named tagName ("stripe_field",
// etc) for that fieldName.
func getTag(m interface{}, tagName, fieldName string) string {
	f, _ := reflect.TypeOf(m).Elem().FieldByName(fieldName)
	return f.Tag.Get(tagName)
}

// addParamsToValues takes an interface (usually *SomeTypeParams) and a pointer
// to a url.Values. It iterates over each field in the interface (using the
// getFieldTypes method), and adds the value of each field to the url.Values.
func parseParams(params interface{}, values *url.Values) {
	var val string

	for fieldName, reflection := range getFieldTypes(params) {
		name := reflection.Name()
		switch name {
		case "string":
			val = getField(params, fieldName).String()
		case "int64":
			val = parseInt64(params, fieldName)
		case "uint64":
			val = parseUInt64(params, fieldName)
		case "float64":
			val = parseFloat64(params, fieldName)
		case "bool":
			val = parseBool(params, fieldName)
		}

		if val != "" {
			values.Add(getTag(params, "stripe_field", fieldName), val)
		}
	}
}

// getFieldTypes takes a struct m and returns a map of strings (field names) to
// reflect.Types (field types).
func getFieldTypes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	attrs := make(map[string]reflect.Type)
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

// getValue gets the reflect.Value of fieldName in the struct m.
func getValue(m interface{}, fieldName string) reflect.Value {
	val := reflect.ValueOf(m)
	return reflect.Indirect(val).FieldByName(fieldName)
}

// parseBool gets the value of fieldName in the struct m (bool), converts it to
// a string, and returns the result. If the value is the zero value (false), it
// returns a blank string.
func parseBool(m interface{}, fieldName string) string {
	val := getValue(m, fieldName).Bool()

	if val {
		return strconv.FormatBool(val)
	} else {
		return ""
	}
}

// parseInt64 gets the value of fieldName in the struct m (int), converts it to a
// string, and returns the result. If the value is the zero value (0), it
// returns a blank string.
func parseInt64(m interface{}, fieldName string) string {
	val := getValue(m, fieldName).Int()

	if val == 0 {
		return ""
	} else {
		return strconv.FormatInt(val, 10)
	}
}

func parseUInt64(m interface{}, fieldName string) string {
	val := getValue(m, fieldName).Uint()

	if val == 0 {
		return ""
	} else {
		return strconv.FormatUint(val, 10)
	}
}

// parseFloat64 gets the value of fieldName in the struct m (float64), converts
// it to a string, and returns the result. If the value is the zero value (0.0),
// it returns a blank string.
func parseFloat64(m interface{}, fieldName string) string {
	val := getValue(m, fieldName).Float()

	if val == 0.0 {
		return ""
	} else {
		return strconv.FormatFloat(val, 'f', 2, 32)
	}
}
