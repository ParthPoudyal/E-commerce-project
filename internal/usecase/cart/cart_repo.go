package cart

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type CartRepo interface {
	// Add Item to Cart
	AddItemToCart(item *domain.CartItem) error  

	// Look for Item in User's Cart
	FindItemInCart(userID, itemID uuid.UUID) (*domain.CartItem, error)
	
	// See all Item in User's Cart
	ListUserCart (userID uuid.UUID) ([]domain.CartItem , error)

	// Remove Item from Cart
	RemoveFromCart(userID, itemID uuid.UUID) error

	// To clear User's Cart
	ClearCart(userID uuid.UUID) error
}