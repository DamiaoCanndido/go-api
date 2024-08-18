package usecases

import (
	"go-api/entities"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (puc *ProductUseCase) GetProducts() ([]entities.Product, error) {
	return puc.repository.GetProducts()
}

func (puc *ProductUseCase) CreateProducts(product entities.Product) (entities.Product, error) {
	product, err := puc.repository.CreateProducts(product)
	if (err != nil) {
		return entities.Product{}, err
	}
	return product, nil
}