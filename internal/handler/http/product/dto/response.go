package dto

import (
	"time"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/pkg/utils"
)

// Product represents the response payload for a product
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`

	Category *Category `json:"category,omitempty"`
} //	@name	Product

// Category represents category information in product response
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
} //	@name	Category

func ProductFromDomain(product *entity.Product) *Product {
	if product == nil {
		return nil
	}

	var category *Category
	if product.Category != nil {
		category = &Category{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		}
	}
	return &Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		SKU:         product.SKU,
		Price:       utils.GetValue(product.Price),
		Stock:       utils.GetValue(product.Stock),
		ImageURL:    utils.GetValue(product.ImageURL),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Category:    category,
	}
}

func ProductsFromDomain(products []entity.Product) []Product {
	productsRes := make([]Product, 0, len(products))
	for _, user := range products {
		product := ProductFromDomain(&user)
		if product == nil {
			continue
		}
		productsRes = append(productsRes, *product)
	}

	return productsRes
}
