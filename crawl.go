package main

import "fmt"

// First version as presented at:
// https://gist.github.com/harryhare/6a4979aa7f8b90db6cbc74400d0beb49#file-exercise-web-crawler-go
func Crawl(reqUrl string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	if c.checkvisited(reqUrl) {
		return
	}
	body, urls, err := fetcher.Fetch(reqUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", reqUrl, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}
