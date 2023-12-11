package user

import (
	"log"
	"net/http"
	utils "tiktok/middleware/utils"

	"github.com/gin-gonic/gin"
)

func SpacePage(c *gin.Context) {
	c.HTML(http.StatusOK, "space.html", gin.H{
		"title": "Space page",
	})
}

func Space(c *gin.Context) {
	token := c.GetHeader("token")
	_, err := utils.GetUserByToken(token)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to find user",
			"error":   err.Error,
		})
		return
	}

}
