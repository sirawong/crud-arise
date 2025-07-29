package product

import (
	"context"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/internal/domain/repository"
	apperr "github.com/sirawong/crud-arise/internal/errors"
)

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

//go:generate mockgen -source=product.go -destination=mocks/mock_product.go -package=mocks
type ProductService interface {
	Create(ctx context.Context, product entity.Product) (string, error)
	Update(ctx context.Context, id string, product entity.Product) error
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	GetAll(ctx context.Context, filter entity.ProductFilter) ([]entity.Product, error)
	Delete(ctx context.Context, id string) error
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (p productService) Create(ctx context.Context, product entity.Product) (string, error) {
	err := p.validateCategory(ctx, product.CategoryID)
	if err != nil {
		return "", err
	}

	id, err := p.productRepo.Create(ctx, &product)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (p productService) validateCategory(ctx context.Context, id string) error {
	_, err := p.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) Update(ctx context.Context, id string, product entity.Product) error {
	if product.CategoryID != "" {
		err := p.validateCategory(ctx, product.CategoryID)
		if err != nil {
			return err
		}
	}

	product.ID = id
	return p.productRepo.Update(ctx, &product)
}

func (p productService) GetByID(ctx context.Context, id string) (*entity.Product, error) {

	return p.productRepo.FindByID(ctx, id)
}

func (p productService) GetAll(ctx context.Context, filter entity.ProductFilter) ([]entity.Product, error) {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	if filter.MinPrice != nil && filter.MaxPrice != nil {
		if *filter.MinPrice > *filter.MaxPrice {
			return nil, apperr.ErrInvalidArgument.WithMessage("min price cannot be greater than max price")
		}
	}

	return p.productRepo.FindAll(ctx, filter)
}

func (p productService) Delete(ctx context.Context, id string) error {

	return p.productRepo.Delete(ctx, id)
}
