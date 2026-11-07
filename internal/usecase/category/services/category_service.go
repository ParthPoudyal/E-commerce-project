package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/ParthPoudyal/E-commerce-project/internal/domain"
	"github.com/ParthPoudyal/E-commerce-project/internal/usecase/category"
)

type CategoryService struct {
	categoryRepo category.CategoryRepo
}

func NewCategoryService(cR category.CategoryRepo) *CategoryService {
	return &CategoryService{
		categoryRepo: cR,
	}
}

func (s *CategoryService) ListCategories() ([]*domain.Category, error) {
	return s.categoryRepo.ListCategories()
}

func (s *CategoryService) GetCategoryBySlug(slug string) (*domain.Category, error) {
	return s.categoryRepo.FindCategoryBySlug(slug)
}

func (s *CategoryService) GetCategoryByID(id uuid.UUID) (*domain.Category, error) {
	return s.categoryRepo.FindCategoryByID(id)
}

func (s *CategoryService) CreateCategory(name, description, imageURL string) (*domain.Category, error) {
	slug := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", "-"))
	existing, _ := s.categoryRepo.FindCategoryBySlug(slug)
	if existing != nil {
		return nil, domain.ErrCategorySlugExists
	}

	cat := &domain.Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug,
		Description: description,
		ImageURL:    imageURL,
	}

	if err := s.categoryRepo.AddCategory(cat); err != nil {
		return nil, fmt.Errorf("create category: %w", err)
	}

	return cat, nil
}

