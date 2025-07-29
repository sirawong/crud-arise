package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/pkg/utils"
	"gorm.io/gorm"
)

type ProductModel struct {
	ID          string  `gorm:"type:uuid;primaryKey"`
	Name        string  `gorm:"size:255;not null;index"`
	Description string  `gorm:"type:text"`
	SKU         string  `gorm:"size:100;unique;not null"`
	Price       float64 `gorm:"type:decimal(10,2);not null;default:0"`
	Stock       int     `gorm:"not null;default:0"`
	ImageURL    string  `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	CategoryID string         `gorm:"type:uuid;not null"`
	Category   *CategoryModel `gorm:"foreignKey:CategoryID"`
}

func (ProductModel) TableName() string {
	return "products"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}

func ToProductEntity(model *ProductModel) *entity.Product {
	if model == nil {
		return nil
	}
	return &entity.Product{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		SKU:         model.SKU,
		Price:       utils.SetPtr(model.Price),
		Stock:       utils.SetPtr(model.Stock),
		ImageURL:    utils.SetPtr(model.ImageURL),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		CategoryID:  model.CategoryID,
		Category:    ToCategoryEntity(model.Category),
	}
}

func ToProductsEntity(models []ProductModel) []entity.Product {
	products := make([]entity.Product, 0, len(models))

	for _, model := range models {
		products = append(products, *ToProductEntity(&model))
	}

	return products
}

func ToProductModel(entity *entity.Product) *ProductModel {
	if entity == nil {
		return nil
	}
	return &ProductModel{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		SKU:         entity.SKU,
		Price:       utils.GetValue(entity.Price),
		Stock:       utils.GetValue(entity.Stock),
		ImageURL:    utils.GetValue(entity.ImageURL),
		CategoryID:  entity.CategoryID,
	}
}
