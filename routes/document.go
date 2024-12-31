package routes

import (
	"github.com/DamiaoCanndido/document-api/controller"
	"github.com/DamiaoCanndido/document-api/repository"
	"github.com/DamiaoCanndido/document-api/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    documentRepo := repository.NewDocumentRepository(db)
    documentUseCase := usecases.NewDocumentUseCase(documentRepo)
    documentController := controller.NewDocumentController(documentUseCase)

    prod := router.Group("/document")
    {
        prod.GET("/:doc", documentController.GetDocuments)
        prod.POST("/", documentController.CreateDocuments)
    }

    return router
}