package safecounter

import (
	"sync"
	"testing"
)

func TestNewCounterStartsAtZero(t *testing.T) {
	c := NewSafeCounter()
	if got := c.Value("home"); got != 0 {
		t.Errorf("NewSafeCounter().Value(\"home\") = %d, want 0", got)
	}
}

func TestSafeCounter(t *testing.T) {
	tests := []struct {
		name       string
		goroutines int
		visits     int // calls to Inc made by each goroutine
		keys       []string
	}{
		{name: "one goroutine", goroutines: 1, visits: 5, keys: []string{"home"}},
		{name: "many goroutines on one key", goroutines: 100, visits: 10, keys: []string{"home"}},
		{name: "many goroutines on several keys", goroutines: 50, visits: 20, keys: []string{"home", "blog", "about"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSafeCounter()
			var wg sync.WaitGroup
			for range tt.goroutines {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for range tt.visits {
						for _, key := range tt.keys {
							c.Inc(key)
							c.Value(key)
						}
					}
				}()
			}
			wg.Wait()

			want := tt.goroutines * tt.visits
			for _, key := range tt.keys {
				if got := c.Value(key); got != want {
					t.Errorf("%d goroutines with %d visits each: Value(%q) = %d, want %d; does Inc lock the Mutex?",
						tt.goroutines, tt.visits, key, got, want)
				}
			}
		})
	}
}
