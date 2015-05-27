package stripe

import "encoding/json"

type ProductParams struct {
	Params
	ID        string
	Active    *bool
	Name      string
	Caption   string
	Desc      string
	Attrs     []string
	Shippable *bool
}

type Product struct {
	ID        string            `json:"id"`
	Created   int64             `json:"created"`
	Updated   int64             `json:"updated"`
	Live      bool              `json:"livemode"`
	Active    bool              `json:"active"`
	Name      string            `json:"name"`
	Caption   string            `json:"caption"`
	Desc      string            `json:"description"`
	Attrs     []string          `json:"attributes"`
	Shippable bool              `json:"shippable"`
	Meta      map[string]string `json:"metdata"`
	Skus      *SKUList          `json:"skus"`
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
		p.ID = string(data[1 : len(data)-1])
	}

	return nil
}
