package resource

import (
	"github.com/CirillaQL/leakedSearch/model"
	"sync"
)

const DefaultPageContentNumber = 20

type Crawler interface {
	GetPageNumber(keyword string) int
	GetVideosList(keyword string, videos chan model.Video, wg *sync.WaitGroup)
	StoreToCache(keyword string)
}
