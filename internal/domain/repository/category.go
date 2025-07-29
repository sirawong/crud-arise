package repository

import (
	"context"

	"github.com/sirawong/crud-arise/internal/domain/entity"
)

//go:generate mockgen -source=category.go -destination=mocks/mock_category.go -package=mocks
type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	FindAll(ctx context.Context, filter entity.CategoriesFilter) ([]entity.Category, error)
	Delete(ctx context.Context, id string) error
}
