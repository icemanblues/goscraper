package main

import (
	"flag"
	"log"
)

// define all command line args here
var beerFlag string
var imgFlag bool
var vidFlag bool
var linkFlag bool

func main() {
	flag.StringVar(&beerFlag, "beer", "", "the name of the beer that you want its ranking")
	flag.BoolVar(&imgFlag, "images", false, "find all the images on a web site")
	flag.BoolVar(&vidFlag, "videos", false, "find all the videos on a web site")
	flag.BoolVar(&linkFlag, "links", false, "find all hyperlinks on a web site")
	flag.Parse()

	if beerFlag != "" {
		score := BeerScore(beerFlag)
		log.Printf("%v : %v\n", beerFlag, score)
	}

	// at this point, the flags we support must have a command line arg that is the URL
	url := flag.Arg(0)
	if url == "" {
		return
	}

	if imgFlag {
		images, err := FindImages(url)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("images:")
		for i, img := range images {
			log.Printf("%d : %s\n", i, img)
		}
	}

	if vidFlag {
		vids, err := FindVideos(url)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("videos:")
		for i, v := range vids {
			log.Printf("%d : %s\n", i, v)
		}
	}

	if linkFlag {
		links, err := FindLinks(url)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("links:")
		for i, l := range links {
			log.Printf("%d : %s\n", i, l)
		}
	}
}
