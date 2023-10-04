package router

import (
	"tiktok/service/user"
	"tiktok/service/video"
	service "tiktok/service/welcome"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", service.Welcome)

	userGroup := r.Group("user")
	userGroup.GET("/", user.UserList)
	userGroup.POST("/register", user.Register)
	userGroup.GET("/register-fail", user.RegisterFail)
	userGroup.GET("/register-success", user.RegisterSuccess)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/logout", user.Logout)

	videoGroup := r.Group("video")
	videoGroup.POST("/publish", video.Publish)
	videoGroup.POST("/feed", video.Feed)
	videoGroup.POST("/delete", video.Delete)

	r.Run(":8080")
}
