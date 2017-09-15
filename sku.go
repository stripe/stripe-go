package stripe

import "encoding/json"

type SKUParams struct {
	Params            `form:"*"`
	ID                string             `form:"id"`
	Active            *bool              `form:"active"`
	Desc              string             `form:"description"`
	Attrs             map[string]string  `form:"attributes"`
	Price             int64              `form:"price"`
	Currency          string             `form:"currency"`
	Image             string             `form:"image"`
	Inventory         Inventory          `form:"inventory"`
	Product           string             `form:"product"`
	PackageDimensions *PackageDimensions `form:"package_dimensions"`
}

type Inventory struct {
	Type     string `json:"type" form:"type"`
	Quantity int64  `json:"quantity" form:"quantity"`
	Value    string `json:"value" form:"value"`
}

type SKU struct {
	ID                string             `json:"id"`
	Created           int64              `json:"created"`
	Updated           int64              `json:"updated"`
	Live              bool               `json:"livemode"`
	Active            bool               `json:"active"`
	Desc              string             `json:"description"`
	Attrs             map[string]string  `json:"attributes"`
	Price             int64              `json:"price"`
	Currency          string             `json:"currency"`
	PackageDimensions *PackageDimensions `json:"package_dimensions"`
	Image             string             `json:"image"`
	Inventory         Inventory          `json:"inventory"`
	Product           Product            `json:"product"`
	Meta              map[string]string  `json:"metadata"`
}

type SKUList struct {
	ListMeta
	Values []*SKU `json:"data"`
}

type SKUListParams struct {
	ListParams `form:"*"`
	Active     *bool             `form:"active"`
	Product    string            `form:"product"`
	Attributes map[string]string `form:"attributes"`
	IDs        []string          `form:"ids"`
	InStock    *bool             `form:"in_stock"`
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
