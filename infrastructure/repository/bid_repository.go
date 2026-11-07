package repository

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/mapper"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

type PostgresBidRepository struct {
	db *gorm.DB
}

func NewPostgresBidRepository(db *gorm.DB) *PostgresBidRepository {
	return &PostgresBidRepository{
		db: db,
	}
}

func (r *PostgresBidRepository) Create(bid *domain.Bid) error {
	bidModel := mapper.ToBidModel(bid)

	result := r.db.Create(bidModel)
	if result.Error != nil {
		return fmt.Errorf("create bid: %w", result.Error)
	}

	// Update domain entity with DB-generated fields
	bid.ID = bidModel.ID
	bid.CreatedAt = bidModel.CreatedAt

	return nil
}

func (r *PostgresBidRepository) GetHighestBidForItem(itemID uuid.UUID) (*domain.Bid, error) {
	var highestBidModel models.BidModel

	result := r.db.
		Where("item_id = ?", itemID).
		Order("amount DESC").
		Order("created_at ASC").
		First(&highestBidModel)

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

	return mapper.ToBidDomain(&highestBidModel), nil
}

func (r *PostgresBidRepository) GetBidHistoryForItem(itemID uuid.UUID) ([]*domain.Bid, error) {
	var bidModels []*models.BidModel

	result := r.db.
		Where("item_id = ?", itemID).
		Order("created_at DESC").
		Find(&bidModels)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"get bid history for item %s: %w",
			itemID,
			result.Error,
		)
	}

	return mapper.ToBidDomainSlice(bidModels), nil
}

func (r *PostgresBidRepository) GetBidsByUser(userID uuid.UUID) ([]*domain.Bid, error) {
	var bidModels []*models.BidModel

	result := r.db.
		Where("bidder_id = ?", userID).
		Order("created_at DESC").
		Find(&bidModels)

	if result.Error != nil {
		return nil, fmt.Errorf(
			"get bids for user %s: %w",
			userID,
			result.Error,
		)
	}

	return mapper.ToBidDomainSlice(bidModels), nil
}