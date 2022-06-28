package porntn

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

const porntnBaseUrl = "https://porntn.com/"

func GetPageNumber(keyword string) int {
	url := fmt.Sprintf("%ssearch/%s/", porntnBaseUrl, keyword)
	fmt.Println(url)
	c := colly.NewCollector()
	var page int
	c.OnHTML("li[class='last']", func(e *colly.HTMLElement) {
		LastPageNumberString := e.ChildAttr("a", "data-parameters")
		fmt.Println(LastPageNumberString)
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
