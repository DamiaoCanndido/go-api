package repository

import (
	"go-api/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]entities.Product, error){
	var product []entities.Product
	pr.connection.Find(&product)
	return product, nil
}

func (pr *ProductRepository) CreateProducts(product entities.Product) (entities.Product, error){
	if err := pr.connection.Create(&product).Error; err != nil {
		return entities.Product{}, err
	}
	return product, nil
}