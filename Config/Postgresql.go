package Config

import (
	"Transaksi/Models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "admin"
	DB_PASSWORD = "lele123"
	DB_NAME     = "eccomerce"
	DB_PORT     = 5432

	DB  *gorm.DB
	err error
)

func StartPostgresqlDB() *gorm.DB {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	dsn := config
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Models.User{}, &Models.Product{}, &Models.Payment{}, &Models.Topup{}, &Models.Wallet{}, &Models.WalletHistori{})

	fmt.Println("Connection Opened to Database")

	return DB

}
