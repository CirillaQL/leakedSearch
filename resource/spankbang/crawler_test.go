package spankbang

import (
	"fmt"
	"testing"
)

func TestSpankBang_CrawlerVideos(t *testing.T) {
	s, err := NewSpankBang()
	if err != nil {
		panic(err)
	}
	s.CrawlerVideos("maimy")
	fmt.Println(s.videosList)
}
