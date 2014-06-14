package stripe

type Balance struct {
	Live      bool     `json:"livemode"`
	Available []Amount `json:"available"`
	Pending   []Amount `json:"pending"`
}

type TransactionStatus string
type TransactionType string

const (
	TxAvailable TransactionStatus = "available"
	TxPending   TransactionStatus = "pending"

	TxCharge         TransactionType = "charge"
	TxRefund         TransactionType = "refund"
	TxAdjust         TransactionType = "adjustment"
	TxFee            TransactionType = "application_fee"
	TxFeeRefund      TransactionType = "application_fee_refund"
	TxTransfer       TransactionType = "transfer"
	TxTransferCancel TransactionType = "transfer_cancel"
	TxTransferFail   TransactionType = "transfer_failure"
)

type Transaction struct {
	Id         string            `json:"id"`
	Amount     int64             `json:"amount"`
	Currency   Currency          `json:"currency"`
	Available  int64             `json:"available_on"`
	Created    int64             `json:"created"`
	Fee        uint64            `json:"fee"`
	FeeDetails []Fee             `json:"fee_details"`
	Net        int64             `json:"net"`
	Status     TransactionStatus `json:"status"`
	Type       TransactionType   `json:"type"`
	Desc       string            `json:"description"`
	Src        string            `json:"source"`
	Recipient  string            `json:"recipient"`
}

type Amount struct {
	Value    int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

type Fee struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Desc        string   `json:"description"`
	Application string   `json:"application"`
}

type BalanceClient struct {
	api   Api
	token string
}

func (c *BalanceClient) Get() (*Balance, error) {
	balance := &Balance{}
	err := c.api.Call("GET", "/balance", c.token, nil, balance)

	return balance, err
}

func (c *BalanceClient) GetTx(id string) (*Transaction, error) {
	balance := &Transaction{}
	err := c.api.Call("GET", "/balance/history/"+id, c.token, nil, balance)

	return balance, err
}
