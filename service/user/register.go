package user

import (
	"errors"
	"net/http"

	"tiktok/common"
	db "tiktok/middleware/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	username := c.PostForm("username")

	err := createUser(account, password, username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "register fail, error type:" + err.Error(),
		})
		c.Redirect(http.StatusFound, "/register-fail")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "register success",
		})
		c.Redirect(http.StatusFound, "/register-success")
	}
}

func RegisterFail(c *gin.Context) {
	confirmed := c.PostForm("confirmed")

	if confirmed == "true" {
		c.JSON(http.StatusOK, gin.H{
			"message": "user confirmed",
		})
		c.Redirect(http.StatusFound, "/register")
	}
}

func RegisterSuccess(c *gin.Context) {
	confirmed := c.PostForm("confirmed")

	if confirmed == "true" {
		c.JSON(http.StatusOK, gin.H{
			"message": "user confirmed",
		})
		c.Redirect(http.StatusFound, "/")
	}
}

func createUser(account, password, username string) error {
	user := common.User{
		Account:   account,
		Password:  password,
		Username:  username,
		Signature: "This user has no signature.",
	}
	if err := db.DB.Where("account = ?", account).First(&user).Error; err != gorm.ErrRecordNotFound {
		return errors.New("Account already exists")
	} else {
		return db.DB.Create(&user).Error
	}
}
