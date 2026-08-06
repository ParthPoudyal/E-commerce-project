package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToWishlistItemDomain(model *models.WishlistItemModel) *domain.WishlistItem {
	if model == nil {
		return nil
	}

	return &domain.WishlistItem{
		ID:      model.ID,
		UserId:  model.UserID,
		ItemId:  model.ItemID,
		AddedAt: model.AddedAt,
	}
}

func ToWishlistItemModel(entity *domain.WishlistItem) *models.WishlistItemModel {
	if entity == nil {
		return nil
	}

	return &models.WishlistItemModel{
		ID:      entity.ID,
		UserID:  entity.UserId,
		ItemID:  entity.ItemId,
		AddedAt: entity.AddedAt,
	}
}

func ToWishlistItemDomainSlice(modelSlice []*models.WishlistItemModel) []*domain.WishlistItem {
	if modelSlice == nil {
		return nil
	}

	domainSlice := make([]*domain.WishlistItem, len(modelSlice))
	for i, m := range modelSlice {
		domainSlice[i] = ToWishlistItemDomain(m)
	}
	return domainSlice
}
