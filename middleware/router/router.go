package router

import (
	jwt "tiktok/middleware/jwt"
	files "tiktok/service/files"
	"tiktok/service/update"
	user "tiktok/service/user"
	vertify "tiktok/service/vertify"
	video "tiktok/service/video"
	welcome "tiktok/service/welcome"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", welcome.WelcomePage)
	r.POST("/", welcome.Welcome)

	filesGroup := r.Group("files")
	filesGroup.GET("/images/*filename", files.Image)
	filesGroup.GET("/videos/*filename", files.Video)

	vertifyGroup := r.Group("vertify")
	vertifyGroup.GET("/jwt", vertify.VertifyJWT)

	userGroup := r.Group("user").Use(jwt.AuthMiddleware())
	userGroup.GET("/", user.UserList)
	userGroup.GET("/register", user.RegisterPage)
	userGroup.POST("/register", user.Register)
	userGroup.GET("/login", user.LoginPage)
	userGroup.POST("/login", user.Login)
	userGroup.GET("/logout", user.LogoutPage)
	userGroup.POST("/logout", user.Logout)
	userGroup.GET("/space", user.SpacePage)
	userGroup.POST("/space", user.Space)

	videoGroup := r.Group("video").Use(jwt.AuthMiddleware())
	videoGroup.GET("/publish", video.PublishPage)
	videoGroup.POST("/publish", video.Publish)
	videoGroup.DELETE("/delete/*video_id", video.Delete)
	videoGroup.GET("/list", video.List)

	updateGroup := r.Group("update")
	updateGroup.POST("/signature", update.Signature)

	r.Run(":8080")
}
