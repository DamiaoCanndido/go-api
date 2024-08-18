package controller

import (
	"go-api/entities"
	"go-api/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := product.Validate(); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	products, _ := p.productUseCase.CreateProducts(product)

	ctx.JSON(http.StatusCreated, products)
}
