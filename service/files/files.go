package files

import (
	"github.com/gin-gonic/gin"
)

func Image(c *gin.Context) {
	filename := c.Param("filename")

	filePath := "public/images" + filename

	c.File(filePath)
}

func Video(c *gin.Context) {
	filename := c.Param("filename")

	filePath := "public/videos" + filename

	c.File(filePath)
}
