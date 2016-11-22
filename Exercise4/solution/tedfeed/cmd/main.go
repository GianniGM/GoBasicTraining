package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
<<<<<<< HEAD
=======
	"sync"
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc

	"encoding/xml"
	"errors"

	"io"

<<<<<<< HEAD
	"strings"
	"sync"

	//TODO change me
	"github.com/GianniGM/GoBasicTraining/Exercise4/solution/tedfeed"
=======
	"tedfeed"
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
)

const (
	// Values used to initialize tedfeed home directory
	tf     = "tedfeed"
	videos = "videos"
<<<<<<< HEAD
	thumbs = "thumbnails"
=======
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc

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

<<<<<<< HEAD
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
=======
// download retrieves the file at a given URL and saves it using the title as name
func download(url string, fPath string, title string) error {

	file, err := os.Create(filepath.Join(fPath, title))
	if err != nil {
		// Something went wrong creating video file, terminate
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		// Something went wrong downloading the video, terminate
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
}

func main() {

	// Initializing tedfeed home directory as Exercise 1 was requesting
	home := os.Getenv("HOME")
<<<<<<< HEAD
	dirs := []string{filepath.Join(home, tf, videos), filepath.Join(home, tf, thumbs)}

	// Create video and thumbnails directories if they are missing
	for _, d := range dirs {
		if _, err := os.Stat(d); os.IsNotExist(err) {
			if err := os.MkdirAll(d, 0755); err != nil {
				// Something went wrong initializing the home, terminate
				log.Fatalf("error: %s while creating directory: %s\n", d, err)
			}
=======
	d := filepath.Join(home, tf, videos)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			// Something went wrong downloading the video, terminate
			log.Fatalf("error: %s while creating directory: %s\n", d, err)
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
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

<<<<<<< HEAD
	// Printing the title of the feed as Exercise 2 was requesting
	log.Printf("The title of the feed is: %s\n", fd.Title)

	//exercise 3 + 4 down here

	//exercise 4: add waitGroup for wait all goroutine
	var waitGroup sync.WaitGroup

	m := fd.GetLinksList()
	for t, link := range m {

		//video Title could have unconconventional characters like '?' or ' " '
		title := strings.Replace(t, "\"", "", -1)
		title = strings.Replace(title, "?", "", -1)

		//launching download message
		log.Printf("Downloading %s", title)

		//Exercise 4: incrementing counter of waitGroups
		waitGroup.Add(1)

		//launch download video task
		go download(link, dirs[0], title+".mp4", waitGroup)
	}

	//Exercise 4: Wait for all downloads to complete.
	waitGroup.Wait()
=======
	// Download videos concurrently as requested by Exercise 4
	m := fd.GetLinksList()
	var wg sync.WaitGroup
	for t, link := range m {
		wg.Add(1)
		go func(t, link string) {
			defer wg.Done()
			log.Printf("Downloading %s\n", t)
			err := download(link, d, t+".mp4")
			if err != nil {
				log.Printf("error downloading video: %s\n", err)
			}
		}(t, link)
	}
	wg.Wait()
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
}
