package update

import (
	"net/http"
	"tiktok/common"
	db "tiktok/middleware/database"
	"tiktok/middleware/utils"

	"github.com/gin-gonic/gin"
)

func Signature(c *gin.Context) {
	// 假设我们使用token来识别用户
	tokenString := c.GetHeader("token")
	user, err := utils.GetUserByToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid token"})
		return
	}

	// 获取新的签名
	var reqBody struct {
		Signature string `json:"signature"`
	}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// 更新数据库中的用户签名
	result := db.DB.Model(&common.User{}).Where("user_id = ?", user.UserId).Update("signature", reqBody.Signature)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update signature"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Signature updated successfully"})
}
