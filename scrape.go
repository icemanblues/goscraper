package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TODO: randomly generate valid user agents
func myGet(url string) (*http.Response, error) {
	return http.Get(url)
}

// Scrape will load a website and return the element at a js selector
func Scrape(url, selector string) (*goquery.Selection, error) {
	resp, err := myGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Did not receive a 200 OK [%v] from %v\n", resp.StatusCode, url)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return document.Find(selector), nil
}

// WriteFile takes a URL and writes it to file via streaming reader writer
func WriteFile(url, filename string) error {
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

// WriteUrl given a url, it will load and write to the local file system
func WriteUrl(url string) error {
	parts := strings.Split(url, "/")
	return WriteFile(url, parts[len(parts)-1])
}

// FindAttribute given a url, selector, and attribute, it will load the URL, find the element(s)
// it will check each element for this attribute and return it
func FindAttribute(url, selector, attr string) ([]string, error) {
	document, err := Scrape(url, selector)
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

// FindLinks returns all hyperlinks in a given URL
func FindLinks(url string) ([]string, error) {
	return FindAttribute(url, "a", "href")
}

// FindImages returns all images URLs in a given URL
func FindImages(url string) ([]string, error) {
	return FindAttribute(url, "img", "src")
}

// FindVideos returns all video URLs in a given URL
func FindVideos(url string) ([]string, error) {
	return FindAttribute(url, "video", "src")
}
