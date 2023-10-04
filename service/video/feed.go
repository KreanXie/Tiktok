package video

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	filename := c.PostForm("filename")

	_, err := os.Open(storePath + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "open file fail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "openfile success",
		"video_url": "http://101.43.169.95:8080" + storePath + filename,
	})
}
