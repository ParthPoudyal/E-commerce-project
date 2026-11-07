package category

import (
	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/google/uuid"
)

type CategoryRepo interface {
	AddCategory(category *domain.Category) error
	ListCategories() ([]*domain.Category, error)
	FindCategoryBySlug(slug string) (*domain.Category, error)
	FindCategoryByID(id uuid.UUID) (*domain.Category, error)
	UpdateCategory(category *domain.Category) error
	DeleteCategory(id uuid.UUID) error
}

