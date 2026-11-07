package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToBidDomain(model *models.BidModel) *domain.Bid {
	if model == nil {
		return nil
	}

	return &domain.Bid{
		ID:        model.ID,
		ItemID:    model.ItemID,
		BidderID:  model.BidderID,
		Amount:    model.Amount,
		CreatedAt: model.CreatedAt,
	}
}

func ToBidModel(entity *domain.Bid) *models.BidModel {
	if entity == nil {
		return nil
	}

	return &models.BidModel{
		ID:        entity.ID,
		ItemID:    entity.ItemID,
		BidderID:  entity.BidderID,
		Amount:    entity.Amount,
		CreatedAt: entity.CreatedAt,
	}
}

func ToBidDomainSlice(modelSlice []*models.BidModel) []*domain.Bid {
	if modelSlice == nil {
		return nil
	}

	domainSlice := make([]*domain.Bid, len(modelSlice))
	for i, m := range modelSlice {
		domainSlice[i] = ToBidDomain(m)
	}
	return domainSlice
}
