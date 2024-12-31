package controller

import (
	"net/http"

	"github.com/DamiaoCanndido/document-api/dto"
	"github.com/DamiaoCanndido/document-api/entities"
	"github.com/DamiaoCanndido/document-api/usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type DocumentController interface {
	GetDocuments(ctx *gin.Context)
	CreateDocuments(ctx *gin.Context)
}

type documentController struct {
	documentUseCase usecases.DocumentUseCase
}

func NewDocumentController(usecase usecases.DocumentUseCase) DocumentController {
	return &documentController{
		documentUseCase: usecase,
	}
}

func (d *documentController) GetDocuments(ctx *gin.Context) {
	docType := ctx.Param("doc")
	documents, err := d.documentUseCase.GetDocuments(docType)

	if (err != nil) {
		ctx.JSON(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, documents)
}

func (d *documentController) CreateDocuments(ctx *gin.Context) {
	var input dto.DocumentCreateDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	document := entities.Document{
		Type: entities.DocumentType(input.Type),
		Order: input.Order,
		Description: input.Description,
		CreatedAt: input.CreatedAt,
	}

	if err := validate.Struct(&document); err != nil {
        validationErrors := err.(validator.ValidationErrors)
        errors := make(map[string]string)
        for _, validationErr := range validationErrors {
            errors[validationErr.Field()] = validationErr.ActualTag()
        }
        ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

	documents, err := d.documentUseCase.CreateDocuments(document)

	if (err != nil) {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusCreated, documents)
}
