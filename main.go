package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TODO: randomly generate valid user agents
func myGet(url string) (*http.Response, error) {
	return http.Get(url)
}

func scrape(url, selector string) (*goquery.Selection, error) {
	resp, err := myGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Not OK: %v\n", resp.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return document.Find(selector), nil
}

func writeFile(url, filename string) error {
	resp, err := myGet(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	return err
}

func writeUrlFile(url string) error {
	parts := strings.Split(url, "/")
	return writeFile(url, parts[len(parts)-1])
}

func FindAttribute(url, selector, attr string) ([]string, error) {
	document, err := scrape(url, selector)
	if err != nil {
		return nil, err
	}

	items := make([]string, 0, document.Length())
	document.Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr(attr)
		if exists {
			items = append(items, href)
		}
	})

	return items, nil
}

func FindLinks(url string) ([]string, error) {
	return FindAttribute(url, "a", "href")
}

func FindImages(url string) ([]string, error) {
	return FindAttribute(url, "img", "src")
}

func FindVideos(url string) ([]string, error) {
	return FindAttribute(url, "video", "src")
}

func main() {
	fmt.Println("Hello World")

	const url string = "https://www.beeradvocate.com"

	images, err := FindImages(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, img := range images {
		log.Println(img)
		writeUrlFile(img)
	}

	links, err := FindLinks(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		log.Println(link)
	}

	videos, err := FindVideos(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, vid := range videos {
		log.Println(vid)
	}
}
