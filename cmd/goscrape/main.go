package main

import (
	"flag"
	"log"
	"net/url"
)

func main() {
	// define all command line args here
	imgFlag := flag.Bool("images", false, "find all the images on a web site")
	vidFlag := flag.Bool("videos", false, "find all the videos on a web site")
	linkFlag := flag.Bool("links", false, "find all hyperlinks on a web site")
	flag.Parse()

	// at this point, the flags we support must have a command line arg that is the URL
	url, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatalf("Unable to parse url %s: %v\n", url, err)
	}

	if *imgFlag {
		images, err := FindImages(url.String())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("images:")
		for i, img := range images {
			log.Printf("%d : %s\n", i, img)
		}
	}

	if *vidFlag {
		vids, err := FindVideos(url.String())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("videos:")
		for i, v := range vids {
			log.Printf("%d : %s\n", i, v)
		}
	}

	if *linkFlag {
		links, err := FindLinks(url.String())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("links:")
		for i, l := range links {
			log.Printf("%d : %s\n", i, l)
		}
	}
}
