package domain


import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	Id      uuid.UUID
	ItemID  uuid.UUID
	UserID  uuid.UUID
	AddedAt time.Time
}

