package resource

import (
	"github.com/CirillaQL/leakedSearch/model"
	"sync"
)

type Crawler interface {
	GetPageNumber(keyword string) int
	GetVideosList(keyword string, videos chan model.Video, wg *sync.WaitGroup)
}
