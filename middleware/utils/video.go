package utils

import (
	"errors"
	"os"
	"tiktok/common"
	db "tiktok/middleware/database"
)

func DeleteVideoById(videoId int) error {
	var video common.Video
	if err := db.DB.Model(&common.Video{}).Where("video_id = ?", videoId).First(&video).Error; err != nil {
		return err
	}
	playUrl := video.PlayUrl
	coverUrl := video.CoverUrl

	videoName := ParseFileName(playUrl)
	coverName := ParseFileName(coverUrl)

	if err := os.Remove("public/images/" + coverName); err != nil {
		return errors.New("Delete cover file fail\n" + err.Error())
	}
	if err := os.Remove("public/videos/" + videoName); err != nil {
		return errors.New("Delete cover file fail\n" + err.Error())
	}
	if err := db.DB.Delete(&common.Video{}, videoId).Error; err != nil {
		return errors.New("Delete database record fail\n" + err.Error())
	}
	return nil
}
