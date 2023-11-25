package models

import "github.com/google/uuid"

type Wallet struct {
	UserID        uuid.UUID `gorm:"type:uuid;unique;not null"`
	Balance       float64
	Topup         []Topup
	Payment       []Payment
	WalletHistori []WalletHistori
}

type WalletHistori struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	WalletID        uuid.UUID `gorm:"type:uuid;not null"`
	UserId          uuid.UUID `gorm:"type:uuid;not null"`
	TraksactionType string    `gorm:"type:varchar(255);not null"`
	PreviousWallet  float64
	AmountPayment   float64
	CurrentWallet   float64
	CreatedAt       string `gorm:"type:varchar(255);not null"`
}
