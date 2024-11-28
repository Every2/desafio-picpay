package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Amount    decimal.Decimal `gorm:"type:decimal(10,2)" json:"amount"`
	SenderID  uint     `json:"sender_id"`
	Sender    User     `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE" json:"sender"`
	ReceiverID uint    `json:"receiver_id"`
	Receiver  User     `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE" json:"receiver"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

func NewTransaction(amount decimal.Decimal, sender User, receiver User) *Transaction {
	return &Transaction{
		Amount:   amount,
		Sender:   sender,
		Receiver: receiver,
		Timestamp: time.Now(),
	}
}