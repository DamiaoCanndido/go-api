package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentType string

const (
    Decree       	DocumentType = "Decree"
    Law          	DocumentType = "Law"
    Notice 			DocumentType = "Notice"
    Ordinance    	DocumentType = "Ordinance"
)

type Document struct {
    ID          uuid.UUID 		`gorm:"type:uuid;primaryKey;"`
    Type        DocumentType   	`gorm:"type:varchar(20);not null"`
    Order       int            	`gorm:"not null"`
    Description string         	`gorm:"not null"`
    CreatedAt   time.Time      	`gorm:"not null autoCreateTime"`
}

func (d *Document) BeforeCreate(db *gorm.DB) (err error) {
	d.ID = uuid.New()

    if d.Type != Law {
        var lastOrder Document
        db.Order("created_at desc").Where("type = ? AND EXTRACT(YEAR FROM created_at) = ?", d.Type, time.Now().Year()).First(&lastOrder)
        if d.Order == 0 {
            d.Order = lastOrder.Order + 1
        }
    } else {
        var lastOrder Document
        db.Order("created_at desc").Where("type = ?", d.Type).First(&lastOrder)
        if d.Order == 0 {
            d.Order = lastOrder.Order + 1
        }
    }
    return
}
