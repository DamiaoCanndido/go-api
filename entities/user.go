package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
    Username    string    `gorm:"type:varchar(50);unique;not null"`
    Email       string    `gorm:"type:varchar(100);unique;not null"`
    Password    string    `gorm:"not null"`
    IsSuperUser bool      `gorm:"default:false"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.New()
    return
}
