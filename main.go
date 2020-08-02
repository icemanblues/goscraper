package main

import (
	"log"
)

func main() {
	runBeer()
}

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
		score := beerScore(beer)
		log.Printf("%v has a score of %v\n", beer, score)
	}
}
