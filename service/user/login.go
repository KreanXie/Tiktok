package user

import (
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"
	jwt "tiktok/middleware/jwt"
	rd "tiktok/middleware/redis"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	token := c.Query("token")
	// Check token
	if ok, err := rd.CheckToken(token); !ok || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid token: blocking",
			"error":   err,
		})
		return
	}

	// Query account and password
	account := c.PostForm("account")
	password := c.PostForm("password")

	// Parse token
	if _, err := jwt.ParseToken(token); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid token: jwt parse fail",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "valid token: jwt pass",
		})
		return
	}

	// Check if password matches the account
	if err := checkAccountAndPassword(account, password); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "login fail, unknown account or wrong password",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "login success, correct account and password",
		})
		token, err = jwt.IssueToken(account)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "issue token fail",
				"error":   err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "issue token success",
			})
		}
	}
}

// Check if given password match the given account, return nil if matches.
func checkAccountAndPassword(account, password string) error {
	user := common.User{
		Account:  account,
		Password: password,
	}
	return db.DB.Where("account = ? and password = ?", account, password).First(&user).Error
}
