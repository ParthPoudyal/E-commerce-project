package models

import (
	"time"
	"github.com/google/uuid"
)

type WishlistItemModel struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID  uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_wishlist_user_item"`
	ItemID  uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_wishlist_user_item"`
	AddedAt time.Time `gorm:"autoCreateTime"`
}

func (WishlistItemModel) TableName() string {
	 return "wishlist_items"
}