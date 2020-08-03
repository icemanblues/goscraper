package main

import (
	"log"
	"strings"
)

const BeerURL = "https://www.beeradvocate.com/search/?q="
const BeerSelector = "#ba-content > div:nth-child(3) > div:nth-child(1) > span"

// BeerScore callee must URL encode the beerName properly
func BeerScore(beerName string) string {
	log.Println(BeerURL + beerName)
	element, _ := Scrape(BeerURL+beerName, BeerSelector)
	html, _ := element.Html()
	parts := strings.Split(html, "<br/>")
	endPart := parts[len(parts)-1]
	pipeParts := strings.Split(endPart, "|")
	return pipeParts[0]
}

// a simple demo that this works
func runBeer() {
	beers := []string{
		"voodoo+ranger",
		"sloop+juice+bomb",
		"budweiser",
		"heineken",
		"hoegaarden",
		"miller+lite",
		"rogue+dead+guy",
		"stone+ipa",
	}
	for _, beer := range beers {
		score := BeerScore(beer)
		log.Printf("%v has a score of %v\n", beer, score)
	}
}
