package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marcosvillanueva9/cheezburger-scrapping/dowloader"
	"github.com/marcosvillanueva9/cheezburger-scrapping/scrapper"
)

func main() {

	// Variables Definition
	firstArg := os.Args[1]
	secondArg := os.Args[2]
	linksNum, err := strconv.Atoi(firstArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	threadsNum, err := strconv.Atoi(secondArg)
	if err != nil {
		fmt.Println(err)
		return
	}

	jobs := make(chan []string, linksNum)
	results := make(chan string, linksNum)

	// Run Scrapper
	err, links := scrapper.Run(linksNum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create workers
	for w := 1; w <= threadsNum; w++ {
		go dowloader.Worker(w, jobs, results)
	}

	// Send links to workers
	for i, job := range links {
		jobs <- []string{job,fmt.Sprintf("%d", i+1)}
	}
	close(jobs)

	// Get the results
	for a := 1; a <= linksNum; a++ {
		fmt.Println(<- results)
	}
	
	fmt.Println("Done! :3")
}