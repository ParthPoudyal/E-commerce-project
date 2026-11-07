package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/mapper"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

type PostgresItemRepository struct {
	db *gorm.DB
}

func NewPostgresItemRepository(db *gorm.DB) *PostgresItemRepository {
	return &PostgresItemRepository{
		db: db,
	}
}

func (r *PostgresItemRepository) AddItem(item *domain.Item) error {
	model := mapper.ToItemModel(item)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("create item: %w", result.Error)
	}

	item.ID = model.ID
	item.CreatedAt = model.CreatedAt
	item.UpdatedAt = model.UpdatedAt

	return nil
}

func (r *PostgresItemRepository) ListAllItems(limit, offset int) ([]domain.Item, error) {
	if limit <= 0 {
		limit = 20
	}
	var itemModels []models.ItemModel
	result := r.db.Where("deleted_at IS NULL").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("list all items: %w", result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) FindAuctionItems(limit, offset int) ([]domain.Item, error) {
	if limit <= 0 {
		limit = 20
	}
	var itemModels []models.ItemModel
	result := r.db.Where("is_auction = ? AND status = ? AND deleted_at IS NULL", true, models.StatusBidding).
		Order("valid_to ASC").
		Limit(limit).
		Offset(offset).
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("find auction items: %w", result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) FindAvailableItems(limit, offset int) ([]domain.Item, error) {
	if limit <= 0 {
		limit = 20
	}
	var itemModels []models.ItemModel
	result := r.db.Where("status = ? AND deleted_at IS NULL", models.StatusAvailable).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("find available items: %w", result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) FindItemByID(id uuid.UUID) (*domain.Item, error) {
	var itemModel models.ItemModel
	result := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&itemModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrItemNotFound
		}
		return nil, fmt.Errorf("find item by id %s: %w", id, result.Error)
	}

	var userModel models.UserModel
	var postedBy *domain.User
	if r.db.Where("id = ?", itemModel.PostedByID).First(&userModel).Error == nil {
		postedBy = mapper.ToUserDomain(&userModel)
	}

	return mapper.ToItemDomain(&itemModel, postedBy), nil
}

func (r *PostgresItemRepository) FindItemByUser(userID uuid.UUID) ([]domain.Item, error) {
	var itemModels []models.ItemModel
	result := r.db.Where("posted_by_id = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("find items by user %s: %w", userID, result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) SearchItemByUser(name string, limit, offset int) ([]domain.Item, error) {
	if limit <= 0 {
		limit = 20
	}
	var itemModels []models.ItemModel
	query := "%" + name + "%"
	result := r.db.Where("(name ILIKE ? OR description ILIKE ?) AND deleted_at IS NULL", query, query).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("search items: %w", result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) FindItemsByCategory(categoryID uuid.UUID, limit, offset int) ([]domain.Item, error) {
	if limit <= 0 {
		limit = 20
	}
	var itemModels []models.ItemModel
	result := r.db.Where("category_id = ? AND deleted_at IS NULL", categoryID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&itemModels)

	if result.Error != nil {
		return nil, fmt.Errorf("find items by category %s: %w", categoryID, result.Error)
	}

	items := make([]domain.Item, len(itemModels))
	for i, m := range itemModels {
		items[i] = *mapper.ToItemDomain(&m, nil)
	}
	return items, nil
}

func (r *PostgresItemRepository) UpdateItem(item *domain.Item) error {
	model := mapper.ToItemModel(item)
	result := r.db.Model(&models.ItemModel{}).Where("id = ? AND deleted_at IS NULL", item.ID).Updates(map[string]interface{}{
		"name":         model.Name,
		"description":  model.Description,
		"category_id":  model.CategoryID,
		"image_url":    model.ImageURL,
		"price":        model.Price,
		"status":       model.Status,
		"valid_from":   model.ValidFrom,
		"valid_to":     model.ValidTo,
		"is_auction":   model.IsAuction,
		"starting_bid": model.StartingBid,
		"sold_price":   model.SoldPrice,
		"sold_at":      model.SoldAt,
		"updated_at":   time.Now(),
	})

	if result.Error != nil {
		return fmt.Errorf("update item %s: %w", item.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrItemNotFound
	}
	return nil
}

func (r *PostgresItemRepository) DeleteItem(id uuid.UUID) error {
	now := time.Now()
	result := r.db.Model(&models.ItemModel{}).Where("id = ?", id).Update("deleted_at", &now)
	if result.Error != nil {
		return fmt.Errorf("delete item %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrItemNotFound
	}
	return nil
}

