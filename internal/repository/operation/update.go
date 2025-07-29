package operation

import "github.com/sirawong/crud-arise/internal/domain/entity"

func ToUpdateProductModel(product *entity.Product) map[string]interface{} {
	if product == nil {
		return nil
	}

	result := make(map[string]interface{})

	if product.Name != "" {
		result["name"] = product.Name
	}
	if product.Description != "" {
		result["description"] = product.Description
	}
	if product.SKU != "" {
		result["sku"] = product.SKU
	}
	if product.Price != nil {
		result["price"] = product.Price
	}
	if product.Stock != nil {
		result["stock"] = product.Stock
	}
	if product.ImageURL != nil {
		result["image_url"] = product.ImageURL
	}

	if product.CategoryID != "" {
		result["category_id"] = product.CategoryID
	}

	return result
}
