package main

import (
	"log"
	"strings"
)

const BeerURL = "https://www.beeradvocate.com/search/?q="
const BeerSelector = "#ba-content > div:nth-child(3) > div:nth-child(1) > span"

// callee must URL encode the beerName properly
func beerScore(beerName string) string {
	log.Println(BeerURL + beerName)
	element, _ := Scrape(BeerURL+beerName, BeerSelector)
	html, _ := element.Html()
	parts := strings.Split(html, "<br/>")
	endPart := parts[len(parts)-1]
	pipeParts := strings.Split(endPart, "|")
	return pipeParts[0]
}
