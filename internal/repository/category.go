package repository

import (
	"context"
	"errors"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/internal/domain/repository"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	"github.com/sirawong/crud-arise/internal/repository/models"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db: db}
}

func (c categoryRepository) Create(ctx context.Context, category *entity.Category) (string, error) {
	if category == nil {
		return "", apperr.ErrInvalidArgument.WithMessage("category cannot be nil")
	}

	createModel := models.ToCategoryModel(category)
	if createModel == nil {
		return "", apperr.ErrInvalidArgument.WithMessage("createModel cannot be nil")
	}
	err := c.db.WithContext(ctx).Create(&createModel).Error
	if err != nil {
		return "", apperr.ErrInternal.Wrap(err)
	}

	return createModel.ID, nil
}

func (c categoryRepository) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	var category models.CategoryModel

	err := c.db.WithContext(ctx).First(&category, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.ErrNotFound.Wrap(err)
		}
		return nil, apperr.ErrInternal.Wrap(err)
	}
	return models.ToCategoryEntity(&category), nil
}

func (c categoryRepository) FindAll(ctx context.Context, filter entity.CategoriesFilter) ([]entity.Category, error) {
	query := c.db.WithContext(ctx).Model(&models.CategoryModel{})
	if filter.Name != nil {
		searchPattern := "%" + *filter.Name + "%"
		query.Where("name ILIKE ?", searchPattern)
	}
	query = query.Limit(filter.Limit).Offset(filter.Offset)

	var categories []models.CategoryModel
	err := query.Find(&categories).Error
	if err != nil {
		return nil, apperr.ErrInternal.Wrap(err)
	}

	return models.ToCategoriesEntity(categories), nil
}

func (c categoryRepository) Update(ctx context.Context, category *entity.Category) error {
	if category == nil {
		return apperr.ErrInvalidArgument.WithMessage("category cannot be nil")
	}

	err := c.db.WithContext(ctx).Model(&models.CategoryModel{}).
		Where("id = ?", category.ID).
		Updates(map[string]interface{}{
			"name": category.Name,
		}).Error
	if err != nil {
		return apperr.ErrInternal.Wrap(err)
	}

	return nil
}

func (c categoryRepository) Delete(ctx context.Context, id string) error {
	err := c.db.WithContext(ctx).Delete(&models.CategoryModel{}, "id = ?", id).Error
	if err != nil {
		return apperr.ErrInternal.Wrap(err)
	}
	return nil
}
