package category

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	"github.com/sirawong/crud-arise/internal/handler/http/category/dto"
	"github.com/sirawong/crud-arise/internal/services/category/mocks"
	"github.com/stretchr/testify/suite"
)

type CategoryHandlerTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	mockService *mocks.MockCategoryService
	handler     *CategoryHandler
	router      *gin.Engine
}

func (suite *CategoryHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockService = mocks.NewMockCategoryService(suite.mockCtrl)
	suite.handler = NewCategoryHandler(suite.mockService)
	suite.router = gin.New()

	v1 := suite.router.Group("/api/v1")
	cate := v1.Group("/categories")
	{
		cate.POST("/", suite.handler.Create)
		cate.GET("/", suite.handler.ListAll)
		cate.GET("/:id", suite.handler.GetByID)
		cate.PUT("/:id", suite.handler.Update)
		cate.DELETE("/:id", suite.handler.Delete)
	}
}

func (suite *CategoryHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *CategoryHandlerTestSuite) TestCreate_Success() {

	request := dto.CategoryRequest{
		Name: "Test Category",
	}
	expectedID := "category-123"

	suite.mockService.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(expectedID, nil).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/categories/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedID, response["status"])
}

func (suite *CategoryHandlerTestSuite) TestCreate_InvalidJSON() {

	invalidJSON := `{"name": 123}`

	req, _ := http.NewRequest("POST", "/api/v1/categories/", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *CategoryHandlerTestSuite) TestCreate_ServiceError() {

	request := dto.CategoryRequest{
		Name: "Test Category",
	}
	expectedErr := apperr.ErrInvalidArgument.WithMessage("name already exists")

	suite.mockService.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return("", expectedErr).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/categories/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *CategoryHandlerTestSuite) TestUpdate_Success() {

	categoryID := "category-123"
	request := dto.CategoryRequest{
		Name: "Updated Category",
	}

	suite.mockService.EXPECT().
		Update(gomock.Any(), categoryID, gomock.Any()).
		Return(nil).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/categories/%s", categoryID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("updated", response["status"])
}

func (suite *CategoryHandlerTestSuite) TestUpdate_ServiceError() {

	categoryID := "category-123"
	request := dto.CategoryRequest{
		Name: "Updated Category",
	}
	expectedErr := apperr.ErrNotFound.WithMessage("category not found")

	suite.mockService.EXPECT().
		Update(gomock.Any(), categoryID, gomock.Any()).
		Return(expectedErr).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/categories/%s", categoryID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *CategoryHandlerTestSuite) TestGetByID_Success() {

	categoryID := "category-123"
	expectedCategory := &entity.Category{
		ID:        categoryID,
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockService.EXPECT().
		GetByID(gomock.Any(), categoryID).
		Return(expectedCategory, nil).
		Times(1)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/categories/%s", categoryID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Category
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedCategory.ID, response.ID)
	suite.Equal(expectedCategory.Name, response.Name)
}

func (suite *CategoryHandlerTestSuite) TestGetByID_ServiceError() {

	categoryID := "non-existent"
	expectedErr := apperr.ErrNotFound.WithMessage("category not found")

	suite.mockService.EXPECT().
		GetByID(gomock.Any(), categoryID).
		Return(nil, expectedErr).
		Times(1)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/categories/%s", categoryID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *CategoryHandlerTestSuite) TestListAll_Success() {

	expectedCategories := []entity.Category{
		{
			ID:        "1",
			Name:      "Category 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Name:      "Category 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.mockService.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		Return(expectedCategories, nil).
		Times(1)

	req, _ := http.NewRequest("GET", "/api/v1/categories/", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response []*dto.Category
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Len(response, 2)
	suite.Equal(expectedCategories[0].ID, response[0].ID)
	suite.Equal(expectedCategories[1].ID, response[1].ID)
}

func (suite *CategoryHandlerTestSuite) TestListAll_WithFilters() {

	suite.mockService.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx interface{}, filter entity.CategoriesFilter) ([]entity.Category, error) {
			suite.Equal("test", *filter.Name)
			suite.Equal(20, filter.Limit)
			return []entity.Category{}, nil
		}).
		Times(1)

	req, _ := http.NewRequest("GET", "/api/v1/categories/?name=test&limit=20", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *CategoryHandlerTestSuite) TestDelete_Success() {

	categoryID := "category-123"

	suite.mockService.EXPECT().
		Delete(gomock.Any(), categoryID).
		Return(nil).
		Times(1)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/categories/%s", categoryID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("deleted", response["status"])
}

func (suite *CategoryHandlerTestSuite) TestDelete_ServiceError() {

	categoryID := "category-123"
	expectedErr := apperr.ErrNotFound.WithMessage("category not found")

	suite.mockService.EXPECT().
		Delete(gomock.Any(), categoryID).
		Return(expectedErr).
		Times(1)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/categories/%s", categoryID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func TestCategoryHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryHandlerTestSuite))
}
