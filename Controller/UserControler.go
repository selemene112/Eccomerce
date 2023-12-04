package Controller

import (
	"Transaksi/Config"
	"Transaksi/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
This Code Controller For Make API USER
Several Function For Controller Tabel User On Database
Several Fitur :
CURD Restfull API
-- Create User
1. Make Username,Password,Email for registrasi
2. after registrasi Auto Generate Wallet use ID User
3. Password harus use hashing password
4. Validate Password
5. Validate Email dont registrasi use same email for 2 account
*/
func CreateUser(c *gin.Context) {
	user := Models.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	// fmt.Println("ini user passowrd : %+v\n", user.Password)
	// fmt.Println("ini user passowrd : %+v\n", &user)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ValidatePassword(Config.DB)
	if err := Config.DB.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := Config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})

}

func GetallUser(c *gin.Context) {
	user := []Models.User{}

	if err := Config.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}
