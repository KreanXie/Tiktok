package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main page",
	})
}

func Welcome(c *gin.Context) {
}
