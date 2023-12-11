package video

import (
	"log"
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"

	"tiktok/middleware/utils"

	"github.com/gin-gonic/gin"
)

func PublishPage(c *gin.Context) {
	c.HTML(http.StatusOK, "publish.html", gin.H{
		"title": "Publish page",
	})
}

func Publish(c *gin.Context) {
	token := c.GetHeader("token")
	user, err := utils.GetUserByToken(token)
	title := c.PostForm("title")
	videoFile, err := c.FormFile("video_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "receive video file fail",
			"error":   err,
		})
		return
	}
	coverFile, err := c.FormFile("cover_file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "receive cover file fail",
			"error":   err,
		})
		return
	}
	playUrl := utils.GetFilePath(videoFile.Filename)
	coverUrl := utils.GetFilePath(coverFile.Filename)
	video := common.Video{
		AuthorId: user.UserId,
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}

	if err := storeVideoInDB(video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "store video fail",
			"error":   err.Error(),
		})
		return
	}
	// Store video file locally
	if err = c.SaveUploadedFile(videoFile, "public/videos/"+videoFile.Filename); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "store video fail",
			"error":   err,
		})
		return
	}
	// Store cover file locally
	if err = c.SaveUploadedFile(coverFile, "public/images/"+coverFile.Filename); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "store cover fail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "publish success",
	})
}

func storeVideoInDB(video common.Video) error {
	return db.DB.Create(&video).Error
}
