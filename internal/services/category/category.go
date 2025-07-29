package category

import (
	"context"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/internal/domain/repository"
)

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

//go:generate mockgen -source=category.go -destination=mocks/mock_category.go -package=mocks
type CategoryService interface {
	Create(ctx context.Context, category entity.Category) (string, error)
	Update(ctx context.Context, id string, category entity.Category) error
	GetByID(ctx context.Context, id string) (*entity.Category, error)
	GetAll(ctx context.Context, filter entity.CategoriesFilter) ([]entity.Category, error)
	Delete(ctx context.Context, id string) error
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (p categoryService) Create(ctx context.Context, category entity.Category) (string, error) {
	id, err := p.categoryRepo.Create(ctx, &category)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (p categoryService) Update(ctx context.Context, id string, category entity.Category) error {
	category.ID = id

	return p.categoryRepo.Update(ctx, &category)
}

func (p categoryService) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	return p.categoryRepo.FindByID(ctx, id)
}

func (p categoryService) GetAll(ctx context.Context, filter entity.CategoriesFilter) ([]entity.Category, error) {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	return p.categoryRepo.FindAll(ctx, filter)
}

func (p categoryService) Delete(ctx context.Context, id string) error {
	return p.categoryRepo.Delete(ctx, id)
}
