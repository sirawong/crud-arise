package dto

import (
	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/pkg/utils"
)

// CategoryRequest represents the request payload for creating/updating a category
type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
} //	@name	CategoryRequest

func (r CategoryRequest) ToDomain() entity.Category {
	return entity.Category{
		Name: r.Name,
	}
}

type FilterCategoriesRequest struct {
	Name   string `form:"name,omitempty"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

func (r FilterCategoriesRequest) ToDomain() entity.CategoriesFilter {
	return entity.CategoriesFilter{
		Name: utils.SetPtr(r.Name),
		Pagination: entity.Pagination{
			Limit:  r.Limit,
			Offset: r.Offset,
		},
	}
}
