package dto

import (
	"time"
)

type DocumentType string

const (
	Decree    DocumentType = "Decree"
	Law       DocumentType = "Law"
	Notice    DocumentType = "Notice"
	Ordinance DocumentType = "Ordinance"
)

type DocumentCreateDTO struct {
	Type        DocumentType 	`json:"type" form:"type" binding:"required"`
	Order       int          	`json:"order" form:"order"`
	Description string       	`json:"description" form:"description" binding:"required"`
	CreatedAt   time.Time   	`json:"created_at" form:"created_at"`
}