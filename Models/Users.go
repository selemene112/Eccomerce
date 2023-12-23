package Models

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (user *User) ValidatePassword(db *gorm.DB) {
	if len(user.Password) < 8 {
		db.Error = errors.New("Password harus memiliki minimal 8 karakter")

		return
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, user.Password); !matched {
		db.AddError(errors.New("Password harus mengandung setidaknya satu huruf besar"))
		return
	}

	if matched, _ := regexp.MatchString(`[!@#$%^&*()_+{}\[\]:;<>,.?~\\/\d]`, user.Password); !matched {
		db.AddError(errors.New("Password harus mengandung setidaknya satu simbol"))
		return
	}
}
