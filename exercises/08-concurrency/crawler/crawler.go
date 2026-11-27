//go:build !solution

// Package crawler follows links in parallel without visiting a page twice.
//
// Write this interface and this function in this file:
//
//	type Fetcher interface {
//		Fetch(url string) (body string, urls []string, err error)
//	}
//
//	func Crawl(url string, depth int, fetcher Fetcher)
//
// A Fetcher downloads a page: it returns its body and the links it holds,
// or an error if the page does not exist.
//
// Crawl fetches url and then, each in its own goroutine, crawls the links
// it finds with one less level of depth. With depth 0 or less it fetches
// nothing. Crawl never fetches the same url twice, even when several pages
// link to it, and it returns only when every goroutine has finished.
//
// Example: with depth 1, Crawl fetches only url; with depth 2, it also
// fetches the pages url links to.
//
// Hint: keep the urls you have already visited in a map protected by a
// sync.Mutex, and count the running goroutines with a sync.WaitGroup. A
// page that fails to load is still visited: do not try it again.
package crawler
