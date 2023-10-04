package user

import (
	"net/http"
	rd "tiktok/middleware/redis"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	confirm := c.Query("confirm")
	if confirm == "true" {
		token := c.Query("token")
		if err := rd.BlockToken(token, tokenExpirationSec); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "block token fail",
				"error":   err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "logout success",
		})
		c.Redirect(http.StatusFound, "/")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user's not confirmed",
		})
	}
}
