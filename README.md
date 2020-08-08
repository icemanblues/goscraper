# Go Scraper

This is a simple command line utility to scrape web pages.
It started off as a simple side project to learn more about go and its
excellent core libraries.

It aims to use the golang core libraries as much as possible.

## goscraper

At the top level of the project, we store the utilities used in web scraping.

### Scrape.go

Scrape is the simple web scraper. Its written in a straight forward, script-like manner

### Beer.go

Beer is a simple wrapper around scrape to look up the scores of beers

This is an example of how to use the scrape utilities in a practical use

## Commands

In the `cmd` directory there are two buildable CLI that will assist you in scraping

If you are aiming to compile and build this yourself, then these are the go files
you should `go build`

* goscraper
* beerscore

### goscraper

The following flags also require a command line arg of the url to scrape

* links
* images
* videos

```
$ goscraper -links https://www.github.com/icemanblues/goscraper
```

You can also combine one or more of the url based flags.

```
$ goscraper -links -images -videos https://www.github.com/icemanblues/goscraper
```

### beerscore

```
$ goscraper -beer czechvar
2020/08/02 22:32:43 czechvar : 80
```
