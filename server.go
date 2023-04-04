package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticrevies/goland-gin-poc/controller"
	"gitlab.com/pragmaticrevies/goland-gin-poc/middlewares"
	"gitlab.com/pragmaticrevies/goland-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setLogOutput() {
	f, _ := os.Create("Log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
