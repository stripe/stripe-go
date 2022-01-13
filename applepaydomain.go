//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List apple pay domains.
type ApplePayDomainListParams struct {
	ListParams `form:"*"`
	DomainName *string `form:"domain_name"`
}

// Create an apple pay domain.
type ApplePayDomainParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name"`
}
type ApplePayDomain struct {
	APIResource
	Created    int64  `json:"created"`
	Deleted    bool   `json:"deleted"`
	DomainName string `json:"domain_name"`
	ID         string `json:"id"`
	Livemode   bool   `json:"livemode"`
	Object     string `json:"object"`
}

// ApplePayDomainList is a list of ApplePayDomains as retrieved from a list endpoint.
type ApplePayDomainList struct {
	APIResource
	ListMeta
	Data []*ApplePayDomain `json:"data"`
}
