package entitites

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Bid struct {
	ID        uuid.UUID       `json:"id"`
	ItemID    uuid.UUID       `json:"item_id"`
	BidderID  uuid.UUID       `json:"bidder_id"`
	Amount    decimal.Decimal `json:"bid_amount"`
	CreatedAt time.Time       `json:"bid_time"`
}
