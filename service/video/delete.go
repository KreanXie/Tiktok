package video

import (
	"fmt"
	"net/http"
	"os"
	"tiktok/common"
	db "tiktok/middleware/database"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	videoId := c.Query("video_id")

	var video common.Video
	if err := db.DB.Where("video_id = ?", videoId).First(&video).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "search video fail",
			"error":   err,
		})
		return
	}

	videoName := parseFileName(video.PlayUrl)
	videoPath := getFilePath(videoName)

	coverName := parseFileName(video.CoverUrl)
	coverPath := getFilePath(coverName)

	if _, err := os.Stat(videoPath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete video file fail",
			"error":   fmt.Sprintf("%v", err),
		})
		return
	}
	if _, err := os.Stat(coverPath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete cover file fail",
			"error":   fmt.Sprintf("%v", err),
		})
		return
	}
	if err := db.DB.Where("video_id = ?", videoId).Delete(&video).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete video record fail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete video success",
	})
}
