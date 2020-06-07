package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/rest-api/controller"
	"github.com/igor-ferreira-almeida/rest-api/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(http.StatusCreated, videoController.Save(context))
	})

	server.Run(":8080")
}
