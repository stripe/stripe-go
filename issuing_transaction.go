//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of fuel that was purchased. One of `diesel`, `unleaded_plus`, `unleaded_regular`, `unleaded_super`, or `other`.
type IssuingTransactionPurchaseDetailsFuelType string

// List of values that IssuingTransactionPurchaseDetailsFuelType can take
const (
	IssuingTransactionPurchaseDetailsFuelTypeDiesel          IssuingTransactionPurchaseDetailsFuelType = "diesel"
	IssuingTransactionPurchaseDetailsFuelTypeOther           IssuingTransactionPurchaseDetailsFuelType = "other"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedPlus    IssuingTransactionPurchaseDetailsFuelType = "unleaded_plus"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedRegular IssuingTransactionPurchaseDetailsFuelType = "unleaded_regular"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedSuper   IssuingTransactionPurchaseDetailsFuelType = "unleaded_super"
)

// The units for `volume_decimal`. One of `us_gallon` or `liter`.
type IssuingTransactionPurchaseDetailsFuelUnit string

// List of values that IssuingTransactionPurchaseDetailsFuelUnit can take
const (
	IssuingTransactionPurchaseDetailsFuelUnitLiter    IssuingTransactionPurchaseDetailsFuelUnit = "liter"
	IssuingTransactionPurchaseDetailsFuelUnitUSGallon IssuingTransactionPurchaseDetailsFuelUnit = "us_gallon"
)

// The nature of the transaction.
type IssuingTransactionType string

// List of values that IssuingTransactionType can take
const (
	IssuingTransactionTypeCapture IssuingTransactionType = "capture"
	IssuingTransactionTypeRefund  IssuingTransactionType = "refund"
)

// The digital wallet used for this transaction. One of `apple_pay`, `google_pay`, or `samsung_pay`.
type IssuingTransactionWallet string

// List of values that IssuingTransactionWallet can take
const (
	IssuingTransactionWalletApplePay   IssuingTransactionWallet = "apple_pay"
	IssuingTransactionWalletGooglePay  IssuingTransactionWallet = "google_pay"
	IssuingTransactionWalletSamsungPay IssuingTransactionWallet = "samsung_pay"
)

// Returns a list of Issuing Transaction objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingTransactionListParams struct {
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Type         *string           `form:"type"`
}

// Retrieves an Issuing Transaction object.
type IssuingTransactionParams struct {
	Params `form:"*"`
}

// Detailed breakdown of amount components. These amounts are denominated in `currency` and in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal).
type IssuingTransactionAmountDetails struct {
	ATMFee int64 `json:"atm_fee"`
}

// The legs of the trip.
type IssuingTransactionPurchaseDetailsFlightSegment struct {
	ArrivalAirportCode   string `json:"arrival_airport_code"`
	Carrier              string `json:"carrier"`
	DepartureAirportCode string `json:"departure_airport_code"`
	FlightNumber         string `json:"flight_number"`
	ServiceClass         string `json:"service_class"`
	StopoverAllowed      bool   `json:"stopover_allowed"`
}

// Information about the flight that was purchased with this transaction.
type IssuingTransactionPurchaseDetailsFlight struct {
	DepartureAt   int64                                             `json:"departure_at"`
	PassengerName string                                            `json:"passenger_name"`
	Refundable    bool                                              `json:"refundable"`
	Segments      []*IssuingTransactionPurchaseDetailsFlightSegment `json:"segments"`
	TravelAgency  string                                            `json:"travel_agency"`
}

// Information about fuel that was purchased with this transaction.
type IssuingTransactionPurchaseDetailsFuel struct {
	Type            IssuingTransactionPurchaseDetailsFuelType `json:"type"`
	Unit            IssuingTransactionPurchaseDetailsFuelUnit `json:"unit"`
	UnitCostDecimal float64                                   `json:"unit_cost_decimal,string"`
	VolumeDecimal   float64                                   `json:"volume_decimal,string"`
}

// Information about lodging that was purchased with this transaction.
type IssuingTransactionPurchaseDetailsLodging struct {
	CheckInAt int64 `json:"check_in_at"`
	Nights    int64 `json:"nights"`
}

// The line items in the purchase.
type IssuingTransactionPurchaseDetailsReceipt struct {
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Total       int64   `json:"total"`
	UnitCost    int64   `json:"unit_cost"`
}

// Additional purchase information that is optionally provided by the merchant.
type IssuingTransactionPurchaseDetails struct {
	Flight    *IssuingTransactionPurchaseDetailsFlight    `json:"flight"`
	Fuel      *IssuingTransactionPurchaseDetailsFuel      `json:"fuel"`
	Lodging   *IssuingTransactionPurchaseDetailsLodging   `json:"lodging"`
	Receipt   []*IssuingTransactionPurchaseDetailsReceipt `json:"receipt"`
	Reference string                                      `json:"reference"`
}

// Any use of an [issued card](https://stripe.com/docs/issuing) that results in funds entering or leaving
// your Stripe account, such as a completed purchase or refund, is represented by an Issuing
// `Transaction` object.
//
// Related guide: [Issued Card Transactions](https://stripe.com/docs/issuing/purchases/transactions).
type IssuingTransaction struct {
	APIResource
	Amount             int64                              `json:"amount"`
	AmountDetails      *IssuingTransactionAmountDetails   `json:"amount_details"`
	Authorization      *IssuingAuthorization              `json:"authorization"`
	BalanceTransaction *BalanceTransaction                `json:"balance_transaction"`
	Card               *IssuingCard                       `json:"card"`
	Cardholder         *IssuingCardholder                 `json:"cardholder"`
	Created            int64                              `json:"created"`
	Currency           Currency                           `json:"currency"`
	Dispute            *IssuingDispute                    `json:"dispute"`
	ID                 string                             `json:"id"`
	Livemode           bool                               `json:"livemode"`
	MerchantAmount     int64                              `json:"merchant_amount"`
	MerchantCurrency   Currency                           `json:"merchant_currency"`
	MerchantData       *IssuingAuthorizationMerchantData  `json:"merchant_data"`
	Metadata           map[string]string                  `json:"metadata"`
	Object             string                             `json:"object"`
	PurchaseDetails    *IssuingTransactionPurchaseDetails `json:"purchase_details"`
	Type               IssuingTransactionType             `json:"type"`
	Wallet             IssuingTransactionWallet           `json:"wallet"`
}

// IssuingTransactionList is a list of Transactions as retrieved from a list endpoint.
type IssuingTransactionList struct {
	APIResource
	ListMeta
	Data []*IssuingTransaction `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingTransaction IssuingTransaction
	var v issuingTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingTransaction(v)
	return nil
}
