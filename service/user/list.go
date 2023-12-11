package user

import (
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"

	"github.com/gin-gonic/gin"
)

const tokenExpirationSec = 900

// UserList return a response to request
func UserList(c *gin.Context) {
	// Check if token's valid
	/*token := c.Query("token")
	if _, err := jwt.ParseToken(token); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid token",
		})
	} else {
		// If valid, return the result from claims directly
		c.JSON(http.StatusOK, gin.H{
			"message": "valid token",
		})
		return
	}*/

	userId := c.Query("user_id")
	var user common.User
	// Read info from DB, and write to user
	if err := db.DB.Table("users").Select("username, username, signature").Where("user_id = ?", userId).Scan(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username":   user.Username,
			"signature": user.Signature,
		})
	}
}
