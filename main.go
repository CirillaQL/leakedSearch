package main

import (
	"fmt"
	"github.com/CirillaQL/leakedSearch/model"
	"github.com/CirillaQL/leakedSearch/resource/dirtyship"
	"github.com/CirillaQL/leakedSearch/resource/spankbang"
)

func main() {
	dirtyshipVideoStream := make(chan model.Video, 50)
	spankbangVideoStream := make(chan model.Video, 50)
	go spankbang.GetVideosList("maimy", spankbangVideoStream)
	go dirtyship.GetVideosList("maimy", dirtyshipVideoStream)
	for {
		dirtyshipVideo := <-dirtyshipVideoStream
		spankbangVideo := <-spankbangVideoStream
		fmt.Printf("%+v \n", dirtyshipVideo)
		fmt.Printf("%+v \n", spankbangVideo)
	}
}
