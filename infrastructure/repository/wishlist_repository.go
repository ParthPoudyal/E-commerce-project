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

type PostgresWishlistRepository struct {
	db *gorm.DB
}

func NewPostgresWishlistRepository(db *gorm.DB) *PostgresWishlistRepository {
	return &PostgresWishlistRepository{
		db: db,
	}
}

func (r *PostgresWishlistRepository) AddItemToWishlist(item *domain.WishlistItem) error {
	model := mapper.ToWishlistItemModel(item)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("add item to wishlist: %w", result.Error)
	}

	item.ID = model.ID
	item.AddedAt = model.AddedAt
	return nil
}

func (r *PostgresWishlistRepository) UserWishList(userID uuid.UUID) ([]domain.WishlistItem, error) {
	var modelList []*models.WishlistItemModel
	result := r.db.Where("user_id = ?", userID).Order("added_at DESC").Find(&modelList)
	if result.Error != nil {
		return nil, fmt.Errorf("list wishlist for user %s: %w", userID, result.Error)
	}

	domainList := mapper.ToWishlistItemDomainSlice(modelList)
	res := make([]domain.WishlistItem, len(domainList))
	for i, d := range domainList {
		res[i] = *d
	}
	return res, nil
}

func (r *PostgresWishlistRepository) RemoveItem(itemID uuid.UUID) error {
	result := r.db.Where("item_id = ?", itemID).Delete(&models.WishlistItemModel{})
	if result.Error != nil {
		return fmt.Errorf("remove item %s from wishlist: %w", itemID, result.Error)
	}
	return nil
}

func (r *PostgresWishlistRepository) RemoveItemForUser(userID, itemID uuid.UUID) error {
	result := r.db.Where("user_id = ? AND item_id = ?", userID, itemID).Delete(&models.WishlistItemModel{})
	if result.Error != nil {
		return fmt.Errorf("remove item %s for user %s from wishlist: %w", itemID, userID, result.Error)
	}
	return nil
}

func (r *PostgresWishlistRepository) IsWishlisted(userID, itemID uuid.UUID) (bool, error) {
	var count int64
	result := r.db.Model(&models.WishlistItemModel{}).
		Where("user_id = ? AND item_id = ?", userID, itemID).
		Count(&count)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("check if wishlisted: %w", result.Error)
	}
	return count > 0, nil
}

