package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/mapper"
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

type PostgresPurchaseRepository struct {
	db *gorm.DB
}

func NewPostgresPurchaseRepository(db *gorm.DB) *PostgresPurchaseRepository {
	return &PostgresPurchaseRepository{
		db: db,
	}
}

func (r *PostgresPurchaseRepository) AddPurchase(purchase *domain.Purchase) error {
	model := mapper.ToPurchaseModel(purchase)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("add purchase: %w", result.Error)
	}

	purchase.ID = model.ID
	purchase.PurchasedAt = model.PurchasedAt
	return nil
}

func (r *PostgresPurchaseRepository) ListPurchasesByUser(userID uuid.UUID) ([]domain.Purchase, error) {
	var modelList []*models.PurchaseModel
	result := r.db.Where("user_id = ?", userID).Order("purchased_at DESC").Find(&modelList)
	if result.Error != nil {
		return nil, fmt.Errorf("list purchases by user %s: %w", userID, result.Error)
	}

	domainList := mapper.ToPurchaseDomainSlice(modelList)
	res := make([]domain.Purchase, len(domainList))
	for i, d := range domainList {
		res[i] = *d
	}
	return res, nil
}

func (r *PostgresPurchaseRepository) ListPurchasesbyDate(userID uuid.UUID, date time.Time) ([]domain.Purchase, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var modelList []*models.PurchaseModel
	result := r.db.Where("user_id = ? AND purchased_at >= ? AND purchased_at < ?", userID, startOfDay, endOfDay).
		Order("purchased_at DESC").
		Find(&modelList)

	if result.Error != nil {
		return nil, fmt.Errorf("list purchases by date: %w", result.Error)
	}

	domainList := mapper.ToPurchaseDomainSlice(modelList)
	res := make([]domain.Purchase, len(domainList))
	for i, d := range domainList {
		res[i] = *d
	}
	return res, nil
}

func (r *PostgresPurchaseRepository) ListPurchasesDuring(userID uuid.UUID, timeA, timeB time.Time) ([]domain.Purchase, error) {
	var modelList []*models.PurchaseModel
	result := r.db.Where("user_id = ? AND purchased_at >= ? AND purchased_at <= ?", userID, timeA, timeB).
		Order("purchased_at DESC").
		Find(&modelList)

	if result.Error != nil {
		return nil, fmt.Errorf("list purchases during interval: %w", result.Error)
	}

	domainList := mapper.ToPurchaseDomainSlice(modelList)
	res := make([]domain.Purchase, len(domainList))
	for i, d := range domainList {
		res[i] = *d
	}
	return res, nil
}

