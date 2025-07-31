package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel mendefinisikan field umum dengan UUID sebagai ID
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// BeforeCreate akan meng-set UUID baru sebelum record dibuat
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
