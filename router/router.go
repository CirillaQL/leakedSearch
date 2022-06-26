package router

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/resource/dirtyship"
	"github.com/CirillaQL/leakedSearch/resource/spankbang"
	"github.com/gin-gonic/gin"
)

type Videos struct {
	Spankbang []model.Video
	Dirtyship []model.Video
}

func StartWebService() {
	g := gin.Default()
	var spankbangVideosList []model.Video
	var dirtyshipvideosList []model.Video
	g.GET("/videos", func(ctx *gin.Context) {
		dirtyshipVideoStream := make(chan model.Video, 100)
		spankbangVideoStream := make(chan model.Video, 100)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go spankbang.GetVideosList("maimy", spankbangVideoStream, &wg)
		go dirtyship.GetVideosList("maimy", dirtyshipVideoStream, &wg)
		wg.Wait()
		fmt.Println("read over")
		for dirtyshipVideo := range dirtyshipVideoStream {
			dirtyshipvideosList = append(dirtyshipvideosList, dirtyshipVideo)
		}
		fmt.Println("dirtyship read over")
		for spankbangVideo := range spankbangVideoStream {
			spankbangVideosList = append(spankbangVideosList, spankbangVideo)
		}
		videos := Videos{}
		videos.Dirtyship = dirtyshipvideosList
		videos.Spankbang = spankbangVideosList
		ctx.JSON(http.StatusOK, videos)
	})
	g.Run(":33333")
}
