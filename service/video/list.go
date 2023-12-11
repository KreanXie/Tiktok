package video

import (
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"
	"tiktok/middleware/utils"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	tokenString := c.GetHeader("token")
	user, err := utils.GetUserByToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var videos []common.Video
	if err = db.DB.Where("author_id = ?", user.UserId).Find(&videos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var videoInfos []gin.H
	for _, video := range videos {
		videoInfos = append(videoInfos, gin.H{
			"title":     video.Title,
			"video_id":  video.VideoId,
			"author_id": user.UserId,
			"cover_url": video.CoverUrl,
			"video_url": video.PlayUrl,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"username":  user.Username,
		"signature": user.Signature,
		"videos":    videoInfos,
	})
}
