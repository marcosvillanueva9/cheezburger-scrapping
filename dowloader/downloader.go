package dowloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Worker(id int, jobs <-chan []string, results chan<- string) {
	for job := range jobs {
		fmt.Printf("worker %d started job %s\n", id, job[1])

		//download image
		fileName := fmt.Sprintf("%s.jpeg", job[1])
		
		err := downloadFile(job[0], fileName)
		if err != nil {
			results<-fmt.Sprintf("%d failed. Error: %s", id, err.Error())
			return
		}

		fmt.Printf("worker %d finished job %s\n", id, job[1])
		results <- fmt.Sprintf("%d ok", id)
	}
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