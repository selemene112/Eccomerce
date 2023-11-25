package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"type:varchar(255);not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	Wallet   Wallet
	Products []Product
	Topup    []Topup
	Payment  []Payment
}
