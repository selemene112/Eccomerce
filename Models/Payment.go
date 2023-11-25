package Models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	WalletID       uuid.UUID `gorm:"type:uuid;not null"`
	PreviousWallet float64
	AmountPayment  float64
	CurrentWallet  float64
	CreatedAt      time.Time
}
