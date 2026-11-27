//go:build solution

// Package crawler follows links in parallel without visiting a page twice.
package crawler

import "sync"

// Fetcher downloads a page and returns its body and its links.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl fetches url and the pages it links to, up to depth levels, each
// page once, and waits until every page is done.
func Crawl(url string, depth int, fetcher Fetcher) {
	c := &crawler{fetcher: fetcher, visited: map[string]bool{}}
	c.visit(url, depth)
	c.wg.Wait()
}

// crawler holds what every goroutine of one Crawl shares.
type crawler struct {
	fetcher Fetcher
	wg      sync.WaitGroup
	mu      sync.Mutex
	visited map[string]bool
}

// visit fetches url and crawls each of its links in a new goroutine.
func (c *crawler) visit(url string, depth int) {
	if depth <= 0 || !c.firstVisit(url) {
		return
	}
	_, urls, err := c.fetcher.Fetch(url)
	if err != nil {
		return
	}
	for _, link := range urls {
		c.wg.Add(1)
		go func() {
			defer c.wg.Done()
			c.visit(link, depth-1)
		}()
	}
}

// firstVisit marks url as visited and reports whether nobody had visited it
// before.
func (c *crawler) firstVisit(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[url] {
		return false
	}
	c.visited[url] = true
	return true
}
