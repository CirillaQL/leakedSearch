package porntn

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"

	"github.com/CirillaQL/leakedSearch/model"
	"github.com/gocolly/colly/v2"
)

const porntnBaseUrl = "https://porntn.com/"

type Porntn struct {
}

func (p *Porntn) GetPageNumber(keyword string) int {
	url := fmt.Sprintf("%ssearch/%s/", porntnBaseUrl, keyword)
	c := colly.NewCollector()
	var page int
	c.OnHTML("li[class='last']", func(e *colly.HTMLElement) {
		LastPageNumberString := e.ChildAttr("a", "data-parameters")
		var numberRex = regexp.MustCompile(`\d+`)
		resultOfRex := numberRex.FindAllString(LastPageNumberString, -1)
		realNumber := resultOfRex[len(resultOfRex)-1]
		if realNumber == "" {
			page = 0
		} else {
			result, err := strconv.Atoi(realNumber)
			if err != nil {
				log.Fatalf("Can't get Page Number, Error: %+v", err)
			}
			page = result
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Can't Connect to Porntn, Error: %+v", err)
	}
	return page
}

func (p *Porntn) GetVideosList(keyword string, videos chan model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(videos)
	url := fmt.Sprintf("%ssearch/%s/?mode=async&function=get_block&block_id=list_videos_videos_list_search_result&q=%s&sort_by=&from_videos=%d&from_albums=%d", porntnBaseUrl, keyword, keyword, 1, 1)
	c := colly.NewCollector()
	pageTotal := p.GetPageNumber(keyword)
	page := 2
	c.OnHTML("div[id='list_videos_videos_list_search_result_items']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='item  ']", func(i int, element *colly.HTMLElement) {
			videoUrl := element.ChildAttr("a", "href")
			name := element.ChildText("strong")
			coverImg := element.ChildAttr("div[class='img'] > img", "data-original")
			video := model.Video{
				Name:     name,
				URL:      videoUrl,
				CoverImg: coverImg,
				Source:   "Porntn",
			}
			videos <- video
		})
		if page <= pageTotal {
			nextPageUrl := fmt.Sprintf("%ssearch/%s/?mode=async&function=get_block&block_id=list_videos_videos_list_search_result&q=%s&sort_by=&from_videos=%s&from_albums=%s", porntnBaseUrl, keyword, keyword, strconv.Itoa(page), strconv.Itoa(page))
			page++
			err := c.Visit(nextPageUrl)
			if err != nil {
				log.Fatalf("Can't Connect to Porntn, Error: %+v", err)
			}
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Can't Connect to Porntn, Error: %+v", err)
	}
}
