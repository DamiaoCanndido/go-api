package dto

type ProductCreateDTO struct {
	Name  string  `json:"name" form:"name" binding:"required"`
	Price float64 `json:"price" form:"price" binding:"required"`
}