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

	documentRepo := repository.NewDocumentRepository(db)
	documentUseCase := usecases.NewDocumentUseCase(documentRepo)
	documentController := controller.NewDocumentController(documentUseCase)

	prod := router.Group("/document")
	{
		prod.GET("/", documentController.GetDocuments)
		prod.POST("/", documentController.CreateDocuments)
	}

	router.Run(":5000")
}