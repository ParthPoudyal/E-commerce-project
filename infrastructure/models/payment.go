package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentTransactionModel struct {
	ID            uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PurchaseID    *uuid.UUID      `gorm:"type:uuid;index"`
	ItemID        uuid.UUID       `gorm:"type:uuid;not null;index"`
	UserID        uuid.UUID       `gorm:"type:uuid;not null;index"`
	Amount        decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	Provider      string          `gorm:"type:varchar(50);not null;default:'khalti'"`
	ProviderTxnID string          `gorm:"type:varchar(255);index"`
	Status        string          `gorm:"type:varchar(50);not null;default:'pending'"`
	PaymentURL    string          `gorm:"type:text"`
	CreatedAt     time.Time       `gorm:"autoCreateTime"`
	CompletedAt   *time.Time
}

func (PaymentTransactionModel) TableName() string {
	return "payment_transactions"
}

