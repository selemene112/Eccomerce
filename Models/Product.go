package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Name   string
	Price  float64
	gorm.Model
}
