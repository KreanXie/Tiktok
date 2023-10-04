package video

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tiktok/common"
	db "tiktok/middleware/database"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	title := c.PostForm("title")
	authorId := c.PostForm("author_id")
	video, err := c.FormFile("video_file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "receive video file fail",
			"error":   err,
		})
		return
	}
	cover, err := c.FormFile("cover_file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "receive cover file fail",
			"error":   err,
		})
		return
	}

	if err := createVideo(video.Filename, cover.Filename, title, authorId); err != nil {
		log.Printf("%v", err.Error)
		c.JSON(http.StatusOK, gin.H{
			"message": "store video info fail",
			"error":   fmt.Sprintf("%s", err),
		})
		return
	}
	if err = c.SaveUploadedFile(video, getFilePath(video.Filename)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "store video fail",
			"error":   err,
		})
		return
	}
	if err = c.SaveUploadedFile(cover, getFilePath(cover.Filename)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "store cover fail",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "publish success",
	})
}

func createVideo(videoName, coverName, title, authorId string) error {
	id, err := strconv.Atoi(authorId)
	if err != nil {
		return err
	}
	return db.DB.Create(&common.Video{
		AuthorId: uint(id),
		Title:    title,
		PlayUrl:  "http://101.43.169.95:8080" + getFilePath(videoName),
		CoverUrl: "http://101.43.169.95:8080" + getFilePath(coverName),
	}).Error
}
