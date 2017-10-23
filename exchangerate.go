package stripe

// ExchangeRate is the resource representing the currency exchange rates at
// a given time.
type ExchangeRate struct {
	ID    string               `json:"id"`
	Rates map[Currency]float64 `json:"rates"`
}

// ExchangeRateList is a list of exchange rates as retrieved from a list endpoint.
type ExchangeRateList struct {
	ListMeta
	Values []*ExchangeRate `json:"data"`
}

// ExchangeRateListParams are the parameters allowed during ExchangeRate listing.
type ExchangeRateListParams struct {
	ListParams `form:"*"`
}
