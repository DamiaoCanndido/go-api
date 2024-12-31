package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentType string

const (
    Decree       	DocumentType = "decree"
    Law          	DocumentType = "law"
    Notice 			DocumentType = "notice"
    Ordinance    	DocumentType = "ordinance"
)

type Document struct {
    ID          uuid.UUID 		`gorm:"type:uuid;primaryKey;" json:"id"`
    Type        DocumentType   	`gorm:"type:varchar(20);not null" json:"type"`
    Order       int            	`gorm:"not null" json:"order"`
    Description string         	`gorm:"not null" json:"description"`
    CreatedAt   time.Time      	`gorm:"not null autoCreateTime" json:"createdAt"`
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
