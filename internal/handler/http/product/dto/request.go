package dto

import (
	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/pkg/utils"
)

// ProductCreateRequest represents the request payload for creating a product
type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required" validate:"required"`
	Description string  `json:"description" binding:"required"`
	SKU         string  `json:"sku" binding:"required"`
	Price       float64 `json:"price" binding:"min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	ImageURL    string  `json:"imageUrl"`
	CategoryID  string  `json:"categoryId" binding:"required"`
} // @name ProductCreateRequest

func (r ProductCreateRequest) ToDomain() entity.Product {
	return entity.Product{
		Name:        r.Name,
		Description: r.Description,
		SKU:         r.SKU,
		Price:       utils.SetPtr(r.Price),
		Stock:       utils.SetPtr(r.Stock),
		ImageURL:    utils.SetPtr(r.ImageURL),
		CategoryID:  r.CategoryID,
	}
}

// ProductUpdateRequest represents the request payload for updating a product
type ProductUpdateRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	SKU         *string  `json:"sku,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
	ImageURL    *string  `json:"imageUrl,omitempty"`
	CategoryID  *string  `json:"categoryId,omitempty"`
} //	@name	ProductUpdateRequest

func (r ProductUpdateRequest) ToDomain() entity.Product {
	return entity.Product{
		Name:        utils.GetValue(r.Name),
		Description: utils.GetValue(r.Description),
		SKU:         utils.GetValue(r.SKU),
		Price:       r.Price,
		Stock:       r.Stock,
		ImageURL:    r.ImageURL,
		CategoryID:  utils.GetValue(r.CategoryID),
	}
}

type FilterProductRequest struct {
	Name       *string  `form:"name,omitempty"`
	CategoryID *string  `form:"categoryId,omitempty"`
	MaxPrice   *float64 `form:"maxPrice,omitempty"`
	MinPrice   *float64 `form:"minPrice,omitempty"`
	Limit      int      `form:"limit"`
	Offset     int      `form:"offset"`
}

func (r FilterProductRequest) ToDomain() entity.ProductFilter {
	return entity.ProductFilter{
		Name:       r.Name,
		CategoryID: r.CategoryID,
		MaxPrice:   r.MaxPrice,
		MinPrice:   r.MinPrice,
		Pagination: entity.Pagination{
			Limit:  r.Limit,
			Offset: r.Offset,
		},
	}
}
