package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirawong/crud-arise/internal/handler/http/category"
	"github.com/sirawong/crud-arise/internal/handler/http/product"
	"github.com/sirawong/crud-arise/pkg/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpServer struct {
	*gin.Engine
}

func NewRouter(productHandler *product.ProductHandler, categoryHandler *category.CategoryHandler) *HttpServer {
	router := gin.New()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		prd := v1.Group("/products")
		{
			prd.POST("/", productHandler.Create)
			prd.GET("/", productHandler.ListAll)
			prd.GET("/:id", productHandler.GetByID)
			prd.PUT("/:id", productHandler.Update)
			prd.DELETE("/:id", productHandler.Delete)
		}
		cate := v1.Group("/categories")
		{
			cate.POST("/", categoryHandler.Create)
			cate.GET("/", categoryHandler.ListAll)
			cate.GET("/:id", categoryHandler.GetByID)
			cate.PUT("/:id", categoryHandler.Update)
			cate.DELETE("/:id", categoryHandler.Delete)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Swagger UI available at http://localhost:8080/swagger/index.html")

	return &HttpServer{router}
}

func (h HttpServer) NewServer(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HttpServerPort),
		Handler: h,
	}
}
