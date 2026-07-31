package bid

import (
	"errors"
	"fmt"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

type PostgresBidRepository struct {
	db *gorm.DB
}

func NewPostgresBidRepository(db *gorm.DB) *PostgresBidRepository {
	return &PostgresBidRepository{
		db: db,
	}
}

func (r *PostgresBidRepository) Create(bid *domain.Bid) error {
	result := r.db.Create(bid)

	if result.Error != nil {
		return fmt.Errorf("create bid: %d", result.Error)
	}

	return nil
}

func (r *PostgresBidRepository) GetHighestBidForItem(itemID uuid.UUID) (*domain.Bid, error) {
	var highestBid domain.Bid

	result := r.db.
		Where("item_id = ?", itemID).
		Order("amount DESC").
		Order("created_at ASC").
		First(&highestBid)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(
				"no bids found for item %s: %w",
				itemID,
				result.Error,
			)
		}

		return nil, fmt.Errorf(
			"get highest bid for item %s: %w",
			itemID,
			result.Error,
		)
	}

	return &highestBid, nil
}

func (r *PostgresBidRepository) GetBidHistoryForItem(itemID uuid.UUID) ([]*domain.Bid, error) {
	var bids []*domain.Bid

	result := r.db.
		Where("item_id = ?", itemID).
		Order("created_at DESC").
		Find(&bids)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"get bid history for item %s: %w",
			itemID,
			result.Error,
		)
	}

	return bids, nil
}

func (r *PostgresBidRepository) GetBidsByUser(userID uuid.UUID) ([]*domain.Bid, error) {
	var bids []*domain.Bid

	result := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&bids)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"get bids for user %s: %w",
			userID,
			result.Error,
		)
	}

	return bids, nil
}
