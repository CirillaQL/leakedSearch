package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"time"
)

// Video Define SpankBang Video Struct
type Video struct {
	URL      string
	CoverIMG string
	Name     string
}

func main() {
	first := colly.NewCollector()
	c := colly.NewCollector()
	var Videos []Video
	var Number int
	page := 2
	// Numbers of pages
	first.OnHTML("div[class='pagination']", func(e *colly.HTMLElement) {
		numbers := e.ChildText("li:nth-child(6)")
		Number, _ = strconv.Atoi(numbers)
	})
	err := first.Visit("https://spankbang.com/s/handjob/?o=new")
	if err != nil {
		return
	}

	// First Page
	c.OnHTML("div[class='video-list video-rotate video-list-with-ads']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='video-item ']", func(i int, element *colly.HTMLElement) {
			url := element.ChildAttr("a[class='thumb ']", "href")
			name := element.ChildText("a:nth-child(2)")
			coverImg := element.ChildAttr("a[class='thumb '] > picture > img", "data-src")
			video := Video{
				URL:      "https://spankbang.com" + url,
				CoverIMG: coverImg,
				Name:     name,
			}
			Videos = append(Videos, video)
		})
		page++
		if page <= Number {
			fmt.Printf("No Page is %d \n", page)
			fmt.Printf("%+v", Videos)
			time.Sleep(4 * time.Second)
			err := c.Visit("https://spankbang.com/s/hanjob/" + strconv.Itoa(page-1) + "/?o=new")
			if err != nil {
				return
			}
		}
	})

	err = c.Visit("https://spankbang.com/s/handjob/?o=new")
	if err != nil {
		return
	}
	fmt.Println(Videos)
}
