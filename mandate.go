//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The mandate includes the type of customer acceptance information, such as: `online` or `offline`.
type MandateCustomerAcceptanceType string

// List of values that MandateCustomerAcceptanceType can take
const (
	MandateCustomerAcceptanceTypeOffline MandateCustomerAcceptanceType = "offline"
	MandateCustomerAcceptanceTypeOnline  MandateCustomerAcceptanceType = "online"
)

// List of Stripe products where this mandate can be selected automatically.
type MandatePaymentMethodDetailsACSSDebitDefaultFor string

// List of values that MandatePaymentMethodDetailsACSSDebitDefaultFor can take
const (
	MandatePaymentMethodDetailsACSSDebitDefaultForInvoice      MandatePaymentMethodDetailsACSSDebitDefaultFor = "invoice"
	MandatePaymentMethodDetailsACSSDebitDefaultForSubscription MandatePaymentMethodDetailsACSSDebitDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type MandatePaymentMethodDetailsACSSDebitPaymentSchedule string

// List of values that MandatePaymentMethodDetailsACSSDebitPaymentSchedule can take
const (
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleCombined MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "combined"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleInterval MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "interval"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleSporadic MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type MandatePaymentMethodDetailsACSSDebitTransactionType string

// List of values that MandatePaymentMethodDetailsACSSDebitTransactionType can take
const (
	MandatePaymentMethodDetailsACSSDebitTransactionTypeBusiness MandatePaymentMethodDetailsACSSDebitTransactionType = "business"
	MandatePaymentMethodDetailsACSSDebitTransactionTypePersonal MandatePaymentMethodDetailsACSSDebitTransactionType = "personal"
)

// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
type MandatePaymentMethodDetailsBACSDebitNetworkStatus string

// List of values that MandatePaymentMethodDetailsBACSDebitNetworkStatus can take
const (
	MandatePaymentMethodDetailsBACSDebitNetworkStatusAccepted MandatePaymentMethodDetailsBACSDebitNetworkStatus = "accepted"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusPending  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "pending"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRefused  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "refused"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRevoked  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "revoked"
)

// When the mandate is revoked on the Bacs network this field displays the reason for the revocation.
type MandatePaymentMethodDetailsBACSDebitRevocationReason string

// List of values that MandatePaymentMethodDetailsBACSDebitRevocationReason can take
const (
	MandatePaymentMethodDetailsBACSDebitRevocationReasonAccountClosed         MandatePaymentMethodDetailsBACSDebitRevocationReason = "account_closed"
	MandatePaymentMethodDetailsBACSDebitRevocationReasonBankAccountRestricted MandatePaymentMethodDetailsBACSDebitRevocationReason = "bank_account_restricted"
	MandatePaymentMethodDetailsBACSDebitRevocationReasonBankOwnershipChanged  MandatePaymentMethodDetailsBACSDebitRevocationReason = "bank_ownership_changed"
	MandatePaymentMethodDetailsBACSDebitRevocationReasonCouldNotProcess       MandatePaymentMethodDetailsBACSDebitRevocationReason = "could_not_process"
	MandatePaymentMethodDetailsBACSDebitRevocationReasonDebitNotAuthorized    MandatePaymentMethodDetailsBACSDebitRevocationReason = "debit_not_authorized"
)

// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
type MandatePaymentMethodDetailsPaytoAmountType string

// List of values that MandatePaymentMethodDetailsPaytoAmountType can take
const (
	MandatePaymentMethodDetailsPaytoAmountTypeFixed   MandatePaymentMethodDetailsPaytoAmountType = "fixed"
	MandatePaymentMethodDetailsPaytoAmountTypeMaximum MandatePaymentMethodDetailsPaytoAmountType = "maximum"
)

// The periodicity at which payments will be collected. Defaults to `adhoc`.
type MandatePaymentMethodDetailsPaytoPaymentSchedule string

// List of values that MandatePaymentMethodDetailsPaytoPaymentSchedule can take
const (
	MandatePaymentMethodDetailsPaytoPaymentScheduleAdhoc       MandatePaymentMethodDetailsPaytoPaymentSchedule = "adhoc"
	MandatePaymentMethodDetailsPaytoPaymentScheduleAnnual      MandatePaymentMethodDetailsPaytoPaymentSchedule = "annual"
	MandatePaymentMethodDetailsPaytoPaymentScheduleDaily       MandatePaymentMethodDetailsPaytoPaymentSchedule = "daily"
	MandatePaymentMethodDetailsPaytoPaymentScheduleFortnightly MandatePaymentMethodDetailsPaytoPaymentSchedule = "fortnightly"
	MandatePaymentMethodDetailsPaytoPaymentScheduleMonthly     MandatePaymentMethodDetailsPaytoPaymentSchedule = "monthly"
	MandatePaymentMethodDetailsPaytoPaymentScheduleQuarterly   MandatePaymentMethodDetailsPaytoPaymentSchedule = "quarterly"
	MandatePaymentMethodDetailsPaytoPaymentScheduleSemiAnnual  MandatePaymentMethodDetailsPaytoPaymentSchedule = "semi_annual"
	MandatePaymentMethodDetailsPaytoPaymentScheduleWeekly      MandatePaymentMethodDetailsPaytoPaymentSchedule = "weekly"
)

// The purpose for which payments are made. Has a default value based on your merchant category code.
type MandatePaymentMethodDetailsPaytoPurpose string

// List of values that MandatePaymentMethodDetailsPaytoPurpose can take
const (
	MandatePaymentMethodDetailsPaytoPurposeDependantSupport MandatePaymentMethodDetailsPaytoPurpose = "dependant_support"
	MandatePaymentMethodDetailsPaytoPurposeGovernment       MandatePaymentMethodDetailsPaytoPurpose = "government"
	MandatePaymentMethodDetailsPaytoPurposeLoan             MandatePaymentMethodDetailsPaytoPurpose = "loan"
	MandatePaymentMethodDetailsPaytoPurposeMortgage         MandatePaymentMethodDetailsPaytoPurpose = "mortgage"
	MandatePaymentMethodDetailsPaytoPurposeOther            MandatePaymentMethodDetailsPaytoPurpose = "other"
	MandatePaymentMethodDetailsPaytoPurposePension          MandatePaymentMethodDetailsPaytoPurpose = "pension"
	MandatePaymentMethodDetailsPaytoPurposePersonal         MandatePaymentMethodDetailsPaytoPurpose = "personal"
	MandatePaymentMethodDetailsPaytoPurposeRetail           MandatePaymentMethodDetailsPaytoPurpose = "retail"
	MandatePaymentMethodDetailsPaytoPurposeSalary           MandatePaymentMethodDetailsPaytoPurpose = "salary"
	MandatePaymentMethodDetailsPaytoPurposeTax              MandatePaymentMethodDetailsPaytoPurpose = "tax"
	MandatePaymentMethodDetailsPaytoPurposeUtility          MandatePaymentMethodDetailsPaytoPurpose = "utility"
)

// This mandate corresponds with a specific payment method type. The `payment_method_details` includes an additional hash with the same name and contains mandate information that's specific to that payment method.
type MandatePaymentMethodDetailsType string

// List of values that MandatePaymentMethodDetailsType can take
const (
	MandatePaymentMethodDetailsTypeACSSDebit     MandatePaymentMethodDetailsType = "acss_debit"
	MandatePaymentMethodDetailsTypeAUBECSDebit   MandatePaymentMethodDetailsType = "au_becs_debit"
	MandatePaymentMethodDetailsTypeBACSDebit     MandatePaymentMethodDetailsType = "bacs_debit"
	MandatePaymentMethodDetailsTypeBLIK          MandatePaymentMethodDetailsType = "blik"
	MandatePaymentMethodDetailsTypeCard          MandatePaymentMethodDetailsType = "card"
	MandatePaymentMethodDetailsTypeLink          MandatePaymentMethodDetailsType = "link"
	MandatePaymentMethodDetailsTypeSEPADebit     MandatePaymentMethodDetailsType = "sepa_debit"
	MandatePaymentMethodDetailsTypeUSBankAccount MandatePaymentMethodDetailsType = "us_bank_account"
)

// Mandate collection method
type MandatePaymentMethodDetailsUSBankAccountCollectionMethod string

// List of values that MandatePaymentMethodDetailsUSBankAccountCollectionMethod can take
const (
	MandatePaymentMethodDetailsUSBankAccountCollectionMethodPaper MandatePaymentMethodDetailsUSBankAccountCollectionMethod = "paper"
)

// The mandate status indicates whether or not you can use it to initiate a payment.
type MandateStatus string

// List of values that MandateStatus can take
const (
	MandateStatusActive   MandateStatus = "active"
	MandateStatusInactive MandateStatus = "inactive"
	MandateStatusPending  MandateStatus = "pending"
)

// The type of the mandate.
type MandateType string

// List of values that MandateType can take
const (
	MandateTypeMultiUse  MandateType = "multi_use"
	MandateTypeSingleUse MandateType = "single_use"
)

// Retrieves a Mandate object.
type MandateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *MandateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a Mandate object.
type MandateRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *MandateRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type MandateCustomerAcceptanceOffline struct{}
type MandateCustomerAcceptanceOnline struct {
	// The customer accepts the mandate from this IP address.
	IPAddress string `json:"ip_address"`
	// The customer accepts the mandate using the user agent of the browser.
	UserAgent string `json:"user_agent"`
}
type MandateCustomerAcceptance struct {
	// The time that the customer accepts the mandate.
	AcceptedAt int64                             `json:"accepted_at"`
	Offline    *MandateCustomerAcceptanceOffline `json:"offline"`
	Online     *MandateCustomerAcceptanceOnline  `json:"online"`
	// The mandate includes the type of customer acceptance information, such as: `online` or `offline`.
	Type MandateCustomerAcceptanceType `json:"type"`
}
type MandateMultiUse struct{}
type MandatePaymentMethodDetailsACSSDebit struct {
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []MandatePaymentMethodDetailsACSSDebitDefaultFor `json:"default_for"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule MandatePaymentMethodDetailsACSSDebitPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType MandatePaymentMethodDetailsACSSDebitTransactionType `json:"transaction_type"`
}
type MandatePaymentMethodDetailsAmazonPay struct{}
type MandatePaymentMethodDetailsAUBECSDebit struct {
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsBACSDebit struct {
	// The display name for the account on this mandate.
	DisplayName string `json:"display_name"`
	// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
	NetworkStatus MandatePaymentMethodDetailsBACSDebitNetworkStatus `json:"network_status"`
	// The unique reference identifying the mandate on the Bacs network.
	Reference string `json:"reference"`
	// When the mandate is revoked on the Bacs network this field displays the reason for the revocation.
	RevocationReason MandatePaymentMethodDetailsBACSDebitRevocationReason `json:"revocation_reason"`
	// The service user number for the account on this mandate.
	ServiceUserNumber string `json:"service_user_number"`
	// The URL that will contain the mandate that the customer has signed.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsCard struct{}
type MandatePaymentMethodDetailsCashApp struct{}
type MandatePaymentMethodDetailsKakaoPay struct{}
type MandatePaymentMethodDetailsKlarna struct{}
type MandatePaymentMethodDetailsKrCard struct{}
type MandatePaymentMethodDetailsLink struct{}
type MandatePaymentMethodDetailsNaverPay struct{}
type MandatePaymentMethodDetailsNzBankAccount struct{}
type MandatePaymentMethodDetailsPaypal struct {
	// The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.
	BillingAgreementID string `json:"billing_agreement_id"`
	// PayPal account PayerID. This identifier uniquely identifies the PayPal customer.
	PayerID string `json:"payer_id"`
}
type MandatePaymentMethodDetailsPayto struct {
	// Amount that will be collected. It is required when `amount_type` is `fixed`.
	Amount int64 `json:"amount"`
	// The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.
	AmountType MandatePaymentMethodDetailsPaytoAmountType `json:"amount_type"`
	// Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.
	EndDate string `json:"end_date"`
	// The periodicity at which payments will be collected. Defaults to `adhoc`.
	PaymentSchedule MandatePaymentMethodDetailsPaytoPaymentSchedule `json:"payment_schedule"`
	// The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.
	PaymentsPerPeriod int64 `json:"payments_per_period"`
	// The purpose for which payments are made. Has a default value based on your merchant category code.
	Purpose MandatePaymentMethodDetailsPaytoPurpose `json:"purpose"`
	// Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.
	StartDate string `json:"start_date"`
}
type MandatePaymentMethodDetailsRevolutPay struct{}
type MandatePaymentMethodDetailsSEPADebit struct {
	// The unique reference of the mandate.
	Reference string `json:"reference"`
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsUSBankAccount struct {
	// Mandate collection method
	CollectionMethod MandatePaymentMethodDetailsUSBankAccountCollectionMethod `json:"collection_method"`
}
type MandatePaymentMethodDetails struct {
	ACSSDebit     *MandatePaymentMethodDetailsACSSDebit     `json:"acss_debit"`
	AmazonPay     *MandatePaymentMethodDetailsAmazonPay     `json:"amazon_pay"`
	AUBECSDebit   *MandatePaymentMethodDetailsAUBECSDebit   `json:"au_becs_debit"`
	BACSDebit     *MandatePaymentMethodDetailsBACSDebit     `json:"bacs_debit"`
	Card          *MandatePaymentMethodDetailsCard          `json:"card"`
	CashApp       *MandatePaymentMethodDetailsCashApp       `json:"cashapp"`
	KakaoPay      *MandatePaymentMethodDetailsKakaoPay      `json:"kakao_pay"`
	Klarna        *MandatePaymentMethodDetailsKlarna        `json:"klarna"`
	KrCard        *MandatePaymentMethodDetailsKrCard        `json:"kr_card"`
	Link          *MandatePaymentMethodDetailsLink          `json:"link"`
	NaverPay      *MandatePaymentMethodDetailsNaverPay      `json:"naver_pay"`
	NzBankAccount *MandatePaymentMethodDetailsNzBankAccount `json:"nz_bank_account"`
	Paypal        *MandatePaymentMethodDetailsPaypal        `json:"paypal"`
	Payto         *MandatePaymentMethodDetailsPayto         `json:"payto"`
	RevolutPay    *MandatePaymentMethodDetailsRevolutPay    `json:"revolut_pay"`
	SEPADebit     *MandatePaymentMethodDetailsSEPADebit     `json:"sepa_debit"`
	// This mandate corresponds with a specific payment method type. The `payment_method_details` includes an additional hash with the same name and contains mandate information that's specific to that payment method.
	Type          MandatePaymentMethodDetailsType           `json:"type"`
	USBankAccount *MandatePaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
}
type MandateSingleUse struct {
	// The amount of the payment on a single use mandate.
	Amount int64 `json:"amount"`
	// The currency of the payment on a single use mandate.
	Currency Currency `json:"currency"`
}

// A Mandate is a record of the permission that your customer gives you to debit their payment method.
type Mandate struct {
	APIResource
	CustomerAcceptance *MandateCustomerAcceptance `json:"customer_acceptance"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool             `json:"livemode"`
	MultiUse *MandateMultiUse `json:"multi_use"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) that the mandate is intended for.
	OnBehalfOf string `json:"on_behalf_of"`
	// ID of the payment method associated with this mandate.
	PaymentMethod        *PaymentMethod               `json:"payment_method"`
	PaymentMethodDetails *MandatePaymentMethodDetails `json:"payment_method_details"`
	SingleUse            *MandateSingleUse            `json:"single_use"`
	// The mandate status indicates whether or not you can use it to initiate a payment.
	Status MandateStatus `json:"status"`
	// The type of the mandate.
	Type MandateType `json:"type"`
}

// UnmarshalJSON handles deserialization of a Mandate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (m *Mandate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		m.ID = id
		return nil
	}

	type mandate Mandate
	var v mandate
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*m = Mandate(v)
	return nil
}
