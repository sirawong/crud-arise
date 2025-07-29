package entity

import "time"

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoriesFilter struct {
	Name *string
	Pagination
}
