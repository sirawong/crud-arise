package repository

import (
	"context"

	"github.com/sirawong/crud-arise/internal/domain/entity"
)

//go:generate mockgen -source=product.go -destination=mocks/mock_product.go -package=mocks
type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	FindAll(ctx context.Context, filter entity.ProductFilter) ([]entity.Product, error)
	Delete(ctx context.Context, id string) error
}
