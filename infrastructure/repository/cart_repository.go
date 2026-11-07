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

type PostgresCartRepository struct {
	db *gorm.DB
}

func NewPostgresCartRepository(db *gorm.DB) *PostgresCartRepository {
	return &PostgresCartRepository{
		db: db,
	}
}

func (r *PostgresCartRepository) AddItemToCart(item *domain.CartItem) error {
	model := mapper.ToCartItemModel(item)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("add item to cart: %w", result.Error)
	}

	item.Id = model.ID
	item.AddedAt = model.AddedAt
	return nil
}

func (r *PostgresCartRepository) FindItemInCart(userID, itemID uuid.UUID) (*domain.CartItem, error) {
	var model models.CartItemModel
	result := r.db.Where("user_id = ? AND item_id = ?", userID, itemID).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrItemNotInCart
		}
		return nil, fmt.Errorf("find item in cart: %w", result.Error)
	}
	return mapper.ToCartItemDomain(&model), nil
}

func (r *PostgresCartRepository) ListUserCart(userID uuid.UUID) ([]domain.CartItem, error) {
	var modelList []*models.CartItemModel
	result := r.db.Where("user_id = ?", userID).Order("added_at DESC").Find(&modelList)
	if result.Error != nil {
		return nil, fmt.Errorf("list user cart: %w", result.Error)
	}

	domainList := mapper.ToCartItemDomainSlice(modelList)
	res := make([]domain.CartItem, len(domainList))
	for i, d := range domainList {
		res[i] = *d
	}
	return res, nil
}

func (r *PostgresCartRepository) RemoveFromCart(userID, itemID uuid.UUID) error {
	result := r.db.Where("user_id = ? AND item_id = ?", userID, itemID).Delete(&models.CartItemModel{})
	if result.Error != nil {
		return fmt.Errorf("remove from cart: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrItemNotInCart
	}
	return nil
}

func (r *PostgresCartRepository) ClearCart(userID uuid.UUID) error {
	result := r.db.Where("user_id = ?", userID).Delete(&models.CartItemModel{})
	if result.Error != nil {
		return fmt.Errorf("clear cart for user %s: %w", userID, result.Error)
	}
	return nil
}

