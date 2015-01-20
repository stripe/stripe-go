package stripe

type SourceParams struct {
	Params
	ID    string
	Token string
	Card  *CardParams
}

type PaymentSource struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

type Chargeable interface {
	PaymentType() string
}
