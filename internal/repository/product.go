package repository

import (
	"context"
	"errors"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/internal/domain/repository"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	"github.com/sirawong/crud-arise/internal/repository/models"
	"github.com/sirawong/crud-arise/internal/repository/operation"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (p productRepository) Create(ctx context.Context, product *entity.Product) (string, error) {
	if product == nil {
		return "", apperr.ErrInvalidArgument.WithMessage("product cannot be nil")
	}

	value := models.ToProductModel(product)
	err := p.db.WithContext(ctx).Create(&value).Error
	if err != nil {
		return "", apperr.ErrInternal.Wrap(err)
	}
	return value.ID, nil
}

func (p productRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	var product models.ProductModel
	err := p.db.WithContext(ctx).Preload("Category").First(&product, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.ErrNotFound.Wrap(err)
		}
		return nil, apperr.ErrInternal.Wrap(err)
	}
	return models.ToProductEntity(&product), nil
}

func (p productRepository) FindAll(ctx context.Context, filter entity.ProductFilter) ([]entity.Product, error) {
	query := p.db.WithContext(ctx).Model([]*models.ProductModel{})
	query = operation.BuildQuery(query, filter)
	query = query.Limit(filter.Limit).Offset(filter.Offset)

	var products []models.ProductModel
	err := query.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, apperr.ErrInternal.Wrap(err)
	}

	return models.ToProductsEntity(products), nil
}

func (p productRepository) Update(ctx context.Context, product *entity.Product) error {
	if product == nil {
		return apperr.ErrInvalidArgument.WithMessage("product cannot be nil")
	}

	update := operation.ToUpdateProductModel(product)
	if len(update) == 0 {
		return apperr.ErrInvalidArgument.WithMessage("product update cannot be nil")
	}

	err := p.db.WithContext(ctx).Model(&models.ProductModel{}).
		Where("id = ?", product.ID).Updates(update).Error
	if err != nil {
		return apperr.ErrInternal.Wrap(err)
	}

	return nil
}

func (p productRepository) Delete(ctx context.Context, id string) error {
	err := p.db.WithContext(ctx).Delete(&models.CategoryModel{}, "id = ?", id).Error
	if err != nil {
		return apperr.ErrInternal.Wrap(err)
	}
	return nil
}
