package mapper

import (
	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToItemStatusDomain(status models.ItemStatus) domain.ItemStatus {
	switch status {
	case models.StatusUnavailable:
		return domain.UNAVAILABLE
	case models.StatusAvailable:
		return domain.AVAILABE
	case models.StatusSold:
		return domain.SOLD
	case models.StatusBidding:
		return domain.BIDDING
	default:
		return domain.ItemStatus(status)
	}
}

func ToItemStatusModel(status domain.ItemStatus) models.ItemStatus {
	switch status {
	case domain.UNAVAILABLE:
		return models.StatusUnavailable
	case domain.AVAILABE:
		return models.StatusAvailable
	case domain.SOLD:
		return models.StatusSold
	case domain.BIDDING:
		return models.StatusBidding
	default:
		return models.ItemStatus(status)
	}
}

func ToItemDomain(model *models.ItemModel, postedBy *domain.User) *domain.Item {
	if model == nil {
		return nil
	}

	return &domain.Item{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		PostedBy:    postedBy,
		Price:       model.Price,
		Status:      ToItemStatusDomain(model.Status),
		ValidFrom:   model.ValidFrom,
		ValidTo:     model.ValidTo,
		IsAuction:   model.IsAuction,
		StartingBid: model.StartingBid,
		SoldPrice:   model.SoldPrice,
		SoldAt:      model.SoldAt,
		Base: domain.Base{
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: model.DeletedAt,
		},
	}
}

func ToItemModel(entity *domain.Item) *models.ItemModel {
	if entity == nil {
		return nil
	}

	var postedByID uuid.UUID
	if entity.PostedBy != nil {
		postedByID = entity.PostedBy.ID
	}

	return &models.ItemModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		PostedByID:  postedByID,
		Price:       entity.Price,
		Status:      ToItemStatusModel(entity.Status),
		ValidFrom:   entity.ValidFrom,
		ValidTo:     entity.ValidTo,
		IsAuction:   entity.IsAuction,
		StartingBid: entity.StartingBid,
		SoldPrice:   entity.SoldPrice,
		SoldAt:      entity.SoldAt,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
	}
}
