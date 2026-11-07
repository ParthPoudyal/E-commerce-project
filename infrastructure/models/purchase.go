package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PurchaseModel struct {
	ID           uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID       uuid.UUID       `gorm:"type:uuid;not null;index"` // buyer
	ItemID       uuid.UUID       `gorm:"type:uuid;not null;index"`
	PricePaid    decimal.Decimal `gorm:"type:numeric(12,2);not null"` // snapshot — item.Price may change/be deleted later
	PurchasedAt  time.Time       `gorm:"autoCreateTime"`
}

func (PurchaseModel) TableName() string {
	return "purchases"
}