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
	PaymentMethodTypeAlipay         PaymentMethodType = "alipay"
	PaymentMethodTypeAUBECSDebit    PaymentMethodType = "au_becs_debit"
	PaymentMethodTypeBACSDebit      PaymentMethodType = "bacs_debit"
	PaymentMethodTypeBancontact     PaymentMethodType = "bancontact"
	PaymentMethodTypeCard           PaymentMethodType = "card"
	PaymentMethodTypeCardPresent    PaymentMethodType = "card_present"
	PaymentMethodTypeEPS            PaymentMethodType = "eps"
	PaymentMethodTypeFPX            PaymentMethodType = "fpx"
	PaymentMethodTypeGiropay        PaymentMethodType = "giropay"
	PaymentMethodTypeGrabpay        PaymentMethodType = "grabpay"
	PaymentMethodTypeIdeal          PaymentMethodType = "ideal"
	PaymentMethodTypeInteracPresent PaymentMethodType = "interac_present"
	PaymentMethodTypeOXXO           PaymentMethodType = "oxxo"
	PaymentMethodTypeP24            PaymentMethodType = "p24"
	PaymentMethodTypeSepaDebit      PaymentMethodType = "sepa_debit"
	PaymentMethodTypeSofort         PaymentMethodType = "sofort"
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

// PaymentMethodCardNetwork is the list of allowed values to represent the network
// used for a card-like transaction.
type PaymentMethodCardNetwork string

// List of values that PaymentMethodCardNetwork can take.
const (
	PaymentMethodCardNetworkAmex       PaymentMethodCardNetwork = "amex"
	PaymentMethodCardNetworkDiners     PaymentMethodCardNetwork = "diners"
	PaymentMethodCardNetworkDiscover   PaymentMethodCardNetwork = "discover"
	PaymentMethodCardNetworkInterac    PaymentMethodCardNetwork = "interac"
	PaymentMethodCardNetworkJCB        PaymentMethodCardNetwork = "jcb"
	PaymentMethodCardNetworkMastercard PaymentMethodCardNetwork = "mastercard"
	PaymentMethodCardNetworkUnionpay   PaymentMethodCardNetwork = "unionpay"
	PaymentMethodCardNetworkUnknown    PaymentMethodCardNetwork = "unknown"
	PaymentMethodCardNetworkVisa       PaymentMethodCardNetwork = "visa"
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

// PaymentMethodAlipayParams is the set of parameters allowed for the `alipay` hash when creating a
// PaymentMethod of type Alipay.
type PaymentMethodAlipayParams struct {
}

// PaymentMethodAUBECSDebitParams is the set of parameters allowed for the `AUBECSDebit` hash when creating a
// PaymentMethod of type AUBECSDebit.
type PaymentMethodAUBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BSBNumber     *string `form:"bsb_number"`
}

// PaymentMethodBACSDebitParams is the set of parameters allowed for BACS Debit payment method.
type PaymentMethodBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}

// PaymentMethodBancontactParams is the set of parameters allowed for the `bancontact` hash when creating a
// PaymentMethod of type Bancontact.
type PaymentMethodBancontactParams struct {
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

// PaymentMethodEPSParams is the set of parameters allowed for the `eps` hash when creating a
// PaymentMethod of type EPS.
type PaymentMethodEPSParams struct {
}

// PaymentMethodFPXParams is the set of parameters allowed for the `fpx` hash when creating a
// PaymentMethod of type fpx.
type PaymentMethodFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// PaymentMethodGiropayParams is the set of parameters allowed for the `giropay` hash when creating a
// PaymentMethod of type Giropay.
type PaymentMethodGiropayParams struct {
}

// PaymentMethodGrabpayParams is the set of parameters allowed for the `grabpay` hash when creating a
// PaymentMethod of type Grabpay.
type PaymentMethodGrabpayParams struct {
}

// PaymentMethodIdealParams is the set of parameters allowed for the `ideal` hash when creating a
// PaymentMethod of type ideal.
type PaymentMethodIdealParams struct {
	Bank *string `form:"bank"`
}

// PaymentMethodInteracPresentParams is the set of parameters allowed for the `interac_present` hash when creating a
// PaymentMethod of type interac_present.
type PaymentMethodInteracPresentParams struct {
}

// PaymentMethodOXXOParams is the set of parameters allowed for the `oxxo` hash when creating a
// PaymentMethod of type OXXO.
type PaymentMethodOXXOParams struct {
}

// PaymentMethodP24Params is the set of parameters allowed for the `p24` hash when creating a
// PaymentMethod of type P24.
type PaymentMethodP24Params struct {
	Bank *string `form:"bank"`
}

// PaymentMethodSepaDebitParams is the set of parameters allowed for the `sepa_debit` hash when
// creating a PaymentMethod of type sepa_debit.
type PaymentMethodSepaDebitParams struct {
	Iban *string `form:"iban"`
}

// PaymentMethodSofortParams is the set of parameters allowed for the `sofort` hash when
// creating a PaymentMethod of type sofort.
type PaymentMethodSofortParams struct {
	Country *string `form:"country"`
}

// PaymentMethodParams is the set of parameters that can be used when creating or updating a
// PaymentMethod.
type PaymentMethodParams struct {
	Params         `form:"*"`
	Alipay         *PaymentMethodAlipayParams         `form:"alipay"`
	AUBECSDebit    *PaymentMethodAUBECSDebitParams    `form:"au_becs_debit"`
	BACSDebit      *PaymentMethodBACSDebitParams      `form:"bacs_debit"`
	Bancontact     *PaymentMethodBancontactParams     `form:"bancontact"`
	BillingDetails *BillingDetailsParams              `form:"billing_details"`
	Card           *PaymentMethodCardParams           `form:"card"`
	EPS            *PaymentMethodEPSParams            `form:"eps"`
	FPX            *PaymentMethodFPXParams            `form:"fpx"`
	Giropay        *PaymentMethodGiropayParams        `form:"giropay"`
	Grabpay        *PaymentMethodGrabpayParams        `form:"grabpay"`
	Ideal          *PaymentMethodIdealParams          `form:"ideal"`
	InteracPresent *PaymentMethodInteracPresentParams `form:"interac_present"`
	OXXO           *PaymentMethodOXXOParams           `form:"oxxo"`
	P24            *PaymentMethodP24Params            `form:"p24"`
	SepaDebit      *PaymentMethodSepaDebitParams      `form:"sepa_debit"`
	Sofort         *PaymentMethodSofortParams         `form:"sofort"`
	Type           *string                            `form:"type"`

	// The following parameters are used when cloning a PaymentMethod to the connected account
	Customer      *string `form:"customer"`
	PaymentMethod *string `form:"payment_method"`
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

// PaymentMethodAlipay represents the Alipay properties.
type PaymentMethodAlipay struct {
}

// PaymentMethodAUBECSDebit represents AUBECSDebit-specific properties (Australia Only).
type PaymentMethodAUBECSDebit struct {
	BSBNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
}

// PaymentMethodBACSDebit represents the BACS Debit properties.
type PaymentMethodBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	SortCode    string `json:"sort_code"`
}

// PaymentMethodBancontact represents the Bancontact properties.
type PaymentMethodBancontact struct {
}

// PaymentMethodCardChecks represents the checks associated with a Card PaymentMethod.
type PaymentMethodCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// PaymentMethodCardNetworks represents the card networks that can be used to process the payment.
type PaymentMethodCardNetworks struct {
	Available []PaymentMethodCardNetwork `json:"available"`
	Preferred PaymentMethodCardNetwork   `json:"preferred"`
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

// PaymentMethodCard represents the card-specific properties.
type PaymentMethodCard struct {
	Brand             PaymentMethodCardBrand              `json:"brand"`
	Checks            *PaymentMethodCardChecks            `json:"checks"`
	Country           string                              `json:"country"`
	ExpMonth          uint64                              `json:"exp_month"`
	ExpYear           uint64                              `json:"exp_year"`
	Fingerprint       string                              `json:"fingerprint"`
	Funding           CardFunding                         `json:"funding"`
	Last4             string                              `json:"last4"`
	Networks          *PaymentMethodCardNetworks          `json:"networks"`
	ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
	Wallet            *PaymentMethodCardWallet            `json:"wallet"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// PaymentMethodCardPresent represents the card-present-specific properties.
type PaymentMethodCardPresent struct {
}

// PaymentMethodEPS represents the EPS properties.
type PaymentMethodEPS struct {
}

// PaymentMethodFPX represents FPX-specific properties (Malaysia Only).
type PaymentMethodFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              string                            `json:"bank"`
	TransactionID     string                            `json:"transaction_id"`
}

// PaymentMethodGiropay represents the Giropay properties.
type PaymentMethodGiropay struct {
}

// PaymentMethodGrabpay represents the Grabpay properties.
type PaymentMethodGrabpay struct {
}

// PaymentMethodIdeal represents the iDEAL-specific properties.
type PaymentMethodIdeal struct {
	Bank string `json:"bank"`
	Bic  string `json:"bic"`
}

// PaymentMethodInteracPresent represents the interac present properties.
type PaymentMethodInteracPresent struct {
}

// PaymentMethodOXXO represents the OXXO-specific properties.
type PaymentMethodOXXO struct {
}

// PaymentMethodP24 represents the P24 properties.
type PaymentMethodP24 struct {
	Bank string `json:"bank"`
}

// PaymentMethodSepaDebitGeneratedFrom represents information about the object
// that generated this PaymentMethod
type PaymentMethodSepaDebitGeneratedFrom struct {
	Charge       *Charge       `json:"charge"`
	SetupAttempt *SetupAttempt `json:"setup_attempt"`
}

// PaymentMethodSepaDebit represents the SEPA-debit-specific properties.
type PaymentMethodSepaDebit struct {
	BankCode      string                               `json:"bank_code"`
	BranchCode    string                               `json:"branch_code"`
	Country       string                               `json:"country"`
	Fingerprint   string                               `json:"fingerprint"`
	Last4         string                               `json:"last4"`
	GeneratedFrom *PaymentMethodSepaDebitGeneratedFrom `json:"generated_from"`
}

// PaymentMethodSofort represents the Sofort-specific properties.
type PaymentMethodSofort struct {
	Country string `json:"country"`
}

// PaymentMethod is the resource representing a PaymentMethod.
type PaymentMethod struct {
	APIResource
	Alipay         *PaymentMethodAlipay         `json:"alipay"`
	AUBECSDebit    *PaymentMethodAUBECSDebit    `json:"au_becs_debit"`
	BACSDebit      *PaymentMethodBACSDebit      `json:"bacs_debit"`
	Bancontact     *PaymentMethodBancontact     `json:"bancontact"`
	BillingDetails *BillingDetails              `json:"billing_details"`
	Card           *PaymentMethodCard           `json:"card"`
	CardPresent    *PaymentMethodCardPresent    `json:"card_present"`
	Created        int64                        `json:"created"`
	Customer       *Customer                    `json:"customer"`
	EPS            *PaymentMethodEPS            `json:"eps"`
	FPX            *PaymentMethodFPX            `json:"fpx"`
	Giropay        *PaymentMethodGiropay        `json:"giropay"`
	Grabpay        *PaymentMethodGrabpay        `json:"grabpay"`
	ID             string                       `json:"id"`
	Ideal          *PaymentMethodIdeal          `json:"ideal"`
	InteracPresent *PaymentMethodInteracPresent `json:"interac_present"`
	Livemode       bool                         `json:"livemode"`
	Metadata       map[string]string            `json:"metadata"`
	Object         string                       `json:"object"`
	OXXO           *PaymentMethodOXXO           `json:"oxxo"`
	P24            *PaymentMethodP24            `json:"p24"`
	SepaDebit      *PaymentMethodSepaDebit      `json:"sepa_debit"`
	Sofort         *PaymentMethodSofort         `json:"sofort"`
	Type           PaymentMethodType            `json:"type"`
}

// PaymentMethodList is a list of PaymentMethods as retrieved from a list endpoint.
type PaymentMethodList struct {
	APIResource
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
