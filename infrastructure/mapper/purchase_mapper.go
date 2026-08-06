package mapper

import (
	"github.com/shopspring/decimal"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToPurchaseDomain(model *models.PurchaseModel) *domain.Purchase {
	if model == nil {
		return nil
	}

	return &domain.Purchase{
		ID:          model.ID,
		UserID:      model.UserID,
		ItemID:      model.ItemID,
		PricePaid:   model.PricePaid.InexactFloat64(),
		PurchasedAt: model.PurchasedAt,
	}
}

func ToPurchaseModel(entity *domain.Purchase) *models.PurchaseModel {
	if entity == nil {
		return nil
	}

	return &models.PurchaseModel{
		ID:          entity.ID,
		UserID:      entity.UserID,
		ItemID:      entity.ItemID,
		PricePaid:   decimal.NewFromFloat(entity.PricePaid),
		PurchasedAt: entity.PurchasedAt,
	}
}

func ToPurchaseDomainSlice(modelSlice []*models.PurchaseModel) []*domain.Purchase {
	if modelSlice == nil {
		return nil
	}

	domainSlice := make([]*domain.Purchase, len(modelSlice))
	for i, m := range modelSlice {
		domainSlice[i] = ToPurchaseDomain(m)
	}
	return domainSlice
}
