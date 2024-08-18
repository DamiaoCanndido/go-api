package entities

import "github.com/go-playground/validator"

type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=100"`
	Price float64 `gorm:"type:decimal(10,2);not null" json:"price" validate:"required,gt=0"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}