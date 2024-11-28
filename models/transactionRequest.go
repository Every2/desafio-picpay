package models

import "github.com/shopspring/decimal"

type TransactionRequest struct {
	SenderID   uint    `json:"sender_id"`
	ReceiverID uint    `json:"receiver_id"`
	Amount     decimal.Decimal `json:"amount"`
}