package router

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/resource/dirtyship"
	"github.com/CirillaQL/leakedSearch/resource/porntn"
	"github.com/CirillaQL/leakedSearch/resource/spankbang"
	"github.com/CirillaQL/leakedSearch/utils/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

type Videos struct {
	Videos []model.Video
}

func (v *Videos) MarshalVideosToBin() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func (v *Videos) UnmarshalBinToVideos(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	return err
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			//Access-Control-Allow-Origin是必须的,他的值要么是请求Origin字段的值,要么是一个*, 表示接受任意域名的请求
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			//该字段可选。CORS请求时，XMLHttpRequest对象的getResponseHeader()方法只能拿到6个基本字段：Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma。
			//如果想拿到其他字段，就必须在Access-Control-Expose-Headers里面指定。上面的例子指定，getResponseHeader('FooBar')可以返回FooBar字段的值。
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			//该字段可选，用来指定本次预检请求的有效期，单位为秒。有效期是20天（1728000秒），即允许缓存该条回应1728000秒（即20天），在此期间，不用发出另一条预检请求。
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func StartWebService() {
	g := gin.Default()
	g.Use(Cors())
	g.GET("/videos/:value", func(ctx *gin.Context) {
		// Check if in Cache
		value := ctx.Param("value")
		videoCache := cache.Cache
		videos, err := videoCache.Get(value)

		if err != nil || videos == nil {
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
			videosResult := Videos{}
			videosResult.Videos = append(videosResult.Videos, dirtyshipvideosList...)
			videosResult.Videos = append(videosResult.Videos, spankbangVideosList...)
			videosResult.Videos = append(videosResult.Videos, porntnvideosList...)

			videosBinary, err := videosResult.MarshalVideosToBin()
			if err != nil {
				fmt.Println(err)
			}
			videoCache.Set(value, videosBinary)
			ctx.JSON(http.StatusOK, videosResult)
		} else {
			// Cached
			var videosResult Videos
			err := videosResult.UnmarshalBinToVideos(videos)
			if err != nil {
				fmt.Println(err)
			}
			ctx.JSON(http.StatusOK, videosResult)
		}
	})
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	g.Run(":33333")
}
