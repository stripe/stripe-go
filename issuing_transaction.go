package stripe

import "encoding/json"

// IssuingTransactionType is the type of an issuing transaction.
type IssuingTransactionType string

// List of values that IssuingTransactionType can take.
const (
	IssuingTransactionTypeCapture IssuingTransactionType = "capture"
	IssuingTransactionTypeRefund  IssuingTransactionType = "refund"
)

// IssuingTransactionPurchaseDetailsFuelType is the type of fuel purchased in a transaction.
type IssuingTransactionPurchaseDetailsFuelType string

// List of values that IssuingTransactionType can take.
const (
	IssuingTransactionPurchaseDetailsFuelTypeDiesel          IssuingTransactionPurchaseDetailsFuelType = "diesel"
	IssuingTransactionPurchaseDetailsFuelTypeOther           IssuingTransactionPurchaseDetailsFuelType = "other"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedPlus    IssuingTransactionPurchaseDetailsFuelType = "unleaded_plus"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedRegular IssuingTransactionPurchaseDetailsFuelType = "unleaded_regular"
	IssuingTransactionPurchaseDetailsFuelTypeUnleadedSuper   IssuingTransactionPurchaseDetailsFuelType = "unleaded_super"
)

// IssuingTransactionPurchaseDetailsFuelUnit is the unit of fuel purchased in a transaction.
type IssuingTransactionPurchaseDetailsFuelUnit string

// List of values that IssuingTransactionPurchaseDetailsFuelUnit can take.
const (
	IssuingTransactionPurchaseDetailsFuelUnitLiter    IssuingTransactionPurchaseDetailsFuelUnit = "liter"
	IssuingTransactionPurchaseDetailsFuelUnitUSGallon IssuingTransactionPurchaseDetailsFuelUnit = "us_gallon"
)

// IssuingTransactionParams is the set of parameters that can be used when creating or updating an issuing transaction.
type IssuingTransactionParams struct {
	Params `form:"*"`
}

// IssuingTransactionListParams is the set of parameters that can be used when listing issuing transactions.
type IssuingTransactionListParams struct {
	ListParams   `form:"*"`
	Card         *string           `form:"card"`
	Cardholder   *string           `form:"cardholder"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// IssuingTransactionAmountDetails is the resource representing the breakdown of the amount.
type IssuingTransactionAmountDetails struct {
	ATMFee int64 `json:"atm_fee"`
}

// IssuingTransactionPurchaseDetailsFlightSegment contains extra information about the flight in this transaction.
type IssuingTransactionPurchaseDetailsFlightSegment struct {
	ArrivalAirportCode   string `json:"arrival_airport_code"`
	Carrier              string `json:"carrier"`
	DepartureAirportCode string `json:"departure_airport_code"`
	FlightNumber         string `json:"flight_number"`
	ServiceClass         string `json:"service_class"`
	StopoverAllowed      bool   `json:"stopover_allowed"`
}

// IssuingTransactionPurchaseDetailsFlight contains extra information about the flight in this transaction.
type IssuingTransactionPurchaseDetailsFlight struct {
	DepartureAt   int64                                             `json:"departure_at"`
	PassengerName string                                            `json:"passenger_name"`
	Refundable    bool                                              `json:"refundable"`
	Segments      []*IssuingTransactionPurchaseDetailsFlightSegment `json:"segments"`
	TravelAgency  string                                            `json:"travel_agency"`
}

// IssuingTransactionPurchaseDetailsFuel contains extra information about the fuel purchase in this transaction.
type IssuingTransactionPurchaseDetailsFuel struct {
	Type            IssuingTransactionPurchaseDetailsFuelType `json:"type"`
	Unit            IssuingTransactionPurchaseDetailsFuelUnit `json:"unit"`
	UnitCostDecimal float64                                   `json:"unit_cost_decimal,string"`
	VolumeDecimal   float64                                   `json:"volume_decimal,string"`
}

// IssuingTransactionPurchaseDetailsLodging contains extra information about the lodging purchase in this transaction.
type IssuingTransactionPurchaseDetailsLodging struct {
	CheckInAt int64 `json:"check_in_at"`
	Nights    int64 `json:"nights"`
}

// IssuingTransactionPurchaseDetailsReceipt contains extra information about the line items this transaction.
type IssuingTransactionPurchaseDetailsReceipt struct {
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Total       int64   `json:"total"`
	UnitCost    int64   `json:"unit_cost"`
}

// IssuingTransactionPurchaseDetails contains extra information provided by the merchant.
type IssuingTransactionPurchaseDetails struct {
	Flight    *IssuingTransactionPurchaseDetailsFlight    `json:"flight"`
	Fuel      *IssuingTransactionPurchaseDetailsFuel      `json:"fuel"`
	Lodging   *IssuingTransactionPurchaseDetailsLodging   `json:"lodging"`
	Receipt   []*IssuingTransactionPurchaseDetailsReceipt `json:"receipt"`
	Reference string                                      `json:"reference"`
}

// IssuingTransaction is the resource representing a Stripe issuing transaction.
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
	MerchantData       *IssuingAuthorizationMerchantData  `json:"merchant_data"`
	MerchantAmount     int64                              `json:"merchant_amount"`
	MerchantCurrency   Currency                           `json:"merchant_currency"`
	Metadata           map[string]string                  `json:"metadata"`
	Object             string                             `json:"object"`
	PurchaseDetails    *IssuingTransactionPurchaseDetails `json:"purchase_details"`
	Type               IssuingTransactionType             `json:"type"`
}

// IssuingTransactionList is a list of issuing transactions as retrieved from a list endpoint.
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
