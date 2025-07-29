package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperr "github.com/sirawong/crud-arise/internal/errors"
)

func mapErrorToHTTPStatus(code string) int {
	switch code {
	case apperr.ErrNotFound.Code:
		return http.StatusNotFound
	case apperr.ErrInvalidArgument.Code:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func RespondWithError(c *gin.Context, err error) {
	code := apperr.GetCode(err)
	httpStatus := mapErrorToHTTPStatus(code)
	c.JSON(httpStatus, gin.H{"error_code": code, "message": err.Error()})
}
