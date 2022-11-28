package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {

	// Variables Definition
	firstArg := os.Args[1]
	linkNum, err := strconv.Atoi(firstArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run Scrapper
	c := colly.NewCollector(
		colly.AllowedDomains("icanhas.cheezburger.com"),
	)

	var links []string

	c.OnHTML(".mu-post", func(e *colly.HTMLElement) {
		links = append(links, e.ChildAttr("img", "data-src"))
	})


	for i:= 1 ; i < 10 * linkNum ; i++{
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

	links = links[:linkNum]

	for i, url := range links {
		fileName := fmt.Sprintf("%d.jpeg", i + 1)
		downloadFile(url, fileName)
	}

	fmt.Println("Done! :3")
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fmt.Sprintf("./images/%s", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}