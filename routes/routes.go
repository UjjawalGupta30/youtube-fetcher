package routes

import (
	"github.com/gin-gonic/gin"
	"youtube-fetcher/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/videos", controllers.GetVideos)
}
