package main

import "fmt"

// Second version using sync.WaitGroup:
// https://gist.github.com/harryhare/6a4979aa7f8b90db6cbc74400d0beb49#gistcomment-2934997
func Crawl(url string, depth int, fetcher Fetcher) {
	defer c.wg.Done()

	if depth <= 0 {
		return
	}
	if c.checkvisited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		c.wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}
