package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	"github.com/sirawong/crud-arise/internal/handler/http/category/dto"
	handlererr "github.com/sirawong/crud-arise/internal/handler/http/errors"
	categorySrv "github.com/sirawong/crud-arise/internal/services/category"
)

type CategoryHandler struct {
	categoryService categorySrv.CategoryService
}

func NewCategoryHandler(categoryService categorySrv.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

// Create godoc
//
//	@Summary		Create a new category
//	@Description	Create a new category with the provided information
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			category	body		dto.CategoryRequest		true	"Category information"
//	@Success		201			{object}	map[string]interface{}	"{"status": "category_id"}"
//	@Failure		400			{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		500			{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/categories [post]
func (h CategoryHandler) Create(c *gin.Context) {
	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	id, err := h.categoryService.Create(c, req.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": id})
}

// Update godoc
//
//	@Summary		Update a category
//	@Description	Update an existing category by ID
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string					true	"Category ID"
//	@Param			category	body		dto.CategoryRequest		true	"Category information"
//	@Success		200			{object}	map[string]interface{}	"{"status": "updated"}"
//	@Failure		400			{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404			{object}	map[string]interface{}	"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500			{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/categories/{id} [put]
func (h CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	err := h.categoryService.Update(c, id, req.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// GetByID godoc
//
//	@Summary		Get a category by ID
//	@Description	Get a single category by its ID
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Category ID"
//	@Success		200	{object}	dto.Category			"Category information"
//	@Failure		400	{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404	{object}	map[string]interface{}	"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500	{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/categories/{id} [get]
func (h CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	category, err := h.categoryService.GetByID(c, id)
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.CategoryFromDomain(category))
}

// ListAll godoc
//
//	@Summary		Get all categories
//	@Description	Get a list of all categories
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string					false	"Search insensitive by category name"
//	@Param			limit	query		int						false	"Limit number of results (default: 10, limit: 100)"
//	@Param			offset	query		int						false	"Offset for pagination (default: 0)"
//	@Success		200		{array}		dto.Category			"List of categories"
//	@Failure		500		{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error	description"}"
//	@Router			/categories [get]
func (h CategoryHandler) ListAll(c *gin.Context) {
	var query dto.FilterCategoriesRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.Wrap(err))
		return
	}

	categories, err := h.categoryService.GetAll(c, query.ToDomain())
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.CategoriesFromDomain(categories))
}

// Delete godoc
//
//	@Summary		Delete a category
//	@Description	Delete a category by ID
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"Category ID"
//	@Success		200	{object}	map[string]interface{}	"{"status": "deleted"}"
//	@Failure		400	{object}	map[string]interface{}	"{"error_code": "INVALID_ARGUMENT", "message": "error	description"}"
//	@Failure		404	{object}	map[string]interface{}	"{"error_code": "NOT_FOUND", "message": "error			description"}"
//	@Failure		500	{object}	map[string]interface{}	"{"error_code": "INTERNAL_ERROR", "message": "error		description"}"
//	@Router			/categories/{id} [delete]
func (h CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handlererr.RespondWithError(c, apperr.ErrInvalidArgument.WithMessage("id is required"))
		return
	}

	err := h.categoryService.Delete(c, id)
	if err != nil {
		handlererr.RespondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
