package stripe

import "encoding/json"

// PaymentMethodFPXAccountHolderType is a list of string values that FPX AccountHolderType accepts.
type PaymentMethodFPXAccountHolderType string

// List of values that PaymentMethodFPXAccountHolderType can take
const (
	PaymentMethodFPXAccountHolderTypeIndividual PaymentMethodFPXAccountHolderType = "individual"
	PaymentMethodFPXAccountHolderTypeCompany    PaymentMethodFPXAccountHolderType = "company"
)

// PaymentMethodType is the list of allowed values for the payment method type.
type PaymentMethodType string

// List of values that PaymentMethodType can take.
const (
	PaymentMethodTypeCard        PaymentMethodType = "card"
	PaymentMethodTypeCardPresent PaymentMethodType = "card_present"
)

// PaymentMethodCardBrand is the list of allowed values for the brand property on a
// Card PaymentMethod.
type PaymentMethodCardBrand string

// List of values that PaymentMethodCardBrand can take.
const (
	PaymentMethodCardBrandAmex       PaymentMethodCardBrand = "amex"
	PaymentMethodCardBrandDiners     PaymentMethodCardBrand = "diners"
	PaymentMethodCardBrandDiscover   PaymentMethodCardBrand = "discover"
	PaymentMethodCardBrandJCB        PaymentMethodCardBrand = "jcb"
	PaymentMethodCardBrandMastercard PaymentMethodCardBrand = "mastercard"
	PaymentMethodCardBrandUnionpay   PaymentMethodCardBrand = "unionpay"
	PaymentMethodCardBrandUnknown    PaymentMethodCardBrand = "unknown"
	PaymentMethodCardBrandVisa       PaymentMethodCardBrand = "visa"
)

// PaymentMethodCardWalletType is the list of allowed values for the type a wallet can take on
// a Card PaymentMethod.
type PaymentMethodCardWalletType string

// List of values that PaymentMethodCardWalletType can take.
const (
	PaymentMethodCardWalletTypeAmexExpressCheckout PaymentMethodCardWalletType = "amex_express_checkout"
	PaymentMethodCardWalletTypeApplePay            PaymentMethodCardWalletType = "apple_pay"
	PaymentMethodCardWalletTypeGooglePay           PaymentMethodCardWalletType = "google_pay"
	PaymentMethodCardWalletTypeMasterpass          PaymentMethodCardWalletType = "masterpass"
	PaymentMethodCardWalletTypeSamsungPay          PaymentMethodCardWalletType = "samsung_pay"
	PaymentMethodCardWalletTypeVisaCheckout        PaymentMethodCardWalletType = "visa_checkout"
)

// BillingDetailsParams is the set of parameters that can be used as billing details
// when creating or updating a PaymentMethod
type BillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// PaymentMethodCardParams is the set of parameters allowed for the `card` hash when creating a
// PaymentMethod of type card.
type PaymentMethodCardParams struct {
	CVC      *string `form:"cvc"`
	ExpMonth *string `form:"exp_month"`
	ExpYear  *string `form:"exp_year"`
	Number   *string `form:"number"`
	Token    *string `form:"token"`
}

// PaymentMethodFPXParams is the set of parameters allowed for the `fpx` hash when creating a
// PaymentMethod of type fpx.
type PaymentMethodFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// PaymentMethodParams is the set of parameters that can be used when creating or updating a
// PaymentMethod.
type PaymentMethodParams struct {
	Params         `form:"*"`
	BillingDetails *BillingDetailsParams    `form:"billing_details"`
	Card           *PaymentMethodCardParams `form:"card"`
	FPX            *PaymentMethodFPXParams  `form:"fpx"`
	PaymentMethod  *string                  `form:"payment_method"`
	Type           *string                  `form:"type"`
}

// PaymentMethodAttachParams is the set of parameters that can be used when attaching a
// PaymentMethod to a Customer.
type PaymentMethodAttachParams struct {
	Params   `form:"*"`
	Customer *string `form:"customer"`
}

// PaymentMethodDetachParams is the set of parameters that can be used when detaching a
// PaymentMethod.
type PaymentMethodDetachParams struct {
	Params `form:"*"`
}

// PaymentMethodListParams is the set of parameters that can be used when listing PaymentMethods.
type PaymentMethodListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Type       *string `form:"type"`
}

// BillingDetails represents the billing details associated with a PaymentMethod.
type BillingDetails struct {
	Address *Address `json:"address"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
}

// PaymentMethodCardChecks represents the checks associated with a Card PaymentMethod.
type PaymentMethodCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// PaymentMethodCardThreeDSecureUsage represents the 3DS usage for that Card PaymentMethod.
type PaymentMethodCardThreeDSecureUsage struct {
	Supported bool `json:"supported"`
}

// PaymentMethodCardWallet represents the details of the card wallet if this Card PaymentMethod
// is part of a card wallet.
type PaymentMethodCardWallet struct {
	DynamicLast4 string                      `json:"dynamic_last4"`
	Type         PaymentMethodCardWalletType `json:"type"`
}

// PaymentMethodCard represents the card-specific properties on a PaymentMethod.
type PaymentMethodCard struct {
	Brand             PaymentMethodCardBrand              `json:"brand"`
	Checks            *PaymentMethodCardChecks            `json:"checks"`
	Country           string                              `json:"country"`
	ExpMonth          uint64                              `json:"exp_month"`
	ExpYear           uint64                              `json:"exp_year"`
	Fingerprint       string                              `json:"fingerprint"`
	Funding           CardFunding                         `json:"funding"`
	Last4             string                              `json:"last4"`
	ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
	Wallet            *PaymentMethodCardWallet            `json:"wallet"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// PaymentMethodCardPresent represents the card-present-specific properties on a PaymentMethod.
type PaymentMethodCardPresent struct {
}

// PaymentMethodFPX represents Malaysia FPX PaymentMethod (Malaysia Only).
type PaymentMethodFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              string                            `json:"bank"`
	TransactionID     string                            `json:"transaction_id"`
}

// PaymentMethod is the resource representing a PaymentMethod.
type PaymentMethod struct {
	BillingDetails *BillingDetails           `json:"billing_details"`
	Card           *PaymentMethodCard        `json:"card"`
	CardPresent    *PaymentMethodCardPresent `json:"card_present"`
	Created        int64                     `json:"created"`
	Customer       *Customer                 `json:"customer"`
	FPX            *PaymentMethodFPX         `json:"fpx"`
	ID             string                    `json:"id"`
	Livemode       bool                      `json:"livemode"`
	Metadata       map[string]string         `json:"metadata"`
	Object         string                    `json:"object"`
	Type           PaymentMethodType         `json:"type"`
}

// PaymentMethodList is a list of PaymentMethods as retrieved from a list endpoint.
type PaymentMethodList struct {
	ListMeta
	Data []*PaymentMethod `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentMethod.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *PaymentMethod) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type pm PaymentMethod
	var v pm
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = PaymentMethod(v)
	return nil
}
