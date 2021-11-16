package main

import "time"

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(5 * time.Second)
}
