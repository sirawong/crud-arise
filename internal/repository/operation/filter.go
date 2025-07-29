package operation

import (
	"github.com/sirawong/crud-arise/internal/domain/entity"
	"gorm.io/gorm"
)

func BuildQuery(db *gorm.DB, filter entity.ProductFilter) *gorm.DB {
	query := db

	if filter.Name != nil {
		query = query.Where("name ILIKE ?", "%"+*filter.Name+"%")
	}
	if filter.CategoryID != nil {
		query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.MinPrice != nil {
		query.Where("price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query.Where("price <= ?", filter.MaxPrice)
	}

	return query
}
