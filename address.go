package stripe

// Standard address parameters.
type AddressParams struct {
	Line1      string `form:"line1"`
	Line2      string `form:"line2"`
	City       string `form:"city"`
	State      string `form:"state"`
	PostalCode string `form:"postal_code"`
	Country    string `form:"country"`
}
