package models

import (
	"time"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ItemStatus string

const (
	StatusUnavailable ItemStatus = "unavailable"
	StatusAvailable   ItemStatus = "available"
	StatusSold        ItemStatus = "sold"
	StatusBidding     ItemStatus = "bidding"
)

type ItemModel struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"not null"`
	Description string
	PostedByID  uuid.UUID  `gorm:"type:uuid;not null;index"`
	Price       decimal.Decimal `gorm:"type:numeric(12,2);not null"` 
	Status      ItemStatus `gorm:"type:varchar(20);not null;default:'available'"`

	ValidFrom time.Time  `gorm:"not null"`
	ValidTo   *time.Time 

	IsAuction   bool             `gorm:"not null;default:false"`
	StartingBid *decimal.Decimal `gorm:"type:numeric(12,2)"`

	SoldPrice *decimal.Decimal `gorm:"type:numeric(12,2)"`
	SoldAt    *time.Time

	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}

func (ItemModel) TableName() string { 
	return "items" 
}