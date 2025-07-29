package category

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

type CategoryServiceTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	mockRepo *mocks.MockCategoryRepository
	service  CategoryService
	ctx      context.Context
}

func (suite *CategoryServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockCategoryRepository(suite.mockCtrl)
	suite.service = NewCategoryService(suite.mockRepo)
	suite.ctx = context.Background()
}

func (suite *CategoryServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *CategoryServiceTestSuite) TestCreate_Success() {

	category := entity.Category{
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	expectedID := "test-id-123"

	suite.mockRepo.EXPECT().
		Create(suite.ctx, &category).
		Return(expectedID, nil).
		Times(1)

	id, err := suite.service.Create(suite.ctx, category)

	suite.NoError(err)
	suite.Equal(expectedID, id)
}

func (suite *CategoryServiceTestSuite) TestCreate_RepositoryError() {

	category := entity.Category{
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	expectedErr := errors.New("repository error")

	suite.mockRepo.EXPECT().
		Create(suite.ctx, &category).
		Return("", expectedErr).
		Times(1)

	id, err := suite.service.Create(suite.ctx, category)

	suite.Error(err)
	suite.Equal("", id)
	suite.Equal(expectedErr, err)
}

func (suite *CategoryServiceTestSuite) TestUpdate_Success() {

	categoryID := "test-id-123"
	category := entity.Category{
		Name:      "Updated Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	expectedCategory := category
	expectedCategory.ID = categoryID

	suite.mockRepo.EXPECT().
		Update(suite.ctx, &expectedCategory).
		Return(nil).
		Times(1)

	err := suite.service.Update(suite.ctx, categoryID, category)

	suite.NoError(err)
}

func (suite *CategoryServiceTestSuite) TestUpdate_RepositoryError() {

	categoryID := "test-id-123"
	category := entity.Category{
		Name:      "Updated Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	expectedCategory := category
	expectedCategory.ID = categoryID
	expectedErr := errors.New("repository error")

	suite.mockRepo.EXPECT().
		Update(suite.ctx, &expectedCategory).
		Return(expectedErr).
		Times(1)

	err := suite.service.Update(suite.ctx, categoryID, category)

	suite.Error(err)
	suite.Equal(expectedErr, err)
}

func (suite *CategoryServiceTestSuite) TestGetByID_Success() {

	categoryID := "test-id-123"
	expectedCategory := &entity.Category{
		ID:        categoryID,
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(expectedCategory, nil).
		Times(1)

	category, err := suite.service.GetByID(suite.ctx, categoryID)

	suite.NoError(err)
	suite.Equal(expectedCategory, category)
}

func (suite *CategoryServiceTestSuite) TestGetByID_NotFound() {

	categoryID := "non-existent-id"
	expectedErr := errors.New("category not found")

	suite.mockRepo.EXPECT().
		FindByID(suite.ctx, categoryID).
		Return(nil, expectedErr).
		Times(1)

	category, err := suite.service.GetByID(suite.ctx, categoryID)

	suite.Error(err)
	suite.Nil(category)
	suite.Equal(expectedErr, err)
}

func (suite *CategoryServiceTestSuite) TestGetAll_Success() {

	filter := entity.CategoriesFilter{
		Name: utils.SetPtr("Test"),
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedCategories := []entity.Category{
		{
			ID:        "1",
			Name:      "Test Category 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Name:      "Test Category 2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.mockRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(expectedCategories, nil).
		Times(1)

	categories, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedCategories, categories)
}

func (suite *CategoryServiceTestSuite) TestGetAll_DefaultLimit() {

	filter := entity.CategoriesFilter{
		Pagination: entity.Pagination{
			Limit:  0,
			Offset: 0,
		},
	}
	expectedFilter := filter
	expectedFilter.Limit = 10

	expectedCategories := []entity.Category{}

	suite.mockRepo.EXPECT().
		FindAll(suite.ctx, expectedFilter).
		Return(expectedCategories, nil).
		Times(1)

	categories, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedCategories, categories)
}

func (suite *CategoryServiceTestSuite) TestGetAll_MaxLimit() {

	filter := entity.CategoriesFilter{
		Pagination: entity.Pagination{
			Limit:  200,
			Offset: 0,
		},
	}
	expectedFilter := filter
	expectedFilter.Limit = 100

	expectedCategories := []entity.Category{}

	suite.mockRepo.EXPECT().
		FindAll(suite.ctx, expectedFilter).
		Return(expectedCategories, nil).
		Times(1)

	categories, err := suite.service.GetAll(suite.ctx, filter)

	suite.NoError(err)
	suite.Equal(expectedCategories, categories)
}

func (suite *CategoryServiceTestSuite) TestGetAll_RepositoryError() {

	filter := entity.CategoriesFilter{
		Pagination: entity.Pagination{
			Limit:  20,
			Offset: 0,
		},
	}
	expectedErr := errors.New("repository error")

	suite.mockRepo.EXPECT().
		FindAll(suite.ctx, filter).
		Return(nil, expectedErr).
		Times(1)

	categories, err := suite.service.GetAll(suite.ctx, filter)

	suite.Error(err)
	suite.Nil(categories)
	suite.Equal(expectedErr, err)
}

func (suite *CategoryServiceTestSuite) TestDelete_Success() {

	categoryID := "test-id-123"

	suite.mockRepo.EXPECT().
		Delete(suite.ctx, categoryID).
		Return(nil).
		Times(1)

	err := suite.service.Delete(suite.ctx, categoryID)

	suite.NoError(err)
}

func (suite *CategoryServiceTestSuite) TestDelete_RepositoryError() {

	categoryID := "test-id-123"
	expectedErr := errors.New("repository error")

	suite.mockRepo.EXPECT().
		Delete(suite.ctx, categoryID).
		Return(expectedErr).
		Times(1)

	err := suite.service.Delete(suite.ctx, categoryID)

	suite.Error(err)
	suite.Equal(expectedErr, err)
}

func TestCategoryServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CategoryServiceTestSuite))
}
