package domain


import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ItemStatus string

const (
	UNAVAILABLE ItemStatus = "unavailable"
	AVAILABE    ItemStatus = "available"
	SOLD        ItemStatus = "sold"
	BIDDING     ItemStatus = "bidding" // for ongoing bidding
)

type Item struct {
	//item description
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	PostedBy    *User      `json:"posted_by"` // * : use foreign key in the model
	Price       decimal.Decimal   `json:"price"`
	Status      ItemStatus `json:"status"`

	//item validity
	ValidFrom time.Time  `json:"valid_from"`
	ValidTo   *time.Time `json:"valid_to"`

	//bidding
	IsAuction    bool       `json:"is_auction"`
	StartingBid  *decimal.Decimal   `json:"starting_bid"`
	CurrentBidID *uuid.UUID `json:"current_bid"` // denormalized pointer to winning Bid, for fast reads
	SoldPrice    *decimal.Decimal  `json:"sold_price"`
	SoldAt       *time.Time `json:"sold_at"`

	//base
	Base
}
