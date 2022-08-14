package server

import (
	"net/http"
	"sync"

	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/utils/cache"
	"github.com/CirillaQL/leakedSearch/utils/logger"
	"github.com/gin-gonic/gin"
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
		videosResult := model.VideoSlice{}
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
		var videosResult model.VideoSlice
		err := videosResult.UnmarshalBinToVideos(videos)
		if err != nil {
			logger.Log().Error("Can't unmarshal binary to videos")
		}
		ctx.JSON(http.StatusOK, videosResult)
	}
}
