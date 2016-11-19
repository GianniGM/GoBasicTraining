package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"encoding/xml"
	"errors"
	"tedfeed"

	"io"

	"strings"
	"sync"
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"
	thumbs = "thumbnails"

	// TED.com atom feed URL
	url = "https://www.ted.com/talks/atom"
)

// parse receive the atom feed, unmarshals it into a Feed instance
// and returns it.
func parse(body []byte) (*tedfeed.Feed, error) {
	var f tedfeed.Feed
	err := xml.Unmarshal(body, &f)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s:error parsing the feed", err))
	}
	return &f, nil
}

func download(url string, fPath string, title string, waitGroup sync.WaitGroup) {

	// Decrement the counter when the goroutine completes.
	defer waitGroup.Done()

	//creating video.file
	file, err := os.Create(filepath.Join(fPath, title))
	if err != nil {
		// Something went wrong creating video file, terminate
		log.Fatalf("%s\n", err)
	}

	//GET the video
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		// Something went wrong downloading the video, terminate
		log.Fatalf("%s\n", err)
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatalf("error: %s while downloading video: %s\n", title, err)
	} else {
		log.Printf("Downloaded file: %s\n", title)
	}

	file.Close()
}

func main() {

	// Initializing tedfeed home directory as Exercise 1 was requesting
	home := os.Getenv("HOME")
	dirs := []string{filepath.Join(home, tf, videos), filepath.Join(home, tf, thumbs)}

	// Create video and thumbnails directories if they are missing
	for _, d := range dirs {
		if _, err := os.Stat(d); os.IsNotExist(err) {
			if err := os.MkdirAll(d, 0755); err != nil {
				// Something went wrong initializing the home, terminate
				log.Fatalf("error: %s while creating directory: %s\n", d, err)
			}
		}
	}

	//GET the atom feed
	resp, err := http.Get(url)
	if err != nil {
		// Something went wrong reading the feed, terminate
		log.Fatalf("%s\n", err)
	}
	defer resp.Body.Close()

	var output []byte
	if output, err = ioutil.ReadAll(resp.Body); err != nil {
		// Something went wrong reading the request body, terminate
		log.Fatalf("%s\n", err)
	}

	fd, err := parse(output)
	if err != nil {
		log.Fatalln("error parsing the feed")
	}

	// Printing the title of the feed as Exercise 2 was reqesting
	log.Printf("The title of the feed is: %s\n", fd.Title)

	//exercise 4
	var waitGroup sync.WaitGroup

	//iterate over tedfeed.Entry[].Link[]
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {

			//must get only Rel == "enclosure" link
			if link.Rel == "enclosure" {

				//launching download task
				log.Printf("Downloading %s", entry.Title)

				videoName := string(entry.Title)

				//videoName could have unconconventional characters
				videoName = strings.Replace(videoName, "\"", "", -1)
				videoName = strings.Replace(videoName, "?", "", -1)

				// Increment the WaitGroup counter.
				waitGroup.Add(1)

				//download video
				go download(link.HRef, dirs[0], videoName+".mp4", waitGroup)
			}
		}
	}

	// Wait for all downloads to complete.
	waitGroup.Wait()
}
