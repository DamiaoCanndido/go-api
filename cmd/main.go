package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	
	router := gin.Default()

	productRepo := repository.NewProductRepository(db)
	ProductUseCase := usecases.NewProductUseCase(productRepo)
	ProductController := controller.NewProductController(ProductUseCase)

	prod := router.Group("/product")
	{
		prod.GET("/", ProductController.GetProducts)
		prod.POST("/", ProductController.CreateProducts)
	}

	router.Run(":5000")
}