package entitites

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	id      uuid.UUID
	ItemID  uuid.UUID
	UserID  uuid.UUID
	AddedAt time.Time
}

