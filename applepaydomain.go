//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Delete an apple pay domain.
type ApplePayDomainParams struct {
	Params     `form:"*"`
	DomainName *string `form:"domain_name"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// List apple pay domains.
type ApplePayDomainListParams struct {
	ListParams `form:"*"`
	DomainName *string `form:"domain_name"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ApplePayDomainListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ApplePayDomain struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created    time.Time `json:"created"`
	Deleted    bool      `json:"deleted"`
	DomainName string    `json:"domain_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// ApplePayDomainList is a list of ApplePayDomains as retrieved from a list endpoint.
type ApplePayDomainList struct {
	APIResource
	ListMeta
	Data []*ApplePayDomain `json:"data"`
}

// UnmarshalJSON handles deserialization of an ApplePayDomain.
// This custom unmarshaling is needed to handle the time fields correctly.
func (a *ApplePayDomain) UnmarshalJSON(data []byte) error {
	type applePayDomain ApplePayDomain
	v := struct {
		Created int64 `json:"created"`
		*applePayDomain
	}{
		applePayDomain: (*applePayDomain)(a),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	a.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of an ApplePayDomain.
// This custom marshaling is needed to handle the time fields correctly.
func (a ApplePayDomain) MarshalJSON() ([]byte, error) {
	type applePayDomain ApplePayDomain
	v := struct {
		Created int64 `json:"created"`
		applePayDomain
	}{
		applePayDomain: (applePayDomain)(a),
		Created:        a.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
