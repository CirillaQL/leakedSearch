package dirtyship

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/CirillaQL/leakedSearch/model"
	"github.com/gocolly/colly/v2"
)

const dirtyshipBaseUrl = "https://dirtyship.com/"

type DirtyShip struct {
}

func (d *DirtyShip) GetPageNumber(keyword string) int {
	url := fmt.Sprintf("%s?search_param=all&s=%s", dirtyshipBaseUrl, keyword)
	c := colly.NewCollector()
	var page int
	c.OnHTML("div[class='pager no-popunder']", func(e *colly.HTMLElement) {
		n := e.ChildText("a:nth-last-child(2)")
		if n == "" {
			page = 0
		} else {
			result, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Can't get Page Number, Error: %+v", err)
			}
			page = result
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Can't Connect to DirtyShip, Error: %+v", err)
	}
	return page
}

func (d *DirtyShip) GetVideosList(keyword string, videos chan model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(videos)
	url := fmt.Sprintf("%s?search_param=all&s=%s", dirtyshipBaseUrl, keyword)
	c := colly.NewCollector()
	pageTotal := d.GetPageNumber(keyword)
	page := 2
	c.OnHTML("ul[class='Thumbnail_List yesPopunder']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			videoUrl := element.ChildAttr("a:nth-child(1)", "href")
			name := element.ChildAttr("a:nth-child(1)", "title")
			coverImg := element.ChildAttr("a:nth-child(1) > img", "src")
			video := model.Video{
				Name:     name,
				URL:      videoUrl,
				CoverImg: coverImg,
				Source:   "DirtyShip",
			}
			videos <- video
		})
		page++
		if (page - 1) <= pageTotal {
			time.Sleep(1 * time.Second)
			nextPageUrl := fmt.Sprintf("%spage/%s/%s", dirtyshipBaseUrl, strconv.Itoa(page-1),
				"?search_param=all&s="+keyword)
			err := c.Visit(nextPageUrl)
			if err != nil {
				log.Fatalf("Can't Connect to DirtyShip, Error: %+v", err)
			}
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Can't Connect to DirtyShip, Error: %+v", err)
	}
}
