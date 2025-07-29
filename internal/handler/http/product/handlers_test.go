package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirawong/crud-arise/pkg/utils"
	"go.uber.org/mock/gomock"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	apperr "github.com/sirawong/crud-arise/internal/errors"
	"github.com/sirawong/crud-arise/internal/handler/http/product/dto"
	"github.com/sirawong/crud-arise/internal/services/product/mocks"
	"github.com/stretchr/testify/suite"
)

type ProductHandlerTestSuite struct {
	suite.Suite
	mockCtrl    *gomock.Controller
	mockService *mocks.MockProductService
	handler     *ProductHandler
	router      *gin.Engine
}

func (suite *ProductHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockService = mocks.NewMockProductService(suite.mockCtrl)
	suite.handler = NewProductHandler(suite.mockService)
	suite.router = gin.New()

	v1 := suite.router.Group("/api/v1")
	prd := v1.Group("/products")
	{
		prd.POST("/", suite.handler.Create)
		prd.GET("/", suite.handler.ListAll)
		prd.GET("/:id", suite.handler.GetByID)
		prd.PUT("/:id", suite.handler.Update)
		prd.DELETE("/:id", suite.handler.Delete)
	}
}

func (suite *ProductHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *ProductHandlerTestSuite) TestCreate_Success() {

	request := dto.ProductCreateRequest{
		Name:        "Test Product",
		Description: "Test Description",
		SKU:         "TEST-001",
		Price:       99.99,
		Stock:       100,
		ImageURL:    "http://example.com/image.jpg",
		CategoryID:  "category-123",
	}
	expectedID := "product-123"

	suite.mockService.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(expectedID, nil).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/products/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedID, response["id"])
}

func (suite *ProductHandlerTestSuite) TestCreate_InvalidJSON() {

	invalidJSON := `{"name": "Test", "price": "invalid"}`

	req, _ := http.NewRequest("POST", "/api/v1/products/", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *ProductHandlerTestSuite) TestCreate_ServiceError() {

	request := dto.ProductCreateRequest{
		Name:        "Test Product",
		Description: "Test Description",
		SKU:         "TEST-001",
		Price:       99.99,
		Stock:       100,
		CategoryID:  "category-123",
	}
	expectedErr := apperr.ErrNotFound.WithMessage("category not found")

	suite.mockService.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return("", expectedErr).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/api/v1/products/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *ProductHandlerTestSuite) TestUpdate_Success() {

	productID := "product-123"
	request := dto.ProductUpdateRequest{
		Name:  utils.SetPtr("Updated Product"),
		Price: utils.SetPtr(149.99),
	}

	suite.mockService.EXPECT().
		Update(gomock.Any(), productID, gomock.Any()).
		Return(nil).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/products/%s", productID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("updated", response["status"])
}

func (suite *ProductHandlerTestSuite) TestUpdate_ServiceError() {

	productID := "product-123"
	request := dto.ProductUpdateRequest{
		Name: utils.SetPtr("Updated Product"),
	}
	expectedErr := apperr.ErrNotFound.WithMessage("product not found")

	suite.mockService.EXPECT().
		Update(gomock.Any(), productID, gomock.Any()).
		Return(expectedErr).
		Times(1)

	body, _ := json.Marshal(request)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/products/%s", productID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *ProductHandlerTestSuite) TestGetByID_Success() {

	productID := "product-123"
	expectedProduct := &entity.Product{
		ID:          productID,
		Name:        "Test Product",
		Description: "Test Description",
		SKU:         "TEST-001",
		Price:       utils.SetPtr(99.99),
		Stock:       utils.SetPtr(100),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CategoryID:  "category-123",
		Category: &entity.Category{
			ID:   "category-123",
			Name: "Test Category",
		},
	}

	suite.mockService.EXPECT().
		GetByID(gomock.Any(), productID).
		Return(expectedProduct, nil).
		Times(1)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s", productID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response dto.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(expectedProduct.ID, response.ID)
	suite.Equal(expectedProduct.Name, response.Name)
	suite.NotNil(response.Category)
}

func (suite *ProductHandlerTestSuite) TestGetByID_ServiceError() {

	productID := "non-existent"
	expectedErr := apperr.ErrNotFound.WithMessage("product not found")

	suite.mockService.EXPECT().
		GetByID(gomock.Any(), productID).
		Return(nil, expectedErr).
		Times(1)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%s", productID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func (suite *ProductHandlerTestSuite) TestListAll_Success() {

	expectedProducts := []entity.Product{
		{
			ID:        "1",
			Name:      "Product 1",
			Price:     utils.SetPtr(50.0),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Name:      "Product 2",
			Price:     utils.SetPtr(75.0),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.mockService.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		Return(expectedProducts, nil).
		Times(1)

	req, _ := http.NewRequest("GET", "/api/v1/products/", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response []dto.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Len(response, 2)
}

func (suite *ProductHandlerTestSuite) TestListAll_WithFilters() {

	suite.mockService.EXPECT().
		GetAll(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx interface{}, filter entity.ProductFilter) ([]entity.Product, error) {
			suite.Equal("test", *filter.Name)
			suite.Equal(20, filter.Limit)
			return []entity.Product{}, nil
		}).
		Times(1)

	req, _ := http.NewRequest("GET", "/api/v1/products/?name=test&limit=20", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *ProductHandlerTestSuite) TestDelete_Success() {

	productID := "product-123"

	suite.mockService.EXPECT().
		Delete(gomock.Any(), productID).
		Return(nil).
		Times(1)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/products/%s", productID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("deleted", response["status"])
}

func (suite *ProductHandlerTestSuite) TestDelete_ServiceError() {

	productID := "product-123"
	expectedErr := apperr.ErrNotFound.WithMessage("product not found")

	suite.mockService.EXPECT().
		Delete(gomock.Any(), productID).
		Return(expectedErr).
		Times(1)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/products/%s", productID), nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code)
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
