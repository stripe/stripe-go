//
//
// File generated from our OpenAPI spec
//
//

package stripe

// ApplePayDomainParams is the set of parameters that can be used when creating an ApplePayDomain object.
type ApplePayDomainParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name"`
}

// ApplePayDomainListParams are the parameters allowed during ApplePayDomain listing.
type ApplePayDomainListParams struct {
	ListParams `form:"*"`
	DomainName *string `form:"domain_name"`
}

// ApplePayDomain is the resource representing a Stripe ApplePayDomain object
type ApplePayDomain struct {
	APIResource
	Created    int64  `json:"created"`
	Deleted    bool   `json:"deleted"`
	DomainName string `json:"domain_name"`
	ID         string `json:"id"`
	Livemode   bool   `json:"livemode"`
	Object     string `json:"object"`
}

// ApplePayDomainList is a list of ApplePayDomains as returned from a list endpoint.
type ApplePayDomainList struct {
	APIResource
	ListMeta
	Data []*ApplePayDomain `json:"data"`
}
