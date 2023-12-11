package user

import (
	"errors"
	"net/http"
	"regexp"

	"tiktok/common"
	db "tiktok/middleware/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "register page",
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	err := createUser(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "register success",
		})
	}
}

func createUser(username, password string) error {
	user := common.User{
		Username:  username,
		Password:  password,
		Signature: "This user has no signature.",
	}
	if err := checkPatternOfUsernameAndPassword(username, password); err != nil {
		return err
	}
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != gorm.ErrRecordNotFound {
		return errors.New("Username already exists")
	}
	return db.DB.Create(&user).Error
}

func checkPatternOfUsernameAndPassword(username, password string) error {
	usernamePattern := `^[a-zA-Z0-9_-]{3,20}$`
	passwordPattern := `^[a-zA-Z0-9@#$%^&+=*!_-]{6,20}$`

	validUsername := regexp.MustCompile(usernamePattern).MatchString(username)
	validPassword := regexp.MustCompile(passwordPattern).MatchString(password)
	if !validUsername {
		return errors.New("Invalid pattern of Username")
	}
	if !validPassword {
		return errors.New("Invalid pattern of Password")
	}
	return nil
}
