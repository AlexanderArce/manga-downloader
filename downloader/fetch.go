package downloader

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/elboletaire/manga-downloader/grabber"
	"github.com/elboletaire/manga-downloader/http"
	"github.com/fatih/color"
)

// File represents a file
type File struct {
	// Data is the file data
	Data []byte
	// Name is the file name
	Name string
}

// Files is a slice of File
type Files []*File

// FetchChapter downloads all the pages of a chapter
// TODO: refactor this function using a RWLock
func FetchChapter(site grabber.Site, chapter grabber.Chapter) (files Files, err error) {
	var wg sync.WaitGroup

	color.Blue("- downloading %s pages...", color.HiBlackString(chapter.GetTitle()))
	guard := make(chan struct{}, site.GetMaxConcurrency().Pages)

	for _, page := range chapter.Pages {
		guard <- struct{}{}
		wg.Add(1)
		go func(page grabber.Page) {
			defer wg.Done()

			filename := fmt.Sprintf("%03d.jpg", page.Number)
			file, err := FetchFile(http.RequestParams{
				URL:     page.URL,
				Referer: site.GetBaseUrl(),
			}, filename)

			if err != nil {
				color.Red("- error downloading page %s", filename)
				return
			}

			files = append(files, file)

			// release guard
			<-guard
		}(page)
	}

	wg.Wait()

	return
}

// FetchFile gets an online file returning a new *File
func FetchFile(params http.RequestParams, filename string) (*File, error) {
	body, err := http.Get(params)
	if err != nil {
		return nil, fmt.Errorf("error downloading file %s: %s", filename, err)
	}

	data := new(bytes.Buffer)
	if _, err = io.Copy(data, body); err != nil {
		return nil, fmt.Errorf("error copying file %s: %s", filename, err)
	}

	return &File{
		Data: data.Bytes(),
		Name: filename,
	}, nil
}
