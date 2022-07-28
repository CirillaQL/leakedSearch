package server

import (
	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/resource/dirtyship"
	"github.com/CirillaQL/leakedSearch/resource/porntn"
	"github.com/CirillaQL/leakedSearch/resource/spankbang"
	"github.com/CirillaQL/leakedSearch/utils/cache"
	"github.com/CirillaQL/leakedSearch/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func GetVideos(ctx *gin.Context) {
	// Check if in Cache
	value := ctx.Param("value")
	videos, err := cache.VideoCache.Get(value)
	if err != nil {
		// No cache
		var spankbangVideosList []model.Video
		var dirtyshipvideosList []model.Video
		var porntnvideosList []model.Video
		dirtyshipVideoStream := make(chan model.Video, 200)
		spankbangVideoStream := make(chan model.Video, 200)
		porntnVideoStream := make(chan model.Video, 200)
		wg := sync.WaitGroup{}
		wg.Add(3)
		go spankbang.GetVideosList(value, spankbangVideoStream, &wg)
		go dirtyship.GetVideosList(value, dirtyshipVideoStream, &wg)
		go porntn.GetVideosList(value, porntnVideoStream, &wg)
		wg.Wait()
		for dirtyshipVideo := range dirtyshipVideoStream {
			dirtyshipvideosList = append(dirtyshipvideosList, dirtyshipVideo)
		}
		for spankbangVideo := range spankbangVideoStream {
			spankbangVideosList = append(spankbangVideosList, spankbangVideo)
		}
		for porntnVideo := range porntnVideoStream {
			porntnvideosList = append(porntnvideosList, porntnVideo)
		}
		videosResult := model.Videos{}
		videosResult.Videos = append(videosResult.Videos, dirtyshipvideosList...)
		videosResult.Videos = append(videosResult.Videos, spankbangVideosList...)
		videosResult.Videos = append(videosResult.Videos, porntnvideosList...)

		videosBinary, err := videosResult.MarshalVideosToBin()
		if err != nil {
			logger.Log().Error("Can't marshal videos list to binary")
		}
		err = cache.VideoCache.Set(value, videosBinary)
		if err != nil {
			logger.Log().Error("Can't put videos into cache")
		}
		ctx.JSON(http.StatusOK, videosResult)
	} else {
		// Cached
		var videosResult model.Videos
		err := videosResult.UnmarshalBinToVideos(videos)
		if err != nil {
			logger.Log().Error("Can't unmarshal binary to videos")
		}
		ctx.JSON(http.StatusOK, videosResult)
	}
}
