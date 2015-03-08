package stripe

import "encoding/json"

type SKUParams struct {
	Params
	Active      bool
	Description string
	Name        string
	Attributes  []string
}

type SKU struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Active    bool              `json:"active"`
	Name      string            `json:"name"`
	Desc      string            `json:"description"`
	Attrs     map[string]string `json:"attributes"`
	Price     int64             `json:"price"`
	Currency  string            `json:"currency"`
	Inventory int64             `json:"inventory"`
	Product   Product           `json:"product"`
	Meta      map[string]string `json:"metadata"`
}

type SKUList struct {
	ListMeta
	Values []*SKU `json:"data"`
}

type SKUListParams struct {
	ListParams
	Created int64
}

func (s *SKU) UnmarshalJSON(data []byte) error {
	type sku SKU
	var sk sku
	err := json.Unmarshal(data, &sk)
	if err == nil {
		*s = SKU(sk)
	} else {
		sk.Id = string(data[1 : len(data)-1])
	}

	return nil
}
