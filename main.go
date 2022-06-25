package main

import (
	"fmt"
	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/resource/spankbang"
)

func main() {
	VideoStream := make(chan model.Video, 50)
	go spankbang.GetVideosList("handjob", VideoStream)
	for {
		video := <-VideoStream
		fmt.Printf("%+v", video)
	}
}
