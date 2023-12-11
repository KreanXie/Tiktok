package service

import (
	"net/http"
	"tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func VertifyJWT(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"isValid": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"isValid": true,
			"user":    claims.Username,
		})
	}
}
