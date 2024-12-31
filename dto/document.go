package dto

import (
	"time"
)

type DocumentType string

const (
	Decree    DocumentType = "decree"
	Law       DocumentType = "law"
	Notice    DocumentType = "notice"
	Ordinance DocumentType = "ordinance"
)

type DocumentCreateDTO struct {
	Type        DocumentType 	`json:"type" form:"type" binding:"required,oneof=decree law notice ordinance"`
	Order       int          	`json:"order" form:"order"`
	Description string       	`json:"description" form:"description" binding:"required"`
	CreatedAt   time.Time   	`json:"created_at" form:"created_at"`
}