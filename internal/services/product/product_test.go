package product

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sirawong/crud-arise/internal/domain/entity"
	"github.com/sirawong/crud-arise/internal/domain/repository/mocks"
	"github.com/sirawong/crud-arise/pkg/utils"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type ProductServiceTestSuite struct {
	suite.Suite
	mockCtrl         *gomock.Controller
	mockProductRepo  *mocks.MockProductRepository
	mockCategoryRepo *mocks.MockCategoryRepository
	service          ProductService
	ctx              context.Context
}

func (suite *ProductServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockProductRepo = mocks.NewMockProductRepository(suite.mockCtrl)
	suite.mockCategoryRepo = mocks.NewMockCategoryRepository(suite.mockCtrl)
	suite.service = NewProductService(suite.mockProductRepo, suite.mockCategoryRepo)
	suite.ctx = context.Background()
}

func (suite *ProductServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *ProductServiceTestSuite) TestCreate_Success() {

	categoryID := "category-123"
	product := entity.Product{
		Name:        "Test Product",
		Description: "Test Description",
		SKU:         "TEST-001",
		Price:       utils.SetPtr(99.99),
		Stock:       utils.SetPtr(100),
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	expectedID := "product-123"
	mockCategory := &entity.Category{
		ID:   categoryID,
		Name: "Test Category",
	}

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(mockCategory, nil).
		Times(1)

	suite.mockProductRepo.EXPECT().
		Create(suite.ctx, &product).
		Return(expectedID, nil).
		Times(1)

	id, err := suite.service.Create(suite.ctx, product)

	suite.NoError(err)
	suite.Equal(expectedID, id)
}

func (suite *ProductServiceTestSuite) TestCreate_InvalidCategory() {

	categoryID := "invalid-category"
	product := entity.Product{
		Name:       "Test Product",
		CategoryID: categoryID,
	}
	expectedErr := errors.New("category not found")

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(nil, expectedErr).
		Times(1)

	id, err := suite.service.Create(suite.ctx, product)

	suite.Error(err)
	suite.Equal("", id)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestCreate_RepositoryError() {

	categoryID := "category-123"
	product := entity.Product{
		Name:       "Test Product",
		CategoryID: categoryID,
	}
	mockCategory := &entity.Category{
		ID:   categoryID,
		Name: "Test Category",
	}
	expectedErr := errors.New("repository error")

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(mockCategory, nil).
		Times(1)

	suite.mockProductRepo.EXPECT().
		Create(suite.ctx, &product).
		Return("", expectedErr).
		Times(1)

	id, err := suite.service.Create(suite.ctx, product)

	suite.Error(err)
	suite.Equal("", id)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestUpdate_Success() {

	productID := "product-123"
	categoryID := "category-123"
	product := entity.Product{
		Name:        "Updated Product",
		Description: "Updated Description",
		CategoryID:  categoryID,
	}
	expectedProduct := product
	expectedProduct.ID = productID
	mockCategory := &entity.Category{
		ID:   categoryID,
		Name: "Test Category",
	}

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(mockCategory, nil).
		Times(1)

	suite.mockProductRepo.EXPECT().
		Update(suite.ctx, &expectedProduct).
		Return(nil).
		Times(1)

	err := suite.service.Update(suite.ctx, productID, product)

	suite.NoError(err)
}

func (suite *ProductServiceTestSuite) TestUpdate_WithoutCategoryID() {

	productID := "product-123"
	product := entity.Product{
		Name:        "Updated Product",
		Description: "Updated Description",
		CategoryID:  "",
	}
	expectedProduct := product
	expectedProduct.ID = productID

	suite.mockProductRepo.EXPECT().
		Update(suite.ctx, &expectedProduct).
		Return(nil).
		Times(1)

	err := suite.service.Update(suite.ctx, productID, product)

	suite.NoError(err)
}

func (suite *ProductServiceTestSuite) TestUpdate_InvalidCategory() {

	productID := "product-123"
	categoryID := "invalid-category"
	product := entity.Product{
		Name:       "Updated Product",
		CategoryID: categoryID,
	}
	expectedErr := errors.New("category not found")

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(nil, expectedErr).
		Times(1)

	err := suite.service.Update(suite.ctx, productID, product)

	suite.Error(err)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestUpdate_RepositoryError() {

	productID := "product-123"
	categoryID := "category-123"
	product := entity.Product{
		Name:       "Updated Product",
		CategoryID: categoryID,
	}
	expectedProduct := product
	expectedProduct.ID = productID
	mockCategory := &entity.Category{
		ID:   categoryID,
		Name: "Test Category",
	}
	expectedErr := errors.New("repository error")

	suite.mockCategoryRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(mockCategory, nil).
		Times(1)

	suite.mockProductRepo.EXPECT().
		Update(suite.ctx, &expectedProduct).
		Return(expectedErr).
		Times(1)

	err := suite.service.Update(suite.ctx, productID, product)

	suite.Error(err)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestGetByID_Success() {

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
	}

	suite.mockProductRepo.EXPECT().
		FindByID(suite.ctx, productID).
		Return(expectedProduct, nil).
		Times(1)

	product, err := suite.service.GetByID(suite.ctx, productID)

	suite.NoError(err)
	suite.Equal(expectedProduct, product)
}

func (suite *ProductServiceTestSuite) TestGetByID_NotFound() {

	productID := "non-existent-id"
	expectedErr := errors.New("product not found")

	suite.mockProductRepo.EXPECT().
		FindByID(suite.ctx, productID).
		Return(nil, expectedErr).
		Times(1)

	product, err := suite.service.GetByID(suite.ctx, productID)

	suite.Error(err)
	suite.Nil(product)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestGetAll_Success() {

	filter := entity.ProductFilter{
		Name:       utils.SetPtr("Test"),
		CategoryID: utils.SetPtr("category-123"),
		MinPrice:   utils.SetPtr(10.0),
		MaxPrice:   utils.SetPtr(100.0),
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedProducts := []entity.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Description 1",
			Price:       utils.SetPtr(50.0),
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Description 2",
			Price:       utils.SetPtr(75.0),
		},
	}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_DefaultLimit() {

	filter := entity.ProductFilter{
		Pagination: entity.Pagination{
			Limit:  0,
			Offset: 0,
		},
	}
	expectedFilter := filter
	expectedFilter.Limit = 10

	expectedProducts := []entity.Product{}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, expectedFilter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_MaxLimit() {

	filter := entity.ProductFilter{
		Pagination: entity.Pagination{
			Limit:  200,
			Offset: 0,
		},
	}
	expectedFilter := filter
	expectedFilter.Limit = 100

	expectedProducts := []entity.Product{}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, expectedFilter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_InvalidPriceRange() {

	filter := entity.ProductFilter{
		MinPrice: utils.SetPtr(100.0),
		MaxPrice: utils.SetPtr(50.0),
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.Error(err)
	suite.Nil(products)
	suite.Contains(err.Error(), "min price cannot be greater than max price")
}

func (suite *ProductServiceTestSuite) TestGetAll_ValidPriceRange() {

	filter := entity.ProductFilter{
		MinPrice: utils.SetPtr(50.0),
		MaxPrice: utils.SetPtr(100.0),
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedProducts := []entity.Product{}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_OnlyMinPrice() {

	filter := entity.ProductFilter{
		MinPrice: utils.SetPtr(50.0),
		MaxPrice: nil,
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedProducts := []entity.Product{}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_OnlyMaxPrice() {

	filter := entity.ProductFilter{
		MinPrice: nil,
		MaxPrice: utils.SetPtr(100.0),
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedProducts := []entity.Product{}

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(expectedProducts, nil).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedProducts, products)
}

func (suite *ProductServiceTestSuite) TestGetAll_RepositoryError() {

	filter := entity.ProductFilter{
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedErr := errors.New("repository error")

	suite.mockProductRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(nil, expectedErr).
		Times(1)

	products, err := suite.service.GetAll(suite.ctx, filter)

	suite.Error(err)
	suite.Nil(products)
	suite.Equal(expectedErr, err)
}

func (suite *ProductServiceTestSuite) TestDelete_Success() {

	productID := "product-123"

	suite.mockProductRepo.EXPECT().
		Delete(suite.ctx, productID).
		Return(nil).
		Times(1)

	err := suite.service.Delete(suite.ctx, productID)

	suite.NoError(err)
}

func (suite *ProductServiceTestSuite) TestDelete_RepositoryError() {

	productID := "product-123"
	expectedErr := errors.New("repository error")

	suite.mockProductRepo.EXPECT().
		Delete(suite.ctx, productID).
		Return(expectedErr).
		Times(1)

	err := suite.service.Delete(suite.ctx, productID)

	suite.Error(err)
	suite.Equal(expectedErr, err)
}

func TestProductServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceTestSuite))
}
