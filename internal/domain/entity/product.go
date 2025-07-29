package entity

import (
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description string
	SKU         string
	Price       *float64
	Stock       *int
	ImageURL    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	CategoryID string
	Category   *Category
}

type ProductFilter struct {
	Name       *string
	CategoryID *string
	MinPrice   *float64
	MaxPrice   *float64
	Pagination
}
