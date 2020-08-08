package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/icemanblues/goscraper"
)

func main() {
	flag.Parse()
	beer := flag.Arg(0)

	if beer != "" {
		beerName := url.PathEscape(beer)
		score := goscraper.BeerScore(beerName)
		log.Printf("%v : %v\n", beer, score)
	}
}
