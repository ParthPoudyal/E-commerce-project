package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type BidModel struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ItemID    uuid.UUID       `gorm:"type:uuid;not null;index"`
	BidderID  uuid.UUID       `gorm:"type:uuid;not null;index"`
	Amount    decimal.Decimal `gorm:"type:numeric(12,2);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
}

func (BidModel) TableName() string { 
	return "bids"
}