package main

func main() {
	c.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	c.wg.Wait()
}
