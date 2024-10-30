package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index" json:"updated_at" `
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" `
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id == "" {
		base.Id = ulid.Make().String()
	}
	return nil
}
