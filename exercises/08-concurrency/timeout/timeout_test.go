package timeout

import (
	"testing"
	"time"
)

// maxWait is how long a test waits for WithTimeout to return.
const maxWait = 300 * time.Millisecond

func TestWithTimeout(t *testing.T) {
	// never is closed when the test ends, so the slow work never really
	// finishes while WithTimeout is waiting for it.
	never := make(chan struct{})
	t.Cleanup(func() { close(never) })

	tests := []struct {
		name   string
		work   func() int
		limit  time.Duration
		want   int
		wantOK bool
	}{
		{
			name:   "a fast result arrives in time",
			work:   func() int { return 42 },
			limit:  200 * time.Millisecond,
			want:   42,
			wantOK: true,
		},
		{
			name:   "a zero result is still a result",
			work:   func() int { return 0 },
			limit:  200 * time.Millisecond,
			want:   0,
			wantOK: true,
		},
		{
			name:   "slow work runs out of time",
			work:   func() int { <-never; return 7 },
			limit:  20 * time.Millisecond,
			want:   0,
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type answer struct {
				v  int
				ok bool
			}
			done := make(chan answer)
			go func() {
				v, ok := WithTimeout(tt.work, tt.limit)
				done <- answer{v, ok}
			}()
			select {
			case got := <-done:
				if got.v != tt.want || got.ok != tt.wantOK {
					t.Errorf("WithTimeout with limit %v = (%d, %t), want (%d, %t)",
						tt.limit, got.v, got.ok, tt.want, tt.wantOK)
				}
			case <-time.After(maxWait):
				t.Fatalf("WithTimeout with limit %v still waits after %v; does it use time.After(limit) in a select?",
					tt.limit, maxWait)
			}
		})
	}
}
