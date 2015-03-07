package stripe

import "encoding/json"

type ProductParams struct {
	Params
	Active      bool
	Description string
	Name        string
	Attributes  []string
}

type Product struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Active    bool              `json:"active"`
	Desc      string            `json:"description"`
	Attrs     []string          `json:"attributes"`
	Shippable bool              `json:"shippable"`
	Meta      map[string]string `json:"metdata"`
	//Skus      []Sku          `json:"skus"`
}

type ProductListParams struct {
	ListParams
	Created int64
}

func (p *Product) UnmarshalJSON(data []byte) error {
	type product Product
	var pr product
	err := json.Unmarshal(data, &pr)
	if err == nil {
		*p = Product(pr)
	} else {
		p.Id = string(data[1 : len(data)-1])
	}

	return nil
}
