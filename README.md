# Go Scraper

This is a simple command line utility to scrape web pages.

It aims to use the golang core libraries as much as possible.

## Command Line Arguments

Flags

* beer

```
$ go-scraper -beer czechvar
2020/08/02 22:32:43 czechvar : 80
```

The following flags also require a command line arg of the url to scrape

* links
* images
* videos

```
$ go-scraper -links https://www.github.com/icemanblues/go-scraper
```

You can also combine one or more of the url based flags.

```
$ go-scraper -links -images -videos https://www.github.com/icemanblues/go-scraper
```
