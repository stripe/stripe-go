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

// Describes an owner of an account.
type FinancialConnectionsAccountOwner struct {
	// The email address of the owner.
	Email string `json:"email"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The full name of the owner.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ownership object that this owner belongs to.
	Ownership string `json:"ownership"`
	// The raw phone number of the owner.
	Phone string `json:"phone"`
	// The raw physical address of the owner.
	RawAddress string `json:"raw_address"`
	// The timestamp of the refresh that updated this owner.
	RefreshedAt time.Time `json:"refreshed_at"`
}

// FinancialConnectionsAccountOwnerList is a list of AccountOwners as retrieved from a list endpoint.
type FinancialConnectionsAccountOwnerList struct {
	APIResource
	ListMeta
	Data []*FinancialConnectionsAccountOwner `json:"data"`
}

// UnmarshalJSON handles deserialization of a FinancialConnectionsAccountOwner.
// This custom unmarshaling is needed to handle the time fields correctly.
func (f *FinancialConnectionsAccountOwner) UnmarshalJSON(data []byte) error {
	type financialConnectionsAccountOwner FinancialConnectionsAccountOwner
	v := struct {
		RefreshedAt int64 `json:"refreshed_at"`
		*financialConnectionsAccountOwner
	}{
		financialConnectionsAccountOwner: (*financialConnectionsAccountOwner)(f),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	f.RefreshedAt = time.Unix(v.RefreshedAt, 0)
	return nil
}

// MarshalJSON handles serialization of a FinancialConnectionsAccountOwner.
// This custom marshaling is needed to handle the time fields correctly.
func (f FinancialConnectionsAccountOwner) MarshalJSON() ([]byte, error) {
	type financialConnectionsAccountOwner FinancialConnectionsAccountOwner
	v := struct {
		RefreshedAt int64 `json:"refreshed_at"`
		financialConnectionsAccountOwner
	}{
		financialConnectionsAccountOwner: (financialConnectionsAccountOwner)(f),
		RefreshedAt:                      f.RefreshedAt.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
