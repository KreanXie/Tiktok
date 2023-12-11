package utils

import (
	"tiktok/common"
	db "tiktok/middleware/database"
)

func GetUserByToken(token string) (common.User, error) {
	var user common.User
	result := db.DB.Where("token = ?", token).First(&user)

	if result.Error != nil {
		return common.User{}, result.Error
	}
	return user, nil
}
