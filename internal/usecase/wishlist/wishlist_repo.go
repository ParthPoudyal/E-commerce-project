package wishlist

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type WishlistRepo interface {
	//Add Item to Wishlist
	AddItemToWishlist(item *domain.WishlistItem) error

	// Retrive all Items in wishlist for a User
	UserWishList (user uuid.UUID) ([]domain.WishlistItem , error) 

	// Remove a wishlist Item
	RemoveItem (itemID uuid.UUID) error
	
	// To check if it is wishlisted by a user 
	IsWishlisted(userID, itemID uuid.UUID) (bool, error)
}