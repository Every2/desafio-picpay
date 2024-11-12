package models

import (
	"github.com/ericlagergren/decimal"
	"time"
)

type Transactions struct {
	ID        int         `json:"id"`
	Amount    decimal.Big `json:"amount"`
	Receiver  []*Users    `json:"receiver,omitempty"`
	Sender    []*Users    `json:"sender,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}
