package entitites

import (
	"time"

	"github.com/google/uuid"
)

type WishlistItem struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id"`
	ItemId uuid.UUID `json:"item_id"`
	AddedAt time.Time `json:"added_at"`
}
