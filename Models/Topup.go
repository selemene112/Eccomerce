package Models

import (
	"time"

	"github.com/google/uuid"
)

type Topup struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	WalletID       uuid.UUID `gorm:"type:uuid;not null;foreignKey:ID"`
	PreviousWallet float64
	AmountTopup    float64
	CurrentWallet  float64
	CreatedAt      time.Time
}
