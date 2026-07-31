package domain

import (
	"time"

	"github.com/google/uuid"
)

// will keep track of who bought what
type Purchase struct {
	ID          uuid.UUID  `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ItemID      uuid.UUID  `json:"item_id"`
	PricePaid   float64    `json:"price_paid"`
	PurchasedAt time.Time  `json:"purchased_at"`
}
