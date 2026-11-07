package mapper

import (
	"github.com/ParthPoudyal/E-commerce-project/infrastructure/models"
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
)

func ToCategoryDomain(model *models.CategoryModel) *domain.Category {
	if model == nil {
		return nil
	}

	return &domain.Category{
		ID:          model.ID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		ImageURL:    model.ImageURL,
	}
}

func ToCategoryModel(entity *domain.Category) *models.CategoryModel {
	if entity == nil {
		return nil
	}

	return &models.CategoryModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Slug:        entity.Slug,
		Description: entity.Description,
		ImageURL:    entity.ImageURL,
	}
}

func ToCategoryDomainSlice(models []*models.CategoryModel) []*domain.Category {
	if models == nil {
		return nil
	}

	domainList := make([]*domain.Category, len(models))
	for i, m := range models {
		domainList[i] = ToCategoryDomain(m)
	}
	return domainList
}

