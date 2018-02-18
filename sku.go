package stripe

import "encoding/json"

type InventoryParams struct {
	Quantity *int64  `form:"quantity"`
	Type     *string `form:"type"`
	Value    *string `form:"value"`
}

type SKUParams struct {
	Params            `form:"*"`
	Active            *bool                    `form:"active"`
	Attributes        map[string]string        `form:"attributes"`
	Currency          *string                  `form:"currency"`
	Description       *string                  `form:"description"`
	ID                *string                  `form:"id"`
	Image             *string                  `form:"image"`
	Inventory         *InventoryParams         `form:"inventory"`
	PackageDimensions *PackageDimensionsParams `form:"package_dimensions"`
	Price             *int64                   `form:"price"`
	Product           *string                  `form:"product"`
}

type Inventory struct {
	Quantity int64  `json:"quantity"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type SKU struct {
	Active            bool               `json:"active"`
	Attributes        map[string]string  `json:"attributes"`
	Created           int64              `json:"created"`
	Currency          Currency           `json:"currency"`
	Description       string             `json:"description"`
	ID                string             `json:"id"`
	Image             string             `json:"image"`
	Inventory         Inventory          `json:"inventory"`
	Livemode          bool               `json:"livemode"`
	Metadata          map[string]string  `json:"metadata"`
	PackageDimensions *PackageDimensions `json:"package_dimensions"`
	Price             int64              `json:"price"`
	Product           Product            `json:"product"`
	Updated           int64              `json:"updated"`
}

type SKUList struct {
	ListMeta
	Data []*SKU `json:"data"`
}

type SKUListParams struct {
	ListParams `form:"*"`
	Active     *bool             `form:"active"`
	Attributes map[string]string `form:"attributes"`
	IDs        []string          `form:"ids"`
	InStock    *bool             `form:"in_stock"`
	Product    *string           `form:"product"`
}

func (s *SKU) UnmarshalJSON(data []byte) error {
	type sku SKU
	var sk sku
	err := json.Unmarshal(data, &sk)
	if err == nil {
		*s = SKU(sk)
	} else {
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}
