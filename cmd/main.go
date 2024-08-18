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
	
	server := gin.Default()

	productRepo := repository.NewProductRepository(db)
	ProductUseCase := usecases.NewProductUseCase(productRepo)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/product", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProducts)

	server.Run(":5000")
}