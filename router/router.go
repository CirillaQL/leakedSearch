package router

import (
	"github.com/CirillaQL/leakedSearch/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartWebService() {
	g := gin.Default()
	//g.StaticFile("/", "./web/dist")
	g.GET("/videos/:value", server.GetVideos)
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	g.Run(":33333")
}
