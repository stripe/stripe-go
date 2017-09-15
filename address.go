package stripe

// Standard address parameters.
type AddressParams struct {
	City       string `form:"city"`
	Country    string `form:"country"`
	Line1      string `form:"line1"`
	Line2      string `form:"line2"`
	PostalCode string `form:"postal_code"`
	State      string `form:"state"`
}
