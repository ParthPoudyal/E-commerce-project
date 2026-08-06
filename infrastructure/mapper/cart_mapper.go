package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToCartItemDomain(model *models.CartItemModel) *domain.CartItem {
	if model == nil {
		return nil
	}

	return &domain.CartItem{
		Id:      model.ID,
		ItemID:  model.ItemID,
		UserID:  model.UserID,
		AddedAt: model.AddedAt,
	}
}

func ToCartItemModel(entity *domain.CartItem) *models.CartItemModel {
	if entity == nil {
		return nil
	}

	return &models.CartItemModel{
		ID:      entity.Id,
		UserID:  entity.UserID,
		ItemID:  entity.ItemID,
		AddedAt: entity.AddedAt,
	}
}

func ToCartItemDomainSlice(modelSlice []*models.CartItemModel) []*domain.CartItem {
	if modelSlice == nil {
		return nil
	}

	domainSlice := make([]*domain.CartItem, len(modelSlice))
	for i, m := range modelSlice {
		domainSlice[i] = ToCartItemDomain(m)
	}
	return domainSlice
}
