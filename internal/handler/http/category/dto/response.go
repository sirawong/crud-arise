package dto

import (
	"time"

	"github.com/sirawong/crud-arise/internal/domain/entity"
)

// Category represents the response payload for a category
type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
} //	@name	Category

func CategoryFromDomain(category *entity.Category) *Category {
	if category == nil {
		return nil
	}
	return &Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func CategoriesFromDomain(categories []entity.Category) []*Category {
	if len(categories) == 0 {
		return nil
	}
	result := make([]*Category, 0, len(categories))
	for _, category := range categories {
		result = append(result, CategoryFromDomain(&category))
	}

	return result
}
