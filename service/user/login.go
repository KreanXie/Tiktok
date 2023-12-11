package user

import (
	"log"
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"
	"tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login page",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isAuthenticated, ok := c.Get("isAuthenticated"); ok && isAuthenticated.(bool) {
		c.JSON(http.StatusOK, gin.H{
			"message": "token passed",
		})
		return
	}
	// Check if password matches the username
	if err := checkUsernameAndPassword(username, password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login fail, unknown username or wrong password",
		})
	} else {
		tokenString, err := jwt.IssueToken(username)
		if err != nil {
			log.Println("Fail to issue token: ", err)
			tokenString = "Wrong token"
		}
		updateToken(username, tokenString)
		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
			"token":   tokenString,
		})
	}
}

// Check if given password match the given username, return nil if matches.
func checkUsernameAndPassword(username, password string) error {
	user := common.User{
		Username: username,
		Password: password,
	}
	return db.DB.Where("username = ? and password = ?", username, password).First(&user).Error
}

func updateToken(username, token string) error {
	return db.DB.Model(&common.User{}).Where("username = ?", username).Update("token", token).Error
}
