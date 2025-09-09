//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The human readable local name of the field.
type V2MoneyManagementPayoutMethodsBankAccountSpecCountriesFieldLocalNameHuman struct {
	// The default content of the localizable string.
	Content string `json:"content"`
	// A unique key representing the instance of this localizable string.
	LocalizationKey string `json:"localization_key"`
}

// The list of fields for a country, along with associated information.
type V2MoneyManagementPayoutMethodsBankAccountSpecCountriesField struct {
	// The local name of the field.
	LocalName string `json:"local_name"`
	// The human readable local name of the field.
	LocalNameHuman *V2MoneyManagementPayoutMethodsBankAccountSpecCountriesFieldLocalNameHuman `json:"local_name_human"`
	// The maximum length of the field.
	MaxLength int64 `json:"max_length"`
	// THe minimum length of the field.
	MinLength int64 `json:"min_length"`
	// The placeholder value of the field.
	Placeholder string `json:"placeholder"`
	// The stripe name of the field.
	StripeName string `json:"stripe_name"`
	// The validation regex of the field.
	ValidationRegex string `json:"validation_regex"`
}

// The list of specs by country.
type V2MoneyManagementPayoutMethodsBankAccountSpecCountries struct {
	// The list of fields for a country, along with associated information.
	Fields []*V2MoneyManagementPayoutMethodsBankAccountSpecCountriesField `json:"fields"`
}

// The PayoutMethodsBankAccountSpec object.
type V2MoneyManagementPayoutMethodsBankAccountSpec struct {
	APIResource
	// The list of specs by country.
	Countries map[string]*V2MoneyManagementPayoutMethodsBankAccountSpecCountries `json:"countries"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
