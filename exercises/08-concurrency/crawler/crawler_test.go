package crawler

import (
	"fmt"
	"maps"
	"sync"
	"testing"
	"time"
)

// maxWait is how long a test waits for Crawl to return.
const maxWait = 300 * time.Millisecond

// page is what the fake fetcher knows about one url.
type page struct {
	body string
	urls []string
}

// fakeFetcher serves the pages of the Go tour's example and counts how many
// times each url is fetched.
type fakeFetcher struct {
	pages map[string]page

	mu      sync.Mutex
	fetched map[string]int
}

func newFakeFetcher() *fakeFetcher {
	return &fakeFetcher{
		fetched: map[string]int{},
		pages: map[string]page{
			"https://golang.org/": {
				"The Go Programming Language",
				[]string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
			},
			"https://golang.org/pkg/": {
				"Packages",
				[]string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
			},
			"https://golang.org/pkg/fmt/": {
				"Package fmt",
				[]string{"https://golang.org/", "https://golang.org/pkg/"},
			},
			"https://golang.org/pkg/os/": {
				"Package os",
				[]string{"https://golang.org/", "https://golang.org/pkg/"},
			},
		},
	}
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	f.mu.Lock()
	f.fetched[url]++
	f.mu.Unlock()
	if p, ok := f.pages[url]; ok {
		return p.body, p.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetchCounts returns a copy of how many times each url was fetched.
func (f *fakeFetcher) fetchCounts() map[string]int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return maps.Clone(f.fetched)
}

func TestCrawl(t *testing.T) {
	tests := []struct {
		name  string
		depth int
		want  []string // every url Crawl must fetch, once each
	}{
		{name: "depth 0 fetches nothing", depth: 0, want: nil},
		{name: "depth 1 fetches only the first page", depth: 1, want: []string{
			"https://golang.org/",
		}},
		{name: "depth 2 also fetches its links", depth: 2, want: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		}},
		{name: "depth 4 fetches every page once", depth: 4, want: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetcher := newFakeFetcher()
			finished := make(chan struct{})
			go func() {
				Crawl("https://golang.org/", tt.depth, fetcher)
				close(finished)
			}()
			select {
			case <-finished:
			case <-time.After(maxWait):
				t.Fatalf("Crawl with depth %d does not return", tt.depth)
			}

			want := map[string]int{}
			for _, url := range tt.want {
				want[url] = 1
			}
			got := fetcher.fetchCounts()
			if !maps.Equal(got, want) {
				t.Errorf("Crawl with depth %d fetched (url: times):\n  %v\nwant each of these once:\n  %v\ndoes it wait for every goroutine and remember the visited pages?",
					tt.depth, got, want)
			}
		})
	}
}
