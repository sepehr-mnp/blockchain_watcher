package dtos

import "errors"

var (
	InvalidNetwork = errors.New("Invalid Network!")
)

type DepositePaymentBody struct {
	PaymentID int64  `json:"payment_id"`
	Amount    string `json:"amount"`
	Symbol    string `json:"symbol"`
}

type Payment struct {
	ID             int64   `json:"id"`
	Network        string  `json:"network"`
	Symbol         string  `json:"symbol"`
	Status         string  `json:"status"`
	Address        string  `json:"address"`
	Amount         string  `json:"amount"`
	ReceivedAmount string  `json:"received_amount"`
	FinalAmount    string  `json:"final_amount"`
	PercentFee     float64 `json:"percent_fee"`
	Fee            string  `json:"fee"`
	UserID         int64   `json:"user_id"`
	OrderID        string  `json:"order_id"`
	Description    string  `json:"Description"`
	ExpireTime     string  `json:"expire_time"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type Transaction struct {
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	PaymentID   uint    `json:"payment_id"`
	ToAddress   string  `json:"to_address"`
	FromAddress string  `json:"from_address"`
	TXHash      string  `json:"tx_hash"`
	BlockNumber int64   `json:"block_number"`
	Amount      float64 `json:"amount"`
	Symbol      string  `json:"symbol"`
	Confirmed   bool    `json:"confirmed"`
	Orphaned    bool    `json:"orphaned"`
	Deposit     bool    `json:"deposit"`
}
