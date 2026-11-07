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

type PostgresCategoryRepository struct {
	db *gorm.DB
}

func NewPostgresCategoryRepository(db *gorm.DB) *PostgresCategoryRepository {
	return &PostgresCategoryRepository{
		db: db,
	}
}

func (r *PostgresCategoryRepository) AddCategory(category *domain.Category) error {
	model := mapper.ToCategoryModel(category)
	if model.ID == uuid.Nil {
		model.ID = uuid.New()
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("create category: %w", result.Error)
	}

	category.ID = model.ID
	return nil
}

func (r *PostgresCategoryRepository) ListCategories() ([]*domain.Category, error) {
	var modelList []*models.CategoryModel
	result := r.db.Order("name ASC").Find(&modelList)
	if result.Error != nil {
		return nil, fmt.Errorf("list categories: %w", result.Error)
	}
	return mapper.ToCategoryDomainSlice(modelList), nil
}

func (r *PostgresCategoryRepository) FindCategoryBySlug(slug string) (*domain.Category, error) {
	var model models.CategoryModel
	result := r.db.Where("slug = ?", slug).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, fmt.Errorf("find category by slug %s: %w", slug, result.Error)
	}
	return mapper.ToCategoryDomain(&model), nil
}

func (r *PostgresCategoryRepository) FindCategoryByID(id uuid.UUID) (*domain.Category, error) {
	var model models.CategoryModel
	result := r.db.Where("id = ?", id).First(&model)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCategoryNotFound
		}
		return nil, fmt.Errorf("find category by id %s: %w", id, result.Error)
	}
	return mapper.ToCategoryDomain(&model), nil
}

func (r *PostgresCategoryRepository) UpdateCategory(category *domain.Category) error {
	model := mapper.ToCategoryModel(category)
	result := r.db.Model(&models.CategoryModel{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"name":        model.Name,
		"slug":        model.Slug,
		"description": model.Description,
		"image_url":   model.ImageURL,
	})
	if result.Error != nil {
		return fmt.Errorf("update category %s: %w", category.ID, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrCategoryNotFound
	}
	return nil
}

func (r *PostgresCategoryRepository) DeleteCategory(id uuid.UUID) error {
	result := r.db.Where("id = ?", id).Delete(&models.CategoryModel{})
	if result.Error != nil {
		return fmt.Errorf("delete category %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrCategoryNotFound
	}
	return nil
}

