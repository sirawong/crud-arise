package di

import (
	"github.com/sirawong/crud-arise/internal/handler/http"
	category2 "github.com/sirawong/crud-arise/internal/handler/http/category"
	product2 "github.com/sirawong/crud-arise/internal/handler/http/product"
	"github.com/sirawong/crud-arise/internal/repository"
	"github.com/sirawong/crud-arise/internal/services/category"
	"github.com/sirawong/crud-arise/internal/services/product"
	"github.com/sirawong/crud-arise/pkg/config"
	"github.com/sirawong/crud-arise/pkg/database"
)

func NewApplication() (*Application, func(), error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, nil, err
	}

	db, cleanup, err := database.NewConnection(cfg)
	if err != nil {
		return nil, nil, err
	}

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryHandler := category2.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := product.NewProductService(productRepo, categoryRepo)
	productHandler := product2.NewProductHandler(productService)

	httpRouter := http.NewRouter(productHandler, categoryHandler)
	httpServer := httpRouter.NewServer(cfg)

	return &Application{
			httpServer: httpServer,
			Cfg:        cfg,
		}, func() {
			cleanup()
		}, nil
}
