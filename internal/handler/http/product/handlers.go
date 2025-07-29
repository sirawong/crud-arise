package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	handlererr "github.com/sirawong/crud-arise/internal/handler/http/errors"
	"github.com/sirawong/crud-arise/internal/handler/http/product/dto"
	productSrv "github.com/sirawong/crud-arise/internal/services/product"
)

type ProductHandler struct {
	productService productSrv.ProductService
}

func NewProductHandler(productService productSrv.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// Create godoc
//
//	@Summary		Create a new product
//	@Description	Create a new product with the provided information
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dto.ProductCreateRequest	true	"Product creation information"
//	@Success		201		{object}	map[string]interface{}		"{"id": "product_id"}"
//	@Failure		400		{object}	map[string]interface{}		"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		500		{object}	map[string]interface{}		"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/products [post]
func (h ProductHandler) Create(c *gin.Context) {
	var req dto.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	id, err := h.productService.Create(c, req.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Update godoc
//
//	@Summary		Update a product
//	@Description	Update an existing product by ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Product ID"
//	@Param			product	body		dto.ProductUpdateRequest	true	"Product update information"
//	@Success		200		{object}	map[string]interface{}		"{"status": "updated"}"
//	@Failure		400		{object}	map[string]interface{}		"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404		{object}	map[string]interface{}		"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500		{object}	map[string]interface{}		"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/products/{id} [put]
func (h ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	var req dto.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	err := h.productService.Update(c, id, req.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// GetByID godoc
//
//	@Summary		Get a product by ID
//	@Description	Get a single product by its ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Product ID"
//	@Success		200	{object}	dto.Product				"Product information"
//	@Failure		400	{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404	{object}	map[string]interface{}	"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500	{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/products/{id} [get]
func (h ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	product, err := h.productService.GetByID(c, id)
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.ProductFromDomain(product))
}

// ListAll godoc
//
//	@Summary		Get all products
//	@Description	Get a list of all products with optional filtering
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			name		query		string					false	"Search insensitive by products name"
//	@Param			categoryId	query		string					false	"Filter by category ID"
//	@Param			minPrice	query		number					false	"Minimum price filter"
//	@Param			maxPrice	query		number					false	"Maximum price filter"
//	@Param			limit		query		int						false	"Limit number of results (default: 10, limit: 100)"
//	@Param			offset		query		int						false	"Offset for pagination (default: 0)"
//	@Success		200			{array}		dto.Product				"List of products"
//	@Failure		500			{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error	description"}"
//	@Router			/products [get]
func (h ProductHandler) ListAll(c *gin.Context) {
	var query dto.FilterProductRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	products, err := h.productService.GetAll(c, query.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.ProductsFromDomain(products))
}

// Delete godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product by ID
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Product ID"
//	@Success		200	{object}	map[string]interface{}	"{"status": "deleted"}"
//	@Failure		400	{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404	{object}	map[string]interface{}	"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500	{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/products/{id} [delete]
func (h ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	err := h.productService.Delete(c, id)
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
