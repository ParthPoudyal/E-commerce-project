package bid

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type BidRepository interface {
	// To Create a Bid
	Create(bid *domain.Bid) error

	// Get Highest Bid for a Item in Bidding
	GetHighestBidForItem(itemID uuid.UUID) (*domain.Bid, error)

	// Get All the Past Bids made for a Item
	GetBidHistoryForItem(itemID uuid.UUID) ([]*domain.Bid, error)

	// Get all the bids placed by a User
	GetBidsByUser(userID uuid.UUID) ([]*domain.Bid, error)
}
