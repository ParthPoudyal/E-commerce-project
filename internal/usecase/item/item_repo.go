package item

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type ItemRepo interface {
	// Add Item to Item List
	AddItem(item *domain.Item) error 

	// Find All Items 
	ListAllItems (limit , offset int) ([]domain.Item , error)
	
	// Find Items in Auction
	FindAuctionItems (limit , offset int)([]domain.Item , error)
	
	// Find Available Items 
	FindAvailableItems (limit , offset int)([]domain.Item , error)
	
	// Find Item 
	FindItemByID (ID uuid.UUID) (*domain.Item , error)

	// Find Items posted by User 
	FindItemByUser (user uuid.UUID) ([]domain.Item , error)
	
	// Search Items by Name
	SearchItemByUser (name string , limit , offset int) ([]domain.Item, error)
	
	// to Update a Item
	UpdateItem(item *domain.Item) error
	
	// to Delete a Existing Item 
	DeleteItem(id uuid.UUID) error
}