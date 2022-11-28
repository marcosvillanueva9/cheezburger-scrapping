package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Run(linkNum int) (error, []string) {

	c := colly.NewCollector(
		colly.AllowedDomains("icanhas.cheezburger.com"),
	)

	var links []string

	c.OnHTML(".mu-post", func(e *colly.HTMLElement) {
		links = append(links, e.ChildAttr("img", "data-src"))
	})


	for i:= 1 ; i < 10 ; i++{
		fmt.Printf("Scraping Page: %d\n", i)
		if i == 1 {
			c.Visit("https://icanhas.cheezburger.com/")
		} else {
			c.Visit(fmt.Sprintf("https://icanhas.cheezburger.com/page/%d", i))
		}
		if len(links) >= linkNum {
			break
		}
	}

	return nil, links[:linkNum]
}