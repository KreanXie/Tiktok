package video

import (
	"net/http"
	"strconv"
	"tiktok/middleware/utils"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	videoId, _ := strconv.Atoi(c.Param("video_id")[1:])
	if err := utils.DeleteVideoById(videoId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete video success",
	})
}
