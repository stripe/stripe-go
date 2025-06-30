package stripe

// Amount describes a monetary amount in a specific currency in minor units.
type Amount struct {
	Currency Currency `form:"currency" json:"currency"`
	Value    int64    `form:"value" json:"value"`
}
