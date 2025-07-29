package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/sirawong/crud-arise/internal/domain/entity"
	"gorm.io/gorm"
)

type CategoryModel struct {
	ID        string         `gorm:"type:uuid;primaryKey"`
	Name      string         `gorm:"size:100;unique;not null"`
	Products  []ProductModel `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CategoryModel) TableName() string {
	return "categories"
}

func (c *CategoryModel) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

func ToCategoryEntity(model *CategoryModel) *entity.Category {
	if model == nil {
		return nil
	}
	return &entity.Category{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToCategoriesEntity(models []CategoryModel) []entity.Category {
	result := make([]entity.Category, 0, len(models))
	for _, model := range models {
		result = append(result, *ToCategoryEntity(&model))
	}
	return result
}

func ToCategoryModel(entity *entity.Category) *CategoryModel {
	if entity == nil {
		return nil
	}
	return &CategoryModel{
		ID:   entity.ID,
		Name: entity.Name,
	}
}
