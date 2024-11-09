package controller

import (
	"go-api/dto"
	"go-api/entities"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ProductController interface {
	GetProducts(ctx *gin.Context)
}

type productController struct {
	productUseCase usecases.ProductUseCase
}

func NewProductController(usecase usecases.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()

	if (err != nil) {
		ctx.JSON(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProducts(ctx *gin.Context) {
	var input dto.ProductCreateDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	product := entities.Product{
		Name: input.Name,
		Price: input.Price,
	}

	if err := validate.Struct(&product); err != nil {
        validationErrors := err.(validator.ValidationErrors)
        errors := make(map[string]string)
        for _, validationErr := range validationErrors {
            errors[validationErr.Field()] = validationErr.ActualTag()
        }
        ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
        return
    }

	products, _ := p.productUseCase.CreateProducts(product)

	ctx.JSON(http.StatusCreated, products)
}
