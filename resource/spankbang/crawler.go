package spankbang

import (
	"fmt"
	"github.com/CirillaQL/leakedSearch/model"
	"github.com/gocolly/colly/v2"
	"log"
	"strconv"
	"time"
)

const spankbangBaseUrl = "https://spankbang.com"

func getPageNumber(keyword string) int {
	url := fmt.Sprintf("%s/s/%s/", spankbangBaseUrl, keyword)
	c := colly.NewCollector()
	var page int
	c.OnHTML("div[class='pagination']", func(e *colly.HTMLElement) {
		n := e.ChildText("li:nth-child(6)")
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
		log.Fatalf("Can't Connect to SpankBang, Error: %+v", err)
	}
	return page
}

func GetVideosList(keyword string, videos chan model.Video) {
	url := fmt.Sprintf("%s/s/%s/", spankbangBaseUrl, keyword)
	c := colly.NewCollector()
	pageTotal := getPageNumber(keyword)
	page := 2
	// First Page
	c.OnHTML("div[class='video-list video-rotate video-list-with-ads']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='video-item ']", func(i int, element *colly.HTMLElement) {
			url := element.ChildAttr("a[class='thumb ']", "href")
			name := element.ChildText("a:nth-child(2)")
			coverImg := element.ChildAttr("a[class='thumb '] > picture > img", "data-src")
			video := model.Video{
				URL:      spankbangBaseUrl + url,
				CoverImg: coverImg,
				Name:     name,
			}
			fmt.Println(video)
			videos <- video
		})
		page++
		if page <= pageTotal {
			time.Sleep(2 * time.Second)
			nextPageUrl := fmt.Sprintf("%s/s/%s/%s/", spankbangBaseUrl, keyword, strconv.Itoa(page-1))
			err := c.Visit(nextPageUrl)
			if err != nil {
				log.Fatalf("Can't Connect to SpankBang, Error: %+v", err)
			}
		}
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Can't Connect to SpankBang, Error: %+v", err)
	}
}
